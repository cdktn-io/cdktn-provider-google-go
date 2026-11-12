// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaievaluationitem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiEvaluationItemEvaluationRequestOutputReference interface {
	cdktn.ComplexObject
	CandidateResponses() VertexAiEvaluationItemEvaluationRequestCandidateResponsesList
	CandidateResponsesInput() interface{}
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
	GoldenResponse() VertexAiEvaluationItemEvaluationRequestGoldenResponseOutputReference
	GoldenResponseInput() *VertexAiEvaluationItemEvaluationRequestGoldenResponse
	InternalValue() *VertexAiEvaluationItemEvaluationRequest
	SetInternalValue(val *VertexAiEvaluationItemEvaluationRequest)
	Prompt() VertexAiEvaluationItemEvaluationRequestPromptOutputReference
	PromptInput() *VertexAiEvaluationItemEvaluationRequestPrompt
	Rubrics() VertexAiEvaluationItemEvaluationRequestRubricsList
	RubricsInput() interface{}
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
	PutCandidateResponses(value interface{})
	PutGoldenResponse(value *VertexAiEvaluationItemEvaluationRequestGoldenResponse)
	PutPrompt(value *VertexAiEvaluationItemEvaluationRequestPrompt)
	PutRubrics(value interface{})
	ResetCandidateResponses()
	ResetGoldenResponse()
	ResetPrompt()
	ResetRubrics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEvaluationItemEvaluationRequestOutputReference
type jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) CandidateResponses() VertexAiEvaluationItemEvaluationRequestCandidateResponsesList {
	var returns VertexAiEvaluationItemEvaluationRequestCandidateResponsesList
	_jsii_.Get(
		j,
		"candidateResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) CandidateResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"candidateResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GoldenResponse() VertexAiEvaluationItemEvaluationRequestGoldenResponseOutputReference {
	var returns VertexAiEvaluationItemEvaluationRequestGoldenResponseOutputReference
	_jsii_.Get(
		j,
		"goldenResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GoldenResponseInput() *VertexAiEvaluationItemEvaluationRequestGoldenResponse {
	var returns *VertexAiEvaluationItemEvaluationRequestGoldenResponse
	_jsii_.Get(
		j,
		"goldenResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) InternalValue() *VertexAiEvaluationItemEvaluationRequest {
	var returns *VertexAiEvaluationItemEvaluationRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) Prompt() VertexAiEvaluationItemEvaluationRequestPromptOutputReference {
	var returns VertexAiEvaluationItemEvaluationRequestPromptOutputReference
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) PromptInput() *VertexAiEvaluationItemEvaluationRequestPrompt {
	var returns *VertexAiEvaluationItemEvaluationRequestPrompt
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) Rubrics() VertexAiEvaluationItemEvaluationRequestRubricsList {
	var returns VertexAiEvaluationItemEvaluationRequestRubricsList
	_jsii_.Get(
		j,
		"rubrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) RubricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rubricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiEvaluationItemEvaluationRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiEvaluationItemEvaluationRequestOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEvaluationItemEvaluationRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEvaluationItemEvaluationRequestOutputReference_Override(v VertexAiEvaluationItemEvaluationRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference)SetInternalValue(val *VertexAiEvaluationItemEvaluationRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) PutCandidateResponses(value interface{}) {
	if err := v.validatePutCandidateResponsesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCandidateResponses",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) PutGoldenResponse(value *VertexAiEvaluationItemEvaluationRequestGoldenResponse) {
	if err := v.validatePutGoldenResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGoldenResponse",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) PutPrompt(value *VertexAiEvaluationItemEvaluationRequestPrompt) {
	if err := v.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPrompt",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) PutRubrics(value interface{}) {
	if err := v.validatePutRubricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRubrics",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ResetCandidateResponses() {
	_jsii_.InvokeVoid(
		v,
		"resetCandidateResponses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ResetGoldenResponse() {
	_jsii_.InvokeVoid(
		v,
		"resetGoldenResponse",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		v,
		"resetPrompt",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ResetRubrics() {
	_jsii_.InvokeVoid(
		v,
		"resetRubrics",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

