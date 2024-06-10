drop schema if exists "project_lotus" cascade;

CREATE SCHEMA "project_lotus"

{{ template "excelusers.sql".}} 

{{ template "users.sql".}} 

{{ template "group_members.sql".}} 