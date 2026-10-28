// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaionlineevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiOnlineEvaluatorCloudObservabilityOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiOnlineEvaluatorCloudObservability
	SetInternalValue(val *VertexAiOnlineEvaluatorCloudObservability)
	LogView() *string
	SetLogView(val *string)
	LogViewInput() *string
	OpenTelemetry() VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference
	OpenTelemetryInput() *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry
	SessionScope() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference
	SessionScopeInput() *VertexAiOnlineEvaluatorCloudObservabilitySessionScope
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TraceScope() VertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference
	TraceScopeInput() *VertexAiOnlineEvaluatorCloudObservabilityTraceScope
	TraceView() *string
	SetTraceView(val *string)
	TraceViewInput() *string
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
	PutOpenTelemetry(value *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry)
	PutSessionScope(value *VertexAiOnlineEvaluatorCloudObservabilitySessionScope)
	PutTraceScope(value *VertexAiOnlineEvaluatorCloudObservabilityTraceScope)
	ResetLogView()
	ResetSessionScope()
	ResetTraceScope()
	ResetTraceView()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiOnlineEvaluatorCloudObservabilityOutputReference
type jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) InternalValue() *VertexAiOnlineEvaluatorCloudObservability {
	var returns *VertexAiOnlineEvaluatorCloudObservability
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) LogView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) LogViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) OpenTelemetry() VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference {
	var returns VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetryOutputReference
	_jsii_.Get(
		j,
		"openTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) OpenTelemetryInput() *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry {
	var returns *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry
	_jsii_.Get(
		j,
		"openTelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) SessionScope() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference {
	var returns VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference
	_jsii_.Get(
		j,
		"sessionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) SessionScopeInput() *VertexAiOnlineEvaluatorCloudObservabilitySessionScope {
	var returns *VertexAiOnlineEvaluatorCloudObservabilitySessionScope
	_jsii_.Get(
		j,
		"sessionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceScope() VertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference {
	var returns VertexAiOnlineEvaluatorCloudObservabilityTraceScopeOutputReference
	_jsii_.Get(
		j,
		"traceScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceScopeInput() *VertexAiOnlineEvaluatorCloudObservabilityTraceScope {
	var returns *VertexAiOnlineEvaluatorCloudObservabilityTraceScope
	_jsii_.Get(
		j,
		"traceScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) TraceViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceViewInput",
		&returns,
	)
	return returns
}


func NewVertexAiOnlineEvaluatorCloudObservabilityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiOnlineEvaluatorCloudObservabilityOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiOnlineEvaluatorCloudObservabilityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiOnlineEvaluator.VertexAiOnlineEvaluatorCloudObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiOnlineEvaluatorCloudObservabilityOutputReference_Override(v VertexAiOnlineEvaluatorCloudObservabilityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiOnlineEvaluator.VertexAiOnlineEvaluatorCloudObservabilityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetInternalValue(val *VertexAiOnlineEvaluatorCloudObservability) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetLogView(val *string) {
	if err := j.validateSetLogViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logView",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference)SetTraceView(val *string) {
	if err := j.validateSetTraceViewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceView",
		val,
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutOpenTelemetry(value *VertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry) {
	if err := v.validatePutOpenTelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOpenTelemetry",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutSessionScope(value *VertexAiOnlineEvaluatorCloudObservabilitySessionScope) {
	if err := v.validatePutSessionScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSessionScope",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) PutTraceScope(value *VertexAiOnlineEvaluatorCloudObservabilityTraceScope) {
	if err := v.validatePutTraceScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTraceScope",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetLogView() {
	_jsii_.InvokeVoid(
		v,
		"resetLogView",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetSessionScope() {
	_jsii_.InvokeVoid(
		v,
		"resetSessionScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetTraceScope() {
	_jsii_.InvokeVoid(
		v,
		"resetTraceScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ResetTraceView() {
	_jsii_.InvokeVoid(
		v,
		"resetTraceView",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

