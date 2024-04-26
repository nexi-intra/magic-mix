<#
title: List Excel sheets
tags: sheets
description: List all sheets in an Excel file
output: excel-sheets.json

#>

# magic-mix excel sheets | Out-File -FilePath $HOME\Documents\list-sheets.txt -Encoding utf8

param (
    [string]$excelFile = "/Users/nielsgregersjohansen/kitchens/magic-apps/.koksmat/workdir/Estrazione Catalogo NEAR_20240404.xlsx"

)
Import-Module ImportExcel

$excelData = Import-Excel -Path $excelFile
write-host "Imported Excel file: $excelFile"


Import-Excel -