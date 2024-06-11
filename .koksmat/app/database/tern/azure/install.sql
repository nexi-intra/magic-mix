drop schema if exists "azure"
-- cascade;

CREATE SCHEMA "azure";

{{ template "group_members.sql".}} 
{{ template "group_owners.sql".}} 
{{ template "groups.sql".}} 
{{ template "users.sql".}} 
