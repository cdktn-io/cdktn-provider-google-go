// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/dialogflowtool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DialogflowToolOpenApiSpecOutputReference interface {
	cdktn.ComplexObject
	Authentication() DialogflowToolOpenApiSpecAuthenticationOutputReference
	AuthenticationInput() *DialogflowToolOpenApiSpecAuthentication
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
	InternalValue() *DialogflowToolOpenApiSpec
	SetInternalValue(val *DialogflowToolOpenApiSpec)
	ServiceDirectoryConfig() DialogflowToolOpenApiSpecServiceDirectoryConfigOutputReference
	ServiceDirectoryConfigInput() *DialogflowToolOpenApiSpecServiceDirectoryConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextSchema() *string
	SetTextSchema(val *string)
	TextSchemaInput() *string
	TlsConfig() DialogflowToolOpenApiSpecTlsConfigOutputReference
	TlsConfigInput() *DialogflowToolOpenApiSpecTlsConfig
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
	PutAuthentication(value *DialogflowToolOpenApiSpecAuthentication)
	PutServiceDirectoryConfig(value *DialogflowToolOpenApiSpecServiceDirectoryConfig)
	PutTlsConfig(value *DialogflowToolOpenApiSpecTlsConfig)
	ResetAuthentication()
	ResetServiceDirectoryConfig()
	ResetTlsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowToolOpenApiSpecOutputReference
type jsiiProxy_DialogflowToolOpenApiSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) Authentication() DialogflowToolOpenApiSpecAuthenticationOutputReference {
	var returns DialogflowToolOpenApiSpecAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) AuthenticationInput() *DialogflowToolOpenApiSpecAuthentication {
	var returns *DialogflowToolOpenApiSpecAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) InternalValue() *DialogflowToolOpenApiSpec {
	var returns *DialogflowToolOpenApiSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ServiceDirectoryConfig() DialogflowToolOpenApiSpecServiceDirectoryConfigOutputReference {
	var returns DialogflowToolOpenApiSpecServiceDirectoryConfigOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ServiceDirectoryConfigInput() *DialogflowToolOpenApiSpecServiceDirectoryConfig {
	var returns *DialogflowToolOpenApiSpecServiceDirectoryConfig
	_jsii_.Get(
		j,
		"serviceDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TextSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TextSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TlsConfig() DialogflowToolOpenApiSpecTlsConfigOutputReference {
	var returns DialogflowToolOpenApiSpecTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) TlsConfigInput() *DialogflowToolOpenApiSpecTlsConfig {
	var returns *DialogflowToolOpenApiSpecTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}


func NewDialogflowToolOpenApiSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowToolOpenApiSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowToolOpenApiSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowToolOpenApiSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolOpenApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowToolOpenApiSpecOutputReference_Override(d DialogflowToolOpenApiSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowTool.DialogflowToolOpenApiSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetInternalValue(val *DialogflowToolOpenApiSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowToolOpenApiSpecOutputReference)SetTextSchema(val *string) {
	if err := j.validateSetTextSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textSchema",
		val,
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) PutAuthentication(value *DialogflowToolOpenApiSpecAuthentication) {
	if err := d.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) PutServiceDirectoryConfig(value *DialogflowToolOpenApiSpecServiceDirectoryConfig) {
	if err := d.validatePutServiceDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceDirectoryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) PutTlsConfig(value *DialogflowToolOpenApiSpecTlsConfig) {
	if err := d.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ResetServiceDirectoryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceDirectoryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowToolOpenApiSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

