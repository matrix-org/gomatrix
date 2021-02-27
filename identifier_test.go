package gomatrix

import (
	"testing"
)

func TestUserIDType(t *testing.T) {
	var userID UserIdentifier
	if idType := userID.Type(); idType != "m.id.user" {
		t.Fatalf("TestUserIDType: expected ID type '%s', got '%s'", "m.id.user", idType)
	}
}

func TestNewUserID(t *testing.T) {
	userID := NewUserIdentifier("user0001")
	if userID.User != "user0001" || userID.IDType != "m.id.user" {
		t.Fatalf(
			"TestNewUserID: invalid UserIdentifier returned, IDType: '%s' User: '%s'",
			userID.IDType,
			userID.User,
		)
	}
}

func TestThirdPartyIDType(t *testing.T) {
	var tpID ThirdpartyIdentifier
	if idType := tpID.Type(); idType != "m.id.thirdparty" {
		t.Fatalf("TestUserIDType: expected ID type '%s', got '%s'", "m.id.user", idType)
	}
}

func TestNewThirdPartyID(t *testing.T) {
	tpID := NewThirdpartyIdentifier("m01", "addr01")
	if tpID.Medium != "m01" || tpID.Address != "addr01" || tpID.IDType != "m.id.thirdparty" {
		t.Fatalf("NewThirdpartyIdentifier: invalid ThirdPartyIdentifier returned, Medium: '%s', Address: '%s', IDType: '%s'", tpID.Medium, tpID.Address, tpID.IDType)
	}
}

func TestPhoneIDType(t *testing.T) {
	var phoneID PhoneIdentifier
	if idType := phoneID.Type(); idType != "m.id.phone" {
		t.Fatalf("TestPhoneIDType: expected ID type '%s', got '%s'", "m.id.phone", idType)
	}
}

func TestNewPhoneID(t *testing.T) {
	phoneID := NewPhoneIdentifier("country01", "(11)111-111")
	if phoneID.Country != "country01" || phoneID.Phone != "(11)111-111" || phoneID.IDType != "m.id.phone" {
		t.Fatalf("TestNewPhoneID: invalid PhoneIdentifier returned, country: '%s', phone: '%s', IDType: '%s'", phoneID.Country, phoneID.Phone, phoneID.IDType)
	}
}
