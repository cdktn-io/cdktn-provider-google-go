// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaievaluationitem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiEvaluationItemEvaluationRequestCandidateResponsesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) VertexAiEvaluationItemEvaluationRequestCandidateResponsesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEvaluationItemEvaluationRequestCandidateResponsesList
type jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVertexAiEvaluationItemEvaluationRequestCandidateResponsesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VertexAiEvaluationItemEvaluationRequestCandidateResponsesList {
	_init_.Initialize()

	if err := validateNewVertexAiEvaluationItemEvaluationRequestCandidateResponsesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestCandidateResponsesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVertexAiEvaluationItemEvaluationRequestCandidateResponsesList_Override(v VertexAiEvaluationItemEvaluationRequestCandidateResponsesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEvaluationItem.VertexAiEvaluationItemEvaluationRequestCandidateResponsesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := v.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		v,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) Get(index *float64) VertexAiEvaluationItemEvaluationRequestCandidateResponsesOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VertexAiEvaluationItemEvaluationRequestCandidateResponsesOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiEvaluationItemEvaluationRequestCandidateResponsesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

