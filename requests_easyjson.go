// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gomatrix

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix(in *jlexer.Lexer, out *ReqUnbanUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix(out *jwriter.Writer, in ReqUnbanUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqUnbanUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqUnbanUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqUnbanUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqUnbanUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix1(in *jlexer.Lexer, out *ReqTyping) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "typing":
			out.Typing = bool(in.Bool())
		case "timeout":
			out.Timeout = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix1(out *jwriter.Writer, in ReqTyping) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"typing\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Typing))
	}
	{
		const prefix string = ",\"timeout\":"
		out.RawString(prefix)
		out.Int64(int64(in.Timeout))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqTyping) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqTyping) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqTyping) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqTyping) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix1(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix2(in *jlexer.Lexer, out *ReqRegister) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "username":
			out.Username = string(in.String())
		case "bind_email":
			out.BindEmail = bool(in.Bool())
		case "password":
			out.Password = string(in.String())
		case "device_id":
			out.DeviceID = string(in.String())
		case "initial_device_display_name":
			out.InitialDeviceDisplayName = string(in.String())
		case "auth":
			if m, ok := out.Auth.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Auth.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Auth = in.Interface()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix2(out *jwriter.Writer, in ReqRegister) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Username != "" {
		const prefix string = ",\"username\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	if in.BindEmail {
		const prefix string = ",\"bind_email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.BindEmail))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if in.DeviceID != "" {
		const prefix string = ",\"device_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DeviceID))
	}
	{
		const prefix string = ",\"initial_device_display_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InitialDeviceDisplayName))
	}
	if in.Auth != nil {
		const prefix string = ",\"auth\":"
		out.RawString(prefix)
		if m, ok := in.Auth.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Auth.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Auth))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqRegister) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqRegister) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqRegister) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqRegister) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix2(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix3(in *jlexer.Lexer, out *ReqRedact) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "reason":
			out.Reason = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix3(out *jwriter.Writer, in ReqRedact) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqRedact) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqRedact) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqRedact) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqRedact) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix3(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix4(in *jlexer.Lexer, out *ReqLogin) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "identifier":
			out.Identifier.UnmarshalEasyJSON(in)
		case "password":
			out.Password = string(in.String())
		case "medium":
			out.Medium = string(in.String())
		case "user":
			out.User = string(in.String())
		case "address":
			out.Address = string(in.String())
		case "token":
			out.Token = string(in.String())
		case "device_id":
			out.DeviceID = string(in.String())
		case "initial_device_display_name":
			out.InitialDeviceDisplayName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix4(out *jwriter.Writer, in ReqLogin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	if in.Identifier != nil {
		const prefix string = ",\"identifier\":"
		out.RawString(prefix)
		in.Identifier.MarshalEasyJSON(out)
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	if in.Medium != "" {
		const prefix string = ",\"medium\":"
		out.RawString(prefix)
		out.String(string(in.Medium))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	if in.Address != "" {
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	if in.Token != "" {
		const prefix string = ",\"token\":"
		out.RawString(prefix)
		out.String(string(in.Token))
	}
	if in.DeviceID != "" {
		const prefix string = ",\"device_id\":"
		out.RawString(prefix)
		out.String(string(in.DeviceID))
	}
	if in.InitialDeviceDisplayName != "" {
		const prefix string = ",\"initial_device_display_name\":"
		out.RawString(prefix)
		out.String(string(in.InitialDeviceDisplayName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqLogin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqLogin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqLogin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqLogin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix4(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix5(in *jlexer.Lexer, out *ReqKickUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "reason":
			out.Reason = string(in.String())
		case "user_id":
			out.UserID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix5(out *jwriter.Writer, in ReqKickUser) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqKickUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqKickUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqKickUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqKickUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix5(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix6(in *jlexer.Lexer, out *ReqInviteUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			out.UserID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix6(out *jwriter.Writer, in ReqInviteUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqInviteUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqInviteUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqInviteUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqInviteUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix6(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix7(in *jlexer.Lexer, out *ReqInvite3PID) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id_server":
			out.IDServer = string(in.String())
		case "medium":
			out.Medium = string(in.String())
		case "address":
			out.Address = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix7(out *jwriter.Writer, in ReqInvite3PID) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_server\":"
		out.RawString(prefix[1:])
		out.String(string(in.IDServer))
	}
	{
		const prefix string = ",\"medium\":"
		out.RawString(prefix)
		out.String(string(in.Medium))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqInvite3PID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqInvite3PID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqInvite3PID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqInvite3PID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix7(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix8(in *jlexer.Lexer, out *ReqCreateRoomAlias) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "room_id":
			out.RoomID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix8(out *jwriter.Writer, in ReqCreateRoomAlias) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"room_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.RoomID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqCreateRoomAlias) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqCreateRoomAlias) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqCreateRoomAlias) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqCreateRoomAlias) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix8(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix9(in *jlexer.Lexer, out *ReqCreateRoom) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "visibility":
			out.Visibility = string(in.String())
		case "room_alias_name":
			out.RoomAliasName = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "topic":
			out.Topic = string(in.String())
		case "invite":
			if in.IsNull() {
				in.Skip()
				out.Invite = nil
			} else {
				in.Delim('[')
				if out.Invite == nil {
					if !in.IsDelim(']') {
						out.Invite = make([]string, 0, 4)
					} else {
						out.Invite = []string{}
					}
				} else {
					out.Invite = (out.Invite)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Invite = append(out.Invite, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "invite_3pid":
			if in.IsNull() {
				in.Skip()
				out.Invite3PID = nil
			} else {
				in.Delim('[')
				if out.Invite3PID == nil {
					if !in.IsDelim(']') {
						out.Invite3PID = make([]ReqInvite3PID, 0, 1)
					} else {
						out.Invite3PID = []ReqInvite3PID{}
					}
				} else {
					out.Invite3PID = (out.Invite3PID)[:0]
				}
				for !in.IsDelim(']') {
					var v2 ReqInvite3PID
					(v2).UnmarshalEasyJSON(in)
					out.Invite3PID = append(out.Invite3PID, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "creation_content":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.CreationContent = make(map[string]interface{})
				} else {
					out.CreationContent = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 interface{}
					if m, ok := v3.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v3.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v3 = in.Interface()
					}
					(out.CreationContent)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "initial_state":
			if in.IsNull() {
				in.Skip()
				out.InitialState = nil
			} else {
				in.Delim('[')
				if out.InitialState == nil {
					if !in.IsDelim(']') {
						out.InitialState = make([]Event, 0, 0)
					} else {
						out.InitialState = []Event{}
					}
				} else {
					out.InitialState = (out.InitialState)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Event
					easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix10(in, &v4)
					out.InitialState = append(out.InitialState, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "preset":
			out.Preset = string(in.String())
		case "is_direct":
			out.IsDirect = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix9(out *jwriter.Writer, in ReqCreateRoom) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Visibility != "" {
		const prefix string = ",\"visibility\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Visibility))
	}
	if in.RoomAliasName != "" {
		const prefix string = ",\"room_alias_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RoomAliasName))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if len(in.Invite) != 0 {
		const prefix string = ",\"invite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Invite {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.Invite3PID) != 0 {
		const prefix string = ",\"invite_3pid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Invite3PID {
				if v7 > 0 {
					out.RawByte(',')
				}
				(v8).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.CreationContent) != 0 {
		const prefix string = ",\"creation_content\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.CreationContent {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				if m, ok := v9Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v9Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v9Value))
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.InitialState) != 0 {
		const prefix string = ",\"initial_state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.InitialState {
				if v10 > 0 {
					out.RawByte(',')
				}
				easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix10(out, v11)
			}
			out.RawByte(']')
		}
	}
	if in.Preset != "" {
		const prefix string = ",\"preset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Preset))
	}
	if in.IsDirect {
		const prefix string = ",\"is_direct\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDirect))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqCreateRoom) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqCreateRoom) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqCreateRoom) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqCreateRoom) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix9(l, v)
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix10(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "state_key":
			if in.IsNull() {
				in.Skip()
				out.StateKey = nil
			} else {
				if out.StateKey == nil {
					out.StateKey = new(string)
				}
				*out.StateKey = string(in.String())
			}
		case "sender":
			out.Sender = string(in.String())
		case "type":
			out.Type = EventType(in.String())
		case "origin_server_ts":
			out.Timestamp = int64(in.Int64())
		case "event_id":
			out.ID = string(in.String())
		case "room_id":
			out.RoomID = string(in.String())
		case "redacts":
			out.Redacts = string(in.String())
		case "unsigned":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Unsigned = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 interface{}
					if m, ok := v12.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v12.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v12 = in.Interface()
					}
					(out.Unsigned)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		case "content":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Content = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 interface{}
					if m, ok := v13.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v13.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v13 = in.Interface()
					}
					(out.Content)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
			}
		case "prev_content":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.PrevContent = make(map[string]interface{})
				} else {
					out.PrevContent = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v14 interface{}
					if m, ok := v14.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v14.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v14 = in.Interface()
					}
					(out.PrevContent)[key] = v14
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix10(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StateKey != nil {
		const prefix string = ",\"state_key\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.StateKey))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"origin_server_ts\":"
		out.RawString(prefix)
		out.Int64(int64(in.Timestamp))
	}
	{
		const prefix string = ",\"event_id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"room_id\":"
		out.RawString(prefix)
		out.String(string(in.RoomID))
	}
	if in.Redacts != "" {
		const prefix string = ",\"redacts\":"
		out.RawString(prefix)
		out.String(string(in.Redacts))
	}
	{
		const prefix string = ",\"unsigned\":"
		out.RawString(prefix)
		if in.Unsigned == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Unsigned {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				if m, ok := v15Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v15Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v15Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix)
		if in.Content == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v16First := true
			for v16Name, v16Value := range in.Content {
				if v16First {
					v16First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v16Name))
				out.RawByte(':')
				if m, ok := v16Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v16Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v16Value))
				}
			}
			out.RawByte('}')
		}
	}
	if len(in.PrevContent) != 0 {
		const prefix string = ",\"prev_content\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.PrevContent {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v17Name))
				out.RawByte(':')
				if m, ok := v17Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v17Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v17Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}
func easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix11(in *jlexer.Lexer, out *ReqBanUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "reason":
			out.Reason = string(in.String())
		case "user_id":
			out.UserID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix11(out *jwriter.Writer, in ReqBanUser) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"user_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqBanUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqBanUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncodeGithubComMatrixOrgGomatrix11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqBanUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqBanUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecodeGithubComMatrixOrgGomatrix11(l, v)
}
