// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaionlineevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexaionlineevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference interface {
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
	Filter() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterList
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	InactivityTrigger() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTriggerOutputReference
	InactivityTriggerInput() *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger
	InternalValue() *VertexAiOnlineEvaluatorCloudObservabilitySessionScope
	SetInternalValue(val *VertexAiOnlineEvaluatorCloudObservabilitySessionScope)
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
	PutFilter(value interface{})
	PutInactivityTrigger(value *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger)
	ResetFilter()
	ResetInactivityTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference
type jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) Filter() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterList {
	var returns VertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) InactivityTrigger() VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTriggerOutputReference {
	var returns VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTriggerOutputReference
	_jsii_.Get(
		j,
		"inactivityTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) InactivityTriggerInput() *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger {
	var returns *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger
	_jsii_.Get(
		j,
		"inactivityTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) InternalValue() *VertexAiOnlineEvaluatorCloudObservabilitySessionScope {
	var returns *VertexAiOnlineEvaluatorCloudObservabilitySessionScope
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiOnlineEvaluator.VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference_Override(v VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiOnlineEvaluator.VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference)SetInternalValue(val *VertexAiOnlineEvaluatorCloudObservabilitySessionScope) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) PutFilter(value interface{}) {
	if err := v.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putFilter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) PutInactivityTrigger(value *VertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger) {
	if err := v.validatePutInactivityTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putInactivityTrigger",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		v,
		"resetFilter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ResetInactivityTrigger() {
	_jsii_.InvokeVoid(
		v,
		"resetInactivityTrigger",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiOnlineEvaluatorCloudObservabilitySessionScopeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

