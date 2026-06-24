// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference interface {
	cdktn.ComplexObject
	CommitSha() *string
	SetCommitSha(val *string)
	CommitShaInput() *string
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
	DataformRepositoryResourceName() *string
	SetDataformRepositoryResourceName(val *string)
	DataformRepositoryResourceNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	SetInternalValue(val *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource)
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
	ResetCommitSha()
	ResetDataformRepositoryResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
type jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) CommitSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) CommitShaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitShaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) DataformRepositoryResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataformRepositoryResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) DataformRepositoryResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataformRepositoryResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) InternalValue() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference_Override(v VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetCommitSha(val *string) {
	if err := j.validateSetCommitShaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitSha",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetDataformRepositoryResourceName(val *string) {
	if err := j.validateSetDataformRepositoryResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataformRepositoryResourceName",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetInternalValue(val *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ResetCommitSha() {
	_jsii_.InvokeVoid(
		v,
		"resetCommitSha",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ResetDataformRepositoryResourceName() {
	_jsii_.InvokeVoid(
		v,
		"resetDataformRepositoryResourceName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

