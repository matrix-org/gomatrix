package gomatrix

type EventType string

const (
	AliasesEventType           EventType = "m.room.aliases"
	CanonicalAliasEventType    EventType = "m.room.canonical_alias"
	CreateEventType            EventType = "m.room.create"
	JoinRulesEventType         EventType = "m.room.join_rules"
	MemberEventType            EventType = "m.room.member"
	PowerLevelsEventType       EventType = "m.room.power_levels"
	RedactionEventType         EventType = "m.room.redaction"
	MessageEventType           EventType = "m.room.message"
	MessageFeedbackEventType   EventType = "m.room.message.feedback"
	NameEventType              EventType = "m.room.name"
	TopicEventType             EventType = "m.room.topic"
	AvatarEventType            EventType = "m.room.avatar"
	TypingEventType            EventType = "m.typing"
	ReceiptEventType           EventType = "m.receipt"
	PresenceEventType          EventType = "m.presence"
	HistoryVisibilityEventType EventType = "m.room.history_visibility"
	ThirdPartyInviteEventType  EventType = "m.room.third_party_invite"
	GuestAccessEventType       EventType = "m.room.guest_access"
	TagEventType               EventType = "m.tag"
)

func (e EventType) String() string {
	return string(e)
}

func (e EventType) KindOf(target EventType) bool {
	return e == target
}
