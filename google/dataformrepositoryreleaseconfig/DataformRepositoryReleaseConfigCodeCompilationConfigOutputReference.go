// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryreleaseconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dataformrepositoryreleaseconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference interface {
	cdktn.ComplexObject
	AssertionSchema() *string
	SetAssertionSchema(val *string)
	AssertionSchemaInput() *string
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
	DatabaseSuffix() *string
	SetDatabaseSuffix(val *string)
	DatabaseSuffixInput() *string
	DefaultDatabase() *string
	SetDefaultDatabase(val *string)
	DefaultDatabaseInput() *string
	DefaultLocation() *string
	SetDefaultLocation(val *string)
	DefaultLocationInput() *string
	DefaultSchema() *string
	SetDefaultSchema(val *string)
	DefaultSchemaInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataformRepositoryReleaseConfigCodeCompilationConfig
	SetInternalValue(val *DataformRepositoryReleaseConfigCodeCompilationConfig)
	SchemaSuffix() *string
	SetSchemaSuffix(val *string)
	SchemaSuffixInput() *string
	TablePrefix() *string
	SetTablePrefix(val *string)
	TablePrefixInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vars() *map[string]*string
	SetVars(val *map[string]*string)
	VarsInput() *map[string]*string
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
	ResetAssertionSchema()
	ResetDatabaseSuffix()
	ResetDefaultDatabase()
	ResetDefaultLocation()
	ResetDefaultSchema()
	ResetSchemaSuffix()
	ResetTablePrefix()
	ResetVars()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference
type jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) AssertionSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) AssertionSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DatabaseSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DatabaseSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) DefaultSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InternalValue() *DataformRepositoryReleaseConfigCodeCompilationConfig {
	var returns *DataformRepositoryReleaseConfigCodeCompilationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) SchemaSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) SchemaSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Vars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) VarsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"varsInput",
		&returns,
	)
	return returns
}


func NewDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataformRepositoryReleaseConfigCodeCompilationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepositoryReleaseConfig.DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataformRepositoryReleaseConfigCodeCompilationConfigOutputReference_Override(d DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepositoryReleaseConfig.DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetAssertionSchema(val *string) {
	if err := j.validateSetAssertionSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertionSchema",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDatabaseSuffix(val *string) {
	if err := j.validateSetDatabaseSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseSuffix",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultDatabase(val *string) {
	if err := j.validateSetDefaultDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDatabase",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultLocation(val *string) {
	if err := j.validateSetDefaultLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLocation",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetDefaultSchema(val *string) {
	if err := j.validateSetDefaultSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSchema",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetInternalValue(val *DataformRepositoryReleaseConfigCodeCompilationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetSchemaSuffix(val *string) {
	if err := j.validateSetSchemaSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaSuffix",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference)SetVars(val *map[string]*string) {
	if err := j.validateSetVarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vars",
		val,
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetAssertionSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetAssertionSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDatabaseSuffix() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseSuffix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetDefaultSchema() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultSchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetSchemaSuffix() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaSuffix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ResetVars() {
	_jsii_.InvokeVoid(
		d,
		"resetVars",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataformRepositoryReleaseConfigCodeCompilationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

