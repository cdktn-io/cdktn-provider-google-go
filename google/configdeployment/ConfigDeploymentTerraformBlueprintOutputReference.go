// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/configdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigDeploymentTerraformBlueprintOutputReference interface {
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
	GcsSource() *string
	SetGcsSource(val *string)
	GcsSourceInput() *string
	GitSource() ConfigDeploymentTerraformBlueprintGitSourceOutputReference
	GitSourceInput() *ConfigDeploymentTerraformBlueprintGitSource
	InputValues() ConfigDeploymentTerraformBlueprintInputValuesList
	InputValuesInput() interface{}
	InternalValue() *ConfigDeploymentTerraformBlueprint
	SetInternalValue(val *ConfigDeploymentTerraformBlueprint)
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
	PutGitSource(value *ConfigDeploymentTerraformBlueprintGitSource)
	PutInputValues(value interface{})
	ResetGcsSource()
	ResetGitSource()
	ResetInputValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigDeploymentTerraformBlueprintOutputReference
type jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GcsSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GcsSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GitSource() ConfigDeploymentTerraformBlueprintGitSourceOutputReference {
	var returns ConfigDeploymentTerraformBlueprintGitSourceOutputReference
	_jsii_.Get(
		j,
		"gitSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GitSourceInput() *ConfigDeploymentTerraformBlueprintGitSource {
	var returns *ConfigDeploymentTerraformBlueprintGitSource
	_jsii_.Get(
		j,
		"gitSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) InputValues() ConfigDeploymentTerraformBlueprintInputValuesList {
	var returns ConfigDeploymentTerraformBlueprintInputValuesList
	_jsii_.Get(
		j,
		"inputValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) InputValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) InternalValue() *ConfigDeploymentTerraformBlueprint {
	var returns *ConfigDeploymentTerraformBlueprint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigDeploymentTerraformBlueprintOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConfigDeploymentTerraformBlueprintOutputReference {
	_init_.Initialize()

	if err := validateNewConfigDeploymentTerraformBlueprintOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.configDeployment.ConfigDeploymentTerraformBlueprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigDeploymentTerraformBlueprintOutputReference_Override(c ConfigDeploymentTerraformBlueprintOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.configDeployment.ConfigDeploymentTerraformBlueprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetGcsSource(val *string) {
	if err := j.validateSetGcsSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsSource",
		val,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetInternalValue(val *ConfigDeploymentTerraformBlueprint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) PutGitSource(value *ConfigDeploymentTerraformBlueprintGitSource) {
	if err := c.validatePutGitSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) PutInputValues(value interface{}) {
	if err := c.validatePutInputValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInputValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ResetGcsSource() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ResetGitSource() {
	_jsii_.InvokeVoid(
		c,
		"resetGitSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ResetInputValues() {
	_jsii_.InvokeVoid(
		c,
		"resetInputValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConfigDeploymentTerraformBlueprintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

