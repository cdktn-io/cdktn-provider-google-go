// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dialogflowconversationprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference interface {
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
	EndOfSpeechSensitivity() *string
	SetEndOfSpeechSensitivity(val *string)
	EndOfSpeechSensitivityInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowConversationProfileSttConfigGeminiAsrConfig
	SetInternalValue(val *DialogflowConversationProfileSttConfigGeminiAsrConfig)
	ModelId() *string
	SetModelId(val *string)
	ModelIdInput() *string
	PrefixPaddingMs() *float64
	SetPrefixPaddingMs(val *float64)
	PrefixPaddingMsInput() *float64
	SilenceDurationMs() *float64
	SetSilenceDurationMs(val *float64)
	SilenceDurationMsInput() *float64
	StartOfSpeechSensitivity() *string
	SetStartOfSpeechSensitivity(val *string)
	StartOfSpeechSensitivityInput() *string
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
	ResetEndOfSpeechSensitivity()
	ResetModelId()
	ResetPrefixPaddingMs()
	ResetSilenceDurationMs()
	ResetStartOfSpeechSensitivity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference
type jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) EndOfSpeechSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfSpeechSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) EndOfSpeechSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfSpeechSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InternalValue() *DialogflowConversationProfileSttConfigGeminiAsrConfig {
	var returns *DialogflowConversationProfileSttConfigGeminiAsrConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) PrefixPaddingMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixPaddingMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) PrefixPaddingMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixPaddingMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) SilenceDurationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"silenceDurationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) SilenceDurationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"silenceDurationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) StartOfSpeechSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startOfSpeechSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) StartOfSpeechSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startOfSpeechSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowConversationProfile.DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference_Override(d DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowConversationProfile.DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetEndOfSpeechSensitivity(val *string) {
	if err := j.validateSetEndOfSpeechSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endOfSpeechSensitivity",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetInternalValue(val *DialogflowConversationProfileSttConfigGeminiAsrConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetPrefixPaddingMs(val *float64) {
	if err := j.validateSetPrefixPaddingMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixPaddingMs",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetSilenceDurationMs(val *float64) {
	if err := j.validateSetSilenceDurationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"silenceDurationMs",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetStartOfSpeechSensitivity(val *string) {
	if err := j.validateSetStartOfSpeechSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startOfSpeechSensitivity",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetEndOfSpeechSensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetEndOfSpeechSensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		d,
		"resetModelId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetPrefixPaddingMs() {
	_jsii_.InvokeVoid(
		d,
		"resetPrefixPaddingMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetSilenceDurationMs() {
	_jsii_.InvokeVoid(
		d,
		"resetSilenceDurationMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetStartOfSpeechSensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetStartOfSpeechSensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

