// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigeeenvironmentdebugmask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/apigeeenvironmentdebugmask/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/apigee_environment_debugmask google_apigee_environment_debugmask}.
type ApigeeEnvironmentDebugmask interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvId() *string
	SetEnvId(val *string)
	EnvIdInput() *string
	FaultXPaths() *[]*string
	SetFaultXPaths(val *[]*string)
	FaultXPathsInput() *[]*string
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	Namespaces() *map[string]*string
	SetNamespaces(val *map[string]*string)
	NamespacesInput() *map[string]*string
	// The tree node.
	Node() constructs.Node
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
	RequestJsonPaths() *[]*string
	SetRequestJsonPaths(val *[]*string)
	RequestJsonPathsInput() *[]*string
	RequestXPaths() *[]*string
	SetRequestXPaths(val *[]*string)
	RequestXPathsInput() *[]*string
	ResponseJsonPaths() *[]*string
	SetResponseJsonPaths(val *[]*string)
	ResponseJsonPathsInput() *[]*string
	ResponseXPaths() *[]*string
	SetResponseXPaths(val *[]*string)
	ResponseXPathsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApigeeEnvironmentDebugmaskTimeoutsOutputReference
	TimeoutsInput() interface{}
	Variables() *[]*string
	SetVariables(val *[]*string)
	VariablesInput() *[]*string
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
	PutTimeouts(value *ApigeeEnvironmentDebugmaskTimeouts)
	ResetFaultXPaths()
	ResetId()
	ResetNamespaces()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestJsonPaths()
	ResetRequestXPaths()
	ResetResponseJsonPaths()
	ResetResponseXPaths()
	ResetTimeouts()
	ResetVariables()
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

// The jsii proxy struct for ApigeeEnvironmentDebugmask
type jsiiProxy_ApigeeEnvironmentDebugmask struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) EnvId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) EnvIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) FaultXPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faultXPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) FaultXPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"faultXPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Namespaces() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"namespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) NamespacesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"namespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) RequestJsonPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestJsonPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) RequestJsonPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestJsonPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) RequestXPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestXPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) RequestXPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestXPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ResponseJsonPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseJsonPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ResponseJsonPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseJsonPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ResponseXPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseXPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) ResponseXPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseXPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Timeouts() ApigeeEnvironmentDebugmaskTimeoutsOutputReference {
	var returns ApigeeEnvironmentDebugmaskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) Variables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask) VariablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/apigee_environment_debugmask google_apigee_environment_debugmask} Resource.
func NewApigeeEnvironmentDebugmask(scope constructs.Construct, id *string, config *ApigeeEnvironmentDebugmaskConfig) ApigeeEnvironmentDebugmask {
	_init_.Initialize()

	if err := validateNewApigeeEnvironmentDebugmaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeEnvironmentDebugmask{}

	_jsii_.Create(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/apigee_environment_debugmask google_apigee_environment_debugmask} Resource.
func NewApigeeEnvironmentDebugmask_Override(a ApigeeEnvironmentDebugmask, scope constructs.Construct, id *string, config *ApigeeEnvironmentDebugmaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetEnvId(val *string) {
	if err := j.validateSetEnvIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envId",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetFaultXPaths(val *[]*string) {
	if err := j.validateSetFaultXPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faultXPaths",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetNamespaces(val *map[string]*string) {
	if err := j.validateSetNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaces",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetRequestJsonPaths(val *[]*string) {
	if err := j.validateSetRequestJsonPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestJsonPaths",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetRequestXPaths(val *[]*string) {
	if err := j.validateSetRequestXPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestXPaths",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetResponseJsonPaths(val *[]*string) {
	if err := j.validateSetResponseJsonPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseJsonPaths",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetResponseXPaths(val *[]*string) {
	if err := j.validateSetResponseXPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseXPaths",
		val,
	)
}

func (j *jsiiProxy_ApigeeEnvironmentDebugmask)SetVariables(val *[]*string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Generates CDKTN code for importing a ApigeeEnvironmentDebugmask resource upon running "cdktn plan <stack-name>".
func ApigeeEnvironmentDebugmask_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApigeeEnvironmentDebugmask_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
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
func ApigeeEnvironmentDebugmask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeEnvironmentDebugmask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeEnvironmentDebugmask_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeEnvironmentDebugmask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApigeeEnvironmentDebugmask_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApigeeEnvironmentDebugmask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApigeeEnvironmentDebugmask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.apigeeEnvironmentDebugmask.ApigeeEnvironmentDebugmask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) PutTimeouts(value *ApigeeEnvironmentDebugmaskTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetFaultXPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetFaultXPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetNamespaces() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespaces",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetRequestJsonPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestJsonPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetRequestXPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestXPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetResponseJsonPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseJsonPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetResponseXPaths() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseXPaths",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ResetVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeEnvironmentDebugmask) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

