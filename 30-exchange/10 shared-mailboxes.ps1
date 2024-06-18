<#---
title: Get Shared Mailboxes
connection: exchange
tag: get-shared-mailboxes

---


#>

if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = Join-Path  $env:WORKDIR "shared-mailboxes"

if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

$workdir = Resolve-Path $workdir

write-host "Workdir: $workdir"

$result = Join-Path $workdir "parents.json"
if (Test-Path $result) {
    $items = Get-Content $result | ConvertFrom-Json
}
else {

    $items = Get-ExoMailbox -RecipientTypeDetails SharedMailbox   -ResultSize Unlimited
    | Select-Object Guid, DisplayName, PrimarySmtpAddress, RecipientTypeDetails, RecipientType, Identity, WindowsEmailAddress, ResourceCapacity
    
    $items | ConvertTo-Json -Depth 10
    | Out-File -FilePath $result -Encoding:utf8NoBOM
}

foreach ($mailbox in $items ) {
    $detailsPath = Join-Path $workdir "permissions-$($mailbox.PrimarySmtpAddress).json"
    if (Test-Path $detailsPath) {

        continue
    }
    else {
        write-host "Getting details for $($mailbox.PrimarySmtpAddress)"
        try {
            
       
            $mailboxDetails = Get-EXOMailboxPermission -Identity $mailbox.Guid
            $data = @{
                details = $mailboxDetails
                id      = $mailbox.Guid
            }
            $data | ConvertTo-Json -Depth 10 | Out-File -FilePath $detailsPath -Encoding:utf8NoBOM
        }
        catch {
     
            write-host "Error getting details for $($mailbox.PrimarySmtpAddress)" $_
            write-host "Process continues"
        }
    }
    <# $mailbox is the current item #>
}


write-host $result

