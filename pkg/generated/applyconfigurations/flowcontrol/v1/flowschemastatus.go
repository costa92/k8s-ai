

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FlowSchemaStatusApplyConfiguration represents a declarative configuration of the FlowSchemaStatus type for use
// with apply.
type FlowSchemaStatusApplyConfiguration struct {
	Conditions []FlowSchemaConditionApplyConfiguration `json:"conditions,omitempty"`
}

// FlowSchemaStatusApplyConfiguration constructs a declarative configuration of the FlowSchemaStatus type for use with
// apply.
func FlowSchemaStatus() *FlowSchemaStatusApplyConfiguration {
	return &FlowSchemaStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *FlowSchemaStatusApplyConfiguration) WithConditions(values ...*FlowSchemaConditionApplyConfiguration) *FlowSchemaStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
