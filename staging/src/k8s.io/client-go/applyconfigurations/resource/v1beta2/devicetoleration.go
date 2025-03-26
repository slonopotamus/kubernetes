/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	resourcev1beta2 "k8s.io/api/resource/v1beta2"
)

// DeviceTolerationApplyConfiguration represents a declarative configuration of the DeviceToleration type for use
// with apply.
type DeviceTolerationApplyConfiguration struct {
	Key               *string                                   `json:"key,omitempty"`
	Operator          *resourcev1beta2.DeviceTolerationOperator `json:"operator,omitempty"`
	Value             *string                                   `json:"value,omitempty"`
	Effect            *resourcev1beta2.DeviceTaintEffect        `json:"effect,omitempty"`
	TolerationSeconds *int64                                    `json:"tolerationSeconds,omitempty"`
}

// DeviceTolerationApplyConfiguration constructs a declarative configuration of the DeviceToleration type for use with
// apply.
func DeviceToleration() *DeviceTolerationApplyConfiguration {
	return &DeviceTolerationApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *DeviceTolerationApplyConfiguration) WithKey(value string) *DeviceTolerationApplyConfiguration {
	b.Key = &value
	return b
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *DeviceTolerationApplyConfiguration) WithOperator(value resourcev1beta2.DeviceTolerationOperator) *DeviceTolerationApplyConfiguration {
	b.Operator = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *DeviceTolerationApplyConfiguration) WithValue(value string) *DeviceTolerationApplyConfiguration {
	b.Value = &value
	return b
}

// WithEffect sets the Effect field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Effect field is set to the value of the last call.
func (b *DeviceTolerationApplyConfiguration) WithEffect(value resourcev1beta2.DeviceTaintEffect) *DeviceTolerationApplyConfiguration {
	b.Effect = &value
	return b
}

// WithTolerationSeconds sets the TolerationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TolerationSeconds field is set to the value of the last call.
func (b *DeviceTolerationApplyConfiguration) WithTolerationSeconds(value int64) *DeviceTolerationApplyConfiguration {
	b.TolerationSeconds = &value
	return b
}
