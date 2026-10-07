// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyLoggingConfigs struct {
	// log_to_big_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#log_to_big_query DataLossPreventionContentPolicy#log_to_big_query}
	LogToBigQuery *DataLossPreventionContentPolicyLoggingConfigsLogToBigQuery `field:"optional" json:"logToBigQuery" yaml:"logToBigQuery"`
}

