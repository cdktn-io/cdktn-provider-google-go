// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesDeploymentWhatsappCredentialsOutputReference interface {
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
	BusinessAccountId() *string
	SetBusinessAccountId(val *string)
	BusinessAccountIdInput() *string
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
	InternalValue() *CesDeploymentWhatsappCredentials
	SetInternalValue(val *CesDeploymentWhatsappCredentials)
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	Pin() *string
	SetPin(val *string)
	PinInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	PinWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetPinWo(val *string)
	PinWoInput() *string
	PinWoVersion() *string
	SetPinWoVersion(val *string)
	PinWoVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WabaId() *string
	SetWabaId(val *string)
	WabaIdInput() *string
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
	ResetPin()
	ResetPinWo()
	ResetPinWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesDeploymentWhatsappCredentialsOutputReference
type jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCodeWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) BusinessAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) BusinessAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ConversationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ConversationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) InternalValue() *CesDeploymentWhatsappCredentials {
	var returns *CesDeploymentWhatsappCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) Pin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PinWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PinWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PinWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) PinWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) WabaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wabaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) WabaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wabaIdInput",
		&returns,
	)
	return returns
}


func NewCesDeploymentWhatsappCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesDeploymentWhatsappCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewCesDeploymentWhatsappCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentWhatsappCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesDeploymentWhatsappCredentialsOutputReference_Override(c CesDeploymentWhatsappCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentWhatsappCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetAuthCodeWo(val *string) {
	if err := j.validateSetAuthCodeWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWo",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetAuthCodeWoVersion(val *string) {
	if err := j.validateSetAuthCodeWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWoVersion",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetBusinessAccountId(val *string) {
	if err := j.validateSetBusinessAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessAccountId",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetConversationProfileId(val *string) {
	if err := j.validateSetConversationProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationProfileId",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetInternalValue(val *CesDeploymentWhatsappCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetPin(val *string) {
	if err := j.validateSetPinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pin",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetPinWo(val *string) {
	if err := j.validateSetPinWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinWo",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetPinWoVersion(val *string) {
	if err := j.validateSetPinWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinWoVersion",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference)SetWabaId(val *string) {
	if err := j.validateSetWabaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wabaId",
		val,
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetAuthCodeWo() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCodeWo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetAuthCodeWoVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthCodeWoVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetConversationProfileId() {
	_jsii_.InvokeVoid(
		c,
		"resetConversationProfileId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetPin() {
	_jsii_.InvokeVoid(
		c,
		"resetPin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetPinWo() {
	_jsii_.InvokeVoid(
		c,
		"resetPinWo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ResetPinWoVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPinWoVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesDeploymentWhatsappCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

