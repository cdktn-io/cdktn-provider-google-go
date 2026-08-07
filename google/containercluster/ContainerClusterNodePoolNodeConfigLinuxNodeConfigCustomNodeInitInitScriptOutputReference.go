// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/containercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference interface {
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
	GcpSecretManagerSecretUri() *string
	SetGcpSecretManagerSecretUri(val *string)
	GcpSecretManagerSecretUriInput() *string
	GcsGeneration() *float64
	SetGcsGeneration(val *float64)
	GcsGenerationInput() *float64
	GcsUri() *string
	SetGcsUri(val *string)
	GcsUriInput() *string
	InternalValue() *ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript
	SetInternalValue(val *ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript)
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
	ResetGcpSecretManagerSecretUri()
	ResetGcsGeneration()
	ResetGcsUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference
type jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcpSecretManagerSecretUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpSecretManagerSecretUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcpSecretManagerSecretUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpSecretManagerSecretUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsGeneration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gcsGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsGenerationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gcsGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InternalValue() *ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript {
	var returns *ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference_Override(c ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcpSecretManagerSecretUri(val *string) {
	if err := j.validateSetGcpSecretManagerSecretUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpSecretManagerSecretUri",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcsGeneration(val *float64) {
	if err := j.validateSetGcsGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsGeneration",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetGcsUri(val *string) {
	if err := j.validateSetGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsUri",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetInternalValue(val *ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcpSecretManagerSecretUri() {
	_jsii_.InvokeVoid(
		c,
		"resetGcpSecretManagerSecretUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcsGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ResetGcsUri() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScriptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

