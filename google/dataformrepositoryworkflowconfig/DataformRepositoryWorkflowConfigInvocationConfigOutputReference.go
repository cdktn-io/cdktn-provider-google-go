// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepositoryworkflowconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dataformrepositoryworkflowconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataformRepositoryWorkflowConfigInvocationConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FullyRefreshIncrementalTablesEnabled() interface{}
	SetFullyRefreshIncrementalTablesEnabled(val interface{})
	FullyRefreshIncrementalTablesEnabledInput() interface{}
	IncludedTags() *[]*string
	SetIncludedTags(val *[]*string)
	IncludedTagsInput() *[]*string
	IncludedTargets() DataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList
	IncludedTargetsInput() interface{}
	InternalValue() *DataformRepositoryWorkflowConfigInvocationConfig
	SetInternalValue(val *DataformRepositoryWorkflowConfigInvocationConfig)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransitiveDependenciesIncluded() interface{}
	SetTransitiveDependenciesIncluded(val interface{})
	TransitiveDependenciesIncludedInput() interface{}
	TransitiveDependentsIncluded() interface{}
	SetTransitiveDependentsIncluded(val interface{})
	TransitiveDependentsIncludedInput() interface{}
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
	PutIncludedTargets(value interface{})
	ResetFullyRefreshIncrementalTablesEnabled()
	ResetIncludedTags()
	ResetIncludedTargets()
	ResetServiceAccount()
	ResetTransitiveDependenciesIncluded()
	ResetTransitiveDependentsIncluded()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataformRepositoryWorkflowConfigInvocationConfigOutputReference
type jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) FullyRefreshIncrementalTablesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullyRefreshIncrementalTablesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) FullyRefreshIncrementalTablesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullyRefreshIncrementalTablesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTargets() DataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList {
	var returns DataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsList
	_jsii_.Get(
		j,
		"includedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) IncludedTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) InternalValue() *DataformRepositoryWorkflowConfigInvocationConfig {
	var returns *DataformRepositoryWorkflowConfigInvocationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependenciesIncluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependenciesIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependenciesIncludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependenciesIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependentsIncluded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependentsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) TransitiveDependentsIncludedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitiveDependentsIncludedInput",
		&returns,
	)
	return returns
}


func NewDataformRepositoryWorkflowConfigInvocationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataformRepositoryWorkflowConfigInvocationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataformRepositoryWorkflowConfigInvocationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepositoryWorkflowConfig.DataformRepositoryWorkflowConfigInvocationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataformRepositoryWorkflowConfigInvocationConfigOutputReference_Override(d DataformRepositoryWorkflowConfigInvocationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepositoryWorkflowConfig.DataformRepositoryWorkflowConfigInvocationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetFullyRefreshIncrementalTablesEnabled(val interface{}) {
	if err := j.validateSetFullyRefreshIncrementalTablesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullyRefreshIncrementalTablesEnabled",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetIncludedTags(val *[]*string) {
	if err := j.validateSetIncludedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedTags",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetInternalValue(val *DataformRepositoryWorkflowConfigInvocationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTransitiveDependenciesIncluded(val interface{}) {
	if err := j.validateSetTransitiveDependenciesIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitiveDependenciesIncluded",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference)SetTransitiveDependentsIncluded(val interface{}) {
	if err := j.validateSetTransitiveDependentsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitiveDependentsIncluded",
		val,
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) PutIncludedTargets(value interface{}) {
	if err := d.validatePutIncludedTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludedTargets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetFullyRefreshIncrementalTablesEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetFullyRefreshIncrementalTablesEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetIncludedTags() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetIncludedTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetTransitiveDependenciesIncluded() {
	_jsii_.InvokeVoid(
		d,
		"resetTransitiveDependenciesIncluded",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ResetTransitiveDependentsIncluded() {
	_jsii_.InvokeVoid(
		d,
		"resetTransitiveDependentsIncluded",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataformRepositoryWorkflowConfigInvocationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

