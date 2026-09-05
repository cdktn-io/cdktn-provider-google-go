// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigquerydatapolicyv2datapolicy


type BigqueryDatapolicyv2DataPolicyDataGovernanceTag struct {
	// Tag keys are globally unique.
	//
	// Tag key is expected to be in the namespaced format, for example "parent-id/pii" where "parent-id" is the ID of the parent organization or project resource for this tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/bigquery_datapolicyv2_data_policy#key BigqueryDatapolicyv2DataPolicy#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag value is expected to be the short name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/bigquery_datapolicyv2_data_policy#value BigqueryDatapolicyv2DataPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

