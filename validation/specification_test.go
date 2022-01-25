package validation

import "testing"

func isHuman(v interface{}) bool {
	return true
}

func isBoy(v interface{}) bool {
	return true
}

func isJimmy(v interface{}) bool {
	return false
}

func TestSpecification(t *testing.T) {
	ret := And(isHuman, isBoy)(struct{}{})
	if !ret {
		t.FailNow()
	}

	ret = Or(isBoy, isJimmy)(struct{}{})
	if !ret {
		t.FailNow()
	}

	ret = All(isHuman, isBoy, isJimmy)(struct{}{})
	if ret {
		t.FailNow()
	}

	ret = Any(isHuman, isBoy, isJimmy)(struct{}{})
	if !ret {
		t.FailNow()
	}
}
