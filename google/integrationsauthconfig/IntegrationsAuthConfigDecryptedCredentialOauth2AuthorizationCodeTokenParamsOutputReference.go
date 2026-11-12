// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/integrationsauthconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference interface {
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
	Entries() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesList
	EntriesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams
	SetInternalValue(val *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams)
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
	PutEntries(value interface{})
	ResetEntries()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference
type jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) Entries() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesList {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsEntriesList
	_jsii_.Get(
		j,
		"entries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) EntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) InternalValue() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference_Override(i IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference)SetInternalValue(val *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) PutEntries(value interface{}) {
	if err := i.validatePutEntriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEntries",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) ResetEntries() {
	_jsii_.InvokeVoid(
		i,
		"resetEntries",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeTokenParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

