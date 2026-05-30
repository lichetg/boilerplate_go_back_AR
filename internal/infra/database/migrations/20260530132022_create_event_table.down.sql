DROP INDEX IF EXISTS events_created_date_index;
DROP INDEX IF EXISTS events_deleted_date_index;
DROP INDEX IF EXISTS events_room_id_index;
DROP INDEX IF EXISTS events_device_id_index;

DROP TABLE IF EXISTS public.events;