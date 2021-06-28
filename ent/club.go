// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/khu-dev/khumu-club/ent/club"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Club is the model entity for the Club schema.
type Club struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Hashtags holds the value of the "hashtags" field.
	Hashtags []string `json:"hashtags,omitempty"`
	// Images holds the value of the "images" field.
	Images []string `json:"images,omitempty"`
	// Homepage holds the value of the "homepage" field.
	Homepage string `json:"homepage,omitempty"`
	// Instagram holds the value of the "instagram" field.
	Instagram string `json:"instagram,omitempty"`
	// Facebook holds the value of the "facebook" field.
	Facebook string `json:"facebook,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Recommended holds the value of the "recommended" field.
	Recommended bool `json:"recommended,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClubQuery when eager-loading is set.
	Edges ClubEdges `json:"edges"`
}

// ClubEdges holds the relations/edges for other nodes in the graph.
type ClubEdges struct {
	// Likes holds the value of the likes edge.
	Likes []*LikeClub `json:"likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e ClubEdges) LikesOrErr() ([]*LikeClub, error) {
	if e.loadedTypes[0] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Club) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case club.FieldHashtags, club.FieldImages:
			values[i] = new([]byte)
		case club.FieldRecommended:
			values[i] = new(sql.NullBool)
		case club.FieldID:
			values[i] = new(sql.NullInt64)
		case club.FieldName, club.FieldSummary, club.FieldDescription, club.FieldHomepage, club.FieldInstagram, club.FieldFacebook, club.FieldPhone, club.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Club", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Club fields.
func (c *Club) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case club.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case club.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case club.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				c.Summary = value.String
			}
		case club.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case club.FieldHashtags:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field hashtags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Hashtags); err != nil {
					return fmt.Errorf("unmarshal field hashtags: %w", err)
				}
			}
		case club.FieldImages:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Images); err != nil {
					return fmt.Errorf("unmarshal field images: %w", err)
				}
			}
		case club.FieldHomepage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field homepage", values[i])
			} else if value.Valid {
				c.Homepage = value.String
			}
		case club.FieldInstagram:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instagram", values[i])
			} else if value.Valid {
				c.Instagram = value.String
			}
		case club.FieldFacebook:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field facebook", values[i])
			} else if value.Valid {
				c.Facebook = value.String
			}
		case club.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case club.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case club.FieldRecommended:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field recommended", values[i])
			} else if value.Valid {
				c.Recommended = value.Bool
			}
		}
	}
	return nil
}

// QueryLikes queries the "likes" edge of the Club entity.
func (c *Club) QueryLikes() *LikeClubQuery {
	return (&ClubClient{config: c.config}).QueryLikes(c)
}

// Update returns a builder for updating this Club.
// Note that you need to call Club.Unwrap() before calling this method if this Club
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Club) Update() *ClubUpdateOne {
	return (&ClubClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Club entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Club) Unwrap() *Club {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Club is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Club) String() string {
	var builder strings.Builder
	builder.WriteString("Club(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", summary=")
	builder.WriteString(c.Summary)
	builder.WriteString(", description=")
	builder.WriteString(c.Description)
	builder.WriteString(", hashtags=")
	builder.WriteString(fmt.Sprintf("%v", c.Hashtags))
	builder.WriteString(", images=")
	builder.WriteString(fmt.Sprintf("%v", c.Images))
	builder.WriteString(", homepage=")
	builder.WriteString(c.Homepage)
	builder.WriteString(", instagram=")
	builder.WriteString(c.Instagram)
	builder.WriteString(", facebook=")
	builder.WriteString(c.Facebook)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", email=")
	builder.WriteString(c.Email)
	builder.WriteString(", recommended=")
	builder.WriteString(fmt.Sprintf("%v", c.Recommended))
	builder.WriteByte(')')
	return builder.String()
}

// Clubs is a parsable slice of Club.
type Clubs []*Club

func (c Clubs) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
