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
	Frequency types.String `tfsdk:"frequency" json:"frequency"`
	Region    types.String `tfsdk:"region" json:"region"`
	URL       types.String `tfsdk:"url" json:"url"`
}

type ObservatoryScheduledTestTestModel struct {
	ID                types.String                                    `tfsdk:"id" json:"id"`
	Date              timetypes.RFC3339                               `tfsdk:"date" json:"date" format:"date-time"`
	DesktopReport     *ObservatoryScheduledTestTestDesktopReportModel `tfsdk:"desktop_report" json:"desktopReport"`
	MobileReport      *ObservatoryScheduledTestTestMobileReportModel  `tfsdk:"mobile_report" json:"mobileReport"`
	Region            *ObservatoryScheduledTestTestRegionModel        `tfsdk:"region" json:"region"`
	ScheduleFrequency types.String                                    `tfsdk:"schedule_frequency" json:"scheduleFrequency"`
	URL               types.String                                    `tfsdk:"url" json:"url"`
}

type ObservatoryScheduledTestTestDesktopReportModel struct {
	Cls              types.Float64                                        `tfsdk:"cls" json:"cls"`
	DeviceType       types.String                                         `tfsdk:"device_type" json:"deviceType"`
	Error            *ObservatoryScheduledTestTestDesktopReportErrorModel `tfsdk:"error" json:"error"`
	Fcp              types.Float64                                        `tfsdk:"fcp" json:"fcp"`
	JsonReportURL    types.String                                         `tfsdk:"json_report_url" json:"jsonReportUrl"`
	Lcp              types.Float64                                        `tfsdk:"lcp" json:"lcp"`
	PerformanceScore types.Float64                                        `tfsdk:"performance_score" json:"performanceScore"`
	Si               types.Float64                                        `tfsdk:"si" json:"si"`
	State            types.String                                         `tfsdk:"state" json:"state"`
	Tbt              types.Float64                                        `tfsdk:"tbt" json:"tbt"`
	Ttfb             types.Float64                                        `tfsdk:"ttfb" json:"ttfb"`
	Tti              types.Float64                                        `tfsdk:"tti" json:"tti"`
}

type ObservatoryScheduledTestTestDesktopReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code"`
	Detail            types.String `tfsdk:"detail" json:"detail"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl"`
}

type ObservatoryScheduledTestTestMobileReportModel struct {
	Cls              types.Float64                                       `tfsdk:"cls" json:"cls"`
	DeviceType       types.String                                        `tfsdk:"device_type" json:"deviceType"`
	Error            *ObservatoryScheduledTestTestMobileReportErrorModel `tfsdk:"error" json:"error"`
	Fcp              types.Float64                                       `tfsdk:"fcp" json:"fcp"`
	JsonReportURL    types.String                                        `tfsdk:"json_report_url" json:"jsonReportUrl"`
	Lcp              types.Float64                                       `tfsdk:"lcp" json:"lcp"`
	PerformanceScore types.Float64                                       `tfsdk:"performance_score" json:"performanceScore"`
	Si               types.Float64                                       `tfsdk:"si" json:"si"`
	State            types.String                                        `tfsdk:"state" json:"state"`
	Tbt              types.Float64                                       `tfsdk:"tbt" json:"tbt"`
	Ttfb             types.Float64                                       `tfsdk:"ttfb" json:"ttfb"`
	Tti              types.Float64                                       `tfsdk:"tti" json:"tti"`
}

type ObservatoryScheduledTestTestMobileReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code"`
	Detail            types.String `tfsdk:"detail" json:"detail"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl"`
}

type ObservatoryScheduledTestTestRegionModel struct {
	Label types.String `tfsdk:"label" json:"label"`
	Value types.String `tfsdk:"value" json:"value"`
}