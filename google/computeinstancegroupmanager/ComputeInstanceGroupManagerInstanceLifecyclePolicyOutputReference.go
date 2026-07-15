// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/computeinstancegroupmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference interface {
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
	DefaultActionOnFailure() *string
	SetDefaultActionOnFailure(val *string)
	DefaultActionOnFailureInput() *string
	ForceUpdateOnRepair() *string
	SetForceUpdateOnRepair(val *string)
	ForceUpdateOnRepairInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeInstanceGroupManagerInstanceLifecyclePolicy
	SetInternalValue(val *ComputeInstanceGroupManagerInstanceLifecyclePolicy)
	OnFailedHealthCheck() *string
	SetOnFailedHealthCheck(val *string)
	OnFailedHealthCheckInput() *string
	OnRepair() ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepairOutputReference
	OnRepairInput() *ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair
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
	PutOnRepair(value *ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair)
	ResetDefaultActionOnFailure()
	ResetForceUpdateOnRepair()
	ResetOnFailedHealthCheck()
	ResetOnRepair()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference
type jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) DefaultActionOnFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) DefaultActionOnFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ForceUpdateOnRepair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateOnRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ForceUpdateOnRepairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateOnRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) InternalValue() *ComputeInstanceGroupManagerInstanceLifecyclePolicy {
	var returns *ComputeInstanceGroupManagerInstanceLifecyclePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) OnFailedHealthCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFailedHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) OnFailedHealthCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFailedHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) OnRepair() ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepairOutputReference {
	var returns ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepairOutputReference
	_jsii_.Get(
		j,
		"onRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) OnRepairInput() *ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair {
	var returns *ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair
	_jsii_.Get(
		j,
		"onRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference_Override(c ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetDefaultActionOnFailure(val *string) {
	if err := j.validateSetDefaultActionOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultActionOnFailure",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetForceUpdateOnRepair(val *string) {
	if err := j.validateSetForceUpdateOnRepairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateOnRepair",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetInternalValue(val *ComputeInstanceGroupManagerInstanceLifecyclePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetOnFailedHealthCheck(val *string) {
	if err := j.validateSetOnFailedHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFailedHealthCheck",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) PutOnRepair(value *ComputeInstanceGroupManagerInstanceLifecyclePolicyOnRepair) {
	if err := c.validatePutOnRepairParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnRepair",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ResetDefaultActionOnFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultActionOnFailure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ResetForceUpdateOnRepair() {
	_jsii_.InvokeVoid(
		c,
		"resetForceUpdateOnRepair",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ResetOnFailedHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetOnFailedHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ResetOnRepair() {
	_jsii_.InvokeVoid(
		c,
		"resetOnRepair",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

