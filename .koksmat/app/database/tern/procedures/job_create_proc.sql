/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild

CREATE OR REPLACE FUNCTION proc.create_job(
    p_actor_name VARCHAR,
    p_params JSONB
   
)
RETURNS JSONB LANGUAGE plpgsql 
AS $$
DECLARE
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_status VARCHAR;
    v_startat TIMESTAMP WITH TIME ZONE;
    v_startedAt TIMESTAMP WITH TIME ZONE;
    v_completedAt TIMESTAMP WITH TIME ZONE;
    v_maxduration INTEGER;
    v_script VARCHAR;
    v_data JSONB;
    v_id INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_status := p_params->>'status';
    v_startat := p_params->>'startat';
    v_startedAt := p_params->>'startedAt';
    v_completedAt := p_params->>'completedAt';
    v_maxduration := p_params->>'maxduration';
    v_script := p_params->>'script';
    v_data := p_params->>'data';
         

    INSERT INTO public.job (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        status,
        startat,
        startedAt,
        completedAt,
        maxduration,
        script,
        data
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
        v_status,
        v_startat,
        v_startedAt,
        v_completedAt,
        v_maxduration,
        v_script,
        v_data
    )
    RETURNING id INTO v_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_job',
        'status', 'success',
        'description', '',
        'action', 'create_job',
        'entity', 'job',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Create Job",
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
    "status": { 
    "type": "string",
    "description":"" },
    "startat": { 
    "type": "string",
    "description":"" },
    "startedAt": { 
    "type": "string",
    "description":"" },
    "completedAt": { 
    "type": "string",
    "description":"" },
    "maxduration": { 
    "type": "number",
    "description":"" },
    "script": { 
    "type": "string",
    "description":"" },
    "data": { 
    "type": "object",
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




