/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.accessholder
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,accesskey character varying COLLATE pg_catalog."default" 


);

                -- lollipop
                CREATE TABLE public.accessholder_m2m_role (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                created_by character varying COLLATE pg_catalog."default"  ,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_by character varying COLLATE pg_catalog."default",
                deleted_at timestamp with time zone
                
                    ,accessholder_id int  
 
                    ,role_id int  
 

                );
            

                ALTER TABLE public.accessholder_m2m_role
                ADD FOREIGN KEY (accessholder_id)
                REFERENCES public.accessholder (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.accessholder_m2m_role
                ADD FOREIGN KEY (role_id)
                REFERENCES public.role (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----
DROP TABLE IF EXISTS public.accessholder_m2m_role;
DROP TABLE public.accessholder;

