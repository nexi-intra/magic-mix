DROP VIEW IF EXISTS "people"."sharedmailboxmemberships";
CREATE OR REPLACE VIEW "people"."sharedmailboxmemberships" AS 

select azure.users.id as userid, exchange.sharedmailboxmembers.SharedMailbox_Guid
from exchange.sharedmailboxmembers
    left join azure.users on azure.users.mail = exchange.sharedmailboxmembers.member
where
    azure.users.mail is not null