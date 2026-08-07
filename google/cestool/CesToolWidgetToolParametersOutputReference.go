// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/cestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesToolWidgetToolParametersOutputReference interface {
	cdktn.ComplexObject
	AdditionalProperties() *string
	SetAdditionalProperties(val *string)
	AdditionalPropertiesInput() *string
	AnyOf() *string
	SetAnyOf(val *string)
	AnyOfInput() *string
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
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	Defs() *string
	SetDefs(val *string)
	DefsInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enum() *[]*string
	SetEnum(val *[]*string)
	EnumInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolWidgetToolParameters
	SetInternalValue(val *CesToolWidgetToolParameters)
	Items() *string
	SetItems(val *string)
	ItemsInput() *string
	Maximum() *float64
	SetMaximum(val *float64)
	MaximumInput() *float64
	MaxItems() *float64
	SetMaxItems(val *float64)
	MaxItemsInput() *float64
	Minimum() *float64
	SetMinimum(val *float64)
	MinimumInput() *float64
	MinItems() *float64
	SetMinItems(val *float64)
	MinItemsInput() *float64
	Nullable() interface{}
	SetNullable(val interface{})
	NullableInput() interface{}
	PrefixItems() *string
	SetPrefixItems(val *string)
	PrefixItemsInput() *string
	Properties() *string
	SetProperties(val *string)
	PropertiesInput() *string
	Ref() *string
	SetRef(val *string)
	RefInput() *string
	Required() *[]*string
	SetRequired(val *[]*string)
	RequiredInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UniqueItems() interface{}
	SetUniqueItems(val interface{})
	UniqueItemsInput() interface{}
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
	ResetAdditionalProperties()
	ResetAnyOf()
	ResetDefault()
	ResetDefs()
	ResetDescription()
	ResetEnum()
	ResetItems()
	ResetMaximum()
	ResetMaxItems()
	ResetMinimum()
	ResetMinItems()
	ResetNullable()
	ResetPrefixItems()
	ResetProperties()
	ResetRef()
	ResetRequired()
	ResetTitle()
	ResetUniqueItems()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolWidgetToolParametersOutputReference
type jsiiProxy_CesToolWidgetToolParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) AdditionalProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) AdditionalPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) AnyOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) AnyOfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Defs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) DefsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) InternalValue() *CesToolWidgetToolParameters {
	var returns *CesToolWidgetToolParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Items() *string {
	var returns *string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) ItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Maximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MaxItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Minimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MinItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) MinItemsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Nullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) NullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) PrefixItems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) PrefixItemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixItemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) PropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) RefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Required() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) RequiredInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) UniqueItems() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference) UniqueItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueItemsInput",
		&returns,
	)
	return returns
}


func NewCesToolWidgetToolParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesToolWidgetToolParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolWidgetToolParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolWidgetToolParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolWidgetToolParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolWidgetToolParametersOutputReference_Override(c CesToolWidgetToolParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolWidgetToolParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetAdditionalProperties(val *string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetAnyOf(val *string) {
	if err := j.validateSetAnyOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyOf",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetDefault(val *string) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetDefs(val *string) {
	if err := j.validateSetDefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defs",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetInternalValue(val *CesToolWidgetToolParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetItems(val *string) {
	if err := j.validateSetItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetMaximum(val *float64) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetMaxItems(val *float64) {
	if err := j.validateSetMaxItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxItems",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetMinimum(val *float64) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetMinItems(val *float64) {
	if err := j.validateSetMinItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minItems",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetNullable(val interface{}) {
	if err := j.validateSetNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullable",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetPrefixItems(val *string) {
	if err := j.validateSetPrefixItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixItems",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetProperties(val *string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetRef(val *string) {
	if err := j.validateSetRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ref",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetRequired(val *[]*string) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolParametersOutputReference)SetUniqueItems(val interface{}) {
	if err := j.validateSetUniqueItemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uniqueItems",
		val,
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetAnyOf() {
	_jsii_.InvokeVoid(
		c,
		"resetAnyOf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		c,
		"resetDefault",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetDefs() {
	_jsii_.InvokeVoid(
		c,
		"resetDefs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		c,
		"resetEnum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetMaxItems() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetMinItems() {
	_jsii_.InvokeVoid(
		c,
		"resetMinItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetNullable() {
	_jsii_.InvokeVoid(
		c,
		"resetNullable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetPrefixItems() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefixItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetRef() {
	_jsii_.InvokeVoid(
		c,
		"resetRef",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		c,
		"resetTitle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ResetUniqueItems() {
	_jsii_.InvokeVoid(
		c,
		"resetUniqueItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesToolWidgetToolParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

