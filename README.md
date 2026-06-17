# baoctx

[![goreleaser](https://github.com/adfinis/baoctx/actions/workflows/release.yml/badge.svg)](https://github.com/adfinis/baoctx/actions/workflows/release.yml)
[![lint](https://github.com/adfinis/baoctx/actions/workflows/lint.yml/badge.svg)](https://github.com/adfinis/baoctx/actions/workflows/lint.yml)

A CLI tool to manage context profiles for [OpenBao](https://openbao.org). This allows you to save connection and configuration details, which would otherwise be set using environment variables — into named context profiles, and easily switch between them.


### Installation

1. Download the latest release from the [Releases](https://github.com/adfinis/baoctx/releases) page.
2. Install the package for your distribution:
   - **Debian/Ubuntu**: `sudo apt install baoctx-<version>.deb`
   - **Fedora/RHEL**: `sudo dnf install baoctx-<version>.rpm`
   - **Arch Linux (AUR)**: `yay -S baoctx-bin`.


### Example use case

There are two OpenBao clusters, one for Dev (<http://dev-openbao:8200>) and one for Prod (<https://prod-openbao:8200>).

Running `bao` CLI commands locally will by default attempt to connect to <https://localhost:8200>. To connect to another cluster, you need to set the appropriate `BAO_*` environment variables. `baoctx` lets you save multiple sets of connection details into context profiles and switch between them easily.

### What Is a Context Profile?

A context profile is a named set of configuration parameters for an OpenBao instance. For example, a `prod` context profile might have an `endpoint` of `https://prod-openbao:8200`, a `namespace` of `admin/prod`, and a `token`. Selecting that profile renders the corresponding `export BAO_ADDR=...; export BAO_NAMESPACE=...; export BAO_TOKEN=...` commands, which can be applied to the current shell with `eval`.

### Example usage

```shell
eval $(baoctx select prod)
```

### Supported Tools

- OpenBao

### Configuring baoctx For Your Shell

`baoctx` can set default context profiles that are automatically loaded into new shell sessions via environment variables. To enable this, configure your shell's startup script with:

```shell
baoctx config --path "~/.zshrc"
```

#### Using fish instead of bash / zsh

```shell
baoctx config --path "~/.config/fish/config.fish" --shell fish
```

This appends a small helper script that sources all defaults when a new shell session starts.

### OpenBao

#### Create Example

```shell
baoctx create staging \
  --endpoint "https://staging-openbao.example.com:8200" \
  --cacert "/path/to/ca.pem" \
  --cert "/path/to/client.pem" \
  --key "/path/to/client-key.pem" \
  --skip-verify true \
  --cli-no-colour true \
  --client-timeout "60s" \
  --format "json"
```

#### Create Example with OIDC
```shell
baoctx create testing \
        --endpoint https://testing-openbao.example.com:8200 \
        --auth-method oidc \
        --oidc-callback-host localhost \
        --oidc-listen-addr 127.0.0.1 \
        --oidc-role device \
        --oidc-callback-mode device
```

#### Update Example

```shell
baoctx update staging \
  --endpoint "https://staging-openbao.example.com:8200" \
  --cacert "/path/to/new-ca.pem" \
  --skip-verify true \
  --format "json"
```

#### Delete Example

```shell
baoctx delete staging
```

#### List Example

```shell
baoctx list
```

### Setting Default Context Profiles

Set a default context profile with the `set-default` sub command:

```shell
baoctx set-default staging
```

Once a default has been set, new shell sessions will spawn with those environment variables already exported.

### Switching Context Profiles

Switch context using the `select` sub command:

```shell
baoctx select dev
```

This prints all `export BAO_*` commands for the selected context profile. To apply them in the current shell:

```shell
eval $(baoctx select dev)
```
