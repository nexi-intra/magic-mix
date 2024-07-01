/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild

CREATE OR REPLACE PROCEDURE proc.create_job(
    p_actor_name VARCHAR,
    p_params JSONB,
    OUT p_id INTEGER
)
LANGUAGE plpgsql
AS $BODY$
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
    RETURNING id INTO p_id;

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

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$
;

/*
###MAGICAPP-START##
{
    "version": "v0.0.1",
    "action": "create",
    "input" : {
  "type": "object",
  "properties": {
  
    "tenant": { "type": "string" },
    "searchindex": { "type": "string" },
    "name": { "type": "string" },
    "description": { "type": "string" },
    "status": { "type": "string" },
    "startat": { "type": "string" },
    "startedAt": { "type": "string" },
    "completedAt": { "type": "string" },
    "maxduration": { "type": "number" },
    "script": { "type": "string" },
    "data": { "type": "object" }
}
    }

##MAGICAPP-END##
*/


