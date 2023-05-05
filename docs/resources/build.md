---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "apko_build Resource - terraform-provider-apko"
subcategory: ""
description: |-
  This performs an apko build from the provided config file
---

# apko_build (Resource)

This performs an apko build from the provided config file

## Example Usage

```terraform
resource "apko_build" "example" {
  # Where to publish the resulting image, e.g. docker.io/user/repo
  repo = "..."

  # Pass in the apko configuration here.  If you'd like to define this in a file
  # so it can be used with apko as well, you can make this something like this
  # instead:  config = file("${path.module}/apko.yaml")
  config = jsonencode({
    contents = {
      repositories = ["https://packages.wolfi.dev/os"]
      keyring      = ["https://packages.wolfi.dev/os/wolfi-signing.rsa.pub"]
      packages = [
        "wolfi-baselayout",
        "ca-certificates-bundle",
        "tzdata"
      ]
    },
    accounts = {
      groups = [{
        groupname = "nonroot",
        gid       = 65532
      }],
      users = [{
        username = "nonroot",
        uid      = 65532,
        gid      = 65532
      }],
      run-as = 65532
    },
    archs = [
      "x86_64",
      "aarch64"
    ]
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (String) The apko configuration file.
- `repo` (String) The name of the container repository to which we should publish the image.

### Read-Only

- `id` (String) The resulting fully-qualified digest (e.g. {repo}@sha256:deadbeef).
- `image_ref` (String) The resulting fully-qualified digest (e.g. {repo}@sha256:deadbeef).
- `sboms` (Map of Object) A map from the APK architecture to the digest for that architecture and its SBOM. (see [below for nested schema](#nestedatt--sboms))

<a id="nestedatt--sboms"></a>
### Nested Schema for `sboms`

Read-Only:

- `digest` (String)
- `predicate` (String)
- `predicate_type` (String)

