package messenger

type Kind string

type kinds struct {
	Unknown                         Kind
	AmeraiconOnlineInstantMessaging Kind
	AppleIMessage                   Kind
	// github.com/boundedinfinity/canonical-model/go/idiomatic                              Kind
	BlackBerryMessenger                 Kind
	Discord                             Kind
	ExtensibleMessagingPresenceProtocol Kind
	FacebookMessenger                   Kind
	GoogleMessages                      Kind
	GoogleTalk                          Kind
	Icq                                 Kind
	InternetRelayChat                   Kind
	Jabber                              Kind
	KakaoTalk                           Kind
	Line                                Kind
	MicrosoftTeams                      Kind
	MyspaceInstantMessaging             Kind
	Signal                              Kind
	Slack                               Kind
	Telegram                            Kind
	Threema                             Kind
	WeChat                              Kind
	Viber                               Kind
	WhatsApp                            Kind
	WindowsMessenger                    Kind
	Zalo                                Kind
}

var Kinds = kinds{
	Unknown:                         "digital.messenger.unknown",
	AmeraiconOnlineInstantMessaging: "digital.messenger.america-online-instant-messaging",
	AppleIMessage:                   "digital.messenger.apple-imessage",
	// github.com/boundedinfinity/canonical-model/go/idiomatic:                              "digital.messenger.github.com/boundedinfinity/canonical-model/go/idiomatic",
	BlackBerryMessenger:                 "digital.messenger.blackberry-messenger",
	Discord:                             "digital.messenger.discord",
	ExtensibleMessagingPresenceProtocol: "digital.messenger.extensible-messaging-presence-protocol",
	FacebookMessenger:                   "digital.messenger.facebook-messenger",
	GoogleMessages:                      "digital.messenger.google-messages",
	GoogleTalk:                          "digital.messenger.google-talk",
	Icq:                                 "digital.messenger.icq",
	InternetRelayChat:                   "digital.messenger.internet-relay-chat",
	Jabber:                              "digital.messenger.jabber",
	KakaoTalk:                           "digital.messenger.kakao-talk",
	Line:                                "digital.messenger.line",
	MicrosoftTeams:                      "digital.messenger.microsoft-teams",
	MyspaceInstantMessaging:             "digital.messenger.myspace-instant-messaging",
	Signal:                              "digital.messenger.signal",
	Slack:                               "digital.messenger.slack",
	Telegram:                            "digital.messenger.telegram",
	Threema:                             "digital.messenger.threema",
	WeChat:                              "digital.messenger.wechat",
	Viber:                               "digital.messenger.viber",
	WhatsApp:                            "digital.messenger.whatsapp",
	WindowsMessenger:                    "digital.messenger.windows-messenger",
	Zalo:                                "digital.messenger.zalo",
}
