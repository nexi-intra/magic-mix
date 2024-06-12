drop schema if exists "people" cascade;

CREATE SCHEMA "people"

{{ template "sharedmailboxmemberships.sql".}} 

