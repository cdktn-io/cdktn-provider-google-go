// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataflowflextemplatejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/dataflowflextemplatejob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job google_dataflow_flex_template_job}.
type DataflowFlexTemplateJob interface {
	cdktn.TerraformResource
	AdditionalExperiments() *[]*string
	SetAdditionalExperiments(val *[]*string)
	AdditionalExperimentsInput() *[]*string
	AdditionalPipelineOptions() *[]*string
	SetAdditionalPipelineOptions(val *[]*string)
	AdditionalPipelineOptionsInput() *[]*string
	AutoscalingAlgorithm() *string
	SetAutoscalingAlgorithm(val *string)
	AutoscalingAlgorithmInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerSpecGcsPath() *string
	SetContainerSpecGcsPath(val *string)
	ContainerSpecGcsPathInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateIgnoreAlreadyExists() interface{}
	SetCreateIgnoreAlreadyExists(val interface{})
	CreateIgnoreAlreadyExistsInput() interface{}
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktn.StringMap
	EnableStreamingEngine() interface{}
	SetEnableStreamingEngine(val interface{})
	EnableStreamingEngineInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpConfiguration() *string
	SetIpConfiguration(val *string)
	IpConfigurationInput() *string
	JobId() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LauncherMachineType() *string
	SetLauncherMachineType(val *string)
	LauncherMachineTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MaxWorkersInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	NumWorkers() *float64
	SetNumWorkers(val *float64)
	NumWorkersInput() *float64
	OnDelete() *string
	SetOnDelete(val *string)
	OnDeleteInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SdkContainerImage() *string
	SetSdkContainerImage(val *string)
	SdkContainerImageInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SkipWaitOnJobTermination() interface{}
	SetSkipWaitOnJobTermination(val interface{})
	SkipWaitOnJobTerminationInput() interface{}
	StagingLocation() *string
	SetStagingLocation(val *string)
	StagingLocationInput() *string
	State() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	TempLocation() *string
	SetTempLocation(val *string)
	TempLocationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransformNameMapping() *map[string]*string
	SetTransformNameMapping(val *map[string]*string)
	TransformNameMappingInput() *map[string]*string
	Type() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAdditionalExperiments()
	ResetAdditionalPipelineOptions()
	ResetAutoscalingAlgorithm()
	ResetCreateIgnoreAlreadyExists()
	ResetDeletionPolicy()
	ResetEnableStreamingEngine()
	ResetId()
	ResetIpConfiguration()
	ResetKmsKeyName()
	ResetLabels()
	ResetLauncherMachineType()
	ResetMachineType()
	ResetMaxWorkers()
	ResetNetwork()
	ResetNumWorkers()
	ResetOnDelete()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetProject()
	ResetRegion()
	ResetSdkContainerImage()
	ResetServiceAccountEmail()
	ResetSkipWaitOnJobTermination()
	ResetStagingLocation()
	ResetSubnetwork()
	ResetTempLocation()
	ResetTransformNameMapping()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataflowFlexTemplateJob
type jsiiProxy_DataflowFlexTemplateJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AdditionalExperiments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AdditionalExperimentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalExperimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AdditionalPipelineOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPipelineOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AdditionalPipelineOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalPipelineOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AutoscalingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) AutoscalingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ContainerSpecGcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ContainerSpecGcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) CreateIgnoreAlreadyExists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIgnoreAlreadyExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) CreateIgnoreAlreadyExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createIgnoreAlreadyExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) EnableStreamingEngine() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) EnableStreamingEngineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStreamingEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) IpConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) IpConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) LauncherMachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launcherMachineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) LauncherMachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launcherMachineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) NumWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) NumWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) OnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) OnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) SdkContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkContainerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) SdkContainerImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sdkContainerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) SkipWaitOnJobTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) SkipWaitOnJobTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipWaitOnJobTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) StagingLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) StagingLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TempLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TempLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TransformNameMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) TransformNameMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataflowFlexTemplateJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job google_dataflow_flex_template_job} Resource.
func NewDataflowFlexTemplateJob(scope constructs.Construct, id *string, config *DataflowFlexTemplateJobConfig) DataflowFlexTemplateJob {
	_init_.Initialize()

	if err := validateNewDataflowFlexTemplateJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataflowFlexTemplateJob{}

	_jsii_.Create(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/dataflow_flex_template_job google_dataflow_flex_template_job} Resource.
func NewDataflowFlexTemplateJob_Override(d DataflowFlexTemplateJob, scope constructs.Construct, id *string, config *DataflowFlexTemplateJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetAdditionalExperiments(val *[]*string) {
	if err := j.validateSetAdditionalExperimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalExperiments",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetAdditionalPipelineOptions(val *[]*string) {
	if err := j.validateSetAdditionalPipelineOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalPipelineOptions",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetAutoscalingAlgorithm(val *string) {
	if err := j.validateSetAutoscalingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetContainerSpecGcsPath(val *string) {
	if err := j.validateSetContainerSpecGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerSpecGcsPath",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetCreateIgnoreAlreadyExists(val interface{}) {
	if err := j.validateSetCreateIgnoreAlreadyExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createIgnoreAlreadyExists",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetEnableStreamingEngine(val interface{}) {
	if err := j.validateSetEnableStreamingEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStreamingEngine",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetIpConfiguration(val *string) {
	if err := j.validateSetIpConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipConfiguration",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetLauncherMachineType(val *string) {
	if err := j.validateSetLauncherMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launcherMachineType",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetMaxWorkers(val *float64) {
	if err := j.validateSetMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetNumWorkers(val *float64) {
	if err := j.validateSetNumWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numWorkers",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetOnDelete(val *string) {
	if err := j.validateSetOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDelete",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetSdkContainerImage(val *string) {
	if err := j.validateSetSdkContainerImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sdkContainerImage",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetSkipWaitOnJobTermination(val interface{}) {
	if err := j.validateSetSkipWaitOnJobTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipWaitOnJobTermination",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetStagingLocation(val *string) {
	if err := j.validateSetStagingLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingLocation",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetTempLocation(val *string) {
	if err := j.validateSetTempLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempLocation",
		val,
	)
}

func (j *jsiiProxy_DataflowFlexTemplateJob)SetTransformNameMapping(val *map[string]*string) {
	if err := j.validateSetTransformNameMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformNameMapping",
		val,
	)
}

// Generates CDKTN code for importing a DataflowFlexTemplateJob resource upon running "cdktn plan <stack-name>".
func DataflowFlexTemplateJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataflowFlexTemplateJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataflowFlexTemplateJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowFlexTemplateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataflowFlexTemplateJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowFlexTemplateJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataflowFlexTemplateJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataflowFlexTemplateJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataflowFlexTemplateJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dataflowFlexTemplateJob.DataflowFlexTemplateJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetAdditionalExperiments() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalExperiments",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetAdditionalPipelineOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalPipelineOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetAutoscalingAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetCreateIgnoreAlreadyExists() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateIgnoreAlreadyExists",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetEnableStreamingEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStreamingEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetIpConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetIpConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetLauncherMachineType() {
	_jsii_.InvokeVoid(
		d,
		"resetLauncherMachineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetMachineType() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetNumWorkers() {
	_jsii_.InvokeVoid(
		d,
		"resetNumWorkers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetOnDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetSdkContainerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetSdkContainerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetSkipWaitOnJobTermination() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipWaitOnJobTermination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetStagingLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetStagingLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetTempLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetTempLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ResetTransformNameMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformNameMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataflowFlexTemplateJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataflowFlexTemplateJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

