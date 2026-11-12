// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowcompanionagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference interface {
	cdktn.ComplexObject
	Actions() DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsActionsList
	ActionsInput() interface{}
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
	DisableSuggestedReply() interface{}
	SetDisableSuggestedReply(val interface{})
	DisableSuggestedReplyInput() interface{}
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggerEvent() *string
	SetTriggerEvent(val *string)
	TriggerEventInput() *string
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
	PutActions(value interface{})
	ResetActions()
	ResetCondition()
	ResetDisableSuggestedReply()
	ResetDisplayDetails()
	ResetDisplayName()
	ResetTriggerEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference
type jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) Actions() DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsActionsList {
	var returns DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsActionsList
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisableSuggestedReply() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSuggestedReply",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisableSuggestedReplyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSuggestedReplyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisplayDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisplayDetailsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) TriggerEvent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) TriggerEventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerEventInput",
		&returns,
	)
	return returns
}


func NewDialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference_Override(d DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetDisableSuggestedReply(val interface{}) {
	if err := j.validateSetDisableSuggestedReplyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSuggestedReply",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetDisplayDetails(val *string) {
	if err := j.validateSetDisplayDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayDetails",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference)SetTriggerEvent(val *string) {
	if err := j.validateSetTriggerEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerEvent",
		val,
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) PutActions(value interface{}) {
	if err := d.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		d,
		"resetActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetDisableSuggestedReply() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableSuggestedReply",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetDisplayDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ResetTriggerEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggerEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigGuidanceInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

