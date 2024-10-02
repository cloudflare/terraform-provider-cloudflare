// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ObservatoryScheduledTestResultEnvelope struct {
	Result ObservatoryScheduledTestModel `json:"result"`
}

type ObservatoryScheduledTestModel struct {
	ID        types.String                                                    `tfsdk:"id" json:"-,computed"`
	URL       types.String                                                    `tfsdk:"url" path:"url,required"`
	ZoneID    types.String                                                    `tfsdk:"zone_id" path:"zone_id,required"`
	Frequency types.String                                                    `tfsdk:"frequency" json:"frequency,computed"`
	ItemCount types.Float64                                                   `tfsdk:"item_count" json:"count,computed"`
	Region    types.String                                                    `tfsdk:"region" json:"region,computed"`
	Schedule  customfield.NestedObject[ObservatoryScheduledTestScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
	Test      customfield.NestedObject[ObservatoryScheduledTestTestModel]     `tfsdk:"test" json:"test,computed"`
}

func (m ObservatoryScheduledTestModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ObservatoryScheduledTestModel) MarshalJSONForUpdate(state ObservatoryScheduledTestModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ObservatoryScheduledTestScheduleModel struct {
	Frequency types.String `tfsdk:"frequency" json:"frequency,computed"`
	Region    types.String `tfsdk:"region" json:"region,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
}

type ObservatoryScheduledTestTestModel struct {
	ID                types.String                                                             `tfsdk:"id" json:"id,computed"`
	Date              timetypes.RFC3339                                                        `tfsdk:"date" json:"date,computed" format:"date-time"`
	DesktopReport     customfield.NestedObject[ObservatoryScheduledTestTestDesktopReportModel] `tfsdk:"desktop_report" json:"desktopReport,computed"`
	MobileReport      customfield.NestedObject[ObservatoryScheduledTestTestMobileReportModel]  `tfsdk:"mobile_report" json:"mobileReport,computed"`
	Region            customfield.NestedObject[ObservatoryScheduledTestTestRegionModel]        `tfsdk:"region" json:"region,computed"`
	ScheduleFrequency types.String                                                             `tfsdk:"schedule_frequency" json:"scheduleFrequency,computed"`
	URL               types.String                                                             `tfsdk:"url" json:"url,computed"`
}

type ObservatoryScheduledTestTestDesktopReportModel struct {
	Cls              types.Float64                                                                 `tfsdk:"cls" json:"cls,computed"`
	DeviceType       types.String                                                                  `tfsdk:"device_type" json:"deviceType,computed"`
	Error            customfield.NestedObject[ObservatoryScheduledTestTestDesktopReportErrorModel] `tfsdk:"error" json:"error,computed"`
	Fcp              types.Float64                                                                 `tfsdk:"fcp" json:"fcp,computed"`
	JsonReportURL    types.String                                                                  `tfsdk:"json_report_url" json:"jsonReportUrl,computed"`
	Lcp              types.Float64                                                                 `tfsdk:"lcp" json:"lcp,computed"`
	PerformanceScore types.Float64                                                                 `tfsdk:"performance_score" json:"performanceScore,computed"`
	Si               types.Float64                                                                 `tfsdk:"si" json:"si,computed"`
	State            types.String                                                                  `tfsdk:"state" json:"state,computed"`
	Tbt              types.Float64                                                                 `tfsdk:"tbt" json:"tbt,computed"`
	Ttfb             types.Float64                                                                 `tfsdk:"ttfb" json:"ttfb,computed"`
	Tti              types.Float64                                                                 `tfsdk:"tti" json:"tti,computed"`
}

type ObservatoryScheduledTestTestDesktopReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code,computed"`
	Detail            types.String `tfsdk:"detail" json:"detail,computed"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl,computed"`
}

type ObservatoryScheduledTestTestMobileReportModel struct {
	Cls              types.Float64                                                                `tfsdk:"cls" json:"cls,computed"`
	DeviceType       types.String                                                                 `tfsdk:"device_type" json:"deviceType,computed"`
	Error            customfield.NestedObject[ObservatoryScheduledTestTestMobileReportErrorModel] `tfsdk:"error" json:"error,computed"`
	Fcp              types.Float64                                                                `tfsdk:"fcp" json:"fcp,computed"`
	JsonReportURL    types.String                                                                 `tfsdk:"json_report_url" json:"jsonReportUrl,computed"`
	Lcp              types.Float64                                                                `tfsdk:"lcp" json:"lcp,computed"`
	PerformanceScore types.Float64                                                                `tfsdk:"performance_score" json:"performanceScore,computed"`
	Si               types.Float64                                                                `tfsdk:"si" json:"si,computed"`
	State            types.String                                                                 `tfsdk:"state" json:"state,computed"`
	Tbt              types.Float64                                                                `tfsdk:"tbt" json:"tbt,computed"`
	Ttfb             types.Float64                                                                `tfsdk:"ttfb" json:"ttfb,computed"`
	Tti              types.Float64                                                                `tfsdk:"tti" json:"tti,computed"`
}

type ObservatoryScheduledTestTestMobileReportErrorModel struct {
	Code              types.String `tfsdk:"code" json:"code,computed"`
	Detail            types.String `tfsdk:"detail" json:"detail,computed"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url" json:"finalDisplayedUrl,computed"`
}

type ObservatoryScheduledTestTestRegionModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
