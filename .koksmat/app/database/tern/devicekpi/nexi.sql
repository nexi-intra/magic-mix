-- View: devicekpi.nexi

-- DROP VIEW devicekpi.nexi;

CREATE OR REPLACE VIEW devicekpi.nexi
 AS
 SELECT 
   ( r.value ->> 'rownumber')::integer  AS row_number,

    r.value ->> ' (column 1)'::text AS column_1,
    r.value ->> 'All PC / XCORP + OFFICE (column 2)'::text AS column_2,
    r.value ->> ' (column 3)'::text AS column_3,
    r.value ->> ' (column 4)'::text AS column_4,
    r.value ->> ' (column 5)'::text AS column_5,
    r.value ->> ' (column 6)'::text AS column_6,
    r.value ->> ' (column 7)'::text AS column_7,
    r.value ->> ' (column 8)'::text AS column_8,
    r.value ->> ' (column 9)'::text AS column_9,
    r.value ->> ' (column 10)'::text AS column_10,
    r.value ->> ' (column 11)'::text AS column_11,
    r.value ->> ' (column 12)'::text AS column_12,
    r.value ->> ' (column 13)'::text AS column_13,
    r.value ->> ' (column 14)'::text AS column_14,
    r.value ->> ' (column 15)'::text AS column_15,
    r.value ->> ' (column 16)'::text AS column_16,
    r.value ->> ' (column 17)'::text AS column_17,
    r.value ->> ' (column 18)'::text AS column_18,
    r.value ->> ' (column 19)'::text AS column_19,
    r.value ->> ' (column 20)'::text AS column_20,
    r.value ->> ' (column 21)'::text AS column_21,
    r.value ->> ' (column 22)'::text AS column_22,
    r.value ->> ' (column 23)'::text AS column_23,
    r.value ->> ' (column 24)'::text AS column_24,
    r.value ->> ' (column 25)'::text AS column_25,
    r.value ->> ' (column 26)'::text AS column_26,
    r.value ->> ' (column 27)'::text AS column_27
   FROM importdata,
    LATERAL jsonb_array_elements(importdata.data) r(value)
  WHERE importdata.name::text ~~* 'devicekpi/table NEXI.json'::text
  order by row_number;

ALTER TABLE devicekpi.nexi
    OWNER TO pgadmin;

