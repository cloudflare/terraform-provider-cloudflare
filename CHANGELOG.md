# Changelog

## 5.1.0 (2025-02-10)

Full Changelog: [v5.0.0...v5.1.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.0.0...v5.1.0)

### Features

* **api_token_permission_groups:** define `get` operation for datasources  ([#5065](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5065)) ([1ee6c93](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1ee6c93ff29f68c72f587b987ac593db219aa5dc))
* **api:** api update ([#5064](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5064)) ([aa22c4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aa22c4de787803640df7807f7f9e5a9b2fadf01f))
* **botnet_feed:** add datasource for get operations ([#5067](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5067)) ([ba39465](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ba394658dc60b5daefc3481358d34de4dab55123))
* **build:** allow for building against private go repos ([#5056](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5056)) ([3fccdc9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3fccdc994bb8005a7e943f60d0d6de7f372acc65))
* **dns:** add terraform resources for `dns_settings` and `dns_settings_internal_view` ([#5060](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5060)) ([b75c1aa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b75c1aac00bfc4a4bcb573177e1ddca1876be835))
* **r2_bucket_cors:** add resource ([#5082](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5082)) ([03a4c40](https://github.com/cloudflare/terraform-provider-cloudflare/commit/03a4c40324c20fd8fd0d7bb7f6f409dcfc39019b))
* **r2_bucket_event_notification:** add resource and flatten method hierarchy ([#5084](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5084)) ([8448659](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8448659d8b17bd378c6daafff571de369237732f))
* **r2_bucket_lifecycle:** add resource ([#5085](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5085)) ([5935439](https://github.com/cloudflare/terraform-provider-cloudflare/commit/59354394b483f28255c8aa7719d0a3c47ad5cb4c))
* **r2_bucket_lock:** add resource ([#5086](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5086)) ([48a4783](https://github.com/cloudflare/terraform-provider-cloudflare/commit/48a478362652cabf6b5967b9fa2ca6f5e9353743))
* **r2_bucket_sippy:** add resource ([#5087](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5087)) ([a00ca25](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a00ca2598c91ba99d4988f8dd2ad7c420e3a6046))
* various codegen changes ([7af724e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7af724e0861a49a6d15031fff5008a6366f158a1))


### Bug Fixes

* **grit:** make pattern names consistent ([0b2ba12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b2ba12bb261244c6e9bd24e4b6fa783f26bc0e7))
* **list_item:** fix schema mappings ([7944314](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7944314f0d2ca8a6aca78f64c17fc353c654be21))
* **pages_domain:** use `name` for anchor identifier ([#5053](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5053)) ([169d4f5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/169d4f581d2980b9adcddf4f62af1424dafe47a3))
* **r2_bucket_event_notification:** dedupe `bucket_name` parameters ([1ebd9a3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1ebd9a32fc5ec04b6994436f8a0faec87a068d15))
* **spectrum_application:** `id` is computed, not required ([2fd2378](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2fd23780cbc201f13761aca3ce9d7ed95cdcbf1c))
* update migration guide to use source, not stdlib ([9d208d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d208d6577f60940ec2bc1d6dcb181fcefbfdfac))
* use correct name for Grit patterns ([2f8d522](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2f8d5220bc1c76003035c5af885883367bf06867))


### Chores

* **internal:** codegen related update ([#5061](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5061)) ([1fa6f58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1fa6f5801761a6aa2baa4148e77600698b4b0407))
* minor change to tests ([#5077](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5077)) ([fcfaec5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fcfaec542597b9beefee370e908bffae11311e57))


### Documentation

* clean out previously set schema_versions ([fba939d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fba939df831ff60e451c2f50310eb626687bcdb8))
* generate ([9b6e1ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9b6e1ecff2578ef99209374a0bfe5dc759d42a5c))
* generate r2 resources ([d72c8c8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d72c8c8d5f63dea25a0bd6be44d637880efc028a))
* handle cloudflare_record data migration ([9eb450b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9eb450bb6f146f49a3f6c9a8c73369d747e164f3))
* update deprecation dates ([7a8b7d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a8b7d2424fe20a0c83e7a76b069d9b7c042f751))
* update list guidance for list_items ([2b72ee6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b72ee6628be7f50ae76756780ed97077a3ff68f))
* update page_rules migration guidance ([45e30b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45e30b1ed3de447d54223d5479ca52fd8609db42))
* update ruleset headers guidance ([e928470](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e928470f6299d2c7620543b17c59c8736b20e3a2))
* **workers_script:** fix example syntax ([d86363d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d86363d58c848ced0d19eaadbe9e4419bec1870d))
