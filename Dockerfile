FROM ubuntu:latest
LABEL authors="thale"

ENTRYPOINT ["top", "-b"]