package validation

import (
	"fmt"
	"strings"
)

type (
	// Specification define a interface
	Specification interface {
		// IsSatisfiedBy check if candidate matched the specification
		IsSatisfiedBy(candidate interface{}) bool
		// Rule print the specification details
		Rule() string
	}

	AndCondition struct {
		left  Specification
		right Specification
	}

	AndNotCondition struct {
		left  Specification
		right Specification
	}

	OrCondition struct {
		left  Specification
		right Specification
	}

	NotCondition struct {
		spec Specification
	}

	OrNotCondition struct {
		left  Specification
		right Specification
	}

	AllCondition struct {
		specs []Specification
	}

	AnyCondition struct {
		specs []Specification
	}
)

var (
	_ Specification = (*AndCondition)(nil)
	_ Specification = (*OrCondition)(nil)
	_ Specification = (*AndNotCondition)(nil)
	_ Specification = (*OrNotCondition)(nil)
	_ Specification = (*NotCondition)(nil)
	_ Specification = (*AllCondition)(nil)
	_ Specification = (*AndCondition)(nil)
)

func NewAndSpecification(s1, s2 Specification) Specification {
	return &AndCondition{s1, s2}
}

func (and *AndCondition) IsSatisfiedBy(candidate interface{}) bool {
	return and.left.IsSatisfiedBy(candidate) && and.right.IsSatisfiedBy(candidate)
}

func (and *AndCondition) Rule() string {
	return fmt.Sprintf("(%s) AND (%s)", and.left.Rule(), and.right.Rule())
}

func NewOrCondition(s1, s2 Specification) Specification {
	return &OrCondition{s1, s2}
}

func (or *OrCondition) IsSatisfiedBy(candidate interface{}) bool {
	return or.left.IsSatisfiedBy(candidate) || or.right.IsSatisfiedBy(candidate)
}

func (or *OrCondition) Rule() string {
	return fmt.Sprintf("(%s) OR (%s)", or.left.Rule(), or.right.Rule())
}

func NewAndNotCondition(s1, s2 Specification) Specification {
	return &AndNotCondition{s1, s2}
}

func (an *AndNotCondition) IsSatisfiedBy(candidate interface{}) bool {
	return an.left.IsSatisfiedBy(candidate) && (!an.right.IsSatisfiedBy(candidate))
}

func (an *AndNotCondition) Rule() string {
	return fmt.Sprintf("(%s) AND !(%s)", an.left.Rule(), an.right.Rule())
}

func NewNotCondition(spec Specification) Specification {
	return &NotCondition{spec: spec}
}

func (not *NotCondition) IsSatisfiedBy(candidate interface{}) bool {
	return !not.spec.IsSatisfiedBy(candidate)
}

func (not *NotCondition) Rule() string {
	return fmt.Sprintf("!(%s)", not.spec.Rule())
}

func NewOrNotCondition(s1, s2 Specification) Specification {
	return &OrNotCondition{s1, s2}
}

func (on *OrNotCondition) IsSatisfiedBy(candidate interface{}) bool {
	return on.left.IsSatisfiedBy(candidate) || (!on.right.IsSatisfiedBy(candidate))
}

func (on *OrNotCondition) Rule() string {
	return fmt.Sprintf("(%s) OR !(%s)", on.left.Rule(), on.right.Rule())
}

func NewAllCondition(specs ...Specification) Specification {
	return &AllCondition{specs: specs}
}

func (all *AllCondition) IsSatisfiedBy(candidate interface{}) bool {
	for _, spec := range all.specs {
		if !spec.IsSatisfiedBy(candidate) {
			return false
		}
	}
	return true
}

func (all *AllCondition) Rule() string {
	rules := make([]string, len(all.specs))
	for i := 0; i < len(all.specs); i++ {
		rules[i] = all.specs[i].Rule()
	}
	return fmt.Sprintf("ALL[%s]", strings.Join(rules, ", "))
}

func NewAnyCondition(specs ...Specification) Specification {
	return &AnyCondition{specs: specs}
}

func (any *AnyCondition) IsSatisfiedBy(candidate interface{}) bool {
	for _, spec := range any.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}

func (any *AnyCondition) Rule() string {
	rules := make([]string, len(any.specs))
	for i := 0; i < len(any.specs); i++ {
		rules[i] = any.specs[i].Rule()
	}
	return fmt.Sprintf("ANY[%s]", strings.Join(rules, ", "))
}
