# Anonymous messaging using mix networks

This is an implementation of an private communications system based on
Ania Piotrowska's PhD research. It implements a 
[Loopix](https://arxiv.org/abs/1703.00536) mixnet as well as the 
[Sphinx](https://cypherpunks.ca/~iang/pubs/Sphinx_Oakland09.pdf) packet format.

## Setup

To build and test the code you need:

* Go 1.12 or later

To perform the unit tests run:

```shell
go test ./...
```

Before first fresh run of the system run:

```shell
./scripts/clean.sh
```

This removes all log files and database.

## Usage

To run the network, i.e., mixnodes and providers run

```shell
./scripts/run_network.sh
```

This spins up 3 mixnodes and 1 provider. To change the number of mixnodes to 5, 
do:

```shell
./scripts/run_network.sh 5
```

To simulate the clients run

```shell
./scripts/run_clients.sh
```
