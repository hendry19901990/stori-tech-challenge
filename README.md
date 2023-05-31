# My Go Application

This is a Go application that can be run using Docker and Docker Compose.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your machine.

## Getting Started

1. Clone the repository:

```
git clone https://github.com/your/repo.git
cd repo
```


2. Build the Docker image using the provided Dockerfile:

```
docker build -t my-go-app .
```


3. Start the application and MongoDB container using Docker Compose:

```
docker-compose up
```


4. View the application logs:

```
docker-compose logs -f app
```


## Command Line Arguments

The application accepts the following command line arguments:

- `email`: The email argument description.
- `txsFile`: The txsFile argument description.
- `senderEmail`: The senderEmail argument description.
- `senderPassword`: The senderPassword argument description.

To pass the command line arguments when running the Docker container, modify the `docker-compose.yaml` file.

## Data and Templates

The application expects two directories to be mounted:

- `data`: Mount this directory to provide the data directory.
- `templates`: Mount this directory to provide the templates directory.

Make sure the necessary files are present in these directories before running the application.

## Stopping the Application

To stop the application and remove the containers, press `Ctrl + C` in the terminal where Docker Compose is running.

