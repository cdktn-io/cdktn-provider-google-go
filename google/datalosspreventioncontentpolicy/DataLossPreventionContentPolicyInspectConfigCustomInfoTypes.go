// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigCustomInfoTypes struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_type DataLossPreventionContentPolicy#info_type}
	InfoType *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// detection_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#detection_rules DataLossPreventionContentPolicy#detection_rules}
	DetectionRules interface{} `field:"optional" json:"detectionRules" yaml:"detectionRules"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#dictionary DataLossPreventionContentPolicy#dictionary}
	Dictionary *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned.
	//
	// It still can be used for rules matching. Possible values: ["EXCLUSION_TYPE_EXCLUDE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#exclusion_type DataLossPreventionContentPolicy#exclusion_type}
	ExclusionType *string `field:"optional" json:"exclusionType" yaml:"exclusionType"`
	// file_label_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#file_label_info_type DataLossPreventionContentPolicy#file_label_info_type}
	FileLabelInfoType *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType `field:"optional" json:"fileLabelInfoType" yaml:"fileLabelInfoType"`
	// Likelihood to return for this CustomInfoType.
	//
	// This base value can be altered by a detection rule if the finding meets the criteria
	// specified by the rule. Default value: "VERY_LIKELY" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#likelihood DataLossPreventionContentPolicy#likelihood}
	Likelihood *string `field:"optional" json:"likelihood" yaml:"likelihood"`
	// metadata_key_value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#metadata_key_value_expression DataLossPreventionContentPolicy#metadata_key_value_expression}
	MetadataKeyValueExpression *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesMetadataKeyValueExpression `field:"optional" json:"metadataKeyValueExpression" yaml:"metadataKeyValueExpression"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#regex DataLossPreventionContentPolicy#regex}
	Regex *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesRegex `field:"optional" json:"regex" yaml:"regex"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#sensitivity_score DataLossPreventionContentPolicy#sensitivity_score}
	SensitivityScore *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// stored_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#stored_type DataLossPreventionContentPolicy#stored_type}
	StoredType *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType `field:"optional" json:"storedType" yaml:"storedType"`
	// surrogate_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#surrogate_type DataLossPreventionContentPolicy#surrogate_type}
	SurrogateType *DataLossPreventionContentPolicyInspectConfigCustomInfoTypesSurrogateType `field:"optional" json:"surrogateType" yaml:"surrogateType"`
}

