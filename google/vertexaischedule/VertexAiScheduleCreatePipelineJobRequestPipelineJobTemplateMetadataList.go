// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
type jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList_Override(v VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) Get(index *float64) VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

