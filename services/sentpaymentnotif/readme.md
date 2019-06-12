# Push Message

API Documentation

```bash
# move directory
cd webhook

# deploy webhook
gcloud functions deploy message --runtime go111 --env-vars-file ../../env.yaml --entry-point SendPaymentNotif --region asia-northeast1 --trigger-http
```

## Monorepo note
```bash
go mod edit -replace github.com/jojoarianto/oli-bot/services/api/line=./line
```
setting up multi module single repo


reference https://github.com/golang/go/issues/27056
