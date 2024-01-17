package verificationcode

import (
	"hotelfortuna/common/config"

	"github.com/Dparty/common/sms"
	"github.com/universalmacro/common/sms/models"
	tencentSms "github.com/universalmacro/common/sms/tencent"
)

// const chainTemplateId = "914223"
// const internationTemplateId = "914222"

// var sendCloud *sms.SendCloud = sms.NewSendCloud(config.GetString("sendCloud.user"), config.GetString("sendCloud.key"))
var tencentSmsSender = tencentSms.NewSmsSender(config.GetString("tencentCloud.secretId"), config.GetString("tencentCloud.secretKey"))

func SendVerificationCode(to sms.PhoneNumber, code string) {
	config := tencentSms.Config{
		AppId:      "2400001857",
		SignName:   "財神酒店",
		TemplateId: "2910166",
	}
	vars := []string{code}
	err := tencentSmsSender.SendWithConfig(models.PhoneNumber{AreaCode: to.AreaCode, Number: to.Number}, config, vars)
	if err != nil {
		println(err)
	}
}

// func getTemplateIdByAreaCode(areaCode string) string {
// 	switch areaCode {
// 	case "86":
// 		return chainTemplateId
// 	case "853", "852":
// 		return internationTemplateId
// 	default:
// 		return ""
// 	}
// }
