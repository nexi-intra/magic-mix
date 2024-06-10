CREATE VIEW "azure"."users" AS
SELECT
    name,
    data_element ->> 'id' AS id,
    data_element ->> 'displayName' AS displayName,
    data_element ->> 'mail' AS mail,
(data_element -> 'onPremisesExtensionAttributes' ->> 'extensionAttribute4')::text AS extensionAttribute4,
(data_element -> 'onPremisesExtensionAttributes' ->> 'extensionAttribute11')::text AS extensionAttribute11
FROM
    importdata,
    jsonb_array_elements(data) AS data_element
WHERE
    name LIKE 'users/parents.json';