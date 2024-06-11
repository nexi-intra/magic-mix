drop schema if exists "exchange";
-- cascade;

CREATE SCHEMA "exchange";

{{ template "sharedmailboxes.sql".}} 