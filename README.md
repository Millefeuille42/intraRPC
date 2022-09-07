# intraRPC

intraRPC is a rich presence client for 42.
## Acquire the binary

### From the prebuilt binary
Go to the [**releases**](https://github.com/Millefeuille42/intraRPC/releases) on this repository
and download the latest release, all binaries are built for the 42 computers (Linux x64).

### Build from source

Download the sources:

```bash
git clone git@github.com:Millefeuille42/intraRPC.git
```
And build the binary for your system: (note: this requires to have `Go >= 1.13` installed)
```bash
go build .
```

## Installation
Move the now acquired binary to your `$PATH` environment variable.
## Usage

You need first to create a `.intraRpc.conf` at your `$HOME` written in YAML, you can find a template
in the root of this repository.

You also need to have an App registered at the 42 API.

Start the program:
```bash
intraRPC [--print]
```

The `--print` flag outputs the RPC data as a string
```
<login>@<location> - Level: <level> - Blackhole: <bh> days <emoji>
```

As an example, here is my output
```
mlabouri@e3r2p22 - Level: 6.37 - Blackhole: 44 days 🤔
```

Please consider that, the RPC protocol using a WebSocket, it is a blocking application, you may run it in the
background with the `&` flag. You can properly interrupt it with a `CTRL-C` or a kill command.

This program is not protected against simultaneous runs, be careful. However, this feature might be implemented
int future releases.

With the print flag, the program shut-downs after printing its string.

## Current Objectives
- Simultaneous runs protection
- Command line override of the selected cursus
- If this program acquire enough notoriety, becoming an official 42 app, this will permit to register the app
on my 42 account only, and it will not be necessary to have your own API app registered