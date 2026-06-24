// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference interface {
	cdktn.ComplexObject
	BucketUri() *string
	SetBucketUri(val *string)
	BucketUriInput() *string
	ChronicleServiceAccount() *string
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
	InternalValue() *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	SetInternalValue(val *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings)
	MaxLookbackDays() *float64
	SetMaxLookbackDays(val *float64)
	MaxLookbackDaysInput() *float64
	PubsubSubscription() *string
	SetPubsubSubscription(val *string)
	PubsubSubscriptionInput() *string
	SourceDeletionOption() *string
	SetSourceDeletionOption(val *string)
	SourceDeletionOptionInput() *string
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
	ResetMaxLookbackDays()
	ResetSourceDeletionOption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference
type jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) BucketUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) BucketUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ChronicleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chronicleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) InternalValue() *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings {
	var returns *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) MaxLookbackDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLookbackDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) MaxLookbackDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLookbackDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) PubsubSubscription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) PubsubSubscriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) SourceDeletionOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDeletionOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) SourceDeletionOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDeletionOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference_Override(c ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetBucketUri(val *string) {
	if err := j.validateSetBucketUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketUri",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetInternalValue(val *ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetMaxLookbackDays(val *float64) {
	if err := j.validateSetMaxLookbackDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLookbackDays",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetPubsubSubscription(val *string) {
	if err := j.validateSetPubsubSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubSubscription",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetSourceDeletionOption(val *string) {
	if err := j.validateSetSourceDeletionOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDeletionOption",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ResetMaxLookbackDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxLookbackDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ResetSourceDeletionOption() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDeletionOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsGoogleCloudStorageEventDrivenSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

