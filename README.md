# go-fuzzing

Example project utilizing [fuzzing](https://pkg.go.dev/testing@master#hdr-Fuzzing) which is currently in beta and planned to be released with Go 1.18.

## How to run

1. `make setup`
    - Only needs to be run once
3. `make test`
    - The test run might take up until ~30 seconds before finding the weak spot
