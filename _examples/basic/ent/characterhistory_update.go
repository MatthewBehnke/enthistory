// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory/_examples/basic/ent/characterhistory"
	"github.com/flume/enthistory/_examples/basic/ent/predicate"
)

// CharacterHistoryUpdate is the builder for updating CharacterHistory entities.
type CharacterHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterHistoryMutation
}

// Where appends a list predicates to the CharacterHistoryUpdate builder.
func (chu *CharacterHistoryUpdate) Where(ps ...predicate.CharacterHistory) *CharacterHistoryUpdate {
	chu.mutation.Where(ps...)
	return chu
}

// SetUpdatedAt sets the "updated_at" field.
func (chu *CharacterHistoryUpdate) SetUpdatedAt(t time.Time) *CharacterHistoryUpdate {
	chu.mutation.SetUpdatedAt(t)
	return chu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (chu *CharacterHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *CharacterHistoryUpdate {
	if t != nil {
		chu.SetUpdatedAt(*t)
	}
	return chu
}

// SetAge sets the "age" field.
func (chu *CharacterHistoryUpdate) SetAge(i int) *CharacterHistoryUpdate {
	chu.mutation.ResetAge()
	chu.mutation.SetAge(i)
	return chu
}

// AddAge adds i to the "age" field.
func (chu *CharacterHistoryUpdate) AddAge(i int) *CharacterHistoryUpdate {
	chu.mutation.AddAge(i)
	return chu
}

// SetName sets the "name" field.
func (chu *CharacterHistoryUpdate) SetName(s string) *CharacterHistoryUpdate {
	chu.mutation.SetName(s)
	return chu
}

// Mutation returns the CharacterHistoryMutation object of the builder.
func (chu *CharacterHistoryUpdate) Mutation() *CharacterHistoryMutation {
	return chu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (chu *CharacterHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CharacterHistoryMutation](ctx, chu.sqlSave, chu.mutation, chu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chu *CharacterHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := chu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (chu *CharacterHistoryUpdate) Exec(ctx context.Context) error {
	_, err := chu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chu *CharacterHistoryUpdate) ExecX(ctx context.Context) {
	if err := chu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chu *CharacterHistoryUpdate) check() error {
	if v, ok := chu.mutation.Age(); ok {
		if err := characterhistory.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "CharacterHistory.age": %w`, err)}
		}
	}
	return nil
}

func (chu *CharacterHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := chu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterhistory.Table, characterhistory.Columns, sqlgraph.NewFieldSpec(characterhistory.FieldID, field.TypeInt))
	if ps := chu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if chu.mutation.RefCleared() {
		_spec.ClearField(characterhistory.FieldRef, field.TypeInt)
	}
	if chu.mutation.UpdatedByCleared() {
		_spec.ClearField(characterhistory.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := chu.mutation.UpdatedAt(); ok {
		_spec.SetField(characterhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := chu.mutation.Age(); ok {
		_spec.SetField(characterhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := chu.mutation.AddedAge(); ok {
		_spec.AddField(characterhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := chu.mutation.Name(); ok {
		_spec.SetField(characterhistory.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, chu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	chu.mutation.done = true
	return n, nil
}

// CharacterHistoryUpdateOne is the builder for updating a single CharacterHistory entity.
type CharacterHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (chuo *CharacterHistoryUpdateOne) SetUpdatedAt(t time.Time) *CharacterHistoryUpdateOne {
	chuo.mutation.SetUpdatedAt(t)
	return chuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (chuo *CharacterHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *CharacterHistoryUpdateOne {
	if t != nil {
		chuo.SetUpdatedAt(*t)
	}
	return chuo
}

// SetAge sets the "age" field.
func (chuo *CharacterHistoryUpdateOne) SetAge(i int) *CharacterHistoryUpdateOne {
	chuo.mutation.ResetAge()
	chuo.mutation.SetAge(i)
	return chuo
}

// AddAge adds i to the "age" field.
func (chuo *CharacterHistoryUpdateOne) AddAge(i int) *CharacterHistoryUpdateOne {
	chuo.mutation.AddAge(i)
	return chuo
}

// SetName sets the "name" field.
func (chuo *CharacterHistoryUpdateOne) SetName(s string) *CharacterHistoryUpdateOne {
	chuo.mutation.SetName(s)
	return chuo
}

// Mutation returns the CharacterHistoryMutation object of the builder.
func (chuo *CharacterHistoryUpdateOne) Mutation() *CharacterHistoryMutation {
	return chuo.mutation
}

// Where appends a list predicates to the CharacterHistoryUpdate builder.
func (chuo *CharacterHistoryUpdateOne) Where(ps ...predicate.CharacterHistory) *CharacterHistoryUpdateOne {
	chuo.mutation.Where(ps...)
	return chuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (chuo *CharacterHistoryUpdateOne) Select(field string, fields ...string) *CharacterHistoryUpdateOne {
	chuo.fields = append([]string{field}, fields...)
	return chuo
}

// Save executes the query and returns the updated CharacterHistory entity.
func (chuo *CharacterHistoryUpdateOne) Save(ctx context.Context) (*CharacterHistory, error) {
	return withHooks[*CharacterHistory, CharacterHistoryMutation](ctx, chuo.sqlSave, chuo.mutation, chuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (chuo *CharacterHistoryUpdateOne) SaveX(ctx context.Context) *CharacterHistory {
	node, err := chuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (chuo *CharacterHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := chuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chuo *CharacterHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := chuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chuo *CharacterHistoryUpdateOne) check() error {
	if v, ok := chuo.mutation.Age(); ok {
		if err := characterhistory.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "CharacterHistory.age": %w`, err)}
		}
	}
	return nil
}

func (chuo *CharacterHistoryUpdateOne) sqlSave(ctx context.Context) (_node *CharacterHistory, err error) {
	if err := chuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterhistory.Table, characterhistory.Columns, sqlgraph.NewFieldSpec(characterhistory.FieldID, field.TypeInt))
	id, ok := chuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CharacterHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := chuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterhistory.FieldID)
		for _, f := range fields {
			if !characterhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != characterhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := chuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if chuo.mutation.RefCleared() {
		_spec.ClearField(characterhistory.FieldRef, field.TypeInt)
	}
	if chuo.mutation.UpdatedByCleared() {
		_spec.ClearField(characterhistory.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := chuo.mutation.UpdatedAt(); ok {
		_spec.SetField(characterhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := chuo.mutation.Age(); ok {
		_spec.SetField(characterhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := chuo.mutation.AddedAge(); ok {
		_spec.AddField(characterhistory.FieldAge, field.TypeInt, value)
	}
	if value, ok := chuo.mutation.Name(); ok {
		_spec.SetField(characterhistory.FieldName, field.TypeString, value)
	}
	_node = &CharacterHistory{config: chuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, chuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	chuo.mutation.done = true
	return _node, nil
}
