// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference interface {
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
	FailurePolicy() *string
	SetFailurePolicy(val *string)
	FailurePolicyInput() *string
	// Experimental.
	Fqn() *string
	GcsOutputDirectory() *string
	SetGcsOutputDirectory(val *string)
	GcsOutputDirectoryInput() *string
	InputArtifacts() *map[string]*string
	SetInputArtifacts(val *map[string]*string)
	InputArtifactsInput() *map[string]*string
	InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig)
	ParameterValues() *map[string]*string
	SetParameterValues(val *map[string]*string)
	ParameterValuesInput() *map[string]*string
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
	ResetFailurePolicy()
	ResetInputArtifacts()
	ResetParameterValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
type jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GcsOutputDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GcsOutputDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InputArtifacts() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InputArtifactsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ParameterValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ParameterValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference_Override(v VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetGcsOutputDirectory(val *string) {
	if err := j.validateSetGcsOutputDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputDirectory",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetInputArtifacts(val *map[string]*string) {
	if err := j.validateSetInputArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputArtifacts",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetParameterValues(val *map[string]*string) {
	if err := j.validateSetParameterValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterValues",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		v,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ResetParameterValues() {
	_jsii_.InvokeVoid(
		v,
		"resetParameterValues",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

