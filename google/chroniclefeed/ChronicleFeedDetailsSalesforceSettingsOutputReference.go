// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsSalesforceSettingsOutputReference interface {
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
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() *ChronicleFeedDetailsSalesforceSettings
	SetInternalValue(val *ChronicleFeedDetailsSalesforceSettings)
	OauthJwtCredentials() ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentialsOutputReference
	OauthJwtCredentialsInput() *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials
	OauthPasswordGrantAuth() ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuthOutputReference
	OauthPasswordGrantAuthInput() *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth
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
	PutOauthJwtCredentials(value *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials)
	PutOauthPasswordGrantAuth(value *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth)
	ResetHostname()
	ResetOauthJwtCredentials()
	ResetOauthPasswordGrantAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsSalesforceSettingsOutputReference
type jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) InternalValue() *ChronicleFeedDetailsSalesforceSettings {
	var returns *ChronicleFeedDetailsSalesforceSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) OauthJwtCredentials() ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentialsOutputReference {
	var returns ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauthJwtCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) OauthJwtCredentialsInput() *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials {
	var returns *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials
	_jsii_.Get(
		j,
		"oauthJwtCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) OauthPasswordGrantAuth() ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuthOutputReference {
	var returns ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuthOutputReference
	_jsii_.Get(
		j,
		"oauthPasswordGrantAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) OauthPasswordGrantAuthInput() *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth {
	var returns *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth
	_jsii_.Get(
		j,
		"oauthPasswordGrantAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsSalesforceSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsSalesforceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsSalesforceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsSalesforceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsSalesforceSettingsOutputReference_Override(c ChronicleFeedDetailsSalesforceSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsSalesforceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetInternalValue(val *ChronicleFeedDetailsSalesforceSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) PutOauthJwtCredentials(value *ChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials) {
	if err := c.validatePutOauthJwtCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthJwtCredentials",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) PutOauthPasswordGrantAuth(value *ChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth) {
	if err := c.validatePutOauthPasswordGrantAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthPasswordGrantAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ResetOauthJwtCredentials() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthJwtCredentials",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ResetOauthPasswordGrantAuth() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthPasswordGrantAuth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsSalesforceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

