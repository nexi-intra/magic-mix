drop schema if exists "sharepoint";
-- cascade;

CREATE SCHEMA "sharepoint";

{{ template "pageviews.sql".}} 
{{ template "sites.sql".}} 
{{ template "allsites.sql".}} 