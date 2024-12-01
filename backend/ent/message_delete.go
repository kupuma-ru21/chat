// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/message"
	"backend/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageDelete is the builder for deleting a Message entity.
type MessageDelete struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageDelete builder.
func (md *MessageDelete) Where(ps ...predicate.Message) *MessageDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MessageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MessageDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MessageDeleteOne is the builder for deleting a single Message entity.
type MessageDeleteOne struct {
	md *MessageDelete
}

// Where appends a list predicates to the MessageDelete builder.
func (mdo *MessageDeleteOne) Where(ps ...predicate.Message) *MessageDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MessageDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{message.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MessageDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
