// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/colabschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference interface {
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
	EncryptionSpec() ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	EncryptionSpecInput() *ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	EndTime() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ColabScheduleCreatePipelineJobRequestPipelineJob
	SetInternalValue(val *ColabScheduleCreatePipelineJobRequestPipelineJob)
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
	PscInterfaceConfig() ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	RuntimeConfig() ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	RuntimeConfigInput() *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	ScheduleName() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	StartTime() *string
	State() *string
	TemplateMetadata() ColabScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
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
	PutEncryptionSpec(value *ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec)
	PutPscInterfaceConfig(value *ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig)
	PutRuntimeConfig(value *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig)
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

// The jsii proxy struct for ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference
type jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpec() ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference {
	var returns ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpecInput() *ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec {
	var returns *ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) InternalValue() *ColabScheduleCreatePipelineJobRequestPipelineJob {
	var returns *ColabScheduleCreatePipelineJobRequestPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfig() ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference {
	var returns ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfigInput() *ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig {
	var returns *ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfig() ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference {
	var returns ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfigInput() *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig {
	var returns *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"runtimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateMetadata() ColabScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList {
	var returns ColabScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
	_jsii_.Get(
		j,
		"templateMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewColabScheduleCreatePipelineJobRequestPipelineJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewColabScheduleCreatePipelineJobRequestPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.colabSchedule.ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewColabScheduleCreatePipelineJobRequestPipelineJobOutputReference_Override(c ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.colabSchedule.ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetInternalValue(val *ColabScheduleCreatePipelineJobRequestPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPipelineSpec(val *string) {
	if err := j.validateSetPipelineSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineSpec",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPreflightValidations(val interface{}) {
	if err := j.validateSetPreflightValidationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preflightValidations",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTemplateUri(val *string) {
	if err := j.validateSetTemplateUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUri",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutEncryptionSpec(value *ColabScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec) {
	if err := c.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutPscInterfaceConfig(value *ColabScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig) {
	if err := c.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPscInterfaceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutRuntimeConfig(value *ColabScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig) {
	if err := c.validatePutRuntimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRuntimeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		c,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPipelineSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetPipelineSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPreflightValidations() {
	_jsii_.InvokeVoid(
		c,
		"resetPreflightValidations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetRuntimeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRuntimeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetTemplateUri() {
	_jsii_.InvokeVoid(
		c,
		"resetTemplateUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreatePipelineJobRequestPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

