# GO API Template

This is an api template built using chi framework for creating REST APIs using Go

## Packages

This template comes with some packages useful for getting started.

- godotenv - For loading env variables
- chi - The core framework for creating REST APIs
- pq - For connecting to postgres DB. Can use anyother DB of choice.
- air - To make development a lot faster with hot reloads

## Steps to Use this template

1. Click on Use template button on the top and create a new repository
2. Clone the repo into a local directory
3. Create a .env file by taking reference from .env.example
4. Change the app_name in the Makefile and the app_name in the Dockerfile to generate a binary with name of your choice
5. Run `make install` to install air
6. In dev enviroment, run `make dev` to start air and run the server
7. Before pushing the code into a repo, add the following variables into github actions as repository secrets
   1. DOCKER_PAT - Create a Personal Access Token to push the images to an account
   2. DOCKER_USER - Username of the docker hub account
   3. SSH_HOST - Hostname of the SSH server
   4. SSH_KEY - Private key to authorize the SSH user
   5. SSH_USER - Username of the SSH server
   6. SSH_PORT - The port to access the SSH server
   7. WORK_DIR - The directory where the deployment files live in
8. The included workflows will be triggered according to the branch once a push event occurs
