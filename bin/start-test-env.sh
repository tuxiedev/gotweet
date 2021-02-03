#!/bin/bash

set -xe

docker-compose -f kafka.docker-compose.yml up -d

#TODO make this command wait until the broker is up
docker exec -it gotweet_kafka_1 kafka-topics --create --topic tweets \
    --partitions 1 --replication-factor 1 \
    --zookeeper zookeeper:2181