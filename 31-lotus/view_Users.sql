CREATE SCHEMA IF NOT EXISTS "azure";
DROP VIEW IF EXISTS "azure"."users";
CREATE VIEW "azure"."users" AS

SELECT
  name,
  data_element  ->> 'id' AS id,
  data_element  ->> 'displayName' AS displayName,
  data_element  ->> 'mail' AS mail,
  (data_element -> 'onPremisesExtensionAttributes' ->> 'extensionAttribute4')::text   as extensionAttribute4,
  (data_element -> 'onPremisesExtensionAttributes' ->> 'extensionAttribute11')::text   as extensionAttribute11
 
FROM
  importdata,
  jsonb_array_elements(data) AS data_element

WHERE
  name like 'users/parents.json'
  --and (data_element -> 'onPremisesExtensionAttributes' ->> 'extensionAttribute4')::text ILIKE 'nets%'

