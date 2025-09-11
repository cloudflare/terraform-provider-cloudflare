# GitHub Actions Runner

Standardize on a GitHub Actions runner.

```grit
language yaml

`runs-on: $runner` where {
	$runner <: or {
		r"ubuntu.+" => `nscloud-ubuntu-22.04-amd64-4x16`,
		r"macos.+" => `nscloud-macos-4x16`
	}
}
```

## Examples

### Ubuntu

Before:

```yaml
name: grit-check

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - '*'

jobs:
  run:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v4
      - name: grit-check
        uses: getgrit/github-action-check@v0
```

After:

```yaml
name: grit-check

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - '*'

jobs:
  run:
    runs-on: nscloud-ubuntu-22.04-amd64-4x16
    steps:
      - name: Check out code
        uses: actions/checkout@v4
      - name: grit-check
        uses: getgrit/github-action-check@v0
```

### macOS

Namespace Cloud also supports macOS runners.

```yaml
name: grit-check

jobs:
  run:
    runs-on: macos-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v4
      - name: grit-check
        uses: getgrit/github-action-check@v0
```

```yaml
name: grit-check

jobs:
  run:
    runs-on: nscloud-macos-4x16
    steps:
      - name: Check out code
        uses: actions/checkout@v4
      - name: grit-check
        uses: getgrit/github-action-check@v0
```
