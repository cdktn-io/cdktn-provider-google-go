// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/geminidataanalyticsdataagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference interface {
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
	DatasourceReferences() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesOutputReference
	DatasourceReferencesInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences
	ExampleQueries() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextExampleQueriesList
	ExampleQueriesInput() interface{}
	// Experimental.
	Fqn() *string
	GlossaryTerms() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextGlossaryTermsList
	GlossaryTermsInput() interface{}
	InternalValue() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext
	SetInternalValue(val *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext)
	LookerGoldenQueries() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesList
	LookerGoldenQueriesInput() interface{}
	Options() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference
	OptionsInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions
	SchemaRelationships() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextSchemaRelationshipsList
	SchemaRelationshipsInput() interface{}
	SystemInstruction() *string
	SetSystemInstruction(val *string)
	SystemInstructionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserFunctions() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsOutputReference
	UserFunctionsInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctions
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
	PutDatasourceReferences(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences)
	PutExampleQueries(value interface{})
	PutGlossaryTerms(value interface{})
	PutLookerGoldenQueries(value interface{})
	PutOptions(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions)
	PutSchemaRelationships(value interface{})
	PutUserFunctions(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctions)
	ResetExampleQueries()
	ResetGlossaryTerms()
	ResetLookerGoldenQueries()
	ResetOptions()
	ResetSchemaRelationships()
	ResetSystemInstruction()
	ResetUserFunctions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference
type jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) DatasourceReferences() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesOutputReference {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesOutputReference
	_jsii_.Get(
		j,
		"datasourceReferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) DatasourceReferencesInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences {
	var returns *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences
	_jsii_.Get(
		j,
		"datasourceReferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ExampleQueries() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextExampleQueriesList {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextExampleQueriesList
	_jsii_.Get(
		j,
		"exampleQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ExampleQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exampleQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GlossaryTerms() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextGlossaryTermsList {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextGlossaryTermsList
	_jsii_.Get(
		j,
		"glossaryTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GlossaryTermsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glossaryTermsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) InternalValue() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext {
	var returns *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) LookerGoldenQueries() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesList {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesList
	_jsii_.Get(
		j,
		"lookerGoldenQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) LookerGoldenQueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lookerGoldenQueriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) Options() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) OptionsInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions {
	var returns *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) SchemaRelationships() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextSchemaRelationshipsList {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextSchemaRelationshipsList
	_jsii_.Get(
		j,
		"schemaRelationships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) SchemaRelationshipsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRelationshipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) SystemInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) SystemInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) UserFunctions() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsOutputReference {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctionsOutputReference
	_jsii_.Get(
		j,
		"userFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) UserFunctionsInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctions {
	var returns *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctions
	_jsii_.Get(
		j,
		"userFunctionsInput",
		&returns,
	)
	return returns
}


func NewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference {
	_init_.Initialize()

	if err := validateNewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.geminiDataAnalyticsDataAgent.GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference_Override(g GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.geminiDataAnalyticsDataAgent.GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetInternalValue(val *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetSystemInstruction(val *string) {
	if err := j.validateSetSystemInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemInstruction",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutDatasourceReferences(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferences) {
	if err := g.validatePutDatasourceReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatasourceReferences",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutExampleQueries(value interface{}) {
	if err := g.validatePutExampleQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExampleQueries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutGlossaryTerms(value interface{}) {
	if err := g.validatePutGlossaryTermsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGlossaryTerms",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutLookerGoldenQueries(value interface{}) {
	if err := g.validatePutLookerGoldenQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLookerGoldenQueries",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutOptions(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions) {
	if err := g.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutSchemaRelationships(value interface{}) {
	if err := g.validatePutSchemaRelationshipsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaRelationships",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) PutUserFunctions(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextUserFunctions) {
	if err := g.validatePutUserFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserFunctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetExampleQueries() {
	_jsii_.InvokeVoid(
		g,
		"resetExampleQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetGlossaryTerms() {
	_jsii_.InvokeVoid(
		g,
		"resetGlossaryTerms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetLookerGoldenQueries() {
	_jsii_.InvokeVoid(
		g,
		"resetLookerGoldenQueries",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetSchemaRelationships() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaRelationships",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetSystemInstruction() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemInstruction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ResetUserFunctions() {
	_jsii_.InvokeVoid(
		g,
		"resetUserFunctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

