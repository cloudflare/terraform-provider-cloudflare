# Changelog

## 5.1.0 (2025-02-06)

Full Changelog: [v5.0.0...v5.1.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.0.0...v5.1.0)

### Features

* **api_token_permission_groups:** define `get` operation for datasources  ([#5065](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5065)) ([1ee6c93](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1ee6c93ff29f68c72f587b987ac593db219aa5dc))
* **api:** api update ([#5064](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5064)) ([aa22c4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aa22c4de787803640df7807f7f9e5a9b2fadf01f))
* **botnet_feed:** add datasource for get operations ([#5067](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5067)) ([ba39465](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ba394658dc60b5daefc3481358d34de4dab55123))
* **build:** allow for building against private go repos ([#5056](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5056)) ([3fccdc9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3fccdc994bb8005a7e943f60d0d6de7f372acc65))
* **dns:** add terraform resources for `dns_settings` and `dns_settings_internal_view` ([#5060](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5060)) ([b75c1aa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b75c1aac00bfc4a4bcb573177e1ddca1876be835))
* various codegen changes ([7af724e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7af724e0861a49a6d15031fff5008a6366f158a1))


### Bug Fixes

* **grit:** make pattern names consistent ([0b2ba12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b2ba12bb261244c6e9bd24e4b6fa783f26bc0e7))
* **list_item:** fix schema mappings ([7944314](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7944314f0d2ca8a6aca78f64c17fc353c654be21))
* **pages_domain:** use `name` for anchor identifier ([#5053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5053)) ([169d4f5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/169d4f581d2980b9adcddf4f62af1424dafe47a3))
* update migration guide to use source, not stdlib ([9d208d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d208d6577f60940ec2bc1d6dcb181fcefbfdfac))
* use correct name for Grit patterns ([2f8d522](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2f8d5220bc1c76003035c5af885883367bf06867))


### Chores

* **internal:** codegen related update ([#5061](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5061)) ([1fa6f58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1fa6f5801761a6aa2baa4148e77600698b4b0407))


### Documentation

* clean out previously set schema_versions ([fba939d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fba939df831ff60e451c2f50310eb626687bcdb8))
* handle cloudflare_record data migration ([9eb450b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9eb450bb6f146f49a3f6c9a8c73369d747e164f3))
* update deprecation dates ([7a8b7d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a8b7d2424fe20a0c83e7a76b069d9b7c042f751))
* update list guidance for list_items ([2b72ee6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b72ee6628be7f50ae76756780ed97077a3ff68f))
* update page_rules migration guidance ([45e30b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45e30b1ed3de447d54223d5479ca52fd8609db42))
* update ruleset headers guidance ([e928470](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e928470f6299d2c7620543b17c59c8736b20e3a2))
