// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AdditionalAttributes() OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesAdditionalAttributesList
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
	DbSystemId() *string
	SetDbSystemId(val *string)
	DbSystemIdInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties)
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
	SslCaFile() *string
	SetSslCaFile(val *string)
	SslCaFileInput() *string
	SslCertFile() *string
	SetSslCertFile(val *string)
	SslCertFileInput() *string
	SslCrlFile() *string
	SetSslCrlFile(val *string)
	SslCrlFileInput() *string
	SslKeyFile() *string
	SetSslKeyFile(val *string)
	SslKeyFileInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
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
	ResetDbSystemId()
	ResetHost()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetPort()
	ResetSecurityProtocol()
	ResetSslCaFile()
	ResetSslCertFile()
	ResetSslCrlFile()
	ResetSslKeyFile()
	ResetSslMode()
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

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) AdditionalAttributes() OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesAdditionalAttributesList {
	var returns OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesAdditionalAttributesList
	_jsii_.Get(
		j,
		"additionalAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) AdditionalAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCaFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCrlFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCrlFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslCrlFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCrlFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSslCaFile(val *string) {
	if err := j.validateSetSslCaFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSslCertFile(val *string) {
	if err := j.validateSetSslCertFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSslCrlFile(val *string) {
	if err := j.validateSetSslCrlFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCrlFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSslKeyFile(val *string) {
	if err := j.validateSetSslKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) PutAdditionalAttributes(value interface{}) {
	if err := o.validatePutAdditionalAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdditionalAttributes",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetAdditionalAttributes() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalAttributes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabase",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		o,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		o,
		"resetHost",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		o,
		"resetPort",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSslCaFile() {
	_jsii_.InvokeVoid(
		o,
		"resetSslCaFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSslCertFile() {
	_jsii_.InvokeVoid(
		o,
		"resetSslCertFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSslCrlFile() {
	_jsii_.InvokeVoid(
		o,
		"resetSslCrlFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSslKeyFile() {
	_jsii_.InvokeVoid(
		o,
		"resetSslKeyFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetSslMode() {
	_jsii_.InvokeVoid(
		o,
		"resetSslMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		o,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesMysqlConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

