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

