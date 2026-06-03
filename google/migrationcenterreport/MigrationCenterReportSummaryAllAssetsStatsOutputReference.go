// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterReportSummaryAllAssetsStatsOutputReference interface {
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
	CoreCountHistogram() MigrationCenterReportSummaryAllAssetsStatsCoreCountHistogramList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterReportSummaryAllAssetsStats
	SetInternalValue(val *MigrationCenterReportSummaryAllAssetsStats)
	MemoryBytesHistogram() MigrationCenterReportSummaryAllAssetsStatsMemoryBytesHistogramList
	MemoryUtilizationChart() MigrationCenterReportSummaryAllAssetsStatsMemoryUtilizationChartList
	OperatingSystem() MigrationCenterReportSummaryAllAssetsStatsOperatingSystemList
	SoftwareInstances() MigrationCenterReportSummaryAllAssetsStatsSoftwareInstancesList
	StorageBytesHistogram() MigrationCenterReportSummaryAllAssetsStatsStorageBytesHistogramList
	StorageUtilizationChart() MigrationCenterReportSummaryAllAssetsStatsStorageUtilizationChartList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalAssets() *string
	TotalCores() *string
	TotalMemoryBytes() *string
	TotalStorageBytes() *string
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

// The jsii proxy struct for MigrationCenterReportSummaryAllAssetsStatsOutputReference
type jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) CoreCountHistogram() MigrationCenterReportSummaryAllAssetsStatsCoreCountHistogramList {
	var returns MigrationCenterReportSummaryAllAssetsStatsCoreCountHistogramList
	_jsii_.Get(
		j,
		"coreCountHistogram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) InternalValue() *MigrationCenterReportSummaryAllAssetsStats {
	var returns *MigrationCenterReportSummaryAllAssetsStats
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) MemoryBytesHistogram() MigrationCenterReportSummaryAllAssetsStatsMemoryBytesHistogramList {
	var returns MigrationCenterReportSummaryAllAssetsStatsMemoryBytesHistogramList
	_jsii_.Get(
		j,
		"memoryBytesHistogram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) MemoryUtilizationChart() MigrationCenterReportSummaryAllAssetsStatsMemoryUtilizationChartList {
	var returns MigrationCenterReportSummaryAllAssetsStatsMemoryUtilizationChartList
	_jsii_.Get(
		j,
		"memoryUtilizationChart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) OperatingSystem() MigrationCenterReportSummaryAllAssetsStatsOperatingSystemList {
	var returns MigrationCenterReportSummaryAllAssetsStatsOperatingSystemList
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) SoftwareInstances() MigrationCenterReportSummaryAllAssetsStatsSoftwareInstancesList {
	var returns MigrationCenterReportSummaryAllAssetsStatsSoftwareInstancesList
	_jsii_.Get(
		j,
		"softwareInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) StorageBytesHistogram() MigrationCenterReportSummaryAllAssetsStatsStorageBytesHistogramList {
	var returns MigrationCenterReportSummaryAllAssetsStatsStorageBytesHistogramList
	_jsii_.Get(
		j,
		"storageBytesHistogram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) StorageUtilizationChart() MigrationCenterReportSummaryAllAssetsStatsStorageUtilizationChartList {
	var returns MigrationCenterReportSummaryAllAssetsStatsStorageUtilizationChartList
	_jsii_.Get(
		j,
		"storageUtilizationChart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TotalAssets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TotalCores() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalCores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TotalMemoryBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalMemoryBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) TotalStorageBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalStorageBytes",
		&returns,
	)
	return returns
}


func NewMigrationCenterReportSummaryAllAssetsStatsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MigrationCenterReportSummaryAllAssetsStatsOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterReportSummaryAllAssetsStatsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryAllAssetsStatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMigrationCenterReportSummaryAllAssetsStatsOutputReference_Override(m MigrationCenterReportSummaryAllAssetsStatsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryAllAssetsStatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference)SetInternalValue(val *MigrationCenterReportSummaryAllAssetsStats) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryAllAssetsStatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

