<p align="center"><img width="200px" src="/_docs/img/logo.png" alt="s2top"/></p>

#

![release][release] ![homebrew][homebrew]

Top-like interface for container metrics

`s2top` provides a concise and condensed overview of real-time metrics for multiple containers:
<p align="center"><img src="_docs/img/grid.gif" alt="s2top"/></p>

as well as an [single container view][single_view] for inspecting a specific container.

`s2top` comes with built-in support for Docker and runC; connectors for other container and cluster systems are planned for future releases.

## Install

Fetch the [latest release](https://github.com/catataw/s2top/releases) for your platform:

#### Linux

```bash
sudo wget https://github.com/catataw/s2top/releases/download/v0.7.3/s2top-0.7.3-linux-amd64 -O /usr/local/bin/s2top
sudo chmod +x /usr/local/bin/s2top
```

#### OS X

```bash
brew install s2top
```
or
```bash
sudo curl -Lo /usr/local/bin/s2top https://github.com/catataw/s2top/releases/download/v0.7.3/s2top-0.7.3-darwin-amd64
sudo chmod +x /usr/local/bin/s2top
```

#### Docker

```bash
docker run --rm -ti \
  --name=s2top \
  --volume /var/run/docker.sock:/var/run/docker.sock:ro \
  quay.io/vektorlab/s2top:latest
```

`s2top` is also available for Arch in the [AUR](https://aur.archlinux.org/packages/s2top-bin/)

## Building

Build steps can be found [here][build].

## Usage

`s2top` requires no arguments and uses Docker host variables by default. See [connectors][connectors] for further configuration options.

### Config file

While running, use `S` to save the current filters, sort field, and other options to a default config path. These settings will be loaded and applied the next time `s2top` is started.

### Options

Option | Description
--- | ---
`-a`	| show active containers only
`-f <string>` | set an initial filter string
`-h`	| display help dialog
`-i`  | invert default colors
`-r`	| reverse container sort order
`-s`  | select initial container sort field
`-scale-cpu`	| show cpu as % of system total
`-v`	| output version information and exit
`-shell` | specify shell (default: sh)

### Keybindings

Key | Action
--- | ---
`<enter>` | Open container menu
`a` | Toggle display of all (running and non-running) containers
`f` | Filter displayed containers (`esc` to clear when open)
`H` | Toggle s2top header
`h` | Open help dialog
`s` | Select container sort field
`r` | Reverse container sort order
`o` | Open single view
`l` | View container logs (`t` to toggle timestamp when open)
`e` | Exec Shell
`S` | Save current configuration to file
`q` | Quit s2top

[build]: _docs/build.md
[connectors]: _docs/connectors.md
[single_view]: _docs/single.md
[release]: https://img.shields.io/github/release/bcicen/s2top.svg "s2top"
[homebrew]: https://img.shields.io/homebrew/v/s2top.svg "s2top"
