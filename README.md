# gotweet

Stream tweets to some number of outputs


## Getting Started

### Get twitter credentials

1. Get an approved Twitter developer account https://developer.twitter.com/en/apply-for-access
2. Create an app in the Twitter Developer portal https://developer.twitter.com/en/docs/apps/app-management
3. Get the following information:
    * API Key
    * API Secret
    * Access Token
    * Access Secret
4. Run the binary
```
./gotweet console
```

The goal of this project is to sink tweets into different outputs. Hence, the submodule of the top level commands will be the name of the output the tweets are to be produced to

```
$ ./gotweet --help
For more information, checkout github.com/tuxiedev

Usage:
  gotweet [command]

Available Commands:
  console     Produces tweets to console
  help        Help about any command
  kafka       Produces tweets to Kafka

Flags:
  -h, --help                           help for gotweet
  -k, --keywords stringArray           keywords to filter the stream on
      --twitter-access-secret string   REQUIRED: Twitter access secret
      --twitter-access-token string    REQUIRED: Twitter access token
      --twitter-api-key string         REQUIRED: Twitter API key
      --twitter-api-secret string      REQUIRED: Twitter API secret

Use "gotweet [command] --help" for more information about a command.
```

### Using Environment variables for twitter credentials
The twitter credentials can be passed as environment variables to the command. Here is a map
| CLI Argument | Environment Variable |
| -------- | ------ |
| `twitter-access-token` | `TWITTER_ACCESS_TOKEN` | 
| `twitter-access-secret` | `TWITTER_ACCESS_SECRET` | 
| `twitter-api-key` | `TWITTER_API_KEY` | 
| `twitter-api-secret` | `TWITTER_API_SECRET` |  

### Producing tweets to different outputs

#### Kafka
You can use the `docker-compose.yml` in the project to start a single broker kafka cluster on local
```
$ docker-compose -f kafka.docker-compose.yml up -d 
```
A topic to produce tweets to can be created using 
```
docker exec -it gotweet_kafka_1 kafka-topics --create --topic tweets \
    --partitions 1 --replication-factor 1 \
    --zookeeper zookeeper:2181
```
Now, the you are ready to produce tweets to the locally running Kafka server
```
$ ./gotweet kafka --bootstrap-brokers localhost:9092 -output-topic tweets
```

## Development
### Build
```
go build
```

