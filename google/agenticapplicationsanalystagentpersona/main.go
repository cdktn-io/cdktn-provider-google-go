// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersona",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersona)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "analystAgentPersonaId", GoGetter: "AnalystAgentPersonaId"},
			_jsii_.MemberProperty{JsiiProperty: "analystAgentPersonaIdInput", GoGetter: "AnalystAgentPersonaIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "artifactExamples", GoGetter: "ArtifactExamples"},
			_jsii_.MemberProperty{JsiiProperty: "artifactExamplesInput", GoGetter: "ArtifactExamplesInput"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsConfig", GoGetter: "ArtifactsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsConfigInput", GoGetter: "ArtifactsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "customerContext", GoGetter: "CustomerContext"},
			_jsii_.MemberProperty{JsiiProperty: "customerContextInput", GoGetter: "CustomerContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "deletionPolicy", GoGetter: "DeletionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "deletionPolicyInput", GoGetter: "DeletionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayDescription", GoGetter: "DisplayDescription"},
			_jsii_.MemberProperty{JsiiProperty: "displayDescriptionInput", GoGetter: "DisplayDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalDataSources", GoGetter: "ExternalDataSources"},
			_jsii_.MemberProperty{JsiiProperty: "externalDataSourcesInput", GoGetter: "ExternalDataSourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "geminiEnterpriseEngine", GoGetter: "GeminiEnterpriseEngine"},
			_jsii_.MemberProperty{JsiiProperty: "geminiEnterpriseEngineInput", GoGetter: "GeminiEnterpriseEngineInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mcpDataSources", GoGetter: "McpDataSources"},
			_jsii_.MemberProperty{JsiiProperty: "mcpDataSourcesInput", GoGetter: "McpDataSourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putArtifactExamples", GoMethod: "PutArtifactExamples"},
			_jsii_.MemberMethod{JsiiMethod: "putArtifactsConfig", GoMethod: "PutArtifactsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalDataSources", GoMethod: "PutExternalDataSources"},
			_jsii_.MemberMethod{JsiiMethod: "putMcpDataSources", GoMethod: "PutMcpDataSources"},
			_jsii_.MemberMethod{JsiiMethod: "putResources", GoMethod: "PutResources"},
			_jsii_.MemberMethod{JsiiMethod: "putSkills", GoMethod: "PutSkills"},
			_jsii_.MemberMethod{JsiiMethod: "putTables", GoMethod: "PutTables"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetArtifactExamples", GoMethod: "ResetArtifactExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resetArtifactsConfig", GoMethod: "ResetArtifactsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerContext", GoMethod: "ResetCustomerContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeletionPolicy", GoMethod: "ResetDeletionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayDescription", GoMethod: "ResetDisplayDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalDataSources", GoMethod: "ResetExternalDataSources"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeminiEnterpriseEngine", GoMethod: "ResetGeminiEnterpriseEngine"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMcpDataSources", GoMethod: "ResetMcpDataSources"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetResources", GoMethod: "ResetResources"},
			_jsii_.MemberMethod{JsiiMethod: "resetRole", GoMethod: "ResetRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkills", GoMethod: "ResetSkills"},
			_jsii_.MemberMethod{JsiiMethod: "resetTables", GoMethod: "ResetTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resources", GoGetter: "Resources"},
			_jsii_.MemberProperty{JsiiProperty: "resourcesInput", GoGetter: "ResourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleInput", GoGetter: "RoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "skills", GoGetter: "Skills"},
			_jsii_.MemberProperty{JsiiProperty: "skillsInput", GoGetter: "SkillsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tables", GoGetter: "Tables"},
			_jsii_.MemberProperty{JsiiProperty: "tablesInput", GoGetter: "TablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateTime", GoGetter: "UpdateTime"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersona{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamples",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamples)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putResource", GoMethod: "PutResource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDataset", GoGetter: "BigqueryDataset"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDatasetInput", GoGetter: "BigqueryDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTable", GoGetter: "BigqueryTable"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTableInput", GoGetter: "BigqueryTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptions", GoGetter: "ColumnDescriptions"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptionsInput", GoGetter: "ColumnDescriptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDataset", GoMethod: "ResetBigqueryDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryTable", GoMethod: "ResetBigqueryTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnDescriptions", GoMethod: "ResetColumnDescriptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceBigqueryResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1Resource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "f1Table", GoGetter: "F1Table"},
			_jsii_.MemberProperty{JsiiProperty: "f1TableInput", GoGetter: "F1TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Table", GoMethod: "ResetF1Table"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceF1ResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObject", GoGetter: "GoogleCloudStorageObject"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObjectInput", GoGetter: "GoogleCloudStorageObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleCloudStorageResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileReference", GoGetter: "FileReference"},
			_jsii_.MemberProperty{JsiiProperty: "fileReferenceInput", GoGetter: "FileReferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileReference", GoMethod: "ResetFileReference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceGoogleDriveResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResource", GoGetter: "BigqueryResource"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResourceInput", GoGetter: "BigqueryResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabel", GoGetter: "DisplayLabel"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabelInput", GoGetter: "DisplayLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "f1Resource", GoGetter: "F1Resource"},
			_jsii_.MemberProperty{JsiiProperty: "f1ResourceInput", GoGetter: "F1ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResource", GoGetter: "GoogleCloudStorageResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResourceInput", GoGetter: "GoogleCloudStorageResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResource", GoGetter: "GoogleDriveResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResourceInput", GoGetter: "GoogleDriveResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryResource", GoMethod: "PutBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "putF1Resource", GoMethod: "PutF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudStorageResource", GoMethod: "PutGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleDriveResource", GoMethod: "PutGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "putRawFileResource", GoMethod: "PutRawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResource", GoGetter: "RawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResourceInput", GoGetter: "RawFileResourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryResource", GoMethod: "ResetBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayLabel", GoMethod: "ResetDisplayLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Resource", GoMethod: "ResetF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudStorageResource", GoMethod: "ResetGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDriveResource", GoMethod: "ResetGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRawFileResource", GoMethod: "ResetRawFileResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseRag", GoMethod: "ResetUseRag"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useRag", GoGetter: "UseRag"},
			_jsii_.MemberProperty{JsiiProperty: "useRagInput", GoGetter: "UseRagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileContentInput", GoGetter: "FileContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitle", GoGetter: "FileTitle"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitleInput", GoGetter: "FileTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mimeType", GoGetter: "MimeType"},
			_jsii_.MemberProperty{JsiiProperty: "mimeTypeInput", GoGetter: "MimeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactExamplesResourceRawFileResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfig",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptions",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamples",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamples)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putResource", GoMethod: "PutResource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDataset", GoGetter: "BigqueryDataset"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDatasetInput", GoGetter: "BigqueryDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTable", GoGetter: "BigqueryTable"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTableInput", GoGetter: "BigqueryTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptions", GoGetter: "ColumnDescriptions"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptionsInput", GoGetter: "ColumnDescriptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDataset", GoMethod: "ResetBigqueryDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryTable", GoMethod: "ResetBigqueryTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnDescriptions", GoMethod: "ResetColumnDescriptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceBigqueryResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1Resource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "f1Table", GoGetter: "F1Table"},
			_jsii_.MemberProperty{JsiiProperty: "f1TableInput", GoGetter: "F1TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Table", GoMethod: "ResetF1Table"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceF1ResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObject", GoGetter: "GoogleCloudStorageObject"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObjectInput", GoGetter: "GoogleCloudStorageObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleCloudStorageResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileReference", GoGetter: "FileReference"},
			_jsii_.MemberProperty{JsiiProperty: "fileReferenceInput", GoGetter: "FileReferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileReference", GoMethod: "ResetFileReference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceGoogleDriveResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResource", GoGetter: "BigqueryResource"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResourceInput", GoGetter: "BigqueryResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabel", GoGetter: "DisplayLabel"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabelInput", GoGetter: "DisplayLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "f1Resource", GoGetter: "F1Resource"},
			_jsii_.MemberProperty{JsiiProperty: "f1ResourceInput", GoGetter: "F1ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResource", GoGetter: "GoogleCloudStorageResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResourceInput", GoGetter: "GoogleCloudStorageResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResource", GoGetter: "GoogleDriveResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResourceInput", GoGetter: "GoogleDriveResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryResource", GoMethod: "PutBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "putF1Resource", GoMethod: "PutF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudStorageResource", GoMethod: "PutGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleDriveResource", GoMethod: "PutGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "putRawFileResource", GoMethod: "PutRawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResource", GoGetter: "RawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResourceInput", GoGetter: "RawFileResourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryResource", GoMethod: "ResetBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayLabel", GoMethod: "ResetDisplayLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Resource", GoMethod: "ResetF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudStorageResource", GoMethod: "ResetGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDriveResource", GoMethod: "ResetGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRawFileResource", GoMethod: "ResetRawFileResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseRag", GoMethod: "ResetUseRag"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useRag", GoGetter: "UseRag"},
			_jsii_.MemberProperty{JsiiProperty: "useRagInput", GoGetter: "UseRagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileContentInput", GoGetter: "FileContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitle", GoGetter: "FileTitle"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitleInput", GoGetter: "FileTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mimeType", GoGetter: "MimeType"},
			_jsii_.MemberProperty{JsiiProperty: "mimeTypeInput", GoGetter: "MimeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsDocumentExamplesResourceRawFileResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentExamples", GoGetter: "DocumentExamples"},
			_jsii_.MemberProperty{JsiiProperty: "documentExamplesInput", GoGetter: "DocumentExamplesInput"},
			_jsii_.MemberProperty{JsiiProperty: "exportFormat", GoGetter: "ExportFormat"},
			_jsii_.MemberProperty{JsiiProperty: "exportFormatInput", GoGetter: "ExportFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDocumentExamples", GoMethod: "PutDocumentExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentExamples", GoMethod: "ResetDocumentExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportFormat", GoMethod: "ResetExportFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documentGenerationOptions", GoGetter: "DocumentGenerationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "documentGenerationOptionsInput", GoGetter: "DocumentGenerationOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDocumentGenerationOptions", GoMethod: "PutDocumentGenerationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putSlideGenerationOptions", GoMethod: "PutSlideGenerationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putVisualizationOptions", GoMethod: "PutVisualizationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentGenerationOptions", GoMethod: "ResetDocumentGenerationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlideGenerationOptions", GoMethod: "ResetSlideGenerationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisualizationOptions", GoMethod: "ResetVisualizationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "slideGenerationOptions", GoGetter: "SlideGenerationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "slideGenerationOptionsInput", GoGetter: "SlideGenerationOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationOptions", GoGetter: "VisualizationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationOptionsInput", GoGetter: "VisualizationOptionsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "exportFormat", GoGetter: "ExportFormat"},
			_jsii_.MemberProperty{JsiiProperty: "exportFormatInput", GoGetter: "ExportFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putSlideExamples", GoMethod: "PutSlideExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportFormat", GoMethod: "ResetExportFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlideExamples", GoMethod: "ResetSlideExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "slideExamples", GoGetter: "SlideExamples"},
			_jsii_.MemberProperty{JsiiProperty: "slideExamplesInput", GoGetter: "SlideExamplesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamples",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamples)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putResource", GoMethod: "PutResource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDataset", GoGetter: "BigqueryDataset"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDatasetInput", GoGetter: "BigqueryDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTable", GoGetter: "BigqueryTable"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTableInput", GoGetter: "BigqueryTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptions", GoGetter: "ColumnDescriptions"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptionsInput", GoGetter: "ColumnDescriptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDataset", GoMethod: "ResetBigqueryDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryTable", GoMethod: "ResetBigqueryTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnDescriptions", GoMethod: "ResetColumnDescriptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceBigqueryResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1Resource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "f1Table", GoGetter: "F1Table"},
			_jsii_.MemberProperty{JsiiProperty: "f1TableInput", GoGetter: "F1TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Table", GoMethod: "ResetF1Table"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceF1ResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObject", GoGetter: "GoogleCloudStorageObject"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObjectInput", GoGetter: "GoogleCloudStorageObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleCloudStorageResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileReference", GoGetter: "FileReference"},
			_jsii_.MemberProperty{JsiiProperty: "fileReferenceInput", GoGetter: "FileReferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileReference", GoMethod: "ResetFileReference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceGoogleDriveResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResource", GoGetter: "BigqueryResource"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResourceInput", GoGetter: "BigqueryResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabel", GoGetter: "DisplayLabel"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabelInput", GoGetter: "DisplayLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "f1Resource", GoGetter: "F1Resource"},
			_jsii_.MemberProperty{JsiiProperty: "f1ResourceInput", GoGetter: "F1ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResource", GoGetter: "GoogleCloudStorageResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResourceInput", GoGetter: "GoogleCloudStorageResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResource", GoGetter: "GoogleDriveResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResourceInput", GoGetter: "GoogleDriveResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryResource", GoMethod: "PutBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "putF1Resource", GoMethod: "PutF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudStorageResource", GoMethod: "PutGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleDriveResource", GoMethod: "PutGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "putRawFileResource", GoMethod: "PutRawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResource", GoGetter: "RawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResourceInput", GoGetter: "RawFileResourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryResource", GoMethod: "ResetBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayLabel", GoMethod: "ResetDisplayLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Resource", GoMethod: "ResetF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudStorageResource", GoMethod: "ResetGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDriveResource", GoMethod: "ResetGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRawFileResource", GoMethod: "ResetRawFileResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseRag", GoMethod: "ResetUseRag"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useRag", GoGetter: "UseRag"},
			_jsii_.MemberProperty{JsiiProperty: "useRagInput", GoGetter: "UseRagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileContentInput", GoGetter: "FileContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitle", GoGetter: "FileTitle"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitleInput", GoGetter: "FileTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mimeType", GoGetter: "MimeType"},
			_jsii_.MemberProperty{JsiiProperty: "mimeTypeInput", GoGetter: "MimeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptionsSlideExamplesResourceRawFileResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptions",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putVisualizationExamples", GoMethod: "PutVisualizationExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisualizationExamples", GoMethod: "ResetVisualizationExamples"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationExamples", GoGetter: "VisualizationExamples"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationExamplesInput", GoGetter: "VisualizationExamplesInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamples",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamples)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putResource", GoMethod: "PutResource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationType", GoGetter: "VisualizationType"},
			_jsii_.MemberProperty{JsiiProperty: "visualizationTypeInput", GoGetter: "VisualizationTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDataset", GoGetter: "BigqueryDataset"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDatasetInput", GoGetter: "BigqueryDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTable", GoGetter: "BigqueryTable"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTableInput", GoGetter: "BigqueryTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptions", GoGetter: "ColumnDescriptions"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptionsInput", GoGetter: "ColumnDescriptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDataset", GoMethod: "ResetBigqueryDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryTable", GoMethod: "ResetBigqueryTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnDescriptions", GoMethod: "ResetColumnDescriptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceBigqueryResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1Resource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "f1Table", GoGetter: "F1Table"},
			_jsii_.MemberProperty{JsiiProperty: "f1TableInput", GoGetter: "F1TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Table", GoMethod: "ResetF1Table"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceF1ResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObject", GoGetter: "GoogleCloudStorageObject"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObjectInput", GoGetter: "GoogleCloudStorageObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileReference", GoGetter: "FileReference"},
			_jsii_.MemberProperty{JsiiProperty: "fileReferenceInput", GoGetter: "FileReferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileReference", GoMethod: "ResetFileReference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleDriveResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResource", GoGetter: "BigqueryResource"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResourceInput", GoGetter: "BigqueryResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabel", GoGetter: "DisplayLabel"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabelInput", GoGetter: "DisplayLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "f1Resource", GoGetter: "F1Resource"},
			_jsii_.MemberProperty{JsiiProperty: "f1ResourceInput", GoGetter: "F1ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResource", GoGetter: "GoogleCloudStorageResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResourceInput", GoGetter: "GoogleCloudStorageResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResource", GoGetter: "GoogleDriveResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResourceInput", GoGetter: "GoogleDriveResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryResource", GoMethod: "PutBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "putF1Resource", GoMethod: "PutF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudStorageResource", GoMethod: "PutGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleDriveResource", GoMethod: "PutGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "putRawFileResource", GoMethod: "PutRawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResource", GoGetter: "RawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResourceInput", GoGetter: "RawFileResourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryResource", GoMethod: "ResetBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayLabel", GoMethod: "ResetDisplayLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Resource", GoMethod: "ResetF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudStorageResource", GoMethod: "ResetGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDriveResource", GoMethod: "ResetGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRawFileResource", GoMethod: "ResetRawFileResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseRag", GoMethod: "ResetUseRag"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useRag", GoGetter: "UseRag"},
			_jsii_.MemberProperty{JsiiProperty: "useRagInput", GoGetter: "UseRagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileContentInput", GoGetter: "FileContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitle", GoGetter: "FileTitle"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitleInput", GoGetter: "FileTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mimeType", GoGetter: "MimeType"},
			_jsii_.MemberProperty{JsiiProperty: "mimeTypeInput", GoGetter: "MimeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceRawFileResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaConfig",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSources",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQuality)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesAirQualityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatistics)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesBureauLaborStatisticsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindesk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesCoindeskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhub)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFinnhubOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFred)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesFredOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "airQuality", GoGetter: "AirQuality"},
			_jsii_.MemberProperty{JsiiProperty: "airQualityInput", GoGetter: "AirQualityInput"},
			_jsii_.MemberProperty{JsiiProperty: "bureauLaborStatistics", GoGetter: "BureauLaborStatistics"},
			_jsii_.MemberProperty{JsiiProperty: "bureauLaborStatisticsInput", GoGetter: "BureauLaborStatisticsInput"},
			_jsii_.MemberProperty{JsiiProperty: "coindesk", GoGetter: "Coindesk"},
			_jsii_.MemberProperty{JsiiProperty: "coindeskInput", GoGetter: "CoindeskInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "finnhub", GoGetter: "Finnhub"},
			_jsii_.MemberProperty{JsiiProperty: "finnhubInput", GoGetter: "FinnhubInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "fred", GoGetter: "Fred"},
			_jsii_.MemberProperty{JsiiProperty: "fredInput", GoGetter: "FredInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAirQuality", GoMethod: "PutAirQuality"},
			_jsii_.MemberMethod{JsiiMethod: "putBureauLaborStatistics", GoMethod: "PutBureauLaborStatistics"},
			_jsii_.MemberMethod{JsiiMethod: "putCoindesk", GoMethod: "PutCoindesk"},
			_jsii_.MemberMethod{JsiiMethod: "putFinnhub", GoMethod: "PutFinnhub"},
			_jsii_.MemberMethod{JsiiMethod: "putFred", GoMethod: "PutFred"},
			_jsii_.MemberMethod{JsiiMethod: "putSecEdgar", GoMethod: "PutSecEdgar"},
			_jsii_.MemberMethod{JsiiMethod: "putTreasurySecuritiesAuctions", GoMethod: "PutTreasurySecuritiesAuctions"},
			_jsii_.MemberMethod{JsiiMethod: "putUsda", GoMethod: "PutUsda"},
			_jsii_.MemberMethod{JsiiMethod: "resetAirQuality", GoMethod: "ResetAirQuality"},
			_jsii_.MemberMethod{JsiiMethod: "resetBureauLaborStatistics", GoMethod: "ResetBureauLaborStatistics"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoindesk", GoMethod: "ResetCoindesk"},
			_jsii_.MemberMethod{JsiiMethod: "resetFinnhub", GoMethod: "ResetFinnhub"},
			_jsii_.MemberMethod{JsiiMethod: "resetFred", GoMethod: "ResetFred"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecEdgar", GoMethod: "ResetSecEdgar"},
			_jsii_.MemberMethod{JsiiMethod: "resetTreasurySecuritiesAuctions", GoMethod: "ResetTreasurySecuritiesAuctions"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsda", GoMethod: "ResetUsda"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secEdgar", GoGetter: "SecEdgar"},
			_jsii_.MemberProperty{JsiiProperty: "secEdgarInput", GoGetter: "SecEdgarInput"},
			_jsii_.MemberProperty{JsiiProperty: "selectionName", GoGetter: "SelectionName"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "treasurySecuritiesAuctions", GoGetter: "TreasurySecuritiesAuctions"},
			_jsii_.MemberProperty{JsiiProperty: "treasurySecuritiesAuctionsInput", GoGetter: "TreasurySecuritiesAuctionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "usda", GoGetter: "Usda"},
			_jsii_.MemberProperty{JsiiProperty: "usdaInput", GoGetter: "UsdaInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgar)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesSecEdgarOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesTreasurySecuritiesAuctionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsda)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaExternalDataSourcesUsdaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaMcpDataSources",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaMcpDataSources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaMcpDataSourcesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaMcpDataSourcesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaMcpDataSourcesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyInput", GoGetter: "ApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyName", GoGetter: "ApiKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyNameInput", GoGetter: "ApiKeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "oauthTokenUrl", GoGetter: "OauthTokenUrl"},
			_jsii_.MemberProperty{JsiiProperty: "oauthTokenUrlInput", GoGetter: "OauthTokenUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "prompt", GoGetter: "Prompt"},
			_jsii_.MemberProperty{JsiiProperty: "promptInput", GoGetter: "PromptInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKey", GoMethod: "ResetApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKeyName", GoMethod: "ResetApiKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthTokenUrl", GoMethod: "ResetOauthTokenUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrompt", GoMethod: "ResetPrompt"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serverUrl", GoGetter: "ServerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "serverUrlInput", GoGetter: "ServerUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaMcpDataSourcesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResources",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDataset", GoGetter: "BigqueryDataset"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDatasetInput", GoGetter: "BigqueryDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTable", GoGetter: "BigqueryTable"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryTableInput", GoGetter: "BigqueryTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptions", GoGetter: "ColumnDescriptions"},
			_jsii_.MemberProperty{JsiiProperty: "columnDescriptionsInput", GoGetter: "ColumnDescriptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDataset", GoMethod: "ResetBigqueryDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryTable", GoMethod: "ResetBigqueryTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumnDescriptions", GoMethod: "ResetColumnDescriptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesBigqueryResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesF1Resource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesF1Resource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "f1Table", GoGetter: "F1Table"},
			_jsii_.MemberProperty{JsiiProperty: "f1TableInput", GoGetter: "F1TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Table", GoMethod: "ResetF1Table"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesF1ResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObject", GoGetter: "GoogleCloudStorageObject"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageObjectInput", GoGetter: "GoogleCloudStorageObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesGoogleCloudStorageResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictions", GoGetter: "FileExtensionRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionRestrictionsInput", GoGetter: "FileExtensionRestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileReference", GoGetter: "FileReference"},
			_jsii_.MemberProperty{JsiiProperty: "fileReferenceInput", GoGetter: "FileReferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtensionRestrictions", GoMethod: "ResetFileExtensionRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileReference", GoMethod: "ResetFileReference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResource", GoGetter: "BigqueryResource"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryResourceInput", GoGetter: "BigqueryResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabel", GoGetter: "DisplayLabel"},
			_jsii_.MemberProperty{JsiiProperty: "displayLabelInput", GoGetter: "DisplayLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "f1Resource", GoGetter: "F1Resource"},
			_jsii_.MemberProperty{JsiiProperty: "f1ResourceInput", GoGetter: "F1ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResource", GoGetter: "GoogleCloudStorageResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleCloudStorageResourceInput", GoGetter: "GoogleCloudStorageResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResource", GoGetter: "GoogleDriveResource"},
			_jsii_.MemberProperty{JsiiProperty: "googleDriveResourceInput", GoGetter: "GoogleDriveResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescription", GoGetter: "ModelDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelDescriptionInput", GoGetter: "ModelDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryResource", GoMethod: "PutBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "putF1Resource", GoMethod: "PutF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleCloudStorageResource", GoMethod: "PutGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "putGoogleDriveResource", GoMethod: "PutGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "putRawFileResource", GoMethod: "PutRawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResource", GoGetter: "RawFileResource"},
			_jsii_.MemberProperty{JsiiProperty: "rawFileResourceInput", GoGetter: "RawFileResourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryResource", GoMethod: "ResetBigqueryResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayLabel", GoMethod: "ResetDisplayLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetF1Resource", GoMethod: "ResetF1Resource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleCloudStorageResource", GoMethod: "ResetGoogleCloudStorageResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDriveResource", GoMethod: "ResetGoogleDriveResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDescription", GoMethod: "ResetModelDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetRawFileResource", GoMethod: "ResetRawFileResource"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseRag", GoMethod: "ResetUseRag"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useRag", GoGetter: "UseRag"},
			_jsii_.MemberProperty{JsiiProperty: "useRagInput", GoGetter: "UseRagInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesRawFileResource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileContentInput", GoGetter: "FileContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitle", GoGetter: "FileTitle"},
			_jsii_.MemberProperty{JsiiProperty: "fileTitleInput", GoGetter: "FileTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mimeType", GoGetter: "MimeType"},
			_jsii_.MemberProperty{JsiiProperty: "mimeTypeInput", GoGetter: "MimeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaResourcesRawFileResourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkills",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkills)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkillsList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkillsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaSkillsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkillsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkillsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentInput", GoGetter: "ContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putReferences", GoMethod: "PutReferences"},
			_jsii_.MemberProperty{JsiiProperty: "references", GoGetter: "References"},
			_jsii_.MemberProperty{JsiiProperty: "referencesInput", GoGetter: "ReferencesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetReferences", GoMethod: "ResetReferences"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skillId", GoGetter: "SkillId"},
			_jsii_.MemberProperty{JsiiProperty: "skillIdInput", GoGetter: "SkillIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaSkillsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkillsReferences",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkillsReferences)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkillsReferencesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkillsReferencesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaSkillsReferencesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaSkillsReferencesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaSkillsReferencesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentInput", GoGetter: "ContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "referenceId", GoGetter: "ReferenceId"},
			_jsii_.MemberProperty{JsiiProperty: "referenceIdInput", GoGetter: "ReferenceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaSkillsReferencesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTables",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTablesColumns",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTablesColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTablesColumnsList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTablesColumnsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaTablesColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTablesColumnsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTablesColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaTablesColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTablesList",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTablesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTablesOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTablesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "columnsInput", GoGetter: "ColumnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putColumns", GoMethod: "PutColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumns", GoMethod: "ResetColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTimeouts",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-google.agenticApplicationsAnalystAgentPersona.AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference",
		reflect.TypeOf((*AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_AgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
