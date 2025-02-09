// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"intern_traning/ent/news"
	"intern_traning/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NewsCreate is the builder for creating a News entity.
type NewsCreate struct {
	config
	mutation *NewsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NewsCreate) SetCreatedAt(t time.Time) *NewsCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NewsCreate) SetNillableCreatedAt(t *time.Time) *NewsCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NewsCreate) SetUpdatedAt(t time.Time) *NewsCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NewsCreate) SetNillableUpdatedAt(t *time.Time) *NewsCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetDeletedAt sets the "deleted_at" field.
func (nc *NewsCreate) SetDeletedAt(t time.Time) *NewsCreate {
	nc.mutation.SetDeletedAt(t)
	return nc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nc *NewsCreate) SetNillableDeletedAt(t *time.Time) *NewsCreate {
	if t != nil {
		nc.SetDeletedAt(*t)
	}
	return nc
}

// SetSlug sets the "slug" field.
func (nc *NewsCreate) SetSlug(s string) *NewsCreate {
	nc.mutation.SetSlug(s)
	return nc
}

// SetTitle sets the "title" field.
func (nc *NewsCreate) SetTitle(s string) *NewsCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetStatus sets the "status" field.
func (nc *NewsCreate) SetStatus(n news.Status) *NewsCreate {
	nc.mutation.SetStatus(n)
	return nc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nc *NewsCreate) SetNillableStatus(n *news.Status) *NewsCreate {
	if n != nil {
		nc.SetStatus(*n)
	}
	return nc
}

// SetDescription sets the "description" field.
func (nc *NewsCreate) SetDescription(s string) *NewsCreate {
	nc.mutation.SetDescription(s)
	return nc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nc *NewsCreate) SetNillableDescription(s *string) *NewsCreate {
	if s != nil {
		nc.SetDescription(*s)
	}
	return nc
}

// SetContent sets the "content" field.
func (nc *NewsCreate) SetContent(s string) *NewsCreate {
	nc.mutation.SetContent(s)
	return nc
}

// SetAuthorID sets the "author_id" field.
func (nc *NewsCreate) SetAuthorID(u uuid.UUID) *NewsCreate {
	nc.mutation.SetAuthorID(u)
	return nc
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (nc *NewsCreate) SetNillableAuthorID(u *uuid.UUID) *NewsCreate {
	if u != nil {
		nc.SetAuthorID(*u)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NewsCreate) SetID(u uuid.UUID) *NewsCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NewsCreate) SetNillableID(u *uuid.UUID) *NewsCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetAuthorEdgeID sets the "author_edge" edge to the User entity by ID.
func (nc *NewsCreate) SetAuthorEdgeID(id uuid.UUID) *NewsCreate {
	nc.mutation.SetAuthorEdgeID(id)
	return nc
}

// SetNillableAuthorEdgeID sets the "author_edge" edge to the User entity by ID if the given value is not nil.
func (nc *NewsCreate) SetNillableAuthorEdgeID(id *uuid.UUID) *NewsCreate {
	if id != nil {
		nc = nc.SetAuthorEdgeID(*id)
	}
	return nc
}

// SetAuthorEdge sets the "author_edge" edge to the User entity.
func (nc *NewsCreate) SetAuthorEdge(u *User) *NewsCreate {
	return nc.SetAuthorEdgeID(u.ID)
}

// Mutation returns the NewsMutation object of the builder.
func (nc *NewsCreate) Mutation() *NewsMutation {
	return nc.mutation
}

// Save creates the News in the database.
func (nc *NewsCreate) Save(ctx context.Context) (*News, error) {
	var (
		err  error
		node *News
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*News)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NewsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NewsCreate) SaveX(ctx context.Context) *News {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NewsCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NewsCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NewsCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := news.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.Status(); !ok {
		v := news.DefaultStatus
		nc.mutation.SetStatus(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := news.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NewsCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "News.created_at"`)}
	}
	if _, ok := nc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "News.slug"`)}
	}
	if v, ok := nc.mutation.Slug(); ok {
		if err := news.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "News.slug": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "News.title"`)}
	}
	if v, ok := nc.mutation.Title(); ok {
		if err := news.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "News.title": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "News.status"`)}
	}
	if v, ok := nc.mutation.Status(); ok {
		if err := news.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "News.status": %w`, err)}
		}
	}
	if v, ok := nc.mutation.Description(); ok {
		if err := news.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "News.description": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "News.content"`)}
	}
	if v, ok := nc.mutation.Content(); ok {
		if err := news.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "News.content": %w`, err)}
		}
	}
	return nil
}

func (nc *NewsCreate) sqlSave(ctx context.Context) (*News, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (nc *NewsCreate) createSpec() (*News, *sqlgraph.CreateSpec) {
	var (
		_node = &News{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: news.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: news.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(news.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(news.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.DeletedAt(); ok {
		_spec.SetField(news.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := nc.mutation.Slug(); ok {
		_spec.SetField(news.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := nc.mutation.Status(); ok {
		_spec.SetField(news.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := nc.mutation.Description(); ok {
		_spec.SetField(news.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := nc.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := nc.mutation.AuthorEdgeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   news.AuthorEdgeTable,
			Columns: []string{news.AuthorEdgeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AuthorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NewsCreateBulk is the builder for creating many News entities in bulk.
type NewsCreateBulk struct {
	config
	builders []*NewsCreate
}

// Save creates the News entities in the database.
func (ncb *NewsCreateBulk) Save(ctx context.Context) ([]*News, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*News, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NewsMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NewsCreateBulk) SaveX(ctx context.Context) []*News {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NewsCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NewsCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
