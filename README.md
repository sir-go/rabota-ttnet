# Application form
[![backend](https://github.com/sir-go/rabota-ttnet/actions/workflows/go.yml/badge.svg)](https://github.com/sir-go/rabota-ttnet/actions/workflows/go.yml)

> https://rabota.ttnet.ru

## What is it
This is a side project aimed at estimating the soft skills of applicants and getting a wide spectrum of 
optional information about applicants who are decided to join us.

The general idea is to check an applicant's patience and ability to focus on 
certain work for several minutes.

The collection of some extended information is an optional side effect.

## Configure
The backend has a mounted to docker container config file
`./backend/config.toml`

The frontend container is based on the `nginx` server and has mounted 
`./frontend/nginx.conf`

## Build & run
```bash
docker compose up -d
```

## Screencast
![](sc.gif)
