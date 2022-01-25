package validation

import (
	"fmt"
	"testing"
)

type testSpecV1 struct{}
type testSpecV2 struct{}

func (ts *testSpecV1) IsSatisfiedBy(candidate interface{}) bool {
	v, ok := candidate.(int)
	if ok {
		return v >= 10
	}
	return false
}

func (ts *testSpecV1) Rule() string {
	return "value >= 10"
}

func (ts2 *testSpecV2) IsSatisfiedBy(candidate interface{}) bool {
	v, ok := candidate.(int)
	if ok {
		return v < 80
	}
	return false
}

func (ts2 *testSpecV2) Rule() string {
	return "value < 80"
}

func TestSpec(t *testing.T) {
	var s1 = new(testSpecV1)
	var s2 = new(testSpecV1)
	var s3 = new(testSpecV2)

	a1 := NewAndSpecification(s1, s2)
	fmt.Println(a1.Rule())
	if !a1.IsSatisfiedBy(20) {
		t.FailNow()
	}
	if a1.IsSatisfiedBy(5) {
		t.FailNow()
	}

	a2 := NewAndSpecification(a1, s3)
	fmt.Println(a2.Rule())
	if !a2.IsSatisfiedBy(50) {
		t.FailNow()
	}

	a3 := NewAndNotCondition(a1, s3)
	fmt.Println(a3.Rule())
	if !a3.IsSatisfiedBy(90) {
		t.FailNow()
	}
	if a3.IsSatisfiedBy(70) {
		t.FailNow()
	}

	a4 := NewAnyCondition(s1, a3)
	fmt.Println(a4.Rule())
	if !a4.IsSatisfiedBy(40) {
		t.FailNow()
	}
}
