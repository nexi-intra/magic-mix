--DROP VIEW "exchange"."sharedmailboxes";
CREATE OR REPLACE  VIEW "exchange"."sharedmailboxes" AS 

SELECT
    name,
    data_element ->> 'DisplayName' AS DisplayName,
    data_element ->> 'RecipientTypeDetails' AS RecipientTypeDetails,
    data_element ->> 'PrimarySmtpAddress' AS PrimarySmtpAddress,
    data_element ->> 'Guid' AS Guid
   

FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name = 'shared-mailboxes/parents.json';