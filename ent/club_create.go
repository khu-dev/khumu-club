// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-club/ent/club"
	"github.com/khu-dev/khumu-club/ent/likeclub"
)

// ClubCreate is the builder for creating a Club entity.
type ClubCreate struct {
	config
	mutation *ClubMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ClubCreate) SetName(s string) *ClubCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetSummary sets the "summary" field.
func (cc *ClubCreate) SetSummary(s string) *ClubCreate {
	cc.mutation.SetSummary(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *ClubCreate) SetDescription(s string) *ClubCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetHashtags sets the "hashtags" field.
func (cc *ClubCreate) SetHashtags(s []string) *ClubCreate {
	cc.mutation.SetHashtags(s)
	return cc
}

// SetImages sets the "images" field.
func (cc *ClubCreate) SetImages(s []string) *ClubCreate {
	cc.mutation.SetImages(s)
	return cc
}

// SetHomepage sets the "homepage" field.
func (cc *ClubCreate) SetHomepage(s string) *ClubCreate {
	cc.mutation.SetHomepage(s)
	return cc
}

// SetNillableHomepage sets the "homepage" field if the given value is not nil.
func (cc *ClubCreate) SetNillableHomepage(s *string) *ClubCreate {
	if s != nil {
		cc.SetHomepage(*s)
	}
	return cc
}

// SetInstagram sets the "instagram" field.
func (cc *ClubCreate) SetInstagram(s string) *ClubCreate {
	cc.mutation.SetInstagram(s)
	return cc
}

// SetNillableInstagram sets the "instagram" field if the given value is not nil.
func (cc *ClubCreate) SetNillableInstagram(s *string) *ClubCreate {
	if s != nil {
		cc.SetInstagram(*s)
	}
	return cc
}

// SetFacebook sets the "facebook" field.
func (cc *ClubCreate) SetFacebook(s string) *ClubCreate {
	cc.mutation.SetFacebook(s)
	return cc
}

// SetNillableFacebook sets the "facebook" field if the given value is not nil.
func (cc *ClubCreate) SetNillableFacebook(s *string) *ClubCreate {
	if s != nil {
		cc.SetFacebook(*s)
	}
	return cc
}

// SetPhone sets the "phone" field.
func (cc *ClubCreate) SetPhone(s string) *ClubCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cc *ClubCreate) SetNillablePhone(s *string) *ClubCreate {
	if s != nil {
		cc.SetPhone(*s)
	}
	return cc
}

// SetEmail sets the "email" field.
func (cc *ClubCreate) SetEmail(s string) *ClubCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cc *ClubCreate) SetNillableEmail(s *string) *ClubCreate {
	if s != nil {
		cc.SetEmail(*s)
	}
	return cc
}

// SetRecommended sets the "recommended" field.
func (cc *ClubCreate) SetRecommended(b bool) *ClubCreate {
	cc.mutation.SetRecommended(b)
	return cc
}

// SetNillableRecommended sets the "recommended" field if the given value is not nil.
func (cc *ClubCreate) SetNillableRecommended(b *bool) *ClubCreate {
	if b != nil {
		cc.SetRecommended(*b)
	}
	return cc
}

// AddLikeIDs adds the "likes" edge to the LikeClub entity by IDs.
func (cc *ClubCreate) AddLikeIDs(ids ...int) *ClubCreate {
	cc.mutation.AddLikeIDs(ids...)
	return cc
}

// AddLikes adds the "likes" edges to the LikeClub entity.
func (cc *ClubCreate) AddLikes(l ...*LikeClub) *ClubCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLikeIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cc *ClubCreate) Mutation() *ClubMutation {
	return cc.mutation
}

// Save creates the Club in the database.
func (cc *ClubCreate) Save(ctx context.Context) (*Club, error) {
	var (
		err  error
		node *Club
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClubCreate) SaveX(ctx context.Context) *Club {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *ClubCreate) defaults() {
	if _, ok := cc.mutation.Recommended(); !ok {
		v := club.DefaultRecommended
		cc.mutation.SetRecommended(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClubCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := cc.mutation.Summary(); !ok {
		return &ValidationError{Name: "summary", err: errors.New("ent: missing required field \"summary\"")}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if v, ok := cc.mutation.Description(); ok {
		if err := club.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Hashtags(); !ok {
		return &ValidationError{Name: "hashtags", err: errors.New("ent: missing required field \"hashtags\"")}
	}
	if _, ok := cc.mutation.Recommended(); !ok {
		return &ValidationError{Name: "recommended", err: errors.New("ent: missing required field \"recommended\"")}
	}
	return nil
}

func (cc *ClubCreate) sqlSave(ctx context.Context) (*Club, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ClubCreate) createSpec() (*Club, *sqlgraph.CreateSpec) {
	var (
		_node = &Club{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: club.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: club.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldSummary,
		})
		_node.Summary = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.Hashtags(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: club.FieldHashtags,
		})
		_node.Hashtags = value
	}
	if value, ok := cc.mutation.Images(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: club.FieldImages,
		})
		_node.Images = value
	}
	if value, ok := cc.mutation.Homepage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldHomepage,
		})
		_node.Homepage = value
	}
	if value, ok := cc.mutation.Instagram(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldInstagram,
		})
		_node.Instagram = value
	}
	if value, ok := cc.mutation.Facebook(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldFacebook,
		})
		_node.Facebook = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := cc.mutation.Recommended(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: club.FieldRecommended,
		})
		_node.Recommended = value
	}
	if nodes := cc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.LikesTable,
			Columns: []string{club.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likeclub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClubCreateBulk is the builder for creating many Club entities in bulk.
type ClubCreateBulk struct {
	config
	builders []*ClubCreate
}

// Save creates the Club entities in the database.
func (ccb *ClubCreateBulk) Save(ctx context.Context) ([]*Club, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Club, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClubMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClubCreateBulk) SaveX(ctx context.Context) []*Club {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
