CREATE VIEW "sharepoint"."pageviews" AS
SELECT
    data_element -> 'auditData' ->> 'CreationTime' AS CreationTime,
    data_element -> 'auditData' ->> 'UserId' AS userid,
    data_element -> 'auditData' ->> 'Site' AS site,
    data_element -> 'auditData' ->> 'ObjectId' AS ObjectId,
    data_element -> 'auditData' ->> 'Operation' AS Operation,
    data_element -> 'auditData' ->> 'Platform' AS Platform,
    data_element ->> 'auditData' AS auditData,
    name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name LIKE '%audit/records%.json'
    AND data <> 'null'
    AND data_element -> 'auditData' ->> 'Operation' LIKE 'Page%'