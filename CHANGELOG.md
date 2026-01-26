# Changelog

## 5.17.0 (2026-01-26)

Full Changelog: [v5.16.0...v5.17.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.16.0...v5.17.0)

### Features

* chore: use 'next' branch of Go SDK in Terraform ([809a3f3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/809a3f35a77e6215a25a10ee19cae72b0fee089e))
* **leaked_credential_check:** add import functionality. add tests for import ([76e44f0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/76e44f06b3ed6ebe99d57fb43e80bd62eaf22e92))
* refactor(terraform): restructure origin_tls_client_auth to peer subresources ([6c12fea](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6c12fead7f3fa947ce4c8bd2a488bb0b001b6cd3))
* **turstile_widget:** add v4 to v5 migration tests ([a1e27af](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a1e27afe229e2812ab7b5e570eb78066b134d6dd))


### Bug Fixes

* prevent unnecessary diffs on consecutive applies for hyperdrive_config ([8755bf9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8755bf9d36e1b994821a3e7a4893845083023f48))
* **zero_trust_access_application:** update v4 version on migration tests ([45a825e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45a825ef3cb0f9f26ce5f5cbd2e343344f277a76))


### Reverts

* **pages_project:** "fix(pages_project) build_config to computed optional" ([b9c13c9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b9c13c9dfb341d9fd7ff49c47c217730afec9abd))


### Chores

* add CODEOWNERS ([3abbb08](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3abbb0810b16d8607653d7be8e1cacd372f758f1))
* improve contribution guide ([85584b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/85584b7ab508fb75cf5ee9213aa87b8a7f18e86f))
* **internal:** codegen related update ([0211418](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0211418984d901b0617c0842e92a99aa1a727f7f))
* **internal:** codegen related update ([2bcbbd5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2bcbbd5d2f2aa17267a10042e864ad7fca0b41ed))
* **internal:** codegen related update ([09f9d99](https://github.com/cloudflare/terraform-provider-cloudflare/commit/09f9d998f1fb0ec5aa4c3f866168bcbf367cda0b))
* Update CHANGELOG.md ([f4a1b58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f4a1b588006049af00ad051c35c1b9453649a0e8))

## 5.16.0 (2026-01-20)

Full Changelog: [v5.15.0...v5.16.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.15.0...v5.16.0)

### Features

* **custom_pages:** add "waf_challenge" as new supported error page type identifier in both resource and data source schemas
* **list:** enhance CIDR validator to check for normalized CIDR notation requiring network address for IPv4 and IPv6
* **magic_wan_gre_tunnel:** add automatic_return_routing attribute for automatic routing control
* **magic_wan_gre_tunnel:** add BGP configuration support with new BGP model attribute
* **magic_wan_gre_tunnel:** add bgp_status computed attribute for BGP connection status information
* **magic_wan_gre_tunnel:** enhance schema with BGP-related attributes and validators
* **magic_wan_ipsec_tunnel:** add automatic_return_routing attribute for automatic routing control
* **magic_wan_ipsec_tunnel:** add BGP configuration support with new BGP model attribute
* **magic_wan_ipsec_tunnel:** add bgp_status computed attribute for BGP connection status information
* **magic_wan_ipsec_tunnel:** add custom_remote_identities attribute for custom identity configuration
* **magic_wan_ipsec_tunnel:** enhance schema with BGP and identity-related attributes
* **ruleset:** add request body buffering support
* **ruleset:** enhance ruleset data source with additional configuration options
* **workers_script:** add observability logs attributes to list data source model
* **workers_script:** enhance list data source schema with additional configuration options

### Bug Fixes

* **dns_record:** remove unnecessary fmt.Sprintf wrapper around LoadTestCase call in test configuration helper function
* **load_balancer:** fix session_affinity_ttl type expectations to match Float64 in initial creation and Int64 after migration
* **workers_kv:** handle special characters correctly in URL encoding

### Documentation

* **account_subscription:** update schema description for rate_plan.sets attribute to clarify it returns an array of strings
* **api_shield:** add resource-level description for API Shield management of auth ID characteristics
* **api_shield:** enhance auth_id_characteristics.name attribute description to include JWT token configuration format requirements
* **api_shield:** specify JSONPath expression format for JWT claim locations
* **hyperdrive_config:** add description attribute to name attribute explaining its purpose in dashboard and API identification
* **hyperdrive_config:** apply description improvements across resource, data source, and list data source schemas
* **hyperdrive_config:** improve schema descriptions for cache settings to clarify default values
* **hyperdrive_config:** update port description to clarify defaults for different database types

## 5.15.0 (2025-12-19)

Full Changelog: [v5.14.0...v5.15.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.14.0...v5.15.0)

### Features

* **ai_search:** add AI Search endpoints ([6f02adb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f02adb420e872457f71f95b49cb527663388915))
* **certificate_pack:** add terraform config for CRUD support. This ensures proper Terraform resource ID handling for path parameters in API calls. ([081f32a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/081f32acab4ce9a194a7ff51c8e9fcabd349895a))
* **leaked_credentials_check:** Add GET endpoint for leaked_credentials_check/detections ([feafd9c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/feafd9c466ec90a2874f2cd6b3316b41f52fd37a))
* **worker_version:** support `startup_time_ms` ([286ab55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/286ab55bea8d5be0faa5a2b5b8b157e4a2214eba))
* **zero_trust_access_group:** v4 to v5 migration acceptance tests ([9c877c7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c877c7f60c8e58cc3f32539e650f1b908a4e628))
* **zero_trust_access_mtls_hostname_settings:** use v2 migrator ([b14aa6d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b14aa6df7598aaf56c7261c1eb4a8e4c2f1d08ab))
* **zero_trust_dlp_custom_entry:** support `upload_status` ([7dc0fe3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0fe3b23726ead8dc075f86728a0540846d90c))
* **zero_trust_dlp_entry:** support `upload_status` ([7dc0fe3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0fe3b23726ead8dc075f86728a0540846d90c))
* **zero_trust_dlp_integration_entry:** support `upload_status` ([7dc0fe3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0fe3b23726ead8dc075f86728a0540846d90c))
* **zero_trust_dlp_predefined_entry:** support `upload_status` ([7dc0fe3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0fe3b23726ead8dc075f86728a0540846d90c))
* **zero_trust_gateway_policy:** support `forensic_copy` ([5741fd0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5741fd0ed9f7270d20731cc47ec45eb0403a628b))
* **zero_trust_list:** support additional types (category, location, device) ([5741fd0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5741fd0ed9f7270d20731cc47ec45eb0403a628b))

### Bug Fixes

* **access_rules:** Add validation to prevent state drift. Ideally we'd use Semantic Equality but since that isn't an option, this will remove a foot-gun. ([4457791](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44577911b3cbe45de6279aefa657bdee73c0794d))
* **cloudflare_pages_project:** addressing drift issues ([6edffcf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6edffcfcf187fdc9b10b624b9a9b90aed2fb2b2e)) ([3db318e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3db318e747423bf10ce587d9149e90edcd8a77b0))
* **cloudflare_worker:** can be cleanly imported ([4859b52](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4859b52968bb25570b680df9813f8e07fd50728f))
* **cloudflare_worker:** ensure clean imports ([5b525bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b525bc478a4e2c9c0d4fd659b92cc7f7c18016a))
* **list_items:** Add validation for IP List items to avoid inconsistent state ([b6733dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b6733dc4be909a5ab35895a88e519fc2582ccada))
* **zero_trust_access_application:** remove all conditions from sweeper ([3197f1a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3197f1aed61be326d507d9e9e3b795b9f1d18fd7))
* map missing fields during spectrum resource import ([#6495](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6495)) ([ddb4e72](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ddb4e722b82c735825a549d651a9da219c142efa))
* update invalid codegen ([d365b98](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d365b9859fddf385220c1e716e8c226651d28905)) ([92f5c9e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/92f5c9e07afec5e2f31a7299fa84b73938530796))

### Chores

* **certificate_pack:** `hosts` is now a set, not a list ([286ab55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/286ab55bea8d5be0faa5a2b5b8b157e4a2214eba))
* **ci:** split acceptance tests into 37 parallel groups for faster ([8c6212b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8c6212b4c4694b9b7ce625e77ffbc2cf35725708))
* **healthcheck:** add test for expected_body default value ([c5afb48](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c5afb48dfd8e7faeea6d5e60aec4e7e75a1a9b6b)) ([e99b43f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e99b43fb877cfc5b80872e8ad20b8d6cf85827c2))
* **magic_wan_ipsec_tunnel:** remove `custom_remote_entities` ([286ab55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/286ab55bea8d5be0faa5a2b5b8b157e4a2214eba))
* **queue_consumer:** Test data fixes for queue consumer acceptance tests ([1b92700](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b92700491da5d189435f9ea37c899970d303dc9))
* update go to point to next ([25d640a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/25d640a4d4b60b23504fae2ceb3250a432dde8af))
* update regional hostnames migration test to use new migrator ([d5ac65f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d5ac65f0965a896508ba2f5ceb2ba87efe3bb049))
* update test to use new migrator ([d5ac65f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d5ac65f0965a896508ba2f5ceb2ba87efe3bb049)) ([ec875bb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ec875bb95701ce155ed64a19c7d5b8ccb4f56fd6))

### Documentation

* Deprecate API Shield Schema Validation resources ([366e1b8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/366e1b8cd631ff8e1b7fc1230def2c13d0aea680))
* generate provider docs ([c23342e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c23342ed882222c0067d24508c45b89e5c258931))


### Refactors

* **healthcheck:** consolidate tests and expand update coverage ([b747d21](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b747d21c01d1d8694e0bfac507c07e27ba22c239)) ([7fa38b3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7fa38b34fc27c856d01a8607237a896648d0b13c))

## 5.14.0 (2025-12-06)

Full Changelog: [v5.13.0...v5.14.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.13.0...v5.14.0)

### Features

* add v4-&gt;v5 migration tests for pages_project and adjust schema ([#6506](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6506)) ([6de0179](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6de0179a033543e4e63053b9db68185f4e4f2c78))
* chore: point Terraform to Go 'next' ([af9a5f8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/af9a5f896f4e65e9808a2b6458279b9a7ff935fe))
* chore: update go sdk to v6.4.0 for provider release ([63cb021](https://github.com/cloudflare/terraform-provider-cloudflare/commit/63cb021ec8318c4aefaacadc025c90d8ef3e618d))
* chore(api_shield_discovery_operation): Deprecate api_shield_discovery_operation ([7dc0a63](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0a63e0fa3de74d8877788ff787294cb603c07))
* feat: BOTS-7562 add bot management feedback endpoints to stainless config (prod) ([f5112e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5112e1e4d4fe53b9e0ec96aa27f69fb7706099b))
* feat(r2_data_catalog): Configure SDKs/Terraform to use R2 Data Catalog routes ([5beb50b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5beb50b30d7d012afae92f8f652e005955d4e430))
* feat(radar): Add origins endpoints to public api docs ([ee51e5f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee51e5f19c5f1f3cd4d2c07acdcf0f368c635c4d))
* improve and standardize sweepers ([#6501](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6501)) ([03fb2d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/03fb2d2f4999ca24b9597093f25fd1dbc0f671b7))


### Bug Fixes

* **account_members:** making member policies a set ([#6488](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6488)) ([f3ecaa5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f3ecaa5938486865698b3956848a8e5f0f6c9054))
* decoder and tests ([#6516](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6516)) ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))
* decoder, build ([#6514](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6514)) ([1935459](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1935459b88e1ec7d656148d92aee6ea45557ce3c))
* **pages_project:** non empty refresh plans ([#6515](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6515)) ([bc526ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bc526ffcf5091195cc143cbf774013b84728296c))
* **pages_project:** use correct field name in test sweeper ([6dc0e53](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6dc0e538754b84e8e31d0c1b7bd8c0e291161811))
* r2 sweeper ([#6512](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6512)) ([fec953c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fec953c98e8764e5b1ed3dc1906d02206718928e))
* **tests:** resolve SDK v6 migration test failures ([#6507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6507)) ([bad9716](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bad971609bfb7ddc8d78de6935b747f350f1ae55))
* update import signature to accept account_id/subscription_id in order to import account subscription ([#6510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6510)) ([c2db582](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c2db582ed27a6e8ac3de6461dba011496b687d05))
* **utils:** test assertions ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))
* **workers_kv:** ignore value import state verify ([#6521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6521)) ([c3e3f89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c3e3f892bcd0d1426eb398688b1f5f8c84144b57))
* **workers_script:** No longer treating the migrations attribute as WriteOnly ([#6489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6489)) ([dc60e38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc60e3883002db7eb9036265aa4000a08a1eb2b6))
* **workers_script:** resource drift when worker has unmanaged secret ([#6504](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6504)) ([505c0fe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/505c0fe78b4f096ed3547f7b564ea5788a64a644))
* **zero_trust_device_posture_rule:** preserve input.version and other fields ([#6500](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6500)) ([4c4e54b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c4e54bdfdfa2c6ce584322cd2ae8562269b4a98))
* **zero_trust_device_posture_rule:** preserve input.version and other fields ([#6503](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6503)) ([d45be9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d45be9a3b6276bee2c9f4413e141db6d5d1fa596))
* **zero_trust_dlp_custom_profile:** add sweepers for dlp_custom_profile ([cbcb325](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cbcb325271e91d6c3a0bb832eb0c347198d9ce6c))
* **zone_subscription|account_subscription:** add partners_ent as valid enum for rate_plan.id ([#6505](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6505)) ([2a70fb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2a70fb49ccc28e5309f7096d1ebc34866c2f07f3))
* **zone:** datasource model schema parity ([#6487](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6487)) ([861c68f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/861c68fcb37627371673e60dfcbc7bf09638e28e))


### Chores

* **account_member:** dont run acceptance with env variable ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))
* **account_member:** fix check for env var ([#6517](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6517)) ([07e9aa5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07e9aa5ef6327499731b17ac334f2f9d4cf3f4bf))
* **account_member:** skip until user is dsr enabled ([#6522](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6522)) ([dd7c2fe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dd7c2fea55e1dc10953b57c2d3792b2e97a8e68e))
* **account_tokens:** adding a simple CRUD test ([#6484](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6484)) ([6869538](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68695382141a2b097c051b38a32fabf5368c2ad9))
* **api:** update composite API spec ([71fc050](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71fc0502fe85faaabd599348b8f46930fa0bd15f))
* **api:** update composite API spec ([68d017a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68d017a1dd2379d7f4c22b83b9ee31e146971b54))
* **cloudflare_api_shield_operation:** Add acceptance tests  ([#6491](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6491)) ([37e1fdf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37e1fdf9d72b9d853e4ba7b751f2680db5abb9f6))
* **docs:** update documentation ([#6523](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6523)) ([a060e61](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a060e61e9882b706ad6386bc2d1ffc8ac295166d))
* **internal:** codegen related update ([923ea1d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/923ea1db2a8accfd846990947a49f44a1e209ea4))
* **internal:** codegen related update ([a110cbe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a110cbe76756173711fbc50c6c13c143d53ae625))
* **internal:** codegen related update ([7b36c06](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7b36c06f10f0fee20c5935b20f6817f7f91d2442))
* **internal:** codegen related update ([c789f91](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c789f91d6ba5c773ad152945114ab6b020949616))
* **internal:** codegen related update ([a91faa6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a91faa61eca4c2cf4ea23aff7f14667e9b042de7))
* **internal:** codegen related update ([dfa745a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dfa745abb71b94428ff1140f592a48cb7eca8969))
* **internal:** codegen related update ([fb3ef1b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb3ef1b5b7fe827753ba82f4f08381bf951f5dbf))
* **internal:** codegen related update ([504e8b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/504e8b61e19091d8f2e66f37889311256cf6590b))
* **internal:** codegen related update ([550d77c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/550d77c15980feac28e1df2b6fc52473d81f57e6))
* **logpush_job:** add v4 to v5 migration tests ([#6483](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6483)) ([2e4b8a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e4b8a090ca9741cd6f7e5861305f82ac582b63d))
* **tests:** cloud connector rules parity tests and add connectivity_directory_service tests ([#6513](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6513)) ([5341c82](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5341c824679002c0168d3ce8696022aec30ff33e))
* update changelog ([#6480](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6480)) ([adba156](https://github.com/cloudflare/terraform-provider-cloudflare/commit/adba156599c44b43277c7d4b5694f3ccde2408a3))
* update changelog ([#6525](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6525)) ([b026b4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b026b4d32a8c394823d91be0a4ed156252ff5fcd))
* **zero_trust_device_default_profile_local_domain_fallback:** add tests ([#6464](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6464)) ([365cb71](https://github.com/cloudflare/terraform-provider-cloudflare/commit/365cb71b906fca75c2be9fedf38f5269eaa0f4dd))
* **zero_trust_device_default|custom_profile:** acceptance test coverage ([#6511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6511)) ([8e4ec1a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8e4ec1a9e376ac294707db9fc320a8855be0ba89))
* **zero_trust_device_managed_networks:** add tests ([#6463](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6463)) ([e9b6783](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e9b67835cbf10fb76282b2526f732f13aa13ca06))
* **zero_trust_device_posture_integration:** update tests for to test with Crowdstrike ([#6470](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6470)) ([e360d6f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e360d6f414799181c76c9eb7b6aaeb51239a2632))
* **zone:** update migration tests ([#6468](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6468)) ([8ff53df](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ff53dfe672acc0dc3b9538a613de77e96517e0e))

## 5.14.0 (2025-12-05)

Full Changelog: [v5.13.0...v5.14.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.13.0...v5.14.0)

### Features

* add v4-&gt;v5 migration tests for pages_project and adjust schema ([#6506](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6506)) ([6de0179](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6de0179a033543e4e63053b9db68185f4e4f2c78))
* chore(api_shield_discovery_operation): Deprecate api_shield_discovery_operation ([7dc0a63](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0a63e0fa3de74d8877788ff787294cb603c07))
* feat(bot_management): add bot management feedback endpoints to stainless config (prod) ([f5112e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5112e1e4d4fe53b9e0ec96aa27f69fb7706099b))
* feat(r2_data_catalog): Configure SDKs/Terraform to use R2 Data Catalog routes ([5beb50b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5beb50b30d7d012afae92f8f652e005955d4e430))
* feat(radar): Add origins endpoints to public api docs ([ee51e5f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee51e5f19c5f1f3cd4d2c07acdcf0f368c635c4d))
* improve and standardize sweepers ([#6501](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6501)) ([03fb2d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/03fb2d2f4999ca24b9597093f25fd1dbc0f671b7))


### Bug Fixes

* **account_members:** making member policies a set ([#6488](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6488)) ([f3ecaa5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f3ecaa5938486865698b3956848a8e5f0f6c9054))
* **pages_project:** non empty refresh plans ([#6515](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6515)) ([bc526ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bc526ffcf5091195cc143cbf774013b84728296c))
* **pages_project:** use correct field name in test sweeper ([6dc0e53](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6dc0e538754b84e8e31d0c1b7bd8c0e291161811))
* **r2** sweeper ([#6512](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6512)) ([fec953c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fec953c98e8764e5b1ed3dc1906d02206718928e))
* **tests:** resolve SDK v6 migration test failures ([#6507](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6507)) ([bad9716](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bad971609bfb7ddc8d78de6935b747f350f1ae55))
* **utils:** test assertions ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))
* **workers_kv:** ignore value import state verify ([#6521](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6521)) ([c3e3f89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c3e3f892bcd0d1426eb398688b1f5f8c84144b57))
* **workers_script:** No longer treating the migrations attribute as WriteOnly ([#6489](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6489)) ([dc60e38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc60e3883002db7eb9036265aa4000a08a1eb2b6))
* **workers_script:** resource drift when worker has unmanaged secret ([#6504](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6504)) ([505c0fe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/505c0fe78b4f096ed3547f7b564ea5788a64a644))
* **zero_trust_device_posture_rule:** preserve input.version and other fields ([#6500](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6500)) ([4c4e54b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c4e54bdfdfa2c6ce584322cd2ae8562269b4a98))
* **zero_trust_device_posture_rule:** preserve input.version and other fields ([#6503](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6503)) ([d45be9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d45be9a3b6276bee2c9f4413e141db6d5d1fa596))
* **zero_trust_dlp_custom_profile:** add sweepers for dlp_custom_profile ([cbcb325](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cbcb325271e91d6c3a0bb832eb0c347198d9ce6c))
* **zone_subscription|account_subscription:** add partners_ent as valid enum for rate_plan.id ([#6505](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6505)) ([2a70fb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2a70fb49ccc28e5309f7096d1ebc34866c2f07f3))
* **zone:** datasource model schema parity ([#6487](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6487)) ([861c68f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/861c68fcb37627371673e60dfcbc7bf09638e28e))
* update import signature to accept account_id/subscription_id in order to import account subscription ([#6510](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6510)) ([c2db582](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c2db582ed27a6e8ac3de6461dba011496b687d05))
* json decoder and tests ([#6516](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6516)) ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))


### Chores
* **api_shield_discovery_operation:** Deprecate api_shield_discovery_operation ([7dc0a63](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc0a63e0fa3de74d8877788ff787294cb603c07))
* **account_member:** dont run acceptance with env variable ([4c3e2db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c3e2db3ec31638d423d0ecc971e5c5ea54298ec))
* **account_member:** fix check for env var ([#6517](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6517)) ([07e9aa5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07e9aa5ef6327499731b17ac334f2f9d4cf3f4bf))
* **account_member:** skip until user is dsr enabled ([#6522](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6522)) ([dd7c2fe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dd7c2fea55e1dc10953b57c2d3792b2e97a8e68e))
* **account_tokens:** adding a simple CRUD test ([#6484](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6484)) ([6869538](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68695382141a2b097c051b38a32fabf5368c2ad9))
* **zero_trust_dlp_entry:** add `profiles` attribute to DLP entry resources ([71fc050](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71fc0502fe85faaabd599348b8f46930fa0bd15f))
* **r2_bucket_sippy:** add S3-compatible provider support with `bucket_url` attribute ([71fc050](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71fc0502fe85faaabd599348b8f46930fa0bd15f))
* **pages_project:** add `source` config block for granular source control settings ([68d017a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68d017a1dd2379d7f4c22b83b9ee31e146971b54))
* **certificate_pack:** add `certificates` and `primary_certificate` attributes ([68d017a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68d017a1dd2379d7f4c22b83b9ee31e146971b54))
* **zero_trust_device_posture_rule:** add `antivirus` posture rule type ([68d017a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68d017a1dd2379d7f4c22b83b9ee31e146971b54))
* **zero_trust_device_settings:** add external emergency signal configuration attributes ([68d017a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68d017a1dd2379d7f4c22b83b9ee31e146971b54))
* **cloudflare_api_shield_operation:** add acceptance tests ([#6491](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6491)) ([37e1fdf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37e1fdf9d72b9d853e4ba7b751f2680db5abb9f6))
* **logpush_job:** add v4 to v5 migration tests ([#6483](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6483)) ([2e4b8a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e4b8a090ca9741cd6f7e5861305f82ac582b63d))
* **tests:** cloud connector rules parity tests and add connectivity_directory_service tests ([#6513](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6513)) ([5341c82](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5341c824679002c0168d3ce8696022aec30ff33e))
* **zero_trust_device_default_profile_local_domain_fallback:** add tests ([#6464](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6464)) ([365cb71](https://github.com/cloudflare/terraform-provider-cloudflare/commit/365cb71b906fca75c2be9fedf38f5269eaa0f4dd))
* **zero_trust_device_default|custom_profile:** acceptance test coverage ([#6511](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6511)) ([8e4ec1a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8e4ec1a9e376ac294707db9fc320a8855be0ba89))
* **zero_trust_device_managed_networks:** add tests ([#6463](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6463)) ([e9b6783](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e9b67835cbf10fb76282b2526f732f13aa13ca06))
* **zero_trust_device_posture_integration:** update tests for to test with Crowdstrike ([#6470](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6470)) ([e360d6f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e360d6f414799181c76c9eb7b6aaeb51239a2632))
* **zone:** update migration tests ([#6468](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6468)) ([8ff53df](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ff53dfe672acc0dc3b9538a613de77e96517e0e))
* **internal:** update cloudflare-go SDK dependency ([923ea1d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/923ea1db2a8accfd846990947a49f44a1e209ea4))
* **internal:** update cloudflare-go SDK dependency ([a110cbe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a110cbe76756173711fbc50c6c13c143d53ae625))
* **internal:** update cloudflare-go SDK dependency ([7b36c06](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7b36c06f10f0fee20c5935b20f6817f7f91d2442))
* **internal:** improve JSON decoder to handle multiple fields with same name ([c789f91](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c789f91d6ba5c773ad152945114ab6b020949616))
* **internal:** update cloudflare-go SDK dependency ([a91faa6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a91faa61eca4c2cf4ea23aff7f14667e9b042de7))
* **internal:** update cloudflare-go SDK dependency ([dfa745a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dfa745abb71b94428ff1140f592a48cb7eca8969))
* **internal:** add example files for email security resources ([fb3ef1b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb3ef1b5b7fe827753ba82f4f08381bf951f5dbf))
* **internal:** update cloudflare-go SDK dependency ([504e8b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/504e8b61e19091d8f2e66f37889311256cf6590b))
* **internal:** update cloudflare-go SDK dependency ([550d77c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/550d77c15980feac28e1df2b6fc52473d81f57e6))

### Deprecations
* **api_shield_discovery_operation**

## 5.13.0 (2025-11-21)

Full Changelog: [v5.12.0...v5.13.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.12.0...v5.13.0)

### ⚠ BREAKING CHANGES: cloudflare_api_token & cloudflare_account_token Schema Update

The 5.13 release includes major updates to the cloudflare_api_token resource to eliminate configuration drift caused by policy ordering differences in the Cloudflare API.

Fixes: cloudflare/terraform-provider-cloudflare#6092

**Whats changed**
- policies are now a Set; order is ignored to prevent drift.
- When defining a policy, resources must use jsonencode(); all policy resource values must now be JSON-encoded strings.
- Removed fields: id, name, and meta have been removed from policy blocks.

**Required Action (v5.13+)**
Customers looking to upgrade to v5.13+ must update all cloudflare_api_token & cloudflare_account_token resources to wrap policy resource values in jsonencode()

Before:
```
resources = {
  "com.cloudflare.api.account.${var.cf_account_id}" = "*"
}
```
After:
```
resources = jsonencode({
  "com.cloudflare.api.account.${var.cf_account_id}" = "*"
})
```

* **account_token:** token policy order and nested resources ([#6440](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6440))

### Features

* add new resources and data sources ([7ce3dec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7ce3dec8fc5b65116750b8bf8209c2ec612d6a61))
* **api_token+account_tokens:** state upgrader and schema bump ([#6472](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6472)) ([42f7db2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/42f7db27659337230aa03094d050c8ebbcbdc24c))
* chore(build): point Terraform to released Go v6.3.0 ([6d06b46](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6d06b462491797af17f086193bbf32ccdffdd4b5))
* **docs:** make docs explicit when a resource does not have import support ([02699f6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/02699f65c082555c54b84288f23eda2272708144))
* **magic_transit_connector:** support self-serve license key ([#6398](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6398)) ([a6ec134](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a6ec1340765d2d9e980ded2b66dce847c142523f))
* **worker_version:** add content_base64 support ([6ff643f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6ff643fda6f0171a07fcd0070fc0e4716f1b1563))
* **worker_version:** boolean support for run_worker_first ([#6407](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6407)) ([116a67b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/116a67bdfaf481200152380a627cd1de8397b1c9))
* **workers_script_subdomains:** add import support  ([#6375](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6375)) ([40f7ed8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/40f7ed8b34adfa42c4dad22ce6e2b0c90d40c8c0))
* **zero_trust_access_application:** add proxy_endpoint for ZT Access Application ([#6453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6453)) ([177f20a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/177f20a46ce5f36d8f1eef70893e76ddb4a3ef05))
* **zero_trust_dlp_predefined_profile:** Switch DLP Predefined Profile endpoints, introduce enabled_entries attribute ([bc69569](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bc695692c86c5af9cfa4a13ce2c6ac5bd38a3538))
* **zero_trust_tunnel_cloudflared:** v4 to v5 migration tests ([#6461](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6461)) ([ffa0fef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ffa0fef80e7175346c6b48f92c3cb0ea89be1d37))


### Bug Fixes

* **account_token:** token policy order and nested resources ([#6440](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6440)) ([86c5972](https://github.com/cloudflare/terraform-provider-cloudflare/commit/86c5972edc65d6190f8fc52f5da2d07a99d0bef0))
* allow r2_bucket_event_notification to be applied twice without failing ([#6419](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6419)) ([6fbd4c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6fbd4c5aeb0f89b935e770dc3fa5c1d89661894f))
* **cloudflare_worker+cloudflare_worker_version:** import for the resources ([#6357](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6357)) ([b98e0be](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b98e0be8d6bbd4f04afbb61c98c36b0ecfa0bea4))
* **dns_record:** inconsistent apply error ([#6452](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6452)) ([f289994](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f289994d58720ec58fc186534a1a5e82776624bc))
* **pages_domain:** resource tests ([#6338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6338)) ([d769e29](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d769e2930016efa73c8e0ac2b4b620a107d03f7d))
* **pages_project:** unintended resource state drift ([#6377](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6377)) ([1a3955a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a3955ad49a4a3d74bb0d5faf08acb0f77d4921b))
* **queue_consumer:** id population ([#6181](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6181)) ([f3c6498](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f3c6498d16e402044160ad38993de188061405fc))
* **workers_kv:** multipart request  ([#6367](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6367)) ([65f8c19](https://github.com/cloudflare/terraform-provider-cloudflare/commit/65f8c19a2269f19d88f4f6edd14c4980ad53c9ac))
* **workers_kv:** updating workers metadata attribute to be read from endpoint ([#6386](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6386)) ([3a35757](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3a35757dff9b6dfdadbd422d36d877e0eec63926))
* **workers_script_subdomain:** add note to cloudflare_workers_script_subdomain about redundancy with cloudflare_worker ([#6383](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6383)) ([9cc9b59](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9cc9b59cb8b79ce8a0cb3988b353e45cd7be07ec))
* **workers_script:** allow config.run_worker_first to accept list input ([fab567c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fab567cc3191feecf1c19d6e4d91125c08dc6121))
* **zero_trust_device_custom_profile_local_domain_fallback:** drift issues ([#6365](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6365)) ([65c0c18](https://github.com/cloudflare/terraform-provider-cloudflare/commit/65c0c1895587b6f61404a04d748ea2ffd5317442))
* **zero_trust_device_custom_profile:** resolve drift issues ([#6364](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6364)) ([4cd2cbd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4cd2cbdd93dbd622cf7f2a29d56f4bf01896a0a5))
* **zero_trust_dex_test:** correct configurability for 'targeted' attribute to fix drift ([cd81178](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cd81178f30e800af3345822d5eee478419d6cd14))
* **zero_trust_tunnel_cloudflared_config:** remove warp_routing from cloudflared_config ([#6471](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6471)) ([dc9d557](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc9d557149289f0bc33d28bb7e31f54dd42e1c82))


### Chores

* **account_member:** add migration test ([#6425](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6425)) ([967a972](https://github.com/cloudflare/terraform-provider-cloudflare/commit/967a9727cd7d0b49e5c3ae1f6a6acee66a925186))
* **byoip:** integrate generated changes for BYOIP resources ([432160e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/432160ef04c30bb13072a2eec84231c197432e69))
* **certificate_pack:** docs show safe rotation instructions ([#6388](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6388)) ([3d37264](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3d3726408c21561acdd1f908a33f2178660ab489))
* **ci:** clean up leftover files in resources ([#6474](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6474)) ([e8aee72](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8aee72c5b9f4a966d1da283af2f3cc941be8ad7))
* **ci:** drop migration tests from CI ([#6476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6476)) ([968565f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/968565ffa098127bf03c354ea148222c6aa4438b))
* **ci:** fix tests ran on release PR ([#6478](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6478)) ([0b43c46](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b43c464a1178d97e083445b2031f3a3f6d178ee))
* **ci:** fixes for parity tests and build failures ([#6475](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6475)) ([3561876](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3561876f17b34f79c16b1a36bab3b2e3129bdeca))
* **ci:** modify sweepers ([#6479](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6479)) ([4c8915d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c8915d6202277724beaf13417eeaae519ad2070))
* **ci:** skip flaky test in CI ([fb14d86](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb14d86b0354e9717caeed87c8b749625fb09f86))
* **cloudflare_zero_trust_dlp_custom_profile:** migration test and ignore order as set ([#6428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6428)) ([1659ff3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1659ff3ee3fa9fd90bc9bba674a2d16927a4e5fe))
* **d1:** integrate generated changes for D1 resources ([cfa3472](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cfa347232730294a359da2eb6187899d56e973ce))
* **dns_record:** improve dns sweepers ([#6430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6430)) ([5e62468](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5e62468963235dfce1cc4d8a87e35063d5203197))
* **docs:** document configurations and examples ([#6449](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6449)) ([59430e0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/59430e0b7bd2a4371e9a817ddf7105690859b40d))
* **docs:** generate docs and examples ([cdd77ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cdd77eca036fcdb6b7ae2ad27cbbc851c5eca95c))
* **email_routing:** improved email routing sweepers ([#6429](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6429)) ([133c81e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/133c81e0b3880071ebec216d558b348676e3b301))
* **iam:** integrate generated changes for IAM resources ([a87806e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a87806ed569c98ae6b301cec30641ffb492b9317))
* include new sections for pr template ([#6395](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6395)) ([81c07e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/81c07e12fc2b2aafbd919af71b2154c16083bf3e))
* **load_balancing:** integrate generated changes for Load Balancing resources ([4c6b34d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c6b34de2bc8a1d191a49d9321dcc6eced60c3a8))
* **logpull_retention:** add migration test for ([#6426](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6426)) ([529f313](https://github.com/cloudflare/terraform-provider-cloudflare/commit/529f31392cc783af1091604900b0611d7385a731))
* **logpull_retention:** update acceptance test ([#6277](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6277)) ([3766b3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3766b3f3346f3f89d23c6613dad98dd7a8a5ed13))
* **logpush_job:** add import tests for resource ([#6402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6402)) ([cded8ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cded8ece37d62b4443ce7da1267d964ed42b7215))
* **logpush:** integrate generated changes for Logpush resources ([06e8446](https://github.com/cloudflare/terraform-provider-cloudflare/commit/06e8446a2c4c253efe4e4687e237edc4158c3392))
* **notification_policy_webhook:** add migration test for notification-policy-webhook ([#6443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6443)) ([742d647](https://github.com/cloudflare/terraform-provider-cloudflare/commit/742d64789205f1ac9d177c57b9e97b0ddf6a5a45))
* **pages:** integrate generated changes for Pages resources ([64855ea](https://github.com/cloudflare/terraform-provider-cloudflare/commit/64855ea4cf8a9396def48d97ceb30bbc0b36b62d))
* **queue_consumer:** testdata refactor ([d301974](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3019745c45fc43ecb1a74f500f846dbce2fce08))
* **r2_bucket:** v4 to v5 migration tests for cloudflare_r2_bucket ([#6437](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6437)) ([99ed1ee](https://github.com/cloudflare/terraform-provider-cloudflare/commit/99ed1ee0f3e8cb761f5e6e712f42ee87bf109039))
* **sso_connector:** add acceptance tests ([#6427](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6427)) ([8b54303](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8b54303138b0a96a2023d15a3f9da59492ebfbae))
* **stainless:** integrate changes from unpinned codegen version ([9cb3b8e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9cb3b8eb7dc6334d8d2ac808cc6adeb02129ca8a))
* **test:** acceptance tests for token validation resources ([#6417](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6417)) ([4d94bdd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4d94bddb8f952072a1b4d5fdfd08ac6a5cc01457))
* **test:** add schema and token validation acceptance tests to CI ([#6421](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6421)) ([b805abc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b805abcf0b6b134bd51c5218d2c22e99d8d28a37))
* **test:** increase legacy migrator test coverage ([#6401](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6401)) ([9a8c48a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a8c48a29df316d319ac84b2b7aa561181e513b2))
* **universal_ssl_setting:** add acceptance tests for universal_ssl_setting ([2601c45](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2601c4542dc6e354552d1b1e2ff2052d5757eea4))
* **worker:** integrate generated changes for Worker resources ([1da2bf2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1da2bf2b663cd086539b9edbde127a970e6b60cd))
* **workers_kv_namespace:** v4 to v5 migration tests for workers_kv_namespace ([#6424](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6424)) ([433010f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/433010f2dba43c4ab909126aff351332220b4907))
* **workers_kv:** v4 to v5 migration tests for workers_kv ([#6435](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6435)) ([58ca912](https://github.com/cloudflare/terraform-provider-cloudflare/commit/58ca912c521729cc9ff0453f589279ec6da9b7c6))
* **workers_script:** add workers scripts sweeper ([#6351](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6351)) ([f439a08](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f439a086e8120b84193e70fa6e426fedb9895b79))
* **workers_script:** fix resource name in TestAccCloudflareWorkerScript_ModuleWithDurableObject ([614d8d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/614d8d3765984e424d5628ad4fd2356bbe422746))
* **workers_script:** fix resource names in tests ([788e73a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/788e73a5a2a1ef43a02ee75ceb1b7da3a05e5ce8))
* **workers:** integrate generated changes for Workers resources ([ab0a330](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ab0a3303f4268783c8b78dc2fda0d1517afc2d16))
* **zero_trust_access_service_token:** add migration test for zero_trust_access_service_token ([#6416](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6416)) ([c77d5d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c77d5d5d2eafc852db0468c13fb880f9a4127e28))
* **zero_trust_gateway_policy:** v4 to v5 migration for zero_trust_gateway_policy ([#6413](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6413)) ([1c1952b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1c1952b61300569bfbfcb731abfe817bcce33fd9))
* **zero_trust_list:** v4 to v5 migration tests for zero trust list records ([#6400](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6400)) ([6ed55d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6ed55d6787ca377c42ddef59fe700b46139cf262))
* **zero_trust_tunnel_cloudflared_route:** v4 to v5 migration tests for zero_trust_tunnel_cloudflared_route ([#6409](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6409)) ([5dc2094](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5dc20940902b3f72594f15b91a7f2f1088dfee94))
* **zero_trust, cfone:** integrate generated changes for ZT and CFONE resources ([b7131b2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b7131b2be2a9fd36b05d71cb4d05182d4b044fa2))
* **zone_dnssec:** v4 to v5 migration tests for zone_dnssec ([#6432](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6432)) ([86abd1f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/86abd1f906b03547e04ad66d185d052461d82251))
* **zone_settings:** acceptance test to repro issue [#6363](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6363) ([#6445](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6445)) ([707c154](https://github.com/cloudflare/terraform-provider-cloudflare/commit/707c1542f7e97c800bef0dfdd0170a7f0594ea33))
* **zones:** data source tests ([#6414](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6414)) ([4d58e56](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4d58e5631cd10b685c7b0b63230ef4a0d6b18a6f))
* **zt_access:** add sweepers for policy and service token ([#6465](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6465)) ([9f4fa94](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9f4fa949610bf27ae4b179cd28232e26be7610b6))

## 5.12.0 (2025-10-30)

Full Changelog: [v5.11.0...v5.12.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.11.0...v5.12.0)

### Features

* chore: pin cloudflare-go for provider release ([61a33f9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/61a33f92db2c788262c166e259b6477aa06cdbb1))
* chore: use cloudflare-go@next for the 'next' branch ([8d8ff6d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8d8ff6d20993ab3127b3fcdbb2f17a93835ca70d))
* chore(abuse): rename path parameter ([cbda07b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cbda07b16992d4007ab38f27ec0ac7ca54bde7a4))
* ci: remove zero_trust_connectivity_directory_service ([23bd535](https://github.com/cloudflare/terraform-provider-cloudflare/commit/23bd5354dfccdc00fd5e38e877e65406abc6a7be))
* ci: trigger prod build ([fffdf5a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fffdf5af95668db4c444856d3fe1b0b3e5b32bc4))
* feat: add connectivity directory service APIs to openapi.stainless.yml ([1a6b304](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a6b304bdf7e0c442b5b651c749fe8f924f17bc4))
* feat: SDKs for Organizations and OrganizationsProfile ([1f6eae3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1f6eae3469f83e99c02cc744462e2c7e0236fed3))
* feat(api): add mcp portals endpoints ([1e317de](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e317dedc0b09f66aea05652cc8c72b2980ecfd0))
* feat(radar): add new group by dimension endpoints; deprecate to_markdown endpoint ([bcb58cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bcb58cb92cbae33ffca77c7646955c5f3ea9e47d))
* fix(content_scanning): content scanning terraform resource ([03b7004](https://github.com/cloudflare/terraform-provider-cloudflare/commit/03b7004e74ba50c2b968133ba5ef37e07781bd58))
* fix(workers_domain): treat `PUT /workers/domains` as a create operation ([8ff0c7d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ff0c7df085a37d62e5e86b76f262a462a60b9fc))
* modernize zero_trust_tunnel_cloudflared_config tests and fix warp_routing ([#6294](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6294)) ([36d38a6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/36d38a602dd53d537c8c1f25ee18910e39bdc36c))
* modernize zero_trust_tunnel_cloudflared_virtual_network tests and improve ([#6293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6293)) ([1b0f6d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b0f6d6c3bd2d22b7ac40a006a20ded3105ffe46))
* **zero_trust_access_application:** Add support for MCP & MCP_PORTAL ([#6326](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6326)) ([9524b60](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9524b60fe59daf86654df29d1fb545eb01b22be3))


### Bug Fixes

* **account_member:** update policies test by selecting correct resource group ([#6352](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6352)) ([693dc9d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/693dc9d47daefaa66cb168cc955a2f7d682d43e0))
* **cloudflare_r2_bucket_sippy:** attribute name in example ([#6336](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6336)) ([208bf81](https://github.com/cloudflare/terraform-provider-cloudflare/commit/208bf815dd926c3304ef18b7a7aaeb607e0b0adc))
* **cloudflare_worker_version:** replace when module content_sha256 value changes ([#6335](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6335)) ([e31395d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e31395df73343dd51ab36673fb4bd69845890b42))
* **cloudflare_workers_script:** Update docs note for resource ([#6304](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6304)) ([f7b4cef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f7b4cef7a44b5514567661aedb71b1185454d7f7))
* **cloudflare_workflow:** download dependencies for workflow resource acceptance tests ([#6302](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6302)) ([84bade9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/84bade9cb649961e9e732ad04147581cb60cddb6))
* correctly detect more ID attributes for data sources ([d5f4e7d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d5f4e7d6eec6e710e24482f7241a8e1ccfc3a836))
* **custom_pages:** fix broken tests ([#6372](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6372)) ([95f344e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/95f344edeca21a7cc25c54649c0bff395a8377c5))
* **custom_pages:** update type enumerations ([#6369](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6369)) ([8bd0d09](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8bd0d091d26aa0b8122f4c595b35977215361d95))
* enable skipped gateway policy tests and simplify quarantine test ([#6296](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6296)) ([b220f2b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b220f2b369693b44fcb3b42ead3ba9e9682c5295))
* ensure model/schema parity across several resources ([#6379](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6379)) ([418aedd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/418aedd21e9cf5d68346b3592a3e2706526d0111))
* fix zero_trust_dex_test tests ([#6301](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6301)) ([0345a4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0345a4d1ba765e87083ddc807e2d913ec1d54d80))
* **internal:** correctly generate schema according to annotations ([529f0ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/529f0ff681b8048401af9d66399584a29ae5fe34))
* **migrate:** add target flag to specify resources ([#6324](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6324)) ([1b94fcd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b94fcd43ac7f6351db5a961d91c49133b8e64ae))
* **notification_policy:** address drift due to unordered lists, converted to sets ([#6316](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6316)) ([7eabe67](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7eabe67756ea6d3239983712473c2fce7e60fef1))
* read by id data sources should have required IDs ([1ca9485](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1ca94850ae165054d84f1e913a091b30f1a3199e))
* restore missing testdata ([#6378](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6378)) ([5cb8dc6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5cb8dc6c807eba0eafc2727d917d423c1e208d6b))
* **workers_version:** inconsistent binding order causing inconsistent result after apply ([#6342](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6342)) ([1de79a4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1de79a4749ed56ba8f8611d4c02ba51eb0779479))
* **zero_trust_access_service_token:** client secret versioning ([#6328](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6328)) ([d6b7107](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d6b71074f87354cd7c3f336ead970f12b324480f))
* **zero_trust_dex_test:** ensure model/schema parity ([#6370](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6370)) ([066ae4f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/066ae4f6b74cfaab4a7c92701a1542fdb436afc7))
* **zero_trust_dex_test:** fix duplicate key, imports ([#6366](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6366)) ([15c05d0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/15c05d05fb6264ef700265b28813bd935f26643b))
* **zero_trust_dlp_custom_profile:** fix read, refresh, import ([#6391](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6391)) ([3154453](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3154453269561d9fb523795bb20fa90296ef74fe))
* **zero_trust_tunnel_cloudflared_virtual_network:** fix sweeper panics ([#6392](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6392)) ([c190bc7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c190bc7de6bc3c6b4a8f01e8e9ca789910143044))


### Chores

* **api_shield:** Acceptance tests increase coverage ([#6325](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6325)) ([3e957c7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3e957c79552895fbc0eb94312acd2eee36bf76e1))
* **api:** update composite API spec ([6d91d6b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6d91d6b476f631f3399d04e048232635ced4c066))
* **api:** update composite API spec ([a1e1df9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a1e1df96caec9d444d3be22251d924166c0039e8))
* **api:** update composite API spec ([1b9a680](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b9a680746943a309142173ec3f4619f33efa376))
* **api:** update composite API spec ([c53a1f5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c53a1f5a8f8126c07480357e9aeef4da3fc7d46f))
* **api:** update composite API spec ([ae642c6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ae642c691b141c31d9c7617cb3fd9d70bd325948))
* **api:** update composite API spec ([86ea5b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/86ea5b74700debeb7367a1fd9a78fc19fe9fc7f6))
* **api:** update composite API spec ([5bf96b0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5bf96b03ccfb01e2bab3ce2ac6443aefe3627336))
* **api:** update composite API spec ([07f3913](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07f39134d760d2ed73a5745675a7206bab98c9be))
* **api:** update composite API spec ([f7c9b47](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f7c9b479a715600aa528efc916de03f7f00d1c7b))
* **api:** update composite API spec ([1519d61](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1519d612068ee8090a7af9ad8abf52e6c1ee2393))
* **api:** update composite API spec ([a78ff01](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a78ff01d1ffa25203b89a94e2462fa164bd75d69))
* **api:** update composite API spec ([fa156c0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fa156c0dc68e9b3d7ad599459f8743e4b3841e89))
* **api:** update composite API spec ([6f4ab90](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f4ab9039465f94bcabc9fb17d871b941adf44ef))
* **api:** update composite API spec ([9455823](https://github.com/cloudflare/terraform-provider-cloudflare/commit/94558236a26d5932e2a31ad789be51d2023286a6))
* **api:** update composite API spec ([e482a4f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e482a4fd14feb2115848d35d54d9bbd4be8010e8))
* **api:** update composite API spec ([98e3585](https://github.com/cloudflare/terraform-provider-cloudflare/commit/98e35853ebf455eb7c8b2d4d49c6050879c0815b))
* **api:** update composite API spec ([d552b8c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d552b8c3da10dfc391ff0dc9970279938364d4fa))
* **api:** update composite API spec ([46b4930](https://github.com/cloudflare/terraform-provider-cloudflare/commit/46b493034b2c404c2df76e95ebb38109157c44f9))
* **api:** update composite API spec ([34eddaf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/34eddaf6d4d712f8dc03931abf53b01c69af8f5b))
* **api:** update composite API spec ([50139f0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/50139f07060123fe6f4a649a95188dec72afd711))
* **api:** update composite API spec ([b151882](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b151882d0951dc0ca43c2ae9f9684c47a8b46869))
* **api:** update composite API spec ([5892196](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5892196685804345b76c30b7043bdab22fa242d6))
* **api:** update composite API spec ([2ef377f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2ef377fe46e3ecb45142596f64a18c8f2380b8db))
* **api:** update composite API spec ([0f029be](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0f029be579b741ef19edfe3de5f6e26e90b03852))
* **api:** update composite API spec ([5483722](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5483722930ba2b3e1381ee9c04b91705cf825343))
* **api:** update composite API spec ([19ec88c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/19ec88cc4e7c86493e6711a62cfb9745b47b3ece))
* **api:** update composite API spec ([3a0fbe7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3a0fbe72169678eda0951c5c64aca0d9f990e03a))
* **api:** update composite API spec ([0ef1be1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0ef1be173722846c7d8b86f1a8e8c27d2973bb47))
* **api:** update composite API spec ([605134b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/605134b2bcddde335729bf114fb456d7229e34e0))
* **api:** update composite API spec ([d134c57](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d134c57fcd15db2c450c725de539c2de83e49796))
* **api:** update composite API spec ([0f0b7ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0f0b7ffe7362d3b1a1f3df12908f239b49277915))
* **api:** update composite API spec ([aebfa72](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aebfa720a4f8bbc49fb39c11cfe2d49edfc99de4))
* **api:** update composite API spec ([a37714a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a37714a14ecd2c878d4827b895d3d1df581da12d))
* **api:** update composite API spec ([e600699](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e600699285ea43f6638a87c12f6817c14ec460dd))
* **api:** update composite API spec ([938d787](https://github.com/cloudflare/terraform-provider-cloudflare/commit/938d78753398b28db1c916e9a7b2522649fac746))
* **api:** update composite API spec ([8ef0127](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ef012715a63cf4682a0f40b5900b65519b5f781))
* **api:** update composite API spec ([8ab046d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ab046d1a456f6188e814b123cc131693b3356ea))
* **api:** update composite API spec ([37a7311](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37a7311176aef8954695f0e29c5283f06053862c))
* **api:** update composite API spec ([038e76d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/038e76d092ab67ca07a805a256cd1c8726aeb498))
* **api:** update composite API spec ([80fc0cd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/80fc0cd8e491a0ac60ee32c0b32774f4f8abe6bb))
* **api:** update composite API spec ([ddb3468](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ddb3468edb2b89acdcb9e1eb263f52e2bdee45b1))
* **api:** update composite API spec ([88b0b8e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/88b0b8e83a075a347ba562533bd20544a6a5bd7c))
* **api:** update composite API spec ([cd64df6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cd64df678a53949708ba9d849a8b489a00c75186))
* **api:** update composite API spec ([41b75fb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/41b75fb77bb6cf07f17539b2283d11d23f9a70ff))
* **api:** update composite API spec ([3a7c0a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3a7c0a5b5d4db53f85690cf24c8e0cace86aae70))
* **api:** update composite API spec ([9c45892](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c458925f39090f7ff08054debad13d67e20a7ce))
* **api:** update composite API spec ([4e791c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4e791c5697e3636477517cd6a61389ad310d6c09))
* **api:** update composite API spec ([39816af](https://github.com/cloudflare/terraform-provider-cloudflare/commit/39816af0d8baa0ff00881fdaea1bf99073ca70a9))
* **api:** update composite API spec ([d156eef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d156eef049aa2bce0a31b5549c03be958aba3c15))
* **api:** update composite API spec ([b2bd81b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b2bd81b5ad50d06a8fb54b3279cd9ffea7492394))
* **api:** update composite API spec ([1111340](https://github.com/cloudflare/terraform-provider-cloudflare/commit/11113407d951e3d6b98ddb8dc01980e7131546b9))
* **api:** update composite API spec ([938003c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/938003c6afedda89407aa7f0e842193a965f41f9))
* **api:** update composite API spec ([0b7e283](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b7e2834c004a44d7e025b368d7514016a4edfa2))
* **api:** update composite API spec ([2e958fb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e958fbf25520b3f38c32c94c81d1098147ea866))
* **api:** update composite API spec ([138735c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/138735c2ce18b9345dbcda986284dcab27697aea))
* **api:** update composite API spec ([002fd57](https://github.com/cloudflare/terraform-provider-cloudflare/commit/002fd573cc71af999c7a6ccf76d72742d213ae18))
* **api:** update composite API spec ([302fa30](https://github.com/cloudflare/terraform-provider-cloudflare/commit/302fa302bdbb9314598eb187b4824877a5b88d54))
* **api:** update composite API spec ([0294b5e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0294b5ee5549ee484e162798d844c4e2812a0c20))
* **api:** update composite API spec ([ddb7a1c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ddb7a1cfea66b7402536dddb07c4f9f4d7b4ea1d))
* **api:** update composite API spec ([4cba4b4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4cba4b44aa939c17a1931ddbc4fdcc9c45f7bde7))
* **api:** update composite API spec ([c88075f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c88075ff98ca945364ce2e6d0e0cb226584f0062))
* **api:** update composite API spec ([23b810f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/23b810fac56a2ccb781863d2e85bf47f1e752ea9))
* **api:** update composite API spec ([8590d26](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8590d2634b29df912c7831bb4fd912a76f091d67))
* **api:** update composite API spec ([2fa0344](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2fa03442859d09b943fd9cc24d0302699e18b2fd))
* **api:** update composite API spec ([e01d191](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e01d191fb48a5ecc52044347bbf468175d5db2b6))
* **api:** update composite API spec ([0b9c36d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b9c36d397cc77c6334a6143d3de326eab67271e))
* **api:** update composite API spec ([644ff5f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/644ff5f2abc4d9dc6582e47f41a5c1c8ad87ac4a))
* **api:** update composite API spec ([fd946b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fd946b1eb24916daaa10ee86f0c69a8092de81fe))
* **api:** update composite API spec ([5b54f0d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b54f0da6c840ea6b6b5c079b41da732b45a2e10))
* **api:** update composite API spec ([0421b6a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0421b6a436786d6623646f470d5d33cd0a0fd440))
* **api:** update composite API spec ([d0f7eb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d0f7eb487eb31eefebdaff92b61879eea36fcf6f))
* **api:** update composite API spec ([e8a2650](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8a2650765ddf2aafda18e93e7f739484087b561))
* **api:** update composite API spec ([d0784b3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d0784b3e77bee624a73d1443242bd1e3927ee63b))
* **api:** update composite API spec ([3490350](https://github.com/cloudflare/terraform-provider-cloudflare/commit/349035066ad623182690511b3b01a96ce7185c33))
* **api:** update composite API spec ([c009139](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c0091397a7cbaffefaf6fb941cc8cdf96f8a36fd))
* **api:** update composite API spec ([385cc44](https://github.com/cloudflare/terraform-provider-cloudflare/commit/385cc44b8ee9c9ef5a5c0662d75667be0a4ec107))
* **api:** update composite API spec ([7dc37da](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7dc37da6a4774e339a7b29302c79b74de2ec8cc2))
* **api:** update composite API spec ([3fc4cf6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3fc4cf6a55c43b0cfdd3af1df0815da408f94500))
* **api:** update composite API spec ([d812ed9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d812ed90d0e153f04ee877de77a0aab68d21c458))
* **api:** update composite API spec ([35aca13](https://github.com/cloudflare/terraform-provider-cloudflare/commit/35aca139e0baf5129b332c7af2b15f524517a5cd))
* **api:** update composite API spec ([f610c3e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f610c3ee159af9d45afb79e9290faf60f6e4c30d))
* **api:** update composite API spec ([b435318](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b4353183a097c6d2aebe5f475d4494b6c7740edf))
* fix errors in `cloudflare_pages_project` acceptance tests ([#6318](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6318)) ([cb63e28](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cb63e287bd9d9de94296ea09580819c0c2e97f39))
* **internal:** codegen related update ([010cc1e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/010cc1ec5f1b9301727d152be6ace1564d760a73))
* **internal:** codegen related update ([fde5364](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fde5364da82867155855e56d7464857079db31bc))
* **internal:** codegen related update ([b451331](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b45133159bf3ba69e9147fba1c7640f4375b28a8))
* **internal:** codegen related update ([cf23a89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cf23a893f24ff84d6afa34a4cfafe948ed906f4c))
* **internal:** codegen related update ([f577253](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5772534cf758bae7e7e63516470dae0bfdd0262))
* **internal:** codegen related update ([70652b4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/70652b459d1de7c429faf4eec5c96fc23885ed6f))
* **internal:** codegen related update ([0504080](https://github.com/cloudflare/terraform-provider-cloudflare/commit/050408034c02c935d48996ca29f5bc31be7a7d6c))
* **logpush_jobs:** Add tests from basic to full fields, and changes on omitempty field ([#6337](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6337)) ([696abcd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/696abcdcb07701aeec498562b920133631a35d4a))
* **organization_profile:** add org id env variable for acceptance tests ([#6382](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6382)) ([37468a7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37468a7048fe9de782a2f9f1a021219e8b6bcea6))
* **organizations and organization_profiles:** Acceptance Tests and wait after create ([#6329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6329)) ([ecfd9bf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ecfd9bf837632e8fcb0505c5bca3feadd952d6d8))
* **organizations:** wire up acceptance test in CI ([#6349](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6349)) ([c1cbe9e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c1cbe9ef6a093be08917be984b5a507057275373))
* **pages_project:** only sweep pages projects resources created during testing ([#6298](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6298)) ([1a2daa3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a2daa3495620ecb6e285580649aab935565f6f0))
* **pages_project:** update CLOUDFLARE_PAGES_OWNER and CLOUDFLARE_PAGES_REPO used for acceptance tests ([#6300](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6300)) ([939499e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/939499e29181e0e0e8ca962e33165c95127bfffd))
* **queue:** Acceptance tests ([#6339](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6339)) ([d9eb75d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d9eb75d148a750ad03ec8fdc56220a888b2ea082))
* **r2_bucket_lock, r2_bucket_lifecycle:** add acceptance tests  ([#6299](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6299)) ([1fdbd28](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1fdbd28c65c693b64a59fcc66584fb51c7c8014f))
* update pr template ([#6359](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6359)) ([a062c51](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a062c514692e10093c6ad7bf9d3e99e95dd9eb97))
* **zero_trust_connectivity_directory_service:** Add wvpc / connectivity directory servic acceptance tests ([#6334](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6334)) ([63e78d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/63e78d560602556d40e984dafc0f3e4d2270cd07))
* **zero_trust_dlp_custom_profile:** shared_entries acceptance tests ([#6317](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6317)) ([83cf87b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/83cf87b8976dce8076c44639b54ae265cdb1e8a7))
* **zero_trust_network_hostname_route:** Add acceptance tests for Hostname Routes ([#6282](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6282)) ([0ec769b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0ec769b0cc2b7f50a8935c4d3051abe36dec8442))
* **zerot trust dl resources:** Add acceptance tests for DLP resources (rebased version of !5751) ([#6233](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6233)) ([cbd0568](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cbd05686898c5d34c0162e7d1b03fb6701e8e370))


### Documentation

* generate provider documentation ([#6394](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6394)) ([44843f0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44843f07ce50df850203cde084b5377faae96cea))
* generate terraform documentation ([#6384](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6384)) ([6bffa7c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6bffa7c23f5d42826b818052e5581745464530dd))

## 5.11.0 (2025-10-02)

Full Changelog: [v5.10.1...v5.11.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.10.1...v5.11.0)

### Features

* add `assets.directory` attribute for handling assets uploads in `cloudflare_workers_script` and `cloudflare_worker_version`  resources ([#6160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6160)) ([50168e5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/50168e591001f93e21250419dd4953b602a2f952))
* add comprehensive test coverage for cloudflare_zero_trust_list types and ([#6258](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6258)) ([6d2746c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6d2746c8903ee50b14b72491e8f074659de9b11a))
* add missing services to CI test runner ([#6271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6271)) ([1477df8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1477df8afdb6b80c5be8c69d612df2fed861f666))
* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([e1faeb8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e1faeb8b1a41d396454b59635292def7f3bdcbbc))
* Add custom origin trust store support ([175f4f5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/175f4f5b049d62ac9ec8a830d96bba5c346577df))
* Add Terraform resource for Workflows ([7533c05](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7533c054484b6993fcef60d4beb9ab2787474d32))
* Add leaked credential check resources ([c6be1c6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c6be1c6d5388b5c867f86ebb57e8e25dea3a2e75))
* Update worker and access application schemas ([ed096e0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ed096e0a5309bd08bb12c046043457b3a3ba34db))
* Adding new self-service SSO APIs ([007bdbc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/007bdbcf314df94887def4f37df6c1f73b772319))
* Changing SSO update from put to patch ([f67fbd5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f67fbd522ccf3eae431cd3bb8b46b3223a951a61))
* Rename duplicate parameter in the to_markdown subresource ([07ccc50](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07ccc503b60e2de61db07dda02feeda7249b313e))
* Add to_markdown subresource to AI resource ([1a71265](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a7126548ce7f6237261c01ff17ee5443e05fee8))

### Bug Fixes

* bugfix for setting JSON keys with special characters ([9a106e3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a106e3db2abf6f741c972866cbb669d0871bbde))
* **build:** fix broken builds on 'next' ([#6280](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6280)) ([2224d8a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2224d8acc57f13dbcbef4299fcb08592eb3045bc))
* **build:** revert cache resources to released state ([#6289](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6289)) ([e62250c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e62250c0e4b59a12bd73849f0a7581e815c9a9e9))
* case-insensitive location handling for R2 bucket resources ([#6026](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6026)) ([78c33ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/78c33ff9a2028c064c42f74b87eb8e9ecc7129e8))
* cloudflare_workers_custom_domain failing to update ([#6082](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6082)) ([46203a3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/46203a3d7045d529697b357c8c4d1f3e07201619))
* fix acceptance tests in CI ([#6286](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6286)) ([c0a9e89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c0a9e89c3410374f426cab0ac90a9bc14dce70c4))
* Fix zero trust access application acceptance tests ([#6243](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6243)) ([4a2cbdb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4a2cbdba426906e310be791c309ebda4d66935b5))
* **list_item:** source url validation ([#6226](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6226)) ([70abffa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/70abffa560bdc74f0901bfb2582ff9b85f2ecd28))
* **migrate:** concatenate static and dynamic rules blocks ([#6215](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6215)) ([be571d8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/be571d890da1cd2dabac44d649edd55a9ba4949c))
* **migrate:** page rules status defaults ([#6212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6212)) ([42a83d1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/42a83d1a1f03264a3a67124c78abef16027d6c05))
* **migrate:** zt access app default type ([#6218](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6218)) ([cea98f8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cea98f8705d40ac2cf4b83171212b33d6a97fa4b))
* **r2_bucket:** case-insensitive location comparison and preserve state case in R2 bucket resource ([#6211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6211)) ([5babbb1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5babbb109a968f8d18c78e2ad5a3dd52ac47fd2c))
* resolve compilation and schema parity errors across multiple services ([#6241](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6241)) ([052cab8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/052cab85afa97d2b0637578af85e75aeff40d0f8))
* resolve compilation errors in zero_trust_access_application and workers_script ([#6230](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6230)) ([2b78333](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b78333e76ac7a9eec9c06796fe396451100274b))
* resolve provider schema validation errors and R2 bucket test failures ([#6222](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6222)) ([2df6eb1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2df6eb168aae57f85a2553490f7e79e19e19a57c))
* resolve zero trust test failures from computed attribute refresh drift ([#6224](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6224)) ([5351c6a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5351c6a041f2cdb0f160438b862088509f3a3c2d))
* **ruleset:** allow rewrite rules to set an empty URL query string ([#6256](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6256)) ([b177cc0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b177cc0ad8e6fd9a0f48687e4ad078656f1ef86b))
* workers script migration ([#6210](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6210)) ([dca249e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dca249eb5737fdb76a958cbe0ad72967031666d9))


### Chores

* add easy sweeper script ([#6220](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6220)) ([7fe36e5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7fe36e5d453663cb3885bc104588c3baec0283f6))
* do not install brew dependencies in ./scripts/bootstrap by default ([44f11c3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44f11c339bf4d86a5a0cdcdf75c746c17675bd66))
* ensure `tfplugindocs` always use `/var/tmp` for compilation on linux ([3ccb727](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3ccb727887e38def156f225983cc1c5b1cdec004))
* improve example values ([1ac2c1e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1ac2c1e71426edace6009b9a4a5b018c92cc7122))
* **logpush_jobs:** Switch to Plan and State Checks from legacy Checks for logpush_jobs resource ([#6083](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6083)) ([5933a83](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5933a83c26663fc36f7f7c7b3eb39116af9a0aa6))
* **mcp:** allow pointing `docs_search` tool at other URLs ([195dbf4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/195dbf40ef6e5955a7ce800fb65e280d292d8837))
* **migrate:** remove debug statements from migration tool ([#6223](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6223)) ([b50ee7c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b50ee7cc3d9c0b310d63c0c452d053d306234ef1))
* run migration tests with sweepers ([#6209](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6209)) ([489795d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/489795d1a78f8a41c12701c253088f5ea9d4abc4))
* run workers_kv and regional_hostname tests in CI ([#6240](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6240)) ([5433d6b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5433d6b2b69c8e0e71a5327cabf2b1589d63518c))
* skip mtls migration test ([#6207](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6207)) ([1730fa2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1730fa26cb6ad6739f99ad3ec6b9989b434f1812))
* **test:** use no-grit by default when running migration tests ([#6214](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6214)) ([267dfc2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/267dfc2e3d13311e0b60afd205006ca0c5785e8b))
* **zero_trust_dex_test:** Updated acceptance tests  ([#6183](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6183)) ([cd7af0a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cd7af0a4c3adb14e2b199e0702da0e3ca3d3dd6b))
* Point to next version ([ac09de9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ac09de93f8bde47288cdaadb2fc413afcadae4b4))
* modernize and expand cloudflare_zero_trust_access_service_token test ([#6260](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6260)) ([79e891e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/79e891e2e2f3f948224bcf06c118e08904f7ccd1))
* Modernize and expand test coverage for zero_trust_device_posture_rule ([#6259](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6259)) ([6666597](https://github.com/cloudflare/terraform-provider-cloudflare/commit/666659775bb190582da9269a0b2766db2ea9141c))
* modernize and expand test coverage for zero_trust_gateway_policy ([#6266](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6266)) ([9ad5fc1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9ad5fc1f2501a68546a8b4d96f8982113e1320ef))
* Modernize and expand test coverage for zero_trust_tunnel_cloudflared_route ([#6264](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6264)) ([2c4a1e3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2c4a1e316dd13fa6cbe4e22c2070cd70215e2c5c))
* modernize and fix cloudflare_zone_dnssec tests with comprehensive ([#6254](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6254)) ([d70cdd1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d70cdd1193c06549a517b41c93a9c28f990599c1))
* modernize and improve cloudflare_pages_project test coverage ([#6274](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6274)) ([ebdbece](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ebdbece27887d249eeffdc91116797a5c46b509c))
* sweepers for workers_kv and zero_trust_list ([#6281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6281)) ([2ed457a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2ed457a406bd7faee78fa5183000c37ccdba6454))
* Update example values ([c889ef1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c889ef143db19299190ab1947a43dd8d169ea49e))
* Use cloudflare-go v6.1.0 for v5.11.0 release ([5ae607d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5ae607d0ff5e95eaa3c18634aae1106be9d31842))
* Fix config for disabling codegen in access_application ([75d8d6b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/75d8d6bcb15616f9a28d518e30847e8e37345d68))
* **api:** update composite API spec ([ee08de7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee08de70b1ae501491d81c7633e98c63df6e438a), [b75aafd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b75aafd6710e78d08f478213d3b8d0c42ad12af4), [d140041](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d1400414d003a7d85a23e00e3c0327190ff12d4f), [d22865a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d22865aad1674bf86130e6faee4746480131dccc), [cbf8d12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cbf8d12fa47ece5b5f09a011286b5707b43f353c), [1b987a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b987a0893529c54356ba1f0945f9ea22ecb9f7b), [d98c4cc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d98c4ccda843b468f3a5223bb92283146928df88), [02f821d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/02f821df0e613e80d7a2d6569f90106dfbd2b90b), [f965e4c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f965e4cd825b50ca2bba7a3d424dcacad11983a1), [9663d30](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9663d30d86181c611d23e8a558b7b58754b77999), [7839869](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7839869aa7af2723d3ce69942aa6305f178dec82), [d15f035](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d15f035fecf11ddaee1cb70fa6cbd2177c985dd7), [2a0c724](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2a0c7249a6ba9b95b43971627658780ba3d805b8), [ca18700](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ca1870030b1980509ebdeca8154fb761be4a6dfa), [7c43df9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7c43df913b5d8f5d5111a283a7bce2e35c250d4c), [0df0b82](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0df0b825cae7f19f04346de13ee1a027d17b4d06), [a50d130](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a50d1307db29ae8b6c7c1ff7725c8bf97f78a95b), [8fabca3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8fabca3d749ab94c3768897fa9b0059d166931cd), [00bb6e6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/00bb6e68964f6f9233422af5df5e4f8077dfd951), [f96417c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f96417c6d36ea4d52d6a18f4f64d4c3a40fd4a81), [86e8862](https://github.com/cloudflare/terraform-provider-cloudflare/commit/86e886284ed179bf9160d2c73bdb3a66d74b0380), [270f556](https://github.com/cloudflare/terraform-provider-cloudflare/commit/270f5561b02beea84b1622517123a5e500d10b01), [226cce5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/226cce5e08ee07449eecc300526b0b59cc92152d), [e8e7241](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8e724178a3059c34b91ae92e46a4419e73fb076), [f720dde](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f720dde7b797f7eed65a821291fec855741f4a1a), [125870b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/125870bb54773db07a53683fb926dee3478c41a7), [e4b3a36](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e4b3a368f020f2171b38903a575f4ceec2dc96e3), [e4c0f73](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e4c0f73327 6b9f339161e2e8bc37fcf7e4ec8b4d), [82c913f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/82c913f60fb38f3f8d855841176a33548673bb61), [64ff1f0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/64ff1f06354b31b0fab27bd841c8eb32be2cca75), [516b468](https://github.com/cloudflare/terraform-provider-cloudflare/commit/516b468aa5f7d4b33ebae7e4269478301c8e6eba), [b9b674c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b9b674c3f88f973ce3ecc9412ddfa08ff35a3c32), [0b6b710](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b6b7107ec1c35662d38026aa5ded32b652a12dd), [fc1f691](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc1f691bdab2ab9596f96d722c00447486e930d4), [7139af1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7139af1ace30e984d2033f5153276383454ad074), [81d79a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/81d79a54d14b26ce7f33e4bb48c4a6910b7b9aba), [fcb14b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fcb14b63e771a88de7c423f0629a001752969d7e), [efebe55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/efebe5571d49bcc4eab520e6094913bcc4603702), [c0fb2a4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c0fb2a47f2a8befafa7169df8e119257012caac5), [2978db8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2978db81702e425f7f8e4d17a64c1f6058e58426), [ced7790](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ced7790f730997ce51f8f10fbaea0f41e5d09427), [0ff7a11](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0ff7a1133179d0f7ed944cd838419779e88561db), [04c89cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/04c89cb6a22d7cc36636b47b2d772a77e0ef240a), [4acb7ab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4acb7ab62079e7c140111e236833772b8bc424ae), [04236ca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/04236ca9926ed8c84242a69053116f82e58a75cf), [7636b30](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7636b3077e49a7b7a303f7ecf11cee01edacf0de), [5c2d985](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5c2d985a1c1de384de89c0f03d02219535951f57), [e52ffb1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e52ffb190c251376ed585f273b7ccba8636687eb), [133bb02](https://github.com/cloudflare/terraform-provider-cloudflare/commit/133bb020d8adc9e8e180f2c9d830e323522140b9), [3fd9175](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3fd91759c7b4214047a1469ec78394843a86c3ef), [b5f2c46](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5f2c46ba266fb1cd3a392f438359e823ddf1699), [f52687d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f52687d91db396787ca0d150cb83f0c5b3415a41), [f1164d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f1164d20453d9cce8f5b0a2effe37afc1f0e9afd), [e4a44a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e4a44a8af68e89ac6ce82ac2aeb66149b1e6f09f), [da6a341](https://github.com/cloudflare/terraform-provider-cloudflare/commit/da6a341e9912b9a128f5edcc674c91b982282972), [002316c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/002316ca9c4ebf5816b352ed61fc7bf0933a1f79))
* **internal:** codegen related update ([3457590](https://github.com/cloudflare/terraform-provider-cloudflare/commit/34575907e189e3fa4d65c0ef5640c7883470ff0a), [290fbad](https://github.com/cloudflare/terraform-provider-cloudflare/commit/290fbad69215e23603da96f6065a1e90bd5aefff), [8d68abc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8d68abcb902560158bec06391b876af5935d2816), [ed99054](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ed99054c30a3fe86997b3d573ef344ca2d3771dd), [95603f6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/95603f6c0bbf2d294d606f28d0b1b3a9d991a84d), [e9b678b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e9b678bf9592f73ba26dab78de9e2254b88e8eb2), [2416723](https://github.com/cloudflare/terraform-provider-cloudflare/commit/24167238cd1e4050916ada729a785d387f2926da), [2e822a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e822a90a32f755f2805052377e8c190f3cd7d5b), [b02b92c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b02b92c0714cb134bf78bcf274a4974ba231f517), [7310778](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7310778d5d8511b8a21425d742b5269850415846), [7fb69e9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7fb69e978e01823fada800f94f203adb4f760e7e), [1f20a56](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1f20a5669a2ed63e279bf16f3863cea91e1ea7f4), [4a181d7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4a181d7ba0b8e3e75881f1b221c2b27599fd93db), [c800edc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c800edc5e7f0c3bed3edc521c19add6e9099b74b), [3a74bdd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3a74bdde6bb1d58b17fa621b58f0da727d0b0ca9), [5fe113d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5fe113dcb36a8d1d390f5394d9c7d4925d8ea8a7))


## 5.10.1 (2025-09-17)

Full Changelog: [v5.10.0...v5.10.1](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.10.0...v5.10.1)

### Features

* grit to go ([#6162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6162)) ([b3c4779](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b3c47796075888b92723ee8888bd8de9e3ab00b3))


### Bug Fixes

* cloudflare_load_balancer transformation issues ([#6171](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6171)) ([92f4a4a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/92f4a4ac3cdf493d0c543a6a234d74772f349236))
* fix grit in migration tests ([#6175](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6175)) ([0a25a5e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0a25a5e05c3f63866e493839419a0a253ac19fee))
* fix zero trust access application state migration ([07a5d06](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07a5d06b6ef6a8b1513ffb74ec7ce501e7c42b33))
* handling of nested arrays in ruleset migration ([#6187](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6187)) ([a00b67f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a00b67f55cf19a0a7d963c16e29530a27e75983e))
* lb and lb pool config migration ([#6170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6170)) ([2af41f8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2af41f86cd166578f227f27ef4edbbb0d2feb4eb))
* lb monitor state migration ([#6180](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6180)) ([c9811ba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c9811bad042a8cb312b7c4db1c734db65730140b))
* **migrate:** block transformations ([#6203](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6203)) ([245166d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/245166d6f4b3e348d9d3b7f5cca55d9fc51e3f60))
* **migrate:** fix main_module value migration ([#6204](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6204)) ([fd24c07](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fd24c0743e745a5954ecc1e4a94a4415fbb5aa8f))
* **migrate:** improve `zone_setting` migrations ([#6169](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6169)) ([6ba251f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6ba251f7c99e04588534399bd03d34aa7e88fd7b))
* remove 'disable_railgun' from state after v4 migration ([#6186](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6186)) ([453d774](https://github.com/cloudflare/terraform-provider-cloudflare/commit/453d774ee162565e2b080ba2291cb2f319eac3d1))
* remove zone settings with null values ([#6201](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6201)) ([f99bac4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f99bac421f6ba694312357a963dfffb6fe431cc6))
* ruleset migration in nogrit ([#6174](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6174)) ([ecb450d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ecb450d4adde5adb19a49bbbf5026d1995792705))
* ruleset migration issues ([#6163](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6163)) ([44b653c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44b653c81b591b3e865fe3d9cae6b88033914c48))
* ruleset migration issues ([#6168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6168)) ([11b3961](https://github.com/cloudflare/terraform-provider-cloudflare/commit/11b396108147fdad462bb0fb839f1f0766a7fc9b))
* ruleset state ([#6191](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6191)) ([dfd36a2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dfd36a22b9c45f7d7a13855cfb8bbcdb0fa9ac3b))
* variable interpolation ([#6193](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6193)) ([332de8d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/332de8dd71f5526cf55087428f5f069378198730))
* **workers_script:** fix incorect model type of `run_worker_first` attribute ([#6199](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6199)) ([13bf28d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/13bf28d11e16eebcbc7ef5c31e7152e47fd02df7))
* **zone_setting:** ensure clean state after migrate ([#6190](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6190)) ([41ae093](https://github.com/cloudflare/terraform-provider-cloudflare/commit/41ae0938374a3a8bde92c002b4a69ba06a60c73f))


### Chores

* compare better ([#6192](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6192)) ([ff67b9e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ff67b9ee390e7964ce0e79fc6cdaadee606d9486))
* enable mconn tests ([#6166](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6166)) ([9c6653b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c6653bab8b8b4ca35fe14bb24bdc68e4d39c268))
* limit max retries ([#6173](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6173)) ([8c1c81c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8c1c81c76d97e932e91287700c44c844d15040cd))
* point transformations to gh/next ([#6177](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6177)) ([c35109a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c35109a5d11d1e03b56f043e1491ccde577d8467))
* zero trust config issues ([#6179](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6179)) ([07a5d06](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07a5d06b6ef6a8b1513ffb74ec7ce501e7c42b33))


### Documentation

* **list_item:** add import documentation ([#6202](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6202)) ([55e12bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/55e12bc5b1faed26af312ddd129b9111073b21a7))

## 5.10.0 (2025-09-12)

Full Changelog: [v5.9.0...v5.10.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.9.0...v5.10.0)

### Features

* add 'ruleset' support in migration tool ([#6104](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6104)) ([82158eb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/82158ebe9c2805b4cfe7533eda70008fbb5fcb12))
* add migration tool support for cloudflare_snippet ([79e19d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/79e19d31e7637e3f2904325d2f19ddbf17072829))
* add migration tool support for cloudflare_snippet_rules ([b1d4e92](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b1d4e92a4e8eca623713465085eefee4ef03b7ab))
* **cloudflare_list:** add nested list items to data source ([0818c2d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0818c2dc98aea55e6777054040ebd131bf8dc370))
* **cloudflare_list:** add nested set list items ([f96b922](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f96b92263421221105b973d506db71793ebe4c26))
* handle list items in v4 ([5c315f2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5c315f207f6c269ea6fd05509088e52bb50462d5))
* **internal:** support CustomMarshaler interface for encoding types ([3ce3cbc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3ce3cbc758bed234da930cbab5e21611850807b5))
* Merge branch 'vaishak/bump-sdk-version' into 'main' ([2d4ae17](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2d4ae176bc2668af22f36c93f99f296bf671d017))
* merge items into list ([0709233](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0709233a0d590f60fdf8f4cc66c3444fe12b7010))
* migrate list with embedded items from v4 to v5 ([40ff2dd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/40ff2dd685dbcf2d4301d6c96c84f0df761001c7))
* **migrate:** add comprehensive workers cross-resource reference support ([39032e3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/39032e3811883fc0e7c41de47fd9c2175de5ecb6))
* **migrate:** fix load_balancer migration test ([#6148](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6148)) ([1d21133](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1d2113344f04e88fea08449ab9f8db2f484f048b))
* **migrate:** implement comprehensive workers_script v4→v5 bindings migration ([59d436b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/59d436bb1188180190c5abdea70ed556839604fe))
* **migrate:** implement remaining workers_script binding migration fixes ([6324582](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6324582362020b435cd0da44fe225d577551970a))
* **migrate:** support migrations for workers_route and workers_script ([3308fa5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3308fa50c601ce533edf49a2625592cf47b99f4f))
* migration tests ([1e35d38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e35d38bfb048726717efe2e22dad796b6c0524a))
* migration tests ([8ec2d24](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ec2d246cdb149dbc729572f844d594e020a47db))
* modernize healthcheck tests ([74a358d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/74a358d018e66a2373ac6bb07d461c6841c30717))
* modernize notification_policy_webhooks tests ([bbab7d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bbab7d571dda978540c6cb6b0f20021be6bc8929))
* modernize r2_bucket tests ([ecf2609](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ecf2609108e1a408b0ce049c0b68f170c834f6d5))
* modernize zero_trust_dlp_custom_profile tests ([cb11079](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cb11079fc2ab25b214aeb5184a633d1a9deb85c0))
* modernize zero_trust_dlp_entry tests ([6f0a6b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f0a6b1efa72b2028235eb74bb7590e684fdda92))
* modernize zero_trust_list tests ([27e8cab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/27e8cabaa116afba22318e11bafc1169893fd727))
* **ruleset:** validate action parameters are used with correct action ([578879e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/578879ee9b5475a8c4d9e0ed11152f490bd4bb47))
* **workers_route, workers_script:** implement migration for workers_(script|route) ([ff3e68e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ff3e68e6fd1714084a97d29e694860d8a2603f53))
* **zone:** add v4 -&gt; v5 migrations ([279070c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/279070c1f677a0bdc91bfdebb42c17e202bb54fe))
* **zone:** implement migrations ([717787a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/717787a5990fd07a4e603874835c18caa04cb2b5))


### Bug Fixes

* broken test data and block attribute conversion ([#6138](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6138)) ([6a07ac2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6a07ac20e6c919b1a4daaf27a4464711b20dc870))
* ci workflows ([ee2117a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee2117a390f9e0deba1e9a8a09af9f5f8c189a0e))
* comment_modified_on drift in DNS records ([b5bdee4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5bdee4baf0c6a30264306b23ec7aa1b89ea03da))
* discord failure ([8b8eb19](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8b8eb195f3209b58269b075fdd46be47b7df46c2))
* dns record empty states ([e8f418e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8f418e608e7cb463cf6aaaa403c10861e1c67c3))
* don't announce to discord ([9e7a495](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e7a495ebb7c8f263a64c1dcd9f5067261159b75))
* dynamic 'origins' blocks migrations ([8d5de51](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8d5de517638d39293a1ef80728a4a45bf638badc))
* dynamic blocks and tests ([4d6855c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4d6855cc71bcc00331a50e2f2684150af09bfce3))
* fix snippets tests ([52dfe49](https://github.com/cloudflare/terraform-provider-cloudflare/commit/52dfe49e665354509a82dce898d24cf9c07b82fd))
* fix zero_trust_dlp_entry acceptance tests ([5007122](https://github.com/cloudflare/terraform-provider-cloudflare/commit/50071226f18a74e58ed7e4eb87bd9b66ccb02611))
* inconsistent apply Issue [#6076](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6076) ([#6139](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6139)) ([0e9650f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0e9650fc1112739a315cdaebcd6157d8451251d2))
* **migrate:** add missing hyperdrive binding attribute renames ([fc6b137](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc6b1376bc56cf945f36ff1d95f87c03c526bdb6))
* **migrate:** correct module transformation and clean up dead code tests ([c364035](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c36403502a8b55e7d3ee646e168e0157d9552010))
* **migrate:** custom_pages state migrations ([8b4e1cf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8b4e1cf913cc6e3f6afd05fa579b53cb614ecdb3))
* **migrate:** implement dispatch_namespace attribute to binding migration ([3a432dd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3a432dd532647f6a02723d52eb7c8d50a278eb0c))
* **migrate:** implement module attribute to main_module/body_part migration ([dd94222](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dd9422222ed36a608c552c41f35a11f00ce0be75))
* **migrate:** implement workers_secret cross-resource migration to secret_text bindings ([5ff5c6c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5ff5c6c9971b416b91675291c14d7b227539b93d))
* migrations for config and state ([24e44b0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/24e44b0bd72f03974d1195bb3a9d3931a933d52b))
* more roboust retry logic for certificate tests ([#6154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6154)) ([25039bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/25039bc51edb81b2ea20d132eea8da46be3535b2))
* nil dereference in `cloudflare_workers_script` resource ([#6158](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6158)) ([bcfc129](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bcfc12947857b8a742b068b32707af4fcf86476a)), closes [#6147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6147)
* prevent resource type corruption in workers_secret state migration ([abc0548](https://github.com/cloudflare/terraform-provider-cloudflare/commit/abc0548e62f82d7aa27122ce607ee501c063ee15))
* prevent resource type corruption in workers_secret state migration ([adbbae2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/adbbae282527fff3f2afc9422e593072b77f5ecd))
* required field ttl ([b88e5b8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b88e5b8d4cf23ecca319d0b2807213d2e053db3a))
* resolve race condition in zero_trust_access_mtls_hostname_settings migration tests ([#6152](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6152)) ([9c6deef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c6deef26dcb8c1141f7494517ee63aae5761e87))
* run spinnets in sequence ([0789979](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0789979eacf3d7350c4958bf12d260609aae6d56))
* skip acceptance tests in unit test scope ([#6155](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6155)) ([e860eb5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e860eb515f334fc44f5df2d50cc5530f672a320c))
* snippet and load balancer migration tests ([#6149](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6149)) ([a347ebc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a347ebccdb5b08a66a5ba98ba36b30b053d43955))
* state ([5742920](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5742920c0afafac57ee3f15da951661fc347d990))
* state test ([16c5fc2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/16c5fc2ec0b0fbcae1f048a134287c8d4a1eedd4))
* tests ([84f27a3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/84f27a375e4e11192dbc19902e0f28755af1cae6))
* tiered cache test ([408a4b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/408a4b6fc9a47fa21dc94aada71144e89115dc10))
* **workers_script:** fix/improve bindings tests ([fca02f3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fca02f34d358d7d2f20a9ecbb5bb87718a4fe2bc))
* **workers_script:** get tests passing again ([404a241](https://github.com/cloudflare/terraform-provider-cloudflare/commit/404a24114484ee16541875a96f5f1ed69ff01891))
* **workers_script:** referenced attribute renames ([#6136](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6136)) ([29d686c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/29d686c2ed811591522202ce28812b2c80771028))
* **workers_script:** resolve binding order infinite loop in v5 provider ([a05f552](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a05f552a51ee3138cf584f4b7a36eeb90db1706c))
* zero trust access indetity provider migration ([7bc2a5d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7bc2a5d3e9dc0dfbb4621f99dd3657ac37732d58))
* zero trust access mtls certificate acceptance tests ([7e91d44](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7e91d44c17fa429c750ad05e9018f8df2806dbca))
* zero trust migrations ([ddc8642](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ddc8642939a84e9a7a71def40d5852f38146f2b1))
* zero_trust_access_mtls_certificate acceptance tests ([b163147](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b16314700cc13a3535304f0db4ba53f406fcec21))
* zero_trust_device_custom_profile sweeper ([f0ed7ca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f0ed7caa94a44ee8b4fbc803875387c3de7022ff))


### Chores

* **account_member:** remove bad test ([21c670d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/21c670d3da4dd30d55b7a5ef68160b55139baa23))
* **account_member:** update acceptance tests ([1988556](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1988556b720e4f7ab608d28709972c1837ced47c))
* **account_token:** update acceptance tests ([1f6be84](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1f6be843ca4b8d516fd7dcd2a9e3d532ab94ed0f))
* **account:** update acceptance tests ([1f49327](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1f49327a607052f4efed1bc052ebfab11ad39a69))
* add migration tool support for load_balancer and load_balancer_pool ([a985fa0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a985fa0980c19940ed9ddb1b999c78a85a0be09a))
* **api_token:** update acceptance tests ([5d02104](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5d02104c772bfa31428827e2c337e529ce4b8338))
* ci tests dependencies and job tracking ([a2142fb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a2142fb0d7c188df0a64f4bf3133d34586defc7e))
* enable more ci tests ([cd96052](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cd96052cfff30a17c8bc80baadd87dbf99644dd3))
* fix list item state migration ([#6146](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6146)) ([7cc6425](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7cc6425756210ab2803fb3844fe1f27cc4cce47d))
* fix transformation source ([#6157](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6157)) ([6cc2cfb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6cc2cfb2b2642bacb10d952a6826c854b6b77077))
* grit to go ([#6143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6143)) ([548f097](https://github.com/cloudflare/terraform-provider-cloudflare/commit/548f0974413728c5091b856967d4a2be3c7dbdf4))
* increase parallel jobs ([9f1a098](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9f1a098b378868c54865ddd9dd846fa6851c922c))
* increase retries ([#6156](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6156)) ([309397c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/309397c3b2f074ce4eb55670eed6a0483e9d5b5c))
* remove files that are not needed ([064c780](https://github.com/cloudflare/terraform-provider-cloudflare/commit/064c780896778d03813d2fe86f4dea121ef34b3b))
* remove grit ([d840087](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d840087e04d87b146fb4677e387512b54bb659c6))
* remove grit for lists ([f7d6229](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f7d62295002beb894c5b4e2df24083cd08170720))
* remove skips ([adee34c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/adee34c10d4e009387f542de1ad6a4f23552852e))
* retry tests ([#6150](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6150)) ([fc88bf8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc88bf855232fcdfeec535692048391a0f8e1107))
* revert grit to go ([#6159](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6159)) ([09bfa12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/09bfa12a51658b0b3879d2916d703de46e661637))
* run goimport ([f837802](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f8378026a3283ee350503f98471f71ba2e0a9d7d))
* sequence magic tests ([#6145](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6145)) ([156694c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/156694ccdb4326e0326385eba48f931fac3dd8a3))
* skip mconn test ([#6161](https://github.com/cloudflare/terraform-provider-cloudflare/issues/6161)) ([e181aef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e181aef8531e4ad75d2bcaaccb5aed27bc9c8567))
* tests ([47a4330](https://github.com/cloudflare/terraform-provider-cloudflare/commit/47a4330a2efd1f3d210f6a5e839295d9939aab87))
* **workers_script:** add lots of missing bindings tests ([dc27751](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc2775112f61adfd7e29ac56b5a7231ee6adc333))
* **workers_script:** remove unused tests ([505c840](https://github.com/cloudflare/terraform-provider-cloudflare/commit/505c8405f93822f1d1a0acb2831eef0751c26ae0))
* **zone:** add migration tests ([b02bf9b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b02bf9b0fa8a40897f56dd33f68c80084c3122be))


### Refactors

* **migrate:** add resource rename support to workers_route and workers_script ([35eaca1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/35eaca1b009cffb755e6700564f91af70bd2d78a))
* **migrate:** clean up duplicate bindings maps ([2393057](https://github.com/cloudflare/terraform-provider-cloudflare/commit/23930572971917df85bda1d7361bb8d91944d402))

## 5.9.0 (2025-08-29)

Full Changelog: [v5.8.4...v5.9.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.8.4...v5.9.0)

### Features

* add comprehensive zero_trust_access_group v4→v5 migration support ([44b55c1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44b55c154c2f203db4ae742484281689d34dd6e4))
* add job IDs ([8bcdbd5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8bcdbd59a65932ea72e61517ce47d58fbf10c2cd))
* **api:** api update ([b9b17cf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b9b17cf612e3865ff8d78a19c748db65c9b8beb6))
* **api:** api update ([8ec5c0e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ec5c0e9e9a5f20caedeb67e25f3d53416755140))
* **api:** api update ([fb4eddb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb4eddbc189646deb65b91fca226c0f8e200d4a8))
* **api:** api update ([12c4328](https://github.com/cloudflare/terraform-provider-cloudflare/commit/12c4328bd833342ccda20fff0aad103062e58160))
* **api:** api update ([511614c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/511614ccc15b8507c7c36d7498e1cbc4f0975cfc))
* **api:** api update ([1d22129](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1d221295bcb9997e237ace7d37e274cdf959d7af))
* **api:** api update ([7391faf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7391faf90673b96356929cbdf809dcf23c1e9033))
* **api:** api update ([4fa333d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4fa333d1694c2d17a15d888f895b78169490a49e))
* **api:** api update ([5f93f24](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f93f2451b21dce4760c019ac809d2659456256c))
* **api:** api update ([b584f87](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b584f872ba8497e30f966c2db110eaebb819cbb9))
* **api:** api update ([ff12699](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ff126990219cd5c4b9270888c4c5ca11b25559e3))
* **api:** api update ([ad38f3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ad38f3f36905edd69e945b18dd3bb810ceec2634))
* **api:** api update ([51bbca9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/51bbca9400189e86632c59c645db5d858ea09906))
* **api:** api update ([d873115](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d87311597d69ba1151ce5ff0d4d4ba61ce981bb9))
* **api:** api update ([a6065ba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a6065ba7e665a61c1154e551d03826f2370db939))
* **api:** api update ([2545b2c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2545b2cbef4c618aaca98ea4a5025749618f1270))
* **api:** api update ([abcd800](https://github.com/cloudflare/terraform-provider-cloudflare/commit/abcd8007805caa6927551780f805bd6b68f96f2c))
* **api:** api update ([9251429](https://github.com/cloudflare/terraform-provider-cloudflare/commit/925142939ca7ad9b4abdc4857195708d466a0973))
* **api:** api update ([4980f38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4980f380b1b91b708ec3f7af82f688646a271310))
* **api:** api update ([76010ab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/76010ab284b1ca21fb47589759b06aed1b71d89e))
* **api:** api update ([627dfd8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/627dfd8603004d8728eb74a4892fa0ab78cda452))
* **api:** api update ([b6e00b8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b6e00b830a5e424a06ff2ad17aa07a7d769753af))
* **api:** api update ([8bfb0c4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8bfb0c4a9694278349313e20957cd00bdeb29f30))
* **api:** api update ([d614e59](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d614e5914d47b088fc4d05e9a559eb1745664821))
* **api:** api update ([b1cb9f3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b1cb9f3104a55c66ee201b8003d9be9565f492b2))
* **api:** api update ([098a710](https://github.com/cloudflare/terraform-provider-cloudflare/commit/098a710a0fcd535843d8a6ba869efed9bf1b4eed))
* **api:** api update ([da44c34](https://github.com/cloudflare/terraform-provider-cloudflare/commit/da44c34ed0c83d316557f40e34cb22ff189a16e5))
* **api:** api update ([276d413](https://github.com/cloudflare/terraform-provider-cloudflare/commit/276d4135214a67b1e263c76a43d3b33b761972c6))
* **api:** api update ([403f6a6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/403f6a60f718bcc73edea84bbfcf02d280eb0562))
* **api:** api update ([51788e7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/51788e7957c0c4dbb5d2ba5d88dd09a2a6b29973))
* **api:** api update ([841fa60](https://github.com/cloudflare/terraform-provider-cloudflare/commit/841fa60466aac847af13272c68a4af7c28a852c5))
* **api:** api update ([cfcd80d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cfcd80d3a0d35c4888023b5ca1ade754a3614490))
* **api:** api update ([8832c77](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8832c7792bdbc325df787d4c0ab1c3e47f665a53))
* **api:** api update ([f02e8ce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f02e8ced752689e0deee321d82df5958707c5f02))
* **api:** api update ([bfd878b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bfd878b5e8f06a5ef0af387294b387750fb80957))
* implement automated v4→v5 migrations for zero trust access application ([344d995](https://github.com/cloudflare/terraform-provider-cloudflare/commit/344d995166dfa359a4336dfa98d6226fd6f3ecde))
* implement automated v4→v5 migrations for zero trust access policy ([204a2c0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/204a2c081ccbb17d997b63a568398d343c3fa1b2))
* implement comprehensive v4 to v5 migration for zero_trust_access_group resources ([c63bd93](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c63bd93f05e50a6a00f4b6a2901688190dde2644))
* implement comprehensive v4→v5 migration for tiered_cache resources ([1d763a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1d763a9236a882f23625d322dcb3c7d76633edb8))
* migrate argo ([21bef66](https://github.com/cloudflare/terraform-provider-cloudflare/commit/21bef664571b2350338a23c8064c730c11f2489a))
* migrate state ([47c25c2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/47c25c2f405a0bfe93581a482f5aa60a85612d61))
* migrate state ([799e3c2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/799e3c2693d8eec516a87d3938e562d0868fbcea))
* **migrate:** fix access application domain_type and destinations migration ([f572e62](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f572e62fa0cc26e19ffc6f6b9924322bdc8be04f))
* **migrate:** remove skip_app_launcher_login_page when type is not app_launcher ([1cfa5ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1cfa5ecdda17b7b3a6dba17ab74183b40647b210))
* migrations for zero_trust_access_application ([bb35c7c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bb35c7cfa814d9e705bd0709eddec2b6c1a17c3d))
* migrations for zero_trust_access_identity_provider ([75e412f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/75e412f7b4b5898832c62cc148cd79f9f43e751d))
* migrations for zero_trust_access_mtls_certificate ([4d547c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4d547c5b9d4f1d691bbdb698ee31d03d87c47b6d))
* migrations for zero_trust_access_mtls_hostname_settings ([82ff1cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/82ff1cb551385c8d27b5f4892e3abb5ef3fa5876))
* parallel test runs ([48d5c35](https://github.com/cloudflare/terraform-provider-cloudflare/commit/48d5c35864199deed796b2bf3decb4100b567731))
* **regional_hostname:** support migration from v4 to v5 ([ffd589d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ffd589ddd8df640f703d5aaa89cd61d8d8cc613c))
* state upgrader ([ed2ad91](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ed2ad911cf4bb76b80ea137d8385e4a65f280eb6))
* sweeper for magic_wan_static_route ([51728f9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/51728f9b2ec8141643de5102fb0b39d8154087c5))
* zero trust access policy migrations ([134df93](https://github.com/cloudflare/terraform-provider-cloudflare/commit/134df93b0fb081435dd3b57c6934379159fdd209))


### Bug Fixes

* 'created_on' API inconsistencies in LB pool ([e8d34db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8d34dbd324f9fe6b71c746514c76a92e69c8e34))
* access application model schema parity ([94311b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/94311b7b7e4e97928242bb7b14bd43727662ac7e))
* account_member tests ([832a05a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/832a05aab2ddbfaf1deeca43aeac411d5a5edc80))
* api_token ([3846ff9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3846ff980577ffce0064dc64e25f63a97ebe7dab))
* drift ([6e5659d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6e5659d04827587c6d33620cb5136b8e6f3f6516))
* dynamic type validators should handle int and floats correctly ([5ae1226](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5ae12261d7b25e7d0d3f18c38682ee00e0a45a2a))
* enable account_member tests ([d3ba4ab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3ba4abacc620728e1c8edb19b3d993841b0e50e))
* encoder crash for nested nils in dynamic types ([91e9782](https://github.com/cloudflare/terraform-provider-cloudflare/commit/91e97825e1abc5ef9fb5bc95127a69a0435ce5b0))
* enhance DNS record sweeper to prevent apex domain test conflicts ([4be8681](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4be86812c652245ad041961b49e417f0f3544a77))
* grit patterns ([736b9ac](https://github.com/cloudflare/terraform-provider-cloudflare/commit/736b9ac6b2361249529d12cc3ad335a6d81a5e57))
* grit patterns for dns records ([6e0785e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6e0785e7bc4d4d661dd5ae88ad2617759fa7f9c9))
* handle empty tags ([35ab9f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/35ab9f7d0de7b0c9439df56ea6fd1490a405aee1))
* implement migration for ZT IDP in the migrate tool ([df2289e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/df2289e7cf65d2f1e7c96bc25671fdcc2d417fc3))
* **load_balancer_monitor:** Fix detected drift on refresh ([024f015](https://github.com/cloudflare/terraform-provider-cloudflare/commit/024f015b2ea7d39861b20249703dd01e3db93525))
* **load_balancer:** Fix detected drift on refresh and update ([3d06582](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3d065828d926b39d839a62c8c0db52885a536f94))
* migrate for managed headers ([8398f4f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8398f4f79c8918919c0717718582a2fb7f767673))
* **migrate:** fix incorrect setting name for `0rtt` ([4943ca2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4943ca2fd8950d2eef892ce298cc9df7c7477f42))
* **migrate:** fix zone_settings migrations ([d24ce96](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d24ce96b219cdd6ceaa0fec81c0fb3eeef47ba37))
* populate computed_optional collections from API responses ([d6c64dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d6c64dc8775540dc5179cf1d655b374fba0c22bc))
* properly handle null nested objects in customfield marshaling ([61c808d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/61c808dae7a7985dcb5b972b1db3c8168abecab8))
* remove state upgrade ([c6716e7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c6716e7f2178f9599141fc06a9b841bb4e74c260))
* resolve test failures in magic_wan_static_route and ([11d91df](https://github.com/cloudflare/terraform-provider-cloudflare/commit/11d91df9134d2d72797449f1c51509d53924d1aa))
* resolve type mismatches in zero trust access policy and application migrations ([8e3b4b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8e3b4b692b78037cfa9a5719d2d3cc99eea5b438))
* **test:** Add plan check validations to cloudflare_zero_trust_tunnel_cloudflared_route resource ([c94fb5b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c94fb5b3ee300dcede4b3af592731df9131eafa7))
* **test:** add planmodifiers for dynamic type ([eaf6f17](https://github.com/cloudflare/terraform-provider-cloudflare/commit/eaf6f176f27a9b16cb0a472eb20c54ae0019a951))
* use planmodifier for zone_settings value normalization ([468f59a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/468f59a09e3e39bca65766ba4b5748db04f3f3b8))
* wire up migrates ([54c3248](https://github.com/cloudflare/terraform-provider-cloudflare/commit/54c32485e50c001f0dc412f5ca863a8f41c8de35))
* zero trust access policy acceptance tests ([4804ca3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4804ca3281b66dff01379d7c4b35b5c810356dc9))
* **zero_trust_device_custom_profile_local_domain_fallback:** fix recurring diffs and add acceptance tests for multiple domains ([c8e790a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c8e790a4cd9915186ae2a7ec325ba4b78330b63d))
* **zero_trust_device_custom_profile:** fix recurring diffs and add acceptance tests ([2b358e6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b358e677e10cdfae0cb1b1a8d555d6fc2ade0ac))
* **zone_setting:** remove grit patterns for `cloudflare_zone_settings_override` ([3b6edda](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3b6edda0c8bc158e869813445d45d814c7e07387))


### Chores

* add load balancer acceptance tests to CI ([b9df93d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b9df93d96bd6d4fa782636d1905bcd1005195315))
* add sweeper for magic_wan_gre_tunnel ([e31976a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e31976a31afb78a8af3561c5e4df1ac1c3c7b2fd))
* add tests for zero_trust_device_custom_profile ([cc567a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cc567a80c61bbbd8c09ce722b243c4cbff4d9f65))
* add zero trust device tests to CI ([bcff1a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bcff1a8460a2776db0943e28ae0bc55a0c621227))
* **api:** upload stainless config from cloudflare-config ([9a32393](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a323936b477add84f6f71c96b06037e5ff1c803))
* **api:** upload stainless config from cloudflare-config ([00d150d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/00d150d5bba8878960146edf64fc2155f99df5ad))
* **api:** upload stainless config from cloudflare-config ([ca57125](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ca57125e28e043799f0f75de5869fb32fa37faeb))
* **api:** upload stainless config from cloudflare-config ([416b5c1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/416b5c10501e443a430709c52dc289bf0c1e47df))
* **ci:** skip ZT mTLS tests ([289ce58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/289ce580ed52c572902d1a92f1598d3d5e9f0fee))
* **ci:** temporarily disable migration tests ([e1c5b3d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e1c5b3d4bd525f406ef49ecd4192d614c03adf0e))
* **ci:** temporarily disable migration tests ([8ab562f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8ab562f95690d41c7f5b9b9f5168a8fe09da7697))
* clean up ([0415065](https://github.com/cloudflare/terraform-provider-cloudflare/commit/04150656a680274cb09c7bcde729921a94a7b304))
* comment modified on test ([d403d5f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d403d5f86fe898b11ee7b815eff0c76b07489ff2))
* disable failing queue and r2bucket sweepers ([e1394a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e1394a8d5bd8231c6f80208fb02db03abba88736))
* **docs:** generate provider docs ([d86327b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d86327bbf93b3523bd8d33967a74ace9cc590c09))
* fix ci script ([5991be1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5991be1343d421535d4ed65515ec0c311e0dea0c))
* fix magic_wan_ipsec_tunnel acceptance tests ([5f5b50a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f5b50acd6b558cc74527e15161e74b31db3c5c5))
* fix sweepers for many resources ([c2f66f6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c2f66f627656a4a0b6a59c3387d37cbacc928dce))
* fix TestAccCloudflareAccessOrganization ([8b0e176](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8b0e176dea5760322c2de3cb3c812d09f79c055c))
* fix TestAccCloudflareAccessPolicy_ApprovalGroup ([52ef60b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/52ef60bdca1bce374318734a016659da5557a566))
* fix workers_script tests ([1116eac](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1116eac5d6df11f527f1a5bc8ece1cfe6730d2dd))
* fix zero_trust_access_identity_provider tests ([2575473](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2575473c8c14078ad437dfbf71fadb0a54de7d43))
* fix zero_trust_access_mtls_certificate acct tests ([a00f421](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a00f4214fe89327ed7b8f483162e226338b675b9))
* fix zero_trust_access_mtls_hostname_settings tests ([166517f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/166517fe66cd791f8d4e3bbaa49d6946addb8e4e))
* fix zero_trust_tunnel_cloudflared_route ([0e961f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0e961f705e585734b4b688cd2e0d7c2036e73879))
* fix zero_trust_tunnel_cloudflared_route tests ([3c51256](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3c512563e3cf3bf6d2442224f4239f5eee89e2fb))
* improve integrity test error messages ([2d410d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2d410d2a1d985a9524785c7ace0277aa7b857fd0))
* increase ci test timeout ([5da9eab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5da9eabfbf2565a6e8aa7925c1d575f651da490b))
* increase number of ci jobs ([50ee749](https://github.com/cloudflare/terraform-provider-cloudflare/commit/50ee749e1ad231577b2f5280a52f58e6dc8555ce))
* **internal:** add test rule to lint for dynamic attributes that do not have planmodifier ([a725465](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a7254653b34797b0f754198606bd81a3a9160ce2))
* **internal:** codegen related update ([42115d7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/42115d73fae3501b50d6294ea4da8a9ab4c83eea))
* **internal:** codegen related update ([8feeada](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8feeadadde65c1ed8cd2d8dd91a16ba4612aed5c))
* **internal:** codegen related update ([17863a2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/17863a2c63d739ba555b5e16a131b691f39cf603))
* **internal:** codegen related update ([0bd098f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0bd098f1f95585334dfc53d18e05039d48767884))
* **internal:** codegen related update ([4e59511](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4e59511354cb3aa567e614a0c82817693532bbef))
* **internal:** codegen related update ([55d5985](https://github.com/cloudflare/terraform-provider-cloudflare/commit/55d5985b2b44f137b7c28b6a76215a95edeae01f))
* **internal:** codegen related update ([5505692](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5505692abda33c286d1ce8b477dde524dfb506b7))
* **internal:** codegen related update ([7765dbf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7765dbf1da1fbd10743b3126bbe35702e481b326))
* **internal:** codegen related update ([1748ea8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1748ea867919cb8ee1dd7db097d4b1378fd0e271))
* **internal:** codegen related update ([e0eff5c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e0eff5c3080169222b05858178063469b25c2d2f))
* **internal:** codegen related update ([cde67a4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cde67a40224cc2633d715f3897f0006748fc3ff5))
* **internal:** codegen related update ([5d2fe3d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5d2fe3d3d1b788b25912d8696947439814b2c384))
* **internal:** codegen related update ([a9610ce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a9610ce8a3ef636a6bb069f7fbd70c037671076b))
* **internal:** codegen related update ([90aa9c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/90aa9c5558e8a1e15cb9e3153973e4b55b8af78d))
* **internal:** codegen related update ([2b37de1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b37de16009c5664d9bb4a895f691ef5b9c167de))
* **internal:** codegen related update ([b27f531](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b27f53154750aae21dd4c44eb4dd70e23dd6f157))
* **internal:** codegen related update ([2275aa2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2275aa2d73e17e0f8ae0dd28f98f3e61ae4554d2))
* **internal:** codegen related update ([4cbcff5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4cbcff5fba7acf268ab5c46e79f93914d2dd8f6d))
* **internal:** codegen related update ([2136d08](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2136d086ba77d5107e86f1acb09b5ccead821cb1))
* **load_balancer_pool:** fix test data and skip broken test ([3be61aa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3be61aa1458c0277d01b3be5e06d9dc2cf40912b))
* merge acct GH steps into a single step ([d27710c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d27710c5a7d8842d86024f3a084df981388cf3aa))
* more tests and drift ([f1cc4ac](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f1cc4ac5ea0828f3572908ecf3cc8f46a00679f2))
* parallel ci test runs ([260a5b9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/260a5b924262eb8148137034d8baf8a10ac555e5))
* remove state upgraders ([ca04b97](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ca04b97d91894b16587304dffa968a089eb79dce))
* remove version from schema ([73f2e17](https://github.com/cloudflare/terraform-provider-cloudflare/commit/73f2e1738bb14f3d2e17ba4ef5fa2af23023fdd7))
* run sweeper before any test run ([5bac861](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5bac86115fcb271aa2678814c228fd25ce914938))
* separate acceptance and migration tests ([e466442](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e46644276c1dbeeff03c007d1351ad584973a89f))
* skip failing ruleset test ([7746f2b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7746f2b951e2d32fdaf07dd61e224d7fa39d6122))
* skip failing sweepers ([46810fe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/46810fe3991ced4886bebdfd9f609b4a53793928))
* update all cf-go v5 -&gt; v6 imports ([2cd840f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2cd840f7569c3ea822364017ea44bafd454edd7d))
* we can have only 3 parallel jobs ([588c710](https://github.com/cloudflare/terraform-provider-cloudflare/commit/588c710a12dde2bb78a57b5f02d86acc90f88ed9))
* wire up migrate commands ([b574a2a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b574a2a6250ec3ad514f9e8ff845533daf643dea))
* zero_trust_access_mtls_certificate tests ([07e374e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/07e374e6a32c00a2a0acfba3c805b8528af62347))

## 5.8.4 (2025-08-15)

Full Changelog: [v5.8.3...v5.8.4](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.8.3...v5.8.4)

### Bug Fixes

* **cloudflare_ruleset:** update for consistency with OpenAPI schema ([837da07](https://github.com/cloudflare/terraform-provider-cloudflare/commit/837da07e1faed0507eae00f602414eac1b1e9f05))


### Chores

* don't announce to discord ([1816fff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1816ffff2ec85325ca622ef1aa72a51b6d2ca9c6))
* generate docs for 5.8.4 release ([ae12b37](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ae12b37d0dff361d7aa463f593a94bca01668c2f))
* only include ones that have tests ([ebc40cf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ebc40cf77c63951364b2b186e93ff19f4374faec))
* run one by one ([1185be7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1185be75c053efa6d04a4d57491ac0be861bf2f8))

## 5.8.3 (2025-08-15)

Full Changelog: [v5.8.2...v5.8.3](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.8.2...v5.8.3)

### Features

* **api:** api update ([23eb89c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/23eb89ca43852af44e8d79ef33d2e67a20ab8f87))
* **api:** api update ([57928dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/57928dc81531d22e4474e9c10bf1f061266d1915))
* **api:** api update ([3351a79](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3351a79400fc48efb8e004a74128c8eb6edc4466))
* **api:** api update ([b1afd55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b1afd55b64af81794cc72cb29c1ee30b6a3663eb))
* **api:** api update ([aa5ac4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aa5ac4da877a5503c11aed8077351428c24153d0))
* **api:** api update ([01bea92](https://github.com/cloudflare/terraform-provider-cloudflare/commit/01bea92e8874a836992693794b022db3187ab430))
* **api:** api update ([5bf6360](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5bf6360e1972846b86c0d588fef10c2b3a7cc806))
* **api:** api update ([698d90c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/698d90c0c97a5d555e78fb81aa0f8e4f413fba3b))
* **api:** api update ([c9f96d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c9f96d310305fbe7105ac1e6b996417eebbbca40))
* **api:** api update ([f5ee559](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5ee559e8db0884c4a8705756961a2ab380192ae))
* **api:** api update ([2568efa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2568efa9b5e008964da1f519d98cb56fa0f5af31))
* **api:** api update ([5afa7cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5afa7cb54445617490fd798ec1731cd02fd0470b))
* **api:** api update ([7cd55d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7cd55d34af5aec6d0b88b8572fb2ecae7f421691))
* **api:** api update ([f1b07f6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f1b07f6af7621fe960e6996c50c0ed4c09a4a283))
* **api:** api update ([0e1f55c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0e1f55cf54afc04d042c688b28c1719cd61fd5ed))
* **api:** api update ([535c250](https://github.com/cloudflare/terraform-provider-cloudflare/commit/535c25089d2dd37a8cc7a8fde5d61616d7217687))
* **api:** api update ([dda8106](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dda8106bb8308f171928ad5ba0fa36198c19fed7))
* **apijson:** add `decode_null_to_zero` tag option ([71538e5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71538e50970a1fd7a27bdba5ed1e285eaefcefcc))
* **apijson:** move changes to new `apijsoncustom` package ([f5b955b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5b955bf49c9c9db58e77bf45fca01e121c67058))
* ensure `internal/apiform` encoder can handle "force_encode" serialization tag ([0d8e9ee](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0d8e9ee977b967ba2477009ebee6dc9e4a7313a8))
* ensure `internal/apiform` encoder can handle "force_encode" serialization tag ([840ee94](https://github.com/cloudflare/terraform-provider-cloudflare/commit/840ee944af524f0e9c4cc4d44e86519b4bfc1f24))


### Bug Fixes

* **cloudflare_ruleset:** handle omitted `rules` in API responses ([7f15668](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7f15668a902d6022dfc442c5cbddd45ea63176df))
* dont run in parallel ([9c952a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c952a96faf80693ba606416757be143d6539191))
* imports ([4369f06](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4369f06866b423d9fa73e63295d0caf866b12a46))
* list item test execution and add managed_transforms ([c3c3bb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c3c3bb4265ade2c671e72e54e17705ffda602d8f))
* **list_item:** Use proper pagination from client ([8f692ad](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8f692adf679f76a5e98eb3ab4980876f352d27af))
* regex to not account order ([1e2d64d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e2d64d2f2b129881f3093a8f4b0b789d5f0ff80))
* remove workers_for_platforms_dispatch_namespace default value for 'trusted_workers' ([dded442](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dded442c162e3f4ebf0d897c7d0e824e3c6ba808))
* test assertion regex ([9a0b7ab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a0b7ab492ce3ae6ade8a45a1756f798bbe1bbc6))
* test data ([1cc94d4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1cc94d4adc4e739817eb10ab58cc3cdc7cf99c6c))
* update schema ([6c97aad](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6c97aad1bbc126af84a9444668c213d18c1596dd))
* **workers_script:** ignore unmanaged secret_text bindings ([a3b6816](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a3b6816b2510577f96668d0e29aa04a4c3a74c28)), closes [#5892](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5892)
* **workers_script:** Obtain migrations directly from config instead of plan ([2602dba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2602dbafcaca02718fd44bf54caca268170dd56b)), closes [#5898](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5898)
* **workers_script:** Revert treating cloudflare_workers_script.bindings as a Set ([757b98f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/757b98f9e0cced8e6ae122f620de6454547ea3cc))
* zero_trust_access_application tests ([d7eccc3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7eccc3f536debcf1fd4e9ea7c3d93f628e65c96))
* zt orgs ([fc9f9a2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc9f9a2ee0f342d2c9e0fabd598dd1d41857429c))
* zt orgs another way ([93f95cd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/93f95cdd892af2b003c666eccf744589c9dd744d))


### Chores

* add bot management test ([6114b7d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6114b7d4949d02cd7e5a6cfbd6f116334c241f8c))
* **api:** upload stainless config from cloudflare-config ([68454d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68454d3d788e8e7c146c085893e0f9c85761d92d))
* **api:** upload stainless config from cloudflare-config ([33e7b03](https://github.com/cloudflare/terraform-provider-cloudflare/commit/33e7b03f955c53a84c0d2b1a844fa0f206c4bd67))
* **api:** upload stainless config from cloudflare-config ([bf113c7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bf113c76383636d0a4971b6f1c35d9ab6444b31c))
* **api:** upload stainless config from cloudflare-config ([7b2fe37](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7b2fe37534d7a1438991d0ab262277b23210e9fa))
* **api:** upload stainless config from cloudflare-config ([9c03a0e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9c03a0e95ede74ac0b7c732f9a108120dd8a2073))
* **api:** upload stainless config from cloudflare-config ([9babc23](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9babc23830d1b7e07bd39240b09e2cd70034d9b4))
* **api:** upload stainless config from cloudflare-config ([872f8ba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/872f8baa356b1ddba6f0537b160b26add9a1d83a))
* **api:** upload stainless config from cloudflare-config ([b21ea4a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b21ea4aede1b39fb80c0d48883b4ca6c2f212c87))
* **api:** upload stainless config from cloudflare-config ([5790fb1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5790fb1d073708b2ce8956337b9cf85290b91faa))
* **api:** upload stainless config from cloudflare-config ([2016bf8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2016bf8833a37bbdd96b472467e148dca79dc890))
* **api:** upload stainless config from cloudflare-config ([ef987a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ef987a81819fabfb0656da8723ea45c60eae5ea9))
* **api:** upload stainless config from cloudflare-config ([26c37c4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/26c37c4ff0dc0698e56cca5f5ea2cb230b6775ee))
* enable token-based auth for acceptance tests ([20ea7d8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/20ea7d8d589ec365083f285e9e1a16f00ef4afff))
* extra checks ([3e5591e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3e5591e6e1ca89771f0c5eaf8fa6ae58d79853d7))
* fix deterministic zone names in cloudflare_zone acceptance tests ([1de0cc3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1de0cc36d965601e7213ae1570f5320dc8201a0a))
* fix TestAccCloudflareAccessIdentityProvider_OAuth_Import ([82c589a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/82c589afe24bfb2b624820f5f160a7c971ebdf7f))
* fix ZT mtls certificate acceptance tests ([cda1128](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cda11280401a989c7f1a2f7c09fff8297af9007a))
* **internal:** upgrade cloudflare/circl ([2df21d4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2df21d4ee0520b977da8c151c6cc94ad41c4ddf2))
* modernize zero_trust_access_mtls_hostname_settings tests ([0a0556b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0a0556b968e2c9bc2514020160de063b5ae760e9))
* new line ([1c0eb79](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1c0eb79c19759c2c7f3f37ff6389e02766d993e9))
* no-op plan ([db1b335](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db1b335a4f218e140eefad3cf7a3f3d6afa40ad8))
* remove line ([1adfcc1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1adfcc100ebce5d31b9907a27228b6dfc83af237))
* Revert "enable token-based auth for acceptance tests" ([d13f98a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d13f98ab90cebb5e7cc53238da2125fa63fbe85a))
* skip api_token test until we configure CI to support token-based auth ([4a84a94](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4a84a9463a15777ae534524ba45519ca7b3ff2bc))
* skip flaky zero_trust_access_application test ([8107984](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8107984ca4500e6b81205358cce062cc74d7b184))
* skip test failing due to inconsistent apply with sensitive value ([9336894](https://github.com/cloudflare/terraform-provider-cloudflare/commit/93368943a3792a8b69f5f4860c85870fe47a5dbd))
* skip unicode 'zone' tests and extend sleeps for mtls certs ([9871bfa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9871bfa9fbfec3074dc9063f8899181df072dfaf))
* sort imports ([dc52145](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc5214537b12ef7700ce2b349d42f0fa45d0ae7d))
* tidy up ([49b1243](https://github.com/cloudflare/terraform-provider-cloudflare/commit/49b1243c2f0e5a4dd6ef71c414c28f35e26e2670))
* uncomment check ([b24e573](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b24e573cf0e104bc45b6f054b1ec8827fe4a5bbb))
* update @stainless-api/prism-cli to v5.15.0 ([7534488](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7534488a919e10a96b9ddb42c7ef54460f1447f1))
* update ci list ([13e5c5e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/13e5c5ec4c83323f6b91e8dd48f21661e98ff7ba))
* update managed_transform tests ([ad82abe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ad82abeb5d5de428634b8f8af429be4c2500e7ac))
* **workers_script:** use resourcevalidator.ExactlyOneOf() to ensure `content` or `content_file` is provided ([e9d000c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e9d000cd2511550748d15bb0313e19c1517d32c4))


### Documentation

* add a warning to workers_script ([5aa1016](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5aa1016b87364f8f2a017463e7825dc8851bd83e))

## 5.8.2 (2025-08-01)

Full Changelog: [v5.8.1...v5.8.2](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.8.1...v5.8.2)

### Features

* **api:** api update ([54b3c10](https://github.com/cloudflare/terraform-provider-cloudflare/commit/54b3c1027b30ae9ad6f388a64ea04941faea6774))
* **api:** api update ([8666096](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8666096e375e55731b286f30f3b722ad70761c1e))


### Chores

* update model for zt gateway settings ([a084e1b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a084e1bfbf36e6fe1b3799d585d20542f792c0f0))

## 5.8.1 (2025-08-01)

Full Changelog: [v5.8.0...v5.8.1](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.8.0...v5.8.1)

### Bug Fixes

* custom page tests ([bd3c296](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bd3c2964bb65d76eb8a03f74599483afba9854b9))


### Chores

* **docs:** generate provider documentation ([21a9cd4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/21a9cd43377897586295c4497f4e9f644ade1fb7))
* skip flaky tests ([890d693](https://github.com/cloudflare/terraform-provider-cloudflare/commit/890d6936dc5be3291b8800215d7ec01876e884f0))
* snippet rules not ready ([e2312e2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2312e242230a64bb012dadba9ff9d419eb1d606))
* update zone name for tests and not in parallel ([d3b0d00](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3b0d00859c7fc2e8e1fc36133fcc8872e11517d))

## 5.8.0 (2025-08-01)

Full Changelog: [v5.7.1...v5.8.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.7.1...v5.8.0)

### Features

* **api:** api update ([b2e32df](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b2e32dfe6cc227005faefc86e38e757d3e90b884))
* **api:** api update ([daa1cbb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/daa1cbb2d29f28d26fa2d231542b629101ba9b77))
* **api:** api update ([5dc00ef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5dc00ef9e173b0a40b304c7800e0824030e2ac80))
* **api:** api update ([b8aac5d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b8aac5d0f6668e7948964aa6168c9e362a8dee86))
* **api:** api update ([b490b66](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b490b66019d66abd4224f823f0b516a42779c797))
* **api:** api update ([2c8a947](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2c8a947b1430a791d2b976232849e7c5c24cd76c))
* **api:** api update ([b8768d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b8768d3610b4ffc4d0eb741fdef7f6d6464c34b4))
* **api:** api update ([1efccfd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1efccfdd2167e268b35fba7a58b819e738fc47c5))
* **api:** api update ([2e1f574](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e1f574bcc13e9bdc333c9a18e935c24484763d8))
* **api:** api update ([2333c4b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2333c4b7d88398bdd6de860990b859b7d084f467))
* **api:** api update ([7d07c43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7d07c43f57017cb352749b695c67a2af2a5801e1))
* **api:** api update ([5a661e8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5a661e884b4f1f24e093adc2aee448870de1ed04))
* **api:** api update ([6f54311](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f54311b6347090929eed0298f915c94d152bb1d))
* **api:** api update ([f2ac6c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f2ac6c523c99974f01ea95290ffbdc8039e3fd23))
* **api:** api update ([cf7e91b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cf7e91b3ebc9f3eb1f83e4f08ff1abf6f2cf1f2c))
* **api:** api update ([04a9002](https://github.com/cloudflare/terraform-provider-cloudflare/commit/04a9002d0c93150390a868586504f8961775ac98))
* **api:** api update ([1a0fdc5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a0fdc58cf03a26a6ee4629dbf8612d3a842aea3))
* **api:** api update ([d42b931](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d42b9317dc1d95f5e4129df3ad575e83bb833b1a))
* **api:** api update ([5280cf8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5280cf8ddfb72a092b4da76308c3412615161855))
* **api:** api update ([4917756](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4917756f413572a0418eaadbd90c5a8dbaa36af5))
* **api:** api update ([7fbafd5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7fbafd549f90c061538ad3d8a0483f0a26ceb9e0))
* **api:** api update ([709e28a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/709e28a0e387be362cc7035509fca3491a6d0515))
* **api:** api update ([9773002](https://github.com/cloudflare/terraform-provider-cloudflare/commit/977300275fae6f492db658064193c0b029036621))
* **api:** api update ([3120f94](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3120f94b6af470eee1c50468cadcc5786268e381))
* **api:** api update ([1b34fa3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b34fa356b39c8055484ff62186b3aff630417f6))
* **api:** api update ([911115b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/911115bc18ff3c25b4fb75bed24168b0948dbe8f))
* **api:** api update ([5086198](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5086198edf30689cb686dc8026f0e6bdcbe91bf1))
* **api:** api update ([83f8053](https://github.com/cloudflare/terraform-provider-cloudflare/commit/83f805386a960ad5f4ebb8d22cf871d66c20d602))
* **api:** api update ([c1de739](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c1de739e65e890f54ebfaf5bd94629b11665ba5c))
* **api:** api update ([9742219](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9742219b61bbd0c0de04f57213600897ce1d713f))
* **api:** api update ([c9f52e4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c9f52e4600e572a6829c94cc5fb269fa35a7ef85))
* **api:** api update ([62dab18](https://github.com/cloudflare/terraform-provider-cloudflare/commit/62dab1834ac5f47989fc07f0d1f3b73a8270c015))
* **api:** api update ([5e36c02](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5e36c02738eefc8096ef0c9a0a9bfa582aaf217d))
* **api:** api update ([2e75b98](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e75b980f470c14d395aa12f7dbddbc051b5ea59))
* **api:** api update ([eb4c002](https://github.com/cloudflare/terraform-provider-cloudflare/commit/eb4c002ff1f4c7a5ca9d79ec6d36000ae504b3ee))
* state migration for custom_pages ([d826ae8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d826ae8ce61833362b757f741293671fae998883))
* **workers_script:** add support for `content_type` attribute (necessary for Python Workers) ([21b235a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/21b235a485f6df13874b73ccf9af05ea2e15ce05))


### Bug Fixes

* add version back to extended email settings for gateway ([81df358](https://github.com/cloudflare/terraform-provider-cloudflare/commit/81df358568f4a94ade380a545ab627b5e3975090))
* initial snippets acceptance tests ([e2d6a0e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2d6a0ea86269235508a092e1593a12c071976a6))
* list item model ([b64e4fc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b64e4fce8d400f7a0411ab79dcacda930e7dd379))
* make gateway settings host selector + inspection optional ([a5cbc80](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a5cbc8076e839aff87d0db21b85fd01f27322f11))
* **workers_script:** add missing `workflow` value to `bindings.type` validator ([cb5d655](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cb5d65508aae48401431fc975a5f8b8f47cde6aa))
* **workers_script:** fix nested computed attributes under `bindings` causing unwarranted plan diffs ([ab9dd38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ab9dd38462cd068e2674d48e02512dbe79c18601))
* **workers_script:** handle `text` value for `secret_text` bindings during refresh plan ([3ff7ac3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3ff7ac3ba2ad2175fb7db1701b2ee9422bb0917a))
* **workers_script:** handle empty `placement`  object during refresh plan ([ee64d6b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee64d6b1bd2e1c072eb02a1d19544aaf554655a8))
* **workers_script:** mark `migrations` as WriteOnly ([409e81b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/409e81bb9fd745ee48f4e07c98eb61fd9879c3b5))
* **zero_trust_gateway_settings:** remove leftover import ([1c3c135](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1c3c13570cfe8df4ae81eca012487275e67864bc))


### Chores

* **api:** update composite API spec ([c1d05a1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c1d05a196f7e94e8aaa2faaf2a562ad7ca2ed9b3))
* **api:** upload OpenAPI schema from api-schemas ([cb6f526](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cb6f526074c22aebd1382f242c32de96979c0d4f))
* **api:** upload stainless config from cloudflare-config ([033cf6d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/033cf6d66754c846e20ae24f240d1ba4d0e8ece3))
* **api:** upload stainless config from cloudflare-config ([d30fb7f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d30fb7f8d014874036260584fb8c978d8a9b2b29))
* **api:** upload stainless config from cloudflare-config ([59efb43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/59efb438e445774ba15aff7adad2f7129e6b29d2))
* **api:** upload stainless config from cloudflare-config ([0ac93c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0ac93c5e74d6a99ae6cd6e91e4b7fd7dbd8680bc))
* **api:** upload stainless config from cloudflare-config ([91aa4e4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/91aa4e4fd303e18304810cb2993946eb9311e38b))
* fix code formatting ([ff642bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ff642bc4e862fafc453cfb713fb938451457cdc2))
* housekeeping ([a42c70a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a42c70aaf496dc79eefc2efe1b7c6083f31e0ba7))
* missing import ([629f60d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/629f60d769360805c21d8fdfb8c2d0970b331a07))
* remove redundant newline ([2a4b8f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2a4b8f765847c25dd9d74d4454714d4fac370951))
* **tests:** Add acceptance tests for WARP Connector resource ([62f19ce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/62f19ce7e89785fecfd3208bdc6f5691c279bcc0))
* **workers_script:** add import verification step to test cases ([f89916c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f89916c70deac4dcd4fa1597819815d97efe7333))
* **workers_script:** update tests to replace legacy check functions with state checks ([6181db3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6181db3438eec8038f6c4070163d249162c40d2e))

## 5.7.1 (2025-07-15)

Full Changelog: [v5.7.0...v5.7.1](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.7.0...v5.7.1)

### Bug Fixes

* **api:** Fix update/read path parameter for zone subscription ([c763edf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c763edf824d7c624a5386d5d0c9c4f10dc7dbb15))


### Chores

* **internal:** version bump ([5d9412c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5d9412ca93034ed24c449ccaa4b5347680d26b99))
* **internal:** version bump ([93e874a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/93e874ac0a2edfa8d72e4eda2b53b1c0cec8c760))
* **internal:** version bump ([5fa684d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5fa684de7c86bd7774940ad8ec191331f769fd6f))
* **internal:** version bump ([68327f2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/68327f2dde2b2ce5b4178c09d9d01c305335a950))

## 5.7.0 (2025-07-14)

Full Changelog: [v5.6.0...v5.7.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.6.0...v5.7.0)

### Features

* **api:** Add 'zero_trust_tunnel_warp_connector' Terraform resource ([204d752](https://github.com/cloudflare/terraform-provider-cloudflare/commit/204d7529af116f0d3104da3de798142e8d75917d))
* **api:** Add DELETE and POST routes for Magic Connector ([b3c8c0a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b3c8c0a2ec82ae06423277f1f23376d61799d0e1))
* **api:** api update ([85a1a2f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/85a1a2f3bdab52cb663d81bee0012cae79a6d560))
* **api:** api update ([c20c04c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c20c04cb8992b98e94a62b507a9933040f8c5e1b))
* **api:** api update ([f936dc9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f936dc9259deecfff96a63bff7deef5d03d8d645))
* **api:** api update ([a5634a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a5634a8547f3f9ca714f000d2f8511cdfcdc4bb6))
* **api:** api update ([d7e118c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7e118c41bd6187321e7346075699fb8c0dc0125))
* **api:** api update ([7a1200e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a1200ebe55933f6dcadb0c198cfdf3250c2b75d))
* **api:** api update ([97ea6d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/97ea6d6100946c3c5581e256b71ddc2403613677))
* **api:** api update ([75e1515](https://github.com/cloudflare/terraform-provider-cloudflare/commit/75e151546d902d622e91ce7a2713c5efa89b92b9))
* **api:** api update ([223c0ff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/223c0ff085394facb2b22ff958429c2571601a75))
* **api:** api update ([7e9304b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7e9304b15fb637574a72729b76596cb8e536070e))
* **api:** api update ([b98281d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b98281d9d84966e39f0b3b1c96d7d0d652c676f5))
* **api:** api update ([d94fb1f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d94fb1f2b8959a0893066fdb4ff28e624d98224e))
* **api:** api update ([7861f45](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7861f45da923902abad48efaa1a02ec3dab5989e))
* **api:** api update ([534cc05](https://github.com/cloudflare/terraform-provider-cloudflare/commit/534cc05eed010839f29d05c5c7e7df843ebb4938))
* **api:** api update ([39676a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/39676a0de28e699286354db0cae59ba083860ebc))
* **api:** api update ([c449ded](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c449dedccade8be6fb9c8b1ac02e293bb9cae709))
* **api:** api update ([de3965a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/de3965a07b17feda13b99b91e2330e22803288d7))
* **api:** api update ([ab0e41a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ab0e41a010524202c0efb3558d369b4c666cb744))
* **api:** api update ([c8168f2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c8168f2009890904f7aca4ac189a632d379885a9))
* **api:** api update ([371b58e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/371b58e01eba0b306b193697f5d99660ff7e064e))
* new option to send computed values back to server ([2b9c5d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b9c5d5bc52f6da073a288142e7af58187f4422f))
* **workers_script:** support `content_file` and `content_sha256` attribute pair as alternative to `content` ([6c850b0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6c850b0bb397f4bb18abd51d7ddfe1d46575fcec))
* **zero_trust_dlp:** Added individual resources and new routes ([2b7185f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b7185fce3cf74d1a8dfedc459dcc2fc2fa351cb))


### Bug Fixes

* assertion ([58392a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/58392a5141bc557a14988e8ea822b144002b1fd5))
* ci jobs ([8fd4d84](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8fd4d8471e0d72caa393e2b52d3c7b3e27db6574))
* ci run setup ([b38f788](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b38f78828f0da27815a86e68f21d5a79bc23f95d))
* **ci:** release-doctor — report correct token name ([87e54a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/87e54a55bf0e7cde2bc1ee44960d58486a357ebc))
* **logpull_retention:** Fix Terraform ID property ([de3811f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/de3811f0f618118051b3a473c2b1042b95321c65))
* null nested attribute decoding ([5ba7d5b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5ba7d5bf759b16d487c3279c6d4df825092b253a))
* **terraform:** strip leading/trailing underscores from attribute names ([e00ca4b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e00ca4b1d504633efcb7c6d47e8b8664b149b5f2))
* **zone_subscription:** Fix incorrect path identifier on Update and Read ([e00223d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e00223d0a3d7db40fe65c7ef73111059993c03e4))


### Chores

* **api:** Specify default value for Zone Lockdown 'paused' property ([808598c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/808598caf1d203221c1dc9670582d8c03c42da6f))
* **api:** Specify default value for Zone Lockdown 'paused' property ([072f9f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/072f9f7f18d637c2294e78dd0487dccc75fb5bbd))
* **ci:** only run for pushes and fork pull requests ([df566b9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/df566b932510480199bd777106fdc97a727c7450))
* **config:** bump cloudflare-go ([8671c9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8671c9af5f95a6e689aff515b3d96d11eb1a0ebd))
* **internal:** codegen related update ([70ba827](https://github.com/cloudflare/terraform-provider-cloudflare/commit/70ba8274bdf4037784f5bbc0aa0e83afd33506d9))
* **internal:** codegen related update ([f29c24a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f29c24aa73da1644f611a2b58d266c930f938f93))
* run steps on failure ([c7360a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c7360a56e7a4c305b63431907b59da3d84a8a542))
* skip flaky list item tests due to rate limit ([85f5b78](https://github.com/cloudflare/terraform-provider-cloudflare/commit/85f5b780a075d6894a47722ddb1882232ddc078f))
* skip flaky list item tests due to rate limit ([c980fdc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c980fdc0553bf89ec7571a36c9a0786fdd49c009))
* **test:** Fix acceptance test runner ([c964479](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c9644793b092dd3e2f60c6f9abdb616c3542e49f))
* **test:** Fix model parity tests ([52446ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/52446eced05e34eabba827971bdbad4f00e146b6))
* **test:** Skip GRE tunnel tests ([7278846](https://github.com/cloudflare/terraform-provider-cloudflare/commit/72788461f0e4a0727670966585f877f22fea2ccf))
* **test:** Skip magic tests when we don't have the right environment ([c8c505b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c8c505bb8002907211224c674415f65506e1f653))
* **test:** Skip some Access IDP tests ([9e13c6f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e13c6fbf5aa7d5ef694ea6786b0714614337ebb))
* **test:** Skip some rulesets tests ([263f43c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/263f43c31c6d0ee82f61e485a387ea2228d2ca50))
* **test:** Skip TestAccCloudflareAPIShieldBasic ([b5a6ba3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5a6ba3a7da82463b4d5147bed9ebd0871a8c0e8))
* **test:** Skip TestAccCloudflareAPITokenData ([11f728d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/11f728d8a47c330c5122f1997cc3713526af36d5))
* **test:** Skip TestAccCloudflareTeamsList_LottaListItems ([7e6c4da](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7e6c4da4777897892b8c746ab177f729f9b4a264))
* **test:** Skip zone lockdown tests ([bfe7436](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bfe7436517c2a58675dcc7fbee5b026cb2b42d5d))
* **test:** Skip zone subscription tests ([a533043](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a533043ab06e32c147f6574fbf5d589722ab3acd))
* update docs ([884dcd8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/884dcd816133381abf6c4edf309f56e36c80b02a))
* update docs ([cd30cb5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cd30cb5f311c9941bf58b9c901c8317cd4776084))
* **zone_subscription:** Fix ID property configuration ([b762cf2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b762cf2484097dbe26d01b9a7dff99d90a1bb4d2))

## 5.6.0 (2025-06-17)

Full Changelog: [v5.5.0...v5.6.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.5.0...v5.6.0)

### Features

* **api:** Add IAM User Groups and AutoRAG ([56dcaf3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/56dcaf31aebbfe8ae99f87c2a0b2fe38d83ecde8))
* **api:** api update ([f5f9fab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f5f9fab8354232c7033163aad87cee972e847d07))
* **api:** api update ([16f59b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/16f59b7ca68f59576d0f1a4afdd419f903bc9ac2))
* **api:** api update ([38cc34f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/38cc34fea8002cd1c5c83617f60c46c2de01ca06))
* **api:** api update ([87dfb9f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/87dfb9f059e9984aa36b5e97d09bd155f1370c37))
* **api:** api update ([9d65aaf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d65aaf7d3e22702c79aba0ed4479116910aaf3c))
* **api:** api update ([db05444](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db0544438c9f56acf0c4be302789daa0e0f991b7))
* **api:** api update ([d373046](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d373046d31d1b8c58551be0a0337e0352a63ff6c))
* **api:** api update ([7f788ec](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7f788ecb6311aff5bc6fa9bf99afa3bf6c6fd777))
* **api:** api update ([c412337](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c4123371db02b61b2d7cbbdf9c863797153f575e))
* **api:** api update ([39a8871](https://github.com/cloudflare/terraform-provider-cloudflare/commit/39a88712180bf132d3ebf9b81983677af9bc21e5))
* **api:** api update ([7af8e9c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7af8e9c1d92dfaeb220cb971570ce81f555720f1))
* **api:** api update ([10873d4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/10873d41e2f2ab50245265a2d512a500e2d3f264))
* **api:** api update ([b298896](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b298896494cc12a127083f12796bb27629225ce0))
* **api:** api update ([2f3e6c2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2f3e6c25eed2d21a0464f3f795ea073d802791a5))
* **api:** api update ([280f050](https://github.com/cloudflare/terraform-provider-cloudflare/commit/280f050d22739a9351e58a76f9f9432bd85e3fd8))
* **api:** api update ([df6b8bf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/df6b8bfe1635c5b1809564a15c6ed0c1a39375f2))
* **api:** api update ([557f7d0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/557f7d0a022fdd6dbc0da0bae8e0ac2325fbe831))
* **api:** api update ([b53ba74](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b53ba74d449817f94ae28d437c732b1a5692d68f))
* **api:** api update ([971ca4a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/971ca4a579c8a50655a1afc9719d191c4a99e3f8))
* **api:** api update ([ef92b4a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ef92b4a7e62baac6c4dfe58064f87ec44cd12549))
* **api:** api update ([79649c8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/79649c8ba974f224895979b8f2d4ee42ae8d297c))
* **api:** api update ([dbc97dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dbc97dc1e0008c9f56a9e1c3bc5c99d8efdff7d1))
* **api:** api update ([d09845d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d09845d053e978111aaa8a9c180439fdeb32ede3))
* **client:** support environments property from Stainless config ([2e9ad1c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e9ad1c69772948c177c9620953d195a1acdd7ef))
* **schema_validation:** add terraform resource mappings ([e2f968e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2f968ef5ed32c60b9c06c66ffb1cedc790221ac))
* support import when resource ID is in read method request body ([500f710](https://github.com/cloudflare/terraform-provider-cloudflare/commit/500f7109d0422f084c30944c81685db9ce61db1a))


### Bug Fixes

* add missing properties ([a21b2bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a21b2bc74d9287580756feb9479075acca78e65b))
* **api:** Update zone subscription paths ([e021998](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e02199818cc018d91895786add49c977d33aa21a))
* page rule panic [#5577](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5577) ([a3c643d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a3c643d61dd8bbd01bcdc4e7a588214da5ebe612))
* **schema:** better support top-level arrays in paginated responses ([e6331d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e6331d547930bcafffbbcb31a1ab2dc1b53e5b73))
* **static_route:** API can accept single routes now ([ac52503](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ac52503e070b2a448130c638eff82be5c3814d72))
* **WDAPI:** Fix 'id_property' for zero_trust_device_default_profile_local_domain_fallback ([c409fd4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c409fd42a57a6f32f09e27067e2f23e7497d8e03))


### Chores

* **api:** TTL is required on DNS records ([6f3f1cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f3f1cbb648aed9944084869630cb6b9c48eea2e))
* **api:** Update Go SDK version for Terraform provider ([27e835d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/27e835d0840e136c5689abebd23d43f5bc690329))
* **api:** Update Go SDK version for Terraform provider ([082cf15](https://github.com/cloudflare/terraform-provider-cloudflare/commit/082cf158dbed6ae55c81d9a2df602aacda2830fe))
* bump deps to avoid GetResourceIdentitySchemas errors for Terraform CLI v1.12+ ([7bceb8f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7bceb8f53cdf0388c096c0c64c735856112a0ea2))
* **ci:** enable for pull requests ([779c686](https://github.com/cloudflare/terraform-provider-cloudflare/commit/779c686271ad03cede6b31c60554de91988a8e56))
* **internal:** codegen related update ([24c025e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/24c025e3c38f0df5e8ed4aea6f099e95b1e19dbe))
* **internal:** codegen related update ([50b8e95](https://github.com/cloudflare/terraform-provider-cloudflare/commit/50b8e95ee60a57f72ff7c34f2c7380b4f4f8442e))

## 5.5.0 (2025-05-19)

Full Changelog: [v5.4.0...v5.5.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.4.0...v5.5.0)

### Features

* **api:** api update ([3823991](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3823991bdf2561b82afbd76849d71fbd98025295))
* **api:** api update ([831ce6c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/831ce6c24b0e7cb18ef46778ffdd48fa3960bc67))
* **api:** api update ([0a3e31a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0a3e31aa6bb542e3427890709ef85475b007f638))
* **api:** api update ([bf3db8c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bf3db8c3318403b060a963e59ce9db03ea997ac3))
* **api:** api update ([315bae3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/315bae3b23e2bf1656138e6f81e2e8563a3997a2))
* **api:** api update ([e8e9f5c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e8e9f5cb9628b3943ccc27e53ec35f881da5dc21))
* **api:** api update ([9275cc7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9275cc75dac0532e35db12e043f4cfdd2430d6ca))
* **api:** api update ([9d82124](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d82124eb6842d62fe627d6e06d2525843e2fc9a))
* **api:** api update ([736d315](https://github.com/cloudflare/terraform-provider-cloudflare/commit/736d315532c10bb1dfa40d5574ec7b9141c7e0b7))
* **api:** api update ([0ae0584](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0ae058461089ed5221185bc16785fbb9f96230ea))
* **api:** api update ([5f69644](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f6964437e80f9adfba6414881a3b447a18a14e6))
* **api:** api update ([275a65a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/275a65aefe0fa1df3b4cecfef35f78e9a84e4fe9))
* **api:** api update ([51f1988](https://github.com/cloudflare/terraform-provider-cloudflare/commit/51f1988cd47bc99c0d73900f07f61bfcbf66272b))
* **api:** api update ([7a7cce4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a7cce48cce9473c1a8b527dc8d6ce9f4096335b))
* **api:** manual updates ([d7f399a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7f399a4f99b00151324d8c6472dc2c3a1b64faa))
* **api:** manual updates ([8356001](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8356001958c00b90444088da60469d2e54be0bde))
* **workers_subdomain:** mark endpoint for upsert ([f3cd535](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f3cd535e42af93e559e48721f7c0ee51a731b179))


### Bug Fixes

* **api:** fix path placeholders ([0964b9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0964b9a2b34439d9ac45c26adddf7b106f02f5e7))
* **cloud_connector_rules:** define upsert operations ([db4e2aa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db4e2aaae61bdd99bf4d02272ba5c1be95681065))
* **cloud_connector_rules:** fix nested schema bodies ([#5559](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5559)) ([64edb91](https://github.com/cloudflare/terraform-provider-cloudflare/commit/64edb916cd0b35a1c8587069d2384cb233556067))
* **cloud_connector_rules:** remove outdated warning ([#5560](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5560)) ([d7d6ad0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7d6ad0b1bfff08027025520e0c89aee15c70071))
* **cloud_connector_rules:** reuse zone_id for anchor aliasing ([bb5cbf3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bb5cbf397f7f5c4c65a75982d73efe080abd66b1))
* **cloud_connector:** alias read methods ([2df31d9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2df31d930ee0fd1a03ee34d2df486484e4d0d95d))
* **cloudforce_one:** fix ID typings ([#5556](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5556)) ([8f30924](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8f30924a22cc651808db8c6e1ecef1c2322beede))
* **cloudforce_one:** fix ID typings ([#5558](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5558)) ([6259852](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6259852a197bee1e7e38f90e1c3a658119f57a9a))
* **docs:** ensure schema docstrings always match the correct schema ([120b0c0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/120b0c01424bb58b0c65e25d7ed1f355317ec223))
* **internal:** more consistent handling of terraform attribute names ([69f06bf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/69f06bf65fdbaaffebddbfd83dcc93423ae55ce6))
* only unmarshal attributes that exist on the read response schema during refresh ([6521853](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6521853a6ac417cbef0c7f997d1bc54a50c7e72b))
* page rule issues ([#5601](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5601)) ([6316235](https://github.com/cloudflare/terraform-provider-cloudflare/commit/631623525caca3ee07ecee03580bfba3755742ed))
* **r2_bucket_event_notification:** add missing queue ID for params ([#5594](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5594)) ([eae6622](https://github.com/cloudflare/terraform-provider-cloudflare/commit/eae662283900dc85c2cae427a3b28026afe89081))
* **r2_bucket_event_notification:** revert incorrect schema update ([#5593](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5593)) ([e86f933](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e86f93338babb2ce01d456c65bb5a351e66b3fe0))
* **r2_bucket:** fix handling of r2_bucket params ([#5562](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5562)) ([aa7ba48](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aa7ba480e0b73d2da0771b96b44b4541ef4d0961))
* **r2_bucket:** support editing attributes in place ([d0f7581](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d0f7581c7d98c4ee002101efc343d86d1eca497e))
* **release:** update README and version correctly in release PRs ([5b2c9d1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b2c9d14452a21c8a9ac8281f9752e713607ca53))
* **schema:** support ID parameters on post bodies in addition to path params ([11b8aa9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/11b8aa96781aa3455d58468b74371f34c4661133))
* **workers_script:** Fix refresh behavior and state thrashing ([#5544](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5544)) ([5c9e166](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5c9e1669d6cf2f6b32d1e128a1aa39f0cfeca696))
* **zone_setting:** update model tags to match schema ([#5597](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5597)) ([624db57](https://github.com/cloudflare/terraform-provider-cloudflare/commit/624db57a2c87c5e2f9e2bf934b4f17f352bfae3a))


### Chores

* **build:** update go.mod indirect dependencies ([b808655](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b80865549413d2c014c137971875728737962674))
* **dep:** bump cloudflare-go to v4.4.0 ([4c54318](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c5431893b6b8539296ecf6a4fc9615a035213af))
* **grit:** make state replacements more flexible ([94617a7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/94617a780eb11c92868b4ecb816b65711d93db7a))
* **internal:** codegen related update ([c805fc4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c805fc42fdd9e12b81e078e5654ec65eb1f42da3))


### Documentation

* generate ([#5557](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5557)) ([0aba524](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0aba5249212a7af4a4970f7ebb4cf2c49503fdc8))
* generate ([#5595](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5595)) ([0c68d86](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0c68d86d39b61cd04edbff5370827245fc954ab9))
* generate ([#5602](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5602)) ([9195f35](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9195f35f1cf06c07fa9a9ad9c41c71defe49558e))
* indicate cloudflare_workers_secret is removed in v5 ([#5539](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5539)) ([3c4c46c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3c4c46c46cd1e881e06fd348af28488efcf7e7fe))

## 5.4.0 (2025-05-05)

Full Changelog: [v5.3.0...v5.4.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.3.0...v5.4.0)

### Features

* **access_settings:** add CRUD support ([c09313d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c09313d3357fe19a910855020ee1a9bff60f6ddc))
* **api:** api update ([9d7422b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d7422beb901133703fb5fc20ace4b937e30e693))
* **api:** api update ([05ee8a1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/05ee8a116e2d343abd9b158882d63279962c4995))
* **api:** api update ([550f36a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/550f36ae8b4a50ffeb9ce6cc755b2e1a4a75d518))
* **api:** api update ([0897751](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0897751567801ff21216a06866c591cb3b66619e))
* **api:** api update ([de87162](https://github.com/cloudflare/terraform-provider-cloudflare/commit/de87162926e250fbc7d30899923fa2525ce15542))
* **api:** api update ([399bbf5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/399bbf5b331a6808b7b8f2910020d7de5120297d))
* **api:** api update ([1a38d89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1a38d89c384c0718f748e831b545adeda04f2eb2))
* **api:** api update ([f68c333](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f68c33347d7e855736f39d5c7a1bb3f52723e728))
* **api:** api update ([61f2d9b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/61f2d9bc7061ce55754b3d730d7cfd27e1bd1dc8))
* **api:** api update ([99de9f5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/99de9f5c1a835c636c0e95c27aef11a592217b3e))
* **api:** api update ([4c88979](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4c889794cee12dbfb9ff798779d67737cd9e06c6))
* **api:** api update ([62aca85](https://github.com/cloudflare/terraform-provider-cloudflare/commit/62aca85d8469b518138ab6f48d283275611cda8c))
* **api:** api update ([fb7b3cc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb7b3cc10d2beb9437caf260b23e859e06da6747))
* **api:** api update ([721c070](https://github.com/cloudflare/terraform-provider-cloudflare/commit/721c070b5354592c423c0997133e8a56e43f03f2))
* **api:** api update ([2459c26](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2459c26a31aa1e45ff71bc6e6ea981aa625f597e))
* **api:** api update ([557a5b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/557a5b7b34290ad4f0ca80dd5e60035fda18bcfc))
* **api:** api update ([b126460](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b126460a5bdae6fd765601aec97f91177286de7a))
* **api:** api update ([dba55a3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dba55a39368afad4227eea5ca7482e9224810771))
* **api:** api update ([ea6738b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ea6738b640385658d527938170709e5e7db0ae37))
* **api:** api update ([db2b0b3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db2b0b38aee30a756933fc6d4d8ef5b68dacf36e))
* **api:** api update ([9488f5a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9488f5a92d7d90ccb02828dafa539bce1fe2ad74))
* **api:** api update ([2b58bcb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b58bcb8dcfc4b45c00aa05d285de2b47f831108))
* **api:** api update ([56d1121](https://github.com/cloudflare/terraform-provider-cloudflare/commit/56d112128664d123d9b53ff70b1b7b2c505e0edb))
* **api:** api update ([447c889](https://github.com/cloudflare/terraform-provider-cloudflare/commit/447c889e38797999eb9fbd6d453143245e9e856d))
* **api:** api update ([28964ac](https://github.com/cloudflare/terraform-provider-cloudflare/commit/28964ac1d34ce08f12d865413553659e96cf305f))
* **api:** api update ([f93d382](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f93d382f97c29cb68d956747ff86bb4af685c02b))
* **api:** api update ([b4c15f2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b4c15f290dbe8a641e1600723b6e6d2e1a736e48))
* **api:** api update ([46e8465](https://github.com/cloudflare/terraform-provider-cloudflare/commit/46e846524dfb46818360245c701da0e5569432a4))
* **api:** api update ([6f9426b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6f9426b58725d163906b5357f99124aa0a6c81ef))
* **api:** api update ([5e9f3a7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5e9f3a790cafb8405789e45606aeab3d43108d41))
* **api:** api update ([94b6154](https://github.com/cloudflare/terraform-provider-cloudflare/commit/94b6154ac8f15b0a43c2eb020fd24539d97ed8ea))
* **api:** api update ([9d49cbf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d49cbf2c2b46db10fe59f229769db4d6ce05a32))
* **api:** api update ([6d3c122](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6d3c122eb97eee2dd09acab985607911ccef0be6))
* **api:** api update ([880408f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/880408fec7ce2994791d2db08bfc0eb2dafcfa78))
* **api:** api update ([201b506](https://github.com/cloudflare/terraform-provider-cloudflare/commit/201b506482688caeca8bdf3e804c4291e4e68708))
* **api:** api update ([762cf56](https://github.com/cloudflare/terraform-provider-cloudflare/commit/762cf56b12b53c4d75c31a0692fa753f780681bc))
* **api:** api update ([4058ccb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4058ccb0448df540464633ce2a72eaf86f645013))
* **api:** api update ([3996601](https://github.com/cloudflare/terraform-provider-cloudflare/commit/399660160d6eb5a9822d57f5d8fc8ba9bd0962e9))
* **api:** api update ([9e54b74](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e54b741272e0ba52cdd19187ab1b0ccc47ddd22))
* **api:** api update ([2521b59](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2521b59f61127b68f2ba347d3c8f07172f664297))
* **api:** api update ([5103861](https://github.com/cloudflare/terraform-provider-cloudflare/commit/510386157228dd59b9e3f0d5e6e79992dd412a67))
* **api:** api update ([c491f51](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c491f513c84da57794c90f57473da0bb2b5336b7))
* **api:** api update ([b281833](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b281833f10a48ad815d14608b47dd8da3518e618))
* **api:** api update ([7d108ca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7d108caeba2033e32dce2c6e529512f3822e7c70))
* **api:** api update ([8dc3fc0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8dc3fc0f3ddd7ef627991bf14dc0c9649f04719f))
* **api:** api update ([a3c6c2e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a3c6c2e9f219f4b8313277b0a995bd3aa25a7030))
* **api:** api update ([7fa3ae4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7fa3ae456d7cae62057a194683fda63fd0992f50))
* **api:** api update ([2011650](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2011650ebc677cdb8f3f54982eb757da60430929))
* **api:** api update ([a4d240a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a4d240a959686ed52118a24738b3433f17f1c3d5))
* **api:** api update ([282dfe0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/282dfe00adef2567de0323d0936fdeafc1bef832))
* **api:** api update ([714c307](https://github.com/cloudflare/terraform-provider-cloudflare/commit/714c30782818d0fc542021333d7232d27542ff05))
* **api:** api update ([e81a4e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e81a4e1bb7687bafc161cc5d67049967f9dd4546))
* **api:** api update ([afc42a7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/afc42a7e7b21f2ba5f6222e0680e574664db9474))
* **api:** api update ([00b46a3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/00b46a3a3443b2754569ecb8a133cb341bcad933))
* **api:** api update ([ae3bbc8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ae3bbc81ea5fec27768d2eb75adac6672eff1c4d))
* **api:** api update ([e2a069f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2a069f237c4bce4bc6b65559ef66c1742e16cb5))
* **api:** api update ([33fea7f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/33fea7f23d56a270127530941eca73bce6540947))
* **api:** api update ([bd83e14](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bd83e145b176c5bcea840d282e4c1a76bfb6f861))
* **api:** api update ([a371171](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a371171b3dd0b1e1cea1a826fb07ba6103ad066e))
* **api:** api update ([d01e6cd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d01e6cdb2a8a040da3bad92830ef307e4dfe2861))
* **api:** api update ([20deac9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/20deac914a69bca130f5ede858096a9b05032a72))
* **api:** api update ([b47cdbe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b47cdbe3fa1d9e2ebeb01c44a39c6f4153cfc5cc))
* **api:** api update ([1836d43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1836d438d0227b2b7dcd728bff1cf35fcc9d59fb))
* **api:** api update ([74eff88](https://github.com/cloudflare/terraform-provider-cloudflare/commit/74eff8838da8416cb90c9dc6070ac588660e72eb))
* **api:** api update ([f96a024](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f96a024d82b60b493e54ec0e0717bfb267ed6038))
* **api:** api update ([266ab95](https://github.com/cloudflare/terraform-provider-cloudflare/commit/266ab95bae24b50eb23bd932c8671d723426447d))
* **api:** api update ([2223159](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2223159703fbf0c9bb6fc5efd3d747c7c9dbb7b0))
* **api:** api update ([66336c3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/66336c3c80992c38cc37e42a3df1144357e2f21d))
* **api:** api update ([b047a2b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b047a2b57e886ad66eec68a3971ad6073a909955))
* **api:** api update ([b512134](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b512134a34ca562859eff72de13f2a618bdaf877))
* **api:** api update ([4cd9d9e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4cd9d9e05cac39f15ba1197bbd3c4d9efe5ce414))
* **api:** api update ([0690194](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0690194ef08be22be8a9a330a3afbdbb30171230))
* **api:** api update ([b00f257](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b00f25767c244998376990728e5da04c41bf0d7c))
* **api:** api update ([ce3313a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ce3313afd4f1a2b3c134794c0aaf38489e6744fe))
* **api:** api update ([dad5dfb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dad5dfbb5d055b908d18a0a327394c41727263c9))
* **api:** api update ([a52d914](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a52d9145b927fbabcc748f9b619a40bfe0181033))
* **api:** api update ([2e1f93c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e1f93c3d648c4938bf36d1f6107f074efd1e30b))
* **api:** api update ([1b3f8b8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b3f8b8cbe6818fce24b884f920f0c1768dd8ebb))
* **api:** api update ([ab9640a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ab9640af4e15aa7d6021e23a77b742bec7569cdc))
* **api:** api update ([b636a9b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b636a9b4718f1392ffc93718ff491308fbd01155))
* **api:** api update ([b6214df](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b6214dfefffc54dfa959f5553e7322e442555982))
* **api:** api update ([54dd3fd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/54dd3fdde0131582639adb7ccc8a4b824209e1d9))
* **api:** api update ([158af31](https://github.com/cloudflare/terraform-provider-cloudflare/commit/158af3116f241c1b66e3fba90f9afb9497c5b1ef))
* **api:** api update ([93729e9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/93729e96f70b150c0ba6afbc8b6ae82862d09bf9))
* **api:** api update ([710186b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/710186b7705cfc14b8c12e83989027164ddab564))
* **api:** api update ([4ca4220](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4ca4220eecdc880433f84ef66c03eb2125675b5d))
* **api:** api update ([9941fce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9941fce3b1b805eb235b36a8873e17d1ac65e7e8))
* **api:** api update ([417f2c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/417f2c521253dda3f6dd36832b9450287cabead8))
* **api:** api update ([bc360d3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bc360d39ecd14a333e30118beff64ec659d26cfc))
* **api:** api update ([d2d7ffd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d2d7ffd59ad598c7c8f261efb5acb197dab2d424))
* **api:** api update ([0b10e03](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b10e03069be5b1ebe9f9c501788098038d17dd8))
* **api:** api update ([664ee29](https://github.com/cloudflare/terraform-provider-cloudflare/commit/664ee291d860e03b00d40f3d8a3be0c23f85b943))
* **api:** api update ([6ee341f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6ee341f76431491a068bb8614524733622f4ca8f))
* **api:** api update ([f40bb1a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f40bb1a98c93de9055d46ccd315a5bea0a52885c))
* **api:** api update ([30f0f5b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/30f0f5bb7e163263ce0aafb8cef362cef18ee976))
* **api:** api update ([820bb7b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/820bb7b37fc32ae817b27ba07627d5a9a11a828d))
* **api:** api update ([c85655c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c85655cb71b8ad376cdfba8b4968a5e6a191ddeb))
* **api:** api update ([e4f02f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e4f02f73156ec396c4d9c8eec872bb2684a922b2))
* **api:** api update ([d0a20d1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d0a20d1df27286a40b892e3767836049c8079537))
* **api:** api update ([e2904bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2904bc043549174b5413448e419a27dcfbfa0f8))
* **api:** api update ([71ff2ca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71ff2ca77b90b3197118d775265675807a9a0c51))
* **api:** api update ([82d7f94](https://github.com/cloudflare/terraform-provider-cloudflare/commit/82d7f947c132ca7ec6ab956c7ef188d212aa4190))
* **api:** api update ([04fdbdf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/04fdbdfc6922b8bbaf3cf989bbfca5164bef9bb3))
* **api:** api update ([0fb8c9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0fb8c9a70ff063f55d878a510e2d3ea4abc392f4))
* **api:** api update ([9d84121](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d84121c7c40024585bc4acde0511cf12010c2b5))
* **api:** api update ([da35236](https://github.com/cloudflare/terraform-provider-cloudflare/commit/da35236e82f9fbe44952f4735485a17347790a69))
* **api:** api update ([d3e4688](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3e468892e9dc37d033570706430ca55115e9331))
* **api:** api update ([5b74331](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b74331e805905a8e2a19222d5dbdde5496b4e88))
* **api:** api update ([18b20e4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/18b20e4787f4dd11b0bef2c833d71099850c084b))
* **api:** api update ([960a31b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/960a31bfb323c6e0ac8731e62fe4c43c0e46fc55))
* **api:** api update ([112c784](https://github.com/cloudflare/terraform-provider-cloudflare/commit/112c7844cbc75d04c1d7a96f400baa8e0e84ed46))
* **api:** api update ([633a9db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/633a9db7fe1ec814edbad09431922c2d654ad01d))
* **api:** api update ([b1bf91c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b1bf91c80d82333abf7e7c6a1233ee6d88c4fbde))
* **api:** api update ([7b69641](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7b6964136368d63f4e69e3c04543c3e66dbefe33))
* **api:** api update ([3caf54c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3caf54c12f0cefd2ce907de355680a6b8c260eb5))
* **api:** api update ([0aa0b76](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0aa0b76e9ed2664e399154443eacab429776a705))
* **api:** api update ([6a32178](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6a32178a80c4d730cabc8d657cfaf5a48a9c2e23))
* **api:** api update ([5772a3c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5772a3c3b1cd781380a60d81d012a0f203459efe))
* **api:** api update ([63e9cff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/63e9cff04170883930f56957bcd15ab74e429d1a))
* **api:** api update ([7520bd1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7520bd1c258a1c95f1870cde6bf61530059847f1))
* **api:** api update ([81d0f38](https://github.com/cloudflare/terraform-provider-cloudflare/commit/81d0f38552a0e8aa6bac025fb4832cff3a0a04e1))
* **api:** api update ([5b962ab](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b962abded01157f27e1f710c760ec5d8868a040))
* **api:** api update ([f1cafb3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f1cafb35fef71c6bb536ea9e37b189e7218965be))
* **api:** api update ([e3e392c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e3e392cf735d2341bd7c95e3de3ad432d0350d7a))
* **api:** api update ([2877191](https://github.com/cloudflare/terraform-provider-cloudflare/commit/28771918cf3b62d298db57119d94dcf5bd5a43b9))
* **api:** api update ([e5c19c8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e5c19c84b2f33acaa44cf5b06e21db5a024eefbd))
* **api:** api update ([3050d3e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3050d3e9b026f5fe90a64e24edcb304bc5036103))
* **api:** api update ([517f362](https://github.com/cloudflare/terraform-provider-cloudflare/commit/517f362cb01671565099a77902b1a6f635d7f013))
* **api:** api update ([e457711](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e45771188d62b6735adef1cbb319399800461ed0))
* **api:** api update ([de19700](https://github.com/cloudflare/terraform-provider-cloudflare/commit/de19700106df7677724234e1d151d9d7f9891f94))
* **api:** api update ([87a5428](https://github.com/cloudflare/terraform-provider-cloudflare/commit/87a54287e73494f2e100e8f62391d2fdb74b6e9e))
* **api:** api update ([5f0b1e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f0b1e123b4f750147a4d986417a184069128861))
* **api:** api update ([9e22553](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e2255360b5de1213b9107b8e1156d90821336a2))
* **api:** api update ([5396fca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5396fca01c13d667846970a71eef37b43c5bdd2b))
* **api:** api update ([ee72dee](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee72dee70b84e8db5478228f5d7d5dc4ff92d51d))
* **api:** api update ([2900abb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2900abb340f48e32f13f25437be82ebf8ee847ba))
* **api:** api update ([f320344](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f3203448803bc34586a11f21fb996c43f5774605))
* **api:** api update ([d0e20c0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d0e20c056502bc96c09195e2edda884344f9a76e))
* **api:** api update ([7c8b299](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7c8b2994ba4b0a3edb84f58dfd5291e2851b1178))
* **api:** api update ([d2c5ef0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d2c5ef036b93cba0044d810b44de15192e99e1b3))
* **api:** api update ([0066259](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0066259736320397ecce63d3b029ce2740587a4e))
* **api:** api update ([1e07973](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e07973b68c78e0eebce1fe2102d3609f7109b67))
* **api:** api update ([e5bd521](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e5bd521c3078247d739fbd1b1216fa0e5807251b))
* **api:** api update ([fc9a3bb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc9a3bbaeb8abffe5980cbadb4238f330d9725b5))
* **api:** api update ([a265879](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a265879f1b7591a38d82f013b995fc56c9aa32ca))
* **api:** api update ([a799f10](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a799f10aa98d619f41f43b69b685957faa8e194b))
* **api:** api update ([358f4cc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/358f4cc792121cedf6ce6569c0a46207141faa1b))
* **api:** api update ([1074306](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1074306f0bb55297e15e3865d15b903cda7ccbdd))
* **api:** api update ([37d82a4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37d82a4a5d45005b736e6bf9ecc674d063b478d2))
* **api:** api update ([5f17801](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f1780158b6c3f34b8e3f286f2c6050d87132a11))
* **api:** api update ([043b87b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/043b87b855f713558b9d7b6f3f408e21e4adc95c))
* **api:** api update ([37ee222](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37ee222045a3c9aba5cf5907411d10cde130dd2e))
* **api:** api update ([3660d3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3660d3fd03a95924e11a7bb6c77eac9ebb12f0bd))
* **api:** api update ([d40f20c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d40f20cbe6d5f2f7ee08df9a5306d73462b785de))
* **api:** api update ([752c380](https://github.com/cloudflare/terraform-provider-cloudflare/commit/752c380b926f991f6a046011eb41e602e0cd5d5f))
* **api:** api update ([54c2bff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/54c2bff14c335aff8663cb810755c14645024210))
* **api:** api update ([063aabe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/063aabe3057673cd62a59aa8873194e8ad8cff74))
* **api:** api update ([24fb419](https://github.com/cloudflare/terraform-provider-cloudflare/commit/24fb4197b43704b3456ece4f27c52108a7d94f47))
* **api:** api update ([88c3bca](https://github.com/cloudflare/terraform-provider-cloudflare/commit/88c3bcafafdc7fa568f23ae5283d68413933265a))
* **api:** api update ([6dc8742](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6dc8742bef2bd414d53378ac68c589fde86a1ff5))
* **api:** api update ([a589930](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a5899308e20dbfd079053672c94a9799a539d52f))
* **api:** api update ([da51741](https://github.com/cloudflare/terraform-provider-cloudflare/commit/da517413242449cb96d49a20101fbc425daa9714))
* **api:** api update ([efaab58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/efaab58187c070a84d9343b39f44cee629019ad4))
* **api:** api update ([edb91c5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/edb91c55ae4f8585e0657b966b652fea1617da6c))
* **api:** api update ([d31d2a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d31d2a8b7dc5500bd38f4dfde80aded3322d6de9))
* **api:** api update ([e89650c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e89650c7e6d80d48a8bb321346bbea8679f9f5a4))
* **api:** api update ([e052d0d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e052d0d9d2ba400b8687a056d94554353a60ac4a))
* **api:** api update ([48d6593](https://github.com/cloudflare/terraform-provider-cloudflare/commit/48d6593016e8311c0cbceed38e5adc0b626590c6))
* **api:** api update ([70ac208](https://github.com/cloudflare/terraform-provider-cloudflare/commit/70ac208018e3e591231fd926d294c2283de32ffe))
* **api:** api update ([1dadd53](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1dadd534f56299f3f3b26f8c203e355c10adf89d))
* **api:** api update ([9b8ee5a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9b8ee5aa9d9f60f91df45d1828abcbebd726cfbd))
* **api:** api update ([8260b16](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8260b1683e4148bc93dc186eca295f9990c4d052))
* **api:** api update ([aec76bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aec76bcd6eaf78648c200e5b67f883038f9a757e))
* **api:** api update ([1edded7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1edded70c13f27d45e92b748257467c87fe166ea))
* **api:** api update ([c1dcf0e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c1dcf0e466a1881299e4e73c7e2aaa329bbf3eee))
* **api:** api update ([805bc1e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/805bc1e4fba8d49d1e34a2703aa8a8aaf8f5fe3b))
* **api:** api update ([a4ba986](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a4ba986de50551c8de0f956a9a7fce638dfb18c4))
* **api:** api update ([2ba103b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2ba103bb7473f665ae4c1cc2bba435d45f1f6861))
* **api:** api update ([67bbaf5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/67bbaf51d6afc9261af82ab2259e57b9650e032c))
* **api:** api update ([8683a24](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8683a249219c29f89351fcf878caea08cd285251))
* **api:** api update ([4d333dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4d333dcd56df895a6d4564ae66c6cc8418d705d2))
* **api:** api update ([33b7228](https://github.com/cloudflare/terraform-provider-cloudflare/commit/33b7228e9916939bf7485a8a380925902df1f614))
* **api:** api update ([a82390e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a82390ee951b700955756e5347c0f1e9e88677d9))
* **api:** api update ([05d8a1f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/05d8a1ff062007d5d70d0674c0c70a86631e60e1))
* **api:** api update ([235fd9a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/235fd9a1a51666cf81d241ac11a88fcd6d4dfe65))
* **api:** api update ([da0d778](https://github.com/cloudflare/terraform-provider-cloudflare/commit/da0d778371532c4c13b5dbbae5f05daf5360cb91))
* **api:** api update ([ea7bb97](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ea7bb975a4c9ad2165346c5bfa8da0dc9c0e4507))
* **api:** api update ([9f82224](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9f8222457fe95d58aab4b1fd35afabd0d3203b8f))
* **api:** api update ([41b6e48](https://github.com/cloudflare/terraform-provider-cloudflare/commit/41b6e483a22e8b145b94bcdb827f8affa6315846))
* **api:** api update ([9fac5b2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9fac5b2c6a96c79403e9941fb7cafe4a4d547d5f))
* **api:** api update ([0a46f20](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0a46f2073f04fc771155b600f3b3bfa15930a215))
* **api:** api update ([9e31062](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e3106225b0ff0d9599c2485a9148406e4708da7))
* **api:** api update ([1dc0992](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1dc09921eba44ac36c1e544cf6463ad2cff1eee5))
* **api:** api update ([076db07](https://github.com/cloudflare/terraform-provider-cloudflare/commit/076db0769a59e535f5765ed694ab0198faa8742d))
* **api:** api update ([6bfc29c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6bfc29cf0033f0701c233b3643247d41c120c99e))
* **api:** api update ([a35ded3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a35ded32a26c284fb5795f9de484710e62601e7b))
* **api:** api update ([069feef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/069feef9d27732897e74d2ae1de320b976da14ec))
* **api:** api update ([f4abe16](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f4abe161843d2a7ed5efcb9ba589619f47b45990))
* **api:** api update ([a1da7c3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a1da7c37a390ce1bc0d8344068e9067eaf623470))
* **api:** api update ([4b3c5c7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4b3c5c73bef7598c7ab6820a908855c7c4b78cd7))
* **api:** api update ([dcdc16f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dcdc16f3ec93070c24e8fa898b64570f9e7bb855))
* **api:** api update ([a3573be](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a3573be1d50994d0829d72d332257cd30995c19c))
* **api:** api update ([7bcc1e8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7bcc1e865565a04275285f5c81967739adc1e18e))
* **api:** api update ([2aba4fc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2aba4fcd9833a077f09eaeac55ab0ce0f7efd51d))
* **api:** api update ([cdc3e81](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cdc3e81d90b566fbfaac8835ab4310de9e056ce0))
* **api:** api update ([de36903](https://github.com/cloudflare/terraform-provider-cloudflare/commit/de369030aab77da125e3172a216a113169436420))
* **api:** api update ([72a58ea](https://github.com/cloudflare/terraform-provider-cloudflare/commit/72a58ea8e9425a773bcc13aba79588ddcac5a83c))
* **api:** api update ([72c3076](https://github.com/cloudflare/terraform-provider-cloudflare/commit/72c3076027e6b792050fe8e53dd8c65b26e7fbff))
* **api:** api update ([90c4745](https://github.com/cloudflare/terraform-provider-cloudflare/commit/90c47455e66f7d10b43fe89bad7aa71fe3c11d97))
* **api:** api update ([b26ac57](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b26ac572ac8305c883030474fc4bd52f4c2a660f))
* **api:** api update ([6530a95](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6530a95c7be5fcb9f042dbdb2fa5e30719b30ccb))
* **api:** api update ([5aa7243](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5aa7243d6d4c625aae77685cb2750b1924f55ec9))
* **api:** api update ([ca44eeb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ca44eeba5777ec8720fceacbee304a164fe59f73))
* **api:** api update ([f1a9c7e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f1a9c7e07daf6eb4909a92028847e6f1020e8ded))
* **api:** api update ([e38c3f2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e38c3f29b4a132c5c1c976490850cda67ac46d17))
* **api:** api update ([9e52ca7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9e52ca70533c5c13b36d10cad021955224c15cd1))
* **api:** api update ([3459942](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3459942ae9d2ff9dbacd896fc8ed4f8d4d76021e))
* **api:** api update ([e7491b8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e7491b83b811765b928c7fc7fc8f6fe0db5ffea3))
* **api:** api update ([8b14770](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8b14770d345ac05fff3585cf3d0c6a773978dd9f))
* **api:** api update ([975f941](https://github.com/cloudflare/terraform-provider-cloudflare/commit/975f941bb7daae86b8f37265a064cbeb57d7af7b))
* **api:** api update ([0970ac7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0970ac7e9c0d2cea70c9bedc7f00da6879a1d03b))
* **api:** api update ([b1be9d8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b1be9d87b96f65c8231a2df93066b652131aea09))
* **api:** api update ([4629e77](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4629e7764b0d310b018dffd13773818fd5dc7317))
* **api:** api update ([db21276](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db21276623047f72b63af96552f544c45ab4595b))
* **api:** api update ([b482742](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b48274217f3d5abffdc6bf7130f83e959a73176c))
* **api:** api update ([b5d537d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5d537da3615bc557cb02d37ed27811f247d9901))
* **api:** api update ([6efbff3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6efbff3b9a1f39117aa6ba07ed6c6f1eb1aa5ccc))
* **api:** api update ([f0a0485](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f0a048515f5540904fceb04df745dd178959b9ff))
* **api:** api update ([2e9b86a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e9b86af32338d903ba36dacdd113ea0c3d0c52f))
* **api:** api update ([c892255](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c8922550ceb57fdf0abd1817781dac1bd5f1986b))
* **api:** api update ([ead90dc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ead90dcc8571dc650002cfb573a64100d1eecf0f))
* **api:** api update ([fd96c7d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fd96c7d3ca662a26555e6fd13574e5fc1d3c4daa))
* **api:** api update ([57e4d5b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/57e4d5bec9bf24949b201443a7a1ea4ecca00b52))
* **api:** api update ([60f6d29](https://github.com/cloudflare/terraform-provider-cloudflare/commit/60f6d29a28362a9764b2f4656d24525f41e64a18))
* **api:** api update ([bf60837](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bf60837d1599266a9b4d7a6a3ab79504c6cc592e))
* **api:** api update ([4e2985e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4e2985ed786e8149911969abb85904081e3822f6))
* **api:** api update ([4755978](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4755978dc9027e1c34090c9569d55eaca4674a0a))
* **api:** api update ([7ee28b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7ee28b773c7611a43677b3292ed85369b8d2467d))
* **api:** api update ([343dbb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/343dbb474e9523076d318d38f18fb9e38072e46d))
* **api:** api update ([27d347d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/27d347def1fef0b0c5b27a3f57cbb9967fac1a85))
* **api:** api update ([f2b88bd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f2b88bdc05cda2e934917be9a30bcfa99fa0ef73))
* **api:** api update ([2b7e6d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b7e6d6482e41c68d3efb461720df325ec928204))
* **api:** api update ([57ae57b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/57ae57b8d9949df85f99edf4b436a96261aa7f1f))
* **api:** api update ([8c212a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8c212a99ae2b217fb318b2bdc4292203b992c4a2))
* **api:** api update ([532199e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/532199e07a1800ecbc59fab8428192b8c1234778))
* **api:** api update ([69d75d1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/69d75d133b3265da702c76b763cf4a857a57e359))
* **api:** api update ([5c41b47](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5c41b474712ab455165c208ae0f81da453b8f87e))
* **api:** api update ([eb7da9e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/eb7da9ed99299a3c9dd06f200cb26f8003e935d6))
* **api:** api update ([f86c697](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f86c6971358022c6267ca811926a63205a7331b9))
* **api:** api update ([ce6d728](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ce6d7288330f5b7e611151e025a97914b96be076))
* **api:** api update ([e292d9c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e292d9c0690ff080fef7e7e875e11184651f7359))
* **api:** api update ([9210f8f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9210f8f1bc04f2bbd3cbde43f8d1f0ffe8903b7b))
* **api:** api update ([4821d8f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4821d8f4d9b89a7cd37d377dfbdb47b794418d60))
* **api:** api update ([e187d08](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e187d08f874d1acfae038938401bf40cbbead9e0))
* **api:** api update ([8158590](https://github.com/cloudflare/terraform-provider-cloudflare/commit/815859088c3720fa9d04a2b7fd4598ad5e30a5b5))
* **api:** api update ([5298bb4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5298bb43af1110a25c8c5e9832f6ff3055b84894))
* **api:** api update ([bd3d354](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bd3d354b9a1fd21ba6f40bda36fa7dca0dffc80b))
* **api:** api update ([f7b0702](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f7b0702a7c8797cc8f6edea5712c5d385a30145d))
* **api:** bump go to 4.3.0 ([1b45f62](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1b45f629f4394aa24bdce3ce1defe9c5839cc330))
* **api:** manual updates ([c3e57cd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c3e57cdb3c7e6fc64fc65fa1914077799c106a09))
* **api:** manual updates ([b8e4dc6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b8e4dc6486c5f10756210a598ff8e30dffecca31))
* **api:** update path placeholders ([ee5d559](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ee5d5598cb872fe66339337498626ecf7adf647d))
* **threat_events:** add list support ([8cad549](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8cad549e71073f708636a50415da210fd719f8e2))
* **workers_for_platforms_script_secret:** remove terraform resource ([0327716](https://github.com/cloudflare/terraform-provider-cloudflare/commit/03277163bb9d3f62230779d2924c1e9df45067ac))


### Bug Fixes

* **api:** Fix workers_route identifier attributes ([cf994ae](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cf994aefc300990293cd0278ded380c3f099cdf4))
* **build:** do not fail if go mod tidy fails during bootstrapping ([99c36f3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/99c36f37aee512b76c09a965fc1fbbe1f170c56a))
* **cloud_connector:** dedupe nested fields ([63e0303](https://github.com/cloudflare/terraform-provider-cloudflare/commit/63e030325a0a2262b45d97aef2ba94e7b95f89b9))
* fix caching issue between Unmarshal and UnmarshalComputed ([f253a26](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f253a26fa25b774c8ffb8752010b46a8e0faa27f))
* **resources:** do not refresh resources that have unsupported response types ([febe04e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/febe04ec7f9341f665280be89bcb5ffea2ee4886))
* **schema:** fix configurability calculation for nested properties without computed values ([428f76b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/428f76bd7a3ac19521b65a8859ff8d46a6eba383))
* **workers_kv:** fix `cloudflare_workers_kv` state refresh ([2e61b3e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e61b3ec6c64993910f638a8606cdf9be53d5c41))
* **workers:** rectify bad merge ([6937388](https://github.com/cloudflare/terraform-provider-cloudflare/commit/69373887686a9be31129b2d80f2def06e34619cc))


### Chores

* **ci:** only use depot for staging repos ([ce5dd65](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ce5dd654a9eee33e5711f80996d5b315deb8fe67))
* **ci:** run on more branches ([b90e5eb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b90e5eb1951a46b2296da87805945b6aa5c5751f))
* **ci:** run on more branches and use depot runners ([0a38714](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0a3871498cac255e4620e381bb39c576ccd77320))
* **workers_kv:** test resource properly supports drift detection ([d3777ed](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3777edff39f4009579c2b011673552bebd8f47e))
* **workers_kv:** update acceptance test to use v2 cloudflare-go client ([d717117](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7171178e4da5d9ea8ef67663abdcb698360b71d))

## 5.3.0 (2025-04-09)

Full Changelog: [v5.2.0...v5.3.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.2.0...v5.3.0)

### Features

* **access_policy:** remove invalid defaults ([#5417](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5417)) ([5f6bec0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5f6bec0604d092f338a0e57024a8107ab912a05a))
* add script to build without optimisations ([#5436](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5436)) ([d1a4f49](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d1a4f4924377f614e8a52c81aa6754e92ad45822))
* **api:** Add workers telemetry routes ([828b20e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/828b20ee9ae399b47b3233ee22dfb6443130fc87))
* **api:** api update ([13ab544](https://github.com/cloudflare/terraform-provider-cloudflare/commit/13ab544fb12f5d4bcad144037160038df43e18b3))
* **api:** api update ([38b779c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/38b779c847a98d76c1fbe41c7120f08e1169c6ae))
* **api:** api update ([35e5ee6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/35e5ee687cfb1504421c4853a107d9a18490bb86))
* **api:** api update ([f8f6637](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f8f6637f12a964c2cc0f03640ffb29d7dea881d3))
* **api:** api update ([e00cc2f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e00cc2fa98890089bf3685ab2f09782222d1b054))
* **api:** api update ([c077795](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c0777953a60d316449956ac40ae2eb91972355b2))
* **api:** api update ([3b8d719](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3b8d71957ca2e5cd89d3d3a4b556f79e6ac7821c))
* **api:** api update ([0497567](https://github.com/cloudflare/terraform-provider-cloudflare/commit/049756736f23cc65090532d17460db5879b64c74))
* **api:** api update ([a5b6d28](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a5b6d283315f789fd3336b1e50a8168b6e442ac6))
* **api:** api update ([85ca5b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/85ca5b656bb5525ab5e1d2e680c5d106bbd0a82e))
* **api:** api update ([119ea92](https://github.com/cloudflare/terraform-provider-cloudflare/commit/119ea928006e646c712dd764473c81a9dad71ec0))
* **api:** api update ([b5b0a8c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5b0a8cf514adbef120d2ac895f1a15868c8f498))
* **api:** api update ([1d07b70](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1d07b70c1c312693ea08cf37290d9a24e60cb704))
* **api:** api update ([72a9f83](https://github.com/cloudflare/terraform-provider-cloudflare/commit/72a9f83b8c206d77eea18413c6efe7a9bdc2d883))
* **api:** api update ([c4180b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c4180b7e7fba960af9b45ebefd9c64bb5527058a))
* **api:** api update ([cff25d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cff25d66216405d353ad817fdb22cd6ddef6f5fd))
* **api:** api update ([72d3824](https://github.com/cloudflare/terraform-provider-cloudflare/commit/72d3824b2e229f454fac17d95efeaf2ae80deff6))
* **api:** api update ([664cb3c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/664cb3cc5aa838e4bc78b398ad845f0965365e50))
* **api:** api update ([44a8c6b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/44a8c6be67f0b0d01e77f033a68c58e99f097fcc))
* **api:** api update ([92919cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/92919cb9ca0cf2ca0e5762a4d6da588524460b7d))
* **api:** api update ([ac6ab4b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ac6ab4bf7d29dcc85dc91fd462cced31964d3cdb))
* **api:** api update ([759abb8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/759abb8584ea6d59c534c0342ad0938c2601c765))
* **api:** api update ([32c8ef6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/32c8ef68a5c8b2289976a6f7d3f50a0c79e998e4))
* **api:** api update ([dabc087](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dabc087a50bc609716ad070b818ec08c4e457d6a))
* **api:** api update ([8611b80](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8611b80185a73a6936e19e7e5f874703b9e23153))
* **api:** api update ([7053012](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7053012fbcdbb7157e4a9ed92def2fcfd06c9004))
* **api:** api update ([45b4ae6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45b4ae604d62fd7f0d7d32bef705b371df79ea6f))
* **api:** api update ([e959fe3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e959fe30e189e64211978f20beaefec8a7ec00bb))
* **api:** api update ([be9ae70](https://github.com/cloudflare/terraform-provider-cloudflare/commit/be9ae70db4cbec6b9251d0d1d76a84ba9904b075))
* **api:** api update ([37e7245](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37e724594c2d7c884dc6059d3385a8aa077f093b))
* **api:** api update ([b3f3126](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b3f312627a23a3030674f4a80d6a49c61191f2fd))
* **api:** api update ([80ac5b7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/80ac5b714c857c9a64f53009c57f7f8b9d68002a))
* **api:** api update ([f281c3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f281c3fc68adba00b116551c123e592dd3e36f0d))
* **api:** api update ([761a96f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/761a96f59d29d571b324d7c0654f6d1563b80317))
* **api:** api update ([2bb8703](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2bb87034fd1cbe46f705d417c89d2b5710ba21cc))
* **api:** api update ([69ead96](https://github.com/cloudflare/terraform-provider-cloudflare/commit/69ead9675f73c116c0a14a12cab65411ff779e11))
* **api:** api update ([e3871a2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e3871a293182a69492adf821e3c86a1d16acb984))
* **api:** api update ([09e6236](https://github.com/cloudflare/terraform-provider-cloudflare/commit/09e62364d60a454beb303244b31e8d5ff13c3bb0))
* **api:** api update ([0be6b20](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0be6b20f470c9e3e2bfe2573a154ae3453cb4aff))
* **api:** api update ([2e1e307](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e1e3078b7f65315bccfd6ddc3226630d6ef9ab3))
* **api:** api update ([78f2553](https://github.com/cloudflare/terraform-provider-cloudflare/commit/78f25532c5233ffc71eb8b8c7128bfecb2737912))
* **api:** api update ([67496e3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/67496e3f5829bda41490209d7a97de663c477c1b))
* **api:** api update ([#5372](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5372)) ([b940618](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b940618db1d989b0cc93c9adc631fb772b9b28e1))
* **api:** api update ([#5375](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5375)) ([a8ea03c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a8ea03ce65382a0c40ca2a7a3880dc4f32d7b5f3))
* **api:** api update ([#5384](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5384)) ([d24be4d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d24be4d0b5fea8d5cfde1eade41e8972ff19a8d0))
* **api:** api update ([#5396](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5396)) ([92c6897](https://github.com/cloudflare/terraform-provider-cloudflare/commit/92c68977133c01bf0dc7f46076a32a2aba6a71b4))
* **api:** api update ([#5408](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5408)) ([e6b94c8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e6b94c8c40af64f564df122c7c4083c74b682850))
* **api:** api update ([#5426](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5426)) ([bf1eec3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bf1eec3e78d4f108c0712a0f8899375725e20231))
* **api:** api update ([#5430](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5430)) ([990d99f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/990d99fbf70fb4b216c265d4db39fb22e3f44d7d))
* **api:** api update ([#5438](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5438)) ([205c1a1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/205c1a1ebf6abea8f05a6edc90f33b1efe62ba19))
* **api:** api update ([#5442](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5442)) ([bcb6b61](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bcb6b615957ee938131840304dcaaa5ed135346d))
* **api:** api update ([#5444](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5444)) ([0c62078](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0c62078656b6d6deba250cbb5f69e41408de274f))
* **api:** api update ([#5447](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5447)) ([3518fb9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3518fb9558d503622e64c97cb5f3fb8e4aca8ed4))
* **api:** api update ([#5449](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5449)) ([a018ca5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a018ca5f012008bfdbc28ef252052154ac6b1c11))
* **api:** api update ([#5453](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5453)) ([e914374](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e91437447fb39867e241cbe4798cde2656a5437b))
* **api:** api update ([#5457](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5457)) ([369b4be](https://github.com/cloudflare/terraform-provider-cloudflare/commit/369b4be85e2135f45cf186557eeed5d115cdcb3a))
* **api:** api update ([#5464](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5464)) ([10dae6c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/10dae6c523958a41af4786a72fe57e11c31f0326))
* **api:** api update ([#5465](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5465)) ([a432bdb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a432bdb8b87822c3e493ba82344b5ba25f5a2ab8))
* **api:** api update ([#5467](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5467)) ([0aee40f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0aee40f098781938906347f0787140996a6035d8))
* **api:** api update ([#5468](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5468)) ([95d34e2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/95d34e2b765410c5a3e1c994b054d7b3de2a40b6))
* **api:** api update ([#5469](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5469)) ([d96661d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d96661df7d86a83ff4850f812d7b384577c3320d))
* **api:** api update ([#5470](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5470)) ([96ea789](https://github.com/cloudflare/terraform-provider-cloudflare/commit/96ea7898e70bcb7f94a8f28d7a3db64e1828df51))
* **api:** api update ([#5473](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5473)) ([f21aa4e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f21aa4ea1cb17cd8dba0edd5bf902f6d86aa9f01))
* **api:** api update ([#5477](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5477)) ([7f1a73f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7f1a73f15d29a30481d326c25da8a745c03d1486))
* **api:** api update ([#5478](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5478)) ([e2e5502](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e2e550282edb3cfc17017cd81b30bce0821119de))
* **api:** api update ([#5479](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5479)) ([1daafb7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1daafb753364f87f32b7c5a8624bd131a62d3b9a))
* **api:** display deprecation messages on resources and attributes ([#5425](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5425)) ([4a0554e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4a0554e6e2f40ea65b6bf145b905e7674f2e0d87))
* **api:** manual updates ([02da0fd](https://github.com/cloudflare/terraform-provider-cloudflare/commit/02da0fde7cc4cccd1de6610de23b254e1c8ddd6a))
* **api:** manual updates ([dc58a0f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc58a0f35c0c66fc630bb3449097d252694bdcc7))
* **api:** manual updates ([abd45b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/abd45b1bfe1efa17094fadcb698a9cfef194e67f))
* **api:** manual updates ([b18b48b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b18b48b270752c2024bb8bb2c325c3b7f19a2471))
* **api:** manual updates ([27e8929](https://github.com/cloudflare/terraform-provider-cloudflare/commit/27e8929392495303be7114cb4f51f8bdf1ccc5cd))
* **api:** manual updates ([b92b2d7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b92b2d7986f6d73b73e9c1e7c60e471f210b73db))
* **api:** manual updates ([2f2436d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2f2436d4e7542b77efac18da6e21316729663290))
* **api:** Update workers telemetry route methods ([2b4db81](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b4db81a6444e13c2f21bd77014bd30430f58a62))
* **devices:** add registrations support ([05bc7a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/05bc7a971707f9adc6557b06dbc24ab3c87f0da5))
* **docs:** add secrets store to navigation ([b5be043](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5be043d7a75121bfb147b54c4b9be49d5e5d0ec))
* **pipelines:** add support ([5068d66](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5068d6699fe455de3cacdbc71fb78048defe036b))
* **secret_store:** add support ([aa76038](https://github.com/cloudflare/terraform-provider-cloudflare/commit/aa7603828e48777a65b997184ee3ec2868286e8a))
* **secrets_store:** remove incorrect bulk edit endpoint ([a8e9411](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a8e9411e2cd87f167472038166bc7433e9d7e5dc))


### Bug Fixes

* **access_policy:** remove invalid defaults ([#5416](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5416)) ([daaceff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/daaceff2c0f22f21275b3277b2103e64d238a936))
* **account_token:** handle unordered `policies` ([#5433](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5433)) ([3f36851](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3f36851222781c0c5b0dc96f77fb5c0e7ba69b25))
* add deserialization annotations to synthetic ID properties ([#5376](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5376)) ([754df3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/754df3fd266522489aa360756cdfc15241bf1c10))
* **api_token:** handle value across updates ([#5414](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5414)) ([1db1294](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1db1294b447e97ce7351c9ffe83bcb2ae3b5e5a6))
* **api:** better support for environment variables as provider properties ([#5377](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5377)) ([a6e7785](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a6e7785d9c72e3cbcf8f6b541cbd1963599f8ffa))
* **api:** skip generation of update endpoint when only a create endpoint exists ([#5397](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5397)) ([9b588cb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9b588cb0db19b6c7d7f71001bd3ca34c9226fa53))
* **dns_record:** don't include defaults for `data` ([#5461](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5461)) ([8a5e8ce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8a5e8ceaa66313e2c200146250b5c754d9802da4))
* **hostnames:** define correct path parameter for updates ([#5428](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5428)) ([c54a6a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c54a6a04bbf4187c554a9df8e09f01c969a11715))
* **logpush:** remove empty default value ([#5401](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5401)) ([0659e1f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0659e1f23852d7d618175666f92e09a2a081bf71))
* **origin_ca_certificate:** persist values across reads ([#5415](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5415)) ([1e6d072](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e6d072f0c1f8313c679c81d5ccade0a983a7990))
* **r2_custom_domain:** add jurisdiction into the schema ([#5404](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5404)) ([7d2b266](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7d2b2663d2f847ea25d580422f7d65530d124779))
* **ruleset:** handle stricter marshaling ([#5391](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5391)) ([b9f0292](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b9f029266ec814c5bb53794a8c53c3c92929a98f))
* **workers_cron_trigger:** remove duplicate struct ([#5435](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5435)) ([7a54b42](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a54b42dab13cb4cfca6dd68b278e76208ac950b))
* **workers_for_platforms_script_secret:** correctly update resource name ([#5459](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5459)) ([e52e4db](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e52e4db7102fc269e226b10db64cc1d78b80a639))
* **workers_script:** cleanup placement duplication ([#5437](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5437)) ([3ac96b6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3ac96b6ba25d4c0a9c1f33be9007dede56ef2604))
* **zero_trust_tunnel_cloudflared_virtual_networ:** persist `is_default` write only value ([#5434](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5434)) ([368ab5d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/368ab5d80f2dcb06355cf2e6c04ec517bb16c66c))
* **zero_trust_tunnel_cloudflared:** persist write only values ([#5432](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5432)) ([7d62813](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7d6281361484121b3a2294801e4b12a3cf0a83d6))


### Chores

* **build:** scripts/format should not fail if generate-docs fails ([#5466](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5466)) ([56b6feb](https://github.com/cloudflare/terraform-provider-cloudflare/commit/56b6feb6319862bf5f0ddd0d7a664a90be3aaeb2))
* **deps:** bump terraform-plugin-docs ([#5399](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5399)) ([ed4e092](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ed4e09202ad3cdfedbb60a071af5235836d9209f))
* **deps:** fix indirect updates ([#5402](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5402)) ([12d9433](https://github.com/cloudflare/terraform-provider-cloudflare/commit/12d943380c251ca03d46520f19b11d0728e3f564))
* **internal:** codegen related update ([0d15cfe](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0d15cfe01b3920cbe13d63778a1d58a2f3c6c142))
* **internal:** codegen related update ([#5441](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5441)) ([ec3412b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ec3412ba07b6d01e87f460aee30365dedcb23c45))
* **internal:** codegen related update ([#5443](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5443)) ([a82d949](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a82d9494ce181036784dd9cbf7208014cb5c23e7))
* **internal:** codegen related update ([#5463](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5463)) ([d3d4be5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d3d4be5044f1b20abc47ea900b14339c055c1403))
* remove unnecessary `toListParams` methods in singular data sources ([#5371](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5371)) ([495ae4c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/495ae4cbae690b700e5c257880162781e1c66cec))
* **tests:** improve enum examples ([#5476](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5476)) ([225ac31](https://github.com/cloudflare/terraform-provider-cloudflare/commit/225ac31022c973b393703f788dc7ca621eef73a0))
* update dependencies ([#5387](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5387)) ([b3bff1d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b3bff1d6831a07e5592fab4ee671ce05cd79aa11))


### Documentation

* generate ([#5400](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5400)) ([ecf3561](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ecf3561e757ca5a42a49e29c17e682bbf923b201))
* generate ([#5429](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5429)) ([8428f6a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8428f6aab6b56973a2b7988ae236cafdca7d063c))
* generate ([#5462](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5462)) ([80b2588](https://github.com/cloudflare/terraform-provider-cloudflare/commit/80b2588ae9eb9010408deb0124cfed9f4dce2c38))
* generate ([#5480](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5480)) ([d7e3b84](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d7e3b84c1e9fcfaea3022ccfad5da983a3855a2f))

## 5.2.0 (2025-03-20)

Full Changelog: [v5.1.0...v5.2.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.1.0...v5.2.0)

### Features

* add docs generation to format script ([#5294](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5294)) ([a199683](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a199683abcd5fcbefc88ad09a88287faf4cb2a66))
* add SKIP_BREW env var to ./scripts/bootstrap ([#5274](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5274)) ([45090a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45090a94f1b2fd65a8c5c204d8abc834e42e35b2))
* **api:** api update ([#5243](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5243)) ([7d287a7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7d287a725ed750935ddb7837fca6af08b8dac94f))
* **api:** api update ([#5249](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5249)) ([9f253d5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9f253d5648823900dd0d883c40c8a10d80a89809))
* **api:** api update ([#5257](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5257)) ([220adc9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/220adc96184f1e8c00710a344cb2c3c8e73ab2ef))
* **api:** api update ([#5265](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5265)) ([fc3045a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fc3045aef7a13601bc9c7a71f32376378c80daa9))
* **api:** api update ([#5267](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5267)) ([c7198d8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c7198d89cedfc2d48c4d1daa08647d7c9a8541e3))
* **api:** api update ([#5269](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5269)) ([3f44f89](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3f44f894e107282b3dae7408b80f669bf8ee47b4))
* **api:** api update ([#5270](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5270)) ([56c1da3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/56c1da3ce85fe18a9e7bf28b38d2f611a2ecd736))
* **api:** api update ([#5271](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5271)) ([b6c093a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b6c093ab6af31273b34c01c246af0b38cfba2de1))
* **api:** api update ([#5293](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5293)) ([941a30a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/941a30afc5025fb28cd2da827b7be00e75c63cb4))
* **api:** api update ([#5295](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5295)) ([86e4e4e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/86e4e4e9ab3644c7189007f42b122691265e76c3))
* **api:** api update ([#5299](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5299)) ([fe8c29d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fe8c29d143b77c1e50bc25f0b59da4abd38d6322))
* **api:** api update ([#5300](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5300)) ([0abdfcf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0abdfcf8129a6519c07ae0a7f29dec7915ba6014))
* **api:** api update ([#5302](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5302)) ([063348c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/063348cb6073bc48b00fb7195b7236b0b9ee937a))
* **api:** api update ([#5309](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5309)) ([b8674a5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b8674a563cc8a83f48a4734888212ef842015da4))
* **api:** api update ([#5325](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5325)) ([9a65852](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9a6585275ca77c3259319fd3d9ab157035501a4a))
* **api:** api update ([#5326](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5326)) ([5cc7f58](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5cc7f589b9a4eb80aeadd2c11075ab81704dcd0b))
* **api:** api update ([#5332](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5332)) ([f16b95e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f16b95e4e95464e74b00b00358f3726fe89b3c5c))
* **api:** api update ([#5338](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5338)) ([6ae5427](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6ae5427bb4a87c07acc3ebeb78009acd529799e7))
* **api:** api update ([#5354](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5354)) ([98b1ec9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/98b1ec9b256c744255956bc2b6a50820acda4437))
* **api:** api update ([#5355](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5355)) ([0fb620e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0fb620eb56a3a7dbb60d7f52c991ebfd05d9a13e))
* **api:** api update ([#5356](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5356)) ([9ca6737](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9ca67378c321bcb99598bfce616f661d8e5d901d))
* **api:** api update ([#5357](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5357)) ([2324e79](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2324e79dfbc54d011b015385d6229a8ff782308b))
* **api:** api update ([#5359](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5359)) ([5b1c190](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5b1c1905a39c6cf515eac121df14b58059b92a25))
* **api:** manual updates ([#5314](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5314)) ([2660117](https://github.com/cloudflare/terraform-provider-cloudflare/commit/26601178a7c3f03eda3d4e66e81271846429884f))
* **custom_pages:** add resource ([#5343](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5343)) ([57d76e2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/57d76e23593011eb96be1570b05e0a3a8a221ffe))
* **custom_pages:** mark `identifier` as `id` ([#5344](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5344)) ([9705e1b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9705e1b119047d466f77155b6268f245d41082d8))
* **custom_pages:** mark `identifier` as `id` ([#5345](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5345)) ([5d1afaa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5d1afaa346df1660aec5dcf7d9da53f3664ac366))
* **custom_pages:** reintroduce endpoints ([#5312](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5312)) ([4653c96](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4653c961cb2a9231deb80da0fd769cbadf6e3421))
* **dns_settings:** fix hierarchy ([#5291](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5291)) ([cb5bee3](https://github.com/cloudflare/terraform-provider-cloudflare/commit/cb5bee340be778e6c0ea07a48e400cd72fb2ed03))
* **dns:** split account and zone DNS settings ([#5283](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5283)) ([3c9e05e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3c9e05ee132a52f049a67c514294c2c486e00711))
* **dns:** split account and zone DNS settings ([#5285](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5285)) ([d669e8f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d669e8f2735356b1c56a0206f7381423797b8f77))
* **internal:** add HA and IO to initialisms ([#5276](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5276)) ([ead063a](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ead063a7742ee151ea9dda46eec10488a5e9d458))
* **internal:** bump cloudflare-go to 4.2.0 ([#5341](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5341)) ([559850d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/559850d44763fef8cdb613f4d96ce4da3f33e1e7))
* **internal:** revert HA and IO to initialisms ([#5279](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5279)) ([8cce7e4](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8cce7e420551cec2953509ce3d32dab4f9ea627d))
* **waiting_rooms:** add account level list API ([#5310](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5310)) ([915f6f7](https://github.com/cloudflare/terraform-provider-cloudflare/commit/915f6f7035e7dd7e35f7ffd234dd0bd65a4905aa))
* **workers:** add in secrets endpoints ([#5329](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5329)) ([0d8f363](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0d8f363ce542866bea12a9214853a29f72aa1652))
* **zero_trust_device_*_profile:** mark include and exclude as mutually exclusive ([2db548c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2db548c60629c165e0f661212b77e978c5fa87e5))


### Bug Fixes

* **accoun_token:** mark `meta` as read only ([84e8c23](https://github.com/cloudflare/terraform-provider-cloudflare/commit/84e8c23b53e790cdb9a87f88121f3a936334e76c))
* **account_token:** fix missing model change ([fff0f2c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fff0f2c123122821d06b799d953cb48e91b91309))
* **account_token:** handle `value` write only field ([4cbb4b5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4cbb4b5e4762f977930285f56ef82df0987c998e))
* **account:** remove recreation on tenant unit ([76fbb98](https://github.com/cloudflare/terraform-provider-cloudflare/commit/76fbb980d75773099dd6c79f0633311b41e5fc2e))
* **api:** avoid spurious replacement plans for computed ID properties ([#5244](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5244)) ([37baea6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/37baea6253e3fbb47ef7ea4450f35f89b5bcd20a))
* **api:** remove min and max validations in mismatched union variants ([#5263](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5263)) ([b5f51a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b5f51a09d84433fe01dcd87b1abab4f1c0171448))
* **authenticated_origin_pulls_certificate:** handle `private_key` write only field ([78b9ff1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/78b9ff10b903843ed2afa364e463c55b62631409))
* **authenticated_origin_pulls_certificate:** populate `certificate_id` from the `id` ([2b53245](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2b53245b4f0c92cf017962ea67270bd10bd4e6cf))
* **dns_record:** relax constraint for overlapping unions ([ac79ff8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ac79ff82a87fc87b9d33b58a438ec9eae1ac8a48))
* find-by style data source attributes should share models with plural data sources ([#5251](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5251)) ([d488159](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d488159a2c80abfde39b5ccc2fe571a56e96905b))
* **r2_custom_domain:** remove duplicated domain value ([e062813](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e0628130a1c6231b2afab0dd8b5bfea4a5aa737a))
* **r2_custom_domain:** update path placeholders to de-duplicate internal values ([#5281](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5281)) ([5ef949d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5ef949d9c0c3f4de1b3d304f24e9406b3a92ce3c))
* **rulesets:** remove unused fields ([dcab45f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dcab45f52e21d6b96668c2c37718aa9b5e429d24))
* **waiting_rooms:** comment out broken struct for now ([3f89aef](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3f89aef24be42d53a31661e80184b8f6a586ec3a))
* **workers_script:** re-resolve the correct schemas ([05b25ba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/05b25bac38ea0f005ea9eb08592c23160f8248e3))


### Chores

* **internal:** codegen related update ([#5286](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5286)) ([1e603a0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1e603a075cb7a0307257dad2139a02ebb2b034ff))


### Documentation

* generate ([e13aae0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e13aae021f469c929bb5c0c59db4589cbc67ffe4))
* generate ([9ce87fa](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9ce87fa6d213b34e175f3f5f4db5f371f00b878c))
* generate ([b33529e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b33529e7ac8a3c392358b1163be36eaaf560f85b))
* generate ([771eaed](https://github.com/cloudflare/terraform-provider-cloudflare/commit/771eaedab9d754ea67ea507cacf62f4ae72f0b73))

## 5.1.0 (2025-02-13)

Full Changelog: [v5.0.0...v5.1.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.0.0...v5.1.0)

### Features

* **api:** disable zero_trust_tunnel_cloudflared_token ([#5128](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5128)) ([df2c3bc](https://github.com/cloudflare/terraform-provider-cloudflare/commit/df2c3bc059f35eddb7b91fef866df7c32165cf05))
* **api:** enable zero_trust_tunnel_cloudflared_token ([#5127](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5127)) ([1bd200e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1bd200e88c78565c0ed879e64f8d91473f30e365))
* **grit:** add more support for dns_record attributes ([3dbe899](https://github.com/cloudflare/terraform-provider-cloudflare/commit/3dbe899a8084a14bb586da181318abe17c7f04ef))
* various codegen changes ([d91aee1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/d91aee1ef6b403177eb2bdfb5c73a8a3d36b79c3))


### Bug Fixes

* **grit:** handle inner objects within the object for records ([e7b7bb1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/e7b7bb1e4d4868b2c94304f9a14069eb5e2822b0))
* **grit:** handle inner objects within the object for records ([c9a5257](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c9a5257d4abbf7f0920415b3ad087789dbfd61ab))
* **grit:** handle inner objects within the object for records ([ae22af5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ae22af5228a5b73e3193862c0eccc1325df94d79))
* **grit:** make pattern names consistent ([0b2ba12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/0b2ba12bb261244c6e9bd24e4b6fa783f26bc0e7))
* update migration guide to use source, not stdlib ([9d208d6](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9d208d6577f60940ec2bc1d6dcb181fcefbfdfac))
* use correct name for Grit patterns ([2f8d522](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2f8d5220bc1c76003035c5af885883367bf06867))


### Documentation

* clean out previously set schema_versions ([fba939d](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fba939df831ff60e451c2f50310eb626687bcdb8))
* handle cloudflare_record data migration ([9eb450b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/9eb450bb6f146f49a3f6c9a8c73369d747e164f3))
* regenerate ([bbf53bf](https://github.com/cloudflare/terraform-provider-cloudflare/commit/bbf53bf80a550844e27097487bcc262fbd4d6516))
* update deprecation dates ([7a8b7d2](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7a8b7d2424fe20a0c83e7a76b069d9b7c042f751))
* update page_rules migration guidance ([45e30b1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/45e30b1ed3de447d54223d5479ca52fd8609db42))
