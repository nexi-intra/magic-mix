drop schema if exists "devicekpi";
-- cascade;

CREATE SCHEMA "devicekpi";

{{ template "nexi.sql".}} 