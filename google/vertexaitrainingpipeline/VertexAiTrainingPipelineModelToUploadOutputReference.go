// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineModelToUploadOutputReference interface {
	cdktn.ComplexObject
	ArtifactUri() *string
	SetArtifactUri(val *string)
	ArtifactUriInput() *string
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
	ContainerSpec() VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference
	ContainerSpecInput() *VertexAiTrainingPipelineModelToUploadContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionSpec() VertexAiTrainingPipelineModelToUploadEncryptionSpecOutputReference
	EncryptionSpecInput() *VertexAiTrainingPipelineModelToUploadEncryptionSpec
	ExplanationSpec() VertexAiTrainingPipelineModelToUploadExplanationSpecOutputReference
	ExplanationSpecInput() *VertexAiTrainingPipelineModelToUploadExplanationSpec
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiTrainingPipelineModelToUpload
	SetInternalValue(val *VertexAiTrainingPipelineModelToUpload)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Name() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VersionAliases() *[]*string
	SetVersionAliases(val *[]*string)
	VersionAliasesInput() *[]*string
	VersionDescription() *string
	SetVersionDescription(val *string)
	VersionDescriptionInput() *string
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
	PutContainerSpec(value *VertexAiTrainingPipelineModelToUploadContainerSpec)
	PutEncryptionSpec(value *VertexAiTrainingPipelineModelToUploadEncryptionSpec)
	PutExplanationSpec(value *VertexAiTrainingPipelineModelToUploadExplanationSpec)
	ResetArtifactUri()
	ResetContainerSpec()
	ResetDescription()
	ResetEncryptionSpec()
	ResetExplanationSpec()
	ResetLabels()
	ResetVersionAliases()
	ResetVersionDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiTrainingPipelineModelToUploadOutputReference
type jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ArtifactUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ArtifactUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ContainerSpec() VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ContainerSpecInput() *VertexAiTrainingPipelineModelToUploadContainerSpec {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) EncryptionSpec() VertexAiTrainingPipelineModelToUploadEncryptionSpecOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) EncryptionSpecInput() *VertexAiTrainingPipelineModelToUploadEncryptionSpec {
	var returns *VertexAiTrainingPipelineModelToUploadEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ExplanationSpec() VertexAiTrainingPipelineModelToUploadExplanationSpecOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadExplanationSpecOutputReference
	_jsii_.Get(
		j,
		"explanationSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ExplanationSpecInput() *VertexAiTrainingPipelineModelToUploadExplanationSpec {
	var returns *VertexAiTrainingPipelineModelToUploadExplanationSpec
	_jsii_.Get(
		j,
		"explanationSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) InternalValue() *VertexAiTrainingPipelineModelToUpload {
	var returns *VertexAiTrainingPipelineModelToUpload
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) VersionAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) VersionAliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) VersionDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescriptionInput",
		&returns,
	)
	return returns
}


func NewVertexAiTrainingPipelineModelToUploadOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiTrainingPipelineModelToUploadOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineModelToUploadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiTrainingPipelineModelToUploadOutputReference_Override(v VertexAiTrainingPipelineModelToUploadOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetArtifactUri(val *string) {
	if err := j.validateSetArtifactUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetInternalValue(val *VertexAiTrainingPipelineModelToUpload) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetVersionAliases(val *[]*string) {
	if err := j.validateSetVersionAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionAliases",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference)SetVersionDescription(val *string) {
	if err := j.validateSetVersionDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) PutContainerSpec(value *VertexAiTrainingPipelineModelToUploadContainerSpec) {
	if err := v.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) PutEncryptionSpec(value *VertexAiTrainingPipelineModelToUploadEncryptionSpec) {
	if err := v.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) PutExplanationSpec(value *VertexAiTrainingPipelineModelToUploadExplanationSpec) {
	if err := v.validatePutExplanationSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putExplanationSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetArtifactUri() {
	_jsii_.InvokeVoid(
		v,
		"resetArtifactUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetExplanationSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetExplanationSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetVersionAliases() {
	_jsii_.InvokeVoid(
		v,
		"resetVersionAliases",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ResetVersionDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetVersionDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

