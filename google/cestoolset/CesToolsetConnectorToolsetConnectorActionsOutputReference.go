// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/cestoolset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CesToolsetConnectorToolsetConnectorActionsOutputReference interface {
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
	ConnectionActionId() *string
	SetConnectionActionId(val *string)
	ConnectionActionIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EntityOperation() CesToolsetConnectorToolsetConnectorActionsEntityOperationOutputReference
	EntityOperationInput() *CesToolsetConnectorToolsetConnectorActionsEntityOperation
	// Experimental.
	Fqn() *string
	InputFields() *[]*string
	SetInputFields(val *[]*string)
	InputFieldsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputFields() *[]*string
	SetOutputFields(val *[]*string)
	OutputFieldsInput() *[]*string
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
	PutEntityOperation(value *CesToolsetConnectorToolsetConnectorActionsEntityOperation)
	ResetConnectionActionId()
	ResetEntityOperation()
	ResetInputFields()
	ResetOutputFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolsetConnectorToolsetConnectorActionsOutputReference
type jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ConnectionActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ConnectionActionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) EntityOperation() CesToolsetConnectorToolsetConnectorActionsEntityOperationOutputReference {
	var returns CesToolsetConnectorToolsetConnectorActionsEntityOperationOutputReference
	_jsii_.Get(
		j,
		"entityOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) EntityOperationInput() *CesToolsetConnectorToolsetConnectorActionsEntityOperation {
	var returns *CesToolsetConnectorToolsetConnectorActionsEntityOperation
	_jsii_.Get(
		j,
		"entityOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) InputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) InputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) OutputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) OutputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolsetConnectorToolsetConnectorActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesToolsetConnectorToolsetConnectorActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolsetConnectorToolsetConnectorActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetConnectorToolsetConnectorActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesToolsetConnectorToolsetConnectorActionsOutputReference_Override(c CesToolsetConnectorToolsetConnectorActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cesToolset.CesToolsetConnectorToolsetConnectorActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetConnectionActionId(val *string) {
	if err := j.validateSetConnectionActionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionActionId",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetInputFields(val *[]*string) {
	if err := j.validateSetInputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFields",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetOutputFields(val *[]*string) {
	if err := j.validateSetOutputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFields",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) PutEntityOperation(value *CesToolsetConnectorToolsetConnectorActionsEntityOperation) {
	if err := c.validatePutEntityOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEntityOperation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ResetConnectionActionId() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectionActionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ResetEntityOperation() {
	_jsii_.InvokeVoid(
		c,
		"resetEntityOperation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ResetInputFields() {
	_jsii_.InvokeVoid(
		c,
		"resetInputFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ResetOutputFields() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CesToolsetConnectorToolsetConnectorActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

