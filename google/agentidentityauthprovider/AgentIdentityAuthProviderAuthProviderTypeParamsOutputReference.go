// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentidentityauthprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agentidentityauthprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference interface {
	cdktn.ComplexObject
	ApiKey() AgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference
	ApiKeyInput() *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey
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
	GeAuthProvider() AgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList
	InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParams
	SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParams)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThreeLeggedOauth() AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
	ThreeLeggedOauthInput() *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	TwoLeggedOauth() AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference
	TwoLeggedOauthInput() *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
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
	PutApiKey(value *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey)
	PutThreeLeggedOauth(value *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth)
	PutTwoLeggedOauth(value *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth)
	ResetApiKey()
	ResetThreeLeggedOauth()
	ResetTwoLeggedOauth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference
type jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ApiKey() AgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference {
	var returns AgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ApiKeyInput() *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GeAuthProvider() AgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList {
	var returns AgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList
	_jsii_.Get(
		j,
		"geAuthProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParams {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ThreeLeggedOauth() AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference {
	var returns AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
	_jsii_.Get(
		j,
		"threeLeggedOauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ThreeLeggedOauthInput() *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	_jsii_.Get(
		j,
		"threeLeggedOauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TwoLeggedOauth() AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference {
	var returns AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference
	_jsii_.Get(
		j,
		"twoLeggedOauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TwoLeggedOauthInput() *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
	_jsii_.Get(
		j,
		"twoLeggedOauthInput",
		&returns,
	)
	return returns
}


func NewAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference {
	_init_.Initialize()

	if err := validateNewAgentIdentityAuthProviderAuthProviderTypeParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference_Override(a AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutApiKey(value *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey) {
	if err := a.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutThreeLeggedOauth(value *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth) {
	if err := a.validatePutThreeLeggedOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreeLeggedOauth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutTwoLeggedOauth(value *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth) {
	if err := a.validatePutTwoLeggedOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTwoLeggedOauth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetThreeLeggedOauth() {
	_jsii_.InvokeVoid(
		a,
		"resetThreeLeggedOauth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetTwoLeggedOauth() {
	_jsii_.InvokeVoid(
		a,
		"resetTwoLeggedOauth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

