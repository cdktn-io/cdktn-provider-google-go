// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesservicelbpolicies


type NetworkServicesServiceLbPoliciesAutoCapacityDrain struct {
	// Optional.
	//
	// If set to 'True', an unhealthy MIG/NEG will be set as drained. - An MIG/NEG is considered unhealthy if less than 25% of the instances/endpoints in the MIG/NEG are healthy. - This option will never result in draining more than 50% of the configured IGs/NEGs for the Backend Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/network_services_service_lb_policies#enable NetworkServicesServiceLbPolicies#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

