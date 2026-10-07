// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/cesagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesAgentTransferRulesDeterministicTransferOutputReference interface {
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
	ExpressionCondition() CesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference
	ExpressionConditionInput() *CesAgentTransferRulesDeterministicTransferExpressionCondition
	// Experimental.
	Fqn() *string
	InternalValue() *CesAgentTransferRulesDeterministicTransfer
	SetInternalValue(val *CesAgentTransferRulesDeterministicTransfer)
	PythonCodeCondition() CesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference
	PythonCodeConditionInput() *CesAgentTransferRulesDeterministicTransferPythonCodeCondition
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
	PutExpressionCondition(value *CesAgentTransferRulesDeterministicTransferExpressionCondition)
	PutPythonCodeCondition(value *CesAgentTransferRulesDeterministicTransferPythonCodeCondition)
	ResetExpressionCondition()
	ResetPythonCodeCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAgentTransferRulesDeterministicTransferOutputReference
type jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ExpressionCondition() CesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference {
	var returns CesAgentTransferRulesDeterministicTransferExpressionConditionOutputReference
	_jsii_.Get(
		j,
		"expressionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ExpressionConditionInput() *CesAgentTransferRulesDeterministicTransferExpressionCondition {
	var returns *CesAgentTransferRulesDeterministicTransferExpressionCondition
	_jsii_.Get(
		j,
		"expressionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) InternalValue() *CesAgentTransferRulesDeterministicTransfer {
	var returns *CesAgentTransferRulesDeterministicTransfer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) PythonCodeCondition() CesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference {
	var returns CesAgentTransferRulesDeterministicTransferPythonCodeConditionOutputReference
	_jsii_.Get(
		j,
		"pythonCodeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) PythonCodeConditionInput() *CesAgentTransferRulesDeterministicTransferPythonCodeCondition {
	var returns *CesAgentTransferRulesDeterministicTransferPythonCodeCondition
	_jsii_.Get(
		j,
		"pythonCodeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesAgentTransferRulesDeterministicTransferOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CesAgentTransferRulesDeterministicTransferOutputReference {
	_init_.Initialize()

	if err := validateNewCesAgentTransferRulesDeterministicTransferOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentTransferRulesDeterministicTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesAgentTransferRulesDeterministicTransferOutputReference_Override(c CesAgentTransferRulesDeterministicTransferOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesAgent.CesAgentTransferRulesDeterministicTransferOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference)SetInternalValue(val *CesAgentTransferRulesDeterministicTransfer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) PutExpressionCondition(value *CesAgentTransferRulesDeterministicTransferExpressionCondition) {
	if err := c.validatePutExpressionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpressionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) PutPythonCodeCondition(value *CesAgentTransferRulesDeterministicTransferPythonCodeCondition) {
	if err := c.validatePutPythonCodeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPythonCodeCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ResetExpressionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetExpressionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ResetPythonCodeCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonCodeCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesAgentTransferRulesDeterministicTransferOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

