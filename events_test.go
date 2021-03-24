package gomatrix

import (
	"encoding/json"
	"strings"
	"testing"
)

// example events from docs
var testEvents = map[string]string{
	"withFields": `{
        "content": {
            "body": "eventbody123",
            "msgtype": "m.text",
            "format": "org.matrix.custom.html",
            "formatted_body": "<b>This is an example text message</b>"
        },
        "type": "m.room.message",
        "event_id": "$143273582443PhrSn:example.org",
        "room_id": "!726s6s6q:example.com",
        "sender": "@example:example.org",
        "origin_server_ts": 1432735824653,
        "unsigned": {
            "age": 1234
        }
    }`,

	"withoutFields": `{
        "content": {
            "avatar_url": "mxc://example.org/SEsfnsuifSDFSSEF",
            "displayname": "Alice Margatroid",
            "membership": "join"
        },
        "event_id": "$143273582443PhrSn:example.org",
        "origin_server_ts": 1432735824653,
        "room_id": "!jEsUZKDJdhlrceRyVU:example.org",
        "sender": "@example:example.org",
        "state_key": "@alice:example.org",
        "type": "m.room.member",
        "unsigned": {
            "age": 1234
        }
    }`,
}

func TestEventWithBody(t *testing.T) {
	var e Event
	err := json.NewDecoder(strings.NewReader(testEvents["withFields"])).Decode(&e)
	if err != nil {
		t.Fatalf("TestFetchEventBody: Something went wrong while parsing: %s", testEvents["withFields"])
	}
	body, ok := e.Body()
	if !ok || body != "eventbody123" {
		t.Fatal("TestEventWithBody: Failed to fetch value of 'body' key in event content")
	}
}

func TestEventWithoutBody(t *testing.T) {
	var e Event
	err := json.NewDecoder(strings.NewReader(testEvents["withoutFields"])).Decode(&e)
	if err != nil {
		t.Fatalf("TestEventWithoutBody: Something went wrong while parsing: %s", testEvents["withFields"])
	}
	body, ok := e.Body()
	if ok || body != "" {
		t.Fatal("TestEventWithoutBody: Failed on 'Event.Body' call for event without a 'body' key")
	}
}

func TestEventWithMessageType(t *testing.T) {
	var e Event
	err := json.NewDecoder(strings.NewReader(testEvents["withFields"])).Decode(&e)
	if err != nil {
		t.Fatalf("TestEventWithMessageType: Something went wrong while parsing: %s", testEvents["withFields"])
	}
	msgtype, ok := e.MessageType()
	if !ok || msgtype != "m.text" {
		t.Fatal("TestEventWithMessageType: Failed to fetch value of 'msgtype' key in event content")
	}
}

func TestEventWithoutMessageType(t *testing.T) {
	var e Event
	err := json.NewDecoder(strings.NewReader(testEvents["withoutFields"])).Decode(&e)
	if err != nil {
		t.Fatalf("TestEventWithMessageType: Something went wrong while parsing: %s", testEvents["withFields"])
	}
	msgtype, ok := e.MessageType()
	if ok || msgtype != "" {
		t.Fatal("TestEventWithoutBody: Failed on 'Event.Body' call for event without a 'msgtype' key")
	}
}

var testHTML = `<div>a<h1>bc</h1>d<p>e<i>fg</i>hi</p>j<p>k<br/>l<b>m</b>no</p>p<small>q</small>rs</div>`

func TestGetHTMLMessage(t *testing.T) {
	msg := GetHTMLMessage("m.text", testHTML)
	if expected := "abcdefghijklmnopqrs"; msg.Body != expected {
		t.Fatalf("TestGetHTMLMessage: got '%s', expected '%s'", msg.Body, expected)
	}
	if msg.FormattedBody != testHTML {
		t.Fatalf("TestGetHTMLMessage: got '%s', expected '%s'", msg.FormattedBody, testHTML)
	}
	if msg.MsgType != "m.text" {
		t.Fatalf("TestGetHTMLMessage: got '%s', expected 'm.text'", msg.FormattedBody)
	}
	if expected := "org.matrix.custom.html"; msg.Format != expected {
		t.Fatalf("TestGetHTMLMessage: got '%s', expected '%s'", msg.Format, expected)
	}
}
