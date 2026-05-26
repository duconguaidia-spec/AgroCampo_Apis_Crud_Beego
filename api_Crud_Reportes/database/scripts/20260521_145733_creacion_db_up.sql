CREATE TABLE ReporteActividad(
    id_reporte_actividad SERIAL PRIMARY KEY,
    tipo_reporte character varying(80)NOT NULL,
    fecha_inicio date NOT NULL,
    fecha_fin date NOT NULL,
    id_usuario_solicitante integer NOT NULL,
    resultado text,
    fecha_generacion timestamp without time zone NOT NULL DEFAULT now()
);
