package spec

import (
	"fmt"
)

type Expression interface {
	Value() string
}

type evaluationExpression struct {
	label    string
	value    string
	operator string
}

func NewEval(label string, value string, operator string) Expression {
	return &evaluationExpression{
		label:    label,
		value:    value,
		operator: operator,
	}
}

func (e *evaluationExpression) Value() string {
	return fmt.Sprintf("%s %s '%s'", e.label, e.operator, e.value)
}

type andExpression struct {
	left  Expression
	right Expression
}

func (e *andExpression) Value() string {
	return fmt.Sprintf("(%s AND %s)", e.left.Value(), e.right.Value())
}

type orExpression struct {
	left  Expression
	right Expression
}

func (e *orExpression) Value() string {
	return fmt.Sprintf(" ( %s OR %s )", e.left.Value(), e.right.Value())
}

type notExpression struct {
	expression Expression
}

func (e *notExpression) Value() string {
	return fmt.Sprintf(" NOT ( %s )", e.expression)
}

func And(l Expression, r Expression) Expression {
	return &andExpression{
		left:  l,
		right: r,
	}
}

func Or(l Expression, r Expression) Expression {
	return &orExpression{
		left:  l,
		right: r,
	}
}

func Not(negate Expression) Expression {
	return &notExpression{expression: negate}
}

type Spec interface {
	Query() string
}

func Specify(e Expression) Spec {
	return &query{expression: e}
}

type query struct {
	expression Expression
}

func (q *query) Query() string {
	return q.expression.Value()
}

func ById(id string) Expression {
	return NewEval("id", id, "=")
}

func ByName(name string, lastName string) Expression {
	return And(
		NewEval("name", name, "="),
		NewEval("last_name", lastName, "="))
}

func ByEmail(email string) Expression {
	return NewEval("email", email, "=")
}
