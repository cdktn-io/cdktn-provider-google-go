// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package discoveryengineprojectprovision


type DiscoveryEngineProjectProvisionSaasParams struct {
	// Set to 'true' to specify that the caller has read and consents to the Terms for business edition quality of service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#accept_biz_qos DiscoveryEngineProjectProvision#accept_biz_qos}
	AcceptBizQos interface{} `field:"optional" json:"acceptBizQos" yaml:"acceptBizQos"`
	// Indicates whether the current request is for Business edition ('true') or not ('false').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#is_biz DiscoveryEngineProjectProvision#is_biz}
	IsBiz interface{} `field:"optional" json:"isBiz" yaml:"isBiz"`
}

