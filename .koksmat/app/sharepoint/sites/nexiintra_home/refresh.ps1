# koksmat sharepoint export-template "https://christianiabpos.sharepoint.com/sites/nexiintra-home" > "$PSScriptRoot/template.xml"
koksmat scaffold "$PSScriptRoot/template.xml" > "$PSScriptRoot/structs.go"