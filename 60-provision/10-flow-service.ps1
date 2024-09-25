<#---
title: FLOW deploy to production
tag: flowdeployproduction
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
  name: $appname-flow
spec:
  selector:
    matchLabels:
      app: $appname-flow
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname-flow
    spec: 
      containers:
      - name: $appname-flow
        image: $image
        command: [$appname]
        args: ["flow"]               
        env:
$configEnv                           

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -