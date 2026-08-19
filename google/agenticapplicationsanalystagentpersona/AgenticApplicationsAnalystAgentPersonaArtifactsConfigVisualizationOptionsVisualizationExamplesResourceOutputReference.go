// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference interface {
	cdktn.ComplexObject
	BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference
	BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource
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
	DisplayLabel() *string
	SetDisplayLabel(val *string)
	DisplayLabelInput() *string
	F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference
	F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource
	// Experimental.
	Fqn() *string
	GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference
	GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource
	GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference
	GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource
	InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource
	SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource)
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference
	RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseRag() interface{}
	SetUseRag(val interface{})
	UseRagInput() interface{}
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
	PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource)
	PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource)
	PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource)
	PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource)
	PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource)
	ResetBigqueryResource()
	ResetDisplayLabel()
	ResetF1Resource()
	ResetGoogleCloudStorageResource()
	ResetGoogleDriveResource()
	ResetModelDescription()
	ResetRawFileResource()
	ResetUseRag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) BigqueryResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference
	_jsii_.Get(
		j,
		"bigqueryResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) BigqueryResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource
	_jsii_.Get(
		j,
		"bigqueryResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) DisplayLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) DisplayLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) F1Resource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference
	_jsii_.Get(
		j,
		"f1Resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) F1ResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource
	_jsii_.Get(
		j,
		"f1ResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GoogleCloudStorageResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference
	_jsii_.Get(
		j,
		"googleCloudStorageResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GoogleCloudStorageResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource
	_jsii_.Get(
		j,
		"googleCloudStorageResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GoogleDriveResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference
	_jsii_.Get(
		j,
		"googleDriveResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GoogleDriveResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource
	_jsii_.Get(
		j,
		"googleDriveResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) InternalValue() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) RawFileResource() AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference
	_jsii_.Get(
		j,
		"rawFileResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) RawFileResourceInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource
	_jsii_.Get(
		j,
		"rawFileResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) UseRag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) UseRagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRagInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetDisplayLabel(val *string) {
	if err := j.validateSetDisplayLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayLabel",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetInternalValue(val *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)SetUseRag(val interface{}) {
	if err := j.validateSetUseRagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRag",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) PutBigqueryResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource) {
	if err := a.validatePutBigqueryResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBigqueryResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) PutF1Resource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource) {
	if err := a.validatePutF1ResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putF1Resource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) PutGoogleCloudStorageResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource) {
	if err := a.validatePutGoogleCloudStorageResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleCloudStorageResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) PutGoogleDriveResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource) {
	if err := a.validatePutGoogleDriveResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleDriveResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) PutRawFileResource(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource) {
	if err := a.validatePutRawFileResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRawFileResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetBigqueryResource() {
	_jsii_.InvokeVoid(
		a,
		"resetBigqueryResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetDisplayLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetF1Resource() {
	_jsii_.InvokeVoid(
		a,
		"resetF1Resource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetGoogleCloudStorageResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleCloudStorageResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetGoogleDriveResource() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleDriveResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetRawFileResource() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFileResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ResetUseRag() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

