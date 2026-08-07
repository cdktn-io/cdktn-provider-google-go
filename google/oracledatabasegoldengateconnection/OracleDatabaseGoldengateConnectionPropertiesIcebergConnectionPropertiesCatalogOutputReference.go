// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference interface {
	cdktn.ComplexObject
	CatalogType() *string
	SetCatalogType(val *string)
	CatalogTypeInput() *string
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
	GlueIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalogOutputReference
	GlueIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog
	InternalValue() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog
	SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog)
	NessieIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalogOutputReference
	NessieIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog
	PolarisIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalogOutputReference
	PolarisIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog
	RestIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalogOutputReference
	RestIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog
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
	PutGlueIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog)
	PutNessieIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog)
	PutPolarisIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog)
	PutRestIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog)
	ResetGlueIcebergCatalog()
	ResetNessieIcebergCatalog()
	ResetPolarisIcebergCatalog()
	ResetRestIcebergCatalog()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) CatalogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) CatalogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GlueIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalogOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalogOutputReference
	_jsii_.Get(
		j,
		"glueIcebergCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GlueIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog
	_jsii_.Get(
		j,
		"glueIcebergCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) NessieIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalogOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalogOutputReference
	_jsii_.Get(
		j,
		"nessieIcebergCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) NessieIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog
	_jsii_.Get(
		j,
		"nessieIcebergCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PolarisIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalogOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalogOutputReference
	_jsii_.Get(
		j,
		"polarisIcebergCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PolarisIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog
	_jsii_.Get(
		j,
		"polarisIcebergCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) RestIcebergCatalog() OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalogOutputReference {
	var returns OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalogOutputReference
	_jsii_.Get(
		j,
		"restIcebergCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) RestIcebergCatalogInput() *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog {
	var returns *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog
	_jsii_.Get(
		j,
		"restIcebergCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetCatalogType(val *string) {
	if err := j.validateSetCatalogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PutGlueIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog) {
	if err := o.validatePutGlueIcebergCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGlueIcebergCatalog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PutNessieIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog) {
	if err := o.validatePutNessieIcebergCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNessieIcebergCatalog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PutPolarisIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog) {
	if err := o.validatePutPolarisIcebergCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPolarisIcebergCatalog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) PutRestIcebergCatalog(value *OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog) {
	if err := o.validatePutRestIcebergCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRestIcebergCatalog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ResetGlueIcebergCatalog() {
	_jsii_.InvokeVoid(
		o,
		"resetGlueIcebergCatalog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ResetNessieIcebergCatalog() {
	_jsii_.InvokeVoid(
		o,
		"resetNessieIcebergCatalog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ResetPolarisIcebergCatalog() {
	_jsii_.InvokeVoid(
		o,
		"resetPolarisIcebergCatalog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ResetRestIcebergCatalog() {
	_jsii_.InvokeVoid(
		o,
		"resetRestIcebergCatalog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

