
# Task Management REST API using Go and Gin

**Comprehensive API Docs**: [View in Postman](https://documenter.getpostman.com/view/33664366/2sB34ijKPv)

## Overview

`task_manager_v1.2` is an enhanced version of `task_manager_v1.1` (originally from `task_6`), now following **clean architecture** principles.

## Getting Started with Docker

### Prerequisites

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)


### Setup Instructions

1. **Create a `.env` file** in the project root with the following variables:

```
APP_ENV=development
SERVER_ADDRESS=:<your_server_port>
CTX_TIMEOUT=<context_timeout_in_seconds>
MONGO_URI="mongodb://<your_mongo_host>:<port>"
DB_NAME=<your_database_name>
REFRESH_TOKEN_SECRET=<your_refresh_token_secret>
ACCESS_TOKEN_SECRET=<your_access_token_secret>
REFRESH_TOKEN_EXPIRY_HOUR=<refresh_token_expiry_hours>
ACCESS_TOKEN_EXPIRY_HOUR=<access_token_expiry_hours>
```
> **Note:** You can freely modify the configuration values, especially `MONGO_URI` and `DB_NAME`, depending on your database setup.
>
> * For **local MongoDB**, use something like `mongodb://localhost:27017`.
> * For **MongoDB Atlas**, use the full connection string provided by Atlas.
> * Or configure however your environment requires.

> Do not commit your `.env` file to version control.

### Run the App

```bash
docker-compose up --build
```

### If You Make Code Changes

Rebuild and restart the app:

```bash
docker-compose up --build --force-recreate
```

### Access the API

Once running, the API is available at:

```
http://localhost:<your_server_port>
```

### Stop & Clean Up

To stop containers:

```bash
docker-compose down
```

To remove volumes and networks:

```bash
docker-compose down -v --remove-orphans
```
