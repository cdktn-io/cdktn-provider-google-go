// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexairagfile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiRagFileRagFileParsingConfigLlmParserOutputReference interface {
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
	CustomParsingPrompt() *string
	SetCustomParsingPrompt(val *string)
	CustomParsingPromptInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiRagFileRagFileParsingConfigLlmParser
	SetInternalValue(val *VertexAiRagFileRagFileParsingConfigLlmParser)
	MaxParsingRequestsPerMin() *float64
	SetMaxParsingRequestsPerMin(val *float64)
	MaxParsingRequestsPerMinInput() *float64
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
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
	ResetCustomParsingPrompt()
	ResetMaxParsingRequestsPerMin()
	ResetModelName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiRagFileRagFileParsingConfigLlmParserOutputReference
type jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CustomParsingPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customParsingPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CustomParsingPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customParsingPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InternalValue() *VertexAiRagFileRagFileParsingConfigLlmParser {
	var returns *VertexAiRagFileRagFileParsingConfigLlmParser
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) MaxParsingRequestsPerMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) MaxParsingRequestsPerMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiRagFileRagFileParsingConfigLlmParserOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiRagFileRagFileParsingConfigLlmParserOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiRagFileRagFileParsingConfigLlmParserOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagFile.VertexAiRagFileRagFileParsingConfigLlmParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiRagFileRagFileParsingConfigLlmParserOutputReference_Override(v VertexAiRagFileRagFileParsingConfigLlmParserOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagFile.VertexAiRagFileRagFileParsingConfigLlmParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetCustomParsingPrompt(val *string) {
	if err := j.validateSetCustomParsingPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customParsingPrompt",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetInternalValue(val *VertexAiRagFileRagFileParsingConfigLlmParser) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetMaxParsingRequestsPerMin(val *float64) {
	if err := j.validateSetMaxParsingRequestsPerMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParsingRequestsPerMin",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetCustomParsingPrompt() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomParsingPrompt",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetMaxParsingRequestsPerMin() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxParsingRequestsPerMin",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		v,
		"resetModelName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

