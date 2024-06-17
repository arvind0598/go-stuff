# Sukasa Go

A simple project to figure out how Go works. REST APIs supported for reserving and resetting seats.

## Table of Contents
- [Local Development Setup](#local-development-setup)
- [Project Structure](#project-structure)
- [Api Documentation](#api-documentation)

## Local Development Setup

To set up the project for local development, follow these steps:
To set up the project for local development, follow these steps:

1. Make sure you have Docker and Go installed on your machine.
1. Clone the project repository.
1. Install the project dependencies by running:
    ```
    go install
    ```
1. Make a .env file based on .env.template
1. Start the project using Docker Compose:
    ```
    docker-compose up
    ```
1. Once the containers are up and running, you can test the project by accessing `http://localhost:8080` in your web browser. If everything is working fine, you should see a 'Hello World!'

Remember to stop the containers when you're done with the local development environment.

## Project Structure

The project follows a standard Go project structure:

## Api Documentation

There are four APIs. They all return in `text/plain` instead of JSON because I couldn't figure out a clean way to implement error handling.

1. `GET /` 
    A _Hello World!_ route.
1. `POST /login`
    Takes credentials, returns JWT.
1. `POST /reserve` 
    Authenticated route that takes seat number, returns success or error message.
1. `POST /reset` 
    Authenticated route that takes seat number, returns success or error message.


