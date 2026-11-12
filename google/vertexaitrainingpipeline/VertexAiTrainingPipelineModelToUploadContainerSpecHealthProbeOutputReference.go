// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference interface {
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
	Exec() VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExecOutputReference
	ExecInput() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExec
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe
	SetInternalValue(val *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutExec(value *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExec)
	ResetExec()
	ResetPeriodSeconds()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference
type jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) Exec() VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExecOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ExecInput() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExec {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) InternalValue() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewVertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference_Override(v VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetInternalValue(val *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) PutExec(value *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeExec) {
	if err := v.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putExec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		v,
		"resetExec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

