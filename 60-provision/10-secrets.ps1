<#---
title: Secrets
tag: secrets
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
  return $value 
}



# $envs += env "PNPAPPID" $env:PNPAPPID
# $envs += env "PNPTENANTID" $env:PNPTENANTID
# $envs += env "PNPCERTIFICATE" $env:PNPCERTIFICATE
# $envs += env "PNPSITE" $env:PNPSITE
# $envs += env "SITEURL" $env:SITEURL
$NATS = env "NATS" "nats://nats:4222"
$DB = env "POSTGRES_DB" $env:POSTGRES_DB


<#
Then we build the deployment file
#>

$image = "$($imagename)-app:$($version)"

$config = @"
apiVersion: v1
kind: Secret
metadata:
  name: postgres-secret
type: Opaque
stringData:
  postgres-url: "$DB"

---
apiVersion: v1
kind: Secret
metadata:
  name: nats-secret
type: Opaque
stringData:
  nats-url: "$NATS"
"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -