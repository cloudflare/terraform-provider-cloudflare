---
page_title: "cloudflare_observatory_scheduled_test Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_observatory_scheduled_test (Resource)



## Example Usage

```terraform
resource "cloudflare_observatory_scheduled_test" "example" {
  zone_id   = "0da42c8d2132a9ddaf714f9e7c920711"
  url       = "example.com"
  region    = "us-central1"
  frequency = "WEEKLY"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `url` (String) A URL.
- `zone_id` (String) Identifier

### Read-Only

- `frequency` (String) The frequency of the test.
- `id` (String) A URL.
- `item_count` (Number) Number of items affected.
- `region` (String) A test region.
- `schedule` (Attributes) The test schedule. (see [below for nested schema](#nestedatt--schedule))
- `test` (Attributes) (see [below for nested schema](#nestedatt--test))

<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Optional:

- `frequency` (String) The frequency of the test.
- `region` (String) A test region.
- `url` (String) A URL.


<a id="nestedatt--test"></a>
### Nested Schema for `test`

Optional:

- `date` (String)
- `desktop_report` (Attributes) The Lighthouse report. (see [below for nested schema](#nestedatt--test--desktop_report))
- `id` (String) UUID
- `mobile_report` (Attributes) The Lighthouse report. (see [below for nested schema](#nestedatt--test--mobile_report))
- `region` (Attributes) A test region with a label. (see [below for nested schema](#nestedatt--test--region))
- `schedule_frequency` (String) The frequency of the test.
- `url` (String) A URL.

<a id="nestedatt--test--desktop_report"></a>
### Nested Schema for `test.desktop_report`

Optional:

- `cls` (Number) Cumulative Layout Shift.
- `device_type` (String) The type of device.
- `error` (Attributes) (see [below for nested schema](#nestedatt--test--desktop_report--error))
- `fcp` (Number) First Contentful Paint.
- `json_report_url` (String) The URL to the full Lighthouse JSON report.
- `lcp` (Number) Largest Contentful Paint.
- `performance_score` (Number) The Lighthouse performance score.
- `si` (Number) Speed Index.
- `state` (String) The state of the Lighthouse report.
- `tbt` (Number) Total Blocking Time.
- `ttfb` (Number) Time To First Byte.
- `tti` (Number) Time To Interactive.

<a id="nestedatt--test--desktop_report--error"></a>
### Nested Schema for `test.desktop_report.error`

Optional:

- `code` (String) The error code of the Lighthouse result.
- `detail` (String) Detailed error message.
- `final_displayed_url` (String) The final URL displayed to the user.



<a id="nestedatt--test--mobile_report"></a>
### Nested Schema for `test.mobile_report`

Optional:

- `cls` (Number) Cumulative Layout Shift.
- `device_type` (String) The type of device.
- `error` (Attributes) (see [below for nested schema](#nestedatt--test--mobile_report--error))
- `fcp` (Number) First Contentful Paint.
- `json_report_url` (String) The URL to the full Lighthouse JSON report.
- `lcp` (Number) Largest Contentful Paint.
- `performance_score` (Number) The Lighthouse performance score.
- `si` (Number) Speed Index.
- `state` (String) The state of the Lighthouse report.
- `tbt` (Number) Total Blocking Time.
- `ttfb` (Number) Time To First Byte.
- `tti` (Number) Time To Interactive.

<a id="nestedatt--test--mobile_report--error"></a>
### Nested Schema for `test.mobile_report.error`

Optional:

- `code` (String) The error code of the Lighthouse result.
- `detail` (String) Detailed error message.
- `final_displayed_url` (String) The final URL displayed to the user.



<a id="nestedatt--test--region"></a>
### Nested Schema for `test.region`

Optional:

- `label` (String)
- `value` (String) A test region.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_observatory_scheduled_test.example <zone_id>:<url>:<region>
```