/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   

-- karry sild

CREATE OR REPLACE PROCEDURE proc.undo_delete_sql(
    p_actor_name VARCHAR,
    p_params JSONB
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_id INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;


BEGIN
    v_id := p_params->>'id';
    
        
    UPDATE public.sql
    SET deleted_at = NULL,
        updated_at = CURRENT_TIMESTAMP,
        updated_by = p_actor_name
    WHERE id = v_id;

           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'undo_delete_sql',
        'status', 'success',
        'description', '',
        'action', 'undo_delete_sql',
        'entity', 'sql',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
    "version": "v0.0.1",
    "action": "undo_delete",
    "input" : {
    "type": "object",
    "properties": {
    "id": { "type": "number" }
}
    }
}
##MAGICAPP-END##*/
END;
$BODY$
;

