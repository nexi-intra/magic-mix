CREATE OR REPLACE VIEW exchange.roomcalendarview AS
SELECT 
    to_timestamp("left"((eventitem -> 'start'::text) ->> 'dateTime'::text, 26), 'YYYY-MM-DD"T"HH24:MI:SS.US'::text) as start,
    to_timestamp("left"((eventitem -> 'end'::text) ->> 'dateTime'::text, 26), 'YYYY-MM-DD"T"HH24:MI:SS.US'::text) as end,
    -- eventitem -> 'start' ->> 'dateTime' AS start_time,
    -- eventitem -> 'end' ->> 'dateTime' AS end_time,
    parent ->> 'emailAddress' AS roomemail,
    eventitem -> 'organizer' -> 'emailAddress' ->> 'address' as organizer,
eventitem ->>  'recurrence' as recurrence

FROM
(
    SELECT 
        importdata.name AS importname,
        -- ((data_element.value -> 'data'::text) -> 'parentId'::text) -> 'emailAddress'::text AS roomemail,
        data_element.value -> 'data' -> 'parentId' AS parent,
        jsonb_array_elements(((data_element.value -> 'data'::text) -> 'details'::text)) AS eventitem
    FROM 
        importdata,
        LATERAL jsonb_array_elements(importdata.data) AS data_element(value)
    WHERE 
        importdata.name::text ILIKE 'roomscalendarview/%'::text
        AND (((data_element.value -> 'data'::text) -> 'details'::text) ) <> ' null'
) AS subquery;
