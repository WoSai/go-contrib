package validation

type (
	Specification func(interface{}) bool
)

func And(s1, s2 Specification) Specification {
	return func(i interface{}) bool {
		return s1(i) && s2(i)
	}
}

func Or(s1, s2 Specification) Specification {
	return func(i interface{}) bool {
		return s1(i) || s2(i)
	}
}
