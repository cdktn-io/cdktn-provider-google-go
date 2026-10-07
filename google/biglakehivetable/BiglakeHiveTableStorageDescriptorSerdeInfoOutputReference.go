// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakehivetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/biglakehivetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference interface {
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
	DeserializerClass() *string
	SetDeserializerClass(val *string)
	DeserializerClassInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BiglakeHiveTableStorageDescriptorSerdeInfo
	SetInternalValue(val *BiglakeHiveTableStorageDescriptorSerdeInfo)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SerdeType() *string
	SetSerdeType(val *string)
	SerdeTypeInput() *string
	SerializationLib() *string
	SetSerializationLib(val *string)
	SerializationLibInput() *string
	SerializerClass() *string
	SetSerializerClass(val *string)
	SerializerClassInput() *string
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
	ResetDescription()
	ResetDeserializerClass()
	ResetParameters()
	ResetSerdeType()
	ResetSerializerClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference
type jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DeserializerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deserializerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) DeserializerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deserializerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InternalValue() *BiglakeHiveTableStorageDescriptorSerdeInfo {
	var returns *BiglakeHiveTableStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerdeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerdeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializationLib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializationLibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializerClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializerClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) SerializerClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializerClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference {
	_init_.Initialize()

	if err := validateNewBiglakeHiveTableStorageDescriptorSerdeInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.biglakeHiveTable.BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBiglakeHiveTableStorageDescriptorSerdeInfoOutputReference_Override(b BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.biglakeHiveTable.BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetDeserializerClass(val *string) {
	if err := j.validateSetDeserializerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deserializerClass",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetInternalValue(val *BiglakeHiveTableStorageDescriptorSerdeInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerdeType(val *string) {
	if err := j.validateSetSerdeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serdeType",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerializationLib(val *string) {
	if err := j.validateSetSerializationLibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializationLib",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetSerializerClass(val *string) {
	if err := j.validateSetSerializerClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializerClass",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetDeserializerClass() {
	_jsii_.InvokeVoid(
		b,
		"resetDeserializerClass",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetSerdeType() {
	_jsii_.InvokeVoid(
		b,
		"resetSerdeType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ResetSerializerClass() {
	_jsii_.InvokeVoid(
		b,
		"resetSerializerClass",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeHiveTableStorageDescriptorSerdeInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

