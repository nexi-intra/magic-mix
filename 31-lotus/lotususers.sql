CREATE SCHEMA IF NOT EXISTS "lotususers";
DROP VIEW IF EXISTS "lotususers"."users";
CREATE VIEW "lotususers"."users" AS
SELECT * FROM lotususers.excel 
 LEFT JOIN "azure"."users"
ON "extensionattribute11" = "nets_username"
