// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v21/datalosspreventioncontentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy google_data_loss_prevention_content_policy}.
type DataLossPreventionContentPolicy interface {
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
	DefaultAction() DataLossPreventionContentPolicyDefaultActionOutputReference
	DefaultActionInput() *DataLossPreventionContentPolicyDefaultAction
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
	Errors() DataLossPreventionContentPolicyErrorsList
	FailedToScanSupportedFileType() DataLossPreventionContentPolicyFailedToScanSupportedFileTypeOutputReference
	FailedToScanSupportedFileTypeInput() *DataLossPreventionContentPolicyFailedToScanSupportedFileType
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
	InputTooLarge() DataLossPreventionContentPolicyInputTooLargeOutputReference
	InputTooLargeInput() *DataLossPreventionContentPolicyInputTooLarge
	InspectConfig() DataLossPreventionContentPolicyInspectConfigOutputReference
	InspectConfigInput() *DataLossPreventionContentPolicyInspectConfig
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoggingConfigs() DataLossPreventionContentPolicyLoggingConfigsList
	LoggingConfigsInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
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
	Rules() DataLossPreventionContentPolicyRulesList
	RulesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataLossPreventionContentPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
	UnsupportedFileType() DataLossPreventionContentPolicyUnsupportedFileTypeOutputReference
	UnsupportedFileTypeInput() *DataLossPreventionContentPolicyUnsupportedFileType
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
	PutDefaultAction(value *DataLossPreventionContentPolicyDefaultAction)
	PutFailedToScanSupportedFileType(value *DataLossPreventionContentPolicyFailedToScanSupportedFileType)
	PutInputTooLarge(value *DataLossPreventionContentPolicyInputTooLarge)
	PutInspectConfig(value *DataLossPreventionContentPolicyInspectConfig)
	PutLoggingConfigs(value interface{})
	PutRules(value interface{})
	PutTimeouts(value *DataLossPreventionContentPolicyTimeouts)
	PutUnsupportedFileType(value *DataLossPreventionContentPolicyUnsupportedFileType)
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
	ResetDefaultAction()
	ResetDeletionPolicy()
	ResetDisplayName()
	ResetFailedToScanSupportedFileType()
	ResetId()
	ResetInputTooLarge()
	ResetInspectConfig()
	ResetLoggingConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetUnsupportedFileType()
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

// The jsii proxy struct for DataLossPreventionContentPolicy
type jsiiProxy_DataLossPreventionContentPolicy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DefaultAction() DataLossPreventionContentPolicyDefaultActionOutputReference {
	var returns DataLossPreventionContentPolicyDefaultActionOutputReference
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DefaultActionInput() *DataLossPreventionContentPolicyDefaultAction {
	var returns *DataLossPreventionContentPolicyDefaultAction
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Errors() DataLossPreventionContentPolicyErrorsList {
	var returns DataLossPreventionContentPolicyErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) FailedToScanSupportedFileType() DataLossPreventionContentPolicyFailedToScanSupportedFileTypeOutputReference {
	var returns DataLossPreventionContentPolicyFailedToScanSupportedFileTypeOutputReference
	_jsii_.Get(
		j,
		"failedToScanSupportedFileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) FailedToScanSupportedFileTypeInput() *DataLossPreventionContentPolicyFailedToScanSupportedFileType {
	var returns *DataLossPreventionContentPolicyFailedToScanSupportedFileType
	_jsii_.Get(
		j,
		"failedToScanSupportedFileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) InputTooLarge() DataLossPreventionContentPolicyInputTooLargeOutputReference {
	var returns DataLossPreventionContentPolicyInputTooLargeOutputReference
	_jsii_.Get(
		j,
		"inputTooLarge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) InputTooLargeInput() *DataLossPreventionContentPolicyInputTooLarge {
	var returns *DataLossPreventionContentPolicyInputTooLarge
	_jsii_.Get(
		j,
		"inputTooLargeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) InspectConfig() DataLossPreventionContentPolicyInspectConfigOutputReference {
	var returns DataLossPreventionContentPolicyInspectConfigOutputReference
	_jsii_.Get(
		j,
		"inspectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) InspectConfigInput() *DataLossPreventionContentPolicyInspectConfig {
	var returns *DataLossPreventionContentPolicyInspectConfig
	_jsii_.Get(
		j,
		"inspectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) LoggingConfigs() DataLossPreventionContentPolicyLoggingConfigsList {
	var returns DataLossPreventionContentPolicyLoggingConfigsList
	_jsii_.Get(
		j,
		"loggingConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) LoggingConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Rules() DataLossPreventionContentPolicyRulesList {
	var returns DataLossPreventionContentPolicyRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) Timeouts() DataLossPreventionContentPolicyTimeoutsOutputReference {
	var returns DataLossPreventionContentPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) UnsupportedFileType() DataLossPreventionContentPolicyUnsupportedFileTypeOutputReference {
	var returns DataLossPreventionContentPolicyUnsupportedFileTypeOutputReference
	_jsii_.Get(
		j,
		"unsupportedFileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) UnsupportedFileTypeInput() *DataLossPreventionContentPolicyUnsupportedFileType {
	var returns *DataLossPreventionContentPolicyUnsupportedFileType
	_jsii_.Get(
		j,
		"unsupportedFileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionContentPolicy) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy google_data_loss_prevention_content_policy} Resource.
func NewDataLossPreventionContentPolicy(scope constructs.Construct, id *string, config *DataLossPreventionContentPolicyConfig) DataLossPreventionContentPolicy {
	_init_.Initialize()

	if err := validateNewDataLossPreventionContentPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionContentPolicy{}

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy google_data_loss_prevention_content_policy} Resource.
func NewDataLossPreventionContentPolicy_Override(d DataLossPreventionContentPolicy, scope constructs.Construct, id *string, config *DataLossPreventionContentPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionContentPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a DataLossPreventionContentPolicy resource upon running "cdktn plan <stack-name>".
func DataLossPreventionContentPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataLossPreventionContentPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
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
func DataLossPreventionContentPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLossPreventionContentPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLossPreventionContentPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLossPreventionContentPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataLossPreventionContentPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLossPreventionContentPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataLossPreventionContentPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dataLossPreventionContentPolicy.DataLossPreventionContentPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DataLossPreventionContentPolicy) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutDefaultAction(value *DataLossPreventionContentPolicyDefaultAction) {
	if err := d.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutFailedToScanSupportedFileType(value *DataLossPreventionContentPolicyFailedToScanSupportedFileType) {
	if err := d.validatePutFailedToScanSupportedFileTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFailedToScanSupportedFileType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutInputTooLarge(value *DataLossPreventionContentPolicyInputTooLarge) {
	if err := d.validatePutInputTooLargeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInputTooLarge",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutInspectConfig(value *DataLossPreventionContentPolicyInspectConfig) {
	if err := d.validatePutInspectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInspectConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutLoggingConfigs(value interface{}) {
	if err := d.validatePutLoggingConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingConfigs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutRules(value interface{}) {
	if err := d.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutTimeouts(value *DataLossPreventionContentPolicyTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) PutUnsupportedFileType(value *DataLossPreventionContentPolicyUnsupportedFileType) {
	if err := d.validatePutUnsupportedFileTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnsupportedFileType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetDefaultAction() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetFailedToScanSupportedFileType() {
	_jsii_.InvokeVoid(
		d,
		"resetFailedToScanSupportedFileType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetInputTooLarge() {
	_jsii_.InvokeVoid(
		d,
		"resetInputTooLarge",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetInspectConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetInspectConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetLoggingConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ResetUnsupportedFileType() {
	_jsii_.InvokeVoid(
		d,
		"resetUnsupportedFileType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionContentPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

