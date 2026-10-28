// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference interface {
	cdktn.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	Env() VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList
	EnvInput() interface{}
	ExecutorImageUri() *string
	SetExecutorImageUri(val *string)
	ExecutorImageUriInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	SetInternalValue(val *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec)
	PackageUris() *[]*string
	SetPackageUris(val *[]*string)
	PackageUrisInput() *[]*string
	PythonModule() *string
	SetPythonModule(val *string)
	PythonModuleInput() *string
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
	ResetArgs()
	ResetEnv()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference
type jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Env() VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList {
	var returns VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ExecutorImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executorImageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ExecutorImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executorImageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InternalValue() *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec {
	var returns *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PackageUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PackageUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"packageUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PythonModule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonModule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PythonModuleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonModuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference_Override(v VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetExecutorImageUri(val *string) {
	if err := j.validateSetExecutorImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executorImageUri",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetInternalValue(val *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetPackageUris(val *[]*string) {
	if err := j.validateSetPackageUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageUris",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetPythonModule(val *string) {
	if err := j.validateSetPythonModuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonModule",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) PutEnv(value interface{}) {
	if err := v.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEnv",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		v,
		"resetArgs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

