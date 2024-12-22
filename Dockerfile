# Dockerfile for running app which interacts with k8s api as a pod in k8s cluster

FROM ubuntu

COPY ./go-k8s ./go-k8s

ENTRYPOINT [ "./go-k8s" ]