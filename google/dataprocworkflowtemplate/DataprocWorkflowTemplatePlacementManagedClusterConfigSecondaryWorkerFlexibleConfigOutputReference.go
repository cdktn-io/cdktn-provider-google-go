// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dataprocworkflowtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference interface {
	cdktn.ComplexObject
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
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
	CoreCount() *float64
	SetCoreCount(val *float64)
	CoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HardwareClass() *string
	SetHardwareClass(val *string)
	HardwareClassInput() *string
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfig
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfig)
	NumInstances() *float64
	SetNumInstances(val *float64)
	NumInstancesInput() *float64
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
	ResetBootDiskSizeGb()
	ResetCoreCount()
	ResetHardwareClass()
	ResetNumInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) CoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) CoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) HardwareClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) HardwareClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hardwareClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) NumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) NumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetCoreCount(val *float64) {
	if err := j.validateSetCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreCount",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetHardwareClass(val *string) {
	if err := j.validateSetHardwareClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardwareClass",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetNumInstances(val *float64) {
	if err := j.validateSetNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numInstances",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		d,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ResetCoreCount() {
	_jsii_.InvokeVoid(
		d,
		"resetCoreCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ResetHardwareClass() {
	_jsii_.InvokeVoid(
		d,
		"resetHardwareClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ResetNumInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetNumInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerFlexibleConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

