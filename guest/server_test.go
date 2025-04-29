package guest

import (
	guestv1 "getting-married/gen/guest/v1"
	"path/filepath"
	"testing"

	"connectrpc.com/connect"
	"github.com/boltdb/bolt"
)

func TestSetGetParty(t *testing.T) {
	boltPath := filepath.Join(t.TempDir(), "test.db")
	boltDB, err := bolt.Open(boltPath, 0600, nil)
	if err != nil {
		t.Fatalf("failed to create BoltDB: %v", err)
	}
	defer boltDB.Close()

	server, err := NewServer(boltDB)
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	createReq := &connect.Request[guestv1.CreatePartyRequest]{
		Msg: &guestv1.CreatePartyRequest{
			Party: &guestv1.Party{
				DisplayName: "my_party",
			},
		},
	}
	createResp, err := server.CreateParty(t.Context(), createReq)
	if err != nil {
		t.Fatalf("failed create party: %v", err)
	}

	getReq := &connect.Request[guestv1.GetPartyRequest]{
		Msg: &guestv1.GetPartyRequest{
			Name: createResp.Msg.GetParty().GetName(),
		},
	}
	getResp, err := server.GetParty(t.Context(), getReq)
	if err != nil {
		t.Fatalf("failed get party: %v", err)
	}

	want := createReq.Msg.GetParty().GetDisplayName()
	got := getResp.Msg.GetParty().GetDisplayName()

	if want != got {
		t.Errorf("unexpected display name, want: %s, got: %s", want, got)
	}
}
