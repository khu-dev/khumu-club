// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/ent/likeclub"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// LikeClub is the model entity for the LikeClub schema.
type LikeClub struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LikeClubQuery when eager-loading is set.
	Edges      LikeClubEdges `json:"edges"`
	club_likes *int
}

// LikeClubEdges holds the relations/edges for other nodes in the graph.
type LikeClubEdges struct {
	// Club holds the value of the club edge.
	Club *Club `json:"club,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeClubEdges) ClubOrErr() (*Club, error) {
	if e.loadedTypes[0] {
		if e.Club == nil {
			// The edge club was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LikeClub) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case likeclub.FieldID:
			values[i] = new(sql.NullInt64)
		case likeclub.FieldUsername:
			values[i] = new(sql.NullString)
		case likeclub.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case likeclub.ForeignKeys[0]: // club_likes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LikeClub", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LikeClub fields.
func (lc *LikeClub) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case likeclub.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lc.ID = int(value.Int64)
		case likeclub.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				lc.Username = value.String
			}
		case likeclub.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lc.CreatedAt = value.Time
			}
		case likeclub.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field club_likes", value)
			} else if value.Valid {
				lc.club_likes = new(int)
				*lc.club_likes = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryClub queries the "club" edge of the LikeClub entity.
func (lc *LikeClub) QueryClub() *ClubQuery {
	return (&LikeClubClient{config: lc.config}).QueryClub(lc)
}

// Update returns a builder for updating this LikeClub.
// Note that you need to call LikeClub.Unwrap() before calling this method if this LikeClub
// was returned from a transaction, and the transaction was committed or rolled back.
func (lc *LikeClub) Update() *LikeClubUpdateOne {
	return (&LikeClubClient{config: lc.config}).UpdateOne(lc)
}

// Unwrap unwraps the LikeClub entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lc *LikeClub) Unwrap() *LikeClub {
	tx, ok := lc.config.driver.(*txDriver)
	if !ok {
		panic("ent: LikeClub is not a transactional entity")
	}
	lc.config.driver = tx.drv
	return lc
}

// String implements the fmt.Stringer.
func (lc *LikeClub) String() string {
	var builder strings.Builder
	builder.WriteString("LikeClub(")
	builder.WriteString(fmt.Sprintf("id=%v", lc.ID))
	builder.WriteString(", username=")
	builder.WriteString(lc.Username)
	builder.WriteString(", created_at=")
	builder.WriteString(lc.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LikeClubs is a parsable slice of LikeClub.
type LikeClubs []*LikeClub

func (lc LikeClubs) config(cfg config) {
	for _i := range lc {
		lc[_i].config = cfg
	}
}
