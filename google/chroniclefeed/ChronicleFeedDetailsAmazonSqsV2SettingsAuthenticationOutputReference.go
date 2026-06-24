// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference interface {
	cdktn.ComplexObject
	AwsIamRoleAuth() ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference
	AwsIamRoleAuthInput() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth
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
	InternalValue() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication
	SetInternalValue(val *ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication)
	SqsV2AccessKeySecretAuth() ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference
	SqsV2AccessKeySecretAuthInput() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth
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
	PutAwsIamRoleAuth(value *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth)
	PutSqsV2AccessKeySecretAuth(value *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference
type jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) AwsIamRoleAuth() ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference {
	var returns ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuthOutputReference
	_jsii_.Get(
		j,
		"awsIamRoleAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) AwsIamRoleAuthInput() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth {
	var returns *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth
	_jsii_.Get(
		j,
		"awsIamRoleAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InternalValue() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication {
	var returns *ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) SqsV2AccessKeySecretAuth() ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference {
	var returns ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuthOutputReference
	_jsii_.Get(
		j,
		"sqsV2AccessKeySecretAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) SqsV2AccessKeySecretAuthInput() *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth {
	var returns *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth
	_jsii_.Get(
		j,
		"sqsV2AccessKeySecretAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference_Override(c ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetInternalValue(val *ChronicleFeedDetailsAmazonSqsV2SettingsAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) PutAwsIamRoleAuth(value *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationAwsIamRoleAuth) {
	if err := c.validatePutAwsIamRoleAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsIamRoleAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) PutSqsV2AccessKeySecretAuth(value *ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationSqsV2AccessKeySecretAuth) {
	if err := c.validatePutSqsV2AccessKeySecretAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqsV2AccessKeySecretAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonSqsV2SettingsAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

