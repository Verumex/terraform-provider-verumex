# Verumex Core Provider

_This repository was generated as a template from the
[provider scaffold](https://github.com/hashicorp/terraform-provider-scaffolding).
See notes there about the Terraform Plugin SDK and other helpful "getting
started" information._

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 

```sh
$ go install
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources! Act accordingly.

```sh
$ make testacc
```

## Deploying a Release

See [here](https://developer.hashicorp.com/terraform/tutorials/providers/provider-release-publish#create-a-provider-release)
for details about how to deploy a new version of the provider. Basically
you just add a tag with `git tag v...` and push it, the rest of the
configuration is already in place. But check the GitHub Action (`release`)
to ensure it worked as you expected.
