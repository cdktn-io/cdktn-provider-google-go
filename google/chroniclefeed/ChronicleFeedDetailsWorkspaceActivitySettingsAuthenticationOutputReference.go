// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	Claims() ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaimsOutputReference
	ClaimsInput() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaims
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
	InternalValue() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsWorkspaceActivitySettingsAuthentication)
	RsCredentials() ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentialsOutputReference
	RsCredentialsInput() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentials
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
	PutClaims(value *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaims)
	PutRsCredentials(value *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentials)
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

// The jsii proxy struct for ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) Claims() ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaimsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaimsOutputReference
	_jsii_.Get(
		j,
		"claims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ClaimsInput() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaims {
	var returns *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaims
	_jsii_.Get(
		j,
		"claimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthentication {
	var returns *ChronicleFeedDetailsWorkspaceActivitySettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) RsCredentials() ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentialsOutputReference {
	var returns ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentialsOutputReference
	_jsii_.Get(
		j,
		"rsCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) RsCredentialsInput() *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentials {
	var returns *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentials
	_jsii_.Get(
		j,
		"rsCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsWorkspaceActivitySettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) PutClaims(value *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationClaims) {
	if err := c.validatePutClaimsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClaims",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) PutRsCredentials(value *ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationRsCredentials) {
	if err := c.validatePutRsCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRsCredentials",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ResetClaims() {
	_jsii_.InvokeVoid(
		c,
		"resetClaims",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ResetRsCredentials() {
	_jsii_.InvokeVoid(
		c,
		"resetRsCredentials",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsWorkspaceActivitySettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

