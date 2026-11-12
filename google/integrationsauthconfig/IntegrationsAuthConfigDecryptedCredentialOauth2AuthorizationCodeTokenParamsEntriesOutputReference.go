// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/integrationsauthconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKeyOutputReference
	KeyInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference
	ValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue
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
	PutKey(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKey)
	PutValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue)
	ResetKey()
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

// The jsii proxy struct for IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference
type jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) Key() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKeyOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKeyOutputReference
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) KeyInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKey {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKey
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) Value() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ValueInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference_Override(i IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) PutKey(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesKey) {
	if err := i.validatePutKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKey",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) PutValue(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesValue) {
	if err := i.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		i,
		"resetKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		i,
		"resetValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

