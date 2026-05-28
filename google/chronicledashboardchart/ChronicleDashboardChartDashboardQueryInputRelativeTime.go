// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicledashboardchart


type ChronicleDashboardChartDashboardQueryInputRelativeTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/chronicle_dashboard_chart#start_time_val ChronicleDashboardChart#start_time_val}.
	StartTimeVal *string `field:"required" json:"startTimeVal" yaml:"startTimeVal"`
	// The time unit for the relative range. Possible values: ["SECOND", "MINUTE", "HOUR", "DAY", "WEEK", "MONTH", "YEAR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/chronicle_dashboard_chart#time_unit ChronicleDashboardChart#time_unit}
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
}

