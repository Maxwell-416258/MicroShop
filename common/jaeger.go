package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// 创建链路追踪实例
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, //恒定采样器，会对所有事件进行追踪
			Param: 1,                       //采样率为1
		},
		Reporter: &config.ReporterConfig{ //设置报告器配置
			BufferFlushInterval: 1 * time.Second, //缓冲刷新间隔
			LogSpans:            true,            //是否记录跟踪日志
			LocalAgentHostPort:  addr,            //jaeger的代理地址
		},
	}
	return cfg.NewTracer() //io.Closer用于关闭追踪器
}
