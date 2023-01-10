<div id="top"></div>

<br />
<div align="center">
  <a href="https://github.com/Skryvvara/common-services">
      <img src="./.github/assets/skryvvara.webp" width=124 height=124 alt="Logo" style="border-radius: 100%; border: 2px solid rgba(255, 255, 255, 0.8)">
  </a>

  <h3 align="center">Common-Services-API</h3>

 <p align="center">
    This is a work in progress.
  </p>
</div>

<!-- ABOUT -->

## About The Project

This API is intended to be consumed by multiple other projects and is a central point to implement simple backend needs without rolling out an entire backend for every little project.

<p align="right">(<a href="#top">back to top</a>)</p>

## Build

### Manual

1. Run make build

```
make build
```

### Docker

1. Run docker build

```
docker build .
```

## Configuration

The configuration of the API can be either done via the toml file or using env variables.

### TOML

1. copy the example-config.toml to a valid destination

`cp example-config.toml ./src/config.toml`

`cp example-config.toml ./src/data/config.toml`

or set a custom config path via env

> :warning: Note that the env var also has to be set when running the application.

`export CONFIG_PATH="./src/my-custom-config/config.toml"; cp example-config.toml ./src/my-custom-config.config.toml`

2. Change the placeholder values to your desired values

3. Run the application

### ENV

1. Set the env variables to the desired values. Possible values are:

```
SERVER_ENVIRONMENT        # string
SERVER_PORT               # int
SERVER_TIMEZONE           # string
SERVER_POWERED_BY_HEADER  # string

MAIL_HOST                 # string
MAIL_PORT                 # int
MAIL_SECURE               # boolean
MAIL_USER                 # string
MAIL_PASSWORD             # string
MAIL_TO                   # string

LOG_DIR                   # string
LOG_FILE                  # string
LOG_MAX_SIZE              # int
LOG_BACKUPS               # int
LOG_MAX_AGE               # int
LOG_COLOR                 # boolean
```

You can also use a mix of toml and env variables. The load order is as followed.

`ENV` wins over `TOMl` wins over `DEFAULTS`

<p align="right">(<a href="#top">back to top</a>)</p>

## Built With 🛠️

- [Go][go]

<p align="right">(<a href="#top">back to top</a>)</p>

[go]: https://go.dev
[react]: https://reactjs.org
