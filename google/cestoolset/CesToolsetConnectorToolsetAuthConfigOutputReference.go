// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/cestoolset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesToolsetConnectorToolsetAuthConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolsetConnectorToolsetAuthConfig
	SetInternalValue(val *CesToolsetConnectorToolsetAuthConfig)
	Oauth2AuthCodeConfig() CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfigOutputReference
	Oauth2AuthCodeConfigInput() *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig
	Oauth2JwtBearerConfig() CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfigOutputReference
	Oauth2JwtBearerConfigInput() *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig
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
	PutOauth2AuthCodeConfig(value *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig)
	PutOauth2JwtBearerConfig(value *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig)
	ResetOauth2AuthCodeConfig()
	ResetOauth2JwtBearerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolsetConnectorToolsetAuthConfigOutputReference
type jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) InternalValue() *CesToolsetConnectorToolsetAuthConfig {
	var returns *CesToolsetConnectorToolsetAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Oauth2AuthCodeConfig() CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfigOutputReference {
	var returns CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthCodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Oauth2AuthCodeConfigInput() *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig {
	var returns *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig
	_jsii_.Get(
		j,
		"oauth2AuthCodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Oauth2JwtBearerConfig() CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfigOutputReference {
	var returns CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2JwtBearerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Oauth2JwtBearerConfigInput() *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig {
	var returns *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig
	_jsii_.Get(
		j,
		"oauth2JwtBearerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolsetConnectorToolsetAuthConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesToolsetConnectorToolsetAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetConnectorToolsetAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetConnectorToolsetAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolsetConnectorToolsetAuthConfigOutputReference_Override(c CesToolsetConnectorToolsetAuthConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetConnectorToolsetAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference)SetInternalValue(val *CesToolsetConnectorToolsetAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) PutOauth2AuthCodeConfig(value *CesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig) {
	if err := c.validatePutOauth2AuthCodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauth2AuthCodeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) PutOauth2JwtBearerConfig(value *CesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig) {
	if err := c.validatePutOauth2JwtBearerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauth2JwtBearerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ResetOauth2AuthCodeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOauth2AuthCodeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ResetOauth2JwtBearerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetOauth2JwtBearerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

