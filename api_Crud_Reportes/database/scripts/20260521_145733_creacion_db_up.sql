CREATE TABLE ReporteActividad(
    id_reporte_actividad SERIAL PRIMARY KEY,
    tipo_reporte character varying(80)NOT NULL,
    fecha_inicio date NOT NULL,
    fecha_fin date NOT NULL,
    id_usuario_solicitante integer NOT NULL,
    resultado text,
    fecha_generacion timestamp without time zone NOT NULL DEFAULT now()
);

INSERT INTO ReporteActividad (tipo_reporte, fecha_inicio, fecha_fin, id_usuario_solicitante) VALUES
    ('Reporte de actividad por usuario', '2026-01-01', '2026-01-31', 1);