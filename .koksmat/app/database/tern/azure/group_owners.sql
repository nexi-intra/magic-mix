CREATE VIEW azure.group_owners AS
SELECT (
        jsonb_array_elements (
            data_element -> 'data' -> 'details' -> 'value'
        ) ->> 'id'
    ) AS user_id,
    data_element -> 'data' ->> 'parentId' AS group_id
FROM
    importdata,
    jsonb_array_elements (data) AS data_element
WHERE
    name LIKE 'groups-all/owners.json';