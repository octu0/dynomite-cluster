# dynomite-cluster

dynomite-cluster is [Dynomite](https://github.com/Netflix/dynomite) cluster management tool.

features below:

- token management for multi-region deployments
- strict validation / realize safe ops
- automatically cold bootstrap

## Build

Build requires Go version 1.11+ installed.

```
$ go version
```

Run `make pkg` to Build and package for linux, darwin.

```
$ git clone https://github.com/octu0/dynomite-cluster
$ make pkg
```

## Usage

### discovery

for pickup token and find neightbors node

```
$ dynomite-cluster discovery --seed xxxx:2101 --token TTTT --replica-host yyyy --replica-port yyy
```

### validate

validation for node setup before cluster

```
$ dynomite-cluster validate --peer-host xxxx --peer-port xxx --replica-host yyyy --replica-port yyy
```

### bootstrap

cold bootstrap

```
$ dynomite-cluster bootstrap --peer-host xxxx --peer-port xxx --replica-host yyyy --replica-port yyy [--join]
```

With `--join`, node to join cluster after replication(makes dynomite state to normal).
Defaults replication only, this is used when creating a backup machine.

## Help

```
NAME:
   dynomite-cluster

USAGE:
   dynomite-cluster [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                  debug mode
   --verbose, -V                verbose. more message
   --help, -h                   show help
   --version, -v                print the version
```
