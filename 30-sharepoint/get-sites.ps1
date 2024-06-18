<#---
title: Get All Sites in Tenant
tag: allsites
api: post
connection: sharepoint
output: allsites.json
---
#>


if ($null -eq $env:WORKDIR ) {
    $env:WORKDIR = join-path $psscriptroot ".." ".koksmat" "workdir"
}
$workdir = $env:WORKDIR

$workdir = Join-Path $workdir "sharepoint"
if (-not (Test-Path $workdir)) {
    New-Item -Path $workdir -ItemType Directory | Out-Null
}

koksmat trace log "Get-PnPTenantSite"
$sites = Get-PnPTenantSite -Detailed   # -Filter "Url -like 'nexiintra'"
koksmat trace log "Got $($sites.Count) sites"

$allsites = @()

foreach ($site in $sites) {
    try {
        $allsites += @{
            SiteId                  = $site.Id
            Url                     = $site.Url
            Title                   = $site.Title
            Owner                   = $site.Owner
            StorageQuota            = $site.StorageQuota
            StorageUsed             = $site.StorageUsed
            LastContentModifiedDate = $site.LastContentModifiedDate
            Created                 = $site.Created
            SharingCapability       = $site.SharingCapability.value__
            Status                  = $site.Status
            IsTeamsConnected        = $site.IsTeamsConnected
            IsTeamsChannelConnected = $site.IsTeamsChannelConnected
            HusSiteId               = $site.HubSiteId
            GroupId                 = $site.GroupId
            IsHubsite               = $site.IsHubsite
            LocaleId                = $site.LocaleId
            RelatedGroupId          = $site.RelatedGroupId
        }
    }
    catch {
        write-host $site.Url "Error: $($_.Exception.Message)" -ForegroundColor Red
        # Do this if a terminating exception happens
    }
 
    # $currentItemName is the current item 
}


$outputfile = Join-Path $workdir  "allsites.json"

$allsites | ConvertTo-Json -Depth 10 | Out-File -FilePath $outputfile -Encoding utf8NoBOM
# magic-files import $outputfile