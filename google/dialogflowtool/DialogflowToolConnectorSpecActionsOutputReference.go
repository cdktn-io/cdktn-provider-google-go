// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/dialogflowtool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowToolConnectorSpecActionsOutputReference interface {
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
	EntityOperation() DialogflowToolConnectorSpecActionsEntityOperationOutputReference
	EntityOperationInput() *DialogflowToolConnectorSpecActionsEntityOperation
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
	PutEntityOperation(value *DialogflowToolConnectorSpecActionsEntityOperation)
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

// The jsii proxy struct for DialogflowToolConnectorSpecActionsOutputReference
type jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ConnectionActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ConnectionActionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionActionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) EntityOperation() DialogflowToolConnectorSpecActionsEntityOperationOutputReference {
	var returns DialogflowToolConnectorSpecActionsEntityOperationOutputReference
	_jsii_.Get(
		j,
		"entityOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) EntityOperationInput() *DialogflowToolConnectorSpecActionsEntityOperation {
	var returns *DialogflowToolConnectorSpecActionsEntityOperation
	_jsii_.Get(
		j,
		"entityOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) InputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) InputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) OutputFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) OutputFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowToolConnectorSpecActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowToolConnectorSpecActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowToolConnectorSpecActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolConnectorSpecActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowToolConnectorSpecActionsOutputReference_Override(d DialogflowToolConnectorSpecActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolConnectorSpecActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetConnectionActionId(val *string) {
	if err := j.validateSetConnectionActionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionActionId",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetInputFields(val *[]*string) {
	if err := j.validateSetInputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFields",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetOutputFields(val *[]*string) {
	if err := j.validateSetOutputFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFields",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) PutEntityOperation(value *DialogflowToolConnectorSpecActionsEntityOperation) {
	if err := d.validatePutEntityOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEntityOperation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ResetConnectionActionId() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionActionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ResetEntityOperation() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityOperation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ResetInputFields() {
	_jsii_.InvokeVoid(
		d,
		"resetInputFields",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ResetOutputFields() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputFields",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolConnectorSpecActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

