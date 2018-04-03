## 0.1.1 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:

* resource/cloudflare_record: Changing `name` or `domain` now force a recreation
  of the record [GH-29]

FEATURES:

* **New Resource:** `cloudflare_rate_limit` [GH-30]
* **New Resource:** `cloudflare_page_rule` [GH-38]
* **New Resource:** `cloudflare_zone_settings_override` [GH-41]
* **New Resource:** `cloudflare_load_balancer_monitor` [GH-42]
* **New Data Source:** `cloudflare_ip_ranges` [GH-28]

IMPROVEMENTS:

* resource/cloudflare_record: Validate `TXT` records [GH-14]
* resource/cloudflare_record: Add `data` input to suppport SRV, LOC records
  [GH-29]
* resource/cloudflare_record: Add computed attributes `created_on`,
  `modified_on`, `proxiable`, and `metadata` to records [GH-29]
* resource/cloudflare_record: Support import of existing records [GH-36]

## 0.1.0 (June 20, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider
  Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
