// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference interface {
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionSpec() VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	EncryptionSpecInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	EndTime() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJob
	SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJob)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	PipelineSpec() *string
	SetPipelineSpec(val *string)
	PipelineSpecInput() *string
	PreflightValidations() interface{}
	SetPreflightValidations(val interface{})
	PreflightValidationsInput() interface{}
	PscInterfaceConfig() VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	RuntimeConfig() VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	RuntimeConfigInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	ScheduleName() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	StartTime() *string
	State() *string
	TemplateMetadata() VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
	TemplateUri() *string
	SetTemplateUri(val *string)
	TemplateUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
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
	PutEncryptionSpec(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec)
	PutPscInterfaceConfig(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig)
	PutRuntimeConfig(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig)
	ResetDisplayName()
	ResetEncryptionSpec()
	ResetLabels()
	ResetNetwork()
	ResetPipelineSpec()
	ResetPreflightValidations()
	ResetPscInterfaceConfig()
	ResetReservedIpRanges()
	ResetRuntimeConfig()
	ResetServiceAccount()
	ResetTemplateUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference
type jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpec() VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference {
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpecInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJob {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfig() VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference {
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfigInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfig() VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference {
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfigInput() *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"runtimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateMetadata() VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList {
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
	_jsii_.Get(
		j,
		"templateMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference_Override(v VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPipelineSpec(val *string) {
	if err := j.validateSetPipelineSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineSpec",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPreflightValidations(val interface{}) {
	if err := j.validateSetPreflightValidationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preflightValidations",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTemplateUri(val *string) {
	if err := j.validateSetTemplateUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutEncryptionSpec(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec) {
	if err := v.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutPscInterfaceConfig(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig) {
	if err := v.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPscInterfaceConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutRuntimeConfig(value *VertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig) {
	if err := v.validatePutRuntimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRuntimeConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		v,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPipelineSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetPipelineSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPreflightValidations() {
	_jsii_.InvokeVoid(
		v,
		"resetPreflightValidations",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		v,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetRuntimeConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetRuntimeConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetTemplateUri() {
	_jsii_.InvokeVoid(
		v,
		"resetTemplateUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

