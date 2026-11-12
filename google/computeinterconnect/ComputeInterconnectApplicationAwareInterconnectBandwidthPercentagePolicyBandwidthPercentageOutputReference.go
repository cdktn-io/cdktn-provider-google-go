// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/computeinterconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficClass() *string
	SetTrafficClass(val *string)
	TrafficClassInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetPercentage()
	ResetTrafficClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference
type jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) TrafficClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) TrafficClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficClassInput",
		&returns,
	)
	return returns
}


func NewComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnect.ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference_Override(c ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnect.ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetPercentage(val *float64) {
	if err := j.validateSetPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference)SetTrafficClass(val *string) {
	if err := j.validateSetTrafficClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficClass",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetPercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ResetTrafficClass() {
	_jsii_.InvokeVoid(
		c,
		"resetTrafficClass",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyBandwidthPercentageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

