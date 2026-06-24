// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/vertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference interface {
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
	DnsPeeringConfigs() VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList
	DnsPeeringConfigsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
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
	PutDnsPeeringConfigs(value interface{})
	ResetDnsPeeringConfigs()
	ResetNetworkAttachment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
type jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) DnsPeeringConfigs() VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList {
	var returns VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigDnsPeeringConfigsList
	_jsii_.Get(
		j,
		"dnsPeeringConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) DnsPeeringConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsPeeringConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InternalValue() *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig {
	var returns *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference_Override(v VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiSchedule.VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetInternalValue(val *VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) PutDnsPeeringConfigs(value interface{}) {
	if err := v.validatePutDnsPeeringConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDnsPeeringConfigs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ResetDnsPeeringConfigs() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsPeeringConfigs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

