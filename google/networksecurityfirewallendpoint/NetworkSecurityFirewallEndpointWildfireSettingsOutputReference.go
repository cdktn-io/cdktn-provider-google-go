// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityfirewallendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/networksecurityfirewallendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecurityFirewallEndpointWildfireSettingsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkSecurityFirewallEndpointWildfireSettings
	SetInternalValue(val *NetworkSecurityFirewallEndpointWildfireSettings)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WildfireInlineCloudAnalysisSettings() NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference
	WildfireInlineCloudAnalysisSettingsInput() *NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	WildfireRealtimeLookupDuration() *string
	SetWildfireRealtimeLookupDuration(val *string)
	WildfireRealtimeLookupDurationInput() *string
	WildfireRealtimeLookupTimeoutAction() *string
	SetWildfireRealtimeLookupTimeoutAction(val *string)
	WildfireRealtimeLookupTimeoutActionInput() *string
	WildfireRegion() *string
	SetWildfireRegion(val *string)
	WildfireRegionInput() *string
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
	PutWildfireInlineCloudAnalysisSettings(value *NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings)
	ResetEnabled()
	ResetWildfireInlineCloudAnalysisSettings()
	ResetWildfireRealtimeLookupDuration()
	ResetWildfireRealtimeLookupTimeoutAction()
	ResetWildfireRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityFirewallEndpointWildfireSettingsOutputReference
type jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InternalValue() *NetworkSecurityFirewallEndpointWildfireSettings {
	var returns *NetworkSecurityFirewallEndpointWildfireSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireInlineCloudAnalysisSettings() NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference {
	var returns NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettingsOutputReference
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireInlineCloudAnalysisSettingsInput() *NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings {
	var returns *NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRealtimeLookupTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) WildfireRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wildfireRegionInput",
		&returns,
	)
	return returns
}


func NewNetworkSecurityFirewallEndpointWildfireSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkSecurityFirewallEndpointWildfireSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityFirewallEndpointWildfireSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkSecurityFirewallEndpoint.NetworkSecurityFirewallEndpointWildfireSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecurityFirewallEndpointWildfireSettingsOutputReference_Override(n NetworkSecurityFirewallEndpointWildfireSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkSecurityFirewallEndpoint.NetworkSecurityFirewallEndpointWildfireSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetInternalValue(val *NetworkSecurityFirewallEndpointWildfireSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRealtimeLookupDuration(val *string) {
	if err := j.validateSetWildfireRealtimeLookupDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookupDuration",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRealtimeLookupTimeoutAction(val *string) {
	if err := j.validateSetWildfireRealtimeLookupTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookupTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference)SetWildfireRegion(val *string) {
	if err := j.validateSetWildfireRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRegion",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) PutWildfireInlineCloudAnalysisSettings(value *NetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings) {
	if err := n.validatePutWildfireInlineCloudAnalysisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireInlineCloudAnalysisSettings",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireInlineCloudAnalysisSettings() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireInlineCloudAnalysisSettings",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRealtimeLookupDuration() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireRealtimeLookupDuration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRealtimeLookupTimeoutAction() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireRealtimeLookupTimeoutAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ResetWildfireRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityFirewallEndpointWildfireSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

