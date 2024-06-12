CREATE OR REPLACE VIEW project_lotus.group_members AS
SELECT DISTINCT
    g.id,
    g.displayname,
    g.mail
FROM azure.groups g
    JOIN azure.group_members m ON g.id = m.group_id
    JOIN azure.users u ON m.user_id = u.id
    JOIN project_lotus.users pu ON u.id = pu.id
WHERE
    grouptypes ILIKE '%Unified%'
    AND creationOptions ILIKE '%Team%'