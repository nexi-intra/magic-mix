CREATE VIEW "sharepoint"."allsites" AS
SELECT
    data_element  ->> 'Url' AS webUrl,
    data_element  ->> 'Title' AS Title,
    data_element  ->> 'GroupId' AS GroupId,
    data_element  ->> 'LocaleId' AS LocaleId,
    data_element  ->> 'HubSiteId' AS HusSiteId,
    data_element  ->> 'IsTeamsChannelConnected' AS IsTeamsChannelConnected,
    data_element  ->> 'IsHubsite' AS IsHubsite,
    data_element  ->> 'IsTeamsConnected' AS IsTeamsConnected,
    data_element  ->> 'SharingCapability' AS SharingCapability,
    data_element  ->> 'RelatedGroupId' AS RelatedGroupId,
    data_element  ->> 'LastContentModifiedDate' AS LastContentModifiedDate,
 
    name as batchname
FROM importdata, jsonb_array_elements(data) AS data_element
WHERE
    name = 'sharepoint/allsites.json'
