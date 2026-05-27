// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/cesdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference interface {
	cdktn.ComplexObject
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	EnableOriginCheck() interface{}
	SetEnableOriginCheck(val interface{})
	EnableOriginCheckInput() interface{}
	EnablePublicAccess() interface{}
	SetEnablePublicAccess(val interface{})
	EnablePublicAccessInput() interface{}
	EnableRecaptcha() interface{}
	SetEnableRecaptcha(val interface{})
	EnableRecaptchaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings
	SetInternalValue(val *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings)
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
	ResetAllowedOrigins()
	ResetEnableOriginCheck()
	ResetEnablePublicAccess()
	ResetEnableRecaptcha()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference
type jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableOriginCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOriginCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableOriginCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOriginCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnablePublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnablePublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableRecaptcha() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRecaptcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) EnableRecaptchaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRecaptchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InternalValue() *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings {
	var returns *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference_Override(c CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesDeployment.CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetAllowedOrigins(val *[]*string) {
	if err := j.validateSetAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnableOriginCheck(val interface{}) {
	if err := j.validateSetEnableOriginCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOriginCheck",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnablePublicAccess(val interface{}) {
	if err := j.validateSetEnablePublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicAccess",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetEnableRecaptcha(val interface{}) {
	if err := j.validateSetEnableRecaptchaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRecaptcha",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetInternalValue(val *CesDeploymentChannelProfileWebWidgetConfigSecuritySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnableOriginCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableOriginCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnablePublicAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePublicAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ResetEnableRecaptcha() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableRecaptcha",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesDeploymentChannelProfileWebWidgetConfigSecuritySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

