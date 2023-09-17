package handler

import (
	"common"
	"context"
	"errors"
	"github.com/plutov/paypal/v3"
	"payment/proto/payment"
	paymentApi "paymentApi/proto/paymentApi"
	"strconv"
)

type PaymentApi struct {
	PaymentService payment.PaymentService
}

var (
	clientID string = "AeSKLbrz0KqsTIZAnhptlSU6qFYHlw-MIMBURCkStgzFIP1CmKaTN0vhSnXTEYukswQFlyY2hivZLiGa"
)

// PaymentApi.PayPalRefund 通过API向外暴露为/paymentApi/payPalRefund,接收http请求
// 即: /paymentApi/payPalrefund.micro.api.paymentApi 服务的PaymentApi.PayPalRefund
func (e *PaymentApi) PayPalRefund(ctx context.Context, req *paymentApi.Request, resp *paymentApi.Response) error {
	if err := isOK("payment_id", req); err != nil {
		resp.StatusCode = 500
		return err
	}
	//退款号
	if err := isOK("refund_id", req); err != nil {
		resp.StatusCode = 500
		return err
	}
	//验证退款金额
	if err := isOK("money", req); err != nil {
		resp.StatusCode = 500
		return err
	}

	payID, err := strconv.ParseInt(req.Get["payment_id"].Values[0], 10, 64)
	if err != nil {
		common.Error(err)
		return err
	}

	paymentInfo, err := e.PaymentService.FindPaymentByID(ctx, &payment.PaymentID{PaymentId: payID})
	if err != nil {
		common.Error(err)
		return err
	}
	//SID获取paymentInfo.PaymentSid
	//支付模式
	status := paypal.APIBaseSandBox
	if paymentInfo.PaymentStatus {
		status = paypal.APIBaseLive
	}

	//退款例子
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			EmailSubject:  req.Get["refund_id"].Values[0] + "xzx提醒你收款",
			EmailMessage:  req.Get["refund_id"].Values[0] + "您有一条收款信息",
			SenderBatchID: req.Get["refund_id"].Values[0],
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType: "EMALL",
				Receiver:      "sb-ztmat27344569@business.example.com",
				Amount: &paypal.AmountPayout{
					//币种
					Currency: "USD",
					Value:    req.Get["money"].Values[0],
				},
				Note:         req.Get["refund_id"].Values[0],
				SenderItemID: req.Get["refund_id"].Values[0],
			},
		},
	}
	//创建支付客户端
	paypalClient, err := paypal.NewClient(clientID, paymentInfo.PaymentSid, status)
	if err != nil {
		common.Error(err)
	}
	//获取token
	_, err = paypalClient.GetAccessToken()
	if err != nil {
		common.Error(err)
	}
	paymentResult, err := paypalClient.CreateSinglePayout(payout)
	if err != nil {
		common.Error(err)
	}
	common.Info(paymentResult)
	resp.Body = req.Get["refund_id"].Values[0] + "支付成功!"
	return err
}

func isOK(key string, req *paymentApi.Request) error {
	if _, ok := req.Get["key"]; !ok {
		err := errors.New(key + "参数异常")
		common.Error(err)
		return err
	}
	return nil

}
