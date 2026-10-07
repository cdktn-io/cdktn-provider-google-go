// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/dialogflowtool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowToolOpenApiSpecAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ApiKeyConfig() DialogflowToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	ApiKeyConfigInput() *DialogflowToolOpenApiSpecAuthenticationApiKeyConfig
	BearerTokenConfig() DialogflowToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	BearerTokenConfigInput() *DialogflowToolOpenApiSpecAuthenticationBearerTokenConfig
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
	InternalValue() *DialogflowToolOpenApiSpecAuthentication
	SetInternalValue(val *DialogflowToolOpenApiSpecAuthentication)
	OauthConfig() DialogflowToolOpenApiSpecAuthenticationOauthConfigOutputReference
	OauthConfigInput() *DialogflowToolOpenApiSpecAuthenticationOauthConfig
	ServiceAgentAuthConfig() DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	ServiceAgentAuthConfigInput() *DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig
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
	PutApiKeyConfig(value *DialogflowToolOpenApiSpecAuthenticationApiKeyConfig)
	PutBearerTokenConfig(value *DialogflowToolOpenApiSpecAuthenticationBearerTokenConfig)
	PutOauthConfig(value *DialogflowToolOpenApiSpecAuthenticationOauthConfig)
	PutServiceAgentAuthConfig(value *DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig)
	ResetApiKeyConfig()
	ResetBearerTokenConfig()
	ResetOauthConfig()
	ResetServiceAgentAuthConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowToolOpenApiSpecAuthenticationOutputReference
type jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfig() DialogflowToolOpenApiSpecAuthenticationApiKeyConfigOutputReference {
	var returns DialogflowToolOpenApiSpecAuthenticationApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ApiKeyConfigInput() *DialogflowToolOpenApiSpecAuthenticationApiKeyConfig {
	var returns *DialogflowToolOpenApiSpecAuthenticationApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfig() DialogflowToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference {
	var returns DialogflowToolOpenApiSpecAuthenticationBearerTokenConfigOutputReference
	_jsii_.Get(
		j,
		"bearerTokenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) BearerTokenConfigInput() *DialogflowToolOpenApiSpecAuthenticationBearerTokenConfig {
	var returns *DialogflowToolOpenApiSpecAuthenticationBearerTokenConfig
	_jsii_.Get(
		j,
		"bearerTokenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) InternalValue() *DialogflowToolOpenApiSpecAuthentication {
	var returns *DialogflowToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) OauthConfig() DialogflowToolOpenApiSpecAuthenticationOauthConfigOutputReference {
	var returns DialogflowToolOpenApiSpecAuthenticationOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) OauthConfigInput() *DialogflowToolOpenApiSpecAuthenticationOauthConfig {
	var returns *DialogflowToolOpenApiSpecAuthenticationOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfig() DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference {
	var returns DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfigOutputReference
	_jsii_.Get(
		j,
		"serviceAgentAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ServiceAgentAuthConfigInput() *DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig {
	var returns *DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig
	_jsii_.Get(
		j,
		"serviceAgentAuthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowToolOpenApiSpecAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowToolOpenApiSpecAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowToolOpenApiSpecAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowToolOpenApiSpecAuthenticationOutputReference_Override(d DialogflowToolOpenApiSpecAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolOpenApiSpecAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference)SetInternalValue(val *DialogflowToolOpenApiSpecAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) PutApiKeyConfig(value *DialogflowToolOpenApiSpecAuthenticationApiKeyConfig) {
	if err := d.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) PutBearerTokenConfig(value *DialogflowToolOpenApiSpecAuthenticationBearerTokenConfig) {
	if err := d.validatePutBearerTokenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBearerTokenConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) PutOauthConfig(value *DialogflowToolOpenApiSpecAuthenticationOauthConfig) {
	if err := d.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) PutServiceAgentAuthConfig(value *DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig) {
	if err := d.validatePutServiceAgentAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceAgentAuthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ResetBearerTokenConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBearerTokenConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ResetServiceAgentAuthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAgentAuthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

