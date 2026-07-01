// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package migrationcenterassetsexportjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/migrationcenterassetsexportjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/migration_center_assets_export_job google_migration_center_assets_export_job}.
type MigrationCenterAssetsExportJob interface {
	cdktn.TerraformResource
	AssetsExportJobId() *string
	SetAssetsExportJobId(val *string)
	AssetsExportJobIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Condition() MigrationCenterAssetsExportJobConditionOutputReference
	ConditionInput() *MigrationCenterAssetsExportJobCondition
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktn.StringMap
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Inventory() MigrationCenterAssetsExportJobInventoryList
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	NetworkDependencies() MigrationCenterAssetsExportJobNetworkDependenciesList
	// The tree node.
	Node() constructs.Node
	PerformanceData() MigrationCenterAssetsExportJobPerformanceDataOutputReference
	PerformanceDataInput() *MigrationCenterAssetsExportJobPerformanceData
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecentExecutions() MigrationCenterAssetsExportJobRecentExecutionsList
	ShowHidden() interface{}
	SetShowHidden(val interface{})
	ShowHiddenInput() interface{}
	SignedUriDestination() MigrationCenterAssetsExportJobSignedUriDestinationOutputReference
	SignedUriDestinationInput() *MigrationCenterAssetsExportJobSignedUriDestination
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MigrationCenterAssetsExportJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCondition(value *MigrationCenterAssetsExportJobCondition)
	PutPerformanceData(value *MigrationCenterAssetsExportJobPerformanceData)
	PutSignedUriDestination(value *MigrationCenterAssetsExportJobSignedUriDestination)
	PutTimeouts(value *MigrationCenterAssetsExportJobTimeouts)
	ResetCondition()
	ResetDeletionPolicy()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceData()
	ResetProject()
	ResetShowHidden()
	ResetSignedUriDestination()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for MigrationCenterAssetsExportJob
type jsiiProxy_MigrationCenterAssetsExportJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) AssetsExportJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsExportJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) AssetsExportJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetsExportJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Condition() MigrationCenterAssetsExportJobConditionOutputReference {
	var returns MigrationCenterAssetsExportJobConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ConditionInput() *MigrationCenterAssetsExportJobCondition {
	var returns *MigrationCenterAssetsExportJobCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Inventory() MigrationCenterAssetsExportJobInventoryList {
	var returns MigrationCenterAssetsExportJobInventoryList
	_jsii_.Get(
		j,
		"inventory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) NetworkDependencies() MigrationCenterAssetsExportJobNetworkDependenciesList {
	var returns MigrationCenterAssetsExportJobNetworkDependenciesList
	_jsii_.Get(
		j,
		"networkDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) PerformanceData() MigrationCenterAssetsExportJobPerformanceDataOutputReference {
	var returns MigrationCenterAssetsExportJobPerformanceDataOutputReference
	_jsii_.Get(
		j,
		"performanceData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) PerformanceDataInput() *MigrationCenterAssetsExportJobPerformanceData {
	var returns *MigrationCenterAssetsExportJobPerformanceData
	_jsii_.Get(
		j,
		"performanceDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) RecentExecutions() MigrationCenterAssetsExportJobRecentExecutionsList {
	var returns MigrationCenterAssetsExportJobRecentExecutionsList
	_jsii_.Get(
		j,
		"recentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ShowHidden() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHidden",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) ShowHiddenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHiddenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) SignedUriDestination() MigrationCenterAssetsExportJobSignedUriDestinationOutputReference {
	var returns MigrationCenterAssetsExportJobSignedUriDestinationOutputReference
	_jsii_.Get(
		j,
		"signedUriDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) SignedUriDestinationInput() *MigrationCenterAssetsExportJobSignedUriDestination {
	var returns *MigrationCenterAssetsExportJobSignedUriDestination
	_jsii_.Get(
		j,
		"signedUriDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) Timeouts() MigrationCenterAssetsExportJobTimeoutsOutputReference {
	var returns MigrationCenterAssetsExportJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/migration_center_assets_export_job google_migration_center_assets_export_job} Resource.
func NewMigrationCenterAssetsExportJob(scope constructs.Construct, id *string, config *MigrationCenterAssetsExportJobConfig) MigrationCenterAssetsExportJob {
	_init_.Initialize()

	if err := validateNewMigrationCenterAssetsExportJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterAssetsExportJob{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/migration_center_assets_export_job google_migration_center_assets_export_job} Resource.
func NewMigrationCenterAssetsExportJob_Override(m MigrationCenterAssetsExportJob, scope constructs.Construct, id *string, config *MigrationCenterAssetsExportJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetAssetsExportJobId(val *string) {
	if err := j.validateSetAssetsExportJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetsExportJobId",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterAssetsExportJob)SetShowHidden(val interface{}) {
	if err := j.validateSetShowHiddenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showHidden",
		val,
	)
}

// Generates CDKTN code for importing a MigrationCenterAssetsExportJob resource upon running "cdktn plan <stack-name>".
func MigrationCenterAssetsExportJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMigrationCenterAssetsExportJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MigrationCenterAssetsExportJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMigrationCenterAssetsExportJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MigrationCenterAssetsExportJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMigrationCenterAssetsExportJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MigrationCenterAssetsExportJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMigrationCenterAssetsExportJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MigrationCenterAssetsExportJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.migrationCenterAssetsExportJob.MigrationCenterAssetsExportJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MigrationCenterAssetsExportJob) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) PutCondition(value *MigrationCenterAssetsExportJobCondition) {
	if err := m.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCondition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) PutPerformanceData(value *MigrationCenterAssetsExportJobPerformanceData) {
	if err := m.validatePutPerformanceDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPerformanceData",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) PutSignedUriDestination(value *MigrationCenterAssetsExportJobSignedUriDestination) {
	if err := m.validatePutSignedUriDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSignedUriDestination",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) PutTimeouts(value *MigrationCenterAssetsExportJobTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetCondition() {
	_jsii_.InvokeVoid(
		m,
		"resetCondition",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetPerformanceData() {
	_jsii_.InvokeVoid(
		m,
		"resetPerformanceData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetProject() {
	_jsii_.InvokeVoid(
		m,
		"resetProject",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetShowHidden() {
	_jsii_.InvokeVoid(
		m,
		"resetShowHidden",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetSignedUriDestination() {
	_jsii_.InvokeVoid(
		m,
		"resetSignedUriDestination",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterAssetsExportJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

