// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalineageopenlineagejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/datalineageopenlineagejob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataLineageOpenLineageJobOutputOutputReference interface {
	cdktn.ComplexObject
	Catalog() DataLineageOpenLineageJobOutputCatalogOutputReference
	CatalogInput() *DataLineageOpenLineageJobOutputCatalog
	ColumnLineage() DataLineageOpenLineageJobOutputColumnLineageOutputReference
	ColumnLineageInput() *DataLineageOpenLineageJobOutputColumnLineage
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Symlink() DataLineageOpenLineageJobOutputSymlinkList
	SymlinkInput() interface{}
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
	PutCatalog(value *DataLineageOpenLineageJobOutputCatalog)
	PutColumnLineage(value *DataLineageOpenLineageJobOutputColumnLineage)
	PutSymlink(value interface{})
	ResetCatalog()
	ResetColumnLineage()
	ResetSymlink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLineageOpenLineageJobOutputOutputReference
type jsiiProxy_DataLineageOpenLineageJobOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Catalog() DataLineageOpenLineageJobOutputCatalogOutputReference {
	var returns DataLineageOpenLineageJobOutputCatalogOutputReference
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) CatalogInput() *DataLineageOpenLineageJobOutputCatalog {
	var returns *DataLineageOpenLineageJobOutputCatalog
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ColumnLineage() DataLineageOpenLineageJobOutputColumnLineageOutputReference {
	var returns DataLineageOpenLineageJobOutputColumnLineageOutputReference
	_jsii_.Get(
		j,
		"columnLineage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ColumnLineageInput() *DataLineageOpenLineageJobOutputColumnLineage {
	var returns *DataLineageOpenLineageJobOutputColumnLineage
	_jsii_.Get(
		j,
		"columnLineageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Symlink() DataLineageOpenLineageJobOutputSymlinkList {
	var returns DataLineageOpenLineageJobOutputSymlinkList
	_jsii_.Get(
		j,
		"symlink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) SymlinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"symlinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLineageOpenLineageJobOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLineageOpenLineageJobOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataLineageOpenLineageJobOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLineageOpenLineageJobOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLineageOpenLineageJob.DataLineageOpenLineageJobOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLineageOpenLineageJobOutputOutputReference_Override(d DataLineageOpenLineageJobOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLineageOpenLineageJob.DataLineageOpenLineageJobOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) PutCatalog(value *DataLineageOpenLineageJobOutputCatalog) {
	if err := d.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCatalog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) PutColumnLineage(value *DataLineageOpenLineageJobOutputColumnLineage) {
	if err := d.validatePutColumnLineageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putColumnLineage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) PutSymlink(value interface{}) {
	if err := d.validatePutSymlinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSymlink",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ResetColumnLineage() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnLineage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ResetSymlink() {
	_jsii_.InvokeVoid(
		d,
		"resetSymlink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLineageOpenLineageJobOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

