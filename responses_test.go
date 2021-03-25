package gomatrix

import (
	"encoding/json"
	"strings"
	"testing"
)

// standard error response example from docs
var testErrorResponse = `{
    "errcode": "M_FORBIDDEN",
    "error": "something went wrong"
}`

// examples interactive response from docs

var testHasSingleStageTrue = `{
    "errcode": "M_FORBIDDEN",
    "error": "Invalid password",
    "completed": [ "example.type.foo" ],
    "flows":[
        {
            "stages":[ "test.type" ]
        },
        {
            "stages": [ "example.type.foo", "example.type.baz" ]
        }
    ],
    "params":{
        "example.type.baz":{
            "example_key":"foobar"
        }
    },
    "session":"xxxxxx"
}`

var testHasSingleStageFalse = `{
    "errcode": "M_FORBIDDEN",
    "error": "Invalid password",
    "completed": [ "example.type.foo" ],
    "flows":[
        {
            "stages": [ "example.type.foo2", "example.type.baz2" ]
        },
        {
            "stages": [ "example.type.foo", "example.type.baz" ]
        }
    ],
    "params":{
        "example.type.baz":{
            "example_key":"foobar"
        }
    },
    "session":"xxxxxx"
}`

func TestHasSingleStageFlow(t *testing.T) {
	var r RespUserInteractive
	err := json.NewDecoder(strings.NewReader(testHasSingleStageTrue)).Decode(&r)
	if err != nil {
		t.Fatalf("TestHasSingleStageFlow: Something went wrong while parsing %s", testHasSingleStageTrue)
	}
	if !r.HasSingleStageFlow("test.type") {
		t.Fatalf("TestHasSingleStageFlow: HasSingleStageFlow('test.type') returned false when true was expected")
	}
	err = json.NewDecoder(strings.NewReader(testHasSingleStageFalse)).Decode(&r)
	if err != nil {
		t.Fatalf("TestHasSingleStageFlow: Something went wrong while parsing %s", testHasSingleStageFalse)
	}
	if r.HasSingleStageFlow("test.type") {
		t.Fatalf("TestHasSingleStageFlow: HasSingleStageFlow('test.type') returned true when false was expected")
	}
}

func TestRespErrorError(t *testing.T) {
	var e RespError
	err := json.NewDecoder(strings.NewReader(testErrorResponse)).Decode(&e)
	if err != nil {
		t.Fatalf("TestRespErrorError: Something went wrong while parsing: %s", testErrorResponse)
	}
	if !strings.Contains(e.Error(), "M_FORBIDDEN") {
		t.Fatal("TestRespErrorError: Error string does not contain expected errorcode 'M_FORBIDDEN'")
	}
	if !strings.Contains(e.Error(), "something went wrong") {
		t.Fatal("TestRespErrorError: Error string does not contain expected error 'something went wrong'")
	}
}
