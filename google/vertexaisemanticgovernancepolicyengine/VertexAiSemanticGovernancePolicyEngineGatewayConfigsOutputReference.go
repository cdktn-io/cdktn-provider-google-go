// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaisemanticgovernancepolicyengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/vertexaisemanticgovernancepolicyengine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference interface {
	cdktn.ComplexObject
	AllowedProjects() *[]*string
	SetAllowedProjects(val *[]*string)
	AllowedProjectsInput() *[]*string
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
	DnsRecord() *string
	DnsZoneName() *string
	SetDnsZoneName(val *string)
	DnsZoneNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	PscEndpoint() *string
	State() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
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
	ResetAllowedProjects()
	ResetDnsZoneName()
	ResetNetwork()
	ResetSubnetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference
type jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) AllowedProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) AllowedProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) DnsRecord() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) DnsZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) DnsZoneNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) PscEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSemanticGovernancePolicyEngine.VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference_Override(v VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSemanticGovernancePolicyEngine.VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetAllowedProjects(val *[]*string) {
	if err := j.validateSetAllowedProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedProjects",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetDnsZoneName(val *string) {
	if err := j.validateSetDnsZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsZoneName",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ResetAllowedProjects() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowedProjects",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ResetDnsZoneName() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsZoneName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetNetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiSemanticGovernancePolicyEngineGatewayConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

