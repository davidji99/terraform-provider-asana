---
layout: "asana"
page_title: "Provider: Asana"
sidebar_current: "docs-asana-index"
description: |-
  Use the Asana provider to interact with the resources in Asana.
---

# Asana Provider

The Asana provider interacts with [Asana APIs](https://developers.asana.com/docs) to create various resources
in your account.

## Contributing

Development happens in the [GitHub repo](https://github.com/davidji99/terraform-provider-asana):

* [Releases](https://github.com/davidji99/terraform-provider-asana/releases)
* [Issues](https://github.com/davidji99/terraform-provider-asana/issues)

## Example Usage

```hcl
# Configure the Asana provider
provider "asana" {
  # ...
}

# Create a new Project
resource "asana_project" "foobar" {
  # ...
}
```

## Authentication

The Asana provider offers a flexible means of providing credentials for authentication.
The following methods are supported, listed in order of precedence, and explained below:

- Static credentials
- Environment variables

### Static credentials

Credentials can be provided statically by adding an `access_token` arguments to the Asana provider block:

```hcl
provider "asana" {
  access_token = "SOME_ACCESS_TOKEN"
}
```

### Environment variables

When the Asana provider block does not contain an `access_token` argument, the missing credential will be sourced
from the environment via the `ASANA_ACCESS_TOKEN` environment variables respectively:

```hcl
provider "asana" {}
```

```shell
$ export ASANA_ACCESS_TOKEN="SOME_KEY"
$ terraform plan
Refreshing Terraform state in-memory prior to plan...
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Required) Asana personal access token. It must be provided, but it can also
  be sourced from [other locations](#Authentication).

* `url` - (Optional) Custom Base Asana API endpoint. Defaults to `https://app.asana.com/api/1.0`.

* `headers` - (Optional) Additional API headers.