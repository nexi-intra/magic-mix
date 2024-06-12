CREATE VIEW "sharepoint"."sites" AS
SELECT
    data_element  ->> 'webUrl' AS webUrl,
        data_element  ->> 'name' AS name,
            data_element  ->> 'id' AS id,
            data_element  ->> 'isPersonalSite' AS isPersonalSite,
    --data_element ,
    name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name = 'sites/parents.json'
