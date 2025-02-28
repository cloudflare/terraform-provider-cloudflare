# Changelog

## 5.2.0 (2025-02-28)

Full Changelog: [v5.1.0...v5.2.0](https://github.com/cloudflare/terraform-provider-cloudflare/compare/v5.1.0...v5.2.0)

### ⚠ BREAKING CHANGES

* **tunnels:** move all cloudflared resources into dedicated namespace ([#5143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5143))

### Features

* add doc string to specify what legal terraform values are for enums ([#5199](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5199)) ([b99d403](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b99d403aa24264e538dcb98387741b1ae70307f6))
* **api_token_permission_groups:** split off account/user into dedicated endpoints ([#5162](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5162)) ([c6d07be](https://github.com/cloudflare/terraform-provider-cloudflare/commit/c6d07be33998cc6f9f1c4064f81d4ed65fe74e89))
* **api:** api update ([#5137](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5137)) ([222d9e1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/222d9e1adcd47b8e459deb946babe42b952500ae))
* **api:** api update ([#5154](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5154)) ([013bfea](https://github.com/cloudflare/terraform-provider-cloudflare/commit/013bfead10af49cc3f244a45b46cfef6f2f198c7))
* **api:** api update ([#5160](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5160)) ([4214c9b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4214c9bb36dd50b752fca6f13d0818f904aa464c))
* **api:** api update ([#5165](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5165)) ([89bf11e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/89bf11ecb9d5cccf887cf1d78acddd66ce93ce26))
* **api:** api update ([#5172](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5172)) ([a55499b](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a55499b0e3f1819e17d66e7eb163e1face86b8aa))
* **api:** api update ([#5188](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5188)) ([dc93f59](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc93f59d743049eaa4b6b3248f26a6550d793535))
* **api:** api update ([#5193](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5193)) ([6077aff](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6077aff3e62a2095c380f0abffb91b4d8c3a073b))
* **api:** api update ([#5194](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5194)) ([8516e97](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8516e979958c19b9bb58d2b6084aaebdf7d29ff2))
* **api:** api update ([#5197](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5197)) ([4574dba](https://github.com/cloudflare/terraform-provider-cloudflare/commit/4574dbaa1ac2f32acf62721c328d6a79f210c250))
* **api:** api update ([#5209](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5209)) ([6fed8b5](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6fed8b57412586625ad5ca4a10ae22c4f75b80d9))
* **api:** api update ([#5210](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5210)) ([ae0d2ce](https://github.com/cloudflare/terraform-provider-cloudflare/commit/ae0d2ce011a0c3da5fd09e706bf3f79196e516d2))
* **api:** api update ([#5216](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5216)) ([05054a8](https://github.com/cloudflare/terraform-provider-cloudflare/commit/05054a845bed0634ecd792980077bb1c87906188))
* **api:** api update ([#5218](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5218)) ([b711a2f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b711a2fe94de774674083b456b6550f257c5bef4))
* **api:** api update ([#5219](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5219)) ([6b6ff43](https://github.com/cloudflare/terraform-provider-cloudflare/commit/6b6ff438badc02c1d11ab823bfac13211b616afb))
* **authenticated_origin_pulls_settings:** add support ([#5180](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5180)) ([2e7eb91](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2e7eb91fe6cd24fdf40f1d6a8213cd304b9913f5))
* **authenticated_origin_pulls_settings:** handle upsert deletion ([b4fcc5c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b4fcc5cabef6be1bf33a26744738b51f6c423799))
* **dns_record:** toggle stricter drift detection ([#5212](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5212)) ([b59d34c](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b59d34c9665826939167b3f487590d8dcfbff6de))
* **firewall_rules:** remove duplicated `id` query parameter ([#5155](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5155)) ([a163794](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a163794281a4517afba103e99b27d6f45cc9b1cd))
* **grit:** add `account_id` =&gt; `account.id` migration for cloudflare_zone ([7bb3f3f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7bb3f3f36c399dce62dcf204328e8ecce69ecef4))
* **grit:** remove `plan` and `jump_start` from zone state ([091523e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/091523e8d106d76e0d228787cf93751e1aea0a8a))
* **r2:** add support for jurisdictions on all resources ([51961a9](https://github.com/cloudflare/terraform-provider-cloudflare/commit/51961a9aa4982637f87375496b5e4fc7a3ac90d6))
* **ruleset:** toggle stricter drift detection ([#5213](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5213)) ([71156b0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71156b002b9336ef42d6438ed7301e81a3fb4395))
* **terraform:** mark some attributes as sensitive ([#5170](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5170)) ([f7ac77e](https://github.com/cloudflare/terraform-provider-cloudflare/commit/f7ac77e81ddf82e258ca88e544b3eaf9f69c703c))
* **tunnels:** move all cloudflared resources into dedicated namespace ([#5143](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5143)) ([010615f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/010615fa7ca4989487bba7a1194c80ee0ebfc0cb))


### Bug Fixes

* **api_token:** store `value` out of band of marshaler ([4165021](https://github.com/cloudflare/terraform-provider-cloudflare/commit/416502103b9187b9d1be9e408cae2dc982a2b351))
* **authenticated_origin_pulls_settings:** remove manual delete handling ([dc66e55](https://github.com/cloudflare/terraform-provider-cloudflare/commit/dc66e55f2cf3ffb318c6b7a28c7b9f372aa10f02))
* **client:** mark some request bodies as optional ([#5168](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5168)) ([2416e12](https://github.com/cloudflare/terraform-provider-cloudflare/commit/2416e12a118347fd06c705c7bb75bf4f85293a16))
* **datasource:** honor query params in non-list data sources ([#5151](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5151)) ([fd0cb35](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fd0cb357bbc806fe991354955ae5868afe028578))
* **datasource:** improve configurability of path parameters on data sources ([#5147](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5147)) ([7033468](https://github.com/cloudflare/terraform-provider-cloudflare/commit/7033468d093806d9376c64cb91e51e9ec282406d))
* **internal:** enforce stricter drift comparison for all resources ([#5223](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5223)) ([b773d03](https://github.com/cloudflare/terraform-provider-cloudflare/commit/b773d03d6445267ef1f7fe95b05ebaa8cd4afd77))
* restrict schema version reset to cloudflare owned resources ([5565468](https://github.com/cloudflare/terraform-provider-cloudflare/commit/5565468988289c097d283a87d40954cf95e7fb30))
* **zero_trust_tunnel_cloudflared_token:** map `token` to top level `result ([db778c1](https://github.com/cloudflare/terraform-provider-cloudflare/commit/db778c11678f5691052f3b894d83996e7e828fd4))


### Chores

* casing change in doc string ([#5200](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5200)) ([71e290f](https://github.com/cloudflare/terraform-provider-cloudflare/commit/71e290fd9e56e877731457695c1ccae3a998abad))
* **internal:** codegen related update ([#5211](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5211)) ([316b5f0](https://github.com/cloudflare/terraform-provider-cloudflare/commit/316b5f0716342f62a10f0885e8f7c041fcda3d07))
* simplify string literals ([#5226](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5226)) ([1f3da51](https://github.com/cloudflare/terraform-provider-cloudflare/commit/1f3da51f4ce9f71a744dac1ba4882a953f1a1f45))


### Documentation

* clarify list_item split ([fb275ea](https://github.com/cloudflare/terraform-provider-cloudflare/commit/fb275ea35b7d21dedbef9248790ff9441a0d279c))
* generate ([a62eeae](https://github.com/cloudflare/terraform-provider-cloudflare/commit/a62eeae98f45392260f8f5928493914fbaaa905f))
* update URLs from stainlessapi.com to stainless.com ([#5220](https://github.com/cloudflare/terraform-provider-cloudflare/issues/5220)) ([8317718](https://github.com/cloudflare/terraform-provider-cloudflare/commit/8317718df20102df8348aaf5b04e9191ca82fd07))

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
