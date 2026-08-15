// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/dataformrepository/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataformRepositoryWorkspaceCompilationOverridesOutputReference interface {
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
	DefaultDatabase() *string
	SetDefaultDatabase(val *string)
	DefaultDatabaseInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataformRepositoryWorkspaceCompilationOverrides
	SetInternalValue(val *DataformRepositoryWorkspaceCompilationOverrides)
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
	ResetDefaultDatabase()
	ResetSchemaSuffix()
	ResetTablePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataformRepositoryWorkspaceCompilationOverridesOutputReference
type jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) DefaultDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) DefaultDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) InternalValue() *DataformRepositoryWorkspaceCompilationOverrides {
	var returns *DataformRepositoryWorkspaceCompilationOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) SchemaSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) SchemaSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataformRepositoryWorkspaceCompilationOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataformRepositoryWorkspaceCompilationOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDataformRepositoryWorkspaceCompilationOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepository.DataformRepositoryWorkspaceCompilationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataformRepositoryWorkspaceCompilationOverridesOutputReference_Override(d DataformRepositoryWorkspaceCompilationOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepository.DataformRepositoryWorkspaceCompilationOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetDefaultDatabase(val *string) {
	if err := j.validateSetDefaultDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDatabase",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetInternalValue(val *DataformRepositoryWorkspaceCompilationOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetSchemaSuffix(val *string) {
	if err := j.validateSetSchemaSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaSuffix",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetDefaultDatabase() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDatabase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetSchemaSuffix() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaSuffix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataformRepositoryWorkspaceCompilationOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

