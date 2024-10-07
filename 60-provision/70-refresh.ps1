<#---
title: REFRESH MATERIALIZED VIEW importdata_summary
---
#>

$config = @"
apiVersion: batch/v1
kind: CronJob
metadata:
  name: refresh-importdata-summary
spec:
  schedule: "*/15 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: magic-mix
            image: ghcr.io/magicbutton/magic-mix-app:latest
            command: ["magic-mix"]
            args: ["sql", "exec", "mix", "REFRESH MATERIALIZED VIEW importdata_summary"]
            env:
            - name: POSTGRES_DB
              valueFrom:
                secretKeyRef:
                  name: postgres-secret
                  key: postgres-url
            - name: NATS
              valueFrom:
                secretKeyRef:
                  name: nats-secret
                  key: nats-url
          restartPolicy: OnFailure
"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -