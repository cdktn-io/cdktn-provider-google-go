// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference interface {
	cdktn.ComplexObject
	AllowedMachineSeries() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesAllowedMachineSeriesList
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
	InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferences
	SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferences)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference
type jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) AllowedMachineSeries() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesAllowedMachineSeriesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesAllowedMachineSeriesList
	_jsii_.Get(
		j,
		"allowedMachineSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferences {
	var returns *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference_Override(m MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference)SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

