// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowgenerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowGeneratorAgentCoachingContextInstructionsOutputReference interface {
	cdktn.ComplexObject
	AgentAction() *string
	SetAgentAction(val *string)
	AgentActionInput() *string
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayDetails() *string
	SetDisplayDetails(val *string)
	DisplayDetailsInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SystemAction() *string
	SetSystemAction(val *string)
	SystemActionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggeringEvent() *string
	SetTriggeringEvent(val *string)
	TriggeringEventInput() *string
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
	ResetAgentAction()
	ResetCondition()
	ResetDisplayDetails()
	ResetDisplayName()
	ResetSystemAction()
	ResetTriggeringEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowGeneratorAgentCoachingContextInstructionsOutputReference
type jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) AgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) AgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayDetailsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) SystemAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) SystemActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TriggeringEvent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggeringEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TriggeringEventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggeringEventInput",
		&returns,
	)
	return returns
}


func NewDialogflowGeneratorAgentCoachingContextInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowGeneratorAgentCoachingContextInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowGeneratorAgentCoachingContextInstructionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowGenerator.DialogflowGeneratorAgentCoachingContextInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowGeneratorAgentCoachingContextInstructionsOutputReference_Override(d DialogflowGeneratorAgentCoachingContextInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowGenerator.DialogflowGeneratorAgentCoachingContextInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetAgentAction(val *string) {
	if err := j.validateSetAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentAction",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetDisplayDetails(val *string) {
	if err := j.validateSetDisplayDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayDetails",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetSystemAction(val *string) {
	if err := j.validateSetSystemActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemAction",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTriggeringEvent(val *string) {
	if err := j.validateSetTriggeringEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggeringEvent",
		val,
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetAgentAction() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetDisplayDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetSystemAction() {
	_jsii_.InvokeVoid(
		d,
		"resetSystemAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetTriggeringEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggeringEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

