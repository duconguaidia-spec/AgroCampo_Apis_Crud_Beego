
CREATE TABLE CategoriaPreguntas (
    id_categoria_preguntas SERIAL PRIMARY KEY,
    nombre_categoria character varying(80) NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TABLE PreguntaFrecuente (
    id_faq SERIAL PRIMARY KEY,
    pregunta character varying(255) NOT NULL,
    respuesta text,
    id_categoria_faq integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
);

