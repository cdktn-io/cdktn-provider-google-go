// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-google-go/google/v19/networkservicesagentgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_services_agent_gateway google_network_services_agent_gateway}.
type NetworkServicesAgentGateway interface {
	cdktn.TerraformResource
	AgentGatewayCard() NetworkServicesAgentGatewayAgentGatewayCardList
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktn.StringMap
	Etag() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleManaged() NetworkServicesAgentGatewayGoogleManagedOutputReference
	GoogleManagedInput() *NetworkServicesAgentGatewayGoogleManaged
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() NetworkServicesAgentGatewayNetworkConfigOutputReference
	NetworkConfigInput() *NetworkServicesAgentGatewayNetworkConfig
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
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
	Registries() *[]*string
	SetRegistries(val *[]*string)
	RegistriesInput() *[]*string
	SelfManaged() NetworkServicesAgentGatewaySelfManagedOutputReference
	SelfManagedInput() *NetworkServicesAgentGatewaySelfManaged
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkServicesAgentGatewayTimeoutsOutputReference
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
	PutGoogleManaged(value *NetworkServicesAgentGatewayGoogleManaged)
	PutNetworkConfig(value *NetworkServicesAgentGatewayNetworkConfig)
	PutSelfManaged(value *NetworkServicesAgentGatewaySelfManaged)
	PutTimeouts(value *NetworkServicesAgentGatewayTimeouts)
	ResetDeletionPolicy()
	ResetDescription()
	ResetGoogleManaged()
	ResetId()
	ResetLabels()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProtocols()
	ResetRegistries()
	ResetSelfManaged()
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

// The jsii proxy struct for NetworkServicesAgentGateway
type jsiiProxy_NetworkServicesAgentGateway struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_NetworkServicesAgentGateway) AgentGatewayCard() NetworkServicesAgentGatewayAgentGatewayCardList {
	var returns NetworkServicesAgentGatewayAgentGatewayCardList
	_jsii_.Get(
		j,
		"agentGatewayCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) GoogleManaged() NetworkServicesAgentGatewayGoogleManagedOutputReference {
	var returns NetworkServicesAgentGatewayGoogleManagedOutputReference
	_jsii_.Get(
		j,
		"googleManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) GoogleManagedInput() *NetworkServicesAgentGatewayGoogleManaged {
	var returns *NetworkServicesAgentGatewayGoogleManaged
	_jsii_.Get(
		j,
		"googleManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) NetworkConfig() NetworkServicesAgentGatewayNetworkConfigOutputReference {
	var returns NetworkServicesAgentGatewayNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) NetworkConfigInput() *NetworkServicesAgentGatewayNetworkConfig {
	var returns *NetworkServicesAgentGatewayNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Registries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"registries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) RegistriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"registriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) SelfManaged() NetworkServicesAgentGatewaySelfManagedOutputReference {
	var returns NetworkServicesAgentGatewaySelfManagedOutputReference
	_jsii_.Get(
		j,
		"selfManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) SelfManagedInput() *NetworkServicesAgentGatewaySelfManaged {
	var returns *NetworkServicesAgentGatewaySelfManaged
	_jsii_.Get(
		j,
		"selfManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) Timeouts() NetworkServicesAgentGatewayTimeoutsOutputReference {
	var returns NetworkServicesAgentGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesAgentGateway) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_services_agent_gateway google_network_services_agent_gateway} Resource.
func NewNetworkServicesAgentGateway(scope constructs.Construct, id *string, config *NetworkServicesAgentGatewayConfig) NetworkServicesAgentGateway {
	_init_.Initialize()

	if err := validateNewNetworkServicesAgentGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesAgentGateway{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/network_services_agent_gateway google_network_services_agent_gateway} Resource.
func NewNetworkServicesAgentGateway_Override(n NetworkServicesAgentGateway, scope constructs.Construct, id *string, config *NetworkServicesAgentGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesAgentGateway)SetRegistries(val *[]*string) {
	if err := j.validateSetRegistriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registries",
		val,
	)
}

// Generates CDKTN code for importing a NetworkServicesAgentGateway resource upon running "cdktn plan <stack-name>".
func NetworkServicesAgentGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkServicesAgentGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
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
func NetworkServicesAgentGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesAgentGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesAgentGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesAgentGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesAgentGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesAgentGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkServicesAgentGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.networkServicesAgentGateway.NetworkServicesAgentGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) PutGoogleManaged(value *NetworkServicesAgentGatewayGoogleManaged) {
	if err := n.validatePutGoogleManagedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putGoogleManaged",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) PutNetworkConfig(value *NetworkServicesAgentGatewayNetworkConfig) {
	if err := n.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) PutSelfManaged(value *NetworkServicesAgentGatewaySelfManaged) {
	if err := n.validatePutSelfManagedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSelfManaged",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) PutTimeouts(value *NetworkServicesAgentGatewayTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetGoogleManaged() {
	_jsii_.InvokeVoid(
		n,
		"resetGoogleManaged",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetProtocols() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocols",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetRegistries() {
	_jsii_.InvokeVoid(
		n,
		"resetRegistries",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetSelfManaged() {
	_jsii_.InvokeVoid(
		n,
		"resetSelfManaged",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesAgentGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesAgentGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		n,
		"with",
		args,
		&returns,
	)

	return returns
}

