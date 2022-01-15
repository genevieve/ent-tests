package schema

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/genevieve/ent-tests/ent"
	"github.com/genevieve/ent-tests/ent/friendrequest"
	"github.com/genevieve/ent-tests/ent/hook"
)

// FriendRequest holds the schema definition for the FriendRequest entity.
type FriendRequest struct {
	ent.Schema
}

// Fields of the FriendRequest.
func (FriendRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("source_id"),
		field.Int("destination_id"),
		field.Enum("status").Values("pending", "accepted", "rejected"),
	}
}

// Edges of the FriendRequest.
func (FriendRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", User.Type).
			Ref("outgoing_friend_requests").
			Field("source_id").
			Unique().
			Required(),
		edge.From("destination", User.Type).
			Ref("incoming_friend_requests").
			Field("destination_id").
			Unique().
			Required(),
	}
}

func (FriendRequest) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.FriendRequestFunc(func(ctx context.Context, m *gen.FriendRequestMutation) (ent.Value, error) {
					oldStatus, err := m.OldStatus(ctx)
					if err != nil {
						return nil, err
					}

					currStatus, exists := m.Status()
					if !exists {
						return nil, errors.New("status field is missing")
					}

					val, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}

					// If request was rejected/pending and is now accepted, make a friend connection
					if oldStatus != friendrequest.StatusAccepted && currStatus == friendrequest.StatusAccepted {
						sourceID, err := m.OldSourceID(ctx)
						if err != nil {
							return nil, errors.New("source_id field is missing")
						}

						destinationID, err := m.OldDestinationID(ctx)
						if err != nil {
							return nil, errors.New("destination_id field is missing")
						}

						source, err := m.Client().User.Get(ctx, sourceID)
						if err != nil {
							return nil, err
						}

						destination, err := m.Client().User.Get(ctx, destinationID)
						if err != nil {
							return nil, err
						}

						err = m.Client().User.UpdateOne(source).AddFriends(destination).Exec(ctx)
						if err != nil {
							return nil, err
						}
					}

					return val, nil
				})
			},
			ent.OpUpdateOne,
		),
	}
}
