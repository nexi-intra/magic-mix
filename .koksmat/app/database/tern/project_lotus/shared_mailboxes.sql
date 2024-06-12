CREATE VIEW "project_lotus"."sharedmailbox_members" AS
select 

--sharedmailbox_guid,

exchange.sharedmailboxes.primarysmtpaddress,exchange.sharedmailboxes.displayname,project_lotus.users.mail,project_lotus.users.displayname
from "people"."sharedmailboxmemberships"
join "project_lotus"."users" on "project_lotus"."users"."id" = "people"."sharedmailboxmemberships"."userid"
join "exchange"."sharedmailboxes" 

on "exchange"."sharedmailboxes".guid = "people"."sharedmailboxmemberships".sharedmailbox_guid
where "project_lotus"."users"."mail" is not null
