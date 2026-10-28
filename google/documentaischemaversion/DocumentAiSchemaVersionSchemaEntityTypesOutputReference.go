// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package documentaischemaversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/documentaischemaversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DocumentAiSchemaVersionSchemaEntityTypesOutputReference interface {
	cdktn.ComplexObject
	BaseTypes() *[]*string
	SetBaseTypes(val *[]*string)
	BaseTypesInput() *[]*string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnumValues() DocumentAiSchemaVersionSchemaEntityTypesEnumValuesOutputReference
	EnumValuesInput() *DocumentAiSchemaVersionSchemaEntityTypesEnumValues
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Properties() DocumentAiSchemaVersionSchemaEntityTypesPropertiesList
	PropertiesInput() interface{}
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
	PutEnumValues(value *DocumentAiSchemaVersionSchemaEntityTypesEnumValues)
	PutProperties(value interface{})
	ResetBaseTypes()
	ResetDisplayName()
	ResetEnumValues()
	ResetProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DocumentAiSchemaVersionSchemaEntityTypesOutputReference
type jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) BaseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"baseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) BaseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"baseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) EnumValues() DocumentAiSchemaVersionSchemaEntityTypesEnumValuesOutputReference {
	var returns DocumentAiSchemaVersionSchemaEntityTypesEnumValuesOutputReference
	_jsii_.Get(
		j,
		"enumValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) EnumValuesInput() *DocumentAiSchemaVersionSchemaEntityTypesEnumValues {
	var returns *DocumentAiSchemaVersionSchemaEntityTypesEnumValues
	_jsii_.Get(
		j,
		"enumValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) Properties() DocumentAiSchemaVersionSchemaEntityTypesPropertiesList {
	var returns DocumentAiSchemaVersionSchemaEntityTypesPropertiesList
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) PropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDocumentAiSchemaVersionSchemaEntityTypesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DocumentAiSchemaVersionSchemaEntityTypesOutputReference {
	_init_.Initialize()

	if err := validateNewDocumentAiSchemaVersionSchemaEntityTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.documentAiSchemaVersion.DocumentAiSchemaVersionSchemaEntityTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDocumentAiSchemaVersionSchemaEntityTypesOutputReference_Override(d DocumentAiSchemaVersionSchemaEntityTypesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.documentAiSchemaVersion.DocumentAiSchemaVersionSchemaEntityTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetBaseTypes(val *[]*string) {
	if err := j.validateSetBaseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseTypes",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) PutEnumValues(value *DocumentAiSchemaVersionSchemaEntityTypesEnumValues) {
	if err := d.validatePutEnumValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnumValues",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) PutProperties(value interface{}) {
	if err := d.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ResetBaseTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetBaseTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ResetEnumValues() {
	_jsii_.InvokeVoid(
		d,
		"resetEnumValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiSchemaVersionSchemaEntityTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

