// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexairagfile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference interface {
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
	InternalValue() *VertexAiRagFileRagFileParsingConfigLayoutParser
	SetInternalValue(val *VertexAiRagFileRagFileParsingConfigLayoutParser)
	MaxParsingRequestsPerMin() *float64
	SetMaxParsingRequestsPerMin(val *float64)
	MaxParsingRequestsPerMinInput() *float64
	ProcessorName() *string
	SetProcessorName(val *string)
	ProcessorNameInput() *string
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
	ResetMaxParsingRequestsPerMin()
	ResetProcessorName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference
type jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) InternalValue() *VertexAiRagFileRagFileParsingConfigLayoutParser {
	var returns *VertexAiRagFileRagFileParsingConfigLayoutParser
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) MaxParsingRequestsPerMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) MaxParsingRequestsPerMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ProcessorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ProcessorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiRagFileRagFileParsingConfigLayoutParserOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiRagFileRagFileParsingConfigLayoutParserOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagFile.VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiRagFileRagFileParsingConfigLayoutParserOutputReference_Override(v VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagFile.VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetInternalValue(val *VertexAiRagFileRagFileParsingConfigLayoutParser) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetMaxParsingRequestsPerMin(val *float64) {
	if err := j.validateSetMaxParsingRequestsPerMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParsingRequestsPerMin",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetProcessorName(val *string) {
	if err := j.validateSetProcessorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorName",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ResetMaxParsingRequestsPerMin() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxParsingRequestsPerMin",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ResetProcessorName() {
	_jsii_.InvokeVoid(
		v,
		"resetProcessorName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiRagFileRagFileParsingConfigLayoutParserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

