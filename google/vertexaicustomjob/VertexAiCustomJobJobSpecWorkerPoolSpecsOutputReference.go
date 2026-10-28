// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference interface {
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
	ContainerSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference
	ContainerSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference
	DiskSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LustreMounts() VertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList
	LustreMountsInput() interface{}
	MachineSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference
	MachineSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec
	NfsMounts() VertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList
	NfsMountsInput() interface{}
	PythonPackageSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
	PythonPackageSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
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
	PutContainerSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec)
	PutDiskSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec)
	PutLustreMounts(value interface{})
	PutMachineSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec)
	PutNfsMounts(value interface{})
	PutPythonPackageSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec)
	ResetContainerSpec()
	ResetDiskSpec()
	ResetLustreMounts()
	ResetMachineSpec()
	ResetNfsMounts()
	ResetPythonPackageSpec()
	ResetReplicaCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference
type jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ContainerSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecOutputReference
	_jsii_.Get(
		j,
		"containerSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ContainerSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec {
	var returns *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec
	_jsii_.Get(
		j,
		"containerSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) DiskSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpecOutputReference
	_jsii_.Get(
		j,
		"diskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) DiskSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec {
	var returns *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec
	_jsii_.Get(
		j,
		"diskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) LustreMounts() VertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsLustreMountsList
	_jsii_.Get(
		j,
		"lustreMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) LustreMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lustreMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) MachineSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) MachineSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec {
	var returns *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) NfsMounts() VertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsNfsMountsList
	_jsii_.Get(
		j,
		"nfsMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) NfsMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsMountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PythonPackageSpec() VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
	_jsii_.Get(
		j,
		"pythonPackageSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PythonPackageSpecInput() *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec {
	var returns *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	_jsii_.Get(
		j,
		"pythonPackageSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference_Override(v VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutContainerSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec) {
	if err := v.validatePutContainerSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putContainerSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutDiskSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec) {
	if err := v.validatePutDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDiskSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutLustreMounts(value interface{}) {
	if err := v.validatePutLustreMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putLustreMounts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutMachineSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec) {
	if err := v.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutNfsMounts(value interface{}) {
	if err := v.validatePutNfsMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putNfsMounts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) PutPythonPackageSpec(value *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec) {
	if err := v.validatePutPythonPackageSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPythonPackageSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetContainerSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetContainerSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetDiskSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetDiskSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetLustreMounts() {
	_jsii_.InvokeVoid(
		v,
		"resetLustreMounts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetNfsMounts() {
	_jsii_.InvokeVoid(
		v,
		"resetNfsMounts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetPythonPackageSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetPythonPackageSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		v,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

