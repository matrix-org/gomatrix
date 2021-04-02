package gomatrix

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_LeaveRoom(t *testing.T) {
	cli := mockClient(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodPost && req.URL.Path == "/_matrix/client/r0/rooms/!foo:bar/leave" {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			}, nil
		}
		return nil, fmt.Errorf("unhandled URL: %s", req.URL.Path)
	})

	if _, err := cli.LeaveRoom("!foo:bar"); err != nil {
		t.Fatalf("LeaveRoom: error, got %s", err.Error())
	}
}

func TestClient_GetAvatarUrl(t *testing.T) {
	cli := mockClient(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet && req.URL.Path == "/_matrix/client/r0/profile/@user:test.gomatrix.org/avatar_url" {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"avatar_url":"mxc://matrix.org/iJaUjkshgdfsdkjfn"}`)),
			}, nil
		}
		return nil, fmt.Errorf("unhandled URL: %s", req.URL.Path)
	})

	if response, err := cli.GetAvatarURL(); err != nil {
		t.Fatalf("GetAvatarURL: Got error: %s", err.Error())
	} else if response == "" {
		t.Fatal("GetAvatarURL: Got empty response")
	} else if response != "mxc://matrix.org/iJaUjkshgdfsdkjfn" {
		t.Fatalf("Unexpected response URL: %s", response)
	}

}

func TestClient_SetAvatarUrl(t *testing.T) {
	cli := mockClient(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodPut && req.URL.Path == "/_matrix/client/r0/profile/@user:test.gomatrix.org/avatar_url" {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			}, nil
		}
		return nil, fmt.Errorf("unhandled URL: %s", req.URL.Path)
	})

	if err := cli.SetAvatarURL("https://foo.com/bar.png"); err != nil {
		t.Fatalf("GetAvatarURL: Got error: %s", err.Error())
	}
}

func TestClient_StateEvent(t *testing.T) {
	cli := mockClient(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet && req.URL.Path == "/_matrix/client/r0/rooms/!foo:bar/state/m.room.name" {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"name":"Room Name Goes Here"}`)),
			}, nil
		}
		return nil, fmt.Errorf("unhandled URL: %s", req.URL.Path)
	})

	content := struct {
		Name string `json:"name"`
	}{}

	if err := cli.StateEvent("!foo:bar", NameEventType, "", &content); err != nil {
		t.Fatalf("StateEvent: error, got %s", err.Error())
	}
	if content.Name != "Room Name Goes Here" {
		t.Fatalf("StateEvent: got %s, want %s", content.Name, "Room Name Goes Here")
	}
}

func TestClient_PublicRooms(t *testing.T) {
	cli := mockClient(func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet && req.URL.Path == "/_matrix/client/r0/publicRooms" {
			return &http.Response{
				StatusCode: 200,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "chunk": [
    {
      "aliases": [
        "#murrays:cheese.bar"
      ],
      "avatar_url": "mxc://bleeker.street/CHEDDARandBRIE",
      "guest_can_join": false,
      "name": "CHEESE",
      "num_joined_members": 37,
      "room_id": "!ol19s:bleecker.street",
      "topic": "Tasty tasty cheese",
      "world_readable": true
    }
  ],
  "next_batch": "p190q",
  "prev_batch": "p1902",
  "total_room_count_estimate": 115
}`)),
			}, nil
		}

		return nil, fmt.Errorf("unhandled URL: %s", req.URL.Path)
	})

	publicRooms, err := cli.PublicRooms(0, "", "")

	if err != nil {
		t.Fatalf("PublicRooms: error, got %s", err.Error())
	}
	if publicRooms.TotalRoomCountEstimate != 115 {
		t.Fatalf("PublicRooms: got %d, want %d", publicRooms.TotalRoomCountEstimate, 115)
	}
	if len(publicRooms.Chunk) != 1 {
		t.Fatalf("PublicRooms: got %d, want %d", len(publicRooms.Chunk), 1)
	}
	if publicRooms.Chunk[0].Name != "CHEESE" {
		t.Fatalf("PublicRooms: got %s, want %s", publicRooms.Chunk[0].Name, "CHEESE")
	}
	if publicRooms.Chunk[0].NumJoinedMembers != 37 {
		t.Fatalf("PublicRooms: got %d, want %d", publicRooms.Chunk[0].NumJoinedMembers, 37)
	}
}

func mockClient(fn func(*http.Request) (*http.Response, error)) *Client {
	mrt := MockRoundTripper{
		RT: fn,
	}

	cli, _ := NewClient("https://test.gomatrix.org", "@user:test.gomatrix.org", "abcdef")
	cli.Client = &http.Client{
		Transport: mrt,
	}
	return cli
}

type MockRoundTripper struct {
	RT func(*http.Request) (*http.Response, error)
}

func (t MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") == "" {
		panic("no auth")
	}
	return t.RT(req)
}
