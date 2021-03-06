package types

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type Decimal float64

func (i Decimal) String() string {
	str, err := conv.String(float64(i))
	if err != nil {
		panic(err)
	}
	return str
}

func (i Decimal) Equals(other TypedValue) bool {
	if i.Kind() == other.Kind() {
		return i == other
	}
	return false
}

func (i Decimal) Compare(other TypedValue) TypeComparision {
	if i.Kind() != other.Kind() {
		panic(fmt.Sprintf(
			"couldn't compare between different kind of types where left is: %s(%v) and right is: %s(%v)",
			i.Kind(), i, other.Kind(), other),
		)
	}

	otherVal, err := conv.Float64(other)
	PanicOnError(err)
	if float64(i) == otherVal {
		return Equals
	} else if float64(i) > otherVal {
		return LeftIsBigger
	} else {
		return RightIsBigger
	}
}

func (i Decimal) Cast(toPtr interface{}) {
	if err := conv.Infer(toPtr, i.NativeType()); err != nil {
		panic(err)
	}
}

func (i Decimal) Kind() TypeKind {
	return KindDecimal
}

func (i Decimal) Ptr() TypedValue {
	return &i
}

func (i Decimal) NativeType() interface{} {
	return float64(i)
}

func (i Decimal) Precedence() int {
	return 2
}
