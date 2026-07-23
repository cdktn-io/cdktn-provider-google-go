// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigMinLikelihoodPerInfoType struct {
	// Only returns findings equal or above this threshold.
	//
	// See https://cloud.google.com/dlp/docs/likelihood for more info. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/data_loss_prevention_inspect_template#min_likelihood DataLossPreventionInspectTemplate#min_likelihood}
	MinLikelihood *string `field:"required" json:"minLikelihood" yaml:"minLikelihood"`
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/data_loss_prevention_inspect_template#info_type DataLossPreventionInspectTemplate#info_type}
	InfoType *DataLossPreventionInspectTemplateInspectConfigMinLikelihoodPerInfoTypeInfoType `field:"optional" json:"infoType" yaml:"infoType"`
}

