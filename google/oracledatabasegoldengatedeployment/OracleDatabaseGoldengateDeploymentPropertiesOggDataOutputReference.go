// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracledatabasegoldengatedeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/oracledatabasegoldengatedeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference interface {
	cdktn.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminPasswordSecretVersion() *string
	SetAdminPasswordSecretVersion(val *string)
	AdminPasswordSecretVersionInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	Certificate() *string
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
	CredentialStore() *string
	Deployment() *string
	SetDeployment(val *string)
	DeploymentInput() *string
	// Experimental.
	Fqn() *string
	GroupRolesMapping() OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference
	GroupRolesMappingInput() *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping
	IdentityDomainId() *string
	InternalValue() *OracleDatabaseGoldengateDeploymentPropertiesOggData
	SetInternalValue(val *OracleDatabaseGoldengateDeploymentPropertiesOggData)
	OggVersion() *string
	SetOggVersion(val *string)
	OggVersionInput() *string
	PasswordSecretId() *string
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
	PutGroupRolesMapping(value *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping)
	ResetAdminPassword()
	ResetAdminPasswordSecretVersion()
	ResetGroupRolesMapping()
	ResetOggVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference
type jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminPasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) CredentialStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Deployment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) DeploymentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GroupRolesMapping() OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference {
	var returns OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMappingOutputReference
	_jsii_.Get(
		j,
		"groupRolesMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GroupRolesMappingInput() *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping
	_jsii_.Get(
		j,
		"groupRolesMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) IdentityDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InternalValue() *OracleDatabaseGoldengateDeploymentPropertiesOggData {
	var returns *OracleDatabaseGoldengateDeploymentPropertiesOggData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) OggVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) OggVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oggVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) PasswordSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference_Override(o OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseGoldengateDeployment.OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminPasswordSecretVersion(val *string) {
	if err := j.validateSetAdminPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetDeployment(val *string) {
	if err := j.validateSetDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deployment",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetInternalValue(val *OracleDatabaseGoldengateDeploymentPropertiesOggData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetOggVersion(val *string) {
	if err := j.validateSetOggVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oggVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) PutGroupRolesMapping(value *OracleDatabaseGoldengateDeploymentPropertiesOggDataGroupRolesMapping) {
	if err := o.validatePutGroupRolesMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGroupRolesMapping",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetAdminPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPasswordSecretVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetGroupRolesMapping() {
	_jsii_.InvokeVoid(
		o,
		"resetGroupRolesMapping",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ResetOggVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetOggVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseGoldengateDeploymentPropertiesOggDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

