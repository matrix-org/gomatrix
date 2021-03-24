package gomatrix

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// sample from docs
var testFilter = `
{
  "room": {
    "state": {
      "types": [
        "m.room.*"
      ],
      "not_rooms": [
        "!726s6s6q:example.com"
      ]
    },
    "timeline": {
      "limit": 10,
      "types": [
        "m.room.message"
      ],
      "not_rooms": [
        "!726s6s6q:example.com"
      ],
      "not_senders": [
        "@spam:example.com"
      ]
    },
    "ephemeral": {
      "types": [
        "m.receipt",
        "m.typing"
      ],
      "not_rooms": [
        "!726s6s6q:example.com"
      ],
      "not_senders": [
        "@spam:example.com"
      ]
    }
  },
  "presence": {
    "types": [
      "m.presence"
    ],
    "not_senders": [
      "@alice:example.com"
    ]
  },
  "event_format": "client",
  "event_fields": [
    "type",
    "content",
    "sender"
  ]
}`

func TestFilterValidate(t *testing.T) {
	var f Filter
	err := json.NewDecoder(strings.NewReader(testFilter)).Decode(&f)
	if err != nil {
		t.Fatalf("TestFilterValidate: Failed to parse %s", testFilter)
	}
	// test validadtion success
	if err = f.Validate(); err != nil {
		t.Fatalf("TestFilterValidate: Filter validation has failed, event_format: '%s'", f.EventFormat)
	}
	// test validation fail
	f.EventFormat = "unkown"
	err = f.Validate()
	if err == nil {
		t.Fatalf("TestFilterValidate: Filter validation false positive, event_format: '%s'", f.EventFormat)
	}
}

func TestDefaultFilter(t *testing.T) {
	defaultFilter := DefaultFilter()
	if reflect.TypeOf(defaultFilter) != reflect.TypeOf(Filter{}) {
		t.Fatal("TestDefaultFilter: Invalid type for default filter")
	}
	if defaultFilter.EventFormat != "client" {
		t.Fatalf("TestDefaultFilter: expected EventFormat %s,  got %s", "client", defaultFilter.EventFormat)
	}
}
