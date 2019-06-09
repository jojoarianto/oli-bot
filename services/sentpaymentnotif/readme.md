# Push Message

API Documentation

```bash
# move directory
cd webhook

# deploy webhook
gcloud functions deploy message --runtime go111 --env-vars-file ../../env.yaml --entry-point SendPaymentNotif --region asia-northeast1 --trigger-http
```
