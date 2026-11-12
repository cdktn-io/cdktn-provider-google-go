// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionVpcAccess struct {
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/app_engine_standard_app_version#network_interfaces AppEngineStandardAppVersion#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The traffic egress setting for the VPC network interface, controlling what traffic is diverted through it.
	//
	// Possible values: ["ALL_TRAFFIC", "PRIVATE_IP_RANGES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/app_engine_standard_app_version#vpc_egress AppEngineStandardAppVersion#vpc_egress}
	VpcEgress *string `field:"optional" json:"vpcEgress" yaml:"vpcEgress"`
}

