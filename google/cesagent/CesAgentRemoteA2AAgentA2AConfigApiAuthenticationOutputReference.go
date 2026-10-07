// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfig
	BearerTokenConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig
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
	InternalValue() *CesAgentRemoteA2AAgentA2AConfigApiAuthentication
	SetInternalValue(val *CesAgentRemoteA2AAgentA2AConfigApiAuthentication)
	OauthConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfigOutputReference
	OauthConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig
	ServiceAccountAuthConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfigOutputReference
	ServiceAccountAuthConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfig
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
	PutApiKeyConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig)
	PutOauthConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig)
	PutServiceAccountAuthConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAccountAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference
type jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ApiKeyConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfigOutputReference {
	var returns CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ApiKeyConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfig {
	var returns *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) BearerTokenConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfigOutputReference {
	var returns CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) BearerTokenConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig {
	var returns *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) InternalValue() *CesAgentRemoteA2AAgentA2AConfigApiAuthentication {
	var returns *CesAgentRemoteA2AAgentA2AConfigApiAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) OauthConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfigOutputReference {
	var returns CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) OauthConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig {
	var returns *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ServiceAccountAuthConfig() CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfigOutputReference {
	var returns CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAccountAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ServiceAccountAuthConfigInput() *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfig {
	var returns *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfig
	_jsii_.Get(
		j,
		"serviceAccountAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference_Override(c CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference)SetInternalValue(val *CesAgentRemoteA2AAgentA2AConfigApiAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) PutApiKeyConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationApiKeyConfig) {
	if err := c.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) PutBearerTokenConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig) {
	if err := c.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) PutOauthConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig) {
	if err := c.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) PutServiceAccountAuthConfig(value *CesAgentRemoteA2AAgentA2AConfigApiAuthenticationServiceAccountAuthConfig) {
	if err := c.validatePutServiceAccountAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceAccountAuthConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ResetServiceAccountAuthConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountAuthConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

