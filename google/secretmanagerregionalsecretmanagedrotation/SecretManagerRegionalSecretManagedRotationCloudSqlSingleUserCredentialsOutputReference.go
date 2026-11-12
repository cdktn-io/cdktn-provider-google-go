// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretmanagerregionalsecretmanagedrotation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/secretmanagerregionalsecretmanagedrotation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InternalValue() *SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials
	SetInternalValue(val *SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	PasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetPasswordWo(val *string)
	PasswordWoInput() *string
	PasswordWoVersion() *string
	SetPasswordWoVersion(val *string)
	PasswordWoVersionInput() *string
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
	ResetPassword()
	ResetPasswordWo()
	ResetPasswordWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference
type jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) InternalValue() *SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials {
	var returns *SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) PasswordWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) PasswordWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewSecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewSecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.secretManagerRegionalSecretManagedRotation.SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference_Override(s SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.secretManagerRegionalSecretManagedRotation.SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetInternalValue(val *SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetPasswordWoVersion(val *string) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		s,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretManagerRegionalSecretManagedRotationCloudSqlSingleUserCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

