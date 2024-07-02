drop schema if exists "koksmat" cascade;

CREATE SCHEMA "koksmat"

{{ template "ensure_roles_and_permissions.sql".}} 

