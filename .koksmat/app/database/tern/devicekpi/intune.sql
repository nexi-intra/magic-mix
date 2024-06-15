-- View: devicekpi.nexi

-- DROP VIEW devicekpi.intune;

CREATE OR REPLACE VIEW devicekpi.intune
 AS
 SELECT 
 ( r.value ->> 'rownumber')::integer  AS row_number,
 r.value ->> ' (column 1)'::text AS column_1,
--    r.value ->> 'INTUNE (column 2)'::text AS column_2,
    r.value ->> ' (column 3)'::text AS column_3,
    r.value ->> ' (column 4)'::text AS column_4,
    r.value ->> ' (column 5)'::text AS column_5,
    r.value ->> ' (column 6)'::text AS column_6
   
   FROM importdata,
    LATERAL jsonb_array_elements(importdata.data) r(value)
  WHERE importdata.name::text = 'devicekpi/intune.json'::text;


ALTER TABLE devicekpi.intune
    OWNER TO pgadmin;

