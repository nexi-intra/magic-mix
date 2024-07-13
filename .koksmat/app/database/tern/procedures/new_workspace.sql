/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
-- tomat sild
CREATE
OR REPLACE PROCEDURE PROC.NEW_WORKSPACE (
	P_ACTOR_NAME VARCHAR,
	P_PARAMS JSONB,
	OUT P_ID INTEGER
) LANGUAGE PLPGSQL AS $BODY$
DECLARE
    v_rows_updated INTEGER;
    v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_key VARCHAR;
     v_data JSONB;
      v_active BOOLEAN;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_user_id INTEGER;

    v_upn VARCHAR COLLATE pg_catalog."default" ;
    v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;
    v_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_key := p_params->>'key';
    v_data := p_params->>'data';
    v_active := p_params->>'active';
    v_name := p_params->>'name';
    v_upn := p_params->>'upn';
    
    -- Lookup user in the user table
    SELECT id INTO v_user_id
    FROM public.user
    WHERE name = v_upn;

    -- If user does not exist, create a new user
    IF v_user_id IS NULL THEN
       -- Construct the JSONB object from variables
        v_params := jsonb_build_object('name', v_upn,'email', v_upn,'tenant','','url','');

        CALL proc.create_user(p_actor_name, v_params, v_user_id);
    END IF;


    v_params := jsonb_build_object('tenant', v_upn,
    'email', v_upn,
    'tenant',v_tenant,
    'searchindex',v_searchindex,
    'name',v_name,
    'description','',
    'user_id',v_user_id,
    'key',v_key,
    'data',v_data,
    'active',true);

    CALL proc.create_workspace(p_actor_name, v_params, p_id);

   
 

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'new_workspace',
        'status', 'success',
        'description', '',
        'action', 'new_workspace',
        'entity', 'workspace',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Create Workspace",
  "description": "Create operation",

  "properties": {
  
    "tenant": { 
    "type": "string",
    "description":"" },
    "searchindex": { 
    "type": "string",
    "description":"Search Index is used for concatenating all searchable fields in a single field making in easier to search\n" },
    "name": { 
    "type": "string",
    "description":"" },
    "description": { 
    "type": "string",
    "description":"" },
    "user_id": { 
    "type": "number",
    "description":"" },
    "key": { 
    "type": "string",
    "description":"" },
    "data": { 
    "type": "object",
    "description":"" },
    "active": { 
    "type": "boolean",
    "description":"" }

    }
}

##MAGICAPP-END##*/

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$;