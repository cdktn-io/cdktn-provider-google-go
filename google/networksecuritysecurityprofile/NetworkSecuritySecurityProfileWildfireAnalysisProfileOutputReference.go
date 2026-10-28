// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/networksecuritysecurityprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference interface {
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
	InternalValue() *NetworkSecuritySecurityProfileWildfireAnalysisProfile
	SetInternalValue(val *NetworkSecuritySecurityProfileWildfireAnalysisProfile)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WildfireInlineCloudAnalysisRules() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList
	WildfireInlineCloudAnalysisRulesInput() interface{}
	WildfireInlineMlOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList
	WildfireInlineMlOverridesInput() interface{}
	WildfireInlineMlSetting() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference
	WildfireInlineMlSettingInput() *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting
	WildfireOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList
	WildfireOverridesInput() interface{}
	WildfireRealtimeLookup() interface{}
	SetWildfireRealtimeLookup(val interface{})
	WildfireRealtimeLookupInput() interface{}
	WildfireSubmissionRules() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList
	WildfireSubmissionRulesInput() interface{}
	WildfireThreatOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList
	WildfireThreatOverridesInput() interface{}
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
	PutWildfireInlineCloudAnalysisRules(value interface{})
	PutWildfireInlineMlOverrides(value interface{})
	PutWildfireInlineMlSetting(value *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting)
	PutWildfireOverrides(value interface{})
	PutWildfireSubmissionRules(value interface{})
	PutWildfireThreatOverrides(value interface{})
	ResetWildfireInlineCloudAnalysisRules()
	ResetWildfireInlineMlOverrides()
	ResetWildfireInlineMlSetting()
	ResetWildfireOverrides()
	ResetWildfireRealtimeLookup()
	ResetWildfireSubmissionRules()
	ResetWildfireThreatOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference
type jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InternalValue() *NetworkSecuritySecurityProfileWildfireAnalysisProfile {
	var returns *NetworkSecuritySecurityProfileWildfireAnalysisProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineCloudAnalysisRules() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineCloudAnalysisRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireInlineCloudAnalysisRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlOverridesList
	_jsii_.Get(
		j,
		"wildfireInlineMlOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireInlineMlOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlSetting() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSettingOutputReference
	_jsii_.Get(
		j,
		"wildfireInlineMlSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireInlineMlSettingInput() *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting {
	var returns *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting
	_jsii_.Get(
		j,
		"wildfireInlineMlSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireOverridesList
	_jsii_.Get(
		j,
		"wildfireOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireRealtimeLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireRealtimeLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireRealtimeLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireRealtimeLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireSubmissionRules() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireSubmissionRulesList
	_jsii_.Get(
		j,
		"wildfireSubmissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireSubmissionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireSubmissionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireThreatOverrides() NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList {
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireThreatOverridesList
	_jsii_.Get(
		j,
		"wildfireThreatOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) WildfireThreatOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildfireThreatOverridesInput",
		&returns,
	)
	return returns
}


func NewNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference_Override(n NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetInternalValue(val *NetworkSecuritySecurityProfileWildfireAnalysisProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference)SetWildfireRealtimeLookup(val interface{}) {
	if err := j.validateSetWildfireRealtimeLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildfireRealtimeLookup",
		val,
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineCloudAnalysisRules(value interface{}) {
	if err := n.validatePutWildfireInlineCloudAnalysisRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireInlineCloudAnalysisRules",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineMlOverrides(value interface{}) {
	if err := n.validatePutWildfireInlineMlOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireInlineMlOverrides",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireInlineMlSetting(value *NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting) {
	if err := n.validatePutWildfireInlineMlSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireInlineMlSetting",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireOverrides(value interface{}) {
	if err := n.validatePutWildfireOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireOverrides",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireSubmissionRules(value interface{}) {
	if err := n.validatePutWildfireSubmissionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireSubmissionRules",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) PutWildfireThreatOverrides(value interface{}) {
	if err := n.validatePutWildfireThreatOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWildfireThreatOverrides",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineCloudAnalysisRules() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireInlineCloudAnalysisRules",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineMlOverrides() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireInlineMlOverrides",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireInlineMlSetting() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireInlineMlSetting",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireOverrides() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireOverrides",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireRealtimeLookup() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireRealtimeLookup",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireSubmissionRules() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireSubmissionRules",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ResetWildfireThreatOverrides() {
	_jsii_.InvokeVoid(
		n,
		"resetWildfireThreatOverrides",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

