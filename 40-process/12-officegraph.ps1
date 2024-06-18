<#---
title: Download audit logs
description: Download audit logs from the auditlog service
tag: batch
api: post
---#>

param (
    $batchname = "users",
    $batchdefition = "users.json"
)
if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
    $path = join-path $PSScriptRoot ".." ".."
}
else {
    $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path 
$jsonFile = join-path $koksmatDir "data" $batchdefition

write-host "Workdir: $workdir"
Push-Location
Set-Location $workdir


if (Test-Path $batchname) {
    Remove-Item -Path $batchname -Recurse -Force
}


magic-mix download batch $batchname $jsonFile
magic-mix sql exec mix "delete from importdata where name ilike '$batchname/%'"
magic-mix upload $batchname
# magic-mix move mix files sharepoint.pageviews events insert_audit_records


Pop-Location