// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiCustomJobJobSpecOutputReference interface {
	cdktn.ComplexObject
	BaseOutputDirectory() VertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference
	BaseOutputDirectoryInput() *VertexAiCustomJobJobSpecBaseOutputDirectory
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
	EnableDashboardAccess() interface{}
	SetEnableDashboardAccess(val interface{})
	EnableDashboardAccessInput() interface{}
	EnableWebAccess() interface{}
	SetEnableWebAccess(val interface{})
	EnableWebAccessInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiCustomJobJobSpec
	SetInternalValue(val *VertexAiCustomJobJobSpec)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	PscInterfaceConfig() VertexAiCustomJobJobSpecPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *VertexAiCustomJobJobSpecPscInterfaceConfig
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	Scheduling() VertexAiCustomJobJobSpecSchedulingOutputReference
	SchedulingInput() *VertexAiCustomJobJobSpecScheduling
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	Tensorboard() *string
	SetTensorboard(val *string)
	TensorboardInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkerPoolSpecs() VertexAiCustomJobJobSpecWorkerPoolSpecsList
	WorkerPoolSpecsInput() interface{}
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
	PutBaseOutputDirectory(value *VertexAiCustomJobJobSpecBaseOutputDirectory)
	PutPscInterfaceConfig(value *VertexAiCustomJobJobSpecPscInterfaceConfig)
	PutScheduling(value *VertexAiCustomJobJobSpecScheduling)
	PutWorkerPoolSpecs(value interface{})
	ResetBaseOutputDirectory()
	ResetEnableDashboardAccess()
	ResetEnableWebAccess()
	ResetNetwork()
	ResetPscInterfaceConfig()
	ResetReservedIpRanges()
	ResetScheduling()
	ResetServiceAccount()
	ResetTensorboard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiCustomJobJobSpecOutputReference
type jsiiProxy_VertexAiCustomJobJobSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) BaseOutputDirectory() VertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference {
	var returns VertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference
	_jsii_.Get(
		j,
		"baseOutputDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) BaseOutputDirectoryInput() *VertexAiCustomJobJobSpecBaseOutputDirectory {
	var returns *VertexAiCustomJobJobSpecBaseOutputDirectory
	_jsii_.Get(
		j,
		"baseOutputDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) EnableDashboardAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDashboardAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) EnableDashboardAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDashboardAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) EnableWebAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWebAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) EnableWebAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWebAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) InternalValue() *VertexAiCustomJobJobSpec {
	var returns *VertexAiCustomJobJobSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PscInterfaceConfig() VertexAiCustomJobJobSpecPscInterfaceConfigOutputReference {
	var returns VertexAiCustomJobJobSpecPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PscInterfaceConfigInput() *VertexAiCustomJobJobSpecPscInterfaceConfig {
	var returns *VertexAiCustomJobJobSpecPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) Scheduling() VertexAiCustomJobJobSpecSchedulingOutputReference {
	var returns VertexAiCustomJobJobSpecSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) SchedulingInput() *VertexAiCustomJobJobSpecScheduling {
	var returns *VertexAiCustomJobJobSpecScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) Tensorboard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tensorboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) TensorboardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tensorboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) WorkerPoolSpecs() VertexAiCustomJobJobSpecWorkerPoolSpecsList {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsList
	_jsii_.Get(
		j,
		"workerPoolSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) WorkerPoolSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerPoolSpecsInput",
		&returns,
	)
	return returns
}


func NewVertexAiCustomJobJobSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiCustomJobJobSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiCustomJobJobSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiCustomJobJobSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiCustomJobJobSpecOutputReference_Override(v VertexAiCustomJobJobSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetEnableDashboardAccess(val interface{}) {
	if err := j.validateSetEnableDashboardAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDashboardAccess",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetEnableWebAccess(val interface{}) {
	if err := j.validateSetEnableWebAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWebAccess",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetInternalValue(val *VertexAiCustomJobJobSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetTensorboard(val *string) {
	if err := j.validateSetTensorboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tensorboard",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PutBaseOutputDirectory(value *VertexAiCustomJobJobSpecBaseOutputDirectory) {
	if err := v.validatePutBaseOutputDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBaseOutputDirectory",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PutPscInterfaceConfig(value *VertexAiCustomJobJobSpecPscInterfaceConfig) {
	if err := v.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPscInterfaceConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PutScheduling(value *VertexAiCustomJobJobSpecScheduling) {
	if err := v.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putScheduling",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) PutWorkerPoolSpecs(value interface{}) {
	if err := v.validatePutWorkerPoolSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putWorkerPoolSpecs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetBaseOutputDirectory() {
	_jsii_.InvokeVoid(
		v,
		"resetBaseOutputDirectory",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetEnableDashboardAccess() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableDashboardAccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetEnableWebAccess() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableWebAccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		v,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		v,
		"resetScheduling",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		v,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ResetTensorboard() {
	_jsii_.InvokeVoid(
		v,
		"resetTensorboard",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

