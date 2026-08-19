// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference interface {
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
	FileExtensionRestrictions() *[]*string
	SetFileExtensionRestrictions(val *[]*string)
	FileExtensionRestrictionsInput() *[]*string
	// Experimental.
	Fqn() *string
	GoogleCloudStorageObject() *string
	SetGoogleCloudStorageObject(val *string)
	GoogleCloudStorageObjectInput() *string
	InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource
	SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource)
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
	ResetFileExtensionRestrictions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) FileExtensionRestrictions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileExtensionRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) FileExtensionRestrictionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileExtensionRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GoogleCloudStorageObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCloudStorageObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GoogleCloudStorageObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCloudStorageObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetFileExtensionRestrictions(val *[]*string) {
	if err := j.validateSetFileExtensionRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtensionRestrictions",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetGoogleCloudStorageObject(val *string) {
	if err := j.validateSetGoogleCloudStorageObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleCloudStorageObject",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) ResetFileExtensionRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetFileExtensionRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

