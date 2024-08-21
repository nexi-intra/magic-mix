--drop view if exists sharepoint.cava3 cascade;
create or replace view sharepoint.cava3 as
select 
*
from (
select 
listname, 
batchname,
detail_element -> 'fields' as field

FROM (
SELECT 
    parent ->> 'name' as listname,
    batchname,
    jsonb_array_elements(details) AS detail_element
  
FROM (
    SELECT
        importdata.data -> 'parentId' AS parent,
        importdata.data -> 'details' AS details,
        importdata.name AS batchname
    FROM
        importdata
    WHERE
        importdata.name::text ILIKE 'sharepoint-cava3/%.json'::text
) AS subquery
WHERE
    jsonb_typeof(details) = 'array') as subquery2
) as subquery3

--LIMIT 100;
