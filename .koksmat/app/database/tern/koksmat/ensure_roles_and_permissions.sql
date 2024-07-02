CREATE OR REPLACE PROCEDURE koksmat.ensure_roles_and_permissions()
LANGUAGE plpgsql
AS $$
BEGIN
    -- Ensure the roles exist
    DO
    $do$
    BEGIN
        -- Create roles if they do not exist
        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
            CREATE ROLE admin;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'editor') THEN
            CREATE ROLE editor;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'reader') THEN
            CREATE ROLE reader;
        END IF;
    END
    $do$;

    -- Loop through all schemas
    FOR schema_record IN
        SELECT schema_name
        FROM information_schema.schemata
        WHERE schema_name NOT IN ('pg_catalog', 'information_schema')
    LOOP
        -- Loop through all tables in the current schema
        FOR table_record IN
            SELECT table_name
            FROM information_schema.tables
            WHERE table_schema = schema_record.schema_name
        LOOP
            -- Grant all privileges to admin
            EXECUTE format('GRANT ALL PRIVILEGES ON TABLE %I.%I TO admin', schema_record.schema_name, table_record.table_name);

            -- Grant select, execute on update functions to editor
            EXECUTE format('GRANT SELECT ON TABLE %I.%I TO editor', schema_record.schema_name, table_record.table_name);
            
            -- Grant select only to reader
            EXECUTE format('GRANT SELECT ON TABLE %I.%I TO reader', schema_record.schema_name, table_record.table_name);
        END LOOP;

        -- Loop through all views in the current schema
        FOR view_record IN
            SELECT table_name
            FROM information_schema.views
            WHERE table_schema = schema_record.schema_name
        LOOP
            -- Grant all privileges to admin
            EXECUTE format('GRANT ALL PRIVILEGES ON VIEW %I.%I TO admin', schema_record.schema_name, view_record.table_name);

            -- Grant select, execute on update functions to editor
            EXECUTE format('GRANT SELECT ON VIEW %I.%I TO editor', schema_record.schema_name, view_record.table_name);
            
            -- Grant select only to reader
            EXECUTE format('GRANT SELECT ON VIEW %I.%I TO reader', schema_record.schema_name, view_record.table_name);
        END LOOP;

        -- Loop through all functions in the current schema
        FOR function_record IN
            SELECT routine_name
            FROM information_schema.routines
            WHERE routine_schema = schema_record.schema_name
              AND routine_type = 'FUNCTION'
        LOOP
            -- Grant execute privilege on all functions to admin
            EXECUTE format('GRANT EXECUTE ON FUNCTION %I.%I TO admin', schema_record.schema_name, function_record.routine_name);

            -- Grant execute privilege on all functions to editor
            EXECUTE format('GRANT EXECUTE ON FUNCTION %I.%I TO editor', schema_record.schema_name, function_record.routine_name);
        END LOOP;

        -- Loop through all procedures in the current schema
        FOR procedure_record IN
            SELECT routine_name
            FROM information_schema.routines
            WHERE routine_schema = schema_record.schema_name
              AND routine_type = 'PROCEDURE'
        LOOP
            -- Grant execute privilege on all procedures to admin
            EXECUTE format('GRANT EXECUTE ON PROCEDURE %I.%I TO admin', schema_record.schema_name, procedure_record.routine_name);

            -- Grant execute privilege on all procedures to editor
            EXECUTE format('GRANT EXECUTE ON PROCEDURE %I.%I TO editor', schema_record.schema_name, procedure_record.routine_name);
        END LOOP;
    END LOOP;
END;
$$;

-- Call the procedure to ensure roles and permissions
--CALL ensure_roles_and_permissions();
