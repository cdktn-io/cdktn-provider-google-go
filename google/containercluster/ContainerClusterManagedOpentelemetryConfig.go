// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterManagedOpentelemetryConfig struct {
	// The scope of the Managed OpenTelemetry pipeline.
	//
	// Available options include NONE and COLLECTION_AND_INSTRUMENTATION_COMPONENTS. See https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#managedopentelemetryconfig for more information. To disable the feature, explicitly set this to NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

