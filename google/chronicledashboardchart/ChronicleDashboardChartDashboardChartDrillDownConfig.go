// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicledashboardchart


type ChronicleDashboardChartDashboardChartDrillDownConfig struct {
	// left_drill_downs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/chronicle_dashboard_chart#left_drill_downs ChronicleDashboardChart#left_drill_downs}
	LeftDrillDowns interface{} `field:"optional" json:"leftDrillDowns" yaml:"leftDrillDowns"`
	// right_drill_downs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/chronicle_dashboard_chart#right_drill_downs ChronicleDashboardChart#right_drill_downs}
	RightDrillDowns interface{} `field:"optional" json:"rightDrillDowns" yaml:"rightDrillDowns"`
}

