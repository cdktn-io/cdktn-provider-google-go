// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference interface {
	cdktn.ComplexObject
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
	CustomFallbackMessages() *map[string]*string
	SetCustomFallbackMessages(val *map[string]*string)
	CustomFallbackMessagesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppErrorHandlingSettingsFallbackResponseConfig
	SetInternalValue(val *CesAppErrorHandlingSettingsFallbackResponseConfig)
	MaxFallbackAttempts() *float64
	SetMaxFallbackAttempts(val *float64)
	MaxFallbackAttemptsInput() *float64
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
	ResetCustomFallbackMessages()
	ResetMaxFallbackAttempts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference
type jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) CustomFallbackMessages() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customFallbackMessages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) CustomFallbackMessagesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customFallbackMessagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) InternalValue() *CesAppErrorHandlingSettingsFallbackResponseConfig {
	var returns *CesAppErrorHandlingSettingsFallbackResponseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) MaxFallbackAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFallbackAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) MaxFallbackAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFallbackAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAppErrorHandlingSettingsFallbackResponseConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppErrorHandlingSettingsFallbackResponseConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppErrorHandlingSettingsFallbackResponseConfigOutputReference_Override(c CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetCustomFallbackMessages(val *map[string]*string) {
	if err := j.validateSetCustomFallbackMessagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFallbackMessages",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetInternalValue(val *CesAppErrorHandlingSettingsFallbackResponseConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetMaxFallbackAttempts(val *float64) {
	if err := j.validateSetMaxFallbackAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFallbackAttempts",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ResetCustomFallbackMessages() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomFallbackMessages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ResetMaxFallbackAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxFallbackAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesAppErrorHandlingSettingsFallbackResponseConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

