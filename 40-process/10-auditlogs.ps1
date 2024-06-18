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

# magic-mix download auditlog audit 2024 5 17
# magic-mix upload audit
# magic-mix move mix files sharepoint.pageviews events


Pop-Location