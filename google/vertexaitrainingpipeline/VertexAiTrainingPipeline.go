// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v22/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v22/vertexaitrainingpipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline google_vertex_ai_training_pipeline}.
type VertexAiTrainingPipeline interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktn.StringMap
	EncryptionSpec() VertexAiTrainingPipelineEncryptionSpecOutputReference
	EncryptionSpecInput() *VertexAiTrainingPipelineEncryptionSpec
	EndTime() *string
	Error() VertexAiTrainingPipelineErrorList
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
	InputDataConfig() VertexAiTrainingPipelineInputDataConfigOutputReference
	InputDataConfigInput() *VertexAiTrainingPipelineInputDataConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ModelId() *string
	SetModelId(val *string)
	ModelIdInput() *string
	ModelToUpload() VertexAiTrainingPipelineModelToUploadOutputReference
	ModelToUploadInput() *VertexAiTrainingPipelineModelToUpload
	Name() *string
	// The tree node.
	Node() constructs.Node
	ParentModel() *string
	SetParentModel(val *string)
	ParentModelInput() *string
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
	StartTime() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VertexAiTrainingPipelineTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrainingTaskDefinition() *string
	SetTrainingTaskDefinition(val *string)
	TrainingTaskDefinitionInput() *string
	TrainingTaskInputs() *string
	SetTrainingTaskInputs(val *string)
	TrainingTaskInputsInput() *string
	TrainingTaskMetadata() *string
	UpdateTime() *string
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
	PutEncryptionSpec(value *VertexAiTrainingPipelineEncryptionSpec)
	PutInputDataConfig(value *VertexAiTrainingPipelineInputDataConfig)
	PutModelToUpload(value *VertexAiTrainingPipelineModelToUpload)
	PutTimeouts(value *VertexAiTrainingPipelineTimeouts)
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
	ResetDeletionPolicy()
	ResetEncryptionSpec()
	ResetId()
	ResetInputDataConfig()
	ResetLabels()
	ResetModelToUpload()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentModel()
	ResetProject()
	ResetRegion()
	ResetTimeouts()
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

// The jsii proxy struct for VertexAiTrainingPipeline
type jsiiProxy_VertexAiTrainingPipeline struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VertexAiTrainingPipeline) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) EncryptionSpec() VertexAiTrainingPipelineEncryptionSpecOutputReference {
	var returns VertexAiTrainingPipelineEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) EncryptionSpecInput() *VertexAiTrainingPipelineEncryptionSpec {
	var returns *VertexAiTrainingPipelineEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Error() VertexAiTrainingPipelineErrorList {
	var returns VertexAiTrainingPipelineErrorList
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) InputDataConfig() VertexAiTrainingPipelineInputDataConfigOutputReference {
	var returns VertexAiTrainingPipelineInputDataConfigOutputReference
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) InputDataConfigInput() *VertexAiTrainingPipelineInputDataConfig {
	var returns *VertexAiTrainingPipelineInputDataConfig
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ModelToUpload() VertexAiTrainingPipelineModelToUploadOutputReference {
	var returns VertexAiTrainingPipelineModelToUploadOutputReference
	_jsii_.Get(
		j,
		"modelToUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ModelToUploadInput() *VertexAiTrainingPipelineModelToUpload {
	var returns *VertexAiTrainingPipelineModelToUpload
	_jsii_.Get(
		j,
		"modelToUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ParentModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ParentModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) Timeouts() VertexAiTrainingPipelineTimeoutsOutputReference {
	var returns VertexAiTrainingPipelineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TrainingTaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingTaskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TrainingTaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingTaskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TrainingTaskInputs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingTaskInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TrainingTaskInputsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingTaskInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) TrainingTaskMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingTaskMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiTrainingPipeline) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline google_vertex_ai_training_pipeline} Resource.
func NewVertexAiTrainingPipeline(scope constructs.Construct, id *string, config *VertexAiTrainingPipelineConfig) VertexAiTrainingPipeline {
	_init_.Initialize()

	if err := validateNewVertexAiTrainingPipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiTrainingPipeline{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline google_vertex_ai_training_pipeline} Resource.
func NewVertexAiTrainingPipeline_Override(v VertexAiTrainingPipeline, scope constructs.Construct, id *string, config *VertexAiTrainingPipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetParentModel(val *string) {
	if err := j.validateSetParentModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentModel",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetTrainingTaskDefinition(val *string) {
	if err := j.validateSetTrainingTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingTaskDefinition",
		val,
	)
}

func (j *jsiiProxy_VertexAiTrainingPipeline)SetTrainingTaskInputs(val *string) {
	if err := j.validateSetTrainingTaskInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingTaskInputs",
		val,
	)
}

// Generates CDKTN code for importing a VertexAiTrainingPipeline resource upon running "cdktn plan <stack-name>".
func VertexAiTrainingPipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVertexAiTrainingPipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
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
func VertexAiTrainingPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiTrainingPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiTrainingPipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiTrainingPipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VertexAiTrainingPipeline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVertexAiTrainingPipeline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VertexAiTrainingPipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.vertexAiTrainingPipeline.VertexAiTrainingPipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VertexAiTrainingPipeline) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := v.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) PutEncryptionSpec(value *VertexAiTrainingPipelineEncryptionSpec) {
	if err := v.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) PutInputDataConfig(value *VertexAiTrainingPipelineInputDataConfig) {
	if err := v.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) PutModelToUpload(value *VertexAiTrainingPipelineModelToUpload) {
	if err := v.validatePutModelToUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putModelToUpload",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) PutTimeouts(value *VertexAiTrainingPipelineTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := v.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetLabels() {
	_jsii_.InvokeVoid(
		v,
		"resetLabels",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetModelToUpload() {
	_jsii_.InvokeVoid(
		v,
		"resetModelToUpload",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetParentModel() {
	_jsii_.InvokeVoid(
		v,
		"resetParentModel",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetProject() {
	_jsii_.InvokeVoid(
		v,
		"resetProject",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetRegion() {
	_jsii_.InvokeVoid(
		v,
		"resetRegion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiTrainingPipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiTrainingPipeline) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		v,
		"with",
		args,
		&returns,
	)

	return returns
}

