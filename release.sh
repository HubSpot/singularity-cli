#!/bin/bash
set -ex

rm -rf release/
mkdir release
go build
mv singularity-cli release
cp -r scripts release
cp config.toml release

tar -cvf singularity-cli.tar release
