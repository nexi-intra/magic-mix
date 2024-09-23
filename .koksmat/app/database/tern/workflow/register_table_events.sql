CREATE OR REPLACE FUNCTION process_table_event(
    p_actor_name VARCHAR,
    p_table_name VARCHAR,
    p_table_recordid VARCHAR,
    p_table_operations VARCHAR,
    p_table_snapshot JSONB
)
RETURNS VOID AS $$
DECLARE
    found_record BOOLEAN;
BEGIN
    RAISE NOTICE 'process_table_event % table % ', p_actor_name, p_table_name;

    -- Check if a matching record exists in the koksmat table
    SELECT EXISTS (
        SELECT 1
        FROM koksmat
        WHERE data ->> 'table' = p_table_name
        AND data ->> 'type' = 'recipe'
    ) INTO found_record;

    -- If a matching record is found, insert the event
    IF found_record THEN
        RAISE NOTICE 'process_table_event % table found_record ', p_actor_name;

        INSERT INTO public.koksmat (
            id,
            created_at,
            updated_at,
            created_by, 
            updated_by, 
            tenant,
            searchindex,
            name,
            description,
            data
        )
        VALUES (  
            DEFAULT,
            DEFAULT,
            DEFAULT,
            p_actor_name,
            p_actor_name,
            '', -- tenant
            '', -- searchindex
            'event ' || p_table_name || ' ' || p_table_recordid || ' ' || p_table_operations, -- name
            '', -- description
            jsonb_build_object(
                'type', 'event',
                'operationtype', 'table',
                'operation', p_table_operations,
                'table_name', p_table_name,
                'table_recordid', p_table_recordid,
                'snapshot', p_table_snapshot
            )
        );
    END IF;
    RAISE NOTICE 'process_table_event % completing', p_actor_name;
    -- Explicitly signal the function's completion
    RETURN;
END;
$$ LANGUAGE plpgsql;