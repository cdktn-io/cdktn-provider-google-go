// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/networksecuritysecurityprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList
type jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList {
	_init_.Initialize()

	if err := validateNewNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList{}

	_jsii_.Create(
		"@cdktn/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList_Override(n NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := n.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		n,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) Get(index *float64) NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineCloudAnalysisRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

