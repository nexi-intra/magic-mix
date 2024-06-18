/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.job
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
    ,status character varying COLLATE pg_catalog."default"  NOT NULL
    ,startat character varying COLLATE pg_catalog."default"  
    ,startedAt character varying COLLATE pg_catalog."default"  
    ,completedAt character varying COLLATE pg_catalog."default"  
    ,maxduration character varying COLLATE pg_catalog."default"  
    ,script character varying COLLATE pg_catalog."default" 
    ,data JSONB  


);




---- create above / drop below ----

DROP TABLE public.job;

