// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication)
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
	ResetClientId()
	ResetClientSecret()
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

// The jsii proxy struct for ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication {
	var returns *ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		c,
		"resetClientId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

