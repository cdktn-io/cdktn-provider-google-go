// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineInputDataConfigOutputReference interface {
	cdktn.ComplexObject
	AnnotationSchemaUri() *string
	SetAnnotationSchemaUri(val *string)
	AnnotationSchemaUriInput() *string
	AnnotationsFilter() *string
	SetAnnotationsFilter(val *string)
	AnnotationsFilterInput() *string
	BigqueryDestination() VertexAiTrainingPipelineInputDataConfigBigqueryDestinationOutputReference
	BigqueryDestinationInput() *VertexAiTrainingPipelineInputDataConfigBigqueryDestination
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
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	FilterSplit() VertexAiTrainingPipelineInputDataConfigFilterSplitOutputReference
	FilterSplitInput() *VertexAiTrainingPipelineInputDataConfigFilterSplit
	// Experimental.
	Fqn() *string
	FractionSplit() VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference
	FractionSplitInput() *VertexAiTrainingPipelineInputDataConfigFractionSplit
	GcsDestination() VertexAiTrainingPipelineInputDataConfigGcsDestinationOutputReference
	GcsDestinationInput() *VertexAiTrainingPipelineInputDataConfigGcsDestination
	InternalValue() *VertexAiTrainingPipelineInputDataConfig
	SetInternalValue(val *VertexAiTrainingPipelineInputDataConfig)
	PersistMlUseAssignment() interface{}
	SetPersistMlUseAssignment(val interface{})
	PersistMlUseAssignmentInput() interface{}
	PredefinedSplit() VertexAiTrainingPipelineInputDataConfigPredefinedSplitOutputReference
	PredefinedSplitInput() *VertexAiTrainingPipelineInputDataConfigPredefinedSplit
	SavedQueryId() *string
	SetSavedQueryId(val *string)
	SavedQueryIdInput() *string
	StratifiedSplit() VertexAiTrainingPipelineInputDataConfigStratifiedSplitOutputReference
	StratifiedSplitInput() *VertexAiTrainingPipelineInputDataConfigStratifiedSplit
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimestampSplit() VertexAiTrainingPipelineInputDataConfigTimestampSplitOutputReference
	TimestampSplitInput() *VertexAiTrainingPipelineInputDataConfigTimestampSplit
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
	PutBigqueryDestination(value *VertexAiTrainingPipelineInputDataConfigBigqueryDestination)
	PutFilterSplit(value *VertexAiTrainingPipelineInputDataConfigFilterSplit)
	PutFractionSplit(value *VertexAiTrainingPipelineInputDataConfigFractionSplit)
	PutGcsDestination(value *VertexAiTrainingPipelineInputDataConfigGcsDestination)
	PutPredefinedSplit(value *VertexAiTrainingPipelineInputDataConfigPredefinedSplit)
	PutStratifiedSplit(value *VertexAiTrainingPipelineInputDataConfigStratifiedSplit)
	PutTimestampSplit(value *VertexAiTrainingPipelineInputDataConfigTimestampSplit)
	ResetAnnotationSchemaUri()
	ResetAnnotationsFilter()
	ResetBigqueryDestination()
	ResetFilterSplit()
	ResetFractionSplit()
	ResetGcsDestination()
	ResetPersistMlUseAssignment()
	ResetPredefinedSplit()
	ResetSavedQueryId()
	ResetStratifiedSplit()
	ResetTimestampSplit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiTrainingPipelineInputDataConfigOutputReference
type jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) AnnotationSchemaUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationSchemaUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) AnnotationSchemaUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationSchemaUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) AnnotationsFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) AnnotationsFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) BigqueryDestination() VertexAiTrainingPipelineInputDataConfigBigqueryDestinationOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigBigqueryDestinationOutputReference
	_jsii_.Get(
		j,
		"bigqueryDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) BigqueryDestinationInput() *VertexAiTrainingPipelineInputDataConfigBigqueryDestination {
	var returns *VertexAiTrainingPipelineInputDataConfigBigqueryDestination
	_jsii_.Get(
		j,
		"bigqueryDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) FilterSplit() VertexAiTrainingPipelineInputDataConfigFilterSplitOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigFilterSplitOutputReference
	_jsii_.Get(
		j,
		"filterSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) FilterSplitInput() *VertexAiTrainingPipelineInputDataConfigFilterSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigFilterSplit
	_jsii_.Get(
		j,
		"filterSplitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) FractionSplit() VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigFractionSplitOutputReference
	_jsii_.Get(
		j,
		"fractionSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) FractionSplitInput() *VertexAiTrainingPipelineInputDataConfigFractionSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigFractionSplit
	_jsii_.Get(
		j,
		"fractionSplitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GcsDestination() VertexAiTrainingPipelineInputDataConfigGcsDestinationOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigGcsDestinationOutputReference
	_jsii_.Get(
		j,
		"gcsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GcsDestinationInput() *VertexAiTrainingPipelineInputDataConfigGcsDestination {
	var returns *VertexAiTrainingPipelineInputDataConfigGcsDestination
	_jsii_.Get(
		j,
		"gcsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) InternalValue() *VertexAiTrainingPipelineInputDataConfig {
	var returns *VertexAiTrainingPipelineInputDataConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PersistMlUseAssignment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistMlUseAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PersistMlUseAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistMlUseAssignmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PredefinedSplit() VertexAiTrainingPipelineInputDataConfigPredefinedSplitOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigPredefinedSplitOutputReference
	_jsii_.Get(
		j,
		"predefinedSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PredefinedSplitInput() *VertexAiTrainingPipelineInputDataConfigPredefinedSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigPredefinedSplit
	_jsii_.Get(
		j,
		"predefinedSplitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) SavedQueryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savedQueryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) SavedQueryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"savedQueryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) StratifiedSplit() VertexAiTrainingPipelineInputDataConfigStratifiedSplitOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigStratifiedSplitOutputReference
	_jsii_.Get(
		j,
		"stratifiedSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) StratifiedSplitInput() *VertexAiTrainingPipelineInputDataConfigStratifiedSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigStratifiedSplit
	_jsii_.Get(
		j,
		"stratifiedSplitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) TimestampSplit() VertexAiTrainingPipelineInputDataConfigTimestampSplitOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigTimestampSplitOutputReference
	_jsii_.Get(
		j,
		"timestampSplit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) TimestampSplitInput() *VertexAiTrainingPipelineInputDataConfigTimestampSplit {
	var returns *VertexAiTrainingPipelineInputDataConfigTimestampSplit
	_jsii_.Get(
		j,
		"timestampSplitInput",
		&returns,
	)
	return returns
}


func NewVertexAiTrainingPipelineInputDataConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiTrainingPipelineInputDataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineInputDataConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiTrainingPipelineInputDataConfigOutputReference_Override(v VertexAiTrainingPipelineInputDataConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineInputDataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetAnnotationSchemaUri(val *string) {
	if err := j.validateSetAnnotationSchemaUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationSchemaUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetAnnotationsFilter(val *string) {
	if err := j.validateSetAnnotationsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationsFilter",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetInternalValue(val *VertexAiTrainingPipelineInputDataConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetPersistMlUseAssignment(val interface{}) {
	if err := j.validateSetPersistMlUseAssignmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistMlUseAssignment",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetSavedQueryId(val *string) {
	if err := j.validateSetSavedQueryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"savedQueryId",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutBigqueryDestination(value *VertexAiTrainingPipelineInputDataConfigBigqueryDestination) {
	if err := v.validatePutBigqueryDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBigqueryDestination",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutFilterSplit(value *VertexAiTrainingPipelineInputDataConfigFilterSplit) {
	if err := v.validatePutFilterSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFilterSplit",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutFractionSplit(value *VertexAiTrainingPipelineInputDataConfigFractionSplit) {
	if err := v.validatePutFractionSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFractionSplit",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutGcsDestination(value *VertexAiTrainingPipelineInputDataConfigGcsDestination) {
	if err := v.validatePutGcsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGcsDestination",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutPredefinedSplit(value *VertexAiTrainingPipelineInputDataConfigPredefinedSplit) {
	if err := v.validatePutPredefinedSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPredefinedSplit",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutStratifiedSplit(value *VertexAiTrainingPipelineInputDataConfigStratifiedSplit) {
	if err := v.validatePutStratifiedSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStratifiedSplit",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) PutTimestampSplit(value *VertexAiTrainingPipelineInputDataConfigTimestampSplit) {
	if err := v.validatePutTimestampSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimestampSplit",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetAnnotationSchemaUri() {
	_jsii_.InvokeVoid(
		v,
		"resetAnnotationSchemaUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetAnnotationsFilter() {
	_jsii_.InvokeVoid(
		v,
		"resetAnnotationsFilter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetBigqueryDestination() {
	_jsii_.InvokeVoid(
		v,
		"resetBigqueryDestination",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetFilterSplit() {
	_jsii_.InvokeVoid(
		v,
		"resetFilterSplit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetFractionSplit() {
	_jsii_.InvokeVoid(
		v,
		"resetFractionSplit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetGcsDestination() {
	_jsii_.InvokeVoid(
		v,
		"resetGcsDestination",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetPersistMlUseAssignment() {
	_jsii_.InvokeVoid(
		v,
		"resetPersistMlUseAssignment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetPredefinedSplit() {
	_jsii_.InvokeVoid(
		v,
		"resetPredefinedSplit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetSavedQueryId() {
	_jsii_.InvokeVoid(
		v,
		"resetSavedQueryId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetStratifiedSplit() {
	_jsii_.InvokeVoid(
		v,
		"resetStratifiedSplit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ResetTimestampSplit() {
	_jsii_.InvokeVoid(
		v,
		"resetTimestampSplit",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineInputDataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

