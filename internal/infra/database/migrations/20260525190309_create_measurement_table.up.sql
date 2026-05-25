CREATE TABLE IF NOT EXISTS public.measurements
(
    id              BIGSERIAL PRIMARY KEY,

    device_id       BIGINT NOT NULL REFERENCES public.devices(id),
    room_id         BIGINT REFERENCES public.rooms(id),

    value           DOUBLE PRECISION NOT NULL,

    created_date    TIMESTAMPTZ NOT NULL,
    updated_date    TIMESTAMPTZ NOT NULL,
    deleted_date    TIMESTAMPTZ
    );

CREATE INDEX IF NOT EXISTS measurements_device_id_index
    ON public.measurements (device_id);

CREATE INDEX IF NOT EXISTS measurements_room_id_index
    ON public.measurements (room_id);

CREATE INDEX IF NOT EXISTS measurements_deleted_date_index
    ON public.measurements (deleted_date);

CREATE INDEX IF NOT EXISTS measurements_created_date_index
    ON public.measurements (created_date);