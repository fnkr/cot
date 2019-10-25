# List of environment variables recognized by cot

| Variable                 | Default | Description |
| ------------------------ | ------- | --- |
| `COT_ARG_*`              | `[]`    | Additional Podman/Docker arguments, one argument per variable |
| `COT_ARGS`               | `""`    | Additional Podman/Docker arguments, separated by whitespace (` `) |
| `COT_CAP_ADD`            | `[]`    | [`--cap-add`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities), separated by comma (`,`) |
| `COT_CAP_DROP`           | `[]`    | [`--cap-drop`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities), separated by comma (`,`) |
| `COT_CPUS`               |  80% of the available cores on Linux, no default value otherwise | [`--cpus`](https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources) |
| `COT_DEBUG`              | `false` | Send debug output to `/dev/stderr` |
| `COT_DRY_RUN`            | `false` | Do not actually execute Docker/Podman command |
| `COT_ENV_*`              | `[]`    | Additional environment variables to set (`COT_ENV_foo=bar` will become `--env=foo=bar`) |
| `COT_IMAGE`              | `fnkr/cot` | Docker/Podman image to use |
| `COT_INTERACTIVE`        | `true`  | `--interactive` |
| `COT_LIMIT`              | `[]`    | Directories in which cot is allowed to run, separated by colon (`:`) |
| `COT_MEMORY`             | -       | [`--memory`](https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources) |
| `COT_MEMORY_RESERVATION` | -       | [`--memory-reservation`](https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources) |
| `COT_NET`                | `"slirp4netns"` (Podman), `"bridge"` (Docker) | [`--net`](https://docs.docker.com/engine/reference/run/#network-settings) |
| `COT_READ_ONLY_ROOT`     | `false` if Podman is used, `true` if Docker is used | Mount root directory (`/`) read-only |
| `COT_SHELL`              | `"/bin/sh"` | Default shell for user in container in `/etc/passwd` |
| `COT_TMP`                | `fmt.Sprintf("/tmp/%s-%s-%s", BinName(), ToolName(), UID())` | Path to temporary directory, used for `/etc/{passwd,group}`, `/tmp` and `/home/$USER` mounts |
| `COT_TOOL`               | `"podman"` if found in `$PATH`, `"docker"` if found in `$PATH` and `$USER` is in `docker` group, otherwise `"sudo docker"` | `podman`, `docker`, or path to Podman or Docker (which must end with `/podman` or `/docker`) |
| `COT_TTY`                | `true`  | `--tty` |
| `COT_VOLUME_*`           | `[]`    | Additional volumes to mount (`COT_VOLUME_foo=/mnt/cot:/mnt:ro,z` will become `--volume=/mnt/cot:/mnt:ro,z`) |
| `EDITOR`                 | -       | Variable will be forwarded to container as-is |
| `SSH_AUTH_SOCK`          | -       | Path to SSH agent socket (will be mounted in container if set) |

`*` is a placeholder, all matching variables will be used. 
