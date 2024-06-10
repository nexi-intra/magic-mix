-- sdfds
CREATE VIEW "azure"."groups" AS SELECT
    name,
    data_element ->> 'id' AS id,
    data_element ->> 'displayName' AS displayName,
    data_element ->> 'mail' AS mail,
    data_element ->> 'groupTypes' AS groupTypes,
    data_element ->> 'membershipRule' AS membershipRule,
    data_element ->> 'onPremisesDomainName' AS onPremisesDomainName,
    data_element ->> 'securityEnabled' AS securityEnabled,
    data_element ->> 'visibility' AS visibility


FROM
    importdata,
    jsonb_array_elements (data) AS data_element
WHERE
    name LIKE 'groups-all/parents.json';