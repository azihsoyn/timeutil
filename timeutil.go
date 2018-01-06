package timex

import "time"

type Time struct {
	time.Time
}

type Operator string

const (
	GT Operator = "gt"
	GE          = "ge"
	LT          = "lt"
	LE          = "le"
	EQ          = "eq"
)

func (t Time) GreaterThan(other time.Time) bool {
	return t.GT(other)
}

func (t Time) GT(other time.Time) bool {
	return t.After(other)
}

func (t Time) GE(other time.Time) bool {
	return !t.Before(other)
}

func (t Time) GreaterEqual(other time.Time) bool {
	return t.GE(other)
}

func (t Time) LT(other time.Time) bool {
	return t.Before(other)
}

func (t Time) LessThan(other time.Time) bool {
	return t.LT(other)
}

func (t Time) LE(other time.Time) bool {
	return !t.After(other)
}

func (t Time) LessEqual(other time.Time) bool {
	return t.LE(other)
}

func (t Time) Between(t1, t2 time.Time) bool {
	return (Time{t1}).LessEqual(t.Time) && t.LessEqual(t2)
}

func (t Time) EQ(other time.Time) bool {
	return t.Time.Equal(other)
}

func (t Time) Equal(other time.Time) bool {
	return t.EQ(other)
}

func Eval(t1 time.Time, op Operator, t2 time.Time) bool {
	t := Time{t1}
	switch op {
	case GT:
		return t.GT(t2)
	case GE:
		return t.GE(t2)
	case LT:
		return t.LT(t2)
	case LE:
		return t.LE(t2)
	case EQ:
		return t.EQ(t2)
	default:
		return false
	}
}
