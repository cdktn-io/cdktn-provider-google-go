// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-google-go/google/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-google-go/google/v20/agenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference interface {
	cdktn.ComplexObject
	AirQuality() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference
	AirQualityInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality
	BureauLaborStatistics() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference
	BureauLaborStatisticsInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics
	Coindesk() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference
	CoindeskInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Finnhub() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference
	FinnhubInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub
	// Experimental.
	Fqn() *string
	Fred() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference
	FredInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecEdgar() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference
	SecEdgarInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar
	SelectionName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TreasurySecuritiesAuctions() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference
	TreasurySecuritiesAuctionsInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions
	Usda() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference
	UsdaInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda
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
	PutAirQuality(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality)
	PutBureauLaborStatistics(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics)
	PutCoindesk(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk)
	PutFinnhub(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub)
	PutFred(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred)
	PutSecEdgar(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar)
	PutTreasurySecuritiesAuctions(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions)
	PutUsda(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda)
	ResetAirQuality()
	ResetBureauLaborStatistics()
	ResetCoindesk()
	ResetFinnhub()
	ResetFred()
	ResetSecEdgar()
	ResetTreasurySecuritiesAuctions()
	ResetUsda()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference
type jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) AirQuality() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference
	_jsii_.Get(
		j,
		"airQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) AirQualityInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality
	_jsii_.Get(
		j,
		"airQualityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) BureauLaborStatistics() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference
	_jsii_.Get(
		j,
		"bureauLaborStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) BureauLaborStatisticsInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics
	_jsii_.Get(
		j,
		"bureauLaborStatisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Coindesk() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference
	_jsii_.Get(
		j,
		"coindesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) CoindeskInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk
	_jsii_.Get(
		j,
		"coindeskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Finnhub() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference
	_jsii_.Get(
		j,
		"finnhub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) FinnhubInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub
	_jsii_.Get(
		j,
		"finnhubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Fred() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference
	_jsii_.Get(
		j,
		"fred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) FredInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred
	_jsii_.Get(
		j,
		"fredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SecEdgar() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference
	_jsii_.Get(
		j,
		"secEdgar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SecEdgarInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar
	_jsii_.Get(
		j,
		"secEdgarInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) SelectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TreasurySecuritiesAuctions() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference
	_jsii_.Get(
		j,
		"treasurySecuritiesAuctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) TreasurySecuritiesAuctionsInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions
	_jsii_.Get(
		j,
		"treasurySecuritiesAuctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Usda() AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference {
	var returns AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference
	_jsii_.Get(
		j,
		"usda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) UsdaInput() *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda {
	var returns *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda
	_jsii_.Get(
		j,
		"usdaInput",
		&returns,
	)
	return returns
}


func NewAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference_Override(a AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutAirQuality(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality) {
	if err := a.validatePutAirQualityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAirQuality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutBureauLaborStatistics(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics) {
	if err := a.validatePutBureauLaborStatisticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBureauLaborStatistics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutCoindesk(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk) {
	if err := a.validatePutCoindeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoindesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutFinnhub(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub) {
	if err := a.validatePutFinnhubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFinnhub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutFred(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred) {
	if err := a.validatePutFredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFred",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutSecEdgar(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar) {
	if err := a.validatePutSecEdgarParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecEdgar",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutTreasurySecuritiesAuctions(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions) {
	if err := a.validatePutTreasurySecuritiesAuctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTreasurySecuritiesAuctions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) PutUsda(value *AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda) {
	if err := a.validatePutUsdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetAirQuality() {
	_jsii_.InvokeVoid(
		a,
		"resetAirQuality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetBureauLaborStatistics() {
	_jsii_.InvokeVoid(
		a,
		"resetBureauLaborStatistics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetCoindesk() {
	_jsii_.InvokeVoid(
		a,
		"resetCoindesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetFinnhub() {
	_jsii_.InvokeVoid(
		a,
		"resetFinnhub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetFred() {
	_jsii_.InvokeVoid(
		a,
		"resetFred",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetSecEdgar() {
	_jsii_.InvokeVoid(
		a,
		"resetSecEdgar",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetTreasurySecuritiesAuctions() {
	_jsii_.InvokeVoid(
		a,
		"resetTreasurySecuritiesAuctions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ResetUsda() {
	_jsii_.InvokeVoid(
		a,
		"resetUsda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

