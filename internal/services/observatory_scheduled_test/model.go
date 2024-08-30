// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultEnvelope struct {
	Result ObservatoryScheduledTestModel `json:"result"`
}

type ObservatoryScheduledTestModel struct {
	ID        types.String                                                    `tfsdk:"id" json:"-,computed"`
	URL       types.String                                                    `tfsdk:"url" path:"url"`
	ZoneID    types.String                                                    `tfsdk:"zone_id" path:"zone_id"`
	Frequency types.String                                                    `tfsdk:"frequency" json:"frequency,computed"`
	ItemCount types.Float64                                                   `tfsdk:"item_count" json:"count,computed"`
	Region    types.String                                                    `tfsdk:"region" json:"region,computed"`
	Schedule  customfield.NestedObject[ObservatoryScheduledTestScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	Test      customfield.NestedObject[ObservatoryScheduledTestTestModel]     `tfsdk:"test" json:"test,computed"`
}

type ObservatoryScheduledTestScheduleModel struct {
	Frequency types.String `tfsdk:"frequency" json:"frequency,computed_optional"`
	Region    types.String `tfsdk:"region" json:"region,computed_optional"`
	URL       types.String `tfsdk:"url" json:"url,computed_optional"`
}

type ObservatoryScheduledTestTestModel struct {
	ID                types.String                                                             `tfsdk:"id" json:"id,computed_optional"`
	Date              timetypes.RFC3339                                                        `tfsdk:"date" json:"date,computed_optional" format:"date-time"`
	DesktopReport     customfield.NestedObject[ObservatoryScheduledTestTestDesktopReportModel] `tfsdk:"desktop_report" json:"desktopReport,computed_optional"`
	MobileReport      customfield.NestedObject[ObservatoryScheduledTestTestMobileReportModel]  `tfsdk:"mobile_report" json:"mobileReport,computed_optional"`
	Region            customfield.NestedObject[ObservatoryScheduledTestTestRegionModel]        `tfsdk:"region" json:"region,computed_optional"`
	ScheduleFrequency types.String                                                             `tfsdk:"schedule_frequency" json:"scheduleFrequency,computed_optional"`
	URL               types.String                                                             `tfsdk:"url" json:"url,computed_optional"`
}

type ObservatoryScheduledTestTestDesktopReportModel struct {
	Cls              types.Float64                                                                 `tfsdk:"cls" json:"cls,computed_optional"`
	DeviceType       types.String                                                                  `tfsdk:"device_type" json:"deviceType,computed_optional"`
	Error            customfield.NestedObject[ObservatoryScheduledTestTestDesktopReportErrorModel] `tfsdk:"error" json:"error,computed_optional"`
	Fcp              types.Float64                                                                 `tfsdk:"fcp" json:"fcp,computed_optional"`
	JsonReportURL    types.String                                                                  `tfsdk:"json_report_url" json:"jsonReportUrl,computed_optional"`
	Lcp              types.Float64                                                                 `tfsdk:"lcp" json:"lcp,computed_optional"`
	PerformanceScore types.Float64                                                                 `tfsdk:"performance_score" json:"performanceScore,computed_optional"`
	Si               types.Float64                                                                 `tfsdk:"si" json:"si,computed_optional"`
	State            types.String                                                                  `tfsdk:"state" json:"state,computed_optional"`
	Tbt              types.Float64                                                                 `tfsdk:"tbt" json:"tbt,computed_optional"`
	Ttfb             types.Float64                                                                 `tfsdk:"ttfb" json:"ttfb,computed_optional"`
	Tti              types.Float64                                                                 `tfsdk:"tti" json:"tti,computed_optional"`
}

type ObservatoryScheduledTestTestDesktopReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code,computed_optional"`
	Detail            types.String `tfsdk:"detail" json:"detail,computed_optional"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl,computed_optional"`
}

type ObservatoryScheduledTestTestMobileReportModel struct {
	Cls              types.Float64                                                                `tfsdk:"cls" json:"cls,computed_optional"`
	DeviceType       types.String                                                                 `tfsdk:"device_type" json:"deviceType,computed_optional"`
	Error            customfield.NestedObject[ObservatoryScheduledTestTestMobileReportErrorModel] `tfsdk:"error" json:"error,computed_optional"`
	Fcp              types.Float64                                                                `tfsdk:"fcp" json:"fcp,computed_optional"`
	JsonReportURL    types.String                                                                 `tfsdk:"json_report_url" json:"jsonReportUrl,computed_optional"`
	Lcp              types.Float64                                                                `tfsdk:"lcp" json:"lcp,computed_optional"`
	PerformanceScore types.Float64                                                                `tfsdk:"performance_score" json:"performanceScore,computed_optional"`
	Si               types.Float64                                                                `tfsdk:"si" json:"si,computed_optional"`
	State            types.String                                                                 `tfsdk:"state" json:"state,computed_optional"`
	Tbt              types.Float64                                                                `tfsdk:"tbt" json:"tbt,computed_optional"`
	Ttfb             types.Float64                                                                `tfsdk:"ttfb" json:"ttfb,computed_optional"`
	Tti              types.Float64                                                                `tfsdk:"tti" json:"tti,computed_optional"`
}

type ObservatoryScheduledTestTestMobileReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code,computed_optional"`
	Detail            types.String `tfsdk:"detail" json:"detail,computed_optional"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl,computed_optional"`
}

type ObservatoryScheduledTestTestRegionModel struct {
	Label types.String `tfsdk:"label" json:"label,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}
