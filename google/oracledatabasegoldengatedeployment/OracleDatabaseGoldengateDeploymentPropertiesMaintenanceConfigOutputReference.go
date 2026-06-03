// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference interface {
	cdktn.ComplexObject
	BundleReleaseUpgradePeriodDays() *float64
	SetBundleReleaseUpgradePeriodDays(val *float64)
	BundleReleaseUpgradePeriodDaysInput() *float64
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
	InterimReleaseUpgradePeriodDays() *float64
	SetInterimReleaseUpgradePeriodDays(val *float64)
	InterimReleaseUpgradePeriodDaysInput() *float64
	InternalValue() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	SetInternalValue(val *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig)
	IsInterimReleaseAutoUpgradeEnabled() interface{}
	SetIsInterimReleaseAutoUpgradeEnabled(val interface{})
	IsInterimReleaseAutoUpgradeEnabledInput() interface{}
	MajorReleaseUpgradePeriodDays() *float64
	SetMajorReleaseUpgradePeriodDays(val *float64)
	MajorReleaseUpgradePeriodDaysInput() *float64
	SecurityPatchUpgradePeriodDays() *float64
	SetSecurityPatchUpgradePeriodDays(val *float64)
	SecurityPatchUpgradePeriodDaysInput() *float64
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
	ResetBundleReleaseUpgradePeriodDays()
	ResetInterimReleaseUpgradePeriodDays()
	ResetIsInterimReleaseAutoUpgradeEnabled()
	ResetMajorReleaseUpgradePeriodDays()
	ResetSecurityPatchUpgradePeriodDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference
type jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) BundleReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bundleReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) BundleReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bundleReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterimReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interimReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterimReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interimReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InternalValue() *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) IsInterimReleaseAutoUpgradeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInterimReleaseAutoUpgradeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) IsInterimReleaseAutoUpgradeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isInterimReleaseAutoUpgradeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) MajorReleaseUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"majorReleaseUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) MajorReleaseUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"majorReleaseUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) SecurityPatchUpgradePeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityPatchUpgradePeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) SecurityPatchUpgradePeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityPatchUpgradePeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference_Override(o OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetBundleReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetBundleReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetInterimReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetInterimReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interimReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetInternalValue(val *OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetIsInterimReleaseAutoUpgradeEnabled(val interface{}) {
	if err := j.validateSetIsInterimReleaseAutoUpgradeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isInterimReleaseAutoUpgradeEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetMajorReleaseUpgradePeriodDays(val *float64) {
	if err := j.validateSetMajorReleaseUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"majorReleaseUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetSecurityPatchUpgradePeriodDays(val *float64) {
	if err := j.validateSetSecurityPatchUpgradePeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPatchUpgradePeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetBundleReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetBundleReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetInterimReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetInterimReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetIsInterimReleaseAutoUpgradeEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsInterimReleaseAutoUpgradeEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetMajorReleaseUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetMajorReleaseUpgradePeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ResetSecurityPatchUpgradePeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityPatchUpgradePeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesMaintenanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

