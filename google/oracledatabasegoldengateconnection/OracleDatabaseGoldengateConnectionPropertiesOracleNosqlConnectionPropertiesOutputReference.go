// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference interface {
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
	InternalValue() *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties)
	PrivateKeyFile() *string
	SetPrivateKeyFile(val *string)
	PrivateKeyFileInput() *string
	PrivateKeyPassphraseSecret() *string
	SetPrivateKeyPassphraseSecret(val *string)
	PrivateKeyPassphraseSecretInput() *string
	PublicKeyFingerprint() *string
	SetPublicKeyFingerprint(val *string)
	PublicKeyFingerprintInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TechnologyType() *string
	SetTechnologyType(val *string)
	TechnologyTypeInput() *string
	TenancyId() *string
	SetTenancyId(val *string)
	TenancyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseResourcePrincipal() interface{}
	SetUseResourcePrincipal(val interface{})
	UseResourcePrincipalInput() interface{}
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
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
	ResetPrivateKeyFile()
	ResetPrivateKeyPassphraseSecret()
	ResetPublicKeyFingerprint()
	ResetRegion()
	ResetTechnologyType()
	ResetTenancyId()
	ResetUseResourcePrincipal()
	ResetUserId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
type jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InternalValue() *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties {
	var returns *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyPassphraseSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyPassphraseSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PublicKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PublicKeyFingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TenancyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TenancyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UseResourcePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UseResourcePrincipalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference_Override(o OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateConnection.OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetInternalValue(val *OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPrivateKeyFile(val *string) {
	if err := j.validateSetPrivateKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyFile",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPrivateKeyPassphraseSecret(val *string) {
	if err := j.validateSetPrivateKeyPassphraseSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyPassphraseSecret",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPublicKeyFingerprint(val *string) {
	if err := j.validateSetPublicKeyFingerprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKeyFingerprint",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTenancyId(val *string) {
	if err := j.validateSetTenancyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancyId",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetUseResourcePrincipal(val interface{}) {
	if err := j.validateSetUseResourcePrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourcePrincipal",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPrivateKeyFile() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateKeyFile",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPrivateKeyPassphraseSecret() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateKeyPassphraseSecret",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPublicKeyFingerprint() {
	_jsii_.InvokeVoid(
		o,
		"resetPublicKeyFingerprint",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		o,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetTenancyId() {
	_jsii_.InvokeVoid(
		o,
		"resetTenancyId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetUseResourcePrincipal() {
	_jsii_.InvokeVoid(
		o,
		"resetUseResourcePrincipal",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		o,
		"resetUserId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

