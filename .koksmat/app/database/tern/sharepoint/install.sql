drop schema if exists "sharepoint";
-- cascade;

CREATE SCHEMA "sharepoint";

{{ template "sharedmailboxes.sql".}} 