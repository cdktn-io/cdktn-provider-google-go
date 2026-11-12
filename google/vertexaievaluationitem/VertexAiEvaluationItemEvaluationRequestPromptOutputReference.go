// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaievaluationitem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiEvaluationItemEvaluationRequestPromptOutputReference interface {
	cdktn.ComplexObject
	AgentData() *string
	SetAgentData(val *string)
	AgentDataInput() *string
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
	InternalValue() *VertexAiEvaluationItemEvaluationRequestPrompt
	SetInternalValue(val *VertexAiEvaluationItemEvaluationRequestPrompt)
	PromptTemplateData() VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateDataOutputReference
	PromptTemplateDataInput() *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextInput() *string
	UserScenario() VertexAiEvaluationItemEvaluationRequestPromptUserScenarioOutputReference
	UserScenarioInput() *VertexAiEvaluationItemEvaluationRequestPromptUserScenario
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	PutPromptTemplateData(value *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData)
	PutUserScenario(value *VertexAiEvaluationItemEvaluationRequestPromptUserScenario)
	ResetAgentData()
	ResetPromptTemplateData()
	ResetText()
	ResetUserScenario()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEvaluationItemEvaluationRequestPromptOutputReference
type jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) AgentData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) AgentDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) InternalValue() *VertexAiEvaluationItemEvaluationRequestPrompt {
	var returns *VertexAiEvaluationItemEvaluationRequestPrompt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) PromptTemplateData() VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateDataOutputReference {
	var returns VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateDataOutputReference
	_jsii_.Get(
		j,
		"promptTemplateData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) PromptTemplateDataInput() *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData {
	var returns *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData
	_jsii_.Get(
		j,
		"promptTemplateDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) UserScenario() VertexAiEvaluationItemEvaluationRequestPromptUserScenarioOutputReference {
	var returns VertexAiEvaluationItemEvaluationRequestPromptUserScenarioOutputReference
	_jsii_.Get(
		j,
		"userScenario",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) UserScenarioInput() *VertexAiEvaluationItemEvaluationRequestPromptUserScenario {
	var returns *VertexAiEvaluationItemEvaluationRequestPromptUserScenario
	_jsii_.Get(
		j,
		"userScenarioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewVertexAiEvaluationItemEvaluationRequestPromptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiEvaluationItemEvaluationRequestPromptOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEvaluationItemEvaluationRequestPromptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEvaluationItemEvaluationRequestPromptOutputReference_Override(v VertexAiEvaluationItemEvaluationRequestPromptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestPromptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetAgentData(val *string) {
	if err := j.validateSetAgentDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentData",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetInternalValue(val *VertexAiEvaluationItemEvaluationRequestPrompt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) PutPromptTemplateData(value *VertexAiEvaluationItemEvaluationRequestPromptPromptTemplateData) {
	if err := v.validatePutPromptTemplateDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPromptTemplateData",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) PutUserScenario(value *VertexAiEvaluationItemEvaluationRequestPromptUserScenario) {
	if err := v.validatePutUserScenarioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putUserScenario",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ResetAgentData() {
	_jsii_.InvokeVoid(
		v,
		"resetAgentData",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ResetPromptTemplateData() {
	_jsii_.InvokeVoid(
		v,
		"resetPromptTemplateData",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		v,
		"resetText",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ResetUserScenario() {
	_jsii_.InvokeVoid(
		v,
		"resetUserScenario",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		v,
		"resetValue",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestPromptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

