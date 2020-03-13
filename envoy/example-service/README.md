# Envoy Example Service
This directory contains a simple Go service for use when learning Envoy. The
example service is a straightforward Server application that implements the
`/service` and `/trace` endpoints. The `/trace` endpoint is hardcoded to call
service #2 from service #1.

This service is built and run using Docker Compose within the Envoy examples.

## Docker Build
The container build leverages a multi-stage Docker build to reduce container
size and enable deployment of the service into a `scratch` container. A base
Envoy container has the built service added to it and the `start-service.sh`
script starts both services.
