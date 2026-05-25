CREATE TABLE public.categoria (
    id_categoria SERIAL PRIMARY KEY,
    nombre_categoria character varying(80) NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TABLE public.respuestaforo (
    id_respuesta SERIAL PRIMARY KEY,
    id_tema integer NOT NULL,
    id_autor integer NOT NULL,
    descripcion text NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TABLE public.temaforo (
    id_tema SERIAL PRIMARY KEY,
    titulo character varying(150) NOT NULL,
    descripcion text NOT NULL,
    id_autor integer NOT NULL,
    id_categoria integer NOT NULL,
    estado character varying(20) NOT NULL DEFAULT 'Pendiente',
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now()
);

CREATE TABLE public.videoeducativo (
    id_video SERIAL PRIMARY KEY,
    titulo character varying(150) NOT NULL,
    descripcion text NOT NULL,
    url_video character varying(255) NOT NULL,
    id_usuario integer NOT NULL,
    estado character varying(20) NOT NULL DEFAULT 'Borrador',
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now()
);

-- INSERTS de ejemplo
INSERT INTO public.categoria (nombre_categoria, activo) VALUES
    ('Agricultura', true),
    ('Ganadería', true),
    ('Tecnología', true);

INSERT INTO public.temaforo (titulo, descripcion, id_autor, id_categoria, estado) VALUES
    ('Mejoras en riego', 'Discusión sobre técnicas de riego eficientes', 1, 1, 'abierto'),
    ('Cuidado del ganado', 'Consejos para salud animal', 2, 2, 'abierto'),
    ('Herramientas digitales', 'Apps útiles para campo', 3, 3, 'abierto');

INSERT INTO public.respuestaforo (id_tema, id_autor, descripcion) VALUES
    (1, 2, 'Excelente aporte, gracias por compartir'),
    (1, 3, '¿Podrías compartir fuentes?'),
    (2, 1, 'Muy útil, aplicaré esto');

INSERT INTO public.videoeducativo (titulo, descripcion, url_video, id_usuario) VALUES
    ('Uso de drones', 'Introducción al uso de drones en agricultura', 'https://example.com/video1', 1),
    ('Manejo de suelos', 'Técnicas para mejorar la fertilidad', 'https://example.com/video2', 2);

-- Fin de inserts de ejemplo
