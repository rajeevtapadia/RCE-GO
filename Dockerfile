FROM alpine:latest

# install compilers and interpreters for c/cpp, js, python
RUN apk update && \
    apk add --no-cache \
    build-base \
    python3 \
    py3-pip \
    nodejs \
    npm \
    bash

# Create a directory for the application
WORKDIR /less-go

# Set the default command to bash
CMD ["bash"]