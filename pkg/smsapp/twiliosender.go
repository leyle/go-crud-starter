package smsapp

import "github.com/leyle/go-crud-starter/configandcontext"

func sendTwilioSMS(ctx *configandcontext.APIContext, info *SMSRequestForm) error {
	if ctx.Cfg.SMS.Debug {
		ctx.Logger.Warn().Msg("current is running on debug mode, no real sms will be sent")
		return nil
	}

	return nil
}
