// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigtableappprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v17/bigtableappprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/bigtable_app_profile google_bigtable_app_profile}.
type BigtableAppProfile interface {
	cdktn.TerraformResource
	AppProfileId() *string
	SetAppProfileId(val *string)
	AppProfileIdInput() *string
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
	DataBoostIsolationReadOnly() BigtableAppProfileDataBoostIsolationReadOnlyOutputReference
	DataBoostIsolationReadOnlyInput() *BigtableAppProfileDataBoostIsolationReadOnly
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IgnoreWarnings() interface{}
	SetIgnoreWarnings(val interface{})
	IgnoreWarningsInput() interface{}
	Instance() *string
	SetInstance(val *string)
	InstanceInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MultiClusterRoutingClusterIds() *[]*string
	SetMultiClusterRoutingClusterIds(val *[]*string)
	MultiClusterRoutingClusterIdsInput() *[]*string
	MultiClusterRoutingUseAny() interface{}
	SetMultiClusterRoutingUseAny(val interface{})
	MultiClusterRoutingUseAnyInput() interface{}
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
	RowAffinity() interface{}
	SetRowAffinity(val interface{})
	RowAffinityInput() interface{}
	SingleClusterRouting() BigtableAppProfileSingleClusterRoutingOutputReference
	SingleClusterRoutingInput() *BigtableAppProfileSingleClusterRouting
	StandardIsolation() BigtableAppProfileStandardIsolationOutputReference
	StandardIsolationInput() *BigtableAppProfileStandardIsolation
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigtableAppProfileTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDataBoostIsolationReadOnly(value *BigtableAppProfileDataBoostIsolationReadOnly)
	PutSingleClusterRouting(value *BigtableAppProfileSingleClusterRouting)
	PutStandardIsolation(value *BigtableAppProfileStandardIsolation)
	PutTimeouts(value *BigtableAppProfileTimeouts)
	ResetDataBoostIsolationReadOnly()
	ResetDescription()
	ResetId()
	ResetIgnoreWarnings()
	ResetInstance()
	ResetMultiClusterRoutingClusterIds()
	ResetMultiClusterRoutingUseAny()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRowAffinity()
	ResetSingleClusterRouting()
	ResetStandardIsolation()
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
}

// The jsii proxy struct for BigtableAppProfile
type jsiiProxy_BigtableAppProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BigtableAppProfile) AppProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) AppProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) DataBoostIsolationReadOnly() BigtableAppProfileDataBoostIsolationReadOnlyOutputReference {
	var returns BigtableAppProfileDataBoostIsolationReadOnlyOutputReference
	_jsii_.Get(
		j,
		"dataBoostIsolationReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) DataBoostIsolationReadOnlyInput() *BigtableAppProfileDataBoostIsolationReadOnly {
	var returns *BigtableAppProfileDataBoostIsolationReadOnly
	_jsii_.Get(
		j,
		"dataBoostIsolationReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) IgnoreWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) IgnoreWarningsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Instance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) InstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) MultiClusterRoutingClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"multiClusterRoutingClusterIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) MultiClusterRoutingClusterIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"multiClusterRoutingClusterIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) MultiClusterRoutingUseAny() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiClusterRoutingUseAny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) MultiClusterRoutingUseAnyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiClusterRoutingUseAnyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) RowAffinity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) RowAffinityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) SingleClusterRouting() BigtableAppProfileSingleClusterRoutingOutputReference {
	var returns BigtableAppProfileSingleClusterRoutingOutputReference
	_jsii_.Get(
		j,
		"singleClusterRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) SingleClusterRoutingInput() *BigtableAppProfileSingleClusterRouting {
	var returns *BigtableAppProfileSingleClusterRouting
	_jsii_.Get(
		j,
		"singleClusterRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) StandardIsolation() BigtableAppProfileStandardIsolationOutputReference {
	var returns BigtableAppProfileStandardIsolationOutputReference
	_jsii_.Get(
		j,
		"standardIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) StandardIsolationInput() *BigtableAppProfileStandardIsolation {
	var returns *BigtableAppProfileStandardIsolation
	_jsii_.Get(
		j,
		"standardIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) Timeouts() BigtableAppProfileTimeoutsOutputReference {
	var returns BigtableAppProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableAppProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/bigtable_app_profile google_bigtable_app_profile} Resource.
func NewBigtableAppProfile(scope constructs.Construct, id *string, config *BigtableAppProfileConfig) BigtableAppProfile {
	_init_.Initialize()

	if err := validateNewBigtableAppProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigtableAppProfile{}

	_jsii_.Create(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/bigtable_app_profile google_bigtable_app_profile} Resource.
func NewBigtableAppProfile_Override(b BigtableAppProfile, scope constructs.Construct, id *string, config *BigtableAppProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetAppProfileId(val *string) {
	if err := j.validateSetAppProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appProfileId",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetIgnoreWarnings(val interface{}) {
	if err := j.validateSetIgnoreWarningsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreWarnings",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetInstance(val *string) {
	if err := j.validateSetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instance",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetMultiClusterRoutingClusterIds(val *[]*string) {
	if err := j.validateSetMultiClusterRoutingClusterIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiClusterRoutingClusterIds",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetMultiClusterRoutingUseAny(val interface{}) {
	if err := j.validateSetMultiClusterRoutingUseAnyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiClusterRoutingUseAny",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigtableAppProfile)SetRowAffinity(val interface{}) {
	if err := j.validateSetRowAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowAffinity",
		val,
	)
}

// Generates CDKTN code for importing a BigtableAppProfile resource upon running "cdktn plan <stack-name>".
func BigtableAppProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBigtableAppProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
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
func BigtableAppProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigtableAppProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigtableAppProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigtableAppProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigtableAppProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigtableAppProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BigtableAppProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.bigtableAppProfile.BigtableAppProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigtableAppProfile) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BigtableAppProfile) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BigtableAppProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BigtableAppProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigtableAppProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BigtableAppProfile) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BigtableAppProfile) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BigtableAppProfile) PutDataBoostIsolationReadOnly(value *BigtableAppProfileDataBoostIsolationReadOnly) {
	if err := b.validatePutDataBoostIsolationReadOnlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDataBoostIsolationReadOnly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigtableAppProfile) PutSingleClusterRouting(value *BigtableAppProfileSingleClusterRouting) {
	if err := b.validatePutSingleClusterRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSingleClusterRouting",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigtableAppProfile) PutStandardIsolation(value *BigtableAppProfileStandardIsolation) {
	if err := b.validatePutStandardIsolationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStandardIsolation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigtableAppProfile) PutTimeouts(value *BigtableAppProfileTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetDataBoostIsolationReadOnly() {
	_jsii_.InvokeVoid(
		b,
		"resetDataBoostIsolationReadOnly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetIgnoreWarnings() {
	_jsii_.InvokeVoid(
		b,
		"resetIgnoreWarnings",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetInstance() {
	_jsii_.InvokeVoid(
		b,
		"resetInstance",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetMultiClusterRoutingClusterIds() {
	_jsii_.InvokeVoid(
		b,
		"resetMultiClusterRoutingClusterIds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetMultiClusterRoutingUseAny() {
	_jsii_.InvokeVoid(
		b,
		"resetMultiClusterRoutingUseAny",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetRowAffinity() {
	_jsii_.InvokeVoid(
		b,
		"resetRowAffinity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetSingleClusterRouting() {
	_jsii_.InvokeVoid(
		b,
		"resetSingleClusterRouting",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetStandardIsolation() {
	_jsii_.InvokeVoid(
		b,
		"resetStandardIsolation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableAppProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableAppProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

