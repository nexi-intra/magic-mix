CREATE OR REPLACE VIEW booking.utilization AS

WITH json_data AS (
    SELECT
        jsonb_array_elements(data) AS record
    FROM
        importdata
    WHERE
        importdata.name ILIKE 'booking/booking_timeslots_output.json%'
)
SELECT
    record->>'resource_id' AS resource_id,
    record->>'booking_id' AS booking_id,
    record->>'start_date' AS start_date,
    record->>'end_date' AS end_date,
    timeslot->>'date' AS date,
    (slot->>'hour')::int AS hour,
    (slot->>'minute')::int AS minute
FROM
    json_data,
    LATERAL (
        SELECT jsonb_array_elements(record->'timeslots') AS timeslot
        WHERE jsonb_typeof(record->'timeslots') = 'array'
    ) AS timeslots_lateral,
    LATERAL (
        SELECT jsonb_array_elements(timeslot->'slots') AS slot
        WHERE jsonb_typeof(timeslot->'slots') = 'array'
    ) AS slots_lateral;
