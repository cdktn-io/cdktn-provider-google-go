// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference interface {
	cdktn.ComplexObject
	AwsIamRoleArn() *string
	SetAwsIamRoleArn(val *string)
	AwsIamRoleArnInput() *string
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
	InternalValue() *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth
	SetInternalValue(val *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth)
	SubjectId() *string
	SetSubjectId(val *string)
	SubjectIdInput() *string
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
	ResetAwsIamRoleArn()
	ResetSubjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference
type jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) AwsIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) AwsIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) InternalValue() *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth {
	var returns *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) SubjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) SubjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference_Override(c ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetAwsIamRoleArn(val *string) {
	if err := j.validateSetAwsIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetInternalValue(val *ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetSubjectId(val *string) {
	if err := j.validateSetSubjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectId",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ResetAwsIamRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsIamRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ResetSubjectId() {
	_jsii_.InvokeVoid(
		c,
		"resetSubjectId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAmazonS3V2SettingsAuthenticationAwsIamRoleAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

