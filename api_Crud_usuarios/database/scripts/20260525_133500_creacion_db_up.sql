CREATE SCHEMA IF NOT EXISTS usuarios;

CREATE SEQUENCE IF NOT EXISTS usuarios."Rol_id_rol_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."Usuario_id_usuario_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."Contrasena_id_contrasena_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."PerfilExtendido_id_perfil_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."TokenRecuperacion_id_token_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."VerificacionDosPasos_id_verificacion_seq";
CREATE SEQUENCE IF NOT EXISTS usuarios."AuditoriaUsuario_id_auditoria_seq";

CREATE TABLE IF NOT EXISTS usuarios."Rol" (
    id_rol integer NOT NULL DEFAULT nextval('usuarios."Rol_id_rol_seq"'::regclass),
    nombre_rol character varying(50) NOT NULL,
    descripcion character varying(150),
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "Rol_pkey" PRIMARY KEY (id_rol),
    CONSTRAINT "Rol_nombre_rol_key" UNIQUE (nombre_rol)
);

CREATE TABLE IF NOT EXISTS usuarios."Usuario" (
    id_usuario integer NOT NULL DEFAULT nextval('usuarios."Usuario_id_usuario_seq"'::regclass),
    nombre_completo character varying(120) NOT NULL,
    correo character varying(120) NOT NULL,
    telefono character varying(30),
    id_rol integer NOT NULL,
    verificacion_dos_pasos boolean NOT NULL DEFAULT false,
    avatar character varying(255),
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "Usuario_pkey" PRIMARY KEY (id_usuario),
    CONSTRAINT "Usuario_correo_key" UNIQUE (correo),
    CONSTRAINT fk_usuario_rol FOREIGN KEY (id_rol)
        REFERENCES usuarios."Rol" (id_rol)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS usuarios."Contrasena" (
    id_contrasena integer NOT NULL DEFAULT nextval('usuarios."Contrasena_id_contrasena_seq"'::regclass),
    id_usuario integer NOT NULL,
    contrasena_hash character varying(255) NOT NULL,
    activa boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "Contrasena_pkey" PRIMARY KEY (id_contrasena),
    CONSTRAINT fk_contrasena_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS usuarios."PerfilExtendido" (
    id_perfil integer NOT NULL DEFAULT nextval('usuarios."PerfilExtendido_id_perfil_seq"'::regclass),
    id_usuario integer NOT NULL,
    direccion character varying(150),
    ciudad character varying(80),
    intereses text,
    redes_sociales text,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "PerfilExtendido_pkey" PRIMARY KEY (id_perfil),
    CONSTRAINT "PerfilExtendido_id_usuario_key" UNIQUE (id_usuario),
    CONSTRAINT fk_perfil_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS usuarios."TokenRecuperacion" (
    id_token integer NOT NULL DEFAULT nextval('usuarios."TokenRecuperacion_id_token_seq"'::regclass),
    token character varying(255) NOT NULL,
    id_usuario integer NOT NULL,
    expiracion timestamp without time zone NOT NULL,
    usado boolean NOT NULL DEFAULT false,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "TokenRecuperacion_pkey" PRIMARY KEY (id_token),
    CONSTRAINT "TokenRecuperacion_token_key" UNIQUE (token),
    CONSTRAINT fk_token_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS usuarios."VerificacionDosPasos" (
    id_verificacion integer NOT NULL DEFAULT nextval('usuarios."VerificacionDosPasos_id_verificacion_seq"'::regclass),
    id_usuario integer NOT NULL,
    tipo_metodo character varying(30) NOT NULL,
    codigo_otp character varying(20),
    expiracion_otp timestamp without time zone,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "VerificacionDosPasos_pkey" PRIMARY KEY (id_verificacion),
    CONSTRAINT fk_verificacion_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS usuarios."AuditoriaUsuario" (
    id_auditoria integer NOT NULL DEFAULT nextval('usuarios."AuditoriaUsuario_id_auditoria_seq"'::regclass),
    id_usuario integer NOT NULL,
    tipo_evento character varying(60) NOT NULL,
    descripcion text,
    ip_origen character varying(45),
    fecha_evento timestamp without time zone NOT NULL,
    fecha_creacion timestamp without time zone NOT NULL DEFAULT now(),
    fecha_modificacion timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT "AuditoriaUsuario_pkey" PRIMARY KEY (id_auditoria),
    CONSTRAINT fk_auditoria_usuario FOREIGN KEY (id_usuario)
        REFERENCES usuarios."Usuario" (id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

ALTER SEQUENCE usuarios."Rol_id_rol_seq" OWNED BY usuarios."Rol".id_rol;
ALTER SEQUENCE usuarios."Usuario_id_usuario_seq" OWNED BY usuarios."Usuario".id_usuario;
ALTER SEQUENCE usuarios."Contrasena_id_contrasena_seq" OWNED BY usuarios."Contrasena".id_contrasena;
ALTER SEQUENCE usuarios."PerfilExtendido_id_perfil_seq" OWNED BY usuarios."PerfilExtendido".id_perfil;
ALTER SEQUENCE usuarios."TokenRecuperacion_id_token_seq" OWNED BY usuarios."TokenRecuperacion".id_token;
ALTER SEQUENCE usuarios."VerificacionDosPasos_id_verificacion_seq" OWNED BY usuarios."VerificacionDosPasos".id_verificacion;
ALTER SEQUENCE usuarios."AuditoriaUsuario_id_auditoria_seq" OWNED BY usuarios."AuditoriaUsuario".id_auditoria;

CREATE INDEX IF NOT EXISTS idx_usuario_id_rol
    ON usuarios."Usuario" (id_rol);

CREATE INDEX IF NOT EXISTS idx_contrasena_id_usuario
    ON usuarios."Contrasena" (id_usuario);

CREATE INDEX IF NOT EXISTS idx_token_id_usuario
    ON usuarios."TokenRecuperacion" (id_usuario);

CREATE INDEX IF NOT EXISTS idx_verificacion_id_usuario
    ON usuarios."VerificacionDosPasos" (id_usuario);

CREATE INDEX IF NOT EXISTS idx_auditoria_id_usuario
    ON usuarios."AuditoriaUsuario" (id_usuario);
