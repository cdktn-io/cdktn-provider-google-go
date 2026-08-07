// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/cestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesToolWidgetToolOutputReference interface {
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
	DataMapping() CesToolWidgetToolDataMappingOutputReference
	DataMappingInput() *CesToolWidgetToolDataMapping
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolWidgetTool
	SetInternalValue(val *CesToolWidgetTool)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() CesToolWidgetToolParametersOutputReference
	ParametersInput() *CesToolWidgetToolParameters
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextResponseConfig() CesToolWidgetToolTextResponseConfigOutputReference
	TextResponseConfigInput() *CesToolWidgetToolTextResponseConfig
	UiConfig() *string
	SetUiConfig(val *string)
	UiConfigInput() *string
	WidgetType() *string
	SetWidgetType(val *string)
	WidgetTypeInput() *string
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
	PutDataMapping(value *CesToolWidgetToolDataMapping)
	PutParameters(value *CesToolWidgetToolParameters)
	PutTextResponseConfig(value *CesToolWidgetToolTextResponseConfig)
	ResetDataMapping()
	ResetDescription()
	ResetParameters()
	ResetTextResponseConfig()
	ResetUiConfig()
	ResetWidgetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolWidgetToolOutputReference
type jsiiProxy_CesToolWidgetToolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) DataMapping() CesToolWidgetToolDataMappingOutputReference {
	var returns CesToolWidgetToolDataMappingOutputReference
	_jsii_.Get(
		j,
		"dataMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) DataMappingInput() *CesToolWidgetToolDataMapping {
	var returns *CesToolWidgetToolDataMapping
	_jsii_.Get(
		j,
		"dataMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) InternalValue() *CesToolWidgetTool {
	var returns *CesToolWidgetTool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) Parameters() CesToolWidgetToolParametersOutputReference {
	var returns CesToolWidgetToolParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) ParametersInput() *CesToolWidgetToolParameters {
	var returns *CesToolWidgetToolParameters
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) TextResponseConfig() CesToolWidgetToolTextResponseConfigOutputReference {
	var returns CesToolWidgetToolTextResponseConfigOutputReference
	_jsii_.Get(
		j,
		"textResponseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) TextResponseConfigInput() *CesToolWidgetToolTextResponseConfig {
	var returns *CesToolWidgetToolTextResponseConfig
	_jsii_.Get(
		j,
		"textResponseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) UiConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) UiConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) WidgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference) WidgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetTypeInput",
		&returns,
	)
	return returns
}


func NewCesToolWidgetToolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesToolWidgetToolOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolWidgetToolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolWidgetToolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolWidgetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolWidgetToolOutputReference_Override(c CesToolWidgetToolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesTool.CesToolWidgetToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetInternalValue(val *CesToolWidgetTool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetUiConfig(val *string) {
	if err := j.validateSetUiConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uiConfig",
		val,
	)
}

func (j *jsiiProxy_CesToolWidgetToolOutputReference)SetWidgetType(val *string) {
	if err := j.validateSetWidgetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetType",
		val,
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) PutDataMapping(value *CesToolWidgetToolDataMapping) {
	if err := c.validatePutDataMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataMapping",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) PutParameters(value *CesToolWidgetToolParameters) {
	if err := c.validatePutParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParameters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) PutTextResponseConfig(value *CesToolWidgetToolTextResponseConfig) {
	if err := c.validatePutTextResponseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTextResponseConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetDataMapping() {
	_jsii_.InvokeVoid(
		c,
		"resetDataMapping",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetTextResponseConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTextResponseConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetUiConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetUiConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ResetWidgetType() {
	_jsii_.InvokeVoid(
		c,
		"resetWidgetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolWidgetToolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesToolWidgetToolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

