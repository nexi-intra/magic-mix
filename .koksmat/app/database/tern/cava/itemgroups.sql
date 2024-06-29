--CREATE VIEW "cava"."items" AS
SELECT
   

    data_element ->> 'ID' as id,
    data_element ->> 'Order' as Order,
    data_element ->> 'Title' as Title,
    data_element ,
    name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name iLIKE '%cava3%/%itemgroup%.json'
    AND data <> 'null'
    and data <> '{}'
    and data_element ->> 'Title' <> 'null'

