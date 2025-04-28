package pkg

import (
	"fmt"
	"github.com/cheng1005/XxCh-Common/common/global"
	"github.com/smartwalle/alipay/v3"
)

func AliPay(name, number string, totalPrice string) string {

	var (
		privateKey = global.Nacos.Alipay.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
		appId      = global.Nacos.Alipay.AppId
	)
	client, err := alipay.New(appId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = global.Nacos.Alipay.NotifyURL
	p.ReturnURL = global.Nacos.Alipay.ReturnURL
	p.Subject = name
	p.OutTradeNo = number
	p.TotalAmount = totalPrice
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
