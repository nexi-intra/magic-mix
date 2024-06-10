CREATE VIEW "project_lotus"."users" AS
SELECT * FROM "project_lotus".excelusers
 LEFT JOIN "azure"."users"
ON "extensionattribute11" = "nets_username";