// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/backupdrrestoreworkload/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload google_backup_dr_restore_workload}.
type BackupDrRestoreWorkload interface {
	cdktn.TerraformResource
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	BackupVaultId() *string
	SetBackupVaultId(val *string)
	BackupVaultIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClearOverridesFieldMask() *string
	SetClearOverridesFieldMask(val *string)
	ClearOverridesFieldMaskInput() *string
	ComputeInstanceRestoreProperties() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
	ComputeInstanceRestorePropertiesInput() *BackupDrRestoreWorkloadComputeInstanceRestoreProperties
	ComputeInstanceTargetEnvironment() BackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference
	ComputeInstanceTargetEnvironmentInput() *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment
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
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
	DeleteRestoredInstance() interface{}
	SetDeleteRestoredInstance(val interface{})
	DeleteRestoredInstanceInput() interface{}
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskRestoreProperties() BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
	DiskRestorePropertiesInput() *BackupDrRestoreWorkloadDiskRestoreProperties
	DiskTargetEnvironment() BackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference
	DiskTargetEnvironmentInput() *BackupDrRestoreWorkloadDiskTargetEnvironment
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RegionDiskTargetEnvironment() BackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference
	RegionDiskTargetEnvironmentInput() *BackupDrRestoreWorkloadRegionDiskTargetEnvironment
	RequestId() *string
	SetRequestId(val *string)
	RequestIdInput() *string
	TargetResource() BackupDrRestoreWorkloadTargetResourceList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BackupDrRestoreWorkloadTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutComputeInstanceRestoreProperties(value *BackupDrRestoreWorkloadComputeInstanceRestoreProperties)
	PutComputeInstanceTargetEnvironment(value *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment)
	PutDiskRestoreProperties(value *BackupDrRestoreWorkloadDiskRestoreProperties)
	PutDiskTargetEnvironment(value *BackupDrRestoreWorkloadDiskTargetEnvironment)
	PutRegionDiskTargetEnvironment(value *BackupDrRestoreWorkloadRegionDiskTargetEnvironment)
	PutTimeouts(value *BackupDrRestoreWorkloadTimeouts)
	ResetClearOverridesFieldMask()
	ResetComputeInstanceRestoreProperties()
	ResetComputeInstanceTargetEnvironment()
	ResetDeleteRestoredInstance()
	ResetDeletionPolicy()
	ResetDiskRestoreProperties()
	ResetDiskTargetEnvironment()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegionDiskTargetEnvironment()
	ResetRequestId()
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

// The jsii proxy struct for BackupDrRestoreWorkload
type jsiiProxy_BackupDrRestoreWorkload struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BackupDrRestoreWorkload) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) BackupVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) BackupVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ClearOverridesFieldMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearOverridesFieldMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ClearOverridesFieldMaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clearOverridesFieldMaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ComputeInstanceRestoreProperties() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
	_jsii_.Get(
		j,
		"computeInstanceRestoreProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ComputeInstanceRestorePropertiesInput() *BackupDrRestoreWorkloadComputeInstanceRestoreProperties {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestoreProperties
	_jsii_.Get(
		j,
		"computeInstanceRestorePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ComputeInstanceTargetEnvironment() BackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"computeInstanceTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ComputeInstanceTargetEnvironmentInput() *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment {
	var returns *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment
	_jsii_.Get(
		j,
		"computeInstanceTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DeleteRestoredInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRestoredInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DeleteRestoredInstanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteRestoredInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DiskRestoreProperties() BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference {
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
	_jsii_.Get(
		j,
		"diskRestoreProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DiskRestorePropertiesInput() *BackupDrRestoreWorkloadDiskRestoreProperties {
	var returns *BackupDrRestoreWorkloadDiskRestoreProperties
	_jsii_.Get(
		j,
		"diskRestorePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DiskTargetEnvironment() BackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference {
	var returns BackupDrRestoreWorkloadDiskTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"diskTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) DiskTargetEnvironmentInput() *BackupDrRestoreWorkloadDiskTargetEnvironment {
	var returns *BackupDrRestoreWorkloadDiskTargetEnvironment
	_jsii_.Get(
		j,
		"diskTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) RegionDiskTargetEnvironment() BackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference {
	var returns BackupDrRestoreWorkloadRegionDiskTargetEnvironmentOutputReference
	_jsii_.Get(
		j,
		"regionDiskTargetEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) RegionDiskTargetEnvironmentInput() *BackupDrRestoreWorkloadRegionDiskTargetEnvironment {
	var returns *BackupDrRestoreWorkloadRegionDiskTargetEnvironment
	_jsii_.Get(
		j,
		"regionDiskTargetEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) RequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) RequestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) TargetResource() BackupDrRestoreWorkloadTargetResourceList {
	var returns BackupDrRestoreWorkloadTargetResourceList
	_jsii_.Get(
		j,
		"targetResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) Timeouts() BackupDrRestoreWorkloadTimeoutsOutputReference {
	var returns BackupDrRestoreWorkloadTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkload) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload google_backup_dr_restore_workload} Resource.
func NewBackupDrRestoreWorkload(scope constructs.Construct, id *string, config *BackupDrRestoreWorkloadConfig) BackupDrRestoreWorkload {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkload{}

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/backup_dr_restore_workload google_backup_dr_restore_workload} Resource.
func NewBackupDrRestoreWorkload_Override(b BackupDrRestoreWorkload, scope constructs.Construct, id *string, config *BackupDrRestoreWorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetBackupVaultId(val *string) {
	if err := j.validateSetBackupVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVaultId",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetClearOverridesFieldMask(val *string) {
	if err := j.validateSetClearOverridesFieldMaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clearOverridesFieldMask",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetDataSourceId(val *string) {
	if err := j.validateSetDataSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetDeleteRestoredInstance(val interface{}) {
	if err := j.validateSetDeleteRestoredInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteRestoredInstance",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkload)SetRequestId(val *string) {
	if err := j.validateSetRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestId",
		val,
	)
}

// Generates CDKTN code for importing a BackupDrRestoreWorkload resource upon running "cdktn plan <stack-name>".
func BackupDrRestoreWorkload_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBackupDrRestoreWorkload_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
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
func BackupDrRestoreWorkload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupDrRestoreWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupDrRestoreWorkload_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupDrRestoreWorkload_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupDrRestoreWorkload_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupDrRestoreWorkload_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupDrRestoreWorkload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutComputeInstanceRestoreProperties(value *BackupDrRestoreWorkloadComputeInstanceRestoreProperties) {
	if err := b.validatePutComputeInstanceRestorePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putComputeInstanceRestoreProperties",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutComputeInstanceTargetEnvironment(value *BackupDrRestoreWorkloadComputeInstanceTargetEnvironment) {
	if err := b.validatePutComputeInstanceTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putComputeInstanceTargetEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutDiskRestoreProperties(value *BackupDrRestoreWorkloadDiskRestoreProperties) {
	if err := b.validatePutDiskRestorePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDiskRestoreProperties",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutDiskTargetEnvironment(value *BackupDrRestoreWorkloadDiskTargetEnvironment) {
	if err := b.validatePutDiskTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDiskTargetEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutRegionDiskTargetEnvironment(value *BackupDrRestoreWorkloadRegionDiskTargetEnvironment) {
	if err := b.validatePutRegionDiskTargetEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRegionDiskTargetEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) PutTimeouts(value *BackupDrRestoreWorkloadTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetClearOverridesFieldMask() {
	_jsii_.InvokeVoid(
		b,
		"resetClearOverridesFieldMask",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetComputeInstanceRestoreProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeInstanceRestoreProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetComputeInstanceTargetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeInstanceTargetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetDeleteRestoredInstance() {
	_jsii_.InvokeVoid(
		b,
		"resetDeleteRestoredInstance",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetDiskRestoreProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetDiskRestoreProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetDiskTargetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetDiskTargetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetRegionDiskTargetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetRegionDiskTargetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetRequestId() {
	_jsii_.InvokeVoid(
		b,
		"resetRequestId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkload) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}

