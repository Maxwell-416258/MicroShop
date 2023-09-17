package main

import (
	"common"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	consul2 "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/opentracing/opentracing-go"
	"log"
	"net"
	"net/http"
	"payment/proto/payment"
	"paymentApi/handler"
	"paymentApi/proto/paymentApi"
)

func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//jaeger 链路追踪
	t, io, err := common.NewTracer("go.micro.api.payment",
		"localhost:6831")
	if err != nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动监听
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9192"),
			hystrixStreamHandler)
	}()

	//监控
	common.PrometheusBoot(9292)

	//New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:9092"),
		micro.Registry(consul),
		//
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//作为服务端时生效
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	//Initialise service
	service.Init()

	paymentService := payment.NewPaymentService("go.micro.service.payment", service.Client())

	//Register Handler

	err = paymentApi.RegisterPaymentApiHandler(service.Server(), &handler.PaymentApi{PaymentService: paymentService})
	if err != nil {
		common.Error(err)
	}
	//Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//正常执行
		common.Info(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, resp, opts...)
	}, func(e error) error {
		common.Error(e)
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
