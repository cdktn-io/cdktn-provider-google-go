// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/biglakeicebergcatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference interface {
	cdktn.ComplexObject
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	AwsRoleArn() *string
	SetAwsRoleArn(val *string)
	AwsRoleArnInput() *string
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
	InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference
type jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) AwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) AwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}


func NewBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference {
	_init_.Initialize()

	if err := validateNewBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference_Override(b BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetAwsRoleArn(val *string) {
	if err := j.validateSetAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRoleArn",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference)SetWarehouse(val *string) {
	if err := j.validateSetWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

