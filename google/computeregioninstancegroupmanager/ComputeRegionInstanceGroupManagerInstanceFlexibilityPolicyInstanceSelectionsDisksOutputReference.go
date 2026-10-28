// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/computeregioninstancegroupmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference interface {
	cdktn.ComplexObject
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	AutoDelete() interface{}
	SetAutoDelete(val interface{})
	AutoDeleteInput() interface{}
	Boot() interface{}
	SetBoot(val interface{})
	BootInput() interface{}
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	DiskEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey
	DiskName() *string
	SetDiskName(val *string)
	DiskNameInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	GuestOsFeatures() *[]*string
	SetGuestOsFeatures(val *[]*string)
	GuestOsFeaturesInput() *[]*string
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList
	LabelsInput() interface{}
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
	ResourceManagerTags() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList
	ResourceManagerTagsInput() interface{}
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Source() *string
	SetSource(val *string)
	SourceImage() *string
	SetSourceImage(val *string)
	SourceImageEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference
	SourceImageEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey
	SourceImageInput() *string
	SourceInput() *string
	SourceSnapshot() *string
	SetSourceSnapshot(val *string)
	SourceSnapshotEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference
	SourceSnapshotEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	SourceSnapshotInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDiskEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey)
	PutLabels(value interface{})
	PutResourceManagerTags(value interface{})
	PutSourceImageEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey)
	PutSourceSnapshotEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey)
	ResetArchitecture()
	ResetAutoDelete()
	ResetBoot()
	ResetDeviceName()
	ResetDiskEncryptionKey()
	ResetDiskName()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetGuestOsFeatures()
	ResetInterface()
	ResetLabels()
	ResetMode()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	ResetResourceManagerTags()
	ResetResourcePolicies()
	ResetSource()
	ResetSourceImage()
	ResetSourceImageEncryptionKey()
	ResetSourceSnapshot()
	ResetSourceSnapshotEncryptionKey()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference
type jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) AutoDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) AutoDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Boot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) BootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey {
	var returns *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GuestOsFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GuestOsFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Labels() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourceManagerTags() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourceManagerTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceImageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey {
	var returns *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey
	_jsii_.Get(
		j,
		"sourceImageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotEncryptionKey() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotEncryptionKeyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey {
	var returns *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference_Override(c ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetAutoDelete(val interface{}) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetBoot(val interface{}) {
	if err := j.validateSetBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boot",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskName(val *string) {
	if err := j.validateSetDiskNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetGuestOsFeatures(val *[]*string) {
	if err := j.validateSetGuestOsFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestOsFeatures",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSourceImage(val *string) {
	if err := j.validateSetSourceImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImage",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSourceSnapshot(val *string) {
	if err := j.validateSetSourceSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSnapshot",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutDiskEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey) {
	if err := c.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutLabels(value interface{}) {
	if err := c.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLabels",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutResourceManagerTags(value interface{}) {
	if err := c.validatePutResourceManagerTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResourceManagerTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutSourceImageEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey) {
	if err := c.validatePutSourceImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceImageEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutSourceSnapshotEncryptionKey(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey) {
	if err := c.validatePutSourceSnapshotEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceSnapshotEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		c,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetBoot() {
	_jsii_.InvokeVoid(
		c,
		"resetBoot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskName() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetGuestOsFeatures() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestOsFeatures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		c,
		"resetInterface",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceImage() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceImageEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceImageEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceSnapshot() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceSnapshot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceSnapshotEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceSnapshotEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

