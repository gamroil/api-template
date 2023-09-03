# API Template

## About
This is an opinionated API template that comes with a lot of the boilerplate code for creating APIs. The API is easily
configurable and comes with a dockerfile, structured logging, and prometheus metrics exported at the `/metrics`
endpoint. There are make targets for building a binary, building and running a docker-image, and linting.

## Quickstart
To quickly start up the API, run `make build-image` and then `make run-image`. This will build and run a docker
image of the API. Alternatively, run `make build` to build a binary and then `./api-template serve` to run the newly
built `api-template` binary.

By default the API will run on port 8080, but this can be changed by configuration

## CLI
The binary comes packaged as a CLI with two commands:

| Command      | Description |
| ----------- | ----------- |
| version      | Print the version of the binary. This should be the git commit of the branch at the time the binary was built       |
| serve   | Start the API        |

The API also comes with flags:
| Flag | Shorthand | Description |
| ----------- | ----------- | ----------- |
| --config |   -c | Run the binary using the specified config file|

## Configuration
The API configuration is managed by the [Viper](https://github.com/spf13/viper) library. This allows customization via
`config.yaml` or any other config file passed in via the `--config` argument. The values in the config yaml can be
overridden by specifying environment variables. The key for configuration environment variables follows the
`{PREFIX}_{ENVIRONMENT_VALUE}` format. The prefix is `OVERRIDE` by default but that can be changed in the config code.
The remaining part of the environment variable key is the path to the yaml in all caps with underscores separating the
different words. Refer to the [Viper](https://github.com/spf13/viper) GitHub page for more details.

## Metrics
Metrics are exported by Prometheus at `/metrics`. There is sample code that creates a new metric that tracks total
requests to one of the endpoints. This can be extended to any metric that Prometheus supports.

## Docker
The API comes with a dockerfile. It builds in two stages and the end result is a roughly 12MB docker image.
`make build-image` builds and `make run-image` runs the image.

## References
https://github.com/spf13/cobra  
https://github.com/spf13/viper  
https://prometheus.io/docs/concepts/metric_types  
