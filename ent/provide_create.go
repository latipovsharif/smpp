// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"smpp/ent/messages"
	"smpp/ent/provide"
	"smpp/ent/usermonthmessage"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProvideCreate is the builder for creating a Provide entity.
type ProvideCreate struct {
	config
	mutation *ProvideMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProvideCreate) SetName(s string) *ProvideCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIPAdres sets the "ip_adres" field.
func (pc *ProvideCreate) SetIPAdres(s string) *ProvideCreate {
	pc.mutation.SetIPAdres(s)
	return pc
}

// SetCreateAt sets the "create_at" field.
func (pc *ProvideCreate) SetCreateAt(t time.Time) *ProvideCreate {
	pc.mutation.SetCreateAt(t)
	return pc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pc *ProvideCreate) SetNillableCreateAt(t *time.Time) *ProvideCreate {
	if t != nil {
		pc.SetCreateAt(*t)
	}
	return pc
}

// SetUpdateAt sets the "update_at" field.
func (pc *ProvideCreate) SetUpdateAt(t time.Time) *ProvideCreate {
	pc.mutation.SetUpdateAt(t)
	return pc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (pc *ProvideCreate) SetNillableUpdateAt(t *time.Time) *ProvideCreate {
	if t != nil {
		pc.SetUpdateAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProvideCreate) SetID(u uuid.UUID) *ProvideCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddProviderIDIDs adds the "provider_id" edge to the UserMonthMessage entity by IDs.
func (pc *ProvideCreate) AddProviderIDIDs(ids ...uuid.UUID) *ProvideCreate {
	pc.mutation.AddProviderIDIDs(ids...)
	return pc
}

// AddProviderID adds the "provider_id" edges to the UserMonthMessage entity.
func (pc *ProvideCreate) AddProviderID(u ...*UserMonthMessage) *ProvideCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddProviderIDIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Messages entity by IDs.
func (pc *ProvideCreate) AddMessageIDs(ids ...uuid.UUID) *ProvideCreate {
	pc.mutation.AddMessageIDs(ids...)
	return pc
}

// AddMessages adds the "messages" edges to the Messages entity.
func (pc *ProvideCreate) AddMessages(m ...*Messages) *ProvideCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMessageIDs(ids...)
}

// Mutation returns the ProvideMutation object of the builder.
func (pc *ProvideCreate) Mutation() *ProvideMutation {
	return pc.mutation
}

// Save creates the Provide in the database.
func (pc *ProvideCreate) Save(ctx context.Context) (*Provide, error) {
	var (
		err  error
		node *Provide
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvideMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProvideCreate) SaveX(ctx context.Context) *Provide {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *ProvideCreate) defaults() {
	if _, ok := pc.mutation.CreateAt(); !ok {
		v := provide.DefaultCreateAt()
		pc.mutation.SetCreateAt(v)
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		v := provide.DefaultUpdateAt()
		pc.mutation.SetUpdateAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := provide.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProvideCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := pc.mutation.IPAdres(); !ok {
		return &ValidationError{Name: "ip_adres", err: errors.New("ent: missing required field \"ip_adres\"")}
	}
	if _, ok := pc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := pc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	return nil
}

func (pc *ProvideCreate) sqlSave(ctx context.Context) (*Provide, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (pc *ProvideCreate) createSpec() (*Provide, *sqlgraph.CreateSpec) {
	var (
		_node = &Provide{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provide.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provide.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provide.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.IPAdres(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provide.FieldIPAdres,
		})
		_node.IPAdres = value
	}
	if value, ok := pc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: provide.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := pc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: provide.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if nodes := pc.mutation.ProviderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provide.ProviderIDTable,
			Columns: []string{provide.ProviderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: usermonthmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provide.MessagesTable,
			Columns: []string{provide.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: messages.FieldID,
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

// ProvideCreateBulk is the builder for creating many Provide entities in bulk.
type ProvideCreateBulk struct {
	config
	builders []*ProvideCreate
}

// Save creates the Provide entities in the database.
func (pcb *ProvideCreateBulk) Save(ctx context.Context) ([]*Provide, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Provide, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvideMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProvideCreateBulk) SaveX(ctx context.Context) []*Provide {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
