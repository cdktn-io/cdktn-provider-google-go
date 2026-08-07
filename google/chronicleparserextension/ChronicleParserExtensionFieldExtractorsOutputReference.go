// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chronicleparserextension/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleParserExtensionFieldExtractorsOutputReference interface {
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
	Extractors() ChronicleParserExtensionFieldExtractorsExtractorsList
	ExtractorsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleParserExtensionFieldExtractors
	SetInternalValue(val *ChronicleParserExtensionFieldExtractors)
	LogFormat() *string
	SetLogFormat(val *string)
	LogFormatInput() *string
	PreprocessConfig() ChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference
	PreprocessConfigInput() *ChronicleParserExtensionFieldExtractorsPreprocessConfig
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
	PutPreprocessConfig(value *ChronicleParserExtensionFieldExtractorsPreprocessConfig)
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

// The jsii proxy struct for ChronicleParserExtensionFieldExtractorsOutputReference
type jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) AppendRepeatedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) AppendRepeatedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendRepeatedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) Extractors() ChronicleParserExtensionFieldExtractorsExtractorsList {
	var returns ChronicleParserExtensionFieldExtractorsExtractorsList
	_jsii_.Get(
		j,
		"extractors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ExtractorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) InternalValue() *ChronicleParserExtensionFieldExtractors {
	var returns *ChronicleParserExtensionFieldExtractors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) LogFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) PreprocessConfig() ChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference {
	var returns ChronicleParserExtensionFieldExtractorsPreprocessConfigOutputReference
	_jsii_.Get(
		j,
		"preprocessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) PreprocessConfigInput() *ChronicleParserExtensionFieldExtractorsPreprocessConfig {
	var returns *ChronicleParserExtensionFieldExtractorsPreprocessConfig
	_jsii_.Get(
		j,
		"preprocessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) TransformedCbnSnippet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformedCbnSnippet",
		&returns,
	)
	return returns
}


func NewChronicleParserExtensionFieldExtractorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleParserExtensionFieldExtractorsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleParserExtensionFieldExtractorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtensionFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleParserExtensionFieldExtractorsOutputReference_Override(c ChronicleParserExtensionFieldExtractorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtensionFieldExtractorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetAppendRepeatedFields(val interface{}) {
	if err := j.validateSetAppendRepeatedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendRepeatedFields",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetInternalValue(val *ChronicleParserExtensionFieldExtractors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetLogFormat(val *string) {
	if err := j.validateSetLogFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) PutExtractors(value interface{}) {
	if err := c.validatePutExtractorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExtractors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) PutPreprocessConfig(value *ChronicleParserExtensionFieldExtractorsPreprocessConfig) {
	if err := c.validatePutPreprocessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreprocessConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ResetAppendRepeatedFields() {
	_jsii_.InvokeVoid(
		c,
		"resetAppendRepeatedFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ResetExtractors() {
	_jsii_.InvokeVoid(
		c,
		"resetExtractors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ResetLogFormat() {
	_jsii_.InvokeVoid(
		c,
		"resetLogFormat",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ResetPreprocessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPreprocessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleParserExtensionFieldExtractorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

