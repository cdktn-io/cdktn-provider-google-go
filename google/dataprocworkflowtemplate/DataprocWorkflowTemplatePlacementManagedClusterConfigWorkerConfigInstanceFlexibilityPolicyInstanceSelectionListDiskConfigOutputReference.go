// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/dataprocworkflowtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference interface {
	cdktn.ComplexObject
	AttachedDiskConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
	AttachedDiskConfigInput() interface{}
	BootDiskProvisionedIops() *float64
	SetBootDiskProvisionedIops(val *float64)
	BootDiskProvisionedIopsInput() *float64
	BootDiskProvisionedThroughput() *float64
	SetBootDiskProvisionedThroughput(val *float64)
	BootDiskProvisionedThroughputInput() *float64
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
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
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig)
	LocalSsdInterface() *string
	SetLocalSsdInterface(val *string)
	LocalSsdInterfaceInput() *string
	NumLocalSsds() *float64
	SetNumLocalSsds(val *float64)
	NumLocalSsdsInput() *float64
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
	PutAttachedDiskConfig(value interface{})
	ResetAttachedDiskConfig()
	ResetBootDiskProvisionedIops()
	ResetBootDiskProvisionedThroughput()
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetLocalSsdInterface()
	ResetNumLocalSsds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigAttachedDiskConfigList
	_jsii_.Get(
		j,
		"attachedDiskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) AttachedDiskConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachedDiskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskProvisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) LocalSsdInterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localSsdInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) NumLocalSsdsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLocalSsdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedIops(val *float64) {
	if err := j.validateSetBootDiskProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedIops",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskProvisionedThroughput(val *float64) {
	if err := j.validateSetBootDiskProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskProvisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetLocalSsdInterface(val *string) {
	if err := j.validateSetLocalSsdInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdInterface",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetNumLocalSsds(val *float64) {
	if err := j.validateSetNumLocalSsdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numLocalSsds",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) PutAttachedDiskConfig(value interface{}) {
	if err := d.validatePutAttachedDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttachedDiskConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetAttachedDiskConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAttachedDiskConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedIops() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskProvisionedIops",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetLocalSsdInterface() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSsdInterface",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ResetNumLocalSsds() {
	_jsii_.InvokeVoid(
		d,
		"resetNumLocalSsds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListDiskConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

