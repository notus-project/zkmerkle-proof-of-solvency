#!/usr/bin/env bash

# source ./setup.sh
source ./reset_docker.sh

sleep 5

cd ../src/witness
go build -o witness.out
./witness.out
rm witness.out

cd ../prover
go build -o prover.out
./prover.out
rm prover.out