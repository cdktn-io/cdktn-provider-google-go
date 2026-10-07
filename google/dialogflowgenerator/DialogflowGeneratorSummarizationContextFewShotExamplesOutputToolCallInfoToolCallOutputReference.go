// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/dialogflowgenerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	InternalValue() *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall
	SetInternalValue(val *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tool() *string
	SetTool(val *string)
	ToolInput() *string
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
	ResetAction()
	ResetTool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference
type jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) InternalValue() *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall {
	var returns *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ToolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}


func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference_Override(d DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetInternalValue(val *DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference)SetTool(val *string) {
	if err := j.validateSetToolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tool",
		val,
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		d,
		"resetAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ResetTool() {
	_jsii_.InvokeVoid(
		d,
		"resetTool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

