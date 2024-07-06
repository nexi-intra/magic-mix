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
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "properties": {
    "title": "Update Job",
  "description": "Update operation",
  
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
END;
$BODY$
;


