// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chronicleparser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleParserLowCodeFieldExtractorsOutputReference interface {
	cdktn.ComplexObject
	AppendRepeatedFields() interface{}
	SetAppendRepeatedFields(val interface{})
	AppendRepeatedFieldsInput() interface{}
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
	Extractors() ChronicleParserLowCodeFieldExtractorsExtractorsList
	ExtractorsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleParserLowCodeFieldExtractors
	SetInternalValue(val *ChronicleParserLowCodeFieldExtractors)
	LogFormat() *string
	SetLogFormat(val *string)
	LogFormatInput() *string
	PreprocessConfig() ChronicleParserLowCodeFieldExtractorsPreprocessConfigOutputReference
	PreprocessConfigInput() *ChronicleParserLowCodeFieldExtractorsPreprocessConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransformedCbnSnippet() *string
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
	PutExtractors(value interface{})
	PutPreprocessConfig(value *ChronicleParserLowCodeFieldExtractorsPreprocessConfig)
	ResetAppendRepeatedFields()
	ResetExtractors()
	ResetLogFormat()
	ResetPreprocessConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleParserLowCodeFieldExtractorsOutputReference
type jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) AppendRepeatedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) AppendRepeatedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) Extractors() ChronicleParserLowCodeFieldExtractorsExtractorsList {
	var returns ChronicleParserLowCodeFieldExtractorsExtractorsList
	_jsii_.Get(
		j,
		"extractors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ExtractorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) InternalValue() *ChronicleParserLowCodeFieldExtractors {
	var returns *ChronicleParserLowCodeFieldExtractors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) LogFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) PreprocessConfig() ChronicleParserLowCodeFieldExtractorsPreprocessConfigOutputReference {
	var returns ChronicleParserLowCodeFieldExtractorsPreprocessConfigOutputReference
	_jsii_.Get(
		j,
		"preprocessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) PreprocessConfigInput() *ChronicleParserLowCodeFieldExtractorsPreprocessConfig {
	var returns *ChronicleParserLowCodeFieldExtractorsPreprocessConfig
	_jsii_.Get(
		j,
		"preprocessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) TransformedCbnSnippet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformedCbnSnippet",
		&returns,
	)
	return returns
}


func NewChronicleParserLowCodeFieldExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleParserLowCodeFieldExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleParserLowCodeFieldExtractorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParser.ChronicleParserLowCodeFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleParserLowCodeFieldExtractorsOutputReference_Override(c ChronicleParserLowCodeFieldExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParser.ChronicleParserLowCodeFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetAppendRepeatedFields(val interface{}) {
	if err := j.validateSetAppendRepeatedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendRepeatedFields",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetInternalValue(val *ChronicleParserLowCodeFieldExtractors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetLogFormat(val *string) {
	if err := j.validateSetLogFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) PutExtractors(value interface{}) {
	if err := c.validatePutExtractorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExtractors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) PutPreprocessConfig(value *ChronicleParserLowCodeFieldExtractorsPreprocessConfig) {
	if err := c.validatePutPreprocessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreprocessConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ResetAppendRepeatedFields() {
	_jsii_.InvokeVoid(
		c,
		"resetAppendRepeatedFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ResetExtractors() {
	_jsii_.InvokeVoid(
		c,
		"resetExtractors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ResetLogFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetLogFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ResetPreprocessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPreprocessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleParserLowCodeFieldExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

