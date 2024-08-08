[![n2x.io](https://github.com/n2x-io/assets/blob/HEAD/images/logo/n2x-logo_black_180x34.png)](https://n2x.io)

[![Discord](https://img.shields.io/discord/654291649572241408?color=%236d82cb&style=flat&logo=discord&logoColor=%23ffffff&label=Chat)](https://n2x.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/n2x-io/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/n2xHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@n2x)

Open source projects from [n2x.io](https://n2x.io).

# n2x-cli

[![Go Report Card](https://goreportcard.com/badge/n2x.dev/x-cli)](https://goreportcard.com/report/n2x.dev/x-cli)
[![Release](https://img.shields.io/github/v/release/n2x-io/c-cli?display_name=tag&style=flat)](https://github.com/n2x-io/x-cli/releases/latest)
[![GitHub](https://img.shields.io/github/license/n2x-io/x-cli?style=flat)](/LICENSE)

This repository contains `n2xctl`, a tool for managing the [n2x.io](https://n2x.io) SASE platform from the command line.

`n2xctl` is available for a variety of Linux platforms, macOS and Windows.

## Minimun Requirements

`n2xctl` has the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements) as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

## Getting Started

See [Quick Start](https://n2x.io/docs/getting-started/quickstart/) to learn how to start building your n2x cloud-agnostic architecture.

See [Installation](#installation) for more details and other platforms.

## Documentation

For the complete n2x.io platform documentation visit [n2x.io/docs](https://n2x.io/docs/).

## Installation

### Download Binaries

Linux, macOS and Windows binary downloads are available from the [Releases](https://github.com/n2x-io/x-cli/releases) page.

You can download the pre-compiled binaries and install them with the appropriate tools.

### Linux Installation

#### Linux binary installation with curl

1. Download the latest release.

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/linux/amd64/n2xctl"
    ```

2. Validate the binary (optional).

    Download the `n2xctl` checksum file:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/linux/amd64/n2xctl_checksum.sha256"
    ```

    Validate the `n2xctl` binary against the checksum file:

    ```bash
    sha256sum --check < n2xctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    n2xctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    n2xctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install `n2xctl`.

    ```shell
    sudo install -o root -g root -m 0755 n2xctl /usr/local/bin/n2xctl
    ```

    > **Note**:
    > If you do not have root access on the target system, you can still install n2xctl to the `~/.local/bin` directory:
    >
    > ```shell
    > chmod +x n2xctl
    > mkdir -p ~/.local/bin
    > mv ./n2xctl ~/.local/bin/n2xctl
    > # and then append (or prepend) ~/.local/bin to $PATH
    > ```

4. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

#### Package Repository

n2x.io provides a package repository that contains both DEB and RPM downloads.

##### **Debian/Ubuntu**

1. Run the following to setup a new APT `sources.list` entry and install `n2x-cli`:

    ```shell
    echo 'deb [trusted=yes] https://repo.n2x.io/apt/ /' | sudo tee /etc/apt/sources.list.d/n2x.list
    sudo apt update
    sudo apt install n2x-cli
    ```

2. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

##### **RHEL/CentOS** 

1. Run the following to create a `n2x.repo` file and install `n2x-cli`:

    ```shell
    cat <<EOF | sudo tee /etc/yum.repos.d/n2x.repo
    [n2x]
    name=n2x repository - stable
    baseurl=https://repo.n2x.io/yum
    enabled=1
    gpgcheck=0
    EOF
    sudo yum install n2x-cli
    ```

2. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

#### Homebrew installation on Linux

If you are on Linux and using [Homebrew](https://docs.brew.sh/Homebrew-on-Linux) package manager, you can install the n2x.io CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install n2x-io/tap/n2x-cli
    ```

2. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

### macOS Installation

#### macOS binary installation with curl

1. Download the latest release.

    **Intel**:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/darwin/amd64/n2xctl"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/darwin/arm64/n2xctl"
    ```

2. Validate the binary (optional).

    Download the `n2xctl` checksum file:

    **Intel**:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/darwin/amd64/n2xctl_checksum.sha256"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/darwin/arm64/n2xctl_checksum.sha256"
    ```

    Validate the `n2xctl` binary against the checksum file:

    ```console
    shasum --algorithm 256 --check n2xctl_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    n2xctl: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    n2xctl: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Make the `n2xctl` binary executable.

    ```shell
    chmod +x n2xctl
    ```

4. Move the `n2xctl` binary to a file location on your system `PATH`.

    ```shell
    sudo mkdir -p /usr/local/bin
    sudo mv n2xctl /usr/local/bin/n2xctl
    sudo chown root: /usr/local/bin/n2xctl
    ```

    > **Note**: Make sure `/usr/local/bin` is in your `PATH` environment variable.

5. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

#### Homebrew installation on macOS

If you are on macOS and using [Homebrew](https://brew.sh/) package manager, you can install the n2x.io CLI with Homebrew.

1. Run the installation command:

    ```shell
    brew install n2x-io/tap/n2x-cli
    ```

2. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

### Windows Installation

#### Windows binary installation with curl

1. Open the Windows PowerShell as Administrator and create the `n2x` folder.

    ```shell
    mkdir 'C:\Program Files\n2x'
    ```

2. Download the latest release into the `n2x` folder.

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/windows/amd64/n2xctl.exe"
    ```

3. Validate the binary (optional).

    Download the `n2xctl.exe` checksum file:

    ```shell
    curl -LO "https://dl.n2x.io/binaries/stable/latest/windows/amd64/n2xctl.exe_checksum.sha256"
    ```

    Validate the `n2xctl.exe` binary against the checksum file:

    - Using Command Prompt to manually compare CertUtil's output to the checksum file downloaded:

         ```shell
         CertUtil -hashfile n2xctl.exe SHA256
         type n2xctl.exe_checksum.sha256
         ```

    - Using PowerShell to automate the verification using the -eq operator to get a `True` or `False` result:

         ```powershell
         $($(CertUtil -hashfile n2xctl.exe SHA256)[1] -replace " ", "") -eq $(type n2xctl.exe_checksum.sha256).split(" ")[0]
         ```

4. Add the folder `C:\Program Files\n2x` to your `PATH` environment variable. You can either append or prepend it to the existing value.

    ```powershell
    $ENV:PATH="$ENV:PATH;C:\Program Files\n2x"
    ```

5. Verify that the installed version is the latest:

    ```shell
    n2xctl version show
    ```

## Running with Docker

You can also run `n2xctl` as a Docker container.

Registry:

- `ghcr.io/n2x-io/n2xctl`

Example usage:

```shell
docker run --rm -ti -v $HOME/.n2x:/root/.n2x:ro ghcr.io/n2x-io/n2xctl help
```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [Releases](https://github.com/n2x-io/x-cli/releases) page:

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```console
COSIGN_EXPERIMENTAL=1 cosign verify ghcr.io/n2x-io/n2xctl
```

## Usage

To view a list of available commands and options, simply run:

```shell
n2xctl help
```

For in-depth information about specific commands and their usage, refer to our dedicated [CLI Command Reference](https://n2x.io/docs/reference/cli-command-reference/) guide.

## Uninstall

### Uninstall n2x-cli in Linux

To remove `n2xctl` from the system, use the following commands:

#### Binary

```shell
sudo rm /usr/local/bin/n2xctl
sudo rm -f $HOME/.n2x
```

#### Package Repository

##### **Debian/Ubuntu**

```shell
sudo apt-get -y remove n2x-cli
sudo rm -f $HOME/.n2x
```

##### **RHEL/Centos**

```shell
sudo yum -y remove n2x-cli
sudo rm -f $HOME/.n2x
```

### Uninstall n2x-cli in macOS

To remove `n2xctl` from the system, use the following commands:

```shell
sudo rm /usr/local/bin/n2xctl
sudo rm -f $HOME/.n2x
```

### Uninstall n2x-cli in macOS

To remove `n2xctl` from the system, open the Windows PowerShell as Administrator and use the following commands:

```powershell
rm 'C:\Program Files\n2x' -r -force
```

## Community

Have questions, need support and or just want to talk?

Get in touch with the n2x community!

[![Discord](https://img.shields.io/badge/Join_us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://n2x.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/n2x-io/discussions)
[![X](https://img.shields.io/badge/Follow_on_X-000000?style=flat&logo=x&logoColor=white)](https://x.com/n2xHQ)
[![Mastodon](https://img.shields.io/badge/Follow_on_Mastodon-2f0c7a?style=flat&logo=mastodon&logoColor=white)](https://mastodon.social/@n2x)

## Code of Conduct

Participation in the n2x community is governed by the Contributor Covenant [Code of Conduct](https://github.com/n2x-io/.github/blob/HEAD/CODE_OF_CONDUCT.md). Please make sure to read and observe this document.

Please make sure to read and observe this document. By participating, you are expected to uphold this code.

## License

The n2x open source projects are licensed under the [Apache 2.0 License](/LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fn2x-io%2Fx-cli.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fn2x-io%2Fx-cli?ref=badge_large)
