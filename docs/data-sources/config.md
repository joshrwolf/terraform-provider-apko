---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "apko_config Data Source - terraform-provider-apko"
subcategory: ""
description: |-
  This reads an apko configuration file into a structured form.
---

# apko_config (Data Source)

This reads an apko configuration file into a structured form.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config_contents` (String) The raw contents of the apko configuration.

### Read-Only

- `config` (Object) The parsed structure of the apko configuration. (see [below for nested schema](#nestedatt--config))
- `id` (String) A unique identifier for this apko config.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Read-Only:

- `accounts` (Object) (see [below for nested schema](#nestedobjatt--config--accounts))
- `annotations` (Map of String)
- `archs` (List of String)
- `cmd` (String)
- `contents` (Object) (see [below for nested schema](#nestedobjatt--config--contents))
- `entrypoint` (Object) (see [below for nested schema](#nestedobjatt--config--entrypoint))
- `environment` (Map of String)
- `include` (String)
- `options` (Map of Object) (see [below for nested schema](#nestedobjatt--config--options))
- `os-release` (Object) (see [below for nested schema](#nestedobjatt--config--os-release))
- `paths` (List of Object) (see [below for nested schema](#nestedobjatt--config--paths))
- `stop-signal` (String)
- `vcs-url` (String)
- `work-dir` (String)

<a id="nestedobjatt--config--accounts"></a>
### Nested Schema for `config.accounts`

Read-Only:

- `groups` (List of Object) (see [below for nested schema](#nestedobjatt--config--accounts--groups))
- `run-as` (String)
- `users` (List of Object) (see [below for nested schema](#nestedobjatt--config--accounts--users))

<a id="nestedobjatt--config--accounts--groups"></a>
### Nested Schema for `config.accounts.groups`

Read-Only:

- `gid` (Number)
- `groupname` (String)
- `members` (List of String)


<a id="nestedobjatt--config--accounts--users"></a>
### Nested Schema for `config.accounts.users`

Read-Only:

- `gid` (Number)
- `uid` (Number)
- `username` (String)



<a id="nestedobjatt--config--contents"></a>
### Nested Schema for `config.contents`

Read-Only:

- `keyring` (List of String)
- `packages` (List of String)
- `repositories` (List of String)


<a id="nestedobjatt--config--entrypoint"></a>
### Nested Schema for `config.entrypoint`

Read-Only:

- `command` (String)
- `services` (Map of String)
- `shell-fragment` (String)
- `type` (String)


<a id="nestedobjatt--config--options"></a>
### Nested Schema for `config.options`

Read-Only:

- `accounts` (Object) (see [below for nested schema](#nestedobjatt--config--options--accounts))
- `contents` (Object) (see [below for nested schema](#nestedobjatt--config--options--contents))
- `entrypoint` (Object) (see [below for nested schema](#nestedobjatt--config--options--entrypoint))
- `environment` (Map of String)

<a id="nestedobjatt--config--options--accounts"></a>
### Nested Schema for `config.options.accounts`

Read-Only:

- `run-as` (String)


<a id="nestedobjatt--config--options--contents"></a>
### Nested Schema for `config.options.contents`

Read-Only:

- `packages` (Object) (see [below for nested schema](#nestedobjatt--config--options--contents--packages))

<a id="nestedobjatt--config--options--contents--packages"></a>
### Nested Schema for `config.options.contents.packages`

Read-Only:

- `add` (List of String)
- `remove` (List of String)



<a id="nestedobjatt--config--options--entrypoint"></a>
### Nested Schema for `config.options.entrypoint`

Read-Only:

- `command` (String)
- `services` (Map of String)
- `shell-fragment` (String)
- `type` (String)



<a id="nestedobjatt--config--os-release"></a>
### Nested Schema for `config.os-release`

Read-Only:

- `bug-report-url` (String)
- `home-url` (String)
- `id` (String)
- `name` (String)
- `pretty-name` (String)
- `version-id` (String)


<a id="nestedobjatt--config--paths"></a>
### Nested Schema for `config.paths`

Read-Only:

- `gid` (Number)
- `path` (String)
- `permissions` (Number)
- `recursive` (Boolean)
- `source` (String)
- `type` (String)
- `uid` (Number)


