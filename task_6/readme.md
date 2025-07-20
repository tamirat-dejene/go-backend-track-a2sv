# Task Management REST API using Go and Gin Framework

## task_manager_v1 is an enhanced version of task_manager in task_4 . It uses
 - MongoDb for data persistence
 - Containerized using Docker

## Comprehensive api documentation is available here: 
[API Documentation](https://documenter.getpostman.com/view/33664366/2sB34ijKPv)

---

## Running with Docker

### rerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed

### Run the app

```bash
docker-compose up --build
```

## If you make changes to code
Rebuild the app with

```bash
docker-compose up --build --force-recreate
```

## Access

Once running, the API is accessible at:

```bash 
http://localhost:1337
```

Database name: taskdb
Collection name: tasks


##To Stop & Clean Up

```bash
docker-compose down
```

## To remove volumes and network:

```bash
docker-compose down -v --remove-orphans
```