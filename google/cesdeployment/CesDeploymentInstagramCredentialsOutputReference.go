// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesDeploymentInstagramCredentialsOutputReference interface {
	cdktn.ComplexObject
	AuthCode() *string
	SetAuthCode(val *string)
	AuthCodeInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AuthCodeWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAuthCodeWo(val *string)
	AuthCodeWoInput() *string
	AuthCodeWoVersion() *string
	SetAuthCodeWoVersion(val *string)
	AuthCodeWoVersionInput() *string
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
	ConversationProfileId() *string
	SetConversationProfileId(val *string)
	ConversationProfileIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesDeploymentInstagramCredentials
	SetInternalValue(val *CesDeploymentInstagramCredentials)
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
	ResetAuthCode()
	ResetAuthCodeWo()
	ResetAuthCodeWoVersion()
	ResetConversationProfileId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesDeploymentInstagramCredentialsOutputReference
type jsiiProxy_CesDeploymentInstagramCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCodeWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCodeWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCodeWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) AuthCodeWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ConversationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ConversationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) InternalValue() *CesDeploymentInstagramCredentials {
	var returns *CesDeploymentInstagramCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesDeploymentInstagramCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesDeploymentInstagramCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewCesDeploymentInstagramCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesDeploymentInstagramCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentInstagramCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesDeploymentInstagramCredentialsOutputReference_Override(c CesDeploymentInstagramCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentInstagramCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetAuthCodeWo(val *string) {
	if err := j.validateSetAuthCodeWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWo",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetAuthCodeWoVersion(val *string) {
	if err := j.validateSetAuthCodeWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWoVersion",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetConversationProfileId(val *string) {
	if err := j.validateSetConversationProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationProfileId",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetInternalValue(val *CesDeploymentInstagramCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ResetAuthCodeWo() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCodeWo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ResetAuthCodeWoVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCodeWoVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ResetConversationProfileId() {
	_jsii_.InvokeVoid(
		c,
		"resetConversationProfileId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesDeploymentInstagramCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

