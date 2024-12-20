// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (dm *Date_Message) Messages(ctx context.Context) (result []*Message, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = dm.NamedMessages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = dm.Edges.MessagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = dm.QueryMessages().All(ctx)
	}
	return result, err
}

func (m *Message) User(ctx context.Context) (*User, error) {
	result, err := m.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryUser().Only(ctx)
	}
	return result, err
}

func (m *Message) DateMessage(ctx context.Context) (*Date_Message, error) {
	result, err := m.Edges.DateMessageOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryDateMessage().Only(ctx)
	}
	return result, err
}

func (u *User) Messages(ctx context.Context) (result []*Message, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedMessages(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.MessagesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryMessages().All(ctx)
	}
	return result, err
}
