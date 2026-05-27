// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/databasemigrationservicemigrationjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference interface {
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
	InternalValue() *DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig
	SetInternalValue(val *DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig)
	IsNativeLogical() interface{}
	SetIsNativeLogical(val interface{})
	IsNativeLogicalInput() interface{}
	MaxAdditionalSubscriptions() *float64
	SetMaxAdditionalSubscriptions(val *float64)
	MaxAdditionalSubscriptionsInput() *float64
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
	ResetMaxAdditionalSubscriptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference
type jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) InternalValue() *DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig {
	var returns *DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) IsNativeLogical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNativeLogical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) IsNativeLogicalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNativeLogicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) MaxAdditionalSubscriptions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAdditionalSubscriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) MaxAdditionalSubscriptionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAdditionalSubscriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference_Override(d DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.databaseMigrationServiceMigrationJob.DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetInternalValue(val *DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetIsNativeLogical(val interface{}) {
	if err := j.validateSetIsNativeLogicalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNativeLogical",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetMaxAdditionalSubscriptions(val *float64) {
	if err := j.validateSetMaxAdditionalSubscriptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAdditionalSubscriptions",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) ResetMaxAdditionalSubscriptions() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxAdditionalSubscriptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceMigrationJobPostgresHomogeneousConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

