// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowcompanionagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowCompanionAgentSkillConfigsOutputReference interface {
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
	GuidanceSkillConfig() DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigOutputReference
	GuidanceSkillConfigInput() *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SkillTriggeringEvent() *string
	SetSkillTriggeringEvent(val *string)
	SkillTriggeringEventInput() *string
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
	PutGuidanceSkillConfig(value *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig)
	ResetGuidanceSkillConfig()
	ResetSkillTriggeringEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCompanionAgentSkillConfigsOutputReference
type jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GuidanceSkillConfig() DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigOutputReference {
	var returns DialogflowCompanionAgentSkillConfigsGuidanceSkillConfigOutputReference
	_jsii_.Get(
		j,
		"guidanceSkillConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GuidanceSkillConfigInput() *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig {
	var returns *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig
	_jsii_.Get(
		j,
		"guidanceSkillConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) SkillTriggeringEvent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skillTriggeringEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) SkillTriggeringEventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skillTriggeringEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCompanionAgentSkillConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCompanionAgentSkillConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCompanionAgentSkillConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentSkillConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCompanionAgentSkillConfigsOutputReference_Override(d DialogflowCompanionAgentSkillConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCompanionAgent.DialogflowCompanionAgentSkillConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetSkillTriggeringEvent(val *string) {
	if err := j.validateSetSkillTriggeringEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skillTriggeringEvent",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) PutGuidanceSkillConfig(value *DialogflowCompanionAgentSkillConfigsGuidanceSkillConfig) {
	if err := d.validatePutGuidanceSkillConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGuidanceSkillConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ResetGuidanceSkillConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGuidanceSkillConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ResetSkillTriggeringEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetSkillTriggeringEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCompanionAgentSkillConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

