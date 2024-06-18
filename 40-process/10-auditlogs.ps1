<#---
title: Download audit logs
description: Download audit logs from the auditlog service
tag: audit
api: post
---#>

param (
    $year = 2024,
    $month = 6,
    $day = 17
)
if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"
Push-Location
Set-Location $workdir
if (Test-Path "audit") {
    Remove-Item -Path "audit" -Recurse -Force
}

magic-mix download auditlog audit $year $month $day
magic-mix upload audit
magic-mix move mix files sharepoint.pageviews events insert_audit_records


Pop-Location