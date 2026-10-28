// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package geminidataanalyticsdataagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/geminidataanalyticsdataagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LookerQuery() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQueryOutputReference
	LookerQueryInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery
	NaturalLanguageQuestions() *[]*string
	SetNaturalLanguageQuestions(val *[]*string)
	NaturalLanguageQuestionsInput() *[]*string
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
	PutLookerQuery(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery)
	ResetLookerQuery()
	ResetNaturalLanguageQuestions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference
type jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) LookerQuery() GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQueryOutputReference {
	var returns GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQueryOutputReference
	_jsii_.Get(
		j,
		"lookerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) LookerQueryInput() *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery {
	var returns *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery
	_jsii_.Get(
		j,
		"lookerQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) NaturalLanguageQuestions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"naturalLanguageQuestions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) NaturalLanguageQuestionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"naturalLanguageQuestionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference {
	_init_.Initialize()

	if err := validateNewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.geminiDataAnalyticsDataAgent.GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference_Override(g GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.geminiDataAnalyticsDataAgent.GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetNaturalLanguageQuestions(val *[]*string) {
	if err := j.validateSetNaturalLanguageQuestionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naturalLanguageQuestions",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) PutLookerQuery(value *GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery) {
	if err := g.validatePutLookerQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLookerQuery",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ResetLookerQuery() {
	_jsii_.InvokeVoid(
		g,
		"resetLookerQuery",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ResetNaturalLanguageQuestions() {
	_jsii_.InvokeVoid(
		g,
		"resetNaturalLanguageQuestions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

