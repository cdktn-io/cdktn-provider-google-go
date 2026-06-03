// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference interface {
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
	ComputeEngineFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindings
	SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindings)
	MachinePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList
	MonthlyCostCompute() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList
	MonthlyCostNetworkEgress() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList
	MonthlyCostOsLicense() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList
	MonthlyCostOther() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList
	MonthlyCostStorage() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList
	MonthlyCostTotal() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList
	SoleTenantFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmwareEngineFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList
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

// The jsii proxy struct for MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference
type jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComputeEngineFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsComputeEngineFindingList
	_jsii_.Get(
		j,
		"computeEngineFinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InternalValue() *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindings {
	var returns *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MachinePreferences() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMachinePreferencesList
	_jsii_.Get(
		j,
		"machinePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostCompute() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostComputeList
	_jsii_.Get(
		j,
		"monthlyCostCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostNetworkEgress() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostNetworkEgressList
	_jsii_.Get(
		j,
		"monthlyCostNetworkEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostOsLicense() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOsLicenseList
	_jsii_.Get(
		j,
		"monthlyCostOsLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostOther() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostOtherList
	_jsii_.Get(
		j,
		"monthlyCostOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostStorage() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostStorageList
	_jsii_.Get(
		j,
		"monthlyCostStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) MonthlyCostTotal() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsMonthlyCostTotalList
	_jsii_.Get(
		j,
		"monthlyCostTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) SoleTenantFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsSoleTenantFindingList
	_jsii_.Get(
		j,
		"soleTenantFinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) VmwareEngineFinding() MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList {
	var returns MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsVmwareEngineFindingList
	_jsii_.Get(
		j,
		"vmwareEngineFinding",
		&returns,
	)
	return returns
}


func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference_Override(m MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetInternalValue(val *MigrationCenterReportSummaryGroupFindingsPreferenceSetFindings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsPreferenceSetFindingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

