// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexairagcorpus/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiRagCorpusVectorDbConfigOutputReference interface {
	cdktn.ComplexObject
	ApiAuth() VertexAiRagCorpusVectorDbConfigApiAuthOutputReference
	ApiAuthInput() *VertexAiRagCorpusVectorDbConfigApiAuth
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
	InternalValue() *VertexAiRagCorpusVectorDbConfig
	SetInternalValue(val *VertexAiRagCorpusVectorDbConfig)
	Pinecone() VertexAiRagCorpusVectorDbConfigPineconeOutputReference
	PineconeInput() *VertexAiRagCorpusVectorDbConfigPinecone
	RagEmbeddingModelConfig() VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference
	RagEmbeddingModelConfigInput() *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig
	RagManagedDb() VertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference
	RagManagedDbInput() *VertexAiRagCorpusVectorDbConfigRagManagedDb
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VertexVectorSearch() VertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference
	VertexVectorSearchInput() *VertexAiRagCorpusVectorDbConfigVertexVectorSearch
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
	PutApiAuth(value *VertexAiRagCorpusVectorDbConfigApiAuth)
	PutPinecone(value *VertexAiRagCorpusVectorDbConfigPinecone)
	PutRagEmbeddingModelConfig(value *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig)
	PutRagManagedDb(value *VertexAiRagCorpusVectorDbConfigRagManagedDb)
	PutVertexVectorSearch(value *VertexAiRagCorpusVectorDbConfigVertexVectorSearch)
	ResetApiAuth()
	ResetPinecone()
	ResetRagEmbeddingModelConfig()
	ResetRagManagedDb()
	ResetVertexVectorSearch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiRagCorpusVectorDbConfigOutputReference
type jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ApiAuth() VertexAiRagCorpusVectorDbConfigApiAuthOutputReference {
	var returns VertexAiRagCorpusVectorDbConfigApiAuthOutputReference
	_jsii_.Get(
		j,
		"apiAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ApiAuthInput() *VertexAiRagCorpusVectorDbConfigApiAuth {
	var returns *VertexAiRagCorpusVectorDbConfigApiAuth
	_jsii_.Get(
		j,
		"apiAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) InternalValue() *VertexAiRagCorpusVectorDbConfig {
	var returns *VertexAiRagCorpusVectorDbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) Pinecone() VertexAiRagCorpusVectorDbConfigPineconeOutputReference {
	var returns VertexAiRagCorpusVectorDbConfigPineconeOutputReference
	_jsii_.Get(
		j,
		"pinecone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PineconeInput() *VertexAiRagCorpusVectorDbConfigPinecone {
	var returns *VertexAiRagCorpusVectorDbConfigPinecone
	_jsii_.Get(
		j,
		"pineconeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) RagEmbeddingModelConfig() VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference {
	var returns VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfigOutputReference
	_jsii_.Get(
		j,
		"ragEmbeddingModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) RagEmbeddingModelConfigInput() *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig {
	var returns *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig
	_jsii_.Get(
		j,
		"ragEmbeddingModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) RagManagedDb() VertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference {
	var returns VertexAiRagCorpusVectorDbConfigRagManagedDbOutputReference
	_jsii_.Get(
		j,
		"ragManagedDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) RagManagedDbInput() *VertexAiRagCorpusVectorDbConfigRagManagedDb {
	var returns *VertexAiRagCorpusVectorDbConfigRagManagedDb
	_jsii_.Get(
		j,
		"ragManagedDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) VertexVectorSearch() VertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference {
	var returns VertexAiRagCorpusVectorDbConfigVertexVectorSearchOutputReference
	_jsii_.Get(
		j,
		"vertexVectorSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) VertexVectorSearchInput() *VertexAiRagCorpusVectorDbConfigVertexVectorSearch {
	var returns *VertexAiRagCorpusVectorDbConfigVertexVectorSearch
	_jsii_.Get(
		j,
		"vertexVectorSearchInput",
		&returns,
	)
	return returns
}


func NewVertexAiRagCorpusVectorDbConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiRagCorpusVectorDbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiRagCorpusVectorDbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagCorpus.VertexAiRagCorpusVectorDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiRagCorpusVectorDbConfigOutputReference_Override(v VertexAiRagCorpusVectorDbConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiRagCorpus.VertexAiRagCorpusVectorDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference)SetInternalValue(val *VertexAiRagCorpusVectorDbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PutApiAuth(value *VertexAiRagCorpusVectorDbConfigApiAuth) {
	if err := v.validatePutApiAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putApiAuth",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PutPinecone(value *VertexAiRagCorpusVectorDbConfigPinecone) {
	if err := v.validatePutPineconeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPinecone",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PutRagEmbeddingModelConfig(value *VertexAiRagCorpusVectorDbConfigRagEmbeddingModelConfig) {
	if err := v.validatePutRagEmbeddingModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRagEmbeddingModelConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PutRagManagedDb(value *VertexAiRagCorpusVectorDbConfigRagManagedDb) {
	if err := v.validatePutRagManagedDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRagManagedDb",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) PutVertexVectorSearch(value *VertexAiRagCorpusVectorDbConfigVertexVectorSearch) {
	if err := v.validatePutVertexVectorSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVertexVectorSearch",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ResetApiAuth() {
	_jsii_.InvokeVoid(
		v,
		"resetApiAuth",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ResetPinecone() {
	_jsii_.InvokeVoid(
		v,
		"resetPinecone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ResetRagEmbeddingModelConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetRagEmbeddingModelConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ResetRagManagedDb() {
	_jsii_.InvokeVoid(
		v,
		"resetRagManagedDb",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ResetVertexVectorSearch() {
	_jsii_.InvokeVoid(
		v,
		"resetVertexVectorSearch",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiRagCorpusVectorDbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

