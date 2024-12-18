<#---
title: REFRESH MATERIALIZED VIEW importdata_summary
---
#>

$envs = @()
function env($name, $value ) {
  if ($null -eq $value) {
    throw "Environment value for $name is not set"
  }
  return @{name = $name; value = $value }
}



# $envs += env "PNPAPPID" $env:PNPAPPID
# $envs += env "PNPTENANTID" $env:PNPTENANTID
# $envs += env "PNPCERTIFICATE" $env:PNPCERTIFICATE
# $envs += env "PNPSITE" $env:PNPSITE
# $envs += env "SITEURL" $env:SITEURL
$envs += env "NATS" "nats://nats:4222"
$envs += env "POSTGRES_DB" $env:POSTGRES_DB
$configEnv = ""
foreach ($item in $envs) {

  $configEnv += @"
          - name: $($item.name)
            value: $($item.value)

"@
}

<#
Then we build the deployment file
#>
$config = @"
apiVersion: batch/v1
kind: CronJob
metadata:
  name: magic-mix-refresh-mviews
spec:
  schedule: "0 4 * * *"  # 04:00 UTC, corresponding to 06:00 AM CET
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: magic-mix
            image: ghcr.io/magicbutton/magic-mix-app:latest
            command:
              - /bin/sh
              - -c
              - |
                cd
                echo "POSTGRES_DB=$POSTGRES_DB" > .env

                echo Starting to refresh materialized views in mix database
                magic-mix sql query mix "SELECT * FROM public.refresh_all_mviews()"
                echo Done refreshing materialized views in mix database

                echo Starting to refresh materialized views in files database
                magic-mix sql query files "SELECT * FROM public.refresh_all_mviews()"
                echo Done refreshing materialized views in files database

            env:
$configEnv                           
          restartPolicy: OnFailure
"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -