FROM alpine:3.14

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY actions_test /actions_test

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/actions_test"]