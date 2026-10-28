// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2instance


type CloudRunV2InstanceContainersStartupProbeGrpc struct {
	// Optional.
	//
	// Port number of the gRPC service. Number must be in the range 1 to 65535. If not specified, defaults to the exposed port of the container, which is the value of container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#port CloudRunV2Instance#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Optional.
	//
	// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md ). If this is not specified, the default behavior is defined by gRPC
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/cloud_run_v2_instance#service CloudRunV2Instance#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

