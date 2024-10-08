// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConditionalUpdateRiskApplyConfiguration represents a declarative configuration of the ConditionalUpdateRisk type for use
// with apply.
type ConditionalUpdateRiskApplyConfiguration struct {
	URL           *string                              `json:"url,omitempty"`
	Name          *string                              `json:"name,omitempty"`
	Message       *string                              `json:"message,omitempty"`
	MatchingRules []ClusterConditionApplyConfiguration `json:"matchingRules,omitempty"`
}

// ConditionalUpdateRiskApplyConfiguration constructs a declarative configuration of the ConditionalUpdateRisk type for use with
// apply.
func ConditionalUpdateRisk() *ConditionalUpdateRiskApplyConfiguration {
	return &ConditionalUpdateRiskApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *ConditionalUpdateRiskApplyConfiguration) WithURL(value string) *ConditionalUpdateRiskApplyConfiguration {
	b.URL = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConditionalUpdateRiskApplyConfiguration) WithName(value string) *ConditionalUpdateRiskApplyConfiguration {
	b.Name = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ConditionalUpdateRiskApplyConfiguration) WithMessage(value string) *ConditionalUpdateRiskApplyConfiguration {
	b.Message = &value
	return b
}

// WithMatchingRules adds the given value to the MatchingRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchingRules field.
func (b *ConditionalUpdateRiskApplyConfiguration) WithMatchingRules(values ...*ClusterConditionApplyConfiguration) *ConditionalUpdateRiskApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchingRules")
		}
		b.MatchingRules = append(b.MatchingRules, *values[i])
	}
	return b
}
