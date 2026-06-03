// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chroniclenativedashboard


type ChronicleNativeDashboardCharts struct {
	// chart_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/chronicle_native_dashboard#chart_layout ChronicleNativeDashboard#chart_layout}
	ChartLayout *ChronicleNativeDashboardChartsChartLayout `field:"optional" json:"chartLayout" yaml:"chartLayout"`
	// The resource name of the associated DashboardChart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/chronicle_native_dashboard#dashboard_chart ChronicleNativeDashboard#dashboard_chart}
	DashboardChart *string `field:"optional" json:"dashboardChart" yaml:"dashboardChart"`
	// List of dashboard filter IDs applied to this chart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/chronicle_native_dashboard#filters_ids ChronicleNativeDashboard#filters_ids}
	FiltersIds *[]*string `field:"optional" json:"filtersIds" yaml:"filtersIds"`
}

