FROM golang:1.14-alpine

# Set the Current Working Directory indide the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Prsent Working Directory) inside the container
COPY bin/filesynchronizer_batch /app/filesynchronizer_batch

RUN chomd 775 /app

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executrable
CMD ["./filesynchronizer_batch"]