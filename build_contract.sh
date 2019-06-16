#!/usr/bin/env bash

# Exit script as soon as a command fails.
set -o errexit

echo "Generating raw Bin Challenge.sol"
solc --bin ./Contract/Challenge.sol -o Contract 

echo "Generating abi of Contract.sol"
solc --abi ./Contract/Challenge.sol -o Contract

echo "Dumping them in a package"
abigen --bin=./Contract/Challenge.bin --abi=./Contract/Challenge.abi --pkg=challenge --out Contract/challenge.go

echo "All Done :)"
