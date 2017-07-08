package main

//import "github.com/kr/pretty"
import (
	"os"
	"time"
	"log"
	"github.com/matrix-org/gomatrix"
	"fmt"
)

// Login to a local homeserver and set the user ID and access token on success.
func Login(server, username, password string) *gomatrix.Client {
	cli, _ := gomatrix.NewClient(server, "", "")
	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)
	return cli
}

func ProcessResponse(resSync *gomatrix.RespSync) {
	for _,v := range resSync.Rooms.Join {
		//fmt.Printf("%v: %# v\n\n\n", k, pretty.Formatter(v.Timeline.Events))
		for _, e := range v.Timeline.Events {
			if e.Type ==  "m.room.message" {
				fmt.Printf("%v: %v\n", e.Sender, e.Content["body"])
			}
		}

	}
}


func ExtractRooms(resSync *gomatrix.RespSync) map[string]string{
	rooms := map[string]string{}
	for roomID,v := range resSync.Rooms.Join {
		//fmt.Printf("%v: %# v\n\n\n", k, pretty.Formatter(v.Timeline.Events))
		for _, e := range v.Timeline.Events {
			if e.Type ==  "m.room.name" {
				rooms[fmt.Sprint(e.Content["name"])] = roomID
			}
		}

	}
	return rooms
}

func main () {
	log.Print("Logging in...")
	username := os.Args[2]
	inviteUser := os.Args[4]
	roomname := "TestRoomFor" + username
	roomalias := roomname
	cli := Login(os.Args[1], username, os.Args[3])
	log.Println("Done!")
	cli.SetDisplayName(os.Args[2])
	log.Print("Creating room...")
	resp, err := cli.CreateRoom(&gomatrix.ReqCreateRoom{
	RoomAliasName: roomalias,
	Name : roomname,
	Preset: "public_chat",
})
	id := "!JHMwXpSMhcSnyAGDWB:matrix.org"
	if err != nil {
		log.Println(err)
		log.Println("Failed!")
		id = "#" + roomalias + ":matrix.org"
	} else {
		log.Println("Done!")
		fmt.Println("Room:", resp.RoomID)
		id =  resp.RoomID
		log.Print("Joining room...")
		if _, err := cli.JoinRoom(id, "", nil); err != nil {
			panic(err)
		}
		log.Println("Done!")
		fmt.Println("Room:", id)
	}

	log.Print("Syncing...")
	filter := `{"room":{"timeline":{"limit":50}}}`
	resSync, err := cli.SyncRequest(5, "", filter, false, "")
	if err != nil {
		panic(err)
	}
	ProcessResponse(resSync)

	nextbatch := resSync.NextBatch

	//get room id
	rooms := ExtractRooms(resSync)
	log.Println(rooms)
	id = rooms[roomname]


	log.Print("Sending message...")
	if _, err := cli.SendText(id, "Hello") ; err != nil {
		panic(err)
	}
	log.Println("Done!")

	log.Print("Inviting user...")
	if _, err := cli.InviteUser(id, &gomatrix.ReqInviteUser{UserID: inviteUser}) ; err != nil {
		log.Println(err)
	}
	log.Println("Done!")



	for {
		resSync, err := cli.SyncRequest(5, nextbatch, filter, false, "")
		ProcessResponse(resSync)
		if err != nil {
			panic(err)
		}
		nextbatch = resSync.NextBatch
		time.Sleep(2 * time.Second)
	}
	log.Println("Done!")
}
