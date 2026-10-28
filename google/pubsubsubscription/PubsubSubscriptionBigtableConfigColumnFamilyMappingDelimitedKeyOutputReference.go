// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/pubsubsubscription/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference interface {
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
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey
	SetInternalValue(val *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey)
	KeyFields() *[]*string
	SetKeyFields(val *[]*string)
	KeyFieldsInput() *[]*string
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
	ResetDelimiter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference
type jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) InternalValue() *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey {
	var returns *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) KeyFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) KeyFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.pubsubSubscription.PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference_Override(p PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.pubsubSubscription.PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetInternalValue(val *PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetKeyFields(val *[]*string) {
	if err := j.validateSetKeyFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFields",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		p,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigtableConfigColumnFamilyMappingDelimitedKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

