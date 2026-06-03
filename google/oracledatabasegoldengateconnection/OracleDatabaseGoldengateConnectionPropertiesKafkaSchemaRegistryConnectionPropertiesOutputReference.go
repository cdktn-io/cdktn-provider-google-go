// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	InternalValue() *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties)
	KeyStoreFile() *string
	SetKeyStoreFile(val *string)
	KeyStoreFileInput() *string
	KeyStorePassword() *string
	SetKeyStorePassword(val *string)
	KeyStorePasswordInput() *string
	KeyStorePasswordSecretVersion() *string
	SetKeyStorePasswordSecretVersion(val *string)
	KeyStorePasswordSecretVersionInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordSecretVersion() *string
	SetPasswordSecretVersion(val *string)
	PasswordSecretVersionInput() *string
	SslKeyPassword() *string
	SetSslKeyPassword(val *string)
	SslKeyPasswordInput() *string
	SslKeyPasswordSecretVersion() *string
	SetSslKeyPasswordSecretVersion(val *string)
	SslKeyPasswordSecretVersionInput() *string
	TechnologyType() *string
	SetTechnologyType(val *string)
	TechnologyTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrustStoreFile() *string
	SetTrustStoreFile(val *string)
	TrustStoreFileInput() *string
	TrustStorePassword() *string
	SetTrustStorePassword(val *string)
	TrustStorePasswordInput() *string
	TrustStorePasswordSecretVersion() *string
	SetTrustStorePasswordSecretVersion(val *string)
	TrustStorePasswordSecretVersionInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetAuthenticationType()
	ResetKeyStoreFile()
	ResetKeyStorePassword()
	ResetKeyStorePasswordSecretVersion()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetSslKeyPassword()
	ResetSslKeyPasswordSecretVersion()
	ResetTechnologyType()
	ResetTrustStoreFile()
	ResetTrustStorePassword()
	ResetTrustStorePasswordSecretVersion()
	ResetUrl()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) KeyStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) SslKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) SslKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) SslKeyPasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) SslKeyPasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) TrustStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetKeyStoreFile(val *string) {
	if err := j.validateSetKeyStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStoreFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetKeyStorePassword(val *string) {
	if err := j.validateSetKeyStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetKeyStorePasswordSecretVersion(val *string) {
	if err := j.validateSetKeyStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetSslKeyPassword(val *string) {
	if err := j.validateSetSslKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetSslKeyPasswordSecretVersion(val *string) {
	if err := j.validateSetSslKeyPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTrustStoreFile(val *string) {
	if err := j.validateSetTrustStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTrustStorePassword(val *string) {
	if err := j.validateSetTrustStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetTrustStorePasswordSecretVersion(val *string) {
	if err := j.validateSetTrustStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		o,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetKeyStoreFile() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyStoreFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetKeyStorePassword() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyStorePassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetKeyStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetKeyStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetSslKeyPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetSslKeyPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetSslKeyPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetSslKeyPasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		o,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetTrustStoreFile() {
	_jsii_.InvokeVoid(
		o,
		"resetTrustStoreFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetTrustStorePassword() {
	_jsii_.InvokeVoid(
		o,
		"resetTrustStorePassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetTrustStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetTrustStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

