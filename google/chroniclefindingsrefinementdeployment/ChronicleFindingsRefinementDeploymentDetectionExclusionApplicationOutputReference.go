// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclefindingsrefinementdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/chroniclefindingsrefinementdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference interface {
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
	CuratedRules() *[]*string
	SetCuratedRules(val *[]*string)
	CuratedRuleSets() *[]*string
	SetCuratedRuleSets(val *[]*string)
	CuratedRuleSetsInput() *[]*string
	CuratedRulesInput() *[]*string
	DeletedCuratedRuleSets() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication
	SetInternalValue(val *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication)
	Rules() *[]*string
	SetRules(val *[]*string)
	RulesInput() *[]*string
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
	ResetCuratedRules()
	ResetCuratedRuleSets()
	ResetRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference
type jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRuleSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRuleSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRuleSetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRuleSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) CuratedRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"curatedRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) DeletedCuratedRuleSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deletedCuratedRuleSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InternalValue() *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication {
	var returns *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Rules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) RulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFindingsRefinementDeployment.ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference_Override(c ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleFindingsRefinementDeployment.ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetCuratedRules(val *[]*string) {
	if err := j.validateSetCuratedRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curatedRules",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetCuratedRuleSets(val *[]*string) {
	if err := j.validateSetCuratedRuleSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curatedRuleSets",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetInternalValue(val *ChronicleFindingsRefinementDeploymentDetectionExclusionApplication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetRules(val *[]*string) {
	if err := j.validateSetRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetCuratedRules() {
	_jsii_.InvokeVoid(
		c,
		"resetCuratedRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetCuratedRuleSets() {
	_jsii_.InvokeVoid(
		c,
		"resetCuratedRuleSets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		c,
		"resetRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChronicleFindingsRefinementDeploymentDetectionExclusionApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

