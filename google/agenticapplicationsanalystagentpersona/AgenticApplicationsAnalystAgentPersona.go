// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona}.
type AgenticApplicationsAnalystAgentPersona interface {
	cdktn.TerraformResource
	AnalystAgentPersonaId() *string
	SetAnalystAgentPersonaId(val *string)
	AnalystAgentPersonaIdInput() *string
	ArtifactExamples() AgenticApplicationsAnalystAgentPersonaArtifactExamplesList
	ArtifactExamplesInput() interface{}
	ArtifactsConfig() AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference
	ArtifactsConfigInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfig
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
	CustomerContext() *[]*string
	SetCustomerContext(val *[]*string)
	CustomerContextInput() *[]*string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayDescription() *string
	SetDisplayDescription(val *string)
	DisplayDescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExternalDataSources() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList
	ExternalDataSourcesInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeminiEnterpriseEngine() *string
	SetGeminiEnterpriseEngine(val *string)
	GeminiEnterpriseEngineInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	McpDataSources() AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList
	McpDataSourcesInput() interface{}
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	Resources() AgenticApplicationsAnalystAgentPersonaResourcesList
	ResourcesInput() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Skills() AgenticApplicationsAnalystAgentPersonaSkillsList
	SkillsInput() interface{}
	Tables() AgenticApplicationsAnalystAgentPersonaTablesList
	TablesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutArtifactExamples(value interface{})
	PutArtifactsConfig(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfig)
	PutExternalDataSources(value interface{})
	PutMcpDataSources(value interface{})
	PutResources(value interface{})
	PutSkills(value interface{})
	PutTables(value interface{})
	PutTimeouts(value *AgenticApplicationsAnalystAgentPersonaTimeouts)
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
	ResetArtifactExamples()
	ResetArtifactsConfig()
	ResetCustomerContext()
	ResetDeletionPolicy()
	ResetDisplayDescription()
	ResetExternalDataSources()
	ResetGeminiEnterpriseEngine()
	ResetId()
	ResetMcpDataSources()
	ResetModelDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResources()
	ResetRole()
	ResetSkills()
	ResetTables()
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

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersona
type jsiiProxy_AgenticApplicationsAnalystAgentPersona struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) AnalystAgentPersonaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analystAgentPersonaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) AnalystAgentPersonaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analystAgentPersonaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ArtifactExamples() AgenticApplicationsAnalystAgentPersonaArtifactExamplesList {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactExamplesList
	_jsii_.Get(
		j,
		"artifactExamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ArtifactExamplesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactExamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ArtifactsConfig() AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference
	_jsii_.Get(
		j,
		"artifactsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ArtifactsConfigInput() *AgenticApplicationsAnalystAgentPersonaArtifactsConfig {
	var returns *AgenticApplicationsAnalystAgentPersonaArtifactsConfig
	_jsii_.Get(
		j,
		"artifactsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) CustomerContext() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) CustomerContextInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DisplayDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DisplayDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ExternalDataSources() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList
	_jsii_.Get(
		j,
		"externalDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ExternalDataSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GeminiEnterpriseEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiEnterpriseEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GeminiEnterpriseEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiEnterpriseEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) McpDataSources() AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList {
	var returns AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList
	_jsii_.Get(
		j,
		"mcpDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) McpDataSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Resources() AgenticApplicationsAnalystAgentPersonaResourcesList {
	var returns AgenticApplicationsAnalystAgentPersonaResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Skills() AgenticApplicationsAnalystAgentPersonaSkillsList {
	var returns AgenticApplicationsAnalystAgentPersonaSkillsList
	_jsii_.Get(
		j,
		"skills",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) SkillsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skillsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Tables() AgenticApplicationsAnalystAgentPersonaTablesList {
	var returns AgenticApplicationsAnalystAgentPersonaTablesList
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) TablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) Timeouts() AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona} Resource.
func NewAgenticApplicationsAnalystAgentPersona(scope constructs.Construct, id *string, config *AgenticApplicationsAnalystAgentPersonaConfig) AgenticApplicationsAnalystAgentPersona {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersona{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona} Resource.
func NewAgenticApplicationsAnalystAgentPersona_Override(a AgenticApplicationsAnalystAgentPersona, scope constructs.Construct, id *string, config *AgenticApplicationsAnalystAgentPersonaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetAnalystAgentPersonaId(val *string) {
	if err := j.validateSetAnalystAgentPersonaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analystAgentPersonaId",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetCustomerContext(val *[]*string) {
	if err := j.validateSetCustomerContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerContext",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetDisplayDescription(val *string) {
	if err := j.validateSetDisplayDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetGeminiEnterpriseEngine(val *string) {
	if err := j.validateSetGeminiEnterpriseEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geminiEnterpriseEngine",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersona)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

// Generates CDKTN code for importing a AgenticApplicationsAnalystAgentPersona resource upon running "cdktn plan <stack-name>".
func AgenticApplicationsAnalystAgentPersona_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAgenticApplicationsAnalystAgentPersona_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
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
func AgenticApplicationsAnalystAgentPersona_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgenticApplicationsAnalystAgentPersona_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AgenticApplicationsAnalystAgentPersona_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgenticApplicationsAnalystAgentPersona_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AgenticApplicationsAnalystAgentPersona_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgenticApplicationsAnalystAgentPersona_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AgenticApplicationsAnalystAgentPersona_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutArtifactExamples(value interface{}) {
	if err := a.validatePutArtifactExamplesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArtifactExamples",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutArtifactsConfig(value *AgenticApplicationsAnalystAgentPersonaArtifactsConfig) {
	if err := a.validatePutArtifactsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArtifactsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutExternalDataSources(value interface{}) {
	if err := a.validatePutExternalDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExternalDataSources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutMcpDataSources(value interface{}) {
	if err := a.validatePutMcpDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMcpDataSources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutResources(value interface{}) {
	if err := a.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutSkills(value interface{}) {
	if err := a.validatePutSkillsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSkills",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutTables(value interface{}) {
	if err := a.validatePutTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTables",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) PutTimeouts(value *AgenticApplicationsAnalystAgentPersonaTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetArtifactExamples() {
	_jsii_.InvokeVoid(
		a,
		"resetArtifactExamples",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetArtifactsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetArtifactsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetCustomerContext() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetDisplayDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetExternalDataSources() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalDataSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetGeminiEnterpriseEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetGeminiEnterpriseEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetMcpDataSources() {
	_jsii_.InvokeVoid(
		a,
		"resetMcpDataSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetModelDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetResources() {
	_jsii_.InvokeVoid(
		a,
		"resetResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetRole() {
	_jsii_.InvokeVoid(
		a,
		"resetRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetSkills() {
	_jsii_.InvokeVoid(
		a,
		"resetSkills",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetTables() {
	_jsii_.InvokeVoid(
		a,
		"resetTables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersona) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

