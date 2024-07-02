CREATE OR REPLACE PROCEDURE koksmat.ensure_roles_and_permissions()
LANGUAGE plpgsql
AS $$
DECLARE
    schema_record RECORD;
    table_record RECORD;
    view_record RECORD;
    routine_record RECORD;
BEGIN
    -- Ensure the roles exist
    RAISE NOTICE 'Ensuring roles exist';
    DO
    $do$
    BEGIN
        -- Create roles if they do not exist
        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'koksmat_admin') THEN
            CREATE ROLE koksmat_admin;
            RAISE NOTICE 'Created role koksmat_admin';
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'koksmat_editor') THEN
            CREATE ROLE koksmat_editor;
            RAISE NOTICE 'Created role koksmat_editor';
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'koksmat_reader') THEN
            CREATE ROLE koksmat_reader;
            RAISE NOTICE 'Created role koksmat_reader';
        END IF;
    END
    $do$;

    -- Loop through all schemas
    RAISE NOTICE 'Looping through schemas';
    FOR schema_record IN
        SELECT schema_name
        FROM information_schema.schemata
        WHERE schema_name NOT IN ('pg_catalog', 'information_schema')
    LOOP
        RAISE NOTICE 'Processing schema: %', schema_record.schema_name;
        -- Loop through all tables in the current schema
        FOR table_record IN
            SELECT table_name
            FROM information_schema.tables
            WHERE table_schema = schema_record.schema_name
        LOOP
            RAISE NOTICE 'Processing table: %.%', schema_record.schema_name, table_record.table_name;
            -- Grant all privileges to admin
            EXECUTE format('GRANT ALL PRIVILEGES ON TABLE %I.%I TO koksmat_admin', schema_record.schema_name, table_record.table_name);
            RAISE NOTICE 'Granted ALL PRIVILEGES on table %.% to koksmat_admin', schema_record.schema_name, table_record.table_name;

            -- Grant select on table to editor
            EXECUTE format('GRANT SELECT ON TABLE %I.%I TO koksmat_editor', schema_record.schema_name, table_record.table_name);
            RAISE NOTICE 'Granted SELECT on table %.% to koksmat_editor', schema_record.schema_name, table_record.table_name;
            
            -- Grant select only to reader
            EXECUTE format('GRANT SELECT ON TABLE %I.%I TO koksmat_reader', schema_record.schema_name, table_record.table_name);
            RAISE NOTICE 'Granted SELECT on table %.% to koksmat_reader', schema_record.schema_name, table_record.table_name;
        END LOOP;

        -- Loop through all views in the current schema
        FOR view_record IN
            SELECT table_name
            FROM information_schema.views
            WHERE table_schema = schema_record.schema_name
        LOOP
            RAISE NOTICE 'Processing view: %.%', schema_record.schema_name, view_record.table_name;
        --     -- Grant all privileges to admin
        --     -- EXECUTE format('GRANT ALL PRIVILEGES ON VIEW %I.%I TO koksmat_admin', schema_record.schema_name, view_record.table_name);
        --     RAISE NOTICE 'Granted ALL PRIVILEGES on view %.% to koksmat_admin', schema_record.schema_name, view_record.table_name;

        --     -- Grant select on view to editor
        --     -- EXECUTE format('GRANT SELECT ON VIEW %I.%I TO koksmat_editor', schema_record.schema_name, view_record.table_name);
        --     RAISE NOTICE 'Granted SELECT on view %.% to koksmat_editor', schema_record.schema_name, view_record.table_name;
            
        --     -- Grant select only to reader
        --     -- EXECUTE format('GRANT SELECT ON VIEW %I.%I TO koksmat_reader', schema_record.schema_name, view_record.table_name);
        --     RAISE NOTICE 'Granted SELECT on view %.% to koksmat_reader', schema_record.schema_name, view_record.table_name;
         END LOOP;

        -- Loop through all functions in the current schema
        FOR routine_record IN
            SELECT routine_name, specific_name
            FROM information_schema.routines
            WHERE routine_schema = schema_record.schema_name
              AND routine_type = 'FUNCTION'
        LOOP
            RAISE NOTICE 'Processing function: %.%', schema_record.schema_name, routine_record.routine_name;
            -- Grant execute privilege on all functions to admin
            EXECUTE format('GRANT EXECUTE ON FUNCTION %I.%I TO koksmat_admin', schema_record.schema_name, routine_record.routine_name);
            RAISE NOTICE 'Granted EXECUTE on function %.% to koksmat_admin', schema_record.schema_name, routine_record.routine_name;

            -- Grant execute privilege on all functions to editor
            EXECUTE format('GRANT EXECUTE ON FUNCTION %I.%I TO koksmat_editor', schema_record.schema_name, routine_record.routine_name);
            RAISE NOTICE 'Granted EXECUTE on function %.% to koksmat_editor', schema_record.schema_name, routine_record.routine_name;
        END LOOP;

        -- Loop through all procedures in the current schema
        FOR routine_record IN
            SELECT routine_name, specific_name
            FROM information_schema.routines
            WHERE routine_schema = schema_record.schema_name
              AND routine_type = 'PROCEDURE'
        LOOP
            RAISE NOTICE 'Processing procedure: %.%', schema_record.schema_name, routine_record.routine_name;
            -- Grant execute privilege on all procedures to admin
            EXECUTE format('GRANT EXECUTE ON PROCEDURE %I.%I TO koksmat_admin', schema_record.schema_name, routine_record.routine_name);
            RAISE NOTICE 'Granted EXECUTE on procedure %.% to koksmat_admin', schema_record.schema_name, routine_record.routine_name;

            -- Grant execute privilege on all procedures to editor
            EXECUTE format('GRANT EXECUTE ON PROCEDURE %I.%I TO koksmat_editor', schema_record.schema_name, routine_record.routine_name);
            RAISE NOTICE 'Granted EXECUTE on procedure %.% to koksmat_editor', schema_record.schema_name, routine_record.routine_name;
        END LOOP;
    END LOOP;
END;
$$;

-- Call the procedure to ensure roles and permissions
-- CALL koksmat.ensure_roles_and_permissions();
