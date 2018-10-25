# Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-cloudflare`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-cloudflare
```

When it comes to building you have two options:

#### `make build` and install it globally

If you don't mind installing the development version of the provider
globally, you can use `make build` in the provider directory which will
build and link the binary into your `$GOPATH/bin` directory.

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-cloudflare
$ make build
```

#### `go build` and install it local to your changes

If you would rather install the provider locally and not impact the
stable version you already have installed, you can use the
`~/.terraformrc` file to tell Terraform where your provider is. You do
this by building the provider using Go.

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-cloudflare
$ go build -o terraform-provider-cloudflare
```

And then update your `~/.terraformrc` file to point at the location
you've built it.

```
providers {
  cloudflare = "{GOPATH}/github.com/terraform-providers/terraform-provider-cloudflare/terraform-cloudflare-provider"
}
```

(Be sure to swap out `{GOPATH}` with your actual `$GOPATH`. This does
not get evaluated)

A caveat with this approach is that you will need to run `terraform
init` whenever the provider is rebuilt. You'll also need to remember to
comment it/remove it when it's not in use to avoid tripping yourself up.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

See above for which option suits your workflow for building the provider.

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Updating a vendored dependency

Terraform providers use [`govendor`][govendor] to manage the vendored
dependencies. To update a dependency, you can run `govendor fetch
<dependency_path>`. An example of updating the `cloudflare-go` library:

```
$ govendor fetch github.com/cloudflare/cloudflare-go
```

This will update the local `vendor` directory and `vendor/vendor.json`
to include the new dependencies.

[govendor]: https://github.com/kardianos/govendor
