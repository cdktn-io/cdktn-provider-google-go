// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterreport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterreport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList
type jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList {
	_init_.Initialize()

	if err := validateNewMigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList_Override(m MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterReport.MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := m.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		m,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) Get(index *float64) MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterReportSummaryGroupFindingsAssetAggregateStatsOperatingSystemList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

