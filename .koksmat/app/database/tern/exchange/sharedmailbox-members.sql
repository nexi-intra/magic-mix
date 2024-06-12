DROP VIEW IF EXISTS "exchange"."sharedmailboxmembers";
CREATE OR REPLACE VIEW "exchange"."sharedmailboxmembers" AS 

SELECT
    data_element -> 'data' ->> 'id' AS SharedMailbox_Guid,
    jsonb_array_elements(
        CASE
            WHEN jsonb_typeof(data_element -> 'data' -> 'details') = 'array'
            THEN data_element -> 'data' -> 'details'
            ELSE '[]'::jsonb
        END
    ) ->> 'User' AS member
FROM importdata,
     jsonb_array_elements(data) AS data_element
WHERE
    name = 'shared-mailboxes/permissions.json'
--limit 100