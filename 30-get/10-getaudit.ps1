if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"
Set-Location $workdir

# magic-mix download auditlog audit 2024 5 19
# magic-mix download auditlog audit 2024 5 20
# magic-mix download auditlog audit 2024 5 21
# magic-mix download auditlog audit 2024 5 22
# magic-mix download auditlog audit 2024 5 23
# magic-mix download auditlog audit 2024 5 24
# magic-mix download auditlog audit 2024 5 25
# magic-mix download auditlog audit 2024 5 26
# magic-mix download auditlog audit 2024 5 27
# magic-mix download auditlog audit 2024 5 28
# magic-mix download auditlog audit 2024 5 29
# magic-mix download auditlog audit 2024 5 30
# magic-mix download auditlog audit 2024 5 31
# magic-mix download auditlog audit 2024 6 1
# magic-mix download auditlog audit 2024 6 2
# magic-mix download auditlog audit 2024 6 3
# magic-mix download auditlog audit 2024 6 4
# magic-mix download auditlog audit 2024 6 5
# magic-mix download auditlog audit 2024 6 6
# magic-mix download auditlog audit 2024 6 7
# magic-mix download auditlog audit 2024 6 8
# magic-mix download auditlog audit 2024 6 9
# magic-mix download auditlog audit 2024 6 10
# magic-mix download auditlog audit 2024 6 11
magic-mix download auditlog audit 2024 6 12


