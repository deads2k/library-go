// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OperatorStatusApplyConfiguration represents an declarative configuration of the OperatorStatus type for use
// with apply.
type OperatorStatusApplyConfiguration struct {
	ObservedGeneration      *int64                                `json:"observedGeneration,omitempty"`
	Conditions              []OperatorConditionApplyConfiguration `json:"conditions,omitempty"`
	Version                 *string                               `json:"version,omitempty"`
	ReadyReplicas           *int32                                `json:"readyReplicas,omitempty"`
	LatestAvailableRevision *int32                                `json:"latestAvailableRevision,omitempty"`
	Generations             []GenerationStatusApplyConfiguration  `json:"generations,omitempty"`
}

// OperatorStatusApplyConfiguration constructs an declarative configuration of the OperatorStatus type for use with
// apply.
func OperatorStatus() *OperatorStatusApplyConfiguration {
	return &OperatorStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *OperatorStatusApplyConfiguration) WithObservedGeneration(value int64) *OperatorStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *OperatorStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *OperatorStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *OperatorStatusApplyConfiguration) WithVersion(value string) *OperatorStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *OperatorStatusApplyConfiguration) WithReadyReplicas(value int32) *OperatorStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *OperatorStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *OperatorStatusApplyConfiguration {
	b.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *OperatorStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *OperatorStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}
