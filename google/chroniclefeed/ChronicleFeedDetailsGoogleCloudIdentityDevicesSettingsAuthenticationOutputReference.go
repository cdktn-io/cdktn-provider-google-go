// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	Claims() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaimsOutputReference
	ClaimsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaims
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
	InternalValue() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication)
	RsCredentials() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentialsOutputReference
	RsCredentialsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentials
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenEndpoint() *string
	SetTokenEndpoint(val *string)
	TokenEndpointInput() *string
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
	PutClaims(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaims)
	PutRsCredentials(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentials)
	ResetClaims()
	ResetRsCredentials()
	ResetTokenEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) Claims() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaimsOutputReference {
	var returns ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaimsOutputReference
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ClaimsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaims {
	var returns *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaims
	_jsii_.Get(
		j,
		"claimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication {
	var returns *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) RsCredentials() ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentialsOutputReference {
	var returns ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentialsOutputReference
	_jsii_.Get(
		j,
		"rsCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) RsCredentialsInput() *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentials {
	var returns *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentials
	_jsii_.Get(
		j,
		"rsCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) PutClaims(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationClaims) {
	if err := c.validatePutClaimsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClaims",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) PutRsCredentials(value *ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationRsCredentials) {
	if err := c.validatePutRsCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRsCredentials",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ResetClaims() {
	_jsii_.InvokeVoid(
		c,
		"resetClaims",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ResetRsCredentials() {
	_jsii_.InvokeVoid(
		c,
		"resetRsCredentials",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

