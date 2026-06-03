// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AdditionalAttributes() OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesAdditionalAttributesList
	AdditionalAttributesInput() interface{}
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordSecretVersion() *string
	SetPasswordSecretVersion(val *string)
	PasswordSecretVersionInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	ServerCertificateValidationRequired() interface{}
	SetServerCertificateValidationRequired(val interface{})
	ServerCertificateValidationRequiredInput() interface{}
	SslCaFile() *string
	SetSslCaFile(val *string)
	SslCaFileInput() *string
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
	PutAdditionalAttributes(value interface{})
	ResetAdditionalAttributes()
	ResetDatabase()
	ResetHost()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetPort()
	ResetSecurityProtocol()
	ResetServerCertificateValidationRequired()
	ResetSslCaFile()
	ResetTechnologyType()
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

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) AdditionalAttributes() OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesAdditionalAttributesList {
	var returns OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesAdditionalAttributesList
	_jsii_.Get(
		j,
		"additionalAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) AdditionalAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ServerCertificateValidationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverCertificateValidationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ServerCertificateValidationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverCertificateValidationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) SslCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) SslCaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetServerCertificateValidationRequired(val interface{}) {
	if err := j.validateSetServerCertificateValidationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateValidationRequired",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetSslCaFile(val *string) {
	if err := j.validateSetSslCaFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) PutAdditionalAttributes(value interface{}) {
	if err := o.validatePutAdditionalAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdditionalAttributes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetAdditionalAttributes() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalAttributes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabase",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		o,
		"resetHost",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		o,
		"resetPort",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetServerCertificateValidationRequired() {
	_jsii_.InvokeVoid(
		o,
		"resetServerCertificateValidationRequired",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetSslCaFile() {
	_jsii_.InvokeVoid(
		o,
		"resetSslCaFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		o,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

