/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/   


-- tomat sild

CREATE OR REPLACE FUNCTION proc.create_user(
    p_actor_name VARCHAR,
    p_params JSONB
   
)
RETURNS JSONB LANGUAGE plpgsql 
AS $$
DECLARE
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_url VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_email VARCHAR;
    v_id INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_url := p_params->>'url';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_email := p_params->>'email';
         

    INSERT INTO public.user (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        url,
        name,
        description,
        email
    )
    VALUES (
        DEFAULT,
        DEFAULT,
        DEFAULT,
        p_actor_name, 
        p_actor_name,  -- Use the same value for updated_by
        v_tenant,
        v_url,
        v_name,
        v_description,
        v_email
    )
    RETURNING id INTO v_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_user',
        'status', 'success',
        'description', '',
        'action', 'create_user',
        'entity', 'user',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Create User",
  "description": "Create operation",

  "properties": {
  
    "tenant": { 
    "type": "string",
    "description":"" },
    "url": { 
    "type": "string",
    "description":"" },
    "name": { 
    "type": "string",
    "description":"" },
    "description": { 
    "type": "string",
    "description":"" },
    "email": { 
    "type": "string",
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




