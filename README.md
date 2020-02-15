pai7
---
pai7 is the Go implementation of [Domino Card Game](https://en.wikipedia.org/wiki/Domino_(card_game)), AKA [排7游戏](https://zh.wikipedia.org/wiki/%E6%8E%92%E4%B8%83) in Chinese words.

Mostly developed as a telegram bot.


## Try in the terminal

```
$ go get github.com/scbizu/pai7
$ pai7 auto
```

## Deploy (On Telegram)

### As a binary

```shell
# your bot token
$ export BOT_KEY = xxxxx
# your webhook server port
$ export LISTENPORT = 1234
# open debug logs
$ export IS_DEBUG_MODE = true
$ pai7 server
```

### As a Docker Container (recommend)

```shell
$ docker pull scnace/pai7:latest
$ docker run  -d \
-it \
-p ${port}:${port} \
-e LISTENPORT=${port} \
-e BOTKEY=${bot_token} \
--rm \
--name pai7 scnace/pai7
```

## Online Bot

t.me/Pai7Bot
