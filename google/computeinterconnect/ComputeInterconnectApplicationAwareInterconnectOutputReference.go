// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/computeinterconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeInterconnectApplicationAwareInterconnectOutputReference interface {
	cdktn.ComplexObject
	BandwidthPercentagePolicy() ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference
	BandwidthPercentagePolicyInput() *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy
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
	InternalValue() *ComputeInterconnectApplicationAwareInterconnect
	SetInternalValue(val *ComputeInterconnectApplicationAwareInterconnect)
	ProfileDescription() *string
	SetProfileDescription(val *string)
	ProfileDescriptionInput() *string
	ShapeAveragePercentage() ComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList
	ShapeAveragePercentageInput() interface{}
	StrictPriorityPolicy() ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference
	StrictPriorityPolicyInput() *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutBandwidthPercentagePolicy(value *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy)
	PutShapeAveragePercentage(value interface{})
	PutStrictPriorityPolicy(value *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy)
	ResetBandwidthPercentagePolicy()
	ResetProfileDescription()
	ResetShapeAveragePercentage()
	ResetStrictPriorityPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectApplicationAwareInterconnectOutputReference
type jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) BandwidthPercentagePolicy() ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference {
	var returns ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference
	_jsii_.Get(
		j,
		"bandwidthPercentagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) BandwidthPercentagePolicyInput() *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy {
	var returns *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy
	_jsii_.Get(
		j,
		"bandwidthPercentagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) InternalValue() *ComputeInterconnectApplicationAwareInterconnect {
	var returns *ComputeInterconnectApplicationAwareInterconnect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ProfileDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ProfileDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ShapeAveragePercentage() ComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList {
	var returns ComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList
	_jsii_.Get(
		j,
		"shapeAveragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ShapeAveragePercentageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapeAveragePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) StrictPriorityPolicy() ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference {
	var returns ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference
	_jsii_.Get(
		j,
		"strictPriorityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) StrictPriorityPolicyInput() *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy {
	var returns *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy
	_jsii_.Get(
		j,
		"strictPriorityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInterconnectApplicationAwareInterconnectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeInterconnectApplicationAwareInterconnectOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInterconnectApplicationAwareInterconnectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnect.ComputeInterconnectApplicationAwareInterconnectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInterconnectApplicationAwareInterconnectOutputReference_Override(c ComputeInterconnectApplicationAwareInterconnectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnect.ComputeInterconnectApplicationAwareInterconnectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetInternalValue(val *ComputeInterconnectApplicationAwareInterconnect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetProfileDescription(val *string) {
	if err := j.validateSetProfileDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileDescription",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) PutBandwidthPercentagePolicy(value *ComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy) {
	if err := c.validatePutBandwidthPercentagePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBandwidthPercentagePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) PutShapeAveragePercentage(value interface{}) {
	if err := c.validatePutShapeAveragePercentageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShapeAveragePercentage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) PutStrictPriorityPolicy(value *ComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy) {
	if err := c.validatePutStrictPriorityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStrictPriorityPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ResetBandwidthPercentagePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidthPercentagePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ResetProfileDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetProfileDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ResetShapeAveragePercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetShapeAveragePercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ResetStrictPriorityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetStrictPriorityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInterconnectApplicationAwareInterconnectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

