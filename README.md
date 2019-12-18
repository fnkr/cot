# [cot](https://github.com/fnkr/cot)

cot is a convenient way to containerize command line applications with Podman or Docker.
The basic concept is to have a wrapper for `docker run` that creates a virtual environment
that grossly matches the host environment and gives the containerized process access to
the current directory only.  

By default, cot will

- run within an ephemeral Podman or Docker container
- run with the same UID/GID as the current user
- have all capabilities dropped
- have access to the current directory
- have access to `$SSH_AUTH_SOCK`
- have read-only access the `/etc/hosts` file
- `/tmp` and `$HOME` will be persisted in the `/tmp` directory of the host

## Build dependencies

### Fedora

```sh
sudo dnf install golang libselinux-devel
```

## Build

Only standard library and golang.org/x is used.

```sh
go get github.com/fnkr/cot/cmd/cot
```

## Install

You can copy cot to some directory that is in your `$PATH` if you want. 

```sh
(eval "$(go env)" && sudo cp "${GOBIN:-$GOPATH/bin}/cot" /usr/local/bin/)
```

## Usage

This example executes `npm install` within a container but you can use it
with any tool that can run within a Podman/Docker container. 

```sh
# You propably want to add this to your ~/.bashrc or ~/.zshrc too.
# cot will refuse to run if called from outside of ~/test or ~/example.
export COT_LIMIT=~/test:~/example

# You can call npm with "cot npm" or link npm to cot and call it just "npm".
sudo ln -s /usr/local/bin/cot /usr/local/bin/npm

# Done! This will run npm in a container.
npm install
```

A full list of configuration options can be found in [`ENVIRONMENT.md`](ENVIRONMENT.md).

## Custom images

By default the `fnkr/cot` image is used.
You can use any image you like, I'd recommend to create your own.
You can use a custom image by setting the `COT_IMAGE` environment variable.

```sh
COT_IMAGE=ubuntu cot uname -a
```

## macOS

Although cot works on macOS, Docker for Mac does not yet support sharing unix sockets
[[1](https://github.com/docker/for-mac/issues/483), [2](https://github.com/docker/for-mac/issues/410)].
To fix SSH agent forwarding on macOS you can install this workaround I made for my employer:
[punktDe/docker-mac-ssh-auth-sock](https://github.com/punktDe/docker-mac-ssh-auth-sock).
