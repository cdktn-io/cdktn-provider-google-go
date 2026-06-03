// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateDeploymentPropertiesOutputReference interface {
	cdktn.ComplexObject
	BackupSchedule() OracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference
	BackupScheduleInput() *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule
	Category() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentBackupId() *string
	DeploymentDiagnosticData() OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference
	DeploymentDiagnosticDataInput() *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData
	DeploymentRole() *string
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DeploymentUrl() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentType() *string
	SetEnvironmentType(val *string)
	EnvironmentTypeInput() *string
	Fqdn() *string
	// Experimental.
	Fqn() *string
	Healthy() cdktn.IResolvable
	IngressIps() OracleDatabaseGoldengateDeploymentPropertiesIngressIpsList
	InternalValue() *OracleDatabaseGoldengateDeploymentProperties
	SetInternalValue(val *OracleDatabaseGoldengateDeploymentProperties)
	IsAutoScalingEnabled() interface{}
	SetIsAutoScalingEnabled(val interface{})
	IsAutoScalingEnabledInput() interface{}
	IsLatestVersion() cdktn.IResolvable
	IsPublic() cdktn.IResolvable
	IsStorageUtilizationLimitExceeded() cdktn.IResolvable
	LastBackupScheduleTime() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	LifecycleDetails() *string
	LifecycleState() *string
	LifecycleSubState() *string
	LoadBalancerId() *string
	LoadBalancerSubnetId() *string
	Locks() OracleDatabaseGoldengateDeploymentPropertiesLocksList
	MaintenanceConfig() OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
	MaintenanceConfigInput() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	MaintenanceWindow() OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference
	MaintenanceWindowInput() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow
	NextBackupScheduleTime() *string
	NextMaintenanceActionType() *string
	NextMaintenanceDescription() *string
	NextMaintenanceTime() *string
	NsgIds() *[]*string
	Ocid() *string
	OggData() OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
	OggDataInput() *OracleDatabaseGoldengateDeploymentPropertiesOggData
	OggVersionSupportEndTime() *string
	Placements() OracleDatabaseGoldengateDeploymentPropertiesPlacementsList
	PrivateIpAddress() *string
	PublicIpAddress() *string
	RoleChangeTime() *string
	StorageUtilizationBytes() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	UpgradeRequiredTime() *string
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
	PutBackupSchedule(value *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule)
	PutDeploymentDiagnosticData(value *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData)
	PutMaintenanceConfig(value *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig)
	PutMaintenanceWindow(value *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow)
	PutOggData(value *OracleDatabaseGoldengateDeploymentPropertiesOggData)
	ResetBackupSchedule()
	ResetCpuCoreCount()
	ResetDeploymentDiagnosticData()
	ResetDescription()
	ResetEnvironmentType()
	ResetIsAutoScalingEnabled()
	ResetLicenseModel()
	ResetMaintenanceConfig()
	ResetMaintenanceWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateDeploymentPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) BackupSchedule() OracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesBackupScheduleOutputReference
	_jsii_.Get(
		j,
		"backupSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) BackupScheduleInput() *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule
	_jsii_.Get(
		j,
		"backupScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentBackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentBackupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentDiagnosticData() OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticDataOutputReference
	_jsii_.Get(
		j,
		"deploymentDiagnosticData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentDiagnosticDataInput() *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData
	_jsii_.Get(
		j,
		"deploymentDiagnosticDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DeploymentUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) EnvironmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) EnvironmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Healthy() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"healthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IngressIps() OracleDatabaseGoldengateDeploymentPropertiesIngressIpsList {
	var returns OracleDatabaseGoldengateDeploymentPropertiesIngressIpsList
	_jsii_.Get(
		j,
		"ingressIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateDeploymentProperties {
	var returns *OracleDatabaseGoldengateDeploymentProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsLatestVersion() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isLatestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsPublic() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) IsStorageUtilizationLimitExceeded() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isStorageUtilizationLimitExceeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LastBackupScheduleTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastBackupScheduleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LifecycleSubState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleSubState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LoadBalancerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) LoadBalancerSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Locks() OracleDatabaseGoldengateDeploymentPropertiesLocksList {
	var returns OracleDatabaseGoldengateDeploymentPropertiesLocksList
	_jsii_.Get(
		j,
		"locks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceConfig() OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
	_jsii_.Get(
		j,
		"maintenanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceConfigInput() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	_jsii_.Get(
		j,
		"maintenanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceWindow() OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) MaintenanceWindowInput() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextBackupScheduleTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextBackupScheduleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) NextMaintenanceTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) NsgIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nsgIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggData() OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
	_jsii_.Get(
		j,
		"oggData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggDataInput() *OracleDatabaseGoldengateDeploymentPropertiesOggData {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesOggData
	_jsii_.Get(
		j,
		"oggDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) OggVersionSupportEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersionSupportEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Placements() OracleDatabaseGoldengateDeploymentPropertiesPlacementsList {
	var returns OracleDatabaseGoldengateDeploymentPropertiesPlacementsList
	_jsii_.Get(
		j,
		"placements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) RoleChangeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleChangeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) StorageUtilizationBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageUtilizationBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) UpgradeRequiredTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRequiredTime",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateDeploymentPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateDeploymentPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateDeploymentPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateDeploymentPropertiesOutputReference_Override(o OracleDatabaseGoldengateDeploymentPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetEnvironmentType(val *string) {
	if err := j.validateSetEnvironmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateDeploymentProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetIsAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutBackupSchedule(value *OracleDatabaseGoldengateDeploymentPropertiesBackupSchedule) {
	if err := o.validatePutBackupScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBackupSchedule",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutDeploymentDiagnosticData(value *OracleDatabaseGoldengateDeploymentPropertiesDeploymentDiagnosticData) {
	if err := o.validatePutDeploymentDiagnosticDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDeploymentDiagnosticData",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutMaintenanceConfig(value *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig) {
	if err := o.validatePutMaintenanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMaintenanceConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutMaintenanceWindow(value *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceWindow) {
	if err := o.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) PutOggData(value *OracleDatabaseGoldengateDeploymentPropertiesOggData) {
	if err := o.validatePutOggDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOggData",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetBackupSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetCpuCoreCount() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuCoreCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetDeploymentDiagnosticData() {
	_jsii_.InvokeVoid(
		o,
		"resetDeploymentDiagnosticData",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetEnvironmentType() {
	_jsii_.InvokeVoid(
		o,
		"resetEnvironmentType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetIsAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoScalingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		o,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetMaintenanceConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

