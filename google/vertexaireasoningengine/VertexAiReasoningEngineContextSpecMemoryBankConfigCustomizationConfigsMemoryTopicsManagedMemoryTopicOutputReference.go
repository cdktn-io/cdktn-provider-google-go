// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexaireasoningengine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference interface {
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
	InternalValue() *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic
	SetInternalValue(val *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic)
	ManagedTopicEnum() *string
	SetManagedTopicEnum(val *string)
	ManagedTopicEnumInput() *string
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
	ResetManagedTopicEnum()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference
type jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) InternalValue() *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic {
	var returns *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ManagedTopicEnum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTopicEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ManagedTopicEnumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTopicEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference_Override(v VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiReasoningEngine.VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetInternalValue(val *VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetManagedTopicEnum(val *string) {
	if err := j.validateSetManagedTopicEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedTopicEnum",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ResetManagedTopicEnum() {
	_jsii_.InvokeVoid(
		v,
		"resetManagedTopicEnum",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

