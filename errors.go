package gomatrix

// RespError is the standard JSON error response from Homeservers. It also implements the Golang "error" interface.
// See http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type RespError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

// Error returns the errcode and error message.
func (e RespError) Error() string {
	return e.ErrCode + ": " + e.Err
}

// Is allow use 'errors.Is' from stdlib
func (e RespError) Is(err error) bool {
	target, ok := err.(RespError)
	if !ok {
		return false
	}

	if e.ErrCode == "M_UNKNOWN" && target.ErrCode == "M_UNKNOWN" {
		return e.Err == target.Err
	}
	return e.ErrCode == target.ErrCode
}

// Common error codes from https://matrix.org/docs/spec/client_server/latest#api-standards
//
// Can be used with errors.Is() to check the response code without casting the error:
//   err := client.SyncRequest()
//   if errors.Is(err, ErrUnknownToken) {
//     // logout
//   }
var (
	// ErrForbidden Forbidden access, e.g. joining a room without permission, failed login.
	ErrForbidden = RespError{ErrCode: "M_FORBIDDEN"}

	// ErrUnknownToken The access token specified was not recognised.
	ErrUnknownToken = RespError{ErrCode: "M_UNKNOWN_TOKEN"}

	// ErrMissingToken No access token was specified for the request.
	ErrMissingToken = RespError{ErrCode: "M_MISSING_TOKEN"}

	// ErrBadJSON Request contained valid JSON, but it was malformed in some way, e.g. missing required keys, invalid values for keys.
	ErrBadJSON = RespError{ErrCode: "M_BAD_JSON"}

	// ErrNotJSON Request did not contain valid JSON.
	ErrNotJSON = RespError{ErrCode: "M_NOT_JSON"}

	// ErrNotFound No resource was found for this request.
	ErrNotFound = RespError{ErrCode: "M_NOT_FOUND"}

	// ErrLimitExceeded Too many requests have been sent in a short period of time. Wait a while then try again.
	ErrLimitExceeded = RespError{ErrCode: "M_LIMIT_EXCEEDED"}

	// ErrUnknown An unknown error has occurred
	ErrUnknown = RespError{ErrCode: "M_UNKNOWN"}

	// ErrUnrecognuzed The server did not understand the request.
	ErrUnrecognuzed = RespError{ErrCode: "M_UNRECOGNIZED"}

	// ErrUnauthorized The request was not correctly authorized. Usually due to login failures.
	ErrUnauthorized = RespError{ErrCode: "M_UNAUTHORIZED"}

	// ErrUserDeactivated The user ID associated with the request has been deactivated.
	// Typically for endpoints that prove authentication, such as /login.
	ErrUserDeactivated = RespError{ErrCode: "M_USER_DEACTIVATED"}

	// ErrUserInUse Encountered when trying to register a user ID which has been taken.
	ErrUserInUse = RespError{ErrCode: "M_USER_IN_USE"}

	// ErrInvalidUsername Encountered when trying to register a user ID which is not valid.
	ErrInvalidUsername = RespError{ErrCode: "M_INVALID_USERNAME"}

	// ErrRoomInUse Sent when the room alias given to the createRoom API is already in use.
	ErrRoomInUse = RespError{ErrCode: "M_ROOM_IN_USE"}

	// ErrInvalidRoomState Sent when the initial state given to the createRoom API is invalid
	ErrInvalidRoomState = RespError{ErrCode: "M_INVALID_ROOM_STATE"}

	// ErrThreepidInUse Sent when a threepid given to an API cannot be used because the same threepid is already in use.
	ErrThreepidInUse = RespError{ErrCode: "M_THREEPID_IN_USE"}

	// ErrThreepidNotFound Sent when a threepid given to an API cannot be used because no record matching the threepid was found.
	ErrThreepidNotFound = RespError{ErrCode: "M_THREEPID_NOT_FOUND"}

	// ErrThreepidAuthFailed Authentication could not be performed on the third party identifier.
	ErrThreepidAuthFailed = RespError{ErrCode: "M_THREEPID_AUTH_FAILED"}

	// ErrThreepidDenied The server does not permit this third party identifier. This may happen if the server only permits, for example, email addresses from a particular domain.
	ErrThreepidDenied = RespError{ErrCode: "M_THREEPID_DENIED"}

	// ErrServerNotTrusted The client's request used a third party server, eg. identity server, that this server does not trust.
	ErrServerNotTrusted = RespError{ErrCode: "M_SERVER_NOT_TRUSTED"}

	// ErrUnsupportedRoomVersion The client's request to create a room used a room version that the server does not support.
	ErrUnsupportedRoomVersion = RespError{ErrCode: "M_UNSUPPORTED_ROOM_VERSION:"}

	// ErrIncompatibleRoomVersion The client attempted to join a room that has a version the server does not support. Inspect the room_version property of the error response for the room's version.
	ErrIncompatibleRoomVersion = RespError{ErrCode: "M_INCOMPATIBLE_ROOM_VERSION"}

	// ErrBadState The state change requested cannot be performed, such as attempting to unban a user who is not banned.
	ErrBadState = RespError{ErrCode: "M_BAD_STATE"}

	// ErrGuestAccessForbidden The room or resource does not permit guests to access it.
	ErrGuestAccessForbidden = RespError{ErrCode: "M_GUEST_ACCESS_FORBIDDEN"}

	// ErrCaptchaNeeded A Captcha is required to complete the request.
	ErrCaptchaNeeded = RespError{ErrCode: "M_CAPTCHA_NEEDED"}

	// ErrCaptchaInvalid The Captcha provided did not match what was expected.
	ErrCaptchaInvalid = RespError{ErrCode: "M_CAPTCHA_INVALID"}

	// ErrMissingParam A required parameter was missing from the request.
	ErrMissingParam = RespError{ErrCode: "M_MISSING_PARAM"}

	// ErrInvalidParam A parameter that was specified has the wrong value. For example, the server expected an integer and instead received a string.
	ErrInvalidParam = RespError{ErrCode: "M_INVALID_PARAM"}

	// ErrTooLarge The request or entity was too large.
	ErrTooLarge = RespError{ErrCode: "M_TOO_LARGE"}

	// ErrExclusive The resource being requested is reserved by an application service, or the application service making the request has not created the resource.
	ErrExclusive = RespError{ErrCode: "M_EXCLUSIVE"}

	// ErrRespuceLimitExceeded The request cannot be completed because the homeserver has reached a resource limit imposed on it. For example, a homeserver held in a shared hosting environment may
	//reach a resource limit if it starts using too much memory or disk space. The error MUST have an admin_contact field to provide the user receiving the error a place to reach out to.
	//Typically, this error will appear on routes which attempt to modify state (eg: sending messages, account data, etc) and not routes which only read state (eg: /sync, get account data, etc).
	ErrRespuceLimitExceeded = RespError{ErrCode: "M_RESOURCE_LIMIT_EXCEEDED"}

	// ErrCannotLeaveServerNoticeRoom The user is unable to reject an invite to join the server notices room.
	ErrCannotLeaveServerNoticeRoom = RespError{ErrCode: "M_CANNOT_LEAVE_SERVER_NOTICE_ROOM"}
)
