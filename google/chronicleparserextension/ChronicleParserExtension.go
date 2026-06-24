// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/chronicleparserextension/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_parser_extension google_chronicle_parser_extension}.
type ChronicleParserExtension interface {
	cdktn.TerraformResource
	CbnSnippet() *string
	SetCbnSnippet(val *string)
	CbnSnippetInput() *string
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
	DynamicParsing() ChronicleParserExtensionDynamicParsingOutputReference
	DynamicParsingInput() *ChronicleParserExtensionDynamicParsing
	ExtensionValidationReport() *string
	FieldExtractors() ChronicleParserExtensionFieldExtractorsOutputReference
	FieldExtractorsInput() *ChronicleParserExtensionFieldExtractors
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
	Instance() *string
	SetInstance(val *string)
	InstanceInput() *string
	LastLiveTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Log() *string
	SetLog(val *string)
	LogInput() *string
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parserextension() *string
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
	State() *string
	StateLastChangedTime() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ChronicleParserExtensionTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValidationReport() *string
	ValidationSkipped() interface{}
	SetValidationSkipped(val interface{})
	ValidationSkippedInput() interface{}
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutDynamicParsing(value *ChronicleParserExtensionDynamicParsing)
	PutFieldExtractors(value *ChronicleParserExtensionFieldExtractors)
	PutTimeouts(value *ChronicleParserExtensionTimeouts)
	ResetCbnSnippet()
	ResetDeletionPolicy()
	ResetDynamicParsing()
	ResetFieldExtractors()
	ResetId()
	ResetLog()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	ResetValidationSkipped()
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

// The jsii proxy struct for ChronicleParserExtension
type jsiiProxy_ChronicleParserExtension struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ChronicleParserExtension) CbnSnippet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cbnSnippet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) CbnSnippetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cbnSnippetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) DynamicParsing() ChronicleParserExtensionDynamicParsingOutputReference {
	var returns ChronicleParserExtensionDynamicParsingOutputReference
	_jsii_.Get(
		j,
		"dynamicParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) DynamicParsingInput() *ChronicleParserExtensionDynamicParsing {
	var returns *ChronicleParserExtensionDynamicParsing
	_jsii_.Get(
		j,
		"dynamicParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ExtensionValidationReport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionValidationReport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) FieldExtractors() ChronicleParserExtensionFieldExtractorsOutputReference {
	var returns ChronicleParserExtensionFieldExtractorsOutputReference
	_jsii_.Get(
		j,
		"fieldExtractors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) FieldExtractorsInput() *ChronicleParserExtensionFieldExtractors {
	var returns *ChronicleParserExtensionFieldExtractors
	_jsii_.Get(
		j,
		"fieldExtractorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Instance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) InstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) LastLiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastLiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Log() *string {
	var returns *string
	_jsii_.Get(
		j,
		"log",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) LogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Parserextension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserextension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) StateLastChangedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateLastChangedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) Timeouts() ChronicleParserExtensionTimeoutsOutputReference {
	var returns ChronicleParserExtensionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ValidationReport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationReport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ValidationSkipped() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationSkipped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleParserExtension) ValidationSkippedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationSkippedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_parser_extension google_chronicle_parser_extension} Resource.
func NewChronicleParserExtension(scope constructs.Construct, id *string, config *ChronicleParserExtensionConfig) ChronicleParserExtension {
	_init_.Initialize()

	if err := validateNewChronicleParserExtensionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleParserExtension{}

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/chronicle_parser_extension google_chronicle_parser_extension} Resource.
func NewChronicleParserExtension_Override(c ChronicleParserExtension, scope constructs.Construct, id *string, config *ChronicleParserExtensionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetCbnSnippet(val *string) {
	if err := j.validateSetCbnSnippetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cbnSnippet",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetInstance(val *string) {
	if err := j.validateSetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instance",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetLog(val *string) {
	if err := j.validateSetLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"log",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ChronicleParserExtension)SetValidationSkipped(val interface{}) {
	if err := j.validateSetValidationSkippedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationSkipped",
		val,
	)
}

// Generates CDKTN code for importing a ChronicleParserExtension resource upon running "cdktn plan <stack-name>".
func ChronicleParserExtension_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateChronicleParserExtension_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
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
func ChronicleParserExtension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChronicleParserExtension_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChronicleParserExtension_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChronicleParserExtension_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChronicleParserExtension_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChronicleParserExtension_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChronicleParserExtension_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.chronicleParserExtension.ChronicleParserExtension",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChronicleParserExtension) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserExtension) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleParserExtension) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleParserExtension) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleParserExtension) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleParserExtension) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleParserExtension) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleParserExtension) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleParserExtension) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChronicleParserExtension) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) PutDynamicParsing(value *ChronicleParserExtensionDynamicParsing) {
	if err := c.validatePutDynamicParsingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDynamicParsing",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) PutFieldExtractors(value *ChronicleParserExtensionFieldExtractors) {
	if err := c.validatePutFieldExtractorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFieldExtractors",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) PutTimeouts(value *ChronicleParserExtensionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetCbnSnippet() {
	_jsii_.InvokeVoid(
		c,
		"resetCbnSnippet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetDynamicParsing() {
	_jsii_.InvokeVoid(
		c,
		"resetDynamicParsing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetFieldExtractors() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldExtractors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetLog() {
	_jsii_.InvokeVoid(
		c,
		"resetLog",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) ResetValidationSkipped() {
	_jsii_.InvokeVoid(
		c,
		"resetValidationSkipped",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleParserExtension) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleParserExtension) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

