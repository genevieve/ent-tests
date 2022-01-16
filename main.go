package main

import (
	"context"
	"log"

	"github.com/genevieve/ent-tests/ent"
	"github.com/genevieve/ent-tests/ent/friendrequest"
	_ "github.com/genevieve/ent-tests/ent/runtime"
	"github.com/genevieve/ent-tests/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	u1, err := client.User.Create().
		SetName("ash").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	u2, err := client.User.Create().
		SetName("echo").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	req, err := client.FriendRequest.Create().
		SetSource(u1).
		SetDestination(u2).
		SetStatus(friendrequest.StatusPending).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	updatedU1, err := client.User.Query().
		Where(user.ID(u1.ID)).
		WithOutgoingFriendRequests().
		WithIncomingFriendRequests().
		First(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(updatedU1.Edges.OutgoingFriendRequests) != 1 {
		log.Fatalf("expected u1 to have 1 outgoing friend requests: %v", updatedU1.Edges)
	}

	err = client.FriendRequest.
		UpdateOne(req).
		SetStatus(friendrequest.StatusAccepted).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	updatedU1, err = client.User.Query().
		Where(user.ID(u1.ID)).
		WithFriends().
		First(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(updatedU1.Edges.Friends) != 1 {
		log.Fatalf("expected u1 to have 1 friend: %v", updatedU1.Edges)
	}
	updatedU2, err := client.User.Query().
		Where(user.ID(u2.ID)).
		WithFriends().
		First(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(updatedU2.Edges.Friends) != 1 {
		log.Fatalf("expected u2 to have 1 friend: %v", updatedU2.Edges)
	}
}
