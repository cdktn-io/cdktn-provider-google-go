// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioncontentpolicy


type DataLossPreventionContentPolicyInspectConfigMinLikelihoodPerInfoType struct {
	// Only returns findings equal or above this threshold.
	//
	// See https://cloud.google.com/dlp/docs/likelihood for more info. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#min_likelihood DataLossPreventionContentPolicy#min_likelihood}
	MinLikelihood *string `field:"required" json:"minLikelihood" yaml:"minLikelihood"`
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/data_loss_prevention_content_policy#info_type DataLossPreventionContentPolicy#info_type}
	InfoType *DataLossPreventionContentPolicyInspectConfigMinLikelihoodPerInfoTypeInfoType `field:"optional" json:"infoType" yaml:"infoType"`
}

