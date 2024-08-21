drop schema if exists "booking";
-- cascade;

-- Create the schema
CREATE SCHEMA IF NOT EXISTS booking;

{{ template "analyse.sql".}} 