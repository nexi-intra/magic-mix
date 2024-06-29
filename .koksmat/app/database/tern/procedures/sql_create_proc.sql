/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild
/*
---
title: Create sql
description: Create a new sql
keep: false
sensivity: private
gdpr: high

---

# Create sql




`json input example
{
    "name": "John Doe",
    "description": "A person",
    "tenant": "default",
    "searchindex": "default"
}
`
`json output example
{
    "id": 1
}

`json zod schema
{
    "name": {
        "type": "string",
        "description": "Name of the entity"
    },
    "description": {
        "type": "string",
        "description": "Description of the entity"
    },
    "tenant": {
        "type": "string",
        "description": "Tenant of the entity"
    },
    "searchindex": {
        "type": "string",
        "description": "Search index of the entity"
    }
}
`

*/
CREATE OR REPLACE PROCEDURE proc.create_sql(
    p_actor_name VARCHAR,
    p_params JSONB,
    OUT p_id INTEGER
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_query VARCHAR;
    v_schema JSONB;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_query := p_params->>'query';
    v_schema := p_params->>'schema';
         

    INSERT INTO public.sql (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        query,
        schema
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
        v_query,
        v_schema
    )
    RETURNING id INTO p_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_sql',
        'status', 'success',
        'description', '',
        'action', 'create_sql',
        'entity', 'sql',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$
;
