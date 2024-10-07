<#---
title: Box deploy to production
tag: boxdeployproduction
xapi: post
---
#>

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
  $path = join-path $PSScriptRoot ".." ".."
}
else {
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path

$inputFile = join-path  $koksmatDir "koksmat.json"

if (!(Test-Path -Path $inputFile) ) {
  Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"
$appname = $json.appname
$imagename = $json.imagename

<#
Envs
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
$envs += env "APPCLIENT_ID" $env:APPCLIENT_ID
$envs += env "APPCLIENT_SECRET" $env:APPCLIENT_SECRET
$envs += env "APPCLIENT_DOMAIN" $env:APPCLIENT_DOMAIN

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

$image = "$($imagename)-app:$($version)"

$config = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname-box
spec:
  selector:
    matchLabels:
      app: $appname-box
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname-box
    spec: 
      containers:
      - name: $appname-box
        image: $image
        command: ["/bin/sh"]
        args: ["-c", "while true; do sleep 3600; done"]               
        env:
$configEnv                           

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -