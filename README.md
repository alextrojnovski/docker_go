## go web application

A simple go web application that displays local and environment variables, packaged in a Docker container.
# Abilities
- Local variable output `MY_VAR`
- Output of the environment variable
- Configuring a port via an environment variable `PORT`
- Docker container with minimal image

To run the app u'll need:
- Go 1.21 or upper
- Git
- Docker 20.10 or upper
# Cloning repository
-git clone https://github.com/alextrojnovski/docker_go

-cd docker_go

# Start in docker

-docker build -t docker_go .

-docker run -d -p 8080:8080 -e MY_VAR="Hello from Docker!" --name myapp docker_go
# Funcional check
-curl http://localhost:8080
