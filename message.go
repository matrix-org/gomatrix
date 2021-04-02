package gomatrix

type MessageType string

const (
	TextMessageType     MessageType = "m.text"
	EmoteMessageType    MessageType = "m.emote"
	NoticeMessageType   MessageType = "m.notice"
	ImageMessageType    MessageType = "m.image"
	FileMessageType     MessageType = "m.file"
	AudioMessageType    MessageType = "m.audio"
	LocationMessageType MessageType = "m.location"
	VideoMessageType    MessageType = "m.video"
)
