# baoctx

[![goreleaser](https://github.com/adfinis/baoctx/actions/workflows/release.yml/badge.svg)](https://github.com/adfinis/baoctx/actions/workflows/release.yml)
[![lint](https://github.com/adfinis/baoctx/actions/workflows/lint.yml/badge.svg)](https://github.com/adfinis/baoctx/actions/workflows/lint.yml)

A CLI tool to manage context profiles for [OpenBao](https://openbao.org). This allows you to save connection and configuration details, which would otherwise be set using environment variables — into named context profiles, and easily switch between them.


### Installation

**Releases**

Binaries can be downloaded from the releases page.

[https://github.com/adfinis/baoctx/releases](https://github.com/adfinis/baoctx/releases)

### Example use case

There are two OpenBao clusters, one for Dev (<http://dev-openbao:8200>) and one for Prod (<https://prod-openbao:8200>).

Running `bao` CLI commands locally will by default attempt to connect to <https://localhost:8200>. To connect to another cluster, you need to set the appropriate `BAO_*` environment variables. `baoctx` lets you save multiple sets of connection details into context profiles and switch between them easily.

### What Is a Context Profile?

A context profile is a named set of configuration parameters for an OpenBao instance. For example, a `prod` context profile might have an `endpoint` of `https://prod-openbao:8200`, a `namespace` of `admin/prod`, and a `token`. Selecting that profile renders the corresponding `export BAO_ADDR=...; export BAO_NAMESPACE=...; export BAO_TOKEN=...` commands, which can be applied to the current shell with `eval`.

### Example usage

```shell
eval $(baoctx openbao select prod)
```

### Supported Tools

- OpenBao

### Configuring baoctx For Your Shell

`baoctx` can set default context profiles that are automatically loaded into new shell sessions via environment variables. To enable this, configure your shell's startup script with:

```shell
baoctx config --path "~/.zshrc"
```

This appends a small helper script that sources all defaults when a new shell session starts.

### OpenBao

#### Create Example

```shell
baoctx openbao create staging \
  --endpoint "https://staging-openbao.example.com:8200" \
  --cacert "/path/to/ca.pem" \
  --cert "/path/to/client.pem" \
  --key "/path/to/client-key.pem" \
  --skip-verify true \
  --cli-no-colour true \
  --client-timeout "60s" \
  --format "json"
```

#### Update Example

```shell
baoctx openbao update staging \
  --endpoint "https://staging-openbao.example.com:8200" \
  --cacert "/path/to/new-ca.pem" \
  --skip-verify true \
  --format "json"
```

#### Delete Example

```shell
baoctx openbao delete staging
```

#### List Example

```shell
baoctx openbao list
```

### Setting Default Context Profiles

Set a default context profile with the `set-default` sub command:

```shell
baoctx openbao set-default staging
```

Once a default has been set, new shell sessions will spawn with those environment variables already exported.

### Switching Context Profiles

Switch context using the `select` sub command:

```shell
baoctx openbao select dev
```

This prints all `export BAO_*` commands for the selected context profile. To apply them in the current shell:

```shell
eval $(baoctx openbao select dev)
```
