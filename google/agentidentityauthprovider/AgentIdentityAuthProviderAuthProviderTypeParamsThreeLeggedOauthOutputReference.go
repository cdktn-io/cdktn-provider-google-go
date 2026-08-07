// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentidentityauthprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agentidentityauthprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference interface {
	cdktn.ComplexObject
	AuthorizationUrl() *string
	SetAuthorizationUrl(val *string)
	AuthorizationUrlInput() *string
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
	DefaultContinueUri() *string
	SetDefaultContinueUri(val *string)
	DefaultContinueUriInput() *string
	EnablePkce() interface{}
	SetEnablePkce(val interface{})
	EnablePkceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth)
	RedirectUrl() *string
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
	ResetAuthorizationUrl()
	ResetClientId()
	ResetClientSecret()
	ResetClientSecretWo()
	ResetClientSecretWoVersion()
	ResetDefaultContinueUri()
	ResetEnablePkce()
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

// The jsii proxy struct for AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
type jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) AuthorizationUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) DefaultContinueUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContinueUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) DefaultContinueUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContinueUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) EnablePkce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) EnablePkceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InternalValue() *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth {
	var returns *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference {
	_init_.Initialize()

	if err := validateNewAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference_Override(a AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agentIdentityAuthProvider.AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetAuthorizationUrl(val *string) {
	if err := j.validateSetAuthorizationUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationUrl",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecretWoVersion(val *string) {
	if err := j.validateSetClientSecretWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWoVersion",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetDefaultContinueUri(val *string) {
	if err := j.validateSetDefaultContinueUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultContinueUri",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetEnablePkce(val interface{}) {
	if err := j.validateSetEnablePkceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePkce",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetInternalValue(val *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetAuthorizationUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecretWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetDefaultContinueUri() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultContinueUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetEnablePkce() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePkce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

