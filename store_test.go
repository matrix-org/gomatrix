package gomatrix

import (
	"testing"
)

func TestNewInMemoryStore(t *testing.T) {
	s := NewInMemoryStore()
	if s.Filters == nil || s.NextBatch == nil || s.Rooms == nil {
		t.Fatal("TestNewInMemoryStore: New InMemoryStore object has not been correctly initialized")
	}
}

func TestSaveFilterID(t *testing.T) {
	uid, fid := "user001", "Filter001"
	s := NewInMemoryStore()
	s.SaveFilterID(uid, fid)
	f, ok := s.Filters[uid]
	if !ok || f != fid {
		t.Fatalf("TestSaveFilterID: The desired value '%s' was not correctly stored under key '%s' and could not be retrieved", fid, uid)
	}
}

func TestLoadFilterID(t *testing.T) {
	uid, fid := "user001", "Filter001"
	s := NewInMemoryStore()
	s.SaveFilterID(uid, fid)
	f := s.LoadFilterID(uid)
	if f != fid {
		t.Fatalf("TestLoadFilterID: The desired value '%s' has not been correctly loaded. Obtained '%s' instead", fid, f)
	}
}

func TestSaveNextBatch(t *testing.T) {
	uid, btok := "user001", "Batch001"
	s := NewInMemoryStore()
	s.SaveNextBatch(uid, btok)
	b, ok := s.NextBatch[uid]
	if !ok || b != btok {
		t.Fatalf("TestSaveNextBatch: The desired value '%s' has not been correctly stored in memory under key '%s'", btok, uid)
	}
}

func TestLoadNextBatch(t *testing.T) {
	uid, btok := "user001", "Batch001"
	s := NewInMemoryStore()
	s.SaveNextBatch(uid, btok)
	b := s.LoadNextBatch(uid)
	if b != btok {
		t.Fatalf("TestSaveNextBatch: The desired value '%s' has not been correctly loaded. Obtained '%s'", btok, b)
	}
}

func TestSaveRoom(t *testing.T) {
	rid := "room001"
	s := NewInMemoryStore()
	r := NewRoom(rid)
	s.SaveRoom(r)
	memr, ok := s.Rooms[rid]
	if !ok || memr != r {
		t.Fatalf("TestSaveRoom: Failed to correctly save room of id '%s' to memory", rid)
	}
}

func TestLoadRoom(t *testing.T) {
	rid := "room001"
	s := NewInMemoryStore()
	r := NewRoom(rid)
	s.SaveRoom(r)
	memr := s.LoadRoom(rid)
	if memr != r {
		t.Fatalf("TestSaveRoom: Failed to correctly Load room of id '%s' from memory", rid)
	}
}
