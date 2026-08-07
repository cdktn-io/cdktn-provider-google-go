// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chronicleparser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference interface {
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
	DestinationPath() *string
	SetDestinationPath(val *string)
	DestinationPathInput() *string
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PreconditionOp() *string
	SetPreconditionOp(val *string)
	PreconditionOpInput() *string
	PreconditionPath() *string
	SetPreconditionPath(val *string)
	PreconditionPathInput() *string
	PreconditionValue() *string
	SetPreconditionValue(val *string)
	PreconditionValueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	ResetDestinationPath()
	ResetFieldPath()
	ResetPreconditionOp()
	ResetPreconditionPath()
	ResetPreconditionValue()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference
type jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionOp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionOpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionOpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) PreconditionValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preconditionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewChronicleParserLowCodeFieldExtractorsExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleParserLowCodeFieldExtractorsExtractorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParser.ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChronicleParserLowCodeFieldExtractorsExtractorsOutputReference_Override(c ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParser.ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionOp(val *string) {
	if err := j.validateSetPreconditionOpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionOp",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionPath(val *string) {
	if err := j.validateSetPreconditionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionPath",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetPreconditionValue(val *string) {
	if err := j.validateSetPreconditionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preconditionValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionOp() {
	_jsii_.InvokeVoid(
		c,
		"resetPreconditionOp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPreconditionPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetPreconditionValue() {
	_jsii_.InvokeVoid(
		c,
		"resetPreconditionValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

