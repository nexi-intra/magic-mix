SELECT 
row_number,
column_13 as Num,
column_14 as Den,

* FROM devicekpi.nexi
where column_4 ilike '%#workstations with antivirus installed and monitored/# total workstations%' 
order by row_number()