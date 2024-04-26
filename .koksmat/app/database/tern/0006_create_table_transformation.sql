/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   




CREATE TABLE public.transformation
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,input_id int  
    ,output_id int  


);

                ALTER TABLE IF EXISTS public.transformation
                ADD FOREIGN KEY (input_id)
                REFERENCES public.connection (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                CREATE TABLE public.transformation_m2m_mapper (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                deleted_at timestamp with time zone
                    ,transformation_id int  
 
                    ,mapper_id int  
 

                );
            

                ALTER TABLE public.transformation_m2m_mapper
                ADD FOREIGN KEY (transformation_id)
                REFERENCES public.transformation (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.transformation_m2m_mapper
                ADD FOREIGN KEY (mapper_id)
                REFERENCES public.mapper (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;                ALTER TABLE IF EXISTS public.transformation
                ADD FOREIGN KEY (output_id)
                REFERENCES public.dataset (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----
DROP TABLE IF EXISTS public.transformation_m2m_mapper;
DROP TABLE public.transformation;

