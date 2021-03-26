package gomatrix

import (
    "encoding/json"
    "testing"
    "strings"
    "reflect"
)

// example events from docs
var testEvent = `{
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
        "state_key": "@alice:example.org",
        "origin_server_ts": 1432735824653,
        "unsigned": {
            "age": 1234
        }
    }`

// IsRoom checks if a given variable is of type *Room 
func IsRoom(t interface{}) bool {
    switch t.(type) {
        case *Room:
            return true
        default:
            return false
    }
}

func TestNewRoom (t *testing.T) {
    var testID = "test-id-001"
    r := NewRoom(testID)
    if r.ID != testID {
        t.Fatalf("TestNewRoom: Expected ID '%s', but got '%s'", testID, r.ID)
    }
    if !IsRoom(r) {
        t.Fatalf("TestNewRoom: Wrong data type returned, expected 'Room'")
    }
}

func TestUpdateState (t *testing.T) {
    var e Event
	err := json.NewDecoder(strings.NewReader(testEvent)).Decode(&e)
	if err != nil {
		t.Fatalf("TestUpdateState: Something went wrong while parsing: %s", testEvent)
	}
    r := NewRoom("test_id_001")
    r.UpdateState(&e)
    // check mapping here first
    if r.State[e.Type][*e.StateKey] != &e {
        t.Fatalf("TestUpdateState: Wrong object reference. The reference is not pointing to the expected event object")
    }
}

func TestGetStateEvent (t *testing.T) {
    var e Event
	err := json.NewDecoder(strings.NewReader(testEvent)).Decode(&e)
	if err != nil {
		t.Fatalf("TestUpdateState: Something went wrong while parsing: %s", testEvent)
	}
    r := NewRoom("test_id_001")
    r.UpdateState(&e)
    e2 := *r.GetStateEvent(e.Type, *e.StateKey)
    if !reflect.DeepEqual(e, e2) {
        t.Fatalf("TestGetStateEvent: Wrong event object returned, expected object at address %p, got %p", &e, &e2)
    }
}

func TestGetMembershipState (t *testing.T) {

    var e Event
	err := json.NewDecoder(strings.NewReader(testEvent)).Decode(&e)
	if err != nil {
		t.Fatalf("TestUpdateState: Something went wrong while parsing: %s", testEvent)
	}
    r := NewRoom("test_id_001")
    r.UpdateState(&e)
    state := r.GetMembershipState("@alice:example.org")
}

