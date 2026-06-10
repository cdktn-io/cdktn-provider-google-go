// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterimportjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterimportjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference interface {
	cdktn.ComplexObject
	ArchiveError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList
	AssetTitle() *string
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
	CsvError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList
	Errors() MigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterImportJobValidationReportFileValidationsRowErrors
	SetInternalValue(val *MigrationCenterImportJobValidationReportFileValidationsRowErrors)
	RowNumber() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmName() *string
	VmUuid() *string
	XlsxError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList
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

// The jsii proxy struct for MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference
type jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ArchiveError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList {
	var returns MigrationCenterImportJobValidationReportFileValidationsRowErrorsArchiveErrorList
	_jsii_.Get(
		j,
		"archiveError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) AssetTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) CsvError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList {
	var returns MigrationCenterImportJobValidationReportFileValidationsRowErrorsCsvErrorList
	_jsii_.Get(
		j,
		"csvError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Errors() MigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList {
	var returns MigrationCenterImportJobValidationReportFileValidationsRowErrorsErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InternalValue() *MigrationCenterImportJobValidationReportFileValidationsRowErrors {
	var returns *MigrationCenterImportJobValidationReportFileValidationsRowErrors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) RowNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) VmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) VmUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) XlsxError() MigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList {
	var returns MigrationCenterImportJobValidationReportFileValidationsRowErrorsXlsxErrorList
	_jsii_.Get(
		j,
		"xlsxError",
		&returns,
	)
	return returns
}


func NewMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterImportJob.MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference_Override(m MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterImportJob.MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetInternalValue(val *MigrationCenterImportJobValidationReportFileValidationsRowErrors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterImportJobValidationReportFileValidationsRowErrorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

