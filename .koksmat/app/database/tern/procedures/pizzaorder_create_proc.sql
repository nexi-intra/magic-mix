/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
 */
-- tomat sild
-- TODO: Figure out why i had this in the public schmea and not in the proc schema
CREATE OR REPLACE FUNCTION proc.create_pizzaorder(p_actor_name varchar, p_params jsonb, p_koksmat_sync jsonb DEFAULT NULL)
    RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    v_rows_updated integer;
    v_tenant varchar COLLATE pg_catalog."default";
    v_searchindex varchar COLLATE pg_catalog."default";
    v_name varchar COLLATE pg_catalog."default";
    v_description varchar COLLATE pg_catalog."default";
    v_status varchar;
    v_size varchar;
    v_toppings varchar;
    v_id integer;
    v_audit_id integer;
    -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;
BEGIN
    RAISE NOTICE 'Actor % Input % ', p_actor_name, p_params;
    v_tenant := p_params ->> 'tenant';
    v_searchindex := p_params ->> 'searchindex';
    v_name := p_params ->> 'name';
    v_description := p_params ->> 'description';
    v_status := p_params ->> 'status';
    v_size := p_params ->> 'size';
    v_toppings := p_params ->> 'toppings';
    INSERT INTO public.pizzaorder(id, created_at, updated_at, created_by, updated_by, tenant, searchindex, name, description, status, size, toppings)
        VALUES (DEFAULT, DEFAULT, DEFAULT, p_actor_name, p_actor_name, -- Use the same value for updated_by
            v_tenant, v_searchindex, v_name, v_description, v_status, v_size, v_toppings)
    RETURNING
        id INTO v_id;
    p_auditlog_params := jsonb_build_object('tenant', '', 'searchindex', '', 'name', 'create_pizzaorder', 'status', 'success', 'description', '', 'action', 'create_pizzaorder', 'entity', 'pizzaorder', 'entityid', v_id::text, 'actor', p_actor_name, 'metadata', p_params);

    /*###MAGICAPP-START##
{
     "$schema": "https://json-schema.org/draft/2020-12/schema",
     "$id": "https://booking.services.koksmat.com/.schema.json",
     
     "type": "object",

     "title": "Create PizzaOrder",
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
     "size": { 
     "type": "string",
     "description":"" },
     "toppings": { 
     "type": "string",
     "description":"" }

     }
}

##MAGICAPP-END##*/
    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
    RETURN jsonb_build_object('comment', 'created', 'id', v_id);
END;
$$;

