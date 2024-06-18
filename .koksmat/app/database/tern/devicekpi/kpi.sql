-- View: devicekpi.nexi

-- DROP VIEW devicekpi.kpi;

CREATE OR REPLACE VIEW devicekpi.kpi as 
select 'w5' as kpi,
       sum(nexi.column_13::integer) as Numerator,
       sum(nexi.column_14::integer) as Denominator
from devicekpi.nexi
where row_number = 12
    or row_number = 34
union all
select 'w6' as kpi,
       sum(nexi.column_13::integer) as Numerator,
       sum(nexi.column_14::integer) as Denominator
from devicekpi.nexi
where row_number = 13
    or row_number = 35
union all
select 'w8' as kpi,
       sum(nexi.column_5::integer) as Numerator,
       sum(nexi.column_6::integer) as Denominator
from devicekpi.nexi
where row_number = 3 
union all
select 'w4' as kpi,
       sum(nexi.column_13::integer) as Numerator,
       sum(nexi.column_14::integer) as Denominator
from devicekpi.nexi
where row_number = 13
    or row_number = 35

union all
select 'w3' as kpi,
       sum(nexi.column_5::integer) as Numerator,
       sum(nexi.column_6::integer) as Denominator
from devicekpi.nexi
where row_number = 24
    or row_number = 37
union all
select 'w7a' as kpi,
       sum(regexp_replace(nexi.column_5, '\D', '', 'g')::integer) as Numerator,
       0 as Denominator
from devicekpi.nexi
where row_number = 16
  
   union all
select 'w7b' as kpi,
       sum(regexp_replace(nexi.column_5, '\D', '', 'g')::integer) as Numerator,
       0 as Denominator
from devicekpi.nexi
where row_number = 17
    or row_number = 39

union all
select 'w2' as kpi,
       sum(intune.column_3::integer) as Numerator,
       sum(intune.column_5::integer) as Denominator
from devicekpi.intune
where row_number = 5
   


union all
select 'cs11' as kpi,
        
       spam.column_3::bigint as Numerator,
      0 as Denominator
from devicekpi.spam
where row_number = 4
    
order by kpi;



ALTER TABLE devicekpi.kpi
    OWNER TO pgadmin;
