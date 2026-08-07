// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAzureEventHubSettingsOutputReference interface {
	cdktn.ComplexObject
	AzureSasToken() *string
	SetAzureSasToken(val *string)
	AzureSasTokenInput() *string
	AzureStorageConnectionString() *string
	SetAzureStorageConnectionString(val *string)
	AzureStorageConnectionStringInput() *string
	AzureStorageContainer() *string
	SetAzureStorageContainer(val *string)
	AzureStorageContainerInput() *string
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
	ConsumerGroup() *string
	SetConsumerGroup(val *string)
	ConsumerGroupInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventHubConnectionString() *string
	SetEventHubConnectionString(val *string)
	EventHubConnectionStringInput() *string
	EventHubNamespace() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleFeedDetailsAzureEventHubSettings
	SetInternalValue(val *ChronicleFeedDetailsAzureEventHubSettings)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetAzureSasToken()
	ResetAzureStorageConnectionString()
	ResetAzureStorageContainer()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsAzureEventHubSettingsOutputReference
type jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureSasToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSasToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureSasTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSasTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) AzureStorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStorageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ConsumerGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ConsumerGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) EventHubNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) InternalValue() *ChronicleFeedDetailsAzureEventHubSettings {
	var returns *ChronicleFeedDetailsAzureEventHubSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAzureEventHubSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAzureEventHubSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAzureEventHubSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureEventHubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAzureEventHubSettingsOutputReference_Override(c ChronicleFeedDetailsAzureEventHubSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureEventHubSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureSasToken(val *string) {
	if err := j.validateSetAzureSasTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureSasToken",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureStorageConnectionString(val *string) {
	if err := j.validateSetAzureStorageConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageConnectionString",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetAzureStorageContainer(val *string) {
	if err := j.validateSetAzureStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureStorageContainer",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetConsumerGroup(val *string) {
	if err := j.validateSetConsumerGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroup",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetEventHubConnectionString(val *string) {
	if err := j.validateSetEventHubConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHubConnectionString",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetInternalValue(val *ChronicleFeedDetailsAzureEventHubSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureSasToken() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureSasToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureStorageConnectionString() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureStorageConnectionString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ResetAzureStorageContainer() {
	_jsii_.InvokeVoid(
		c,
		"resetAzureStorageContainer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureEventHubSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

