// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/date_message"
	"backend/ent/message"
	"backend/ent/predicate"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUserID sets the "user_id" field.
func (mu *MessageUpdate) SetUserID(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetUserID(u)
	return mu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableUserID(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetUserID(*u)
	}
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(s string) *MessageUpdate {
	mu.mutation.SetContent(s)
	return mu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableContent(s *string) *MessageUpdate {
	if s != nil {
		mu.SetContent(*s)
	}
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MessageUpdate) SetCreatedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableCreatedAt(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetDateID sets the "date_id" field.
func (mu *MessageUpdate) SetDateID(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetDateID(u)
	return mu
}

// SetNillableDateID sets the "date_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableDateID(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetDateID(*u)
	}
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MessageUpdate) SetUser(u *User) *MessageUpdate {
	return mu.SetUserID(u.ID)
}

// SetDateMessageID sets the "date_message" edge to the Date_Message entity by ID.
func (mu *MessageUpdate) SetDateMessageID(id uuid.UUID) *MessageUpdate {
	mu.mutation.SetDateMessageID(id)
	return mu
}

// SetDateMessage sets the "date_message" edge to the Date_Message entity.
func (mu *MessageUpdate) SetDateMessage(d *Date_Message) *MessageUpdate {
	return mu.SetDateMessageID(d.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MessageUpdate) ClearUser() *MessageUpdate {
	mu.mutation.ClearUser()
	return mu
}

// ClearDateMessage clears the "date_message" edge to the Date_Message entity.
func (mu *MessageUpdate) ClearDateMessage() *MessageUpdate {
	mu.mutation.ClearDateMessage()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.Content(); ok {
		if err := message.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Message.content": %w`, err)}
		}
	}
	if mu.mutation.UserCleared() && len(mu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.user"`)
	}
	if mu.mutation.DateMessageCleared() && len(mu.mutation.DateMessageIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.date_message"`)
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.DateMessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.DateMessageTable,
			Columns: []string{message.DateMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(date_message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DateMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.DateMessageTable,
			Columns: []string{message.DateMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(date_message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetUserID sets the "user_id" field.
func (muo *MessageUpdateOne) SetUserID(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetUserID(u)
	return muo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableUserID(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetUserID(*u)
	}
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(s string) *MessageUpdateOne {
	muo.mutation.SetContent(s)
	return muo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableContent(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetContent(*s)
	}
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MessageUpdateOne) SetCreatedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableCreatedAt(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetDateID sets the "date_id" field.
func (muo *MessageUpdateOne) SetDateID(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetDateID(u)
	return muo
}

// SetNillableDateID sets the "date_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableDateID(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetDateID(*u)
	}
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MessageUpdateOne) SetUser(u *User) *MessageUpdateOne {
	return muo.SetUserID(u.ID)
}

// SetDateMessageID sets the "date_message" edge to the Date_Message entity by ID.
func (muo *MessageUpdateOne) SetDateMessageID(id uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetDateMessageID(id)
	return muo
}

// SetDateMessage sets the "date_message" edge to the Date_Message entity.
func (muo *MessageUpdateOne) SetDateMessage(d *Date_Message) *MessageUpdateOne {
	return muo.SetDateMessageID(d.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MessageUpdateOne) ClearUser() *MessageUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// ClearDateMessage clears the "date_message" edge to the Date_Message entity.
func (muo *MessageUpdateOne) ClearDateMessage() *MessageUpdateOne {
	muo.mutation.ClearDateMessage()
	return muo
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.Content(); ok {
		if err := message.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Message.content": %w`, err)}
		}
	}
	if muo.mutation.UserCleared() && len(muo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.user"`)
	}
	if muo.mutation.DateMessageCleared() && len(muo.mutation.DateMessageIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.date_message"`)
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.UserTable,
			Columns: []string{message.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.DateMessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.DateMessageTable,
			Columns: []string{message.DateMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(date_message.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DateMessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.DateMessageTable,
			Columns: []string{message.DateMessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(date_message.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
