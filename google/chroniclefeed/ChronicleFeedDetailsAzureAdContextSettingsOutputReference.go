// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFeedDetailsAzureAdContextSettingsOutputReference interface {
	cdktn.ComplexObject
	AuthEndpoint() *string
	SetAuthEndpoint(val *string)
	AuthEndpointInput() *string
	Authentication() ChronicleFeedDetailsAzureAdContextSettingsAuthenticationOutputReference
	AuthenticationInput() *ChronicleFeedDetailsAzureAdContextSettingsAuthentication
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
	InternalValue() *ChronicleFeedDetailsAzureAdContextSettings
	SetInternalValue(val *ChronicleFeedDetailsAzureAdContextSettings)
	RetrieveDevices() interface{}
	SetRetrieveDevices(val interface{})
	RetrieveDevicesInput() interface{}
	RetrieveGroups() interface{}
	SetRetrieveGroups(val interface{})
	RetrieveGroupsInput() interface{}
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	PutAuthentication(value *ChronicleFeedDetailsAzureAdContextSettingsAuthentication)
	ResetAuthEndpoint()
	ResetAuthentication()
	ResetHostname()
	ResetRetrieveDevices()
	ResetRetrieveGroups()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFeedDetailsAzureAdContextSettingsOutputReference
type jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) AuthEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) AuthEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) Authentication() ChronicleFeedDetailsAzureAdContextSettingsAuthenticationOutputReference {
	var returns ChronicleFeedDetailsAzureAdContextSettingsAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) AuthenticationInput() *ChronicleFeedDetailsAzureAdContextSettingsAuthentication {
	var returns *ChronicleFeedDetailsAzureAdContextSettingsAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) InternalValue() *ChronicleFeedDetailsAzureAdContextSettings {
	var returns *ChronicleFeedDetailsAzureAdContextSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) RetrieveDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) RetrieveDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) RetrieveGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) RetrieveGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFeedDetailsAzureAdContextSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFeedDetailsAzureAdContextSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFeedDetailsAzureAdContextSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureAdContextSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFeedDetailsAzureAdContextSettingsOutputReference_Override(c ChronicleFeedDetailsAzureAdContextSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFeed.ChronicleFeedDetailsAzureAdContextSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetAuthEndpoint(val *string) {
	if err := j.validateSetAuthEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authEndpoint",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetInternalValue(val *ChronicleFeedDetailsAzureAdContextSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetRetrieveDevices(val interface{}) {
	if err := j.validateSetRetrieveDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrieveDevices",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetRetrieveGroups(val interface{}) {
	if err := j.validateSetRetrieveGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrieveGroups",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) PutAuthentication(value *ChronicleFeedDetailsAzureAdContextSettingsAuthentication) {
	if err := c.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetAuthEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetRetrieveDevices() {
	_jsii_.InvokeVoid(
		c,
		"resetRetrieveDevices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetRetrieveGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetRetrieveGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		c,
		"resetTenantId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFeedDetailsAzureAdContextSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

