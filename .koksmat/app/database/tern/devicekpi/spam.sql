
-- DROP VIEW devicekpi.spam ;

CREATE OR REPLACE VIEW devicekpi.spam
 AS
 SELECT 
 ( r.value ->> 'rownumber')::integer  AS row_number,
 

     r.value ->> 'SPAM (column 2)'::text AS column_2,
     r.value ->> ' (column 3)'::text AS column_3
     
     
   
   FROM importdata,
    LATERAL jsonb_array_elements(importdata.data) r(value)
  WHERE importdata.name::text = 'devicekpi/SPAM.json'::text;


ALTER TABLE devicekpi.spam
    OWNER TO pgadmin;

