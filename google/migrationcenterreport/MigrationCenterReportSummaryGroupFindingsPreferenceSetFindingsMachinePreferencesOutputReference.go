// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference interface {
	cdktn.ComplexObject
	CommitmentPlan() *string
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
	ComputeEnginePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferences
	SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferences)
	RegionPreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesRegionPreferencesList
	SizingOptimizationStrategy() *string
	SoleTenancyPreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesSoleTenancyPreferencesList
	TargetProduct() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmwareEnginePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList
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

// The jsii proxy struct for MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference
type jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) CommitmentPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) ComputeEnginePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesComputeEnginePreferencesList
	_jsii_.Get(
		j,
		"computeEnginePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferences {
	var returns *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) RegionPreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesRegionPreferencesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesRegionPreferencesList
	_jsii_.Get(
		j,
		"regionPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) SizingOptimizationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingOptimizationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) SoleTenancyPreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesSoleTenancyPreferencesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesSoleTenancyPreferencesList
	_jsii_.Get(
		j,
		"soleTenancyPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) TargetProduct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) VmwareEnginePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesVmwareEnginePreferencesList
	_jsii_.Get(
		j,
		"vmwareEnginePreferences",
		&returns,
	)
	return returns
}


func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference_Override(m MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference)SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

