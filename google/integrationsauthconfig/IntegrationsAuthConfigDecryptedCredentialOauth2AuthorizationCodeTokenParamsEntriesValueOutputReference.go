// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/integrationsauthconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference interface {
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
	InternalValue() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue
	SetInternalValue(val *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue)
	LiteralValue() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValueOutputReference
	LiteralValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValue
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
	PutLiteralValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValue)
	ResetLiteralValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference
type jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) InternalValue() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) LiteralValue() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValueOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValueOutputReference
	_jsii_.Get(
		j,
		"literalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) LiteralValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValue {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValue
	_jsii_.Get(
		j,
		"literalValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference_Override(i IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference)SetInternalValue(val *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) PutLiteralValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueLiteralValue) {
	if err := i.validatePutLiteralValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLiteralValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) ResetLiteralValue() {
	_jsii_.InvokeVoid(
		i,
		"resetLiteralValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

