# Registry documentation

This provider automatically generates the documentation at
https://registry.terraform.io/providers/cloudflare/cloudflare using
[terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs).

We use the data-source or resource schema definitions alongside example Terraform
definitions to generate consistent and self-updating documentation.

You can generate the documentation locally using `make docs`.

## Structure

```
docs/
├── data-sources
│   ├── <individual data sources>
├── guides
│   ├── version-2-upgrade.md
│   └── version-3-upgrade.md
├── index.md
└── resources
    └── <individual data sources>

examples/
├── provider
│   └── provider.tf
└── resources
    ├── ...
    │   ├── import.sh
    │   └── resource.tf
    └── <resource_name>
        ├── import.sh
        └── resource.tf

templates/
├── data-sources
│   ├── ...
│   └── <resource_name>.md
├── guide
│   ├── version-2-upgrade.md
│   └── version-3-upgrade.md
├── index.md.tmpl
├── resources
│   ├── ...
│   └── <resource_name>.md
└── resources.md.tmpl
```

### `docs/`

The `docs` repository is the generated output from the documentation build
process. **This directory SHOULD NOT be manually changed**. Changes are either
made via the resource/data-source schema or the `templates` directory otherwise
they will be lost next time documentation is generated.

### `examples/`

`examples` is a set of nested directories that matches the following conventions:

- Inside of `data-sources` or `resources`, there is a directory per data-source
  or resource. I.e. `cloudflare_ruleset`.
- In each entity directory there are two files, `resource.tf` and `import.sh`.
  `resource.tf` is an example of how to use the entity and `import.sh` is the
  `terraform import` syntax and command (if applicable).

Full directory layout.

```
$ tree examples/resources/cloudflare_access_application/
examples/resources/cloudflare_access_application/
├── import.sh
└── resource.tf
```

### `templates/`

`templates` contains either:

- The hardcoded documentation that is copied verbatim into the `docs` directory
  during generation (files must end in `.md`); or
- A template that overrides the default template used to generate the
  documentation (files must end in `.md.tmpl`). This option is used if an
  additional banner or layout configuration is needed. Avoid duplicating if your
  data-source or resource doesn't need any modification.

## Migration

Where possible, you should migrate the entity you modify to use the generation
process and remove the manual steps of keeping these in sync. The steps to take
are:

- Add a `Description` to the data-source or resource definition.
- Update all schema references to include a `Description` property.
- Determine if you need a custom template (most use cases won't need this).
  - If you need a custom template, add it to the corresponding `templates`
    directory.
- Delete the old documentation at
  `templates/{data-source,resource}/<resource_name>.md`.
- Run `make docs`.
- Commit the changes.

## Previewing registry documentation changes

For individual documentation pages, you can use
https://registry.terraform.io/tools/doc-preview by copying and pasting in the
updated markdown.
