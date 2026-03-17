package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareObservatoryScheduledTestModel represents the source cloudflare_observatory_scheduled_test state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromV4 to parse legacy state.
//
// All fields are simple strings - no nested structures or complex types.
type SourceCloudflareObservatoryScheduledTestModel struct {
	ID        types.String `tfsdk:"id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	URL       types.String `tfsdk:"url"`
	Region    types.String `tfsdk:"region"`
	Frequency types.String `tfsdk:"frequency"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetObservatoryScheduledTestModel represents the target cloudflare_observatory_scheduled_test state structure (v500).
// This should match the model in the parent package's model.go file.
//
// Note: We only include fields that need to be migrated from v4.
// New computed-only fields (schedule, test) are not included here as they will be
// populated by the API on the first refresh after migration.
type TargetObservatoryScheduledTestModel struct {
	ID        types.String `tfsdk:"id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	URL       types.String `tfsdk:"url"`
	Region    types.String `tfsdk:"region"`
	Frequency types.String `tfsdk:"frequency"`
	// schedule and test are computed-only in v5, not migrated from v4
	Schedule customfield.NestedObject[TargetObservatoryScheduledTestScheduleModel] `tfsdk:"schedule"`
	Test     customfield.NestedObject[TargetObservatoryScheduledTestTestModel]     `tfsdk:"test"`
}

// TargetObservatoryScheduledTestScheduleModel represents the schedule nested object (v500).
// This is a computed-only field in v5, not present in v4.
type TargetObservatoryScheduledTestScheduleModel struct {
	Frequency types.String `tfsdk:"frequency"`
	Region    types.String `tfsdk:"region"`
	URL       types.String `tfsdk:"url"`
}

// TargetObservatoryScheduledTestTestModel represents the test nested object (v500).
// This is a computed-only field in v5, not present in v4.
type TargetObservatoryScheduledTestTestModel struct {
	ID                types.String                                                             `tfsdk:"id"`
	Date              timetypes.RFC3339                                                        `tfsdk:"date"`
	DesktopReport     customfield.NestedObject[TargetObservatoryScheduledTestTestDesktopReportModel] `tfsdk:"desktop_report"`
	MobileReport      customfield.NestedObject[TargetObservatoryScheduledTestTestMobileReportModel]  `tfsdk:"mobile_report"`
	Region            customfield.NestedObject[TargetObservatoryScheduledTestTestRegionModel]        `tfsdk:"region"`
	ScheduleFrequency types.String                                                             `tfsdk:"schedule_frequency"`
	URL               types.String                                                             `tfsdk:"url"`
}

// Nested report models (computed-only, for completeness)
type TargetObservatoryScheduledTestTestDesktopReportModel struct {
	CLS              types.Float64                                                                 `tfsdk:"cls"`
	DeviceType       types.String                                                                  `tfsdk:"device_type"`
	Error            customfield.NestedObject[TargetObservatoryScheduledTestTestDesktopReportErrorModel] `tfsdk:"error"`
	FCP              types.Float64                                                                 `tfsdk:"fcp"`
	JsonReportURL    types.String                                                                  `tfsdk:"json_report_url"`
	LCP              types.Float64                                                                 `tfsdk:"lcp"`
	PerformanceScore types.Float64                                                                 `tfsdk:"performance_score"`
	Si               types.Float64                                                                 `tfsdk:"si"`
	State            types.String                                                                  `tfsdk:"state"`
	TBT              types.Float64                                                                 `tfsdk:"tbt"`
	TTFB             types.Float64                                                                 `tfsdk:"ttfb"`
	TTI              types.Float64                                                                 `tfsdk:"tti"`
}

type TargetObservatoryScheduledTestTestDesktopReportErrorModel struct {
	Code              types.String `tfsdk:"code"`
	Detail            types.String `tfsdk:"detail"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url"`
}

type TargetObservatoryScheduledTestTestMobileReportModel struct {
	CLS              types.Float64                                                                `tfsdk:"cls"`
	DeviceType       types.String                                                                 `tfsdk:"device_type"`
	Error            customfield.NestedObject[TargetObservatoryScheduledTestTestMobileReportErrorModel] `tfsdk:"error"`
	FCP              types.Float64                                                                `tfsdk:"fcp"`
	JsonReportURL    types.String                                                                 `tfsdk:"json_report_url"`
	LCP              types.Float64                                                                `tfsdk:"lcp"`
	PerformanceScore types.Float64                                                                `tfsdk:"performance_score"`
	Si               types.Float64                                                                `tfsdk:"si"`
	State            types.String                                                                 `tfsdk:"state"`
	TBT              types.Float64                                                                `tfsdk:"tbt"`
	TTFB             types.Float64                                                                `tfsdk:"ttfb"`
	TTI              types.Float64                                                                `tfsdk:"tti"`
}

type TargetObservatoryScheduledTestTestMobileReportErrorModel struct {
	Code              types.String `tfsdk:"code"`
	Detail            types.String `tfsdk:"detail"`
	FinalDisplayedURL types.String `tfsdk:"final_displayed_url"`
}

type TargetObservatoryScheduledTestTestRegionModel struct {
	Label types.String `tfsdk:"label"`
	Value types.String `tfsdk:"value"`
}
