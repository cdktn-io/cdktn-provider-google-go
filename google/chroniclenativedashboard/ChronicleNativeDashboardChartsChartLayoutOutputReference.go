// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclenativedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chroniclenativedashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleNativeDashboardChartsChartLayoutOutputReference interface {
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
	InternalValue() *ChronicleNativeDashboardChartsChartLayout
	SetInternalValue(val *ChronicleNativeDashboardChartsChartLayout)
	SpanX() *float64
	SetSpanX(val *float64)
	SpanXInput() *float64
	SpanY() *float64
	SetSpanY(val *float64)
	SpanYInput() *float64
	StartX() *float64
	SetStartX(val *float64)
	StartXInput() *float64
	StartY() *float64
	SetStartY(val *float64)
	StartYInput() *float64
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
	ResetStartX()
	ResetStartY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleNativeDashboardChartsChartLayoutOutputReference
type jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) InternalValue() *ChronicleNativeDashboardChartsChartLayout {
	var returns *ChronicleNativeDashboardChartsChartLayout
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) SpanX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) SpanXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) SpanY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) SpanYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spanYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) StartX() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startX",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) StartXInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startXInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) StartY() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startY",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) StartYInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startYInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleNativeDashboardChartsChartLayoutOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleNativeDashboardChartsChartLayoutOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleNativeDashboardChartsChartLayoutOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleNativeDashboard.ChronicleNativeDashboardChartsChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleNativeDashboardChartsChartLayoutOutputReference_Override(c ChronicleNativeDashboardChartsChartLayoutOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleNativeDashboard.ChronicleNativeDashboardChartsChartLayoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetInternalValue(val *ChronicleNativeDashboardChartsChartLayout) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetSpanX(val *float64) {
	if err := j.validateSetSpanXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanX",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetSpanY(val *float64) {
	if err := j.validateSetSpanYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanY",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetStartX(val *float64) {
	if err := j.validateSetStartXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startX",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetStartY(val *float64) {
	if err := j.validateSetStartYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startY",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ResetStartX() {
	_jsii_.InvokeVoid(
		c,
		"resetStartX",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ResetStartY() {
	_jsii_.InvokeVoid(
		c,
		"resetStartY",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleNativeDashboardChartsChartLayoutOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

