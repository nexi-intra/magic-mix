--drop schema if exists "koksmat" cascade;
CREATE SCHEMA if not exists "koksmat";

{{ template "model_create_proc.sql".}} 
{{ template "model_delete_proc.sql".}} 
{{ template "model_undo_delete_proc.sql".}} 
{{ template "model_update_proc.sql".}} 
{{ template "ensure_roles_and_permissions.sql".}} 
