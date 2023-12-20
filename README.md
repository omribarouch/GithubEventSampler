# Github Event Handler

Welcome to the Github Event Handler! This application shows github latest events.

## Setup Instructions

first you'll need to generate a github api token and paste it in the docker-compose.yaml
instead of the `<github api token>`.

then simply run the following cmd command (make sure that you have docker and docker compose installed):
```cmd
docker-compose up -d
```

## Architecture

The system consists of 3 services:
- Event Sampler: every 2 minutes fetches public events from github and produces them into a kafka messaging queue.
- Event Processor: subscribes to the kafka and process the raw messages to storable db models, which then saved in a postgres database.
- Api Server: serves an api for retrieving data on the processed events and their related actors and repositories.

## API

The application is accessible via the following api:

## Events

#### Get all events

**Endpoint:** `GET /events`

## Repositories

#### Get 20 most recent repositories

**Endpoint:** `GET /repositories/recent`

## Actors

#### Create an event

**Endpoint:** `POST /repositories/recent`