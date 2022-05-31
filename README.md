# intraRPC

intraRPC is a rich presence client for 42
## Installation

Clone this repository and build it with

```bash
go build .
```

Then move the produced binary to your `$PATH`

## Usage

```bash
intraRPC <cursus> [--print]
```

With `cursus` being the cursus you want to display <br>
The `--print` flag outputs the RPC data as a string
```
<login>@<location> - Level: <level> - Blackhole: <bh> days <emoji>
```

As an example, here is my output
```
mlabouri@e3r2p22 - Level: 6.37 - Blackhole: 44 days 🤔
```