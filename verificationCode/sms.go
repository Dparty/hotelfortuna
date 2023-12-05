package verificationcode

import (
	"hotelfortuna/common/config"

	"github.com/Dparty/common/sms"
)

const chainTemplateId = "914223"
const internationTemplateId = "914222"

var sendCloud *sms.SendCloud = sms.NewSendCloud(config.GetString("sendCloud.user"), config.GetString("sendCloud.key"))

func SendVerificationCode(to sms.PhoneNumber, code string) {
	sendCloud.SendWithTemplate(to, getTemplateIdByAreaCode(to.AreaCode), map[string]string{"code": code})
}

func getTemplateIdByAreaCode(areaCode string) string {
	switch areaCode {
	case "86":
		return chainTemplateId
	case "853", "852":
		return internationTemplateId
	default:
		return ""
	}
}
