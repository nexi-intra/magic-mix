CREATE VIEW "project_lotus"."excelusers" AS
select  
-- id
-- ,"Internal/External (column 0)"
-- ,"Vendor (column 1)"
-- ,"BU (column 2)"
-- ,"Hierarchy Level 5 (column 3)" 
"Employee No. (column 4)" as nets_employeeid
--,"Employee Name (column 5)" as name
,"User Name (column 6)" as nets_username
-- ,"Position Name (column 7)"
-- ,"Employed first time (column 8)"
-- ,"Location (column 9)"
-- ,"Hiring Status (column 10)"
-- ,"Total (column 11)"
from excelimport."Sheet1";