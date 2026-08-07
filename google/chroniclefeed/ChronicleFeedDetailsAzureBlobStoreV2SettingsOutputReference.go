// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference interface {
	cdktn.ComplexObject
	Authentication() ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationOutputReference
	AuthenticationInput() *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication
	AzureUri() *string
	SetAzureUri(val *string)
	AzureUriInput() *string
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
	InternalValue() *ChronicleFeedDetailsAzureBlobStoreV2Settings
	SetInternalValue(val *ChronicleFeedDetailsAzureBlobStoreV2Settings)
	MaxLookbackDays() *float64
	SetMaxLookbackDays(val *float64)
	MaxLookbackDaysInput() *float64
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
	PutAuthentication(value *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication)
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

// The jsii proxy struct for ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference
type jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) Authentication() ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationOutputReference {
	var returns ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) AuthenticationInput() *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication {
	var returns *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) AzureUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) AzureUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ChronicleServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chronicleServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) InternalValue() *ChronicleFeedDetailsAzureBlobStoreV2Settings {
	var returns *ChronicleFeedDetailsAzureBlobStoreV2Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) MaxLookbackDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLookbackDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) MaxLookbackDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLookbackDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) SourceDeletionOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDeletionOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) SourceDeletionOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDeletionOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference_Override(c ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetAzureUri(val *string) {
	if err := j.validateSetAzureUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureUri",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetInternalValue(val *ChronicleFeedDetailsAzureBlobStoreV2Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetMaxLookbackDays(val *float64) {
	if err := j.validateSetMaxLookbackDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLookbackDays",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetSourceDeletionOption(val *string) {
	if err := j.validateSetSourceDeletionOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDeletionOption",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) PutAuthentication(value *ChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication) {
	if err := c.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ResetMaxLookbackDays() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxLookbackDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ResetSourceDeletionOption() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDeletionOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureBlobStoreV2SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

