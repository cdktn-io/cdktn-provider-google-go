// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference interface {
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
	CustomEnvironmentSpec() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
	CustomEnvironmentSpecInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	DataformRepositorySource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	DataformRepositorySourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	DirectNotebookSource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference
	DirectNotebookSourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionSpec() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference
	EncryptionSpecInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
	ExecutionUser() *string
	SetExecutionUser(val *string)
	ExecutionUserInput() *string
	// Experimental.
	Fqn() *string
	GcsNotebookSource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	GcsNotebookSourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	GcsOutputUri() *string
	SetGcsOutputUri(val *string)
	GcsOutputUriInput() *string
	InternalValue() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	SetInternalValue(val *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob)
	JobState() *string
	KernelName() *string
	SetKernelName(val *string)
	KernelNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Name() *string
	NotebookRuntimeTemplateResourceName() *string
	SetNotebookRuntimeTemplateResourceName(val *string)
	NotebookRuntimeTemplateResourceNameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	ScheduleResourceName() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	WorkbenchRuntime() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference
	WorkbenchRuntimeInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime
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
	PutCustomEnvironmentSpec(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec)
	PutDataformRepositorySource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource)
	PutDirectNotebookSource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource)
	PutEncryptionSpec(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec)
	PutGcsNotebookSource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource)
	PutWorkbenchRuntime(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime)
	ResetCustomEnvironmentSpec()
	ResetDataformRepositorySource()
	ResetDirectNotebookSource()
	ResetDisplayName()
	ResetEncryptionSpec()
	ResetExecutionTimeout()
	ResetExecutionUser()
	ResetGcsNotebookSource()
	ResetGcsOutputUri()
	ResetKernelName()
	ResetLabels()
	ResetNotebookRuntimeTemplateResourceName()
	ResetParameters()
	ResetServiceAccount()
	ResetWorkbenchRuntime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference
type jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CustomEnvironmentSpec() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
	_jsii_.Get(
		j,
		"customEnvironmentSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CustomEnvironmentSpecInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"customEnvironmentSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	_jsii_.Get(
		j,
		"dataformRepositorySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	_jsii_.Get(
		j,
		"dataformRepositorySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DirectNotebookSource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"directNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DirectNotebookSourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource
	_jsii_.Get(
		j,
		"directNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) EncryptionSpec() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) EncryptionSpecInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSource() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"gcsNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSourceInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	_jsii_.Get(
		j,
		"gcsNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InternalValue() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) JobState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) KernelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) KernelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ScheduleResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) WorkbenchRuntime() VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference {
	var returns VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference
	_jsii_.Get(
		j,
		"workbenchRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) WorkbenchRuntimeInput() *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime {
	var returns *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime
	_jsii_.Get(
		j,
		"workbenchRuntimeInput",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference_Override(v VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionTimeout(val *string) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionUser(val *string) {
	if err := j.validateSetExecutionUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUser",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetGcsOutputUri(val *string) {
	if err := j.validateSetGcsOutputUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetInternalValue(val *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetKernelName(val *string) {
	if err := j.validateSetKernelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelName",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetNotebookRuntimeTemplateResourceName(val *string) {
	if err := j.validateSetNotebookRuntimeTemplateResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookRuntimeTemplateResourceName",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutCustomEnvironmentSpec(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec) {
	if err := v.validatePutCustomEnvironmentSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomEnvironmentSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDataformRepositorySource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource) {
	if err := v.validatePutDataformRepositorySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDataformRepositorySource",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDirectNotebookSource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource) {
	if err := v.validatePutDirectNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDirectNotebookSource",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutEncryptionSpec(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec) {
	if err := v.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutGcsNotebookSource(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource) {
	if err := v.validatePutGcsNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGcsNotebookSource",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutWorkbenchRuntime(value *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime) {
	if err := v.validatePutWorkbenchRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putWorkbenchRuntime",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetCustomEnvironmentSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomEnvironmentSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDataformRepositorySource() {
	_jsii_.InvokeVoid(
		v,
		"resetDataformRepositorySource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDirectNotebookSource() {
	_jsii_.InvokeVoid(
		v,
		"resetDirectNotebookSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		v,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionUser() {
	_jsii_.InvokeVoid(
		v,
		"resetExecutionUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsNotebookSource() {
	_jsii_.InvokeVoid(
		v,
		"resetGcsNotebookSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsOutputUri() {
	_jsii_.InvokeVoid(
		v,
		"resetGcsOutputUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetKernelName() {
	_jsii_.InvokeVoid(
		v,
		"resetKernelName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetNotebookRuntimeTemplateResourceName() {
	_jsii_.InvokeVoid(
		v,
		"resetNotebookRuntimeTemplateResourceName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		v,
		"resetParameters",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetWorkbenchRuntime() {
	_jsii_.InvokeVoid(
		v,
		"resetWorkbenchRuntime",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

