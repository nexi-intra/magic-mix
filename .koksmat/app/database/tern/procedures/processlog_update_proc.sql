/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sherry sild

CREATE OR REPLACE PROCEDURE proc.update_processlog(
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
    v_transformation_id INTEGER;
    v_status VARCHAR;
    v_message VARCHAR;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

    
BEGIN
    v_id := p_params->>'id';
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_transformation_id := p_params->>'transformation_id';
    v_status := p_params->>'status';
    v_message := p_params->>'message';
         
    
        
    UPDATE public.processlog
    SET updated_by = p_actor_name,
        updated_at = CURRENT_TIMESTAMP,
        tenant = v_tenant,
        searchindex = v_searchindex,
        name = v_name,
        description = v_description,
        transformation_id = v_transformation_id,
        status = v_status,
        message = v_message
    WHERE id = v_id;

    GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
    IF v_rows_updated < 1 THEN
        RAISE EXCEPTION 'No records updated. processlog ID % not found', v_id ;
    END IF;

           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'update_processlog',
        'status', 'success',
        'description', '',
        'action', 'update_processlog',
        'entity', 'processlog',
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
    "transformation_id": { "type": "number" },
    "status": { "type": "string" },
    "message": { "type": "string" }
}
    }

##MAGICAPP-END##
*/
END;
$BODY$
;


