// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataformrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/dataformrepository/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataformRepositoryGitRemoteSettingsOutputReference interface {
	cdktn.ComplexObject
	AuthenticationTokenSecretVersion() *string
	SetAuthenticationTokenSecretVersion(val *string)
	AuthenticationTokenSecretVersionInput() *string
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
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	// Experimental.
	Fqn() *string
	GitRepositoryLink() *string
	SetGitRepositoryLink(val *string)
	GitRepositoryLinkInput() *string
	InternalValue() *DataformRepositoryGitRemoteSettings
	SetInternalValue(val *DataformRepositoryGitRemoteSettings)
	SshAuthenticationConfig() DataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference
	SshAuthenticationConfigInput() *DataformRepositoryGitRemoteSettingsSshAuthenticationConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenStatus() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutSshAuthenticationConfig(value *DataformRepositoryGitRemoteSettingsSshAuthenticationConfig)
	ResetAuthenticationTokenSecretVersion()
	ResetGitRepositoryLink()
	ResetSshAuthenticationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataformRepositoryGitRemoteSettingsOutputReference
type jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) AuthenticationTokenSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTokenSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) AuthenticationTokenSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTokenSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GitRepositoryLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GitRepositoryLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) InternalValue() *DataformRepositoryGitRemoteSettings {
	var returns *DataformRepositoryGitRemoteSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) SshAuthenticationConfig() DataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference {
	var returns DataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference
	_jsii_.Get(
		j,
		"sshAuthenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) SshAuthenticationConfigInput() *DataformRepositoryGitRemoteSettingsSshAuthenticationConfig {
	var returns *DataformRepositoryGitRemoteSettingsSshAuthenticationConfig
	_jsii_.Get(
		j,
		"sshAuthenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) TokenStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewDataformRepositoryGitRemoteSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataformRepositoryGitRemoteSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataformRepositoryGitRemoteSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepository.DataformRepositoryGitRemoteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataformRepositoryGitRemoteSettingsOutputReference_Override(d DataformRepositoryGitRemoteSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataformRepository.DataformRepositoryGitRemoteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetAuthenticationTokenSecretVersion(val *string) {
	if err := j.validateSetAuthenticationTokenSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationTokenSecretVersion",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetGitRepositoryLink(val *string) {
	if err := j.validateSetGitRepositoryLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitRepositoryLink",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetInternalValue(val *DataformRepositoryGitRemoteSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) PutSshAuthenticationConfig(value *DataformRepositoryGitRemoteSettingsSshAuthenticationConfig) {
	if err := d.validatePutSshAuthenticationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSshAuthenticationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ResetAuthenticationTokenSecretVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationTokenSecretVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ResetGitRepositoryLink() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepositoryLink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ResetSshAuthenticationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSshAuthenticationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataformRepositoryGitRemoteSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

