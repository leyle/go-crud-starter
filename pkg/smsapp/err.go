package smsapp

import "errors"

var (
	ErrNotSupportedCountry = errors.New("not supported country or area")
	ErrRequestTooFast      = errors.New("request too fast")
	ErrSendSMSFailed       = errors.New("send sms failed")
	ErrInternalError       = errors.New("internal error")
	ErrInvalidCode         = errors.New("invalid code")
)
