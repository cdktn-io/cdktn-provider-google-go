// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Examples() *[]*string
	SetExamples(val *[]*string)
	ExamplesInput() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InputModes() *[]*string
	SetInputModes(val *[]*string)
	InputModesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OutputModes() *[]*string
	SetOutputModes(val *[]*string)
	OutputModesInput() *[]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
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
	ResetExamples()
	ResetInputModes()
	ResetOutputModes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference
type jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Examples() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"examples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ExamplesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"examplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) InputModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) InputModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) OutputModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) OutputModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference {
	_init_.Initialize()

	if err := validateNewCesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference_Override(c CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetExamples(val *[]*string) {
	if err := j.validateSetExamplesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"examples",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetInputModes(val *[]*string) {
	if err := j.validateSetInputModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputModes",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetOutputModes(val *[]*string) {
	if err := j.validateSetOutputModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputModes",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ResetExamples() {
	_jsii_.InvokeVoid(
		c,
		"resetExamples",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ResetInputModes() {
	_jsii_.InvokeVoid(
		c,
		"resetInputModes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ResetOutputModes() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputModes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigAgentCardSkillsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

