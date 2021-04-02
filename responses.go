package gomatrix

// RespCreateFilter is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-user-userid-filter
//easyjson:json
type RespCreateFilter struct {
	FilterID string `json:"filter_id"`
}

// RespVersions is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-versions
//easyjson:json
type RespVersions struct {
	Versions []string `json:"versions"`
}

// RespPublicRooms is the JSON response for http://matrix.org/speculator/spec/HEAD/client_server/unstable.html#get-matrix-client-unstable-publicrooms
//easyjson:json
type RespPublicRooms struct {
	TotalRoomCountEstimate int          `json:"total_room_count_estimate"`
	PrevBatch              string       `json:"prev_batch"`
	NextBatch              string       `json:"next_batch"`
	Chunk                  []PublicRoom `json:"chunk"`
}

// RespJoinRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-join
//easyjson:json
type RespJoinRoom struct {
	RoomID string `json:"room_id"`
}

// RespLeaveRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-leave
//easyjson:json
type RespLeaveRoom struct{}

// RespForgetRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-forget
//easyjson:json
type RespForgetRoom struct{}

// RespInviteUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-invite
//easyjson:json
type RespInviteUser struct{}

// RespKickUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-kick
//easyjson:json
type RespKickUser struct{}

// RespBanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-ban
//easyjson:json
type RespBanUser struct{}

// RespUnbanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-unban
//easyjson:json
type RespUnbanUser struct{}

// RespTyping is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-typing-userid
//easyjson:json
type RespTyping struct{}

// RespJoinedRooms is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
//easyjson:json
type RespJoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// RespJoinedMembers is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
//easyjson:json
type RespJoinedMembers struct {
	Joined map[string]struct {
		DisplayName *string `json:"display_name"`
		AvatarURL   *string `json:"avatar_url"`
	} `json:"joined"`
}

// RespMessages is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-rooms-roomid-messages
//easyjson:json
type RespMessages struct {
	Start string  `json:"start"`
	Chunk []Event `json:"chunk"`
	End   string  `json:"end"`
}

// RespSendEvent is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-send-eventtype-txnid
//easyjson:json
type RespSendEvent struct {
	EventID string `json:"event_id"`
}

// RespMediaUpload is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-media-r0-upload
//easyjson:json
type RespMediaUpload struct {
	ContentURI string `json:"content_uri"`
}

// RespUserInteractive is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#user-interactive-authentication-api
//easyjson:json
type RespUserInteractive struct {
	Flows []struct {
		Stages []string `json:"stages"`
	} `json:"flows"`
	Params    map[string]interface{} `json:"params"`
	Session   string                 `json:"session"`
	Completed []string               `json:"completed"`
	ErrCode   string                 `json:"errcode"`
	Error     string                 `json:"error"`
}

// HasSingleStageFlow returns true if there exists at least 1 Flow with a single stage of stageName.
func (r RespUserInteractive) HasSingleStageFlow(stageName string) bool {
	for _, f := range r.Flows {
		if len(f.Stages) == 1 && f.Stages[0] == stageName {
			return true
		}
	}
	return false
}

// RespUserDisplayName is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-profile-userid-displayname
//easyjson:json
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// RespUserStatus is the JSON response for https://matrix.org/docs/spec/client_server/r0.6.0#get-matrix-client-r0-presence-userid-status
//easyjson:json
type RespUserStatus struct {
	Presence        string `json:"presence"`
	StatusMsg       string `json:"status_msg"`
	LastActiveAgo   int    `json:"last_active_ago"`
	CurrentlyActive bool   `json:"currently_active"`
}

// RespRegister is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-register
//easyjson:json
type RespRegister struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	HomeServer   string `json:"home_server"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

// RespLogin is the JSON response for http://matrix.org/docs/spec/client_server/r0.6.0.html#post-matrix-client-r0-login
//easyjson:json
type RespLogin struct {
	AccessToken string               `json:"access_token"`
	DeviceID    string               `json:"device_id"`
	HomeServer  string               `json:"home_server"`
	UserID      string               `json:"user_id"`
	WellKnown   DiscoveryInformation `json:"well_known"`
}

// DiscoveryInformation is the JSON Response for https://matrix.org/docs/spec/client_server/r0.6.0#get-well-known-matrix-client and a part of the JSON Response for https://matrix.org/docs/spec/client_server/r0.6.0#post-matrix-client-r0-login
type DiscoveryInformation struct {
	Homeserver struct {
		BaseURL string `json:"base_url"`
	} `json:"m.homeserver"`
	IdentityServer struct {
		BaseURL string `json:"base_url"`
	} `json:"m.identitiy_server"`
}

// RespLogout is the JSON response for http://matrix.org/docs/spec/client_server/r0.6.0.html#post-matrix-client-r0-logout
//easyjson:json
type RespLogout struct{}

// RespLogoutAll is the JSON response for https://matrix.org/docs/spec/client_server/r0.6.0#post-matrix-client-r0-logout-all
//easyjson:json
type RespLogoutAll struct{}

// RespCreateRoom is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-createroom
//easyjson:json
type RespCreateRoom struct {
	RoomID string `json:"room_id"`
}

// RespSync is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-sync
//easyjson:json
type RespSync struct {
	NextBatch   string `json:"next_batch"`
	AccountData struct {
		Events []Event `json:"events"`
	} `json:"account_data"`
	Presence struct {
		Events []Event `json:"events"`
	} `json:"presence"`
	Rooms struct {
		Leave map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"leave"`
		Join map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
			Ephemeral struct {
				Events []Event `json:"events"`
			} `json:"ephemeral"`
		} `json:"join"`
		Invite map[string]struct {
			State struct {
				Events []Event
			} `json:"invite_state"`
		} `json:"invite"`
	} `json:"rooms"`
}

// RespTurnServer is the JSON response from a Turn Server
//easyjson:json
type RespTurnServer struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	TTL      int      `json:"ttl"`
	URIs     []string `json:"uris"`
}

// RespResolveRoomsIDs is the JSON response for https://matrix.org/docs/spec/client_server/r0.6.1#get-matrix-client-r0-directory-room-roomalias
//easyjson:json
type RespResolveRoomsIDs struct {
	RoomID  string   `json:"room_id"`
	Servers []string `json:"servers"`
}

// RespRoomAliases is the JSON response for https://matrix.org/docs/spec/client_server/r0.6.1#get-matrix-client-r0-rooms-roomid-aliases
//easyjson:json
type RespRoomAliases struct {
	Aliases []string `json:"aliases"`
}
