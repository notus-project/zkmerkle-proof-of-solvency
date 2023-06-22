#!/usr/bin/env bash

# source ./setup.sh
source ./reset_docker.sh

cd ../src/witness
go build -o witness.out
./witness.out
rm witness.out

cd ../prover
go build -o prover.out
./prover.out
rm prover.out