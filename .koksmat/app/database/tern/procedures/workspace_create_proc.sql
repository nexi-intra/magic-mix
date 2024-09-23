/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild
-- TODO: Figure out why i had this in the public schmea and not in the proc schema 
CREATE OR REPLACE FUNCTION proc.create_workspace(
    p_actor_name VARCHAR,
    p_params JSONB,
    p_koksmat_sync JSONB DEFAULT NULL
   
)
RETURNS JSONB LANGUAGE plpgsql 
AS $$
DECLARE
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_user_id INTEGER;
    v_key VARCHAR;
    v_data JSONB;
    v_active BOOLEAN;
    v_id INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    RAISE NOTICE 'Actor % Input % ', p_actor_name,p_params;
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_user_id := p_params->>'user_id';
    v_key := p_params->>'key';
    v_data := p_params->>'data';
    v_active := p_params->>'active';
         
    
    INSERT INTO public.workspace (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        user_id,
        key,
        data,
        active
    )
    VALUES (
        DEFAULT,
        DEFAULT,
        DEFAULT,
        p_actor_name, 
        p_actor_name,  -- Use the same value for updated_by
        v_tenant,
        v_searchindex,
        v_name,
        v_description,
        v_user_id,
        v_key,
        v_data,
        v_active
    )
    RETURNING id INTO v_id;

    

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_workspace',
        'status', 'success',
        'description', '',
        'action', 'create_workspace',
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

    return jsonb_build_object(
    'comment','created',
    'id',v_id);

END;
$$ 
;




