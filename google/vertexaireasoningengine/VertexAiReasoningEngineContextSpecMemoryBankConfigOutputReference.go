// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexaireasoningengine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference interface {
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
	CustomizationConfigs() VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsList
	CustomizationConfigsInput() interface{}
	DisableMemoryRevisions() interface{}
	SetDisableMemoryRevisions(val interface{})
	DisableMemoryRevisionsInput() interface{}
	// Experimental.
	Fqn() *string
	GenerationConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference
	GenerationConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig
	InternalValue() *VertexAiReasoningEngineContextSpecMemoryBankConfig
	SetInternalValue(val *VertexAiReasoningEngineContextSpecMemoryBankConfig)
	SimilaritySearchConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference
	SimilaritySearchConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig
	StructuredMemoryConfigs() VertexAiReasoningEngineContextSpecMemoryBankConfigStructuredMemoryConfigsList
	StructuredMemoryConfigsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TtlConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference
	TtlConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig
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
	PutCustomizationConfigs(value interface{})
	PutGenerationConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig)
	PutSimilaritySearchConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig)
	PutStructuredMemoryConfigs(value interface{})
	PutTtlConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig)
	ResetCustomizationConfigs()
	ResetDisableMemoryRevisions()
	ResetGenerationConfig()
	ResetSimilaritySearchConfig()
	ResetStructuredMemoryConfigs()
	ResetTtlConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference
type jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) CustomizationConfigs() VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsList {
	var returns VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsList
	_jsii_.Get(
		j,
		"customizationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) CustomizationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customizationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) DisableMemoryRevisions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableMemoryRevisions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) DisableMemoryRevisionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableMemoryRevisionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GenerationConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference {
	var returns VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigOutputReference
	_jsii_.Get(
		j,
		"generationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GenerationConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig {
	var returns *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig
	_jsii_.Get(
		j,
		"generationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InternalValue() *VertexAiReasoningEngineContextSpecMemoryBankConfig {
	var returns *VertexAiReasoningEngineContextSpecMemoryBankConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) SimilaritySearchConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference {
	var returns VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfigOutputReference
	_jsii_.Get(
		j,
		"similaritySearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) SimilaritySearchConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig {
	var returns *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig
	_jsii_.Get(
		j,
		"similaritySearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) StructuredMemoryConfigs() VertexAiReasoningEngineContextSpecMemoryBankConfigStructuredMemoryConfigsList {
	var returns VertexAiReasoningEngineContextSpecMemoryBankConfigStructuredMemoryConfigsList
	_jsii_.Get(
		j,
		"structuredMemoryConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) StructuredMemoryConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"structuredMemoryConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TtlConfig() VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference {
	var returns VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigOutputReference
	_jsii_.Get(
		j,
		"ttlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) TtlConfigInput() *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig {
	var returns *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig
	_jsii_.Get(
		j,
		"ttlConfigInput",
		&returns,
	)
	return returns
}


func NewVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference_Override(v VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetDisableMemoryRevisions(val interface{}) {
	if err := j.validateSetDisableMemoryRevisionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableMemoryRevisions",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetInternalValue(val *VertexAiReasoningEngineContextSpecMemoryBankConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutCustomizationConfigs(value interface{}) {
	if err := v.validatePutCustomizationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomizationConfigs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutGenerationConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig) {
	if err := v.validatePutGenerationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGenerationConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutSimilaritySearchConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig) {
	if err := v.validatePutSimilaritySearchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSimilaritySearchConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutStructuredMemoryConfigs(value interface{}) {
	if err := v.validatePutStructuredMemoryConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStructuredMemoryConfigs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) PutTtlConfig(value *VertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig) {
	if err := v.validatePutTtlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTtlConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetCustomizationConfigs() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomizationConfigs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetDisableMemoryRevisions() {
	_jsii_.InvokeVoid(
		v,
		"resetDisableMemoryRevisions",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetGenerationConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetGenerationConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetSimilaritySearchConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetSimilaritySearchConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetStructuredMemoryConfigs() {
	_jsii_.InvokeVoid(
		v,
		"resetStructuredMemoryConfigs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ResetTtlConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetTtlConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

