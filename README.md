# oli-bot

Olimpiade.id bot, Line App

## Purpose

As an organizer who held olimpiade event, we need some information faster for some cases, especially about receiving payment proof from participants. so we create this tools (line app) to send notifcation via line when new payment comes. further we'll also develope some another necessary features.

## How to use it

- 

## Official line

Barcode

### Version 0.1

- Feature Line notification when new payment proof from participant comes

## Deploy

Deploy function on google cloud function

first setup project id

```bash
# fill with project id
PROJECT_ID=<project-id>

# set config to spesific projectid
gcloud config set project $PROJECT_ID
```

go to directory you want to deploy
```bash
# move directory
cd webhook

# deploy webhook
gcloud functions deploy webhook --runtime go111 --env-vars-file ../env.yaml --entry-point Webhook --region asia-northeast1 --trigger-http
```

delete function
```bash
# hello is function name
gcloud functions delete functionname
```
