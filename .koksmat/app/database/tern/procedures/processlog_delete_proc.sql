/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- krydder sild

CREATE OR REPLACE PROCEDURE proc.delete_processlog(
    p_actor_name VARCHAR,
    p_params JSONB
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_id INTEGER;
    v_hard BOOLEAN;
     v_rows_updated INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;


BEGIN
    v_id := p_params->>'id';
    v_hard := p_params->>'hard';
  
    IF v_hard THEN
     DELETE FROM public.processlog
        WHERE id = v_id;
       
    ELSE
        UPDATE public.processlog
        SET deleted_at = CURRENT_TIMESTAMP,
            updated_at = CURRENT_TIMESTAMP,
            updated_by = p_actor_name
        WHERE id = v_id;
        
        GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
        IF v_rows_updated < 1 THEN
            RAISE EXCEPTION 'No records updated. processlog ID % not found', v_id ;
        END IF;
    END IF;
           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'delete_processlog',
        'status', 'success',
        'description', '',
        'action', 'delete_processlog',
        'entity', 'processlog',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
    "version": "v0.0.1",
    "action": "delete",
    "input" : {
  "type": "object",
  "properties": {
   "id": { "type": "number" },
    "hard": { "type": "boolean" }
}
    }
}
##MAGICAPP-END##*/
END;
$BODY$
;

