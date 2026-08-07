// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentidentityauthprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agentidentityauthprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference interface {
	cdktn.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientSecretWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientSecretWo(val *string)
	ClientSecretWoInput() *string
	ClientSecretWoVersion() *string
	SetClientSecretWoVersion(val *string)
	ClientSecretWoVersionInput() *string
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
	InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
	SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
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
	ResetClientId()
	ResetClientSecret()
	ResetClientSecretWo()
	ResetClientSecretWoVersion()
	ResetTokenUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference
type jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecretWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ClientSecretWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference {
	_init_.Initialize()

	if err := validateNewAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference_Override(a AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetClientSecretWoVersion(val *string) {
	if err := j.validateSetClientSecretWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWoVersion",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ResetClientSecretWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

