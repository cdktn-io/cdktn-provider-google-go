// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesAction struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_grpc_route#destinations NetworkServicesGrpcRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_grpc_route#fault_injection_policy NetworkServicesGrpcRoute#fault_injection_policy}
	FaultInjectionPolicy *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// Specifies the idle timeout for the selected route.
	//
	// The idle timeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection. If not set, the default idle timeout is 1 hour. If set to 0s, the timeout will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_grpc_route#idle_timeout NetworkServicesGrpcRoute#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_grpc_route#retry_policy NetworkServicesGrpcRoute#retry_policy}
	RetryPolicy *NetworkServicesGrpcRouteRulesActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Specifies the timeout for selected route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/network_services_grpc_route#timeout NetworkServicesGrpcRoute#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

