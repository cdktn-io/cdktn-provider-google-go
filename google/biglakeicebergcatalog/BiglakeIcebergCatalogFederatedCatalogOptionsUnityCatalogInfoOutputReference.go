// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/biglakeicebergcatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference interface {
	cdktn.ComplexObject
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
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
	InstanceName() *string
	SetInstanceName(val *string)
	InstanceNameInput() *string
	InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
	SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo)
	ServicePrincipalApplicationId() *string
	SetServicePrincipalApplicationId(val *string)
	ServicePrincipalApplicationIdInput() *string
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
	ResetServicePrincipalApplicationId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference
type jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) InstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) InstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ServicePrincipalApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ServicePrincipalApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference {
	_init_.Initialize()

	if err := validateNewBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference_Override(b BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetInstanceName(val *string) {
	if err := j.validateSetInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceName",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetServicePrincipalApplicationId(val *string) {
	if err := j.validateSetServicePrincipalApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicePrincipalApplicationId",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ResetServicePrincipalApplicationId() {
	_jsii_.InvokeVoid(
		b,
		"resetServicePrincipalApplicationId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

