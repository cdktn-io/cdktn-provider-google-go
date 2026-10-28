// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiCustomJobJobSpecSchedulingOutputReference interface {
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
	DisableRetries() interface{}
	SetDisableRetries(val interface{})
	DisableRetriesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiCustomJobJobSpecScheduling
	SetInternalValue(val *VertexAiCustomJobJobSpecScheduling)
	MaxWaitDuration() *string
	SetMaxWaitDuration(val *string)
	MaxWaitDurationInput() *string
	RestartJobOnWorkerRestart() interface{}
	SetRestartJobOnWorkerRestart(val interface{})
	RestartJobOnWorkerRestartInput() interface{}
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	ResetDisableRetries()
	ResetMaxWaitDuration()
	ResetRestartJobOnWorkerRestart()
	ResetStrategy()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiCustomJobJobSpecSchedulingOutputReference
type jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) DisableRetries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) DisableRetriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) InternalValue() *VertexAiCustomJobJobSpecScheduling {
	var returns *VertexAiCustomJobJobSpecScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) MaxWaitDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWaitDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) MaxWaitDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWaitDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) RestartJobOnWorkerRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartJobOnWorkerRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) RestartJobOnWorkerRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartJobOnWorkerRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewVertexAiCustomJobJobSpecSchedulingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiCustomJobJobSpecSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiCustomJobJobSpecSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiCustomJobJobSpecSchedulingOutputReference_Override(v VertexAiCustomJobJobSpecSchedulingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiCustomJob.VertexAiCustomJobJobSpecSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetDisableRetries(val interface{}) {
	if err := j.validateSetDisableRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRetries",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetInternalValue(val *VertexAiCustomJobJobSpecScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetMaxWaitDuration(val *string) {
	if err := j.validateSetMaxWaitDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWaitDuration",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetRestartJobOnWorkerRestart(val interface{}) {
	if err := j.validateSetRestartJobOnWorkerRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartJobOnWorkerRestart",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ResetDisableRetries() {
	_jsii_.InvokeVoid(
		v,
		"resetDisableRetries",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ResetMaxWaitDuration() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxWaitDuration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ResetRestartJobOnWorkerRestart() {
	_jsii_.InvokeVoid(
		v,
		"resetRestartJobOnWorkerRestart",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		v,
		"resetStrategy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeout",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiCustomJobJobSpecSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

