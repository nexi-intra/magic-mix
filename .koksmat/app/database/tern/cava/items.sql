--CREATE VIEW "cava"."items" AS
SELECT
   

    data_element ->> 'ID' as id,
    data_element ->> 'Order' as Order,
    data_element ->> 'Price' as Price,
    data_element ->> 'Title' as Title,
  data_element -> 'Currency' ->> 'LookupValue' as Currency,
data_element -> 'Provider' ->> 'LookupValue' as Provider,
data_element ->> 'Description' as Description,
    data_element ,
        name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name iLIKE '%cava3%/items.json'
    AND data <> 'null'
    and data <> '{}'


