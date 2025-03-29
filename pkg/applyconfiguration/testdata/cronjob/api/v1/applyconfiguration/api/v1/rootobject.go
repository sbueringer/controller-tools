// Code generated by applyconfiguration. DO NOT EDIT.

package v1

// RootObjectApplyConfiguration represents a declarative configuration of the RootObject type for use
// with apply.
type RootObjectApplyConfiguration struct {
	Nested *NestedObjectApplyConfiguration `json:"nested,omitempty"`
}

// RootObjectApplyConfiguration constructs a declarative configuration of the RootObject type for use with
// apply.
func RootObject() *RootObjectApplyConfiguration {
	return &RootObjectApplyConfiguration{}
}

// WithNested sets the Nested field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Nested field is set to the value of the last call.
func (b *RootObjectApplyConfiguration) WithNested(value *NestedObjectApplyConfiguration) *RootObjectApplyConfiguration {
	b.Nested = value
	return b
}
