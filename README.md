# GO API Template

This is an api template built using chi framework for creating REST APIs using Go

## Packages

This template comes with some packages useful for getting started with Go.

- godotenv - For loading env variables
- chi - The core framework for creating REST APIs
- pq - For connecting to postgres DB. Can use anyother DB of choice.
- air - To make development a lot faster with hot reloads

## Steps to Use this template

1. Click on Use template button on the top and create a new repository
2. Create a .env file by taking reference from .env.example
3. Change the app_name in the Makefile and the app_name in the Dockerfile to generate a binary with name of your choice
4. Run `make install` to install air
5. In dev enviroment, run `make dev` to start air and run the server
6. Use `make run` for prod environments
7. The included workflows will be triggered according to the branch once a push event occurs

