DO $$
DECLARE
    r RECORD;
    column_exists BOOLEAN;
BEGIN
    -- Loop through each table in the public schema
    FOR r IN (SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE') 
    LOOP
        -- Check if koksmat_masterdataref column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_masterdataref'
        ) INTO column_exists;

        -- If koksmat_masterdataref does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_masterdata_ref VARCHAR COLLATE pg_catalog."default"';
        END IF;

        -- Check if koksmat_masterdata_id column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_masterdata_id'
        ) INTO column_exists;

        -- If koksmat_masterdata_id does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_masterdata_id VARCHAR COLLATE pg_catalog."default"';
        END IF;


        -- Check if koksmat_masterdata_etag column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_masterdata_etag'
        ) INTO column_exists;

        -- If koksmat_masterdata_etag does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_masterdata_etag VARCHAR COLLATE pg_catalog."default"';
        END IF;

        -- Check if koksmat__compliancetag column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_state'
        ) INTO column_exists;

        -- If koksmat_masterdata_etag does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_state VARCHAR COLLATE pg_catalog."default"';
        END IF;

        -- Check if koksmat__compliancetag column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat__compliancetag'
        ) INTO column_exists;

        -- If koksmat_masterdata_etag does not exist, add it
        IF  column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' DROP COLUMN koksmat__compliancetag';
        END IF;

        -- Check if koksmat_compliancetag column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_compliancetag'
        ) INTO column_exists;

        -- If koksmat_masterdata_etag does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_compliancetag VARCHAR COLLATE pg_catalog."default"';
        END IF;

        -- Check if koksmat_bucket column exists
        SELECT EXISTS (
            SELECT 1 
            FROM information_schema.columns 
            WHERE table_name = r.table_name 
              AND column_name = 'koksmat_bucket'
        ) INTO column_exists;

        -- If koksmat_bucket does not exist, add it
        IF NOT column_exists THEN
            EXECUTE 'ALTER TABLE public.' || quote_ident(r.table_name) || ' ADD COLUMN koksmat_bucket JSONB ';
        END IF;
    END LOOP;
END $$;
