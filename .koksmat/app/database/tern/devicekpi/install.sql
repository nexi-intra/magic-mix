drop schema if exists "devicekpi";
-- cascade;

CREATE SCHEMA "devicekpi";

{{ template "nexi.sql".}} 
{{ template "intune.sql".}} 
{{ template "spam.sql".}} 
{{ template "kpi.sql".}} 