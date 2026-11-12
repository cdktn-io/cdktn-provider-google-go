// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference interface {
	cdktn.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	DeploymentTimeout() *string
	SetDeploymentTimeout(val *string)
	DeploymentTimeoutInput() *string
	Env() VertexAiTrainingPipelineModelToUploadContainerSpecEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	GrpcPorts() VertexAiTrainingPipelineModelToUploadContainerSpecGrpcPortsList
	GrpcPortsInput() interface{}
	HealthProbe() VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference
	HealthProbeInput() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe
	HealthRoute() *string
	SetHealthRoute(val *string)
	HealthRouteInput() *string
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InternalValue() *VertexAiTrainingPipelineModelToUploadContainerSpec
	SetInternalValue(val *VertexAiTrainingPipelineModelToUploadContainerSpec)
	Ports() VertexAiTrainingPipelineModelToUploadContainerSpecPortsList
	PortsInput() interface{}
	PredictRoute() *string
	SetPredictRoute(val *string)
	PredictRouteInput() *string
	SharedMemorySizeMb() *string
	SetSharedMemorySizeMb(val *string)
	SharedMemorySizeMbInput() *string
	StartupProbe() VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbeOutputReference
	StartupProbeInput() *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe
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
	PutEnv(value interface{})
	PutGrpcPorts(value interface{})
	PutHealthProbe(value *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe)
	PutPorts(value interface{})
	PutStartupProbe(value *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe)
	ResetArgs()
	ResetCommand()
	ResetDeploymentTimeout()
	ResetEnv()
	ResetGrpcPorts()
	ResetHealthProbe()
	ResetHealthRoute()
	ResetPorts()
	ResetPredictRoute()
	ResetSharedMemorySizeMb()
	ResetStartupProbe()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference
type jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) DeploymentTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) DeploymentTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Env() VertexAiTrainingPipelineModelToUploadContainerSpecEnvList {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GrpcPorts() VertexAiTrainingPipelineModelToUploadContainerSpecGrpcPortsList {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecGrpcPortsList
	_jsii_.Get(
		j,
		"grpcPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GrpcPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) HealthProbe() VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbeOutputReference
	_jsii_.Get(
		j,
		"healthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) HealthProbeInput() *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe
	_jsii_.Get(
		j,
		"healthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) HealthRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) HealthRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) InternalValue() *VertexAiTrainingPipelineModelToUploadContainerSpec {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Ports() VertexAiTrainingPipelineModelToUploadContainerSpecPortsList {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecPortsList
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PredictRoute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PredictRouteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) SharedMemorySizeMb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) SharedMemorySizeMbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedMemorySizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) StartupProbe() VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbeOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbeOutputReference
	_jsii_.Get(
		j,
		"startupProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) StartupProbeInput() *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe {
	var returns *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe
	_jsii_.Get(
		j,
		"startupProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiTrainingPipelineModelToUploadContainerSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineModelToUploadContainerSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiTrainingPipelineModelToUploadContainerSpecOutputReference_Override(v VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetDeploymentTimeout(val *string) {
	if err := j.validateSetDeploymentTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentTimeout",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetHealthRoute(val *string) {
	if err := j.validateSetHealthRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthRoute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetInternalValue(val *VertexAiTrainingPipelineModelToUploadContainerSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetPredictRoute(val *string) {
	if err := j.validateSetPredictRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictRoute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetSharedMemorySizeMb(val *string) {
	if err := j.validateSetSharedMemorySizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySizeMb",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PutEnv(value interface{}) {
	if err := v.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEnv",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PutGrpcPorts(value interface{}) {
	if err := v.validatePutGrpcPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGrpcPorts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PutHealthProbe(value *VertexAiTrainingPipelineModelToUploadContainerSpecHealthProbe) {
	if err := v.validatePutHealthProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putHealthProbe",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PutPorts(value interface{}) {
	if err := v.validatePutPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPorts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) PutStartupProbe(value *VertexAiTrainingPipelineModelToUploadContainerSpecStartupProbe) {
	if err := v.validatePutStartupProbeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStartupProbe",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		v,
		"resetArgs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		v,
		"resetCommand",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetDeploymentTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetDeploymentTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetGrpcPorts() {
	_jsii_.InvokeVoid(
		v,
		"resetGrpcPorts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetHealthProbe() {
	_jsii_.InvokeVoid(
		v,
		"resetHealthProbe",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetHealthRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetHealthRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		v,
		"resetPorts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetPredictRoute() {
	_jsii_.InvokeVoid(
		v,
		"resetPredictRoute",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetSharedMemorySizeMb() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedMemorySizeMb",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ResetStartupProbe() {
	_jsii_.InvokeVoid(
		v,
		"resetStartupProbe",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipelineModelToUploadContainerSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

