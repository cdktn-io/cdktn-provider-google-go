// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference interface {
	cdktn.ComplexObject
	AccessKeyId() *string
	SetAccessKeyId(val *string)
	AccessKeyIdInput() *string
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
	InternalValue() *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth
	SetInternalValue(val *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth)
	SecretAccessKey() *string
	SetSecretAccessKey(val *string)
	SecretAccessKeyInput() *string
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
	ResetAccessKeyId()
	ResetSecretAccessKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference
type jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) AccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) InternalValue() *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth {
	var returns *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) SecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) SecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference_Override(c ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetAccessKeyId(val *string) {
	if err := j.validateSetAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeyId",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetInternalValue(val *ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetSecretAccessKey(val *string) {
	if err := j.validateSetSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKey",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ResetAccessKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ResetSecretAccessKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretAccessKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsSettingsAuthenticationSqsAccessKeySecretAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

