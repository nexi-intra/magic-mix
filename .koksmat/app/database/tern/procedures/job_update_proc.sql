/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sherry sild

CREATE OR REPLACE PROCEDURE proc.update_job(
    p_actor_name VARCHAR,
    p_params JSONB
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_id INTEGER;
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
    v_id := p_params->>'id';
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
         
    
        
    UPDATE public.job
    SET updated_by = p_actor_name,
        updated_at = CURRENT_TIMESTAMP,
        tenant = v_tenant,
        searchindex = v_searchindex,
        name = v_name,
        description = v_description,
        status = v_status,
        startat = v_startat,
        startedAt = v_startedAt,
        completedAt = v_completedAt,
        maxduration = v_maxduration,
        script = v_script,
        data = v_data
    WHERE id = v_id;

    GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
    IF v_rows_updated < 1 THEN
        RAISE EXCEPTION 'No records updated. job ID % not found', v_id ;
    END IF;

           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'update_job',
        'status', 'success',
        'description', '',
        'action', 'update_job',
        'entity', 'job',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
    /*
###MAGICAPP-START##
{
    "version": "v0.0.1",
    "action": "update",
    "input" : {
  "type": "object",
  "properties": {
   "id": { "type": "number" },
  
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
END;
$BODY$
;


