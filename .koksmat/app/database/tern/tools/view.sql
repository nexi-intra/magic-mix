select 
weburl,
data_element -> 'fields' ->> 'Title' as title,
data_element -> 'fields' ->> 'Guides' as Guides,
data_element -> 'fields' ->> 'Description' as Description,
data_element -> 'fields' -> 'Document_x0020_Link' ->> 'Url' as Url,
data_element -> 'fields' ->> 'AreaLookupId' as AreaLookupId,
data_element 

from
(SELECT item as data_element, item ->> 'webUrl' as weburl FROM (
SELECT
    jsonb_array_elements(data_element  -> 'data' -> 'details' -> 'value')  AS item
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name ilike 'list-nexiintra-country-it/listitems.json') 
    AS items
    
    
    ) as x 


where weburl ilike '%/Lists/Applications/%'