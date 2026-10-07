// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentconnectivitytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/networkservicesagentconnectivitytemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference interface {
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
	DnsPeeringConfig() NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference
	DnsPeeringConfigInput() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfig
	SetInternalValue(val *NetworkServicesAgentConnectivityTemplateEgressNetworkConfig)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsConfig() NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference
	TlsConfigInput() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig
	VpcEgress() *string
	SetVpcEgress(val *string)
	VpcEgressInput() *string
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
	PutDnsPeeringConfig(value *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig)
	PutTlsConfig(value *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig)
	ResetDnsPeeringConfig()
	ResetNetworkAttachment()
	ResetTlsConfig()
	ResetVpcEgress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference
type jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) DnsPeeringConfig() NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference {
	var returns NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference
	_jsii_.Get(
		j,
		"dnsPeeringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) DnsPeeringConfigInput() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig {
	var returns *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig
	_jsii_.Get(
		j,
		"dnsPeeringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InternalValue() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfig {
	var returns *NetworkServicesAgentConnectivityTemplateEgressNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TlsConfig() NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference {
	var returns NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TlsConfigInput() *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig {
	var returns *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) VpcEgress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) VpcEgressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEgressInput",
		&returns,
	)
	return returns
}


func NewNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentConnectivityTemplate.NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference_Override(n NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentConnectivityTemplate.NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetInternalValue(val *NetworkServicesAgentConnectivityTemplateEgressNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetVpcEgress(val *string) {
	if err := j.validateSetVpcEgressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEgress",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) PutDnsPeeringConfig(value *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig) {
	if err := n.validatePutDnsPeeringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDnsPeeringConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) PutTlsConfig(value *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig) {
	if err := n.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetDnsPeeringConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetDnsPeeringConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetVpcEgress() {
	_jsii_.InvokeVoid(
		n,
		"resetVpcEgress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

