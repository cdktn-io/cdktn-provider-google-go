// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/computeglobalvmextensionpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference interface {
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
	ConflictBehavior() *string
	SetConflictBehavior(val *string)
	ConflictBehaviorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput
	SetInternalValue(val *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput)
	Name() *string
	SetName(val *string)
	NameInput() *string
	PredefinedRolloutPlan() *string
	SetPredefinedRolloutPlan(val *string)
	PredefinedRolloutPlanInput() *string
	RetryUuid() *string
	SetRetryUuid(val *string)
	RetryUuidInput() *string
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
	ResetConflictBehavior()
	ResetName()
	ResetPredefinedRolloutPlan()
	ResetRetryUuid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference
type jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ConflictBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ConflictBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InternalValue() *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput {
	var returns *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) PredefinedRolloutPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedRolloutPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) PredefinedRolloutPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedRolloutPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) RetryUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) RetryUuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryUuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference {
	_init_.Initialize()

	if err := validateNewComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeGlobalVmExtensionPolicy.ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference_Override(c ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeGlobalVmExtensionPolicy.ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetConflictBehavior(val *string) {
	if err := j.validateSetConflictBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conflictBehavior",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetInternalValue(val *ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetPredefinedRolloutPlan(val *string) {
	if err := j.validateSetPredefinedRolloutPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedRolloutPlan",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetRetryUuid(val *string) {
	if err := j.validateSetRetryUuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryUuid",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetConflictBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetConflictBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetPredefinedRolloutPlan() {
	_jsii_.InvokeVoid(
		c,
		"resetPredefinedRolloutPlan",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ResetRetryUuid() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryUuid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeGlobalVmExtensionPolicyRolloutOperationRolloutInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

