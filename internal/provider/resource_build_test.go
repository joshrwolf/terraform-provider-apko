package provider

import (
	"fmt"
	"regexp"
	"testing"

	ocitesting "github.com/chainguard-dev/terraform-provider-oci/testing"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceApkoBuild(t *testing.T) {
	repo, cleanup := ocitesting.SetupRepository(t, "test")
	defer cleanup()

	repostr := repo.String()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "apko_config" "foo" {
  config_contents = <<EOF
contents:
  repositories:
    - https://packages.wolfi.dev/os
  keyring:
    - https://packages.wolfi.dev/os/wolfi-signing.rsa.pub
  packages:
    - wolfi-baselayout
    - ca-certificates-bundle
    - tzdata

accounts:
  groups:
    - groupname: nonroot
      gid: 65532
  users:
    - username: nonroot
      uid: 65532
      gid: 65532
  run-as: 65532

archs:
  - x86_64
  - aarch64
EOF
}

resource "apko_build" "foo" {
  repo   = %q
  config = data.apko_config.foo.config
}
`, repostr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"apko_build.foo", "repo", regexp.MustCompile("^"+repostr)),
					resource.TestMatchResourceAttr(
						"apko_build.foo", "image_ref", regexp.MustCompile("^"+repostr+"@sha256:")),
					resource.TestCheckResourceAttr("apko_build.foo", "sboms.%", "3"),
				),
			},
			// Update the config and make sure the image gets rebuilt.
			{
				Config: fmt.Sprintf(`
data "apko_config" "foo" {
  config_contents = <<EOF
contents:
  repositories:
    - https://packages.wolfi.dev/os
  keyring:
    - https://packages.wolfi.dev/os/wolfi-signing.rsa.pub
  packages:
    - wolfi-baselayout
    - ca-certificates-bundle
    - tzdata
    - git # <-- add git

accounts:
  groups:
    - groupname: nonroot
      gid: 65532
  users:
    - username: nonroot
      uid: 65532
      gid: 65532
  run-as: 65532

archs:
  - x86_64
  - aarch64
EOF
}

resource "apko_build" "foo" {
	repo   = %q
	config = data.apko_config.foo.config
}`, repostr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"apko_build.foo", "repo", regexp.MustCompile("^"+repostr)),
					resource.TestMatchResourceAttr(
						"apko_build.foo", "image_ref", regexp.MustCompile("^"+repostr+"@sha256:")),
				),
			},
		},
	})
}

func TestAccResourceApkoBuild_ProviderOpts(t *testing.T) {
	repo, cleanup := ocitesting.SetupRepository(t, "test")
	defer cleanup()

	repostr := repo.String()

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"apko": providerserver.NewProtocol6WithError(&Provider{
				repositories: []string{"https://packages.wolfi.dev/os"},
				keyring:      []string{"https://packages.wolfi.dev/os/wolfi-signing.rsa.pub"},
				archs:        []string{"x86_64", "aarch64"},
				packages:     []string{"wolfi-baselayout"},
			}),
		}, Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "apko_config" "foo" {
  config_contents = <<EOF
contents:
  packages:
    - ca-certificates-bundle
    - tzdata
EOF
}

resource "apko_build" "foo" {
	repo   = %q
	config = data.apko_config.foo.config
}
`, repostr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"apko_build.foo", "repo", regexp.MustCompile("^"+repostr)),
					resource.TestMatchResourceAttr(
						"apko_build.foo", "image_ref", regexp.MustCompile("^"+repostr+"@sha256:")),
				),
			},
			// Update the config and make sure the image gets rebuilt.
			{
				Config: fmt.Sprintf(`
data "apko_config" "foo" {
  config_contents = <<EOF
contents:
  packages:
    - ca-certificates-bundle
    - tzdata
    - busybox # <-- add busybox
EOF
}

resource "apko_build" "foo" {
	repo   = %q
	config = data.apko_config.foo.config
}
`, repostr),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"apko_build.foo", "repo", regexp.MustCompile("^"+repostr)),
					resource.TestMatchResourceAttr(
						"apko_build.foo", "image_ref", regexp.MustCompile("^"+repostr+"@sha256:")),
				),
			},
		},
	})
}
