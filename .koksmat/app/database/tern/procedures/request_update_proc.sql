/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sherry sild

CREATE OR REPLACE FUNCTION proc.update_request(
    p_actor_name VARCHAR,
    p_params JSONB,
    p_koksmat_sync JSONB DEFAULT NULL
   
)
RETURNS JSONB LANGUAGE plpgsql 
AS $$
DECLARE
    v_id INTEGER;
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_application VARCHAR;
    v_method VARCHAR;
    v_route VARCHAR;
    v_headers JSONB;
    v_body JSONB;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

    
BEGIN
    RAISE NOTICE 'Actor % Input % ', p_actor_name,p_params;
    v_id := p_params->>'id';
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_application := p_params->>'application';
    v_method := p_params->>'method';
    v_route := p_params->>'route';
    v_headers := p_params->>'headers';
    v_body := p_params->>'body';
         
    

        
    UPDATE public.request
    SET updated_by = p_actor_name,
        updated_at = CURRENT_TIMESTAMP,
        tenant = v_tenant,
        searchindex = v_searchindex,
        name = v_name,
        description = v_description,
        application = v_application,
        method = v_method,
        route = v_route,
        headers = v_headers,
        body = v_body
    WHERE id = v_id;

    GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
    IF v_rows_updated < 1 THEN
        RAISE EXCEPTION 'No records updated. request ID % not found', v_id ;
    END IF;


           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'update_request',
        'status', 'success',
        'description', '',
        'action', 'update_request',
        'entity', 'request',
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
    "title": "Update Request",
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
    "application": { 
    "type": "string",
    "description":"" },
    "method": { 
    "type": "string",
    "description":"" },
    "route": { 
    "type": "string",
    "description":"" },
    "headers": { 
    "type": "object",
    "description":"" },
    "body": { 
    "type": "object",
    "description":"" }

    }
}
##MAGICAPP-END##*/
    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);

    return jsonb_build_object(
    'comment','updated',
    'id',v_id
    );
END;
$$ 
;


