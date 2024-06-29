--CREATE VIEW "cava"."items" AS
SELECT
   

    data_element ->> 'ID' as id,
    data_element ->> 'Order' as Order,
 
    data_element ->> 'Title' as Title,
  data_element -> 'Location' ->> 'LookupValue' as Location,
data_element -> 'Provider' ->> 'LookupValue' as Provider,
data_element ->> 'Description' as Description,
data_element ->> 'Delivery_x0020_Instructions_x002' as Delivery_Instructions,

    data_element ,
        name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name iLIKE '%cava3%/pricelists2.json'
    AND data <> 'null'
    and data <> '{}'
    and data_element ->> 'Title' <> 'null'

