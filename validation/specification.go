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

func All(ss ...Specification) Specification {
	return func(i interface{}) bool {
		for _, spec := range ss {
			if !spec(i) {
				return false
			}
		}
		return true
	}
}

func Any(ss ...Specification) Specification {
	return func(i interface{}) bool {
		for _, spec := range ss {
			if spec(i) {
				return true
			}
		}
		return false
	}
}
