SELECT 
    details ->> 'id' as id,
    details ->> 'eTag' as eTag,
    details -> 'fields' as fields,
    parent ->> 'name' as listname 
FROM (
    SELECT 
        data_element -> 'data' -> 'parentId' as parent,
        jsonb_array_elements(CASE 
            WHEN jsonb_typeof(data_element -> 'data' -> 'details') = 'array' 
            THEN data_element -> 'data' -> 'details'
            ELSE '[]'::jsonb 
        END) as details
    FROM 
        importdata, 
        jsonb_array_elements(data) AS data_element
    WHERE 
        name ILIKE '%cava3%'
) as items
--LIMIT 200;
