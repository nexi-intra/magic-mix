/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild

CREATE OR REPLACE PROCEDURE proc.create_column(
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
    v_datatype VARCHAR;
    v_sortorder VARCHAR;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_datatype := p_params->>'datatype';
    v_sortorder := p_params->>'sortorder';
         

    INSERT INTO public.column (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        datatype,
        sortorder
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
        v_datatype,
        v_sortorder
    )
    RETURNING id INTO p_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_column',
        'status', 'success',
        'description', '',
        'action', 'create_column',
        'entity', 'column',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
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
    "datatype": { "type": "string" },
    "sortorder": { "type": "string" }
}
    }
}

##MAGICAPP-END##*/

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$
;




