// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesapp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference interface {
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
	ExtraToolCallBehavior() *string
	SetExtraToolCallBehavior(val *string)
	ExtraToolCallBehaviorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings
	SetInternalValue(val *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings)
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
	ResetExtraToolCallBehavior()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference
type jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ExtraToolCallBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraToolCallBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ExtraToolCallBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraToolCallBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) InternalValue() *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings {
	var returns *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference_Override(c CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesApp.CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetExtraToolCallBehavior(val *string) {
	if err := j.validateSetExtraToolCallBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraToolCallBehavior",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetInternalValue(val *CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ResetExtraToolCallBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetExtraToolCallBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

