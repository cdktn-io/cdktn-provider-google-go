// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicledashboardchart


type ChronicleDashboardChartDashboardChartVisualizationGoogleMapsConfigPointSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/chronicle_dashboard_chart#color ChronicleDashboardChart#color}.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Possible values: ["POINT_SIZE_TYPE_UNSPECIFIED", "POINT_SIZE_TYPE_FIXED", "POINT_SIZE_TYPE_PROPORTIONAL_TO_SIZE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/chronicle_dashboard_chart#point_size_type ChronicleDashboardChart#point_size_type}
	PointSizeType *string `field:"optional" json:"pointSizeType" yaml:"pointSizeType"`
}

