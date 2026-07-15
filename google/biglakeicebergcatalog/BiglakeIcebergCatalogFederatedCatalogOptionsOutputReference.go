// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/biglakeicebergcatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference interface {
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
	GlueCatalogInfo() BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference
	GlueCatalogInfoInput() *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptions
	SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptions)
	RefreshOptions() BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference
	RefreshOptionsInput() *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions
	RefreshStatus() BiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
	ServiceDirectoryName() *string
	SetServiceDirectoryName(val *string)
	ServiceDirectoryNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnityCatalogInfo() BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference
	UnityCatalogInfoInput() *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
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
	PutGlueCatalogInfo(value *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo)
	PutRefreshOptions(value *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions)
	PutUnityCatalogInfo(value *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo)
	ResetGlueCatalogInfo()
	ResetRefreshOptions()
	ResetSecretName()
	ResetServiceDirectoryName()
	ResetUnityCatalogInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference
type jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GlueCatalogInfo() BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference {
	var returns BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"glueCatalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GlueCatalogInfoInput() *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo
	_jsii_.Get(
		j,
		"glueCatalogInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InternalValue() *BiglakeIcebergCatalogFederatedCatalogOptions {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshOptions() BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference {
	var returns BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsOutputReference
	_jsii_.Get(
		j,
		"refreshOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshOptionsInput() *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions
	_jsii_.Get(
		j,
		"refreshOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) RefreshStatus() BiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList {
	var returns BiglakeIcebergCatalogFederatedCatalogOptionsRefreshStatusList
	_jsii_.Get(
		j,
		"refreshStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ServiceDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ServiceDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) UnityCatalogInfo() BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference {
	var returns BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"unityCatalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) UnityCatalogInfoInput() *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo {
	var returns *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo
	_jsii_.Get(
		j,
		"unityCatalogInfoInput",
		&returns,
	)
	return returns
}


func NewBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewBiglakeIcebergCatalogFederatedCatalogOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBiglakeIcebergCatalogFederatedCatalogOptionsOutputReference_Override(b BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.biglakeIcebergCatalog.BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetInternalValue(val *BiglakeIcebergCatalogFederatedCatalogOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetSecretName(val *string) {
	if err := j.validateSetSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetServiceDirectoryName(val *string) {
	if err := j.validateSetServiceDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceDirectoryName",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutGlueCatalogInfo(value *BiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo) {
	if err := b.validatePutGlueCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGlueCatalogInfo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutRefreshOptions(value *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions) {
	if err := b.validatePutRefreshOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRefreshOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) PutUnityCatalogInfo(value *BiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo) {
	if err := b.validatePutUnityCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUnityCatalogInfo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetGlueCatalogInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetGlueCatalogInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetRefreshOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetRefreshOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetSecretName() {
	_jsii_.InvokeVoid(
		b,
		"resetSecretName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetServiceDirectoryName() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceDirectoryName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ResetUnityCatalogInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetUnityCatalogInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BiglakeIcebergCatalogFederatedCatalogOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

