CREATE SCHEMA IF NOT EXISTS agropecuario;

CREATE SEQUENCE IF NOT EXISTS agropecuario."CategoriaGanado_id_categoria_ganado_seq";
CREATE SEQUENCE IF NOT EXISTS agropecuario."Subasta_id_subasta_seq";
CREATE SEQUENCE IF NOT EXISTS agropecuario."Tr_PrecioSubastaGanado_id_precio_subasta_seq";

CREATE TABLE IF NOT EXISTS agropecuario.categoria_ganado (
    id_categoria_ganado integer NOT NULL DEFAULT nextval('agropecuario."CategoriaGanado_id_categoria_ganado_seq"'::regclass),
    codigo character varying(10) NOT NULL,
    descripcion character varying(100) NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "CategoriaGanado_pkey" PRIMARY KEY (id_categoria_ganado),
    CONSTRAINT "CategoriaGanado_codigo_key" UNIQUE (codigo)
);

CREATE TABLE IF NOT EXISTS agropecuario.subasta (
    id_subasta integer NOT NULL DEFAULT nextval('agropecuario."Subasta_id_subasta_seq"'::regclass),
    fecha_subasta date NOT NULL,
    ubicacion character varying(150),
    precio numeric(10, 2) NOT NULL,
    id_usuario_registra integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "Subasta_pkey" PRIMARY KEY (id_subasta),
    CONSTRAINT fk_subasta_usuario FOREIGN KEY (id_usuario_registra)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS agropecuario."tr_precioSubastaGanado" (
    id_precio_subasta integer NOT NULL DEFAULT nextval('agropecuario."Tr_PrecioSubastaGanado_id_precio_subasta_seq"'::regclass),
    id_subasta integer NOT NULL,
    id_categoria_ganado integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "Tr_PrecioSubastaGanado_pkey" PRIMARY KEY (id_precio_subasta),
    CONSTRAINT uq_subasta_categoria UNIQUE (id_subasta, id_categoria_ganado),
    CONSTRAINT fk_precio_subasta_categoria FOREIGN KEY (id_categoria_ganado)
        REFERENCES agropecuario.categoria_ganado (id_categoria_ganado)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_precio_subasta_subasta FOREIGN KEY (id_subasta)
        REFERENCES agropecuario.subasta (id_subasta)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

ALTER SEQUENCE agropecuario."CategoriaGanado_id_categoria_ganado_seq" OWNED BY agropecuario.categoria_ganado.id_categoria_ganado;
ALTER SEQUENCE agropecuario."Subasta_id_subasta_seq" OWNED BY agropecuario.subasta.id_subasta;
ALTER SEQUENCE agropecuario."Tr_PrecioSubastaGanado_id_precio_subasta_seq" OWNED BY agropecuario."tr_precioSubastaGanado".id_precio_subasta;

CREATE INDEX IF NOT EXISTS idx_precio_subasta_id_categoria_ganado
    ON agropecuario."tr_precioSubastaGanado" (id_categoria_ganado);

CREATE INDEX IF NOT EXISTS idx_precio_subasta_id_subasta
    ON agropecuario."tr_precioSubastaGanado" (id_subasta);
