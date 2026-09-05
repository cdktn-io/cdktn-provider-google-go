// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package developerconnectinsightsconfig


type DeveloperConnectInsightsConfigTargetProjects struct {
	// The project IDs. Format {project}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/developer_connect_insights_config#project_ids DeveloperConnectInsightsConfig#project_ids}
	ProjectIds *[]*string `field:"optional" json:"projectIds" yaml:"projectIds"`
}

