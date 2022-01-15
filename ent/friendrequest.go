// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/genevieve/ent-tests/ent/friendrequest"
	"github.com/genevieve/ent-tests/ent/user"
)

// FriendRequest is the model entity for the FriendRequest schema.
type FriendRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID int `json:"source_id,omitempty"`
	// DestinationID holds the value of the "destination_id" field.
	DestinationID int `json:"destination_id,omitempty"`
	// Status holds the value of the "status" field.
	Status friendrequest.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FriendRequestQuery when eager-loading is set.
	Edges FriendRequestEdges `json:"edges"`
}

// FriendRequestEdges holds the relations/edges for other nodes in the graph.
type FriendRequestEdges struct {
	// Source holds the value of the source edge.
	Source *User `json:"source,omitempty"`
	// Destination holds the value of the destination edge.
	Destination *User `json:"destination,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SourceOrErr returns the Source value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendRequestEdges) SourceOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Source == nil {
			// The edge source was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Source, nil
	}
	return nil, &NotLoadedError{edge: "source"}
}

// DestinationOrErr returns the Destination value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendRequestEdges) DestinationOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Destination == nil {
			// The edge destination was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Destination, nil
	}
	return nil, &NotLoadedError{edge: "destination"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FriendRequest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case friendrequest.FieldID, friendrequest.FieldSourceID, friendrequest.FieldDestinationID:
			values[i] = new(sql.NullInt64)
		case friendrequest.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FriendRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FriendRequest fields.
func (fr *FriendRequest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friendrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fr.ID = int(value.Int64)
		case friendrequest.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				fr.SourceID = int(value.Int64)
			}
		case friendrequest.FieldDestinationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field destination_id", values[i])
			} else if value.Valid {
				fr.DestinationID = int(value.Int64)
			}
		case friendrequest.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				fr.Status = friendrequest.Status(value.String)
			}
		}
	}
	return nil
}

// QuerySource queries the "source" edge of the FriendRequest entity.
func (fr *FriendRequest) QuerySource() *UserQuery {
	return (&FriendRequestClient{config: fr.config}).QuerySource(fr)
}

// QueryDestination queries the "destination" edge of the FriendRequest entity.
func (fr *FriendRequest) QueryDestination() *UserQuery {
	return (&FriendRequestClient{config: fr.config}).QueryDestination(fr)
}

// Update returns a builder for updating this FriendRequest.
// Note that you need to call FriendRequest.Unwrap() before calling this method if this FriendRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FriendRequest) Update() *FriendRequestUpdateOne {
	return (&FriendRequestClient{config: fr.config}).UpdateOne(fr)
}

// Unwrap unwraps the FriendRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FriendRequest) Unwrap() *FriendRequest {
	tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FriendRequest is not a transactional entity")
	}
	fr.config.driver = tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FriendRequest) String() string {
	var builder strings.Builder
	builder.WriteString("FriendRequest(")
	builder.WriteString(fmt.Sprintf("id=%v", fr.ID))
	builder.WriteString(", source_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.SourceID))
	builder.WriteString(", destination_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.DestinationID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", fr.Status))
	builder.WriteByte(')')
	return builder.String()
}

// FriendRequests is a parsable slice of FriendRequest.
type FriendRequests []*FriendRequest

func (fr FriendRequests) config(cfg config) {
	for _i := range fr {
		fr[_i].config = cfg
	}
}