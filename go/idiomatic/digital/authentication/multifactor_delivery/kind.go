package multifactor_delivery

type Kind string

type kinds struct {
	Unknown                  Kind
	AuthenticatorApplication Kind
	ShortMessageService      Kind
	Voice                    Kind
	Email                    Kind
	PushNotification         Kind
}

var Kinds = kinds{
	Unknown:                  "unknown",
	AuthenticatorApplication: "authenticator-application",
	ShortMessageService:      "short-message-service",
	Voice:                    "voice",
	Email:                    "email",
	PushNotification:         "push-notification",
}
