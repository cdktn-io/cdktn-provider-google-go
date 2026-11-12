// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowcompanionagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowCompanionAgentCesToolSpecsOutputReference interface {
	cdktn.ComplexObject
	CesTool() *string
	SetCesTool(val *string)
	CesToolInput() *string
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
	ParameterOverrides() DialogflowCompanionAgentCesToolSpecsParameterOverridesList
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
	ResetCesTool()
	ResetConfirmationRequirement()
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

// The jsii proxy struct for DialogflowCompanionAgentCesToolSpecsOutputReference
type jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) CesTool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cesTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) CesToolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cesToolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ConfirmationRequirement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confirmationRequirement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ConfirmationRequirementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confirmationRequirementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ParameterOverrides() DialogflowCompanionAgentCesToolSpecsParameterOverridesList {
	var returns DialogflowCompanionAgentCesToolSpecsParameterOverridesList
	_jsii_.Get(
		j,
		"parameterOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ParameterOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ProactiveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ProactiveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proactiveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ReactiveEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reactiveEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ReactiveEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reactiveEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCompanionAgentCesToolSpecsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCompanionAgentCesToolSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCompanionAgentCesToolSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentCesToolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCompanionAgentCesToolSpecsOutputReference_Override(d DialogflowCompanionAgentCesToolSpecsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentCesToolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetCesTool(val *string) {
	if err := j.validateSetCesToolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cesTool",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetConfirmationRequirement(val *string) {
	if err := j.validateSetConfirmationRequirementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confirmationRequirement",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetProactiveEnabled(val interface{}) {
	if err := j.validateSetProactiveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proactiveEnabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetReactiveEnabled(val interface{}) {
	if err := j.validateSetReactiveEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reactiveEnabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) PutParameterOverrides(value interface{}) {
	if err := d.validatePutParameterOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParameterOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ResetCesTool() {
	_jsii_.InvokeVoid(
		d,
		"resetCesTool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ResetConfirmationRequirement() {
	_jsii_.InvokeVoid(
		d,
		"resetConfirmationRequirement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ResetParameterOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ResetProactiveEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetProactiveEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ResetReactiveEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetReactiveEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentCesToolSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

