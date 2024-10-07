
CREATE OR REPLACE FUNCTION koksmat_acl_init()
RETURNS void LANGUAGE plpgsql AS $$
DECLARE
    rec RECORD;  -- Declare the loop variable as a record
    dbname TEXT := current_database();  -- Get the current database name
    reader_role TEXT := 'magic_' || dbname || '_reader';
    writer_role TEXT := 'magic_' || dbname || '_writer';
BEGIN
    -- Notify starting process
    RAISE NOTICE 'Starting koksmat_acl_init procedure';

    -- Step 1: Create global roles if they do not exist and grant LOGIN privilege
    RAISE NOTICE 'Step 1: Checking and creating global roles if they do not exist, and granting LOGIN';

    PERFORM koksmat_create_role('koksmat_reader');
    PERFORM koksmat_create_role('koksmat_contributor');
    PERFORM koksmat_create_role('koksmat_owner');

    -- Step 2: Create database-specific roles if they do not exist
    RAISE NOTICE 'Step 2: Checking and creating database-specific roles magic_%_reader and magic_%_writer',dbname,dbname;

    PERFORM koksmat_create_role(reader_role);
    PERFORM koksmat_create_role(writer_role);

    -- Step 3: Grant SELECT permissions to magic_DATABASENAME_reader on user tables and views
    RAISE NOTICE 'Step 3: Granting SELECT permissions to % on user tables and views', reader_role;

    -- Only grant on user-defined tables, excluding system tables
    FOR rec IN 
        SELECT tablename FROM pg_tables 
        WHERE schemaname = 'public' AND tablename NOT LIKE 'pg_%' AND tablename NOT LIKE 'sql_%'
    LOOP
        -- Grant SELECT permissions on tables to magic_DATABASENAME_reader
        EXECUTE format('GRANT SELECT ON TABLE %I TO %I', rec.tablename, reader_role);
        RAISE NOTICE 'Granted SELECT on table % to %', rec.tablename, reader_role;
    END LOOP;

    -- Grant SELECT permissions on user-defined views to magic_DATABASENAME_reader
    FOR rec IN 
        SELECT viewname FROM pg_views 
        WHERE schemaname = 'public' AND viewname NOT LIKE 'pg_%' AND viewname NOT LIKE 'sql_%' AND viewname NOT LIKE 'hypopg_%'
    LOOP
        EXECUTE format('GRANT SELECT ON VIEW %I TO %I', rec.viewname, reader_role);
        RAISE NOTICE 'Granted SELECT on view % to %', rec.viewname, reader_role;
    END LOOP;

    -- Step 4: Grant EXECUTE and SELECT permissions to magic_DATABASENAME_writer on functions and procedures
    RAISE NOTICE 'Step 4: Granting EXECUTE and SELECT permissions to % on functions and procedures', writer_role;

    FOR rec IN 
        SELECT routine_name 
        FROM information_schema.routines 
        WHERE routine_schema = 'public' AND routine_type = 'FUNCTION' AND specific_name NOT LIKE 'pg_%' AND specific_name NOT LIKE 'sql_%' AND specific_name NOT LIKE 'hypopg_%' AND specific_name NOT LIKE 'koksmat_%'
    LOOP
        -- Grant EXECUTE permissions on functions and stored procedures to magic_DATABASENAME_writer
        EXECUTE format('GRANT EXECUTE ON FUNCTION %I() TO %I', rec.routine_name, writer_role);
        RAISE NOTICE 'Granted EXECUTE on function % to %', rec.routine_name, writer_role;
    END LOOP;

    -- Step 5: Ensure koksmat_owner cannot change access control
    RAISE NOTICE 'Step 5: Revoking access control modification privileges from koksmat_owner';
    
    PERFORM koksmat_revoke_acl_privileges('koksmat_owner');
    RAISE NOTICE 'Access control modification privileges revoked from koksmat_owner';

    -- Notify completion of process
    RAISE NOTICE 'Completed koksmat_acl_init procedure';
    
END;
$$;

-- Helper function to create role if it does not exist, and grant LOGIN
CREATE OR REPLACE FUNCTION koksmat_create_role(role_name TEXT)
RETURNS void LANGUAGE plpgsql AS $$
BEGIN
    -- Check if the role already exists
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = role_name) THEN
        -- Create the role if it does not exist
        EXECUTE format('CREATE ROLE %I', role_name);
        RAISE NOTICE 'Created role %', role_name;

        -- Grant LOGIN privilege to the role
        EXECUTE format('ALTER ROLE %I WITH LOGIN', role_name);
        RAISE NOTICE 'Granted LOGIN to role %', role_name;
    ELSE
        RAISE NOTICE 'Role % already exists', role_name;
    END IF;
END;
$$;

-- Helper function to revoke access control modification privileges
CREATE OR REPLACE FUNCTION koksmat_revoke_acl_privileges(role_name TEXT)
RETURNS void LANGUAGE plpgsql AS $$
BEGIN
    -- Ensure koksmat_owner does not have privileges to grant or revoke access control
    EXECUTE format('REVOKE GRANT OPTION FOR ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM %I', role_name);
    EXECUTE format('REVOKE GRANT OPTION FOR ALL PRIVILEGES ON SCHEMA public FROM %I', role_name);
    RAISE NOTICE 'Revoked GRANT OPTION FOR ALL PRIVILEGES on tables and schema from %', role_name;
END;
$$;

-- Helper function to generate a random password and assign it to a role
CREATE OR REPLACE FUNCTION koksmat_assign_password(role_name TEXT)
RETURNS TEXT LANGUAGE plpgsql AS $$
DECLARE
    role_password TEXT;
BEGIN
    -- Generate a random strong password (example: 32 characters long with letters, digits, special characters)
    role_password := (
        SELECT string_agg(
            substring('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+' FROM floor(random() * 72 + 1)::int FOR 1), ''
        )
        FROM generate_series(1, 32)
    );

    -- Assign the password to the role
    EXECUTE format('ALTER ROLE %I WITH PASSWORD %L', role_name, role_password);

    -- Notify that password was assigned
    RAISE NOTICE 'Assigned password to role %', role_name;

    -- Return the generated password
    RETURN role_password;
END;
$$;


