// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowcompanionagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowCompanionAgentToolsetToolsOutputReference interface {
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
	ConfirmationRequirement() *string
	SetConfirmationRequirement(val *string)
	ConfirmationRequirementInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperationId() *string
	SetOperationId(val *string)
	OperationIdInput() *string
	ParameterOverrides() DialogflowCompanionAgentToolsetToolsParameterOverridesList
	ParameterOverridesInput() interface{}
	ProactiveEnabled() interface{}
	SetProactiveEnabled(val interface{})
	ProactiveEnabledInput() interface{}
	ReactiveEnabled() interface{}
	SetReactiveEnabled(val interface{})
	ReactiveEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Toolset() *string
	SetToolset(val *string)
	ToolsetInput() *string
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
	PutParameterOverrides(value interface{})
	ResetConfirmationRequirement()
	ResetOperationId()
	ResetParameterOverrides()
	ResetProactiveEnabled()
	ResetReactiveEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCompanionAgentToolsetToolsOutputReference
type jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ConfirmationRequirement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confirmationRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ConfirmationRequirementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confirmationRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) OperationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) OperationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ParameterOverrides() DialogflowCompanionAgentToolsetToolsParameterOverridesList {
	var returns DialogflowCompanionAgentToolsetToolsParameterOverridesList
	_jsii_.Get(
		j,
		"parameterOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ParameterOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ProactiveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ProactiveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ReactiveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reactiveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ReactiveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reactiveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) Toolset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ToolsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolsetInput",
		&returns,
	)
	return returns
}


func NewDialogflowCompanionAgentToolsetToolsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCompanionAgentToolsetToolsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCompanionAgentToolsetToolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentToolsetToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCompanionAgentToolsetToolsOutputReference_Override(d DialogflowCompanionAgentToolsetToolsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentToolsetToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetConfirmationRequirement(val *string) {
	if err := j.validateSetConfirmationRequirementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confirmationRequirement",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetOperationId(val *string) {
	if err := j.validateSetOperationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationId",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetProactiveEnabled(val interface{}) {
	if err := j.validateSetProactiveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveEnabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetReactiveEnabled(val interface{}) {
	if err := j.validateSetReactiveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reactiveEnabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference)SetToolset(val *string) {
	if err := j.validateSetToolsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolset",
		val,
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) PutParameterOverrides(value interface{}) {
	if err := d.validatePutParameterOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParameterOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ResetConfirmationRequirement() {
	_jsii_.InvokeVoid(
		d,
		"resetConfirmationRequirement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ResetOperationId() {
	_jsii_.InvokeVoid(
		d,
		"resetOperationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ResetParameterOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ResetProactiveEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetProactiveEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ResetReactiveEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetReactiveEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentToolsetToolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

