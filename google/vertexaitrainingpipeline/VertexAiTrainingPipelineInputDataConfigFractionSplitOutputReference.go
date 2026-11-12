// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference interface {
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
	InternalValue() *VertexAiTrainingPipelineInputDataConfigFractionSplit
	SetInternalValue(val *VertexAiTrainingPipelineInputDataConfigFractionSplit)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TestFraction() *float64
	SetTestFraction(val *float64)
	TestFractionInput() *float64
	TrainingFraction() *float64
	SetTrainingFraction(val *float64)
	TrainingFractionInput() *float64
	ValidationFraction() *float64
	SetValidationFraction(val *float64)
	ValidationFractionInput() *float64
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
	ResetTestFraction()
	ResetTrainingFraction()
	ResetValidationFraction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference
type jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) InternalValue() *VertexAiTrainingPipelineInputDataConfigFractionSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigFractionSplit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TestFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"testFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TestFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"testFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TrainingFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trainingFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) TrainingFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trainingFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ValidationFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validationFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ValidationFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validationFractionInput",
		&returns,
	)
	return returns
}


func NewVertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineInputDataConfigFractionSplitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference_Override(v VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetInternalValue(val *VertexAiTrainingPipelineInputDataConfigFractionSplit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetTestFraction(val *float64) {
	if err := j.validateSetTestFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"testFraction",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetTrainingFraction(val *float64) {
	if err := j.validateSetTrainingFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingFraction",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference)SetValidationFraction(val *float64) {
	if err := j.validateSetValidationFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationFraction",
		val,
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ResetTestFraction() {
	_jsii_.InvokeVoid(
		v,
		"resetTestFraction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ResetTrainingFraction() {
	_jsii_.InvokeVoid(
		v,
		"resetTrainingFraction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ResetValidationFraction() {
	_jsii_.InvokeVoid(
		v,
		"resetValidationFraction",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

