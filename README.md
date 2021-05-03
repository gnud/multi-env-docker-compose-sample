# Instructions

This is a demonstration of a Multi Docker Compose.


To create a new environment

```bash
mkdir platform-new
touch platform-new/.env
ln -s docker-compose.base.yml platform-new/docker-compose.yml
touch platform-new/docker-compose.override.yml
```

Basically we link the original docker base compose and we have new .env file for each platform environment.
Inside each .env we can control the inside port and the outside port.

Each override has the example structure:

```yaml
version: '3'

services:
    backend1:
        # add content here
        volumes:
            - ./data/backend3/config:/config
```

We can create a data directory for each service an mount/bind few things.

Note: each platform environment will not to be committed to GIT.
The platform environments will be created on the dev machine or production.


Environment keys:

- PORT_BACKEND1 is the outside port

- PORT_BACKEND1_IN is the inside port

- BACKEND1_PORT is the listen port for the service

# Commands

## Rebuild

```
cd platform-dev  # sample env
docker-compose build
docker-compose up -d
```

Test: http://localhost:8003

## Status

```
cd platform-dev  # sample env
docker-compose ps
```

# Features

- Multi environment docker compose

- Building an image

- Controlling inside env vars

- Controlling outside env vars in docker -compose


# Reference

- https://pspdfkit.com/blog/2018/how-to-manage-multiple-system-configurations-using-docker-compose/


- https://github.com/docker/awesome-compose/ (for the Go sample insipiration)


# TODOS

Please check the Issues section for a complete list of pending features **[here](https://github.com/gnud/multi-env-docker-compose-sample/issues/)**.
