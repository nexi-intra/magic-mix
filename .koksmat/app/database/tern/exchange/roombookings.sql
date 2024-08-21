-- View: exchange.sharedmailboxmembers

-- DROP VIEW exchange.sharedmailboxmembers;

CREATE OR REPLACE VIEW exchange.roombookings AS
SELECT 

eventitem -> 'end' ->> 'dateTime' as enddatetime,
eventitem -> 'start' ->> 'dateTime' as startdatetime,
eventitem ->> 'type' as bookingtype,
eventitem -> 'organizer' -> 'emailAddress' ->> 'address' as organizer,
eventitem ->>  'recurrence' as recurrence,
eventitem ->>  'webLink' as webLink,
jsonb_array_length(eventitem -> 'attendees') as attendees

--*


FROM 
(
 SELECT 
 --*,
    importdata.name as importname
    ,data_element.value -> 'data'-> 'parentId' -> 'emailAddress'::text as roomemail
    ,jsonb_array_elements(data_element.value -> 'data'-> 'details' ->  'value') as eventitem
    
    -- data_element.value ->> 'RecipientTypeDetails'::text AS recipienttypedetails,
    -- data_element.value ->> 'PrimarySmtpAddress'::text AS primarysmtpaddress,
    -- data_element.value ->> 'Guid'::text AS guid
   FROM importdata
   
   ,LATERAL jsonb_array_elements(importdata.data) data_element(value)
  WHERE importdata.name ilike 'roomscalendarview/calendarview%'
    and data_element.value -> 'data'-> 'details' ->  'value'  <> '[]'
) as eventitems;

ALTER TABLE exchange.roombookings
     OWNER TO pgadmin;
