// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	// cloud_storage_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#cloud_storage_path DataLossPreventionContentPolicy#cloud_storage_path}
	CloudStoragePath *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `field:"optional" json:"cloudStoragePath" yaml:"cloudStoragePath"`
	// word_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#word_list DataLossPreventionContentPolicy#word_list}
	WordList *DataLossPreventionContentPolicyInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct `field:"optional" json:"wordList" yaml:"wordList"`
}

