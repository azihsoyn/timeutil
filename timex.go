package timex

import (
	"time"
)

type Comparable interface {
	GreaterThan() ComparableResult
	GreaterEqual() ComparableResult
	LessThan() ComparableResult
	LessEqual() ComparableResult
	Equal() ComparableResult
}

type Betweenable interface {
	Between() BetweenableY
}

type BetweenableY interface {
	Y(t time.Time) BetweenableAnd
}

type BetweenableAnd interface {
	And() BetweenResult
}

type BetweenResult interface {
	Z(t time.Time) bool
}

type ComparableResult interface {
	B(t time.Time) bool
}

type Timex struct {
	op Operator
	t1 time.Time
	t2 time.Time
}

func A(t time.Time) Comparable {
	return &Timex{
		t1 : t,
	}
}

func (x *Timex) GreaterThan() ComparableResult {
	x.op = GT
	return x
}

func (x *Timex) GreaterEqual() ComparableResult {
	x.op = GE
	return x
}

func (x *Timex) LessThan() ComparableResult {
	x.op = LT
	return x
}

func (x *Timex) LessEqual() ComparableResult {
	x.op = LE
	return x
}

func (x *Timex) Equal() ComparableResult {
	x.op = EQ
	return x
}

func (x *Timex) B(t time.Time) bool {
	return Eval(x.t1, x.op, t)
}

func X(t time.Time) Betweenable {
	return &Timex{
		t1: t,
	}
}

func (x *Timex) Between() BetweenableY {
	return x
}

func (x *Timex) Y(t time.Time) BetweenableAnd {
	x.t2 = t
	return x
}

func (x *Timex) And() BetweenResult {
	return x
}

func (x *Timex) Z(t time.Time) bool {
	return (Time{x.t1}).Between(x.t2, t)
}
