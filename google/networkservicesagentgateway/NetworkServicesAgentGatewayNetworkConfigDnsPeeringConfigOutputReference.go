// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/networkservicesagentgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference interface {
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
	Domains() *[]*string
	SetDomains(val *[]*string)
	DomainsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig
	SetInternalValue(val *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig)
	TargetNetwork() *string
	SetTargetNetwork(val *string)
	TargetNetworkInput() *string
	TargetProject() *string
	SetTargetProject(val *string)
	TargetProjectInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference
type jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) DomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) InternalValue() *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig {
	var returns *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TargetNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TargetNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TargetProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TargetProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference_Override(n NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetDomains(val *[]*string) {
	if err := j.validateSetDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetInternalValue(val *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetTargetNetwork(val *string) {
	if err := j.validateSetTargetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNetwork",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetTargetProject(val *string) {
	if err := j.validateSetTargetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProject",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

