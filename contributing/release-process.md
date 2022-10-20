## Release Process

We aim to release on a fortnightly cadence. This is to accommodate downstream tools 
(such as [cf-terraforming](https://github.com/cloudflare/cf-terraforming)) and 
systems that may need to integrate the changes made in the Terraform provider. 
To determine when the next release is due, you can either:

- Review the latest [releases](https://github.com/cloudflare/terraform-provider-cloudflare/releases); or
- Review the [current milestones](https://github.com/cloudflare/terraform-provider-cloudflare/milestones).

If a hotfix is needed, the same process outlined below is used however only the 
semantic versioning patch version is bumped.

- Ensure CI is passing for [`master` branch](https://github.com/cloudflare/terraform-provider-cloudflare/actions?query=branch%3Amaster).
- Remove "(Unreleased)" portion from the header for the version you are intending 
  to release (here, 2.27.0). Create a new H2 above for the next unreleased 
  version (here 2.28.0). Example diff:
  ```diff
  + ## 2.28.0 (Unreleased)

  + ## 2.27.0
  - ## 2.27.0 (Unreleased)

  NOTES:

  * provider: Update to cloudflare-go v0.22.0 ([#1184](https://github.com/cloudflare/terraform-provider-cloudflare/issues/1184))
  ```
  Bumping the minor version is usually fine here unless you are intending on 
  releasing a major version bump. 
- Create a new GitHub release with the release title exactly matching the tag 
  (e.g. `v2.27.0`) and copy the entries from the CHANGELOG to the release notes.
- A GitHub Action will now build the binaries, documentation and distribute them 
  to the Terraform registry for the [Cloudflare provider](https://registry.terraform.io/providers/cloudflare/cloudflare/latest).
- Once this is completed, close off the milestone for the current release and 
  open the next that matches the CHANGELOG additions from earlier. Example: close 
  v2.27.0 but open a v2.28.0.
