BEGIN;

CREATE SCHEMA IF NOT EXISTS "Usuarios";
CREATE SCHEMA IF NOT EXISTS "Agropecuario";
CREATE SCHEMA IF NOT EXISTS "Veterinarias";
CREATE SCHEMA IF NOT EXISTS "Contenido";
CREATE SCHEMA IF NOT EXISTS "Noticias";
CREATE SCHEMA IF NOT EXISTS "Reportes";
CREATE SCHEMA IF NOT EXISTS "Soporte";


-- Usuarios
CREATE TABLE IF NOT EXISTS "Usuarios"."Rol" (
    id_rol SERIAL PRIMARY KEY,
    nombre_rol VARCHAR(50) NOT NULL,
    descripcion VARCHAR(200),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "Usuarios"."Usuario" (
    id_usuario SERIAL PRIMARY KEY,
    nombre_completo VARCHAR(100) NOT NULL,
    correo VARCHAR(100) NOT NULL UNIQUE,
    telefono VARCHAR(12),
    id_rol INT NOT NULL,
    verificacion_dos_pasos BOOLEAN NOT NULL DEFAULT FALSE,
    avatar VARCHAR(255),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_usuario_rol
        FOREIGN KEY (id_rol)
        REFERENCES "Usuarios"."Rol"(id_rol)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

-- Agropecuario
CREATE TABLE IF NOT EXISTS "Agropecuario"."CategoriaGanado" (
    id_categoria_ganado SERIAL PRIMARY KEY,
    codigo VARCHAR(10) NOT NULL UNIQUE,
    descripcion VARCHAR(100) NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Veterinarias
CREATE TABLE IF NOT EXISTS "Veterinarias"."Especialidad" (
    id_especialidad SERIAL PRIMARY KEY,
    nombre_especialidad VARCHAR(80) NOT NULL UNIQUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."ServicioGeneral" (
    id_servicio_general SERIAL PRIMARY KEY,
    nombre_servicio VARCHAR(100) NOT NULL UNIQUE,
    descripcion VARCHAR(255),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Contenido
CREATE TABLE IF NOT EXISTS "Contenido"."Categoria" (
    id_categoria SERIAL PRIMARY KEY,
    nombre_categoria VARCHAR(80) NOT NULL UNIQUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Soporte
CREATE TABLE IF NOT EXISTS "Soporte"."CategoriaPreguntas" (
    id_categoria_preguntas SERIAL PRIMARY KEY,
    nombre_categoria VARCHAR(80) NOT NULL UNIQUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Usuarios 
CREATE TABLE IF NOT EXISTS "Usuarios"."Contrasena" (
    id_contrasena SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL,
    contrasena_hash VARCHAR(255) NOT NULL,
    activa BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_contrasena_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Usuarios"."TokenRecuperacion" (
    id_token SERIAL PRIMARY KEY,
    token VARCHAR(10) NOT NULL UNIQUE,
    id_usuario INT NOT NULL,
    expiracion TIMESTAMP NOT NULL,
    usado BOOLEAN NOT NULL DEFAULT FALSE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tokenrecuperacion_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Usuarios"."VerificacionDosPasos" (
    id_verificacion SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL,
    tipo_metodo VARCHAR(20) NOT NULL,
    codigo_otp VARCHAR(10),
    expiracion_otp TIMESTAMP,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_verificacion_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Usuarios"."AuditoriaUsuario" (
    id_auditoria SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL,
    tipo_evento VARCHAR(80) NOT NULL,
    descripcion VARCHAR(255),
    ip_origen VARCHAR(45),
    fecha_evento TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_auditoria_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Usuarios"."PerfilExtendido" (
    id_perfil SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL UNIQUE,
    direccion VARCHAR(150),
    ciudad VARCHAR(80),
    intereses VARCHAR(255),
    redes_sociales VARCHAR(255),
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_perfil_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

-- Agropecuario 
CREATE TABLE IF NOT EXISTS "Agropecuario"."Subasta" (
    id_subasta SERIAL PRIMARY KEY,
    fecha_subasta DATE NOT NULL,
    ubicacion VARCHAR(150),
    precio NUMERIC(10,2) NOT NULL,
    id_usuario_registra INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_subasta_usuario
        FOREIGN KEY (id_usuario_registra)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS "Agropecuario"."Tr_PrecioSubastaGanado" (
    id_precio_subasta SERIAL PRIMARY KEY,
    id_subasta INT NOT NULL,
    id_categoria_ganado INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_precio_subasta_subasta
        FOREIGN KEY (id_subasta)
        REFERENCES "Agropecuario"."Subasta"(id_subasta)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_precio_subasta_categoria
        FOREIGN KEY (id_categoria_ganado)
        REFERENCES "Agropecuario"."CategoriaGanado"(id_categoria_ganado)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_subasta_categoria UNIQUE (id_subasta, id_categoria_ganado)
);

-- Veterinarias 
CREATE TABLE IF NOT EXISTS "Veterinarias"."Veterinaria" (
    id_veterinaria SERIAL PRIMARY KEY,
    nombre_clinica VARCHAR(100) NOT NULL,
    direccion VARCHAR(150) NOT NULL,
    telefono VARCHAR(12) NOT NULL,
    correo_publico VARCHAR(100),
    horario_atencion VARCHAR(150),
    calificacion_promedio NUMERIC(3,2) NOT NULL DEFAULT 0,
    id_usuario INT NOT NULL,
    aprobado BOOLEAN NOT NULL DEFAULT FALSE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_veterinaria_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_veterinaria_calificacion
        CHECK (calificacion_promedio >= 0 AND calificacion_promedio <= 5)
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."tr_VeterinariaEspecialidad" (
    id_vet_especialidad SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_especialidad INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_vet_especialidad_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES "Veterinarias"."Veterinaria"(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_vet_especialidad_especialidad
        FOREIGN KEY (id_especialidad)
        REFERENCES "Veterinarias"."Especialidad"(id_especialidad)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_veterinaria_especialidad UNIQUE (id_veterinaria, id_especialidad)
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."tr_VeterinariaServicio" (
    id_vet_servicio SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_servicio_general INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_vet_servicio_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES "Veterinarias"."Veterinaria"(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_vet_servicio_general
        FOREIGN KEY (id_servicio_general)
        REFERENCES "Veterinarias"."ServicioGeneral"(id_servicio_general)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_veterinaria_servicio_general UNIQUE (id_veterinaria, id_servicio_general)
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."ServicioVeterinaria" (
    id_servicio SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    nombre_servicio VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255),
    precio NUMERIC(10,2),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_servicio_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES "Veterinarias"."Veterinaria"(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_servicio_precio
        CHECK (precio IS NULL OR precio >= 0)
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."ProductoVeterinaria" (
    id_producto_vet SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    nombre_producto VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255),
    precio NUMERIC(10,2),
    disponible BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_producto_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES "Veterinarias"."Veterinaria"(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_producto_precio
        CHECK (precio IS NULL OR precio >= 0)
);

CREATE TABLE IF NOT EXISTS "Veterinarias"."ResenaVeterinaria" (
    id_resena SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_usuario INT NOT NULL,
    calificacion INT NOT NULL,
    comentario VARCHAR(500),
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_resena_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES "Veterinarias"."Veterinaria"(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_resena_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_resena_calificacion
        CHECK (calificacion BETWEEN 1 AND 5)
);

-- Contenido 
CREATE TABLE IF NOT EXISTS "Contenido"."VideoEducativo" (
    id_video SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    descripcion TEXT NOT NULL,
    url_video VARCHAR(255) NOT NULL,
    id_usuario INT NOT NULL,
    estado VARCHAR(20) NOT NULL DEFAULT 'Borrador',
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_video_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_video_estado
        CHECK (estado IN ('Borrador', 'Publicado', 'Archivado'))
);

CREATE TABLE IF NOT EXISTS "Contenido"."TemaForo" (
    id_tema SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    descripcion TEXT NOT NULL,
    id_autor INT NOT NULL,
    id_categoria INT NOT NULL,
    estado VARCHAR(20) NOT NULL DEFAULT 'Pendiente',
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_tema_autor
        FOREIGN KEY (id_autor)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_tema_categoria
        FOREIGN KEY (id_categoria)
        REFERENCES "Contenido"."Categoria"(id_categoria)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_tema_estado
        CHECK (estado IN ('Pendiente', 'Aprobado', 'Rechazado', 'Cerrado'))
);

CREATE TABLE IF NOT EXISTS "Contenido"."RespuestaForo" (
    id_respuesta SERIAL PRIMARY KEY,
    id_tema INT NOT NULL,
    id_autor INT NOT NULL,
    descripcion TEXT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_respuesta_tema
        FOREIGN KEY (id_tema)
        REFERENCES "Contenido"."TemaForo"(id_tema)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_respuesta_autor
        FOREIGN KEY (id_autor)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

-- Noticias 
CREATE TABLE IF NOT EXISTS "Noticias"."Noticia" (
    id_noticia SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    tipo VARCHAR(20) NOT NULL,
    cuerpo TEXT,
    url_video VARCHAR(255),
    imagen_destacada VARCHAR(255) NOT NULL,
    fuente VARCHAR(150),
    fecha_noticia DATE,
    id_usuario INT NOT NULL,
    acceso_limitado BOOLEAN NOT NULL DEFAULT FALSE,
    estado VARCHAR(20) NOT NULL DEFAULT 'Borrador',
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_noticia_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_noticia_estado
        CHECK (estado IN ('Borrador', 'Publicado', 'Archivado')),
    CONSTRAINT ck_noticia_tipo
        CHECK (tipo IN ('Articulo', 'Video', 'Boletin', 'Comunicado'))
);

-- Reportes 
CREATE TABLE IF NOT EXISTS "Reportes"."ReporteActividad" (
    id_reporte_actividad SERIAL PRIMARY KEY,
    tipo_reporte VARCHAR(80) NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    id_usuario_solicitante INT NOT NULL,
    resultado TEXT,
    fecha_generacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_reporte_usuario
        FOREIGN KEY (id_usuario_solicitante)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_reporte_fechas
        CHECK (fecha_fin >= fecha_inicio)
);

CREATE TABLE IF NOT EXISTS "Reportes"."tr_ExportacionDatos" (
    id_tr_exportacion_datos SERIAL PRIMARY KEY,
    id_usuario INT NOT NULL,
    id_reporte INT NOT NULL,
    formato VARCHAR(10) NOT NULL,
    url_archivo VARCHAR(255),
    estado VARCHAR(20) NOT NULL DEFAULT 'Procesando',
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_exportacion_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES "Usuarios"."Usuario"(id_usuario)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_exportacion_reporte
        FOREIGN KEY (id_reporte)
        REFERENCES "Reportes"."ReporteActividad"(id_reporte_actividad)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_exportacion_formato
        CHECK (formato IN ('PDF', 'CSV', 'XLSX', 'JSON')),
    CONSTRAINT ck_exportacion_estado
        CHECK (estado IN ('Procesando', 'Generado', 'Fallido'))
);

-- Soporte 
CREATE TABLE IF NOT EXISTS "Soporte"."PreguntaFrecuente" (
    id_faq SERIAL PRIMARY KEY,
    pregunta VARCHAR(255) NOT NULL,
    respuesta TEXT NOT NULL,
    id_categoria_faq INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_faq_categoria
        FOREIGN KEY (id_categoria_faq)
        REFERENCES "Soporte"."CategoriaPreguntas"(id_categoria_preguntas)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_usuario_id_rol
    ON "Usuarios"."Usuario"(id_rol);

CREATE INDEX IF NOT EXISTS idx_contrasena_id_usuario
    ON "Usuarios"."Contrasena"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_tokenrecuperacion_id_usuario
    ON "Usuarios"."TokenRecuperacion"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_verificacion_id_usuario
    ON "Usuarios"."VerificacionDosPasos"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_auditoria_id_usuario
    ON "Usuarios"."AuditoriaUsuario"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_subasta_id_usuario_registra
    ON "Agropecuario"."Subasta"(id_usuario_registra);

CREATE INDEX IF NOT EXISTS idx_precio_subasta_id_subasta
    ON "Agropecuario"."Tr_PrecioSubastaGanado"(id_subasta);

CREATE INDEX IF NOT EXISTS idx_precio_subasta_id_categoria_ganado
    ON "Agropecuario"."Tr_PrecioSubastaGanado"(id_categoria_ganado);

CREATE INDEX IF NOT EXISTS idx_veterinaria_id_usuario
    ON "Veterinarias"."Veterinaria"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_vet_especialidad_id_veterinaria
    ON "Veterinarias"."tr_VeterinariaEspecialidad"(id_veterinaria);

CREATE INDEX IF NOT EXISTS idx_vet_especialidad_id_especialidad
    ON "Veterinarias"."tr_VeterinariaEspecialidad"(id_especialidad);

CREATE INDEX IF NOT EXISTS idx_vet_servicio_id_veterinaria
    ON "Veterinarias"."tr_VeterinariaServicio"(id_veterinaria);

CREATE INDEX IF NOT EXISTS idx_vet_servicio_id_servicio_general
    ON "Veterinarias"."tr_VeterinariaServicio"(id_servicio_general);

CREATE INDEX IF NOT EXISTS idx_servicio_veterinaria_id_veterinaria
    ON "Veterinarias"."ServicioVeterinaria"(id_veterinaria);

CREATE INDEX IF NOT EXISTS idx_producto_veterinaria_id_veterinaria
    ON "Veterinarias"."ProductoVeterinaria"(id_veterinaria);

CREATE INDEX IF NOT EXISTS idx_resena_veterinaria_id_veterinaria
    ON "Veterinarias"."ResenaVeterinaria"(id_veterinaria);

CREATE INDEX IF NOT EXISTS idx_resena_veterinaria_id_usuario
    ON "Veterinarias"."ResenaVeterinaria"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_videoeducativo_id_usuario
    ON "Contenido"."VideoEducativo"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_temaforo_id_autor
    ON "Contenido"."TemaForo"(id_autor);

CREATE INDEX IF NOT EXISTS idx_temaforo_id_categoria
    ON "Contenido"."TemaForo"(id_categoria);

CREATE INDEX IF NOT EXISTS idx_respuestaforo_id_tema
    ON "Contenido"."RespuestaForo"(id_tema);

CREATE INDEX IF NOT EXISTS idx_respuestaforo_id_autor
    ON "Contenido"."RespuestaForo"(id_autor);

CREATE INDEX IF NOT EXISTS idx_noticia_id_usuario
    ON "Noticias"."Noticia"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_reporteactividad_id_usuario_solicitante
    ON "Reportes"."ReporteActividad"(id_usuario_solicitante);

CREATE INDEX IF NOT EXISTS idx_exportaciondatos_id_usuario
    ON "Reportes"."tr_ExportacionDatos"(id_usuario);

CREATE INDEX IF NOT EXISTS idx_exportaciondatos_id_reporte
    ON "Reportes"."tr_ExportacionDatos"(id_reporte);

CREATE INDEX IF NOT EXISTS idx_preguntafrecuente_id_categoria_faq
    ON "Soporte"."PreguntaFrecuente"(id_categoria_faq);

COMMIT;
BEGIN;

-- Usuarios.Rol
INSERT INTO "Usuarios"."Rol" (nombre_rol, descripcion) VALUES 
('Administrador', 'Control total del sistema'),
('Ganadero', 'Usuario propietario de ganado y fincas'),
('Veterinario', 'Profesional de salud animal registrado');

-- Agropecuario.CategoriaGanado
INSERT INTO "Agropecuario"."CategoriaGanado" (codigo, descripcion) VALUES 
('BOV-LEC', 'Ganado bovino de leche'),
('BOV-CAR', 'Ganado bovino de carne'),
('POR-ENG', 'Ganado porcino de engorde');

-- Veterinarias.Especialidad
INSERT INTO "Veterinarias"."Especialidad" (nombre_especialidad) VALUES
('Medicina bovina'),
('Cirugía de grandes animales'),
('Nutrición y alimentación animal');

-- Veterinarias.ServicioGeneral
INSERT INTO "Veterinarias"."ServicioGeneral" (nombre_servicio, descripcion) VALUES
('Consulta general', 'Atención veterinaria general para animales de producción'),
('Vacunación', 'Aplicación de vacunas y esquemas preventivos'),
('Nutrición animal', 'Asesoría en alimentación, suplementación y balance nutricional');

-- Contenido.Categoria
INSERT INTO "Contenido"."Categoria" (nombre_categoria) VALUES 
('Sanidad animal'),
('Manejo de pasturas'),
('Mercados y precios ganaderos');

-- Soporte.CategoriaPreguntas
INSERT INTO "Soporte"."CategoriaPreguntas" (nombre_categoria) VALUES
('Registro y acceso'),
('Subastas y precios'),
('Veterinarias y servicios');

-- Usuarios.Usuario
INSERT INTO "Usuarios"."Usuario" (nombre_completo, correo, telefono, id_rol) VALUES 
('Carlos Medina', 'carlos.medina@agrocampo.co', '3101234567', 1),
('Juan Pablo Ríos', 'juanpablo.rios@gmail.com', '3209876543', 2),
('María Fernanda López', 'mfernanda.lopez@vetcol.co', '3154567890', 3);

-- Usuarios.Contrasena
INSERT INTO "Usuarios"."Contrasena" (id_usuario, contrasena_hash) VALUES
(1, 'KIXabcHashAdmin'),
(2, 'KIXabcHashGanadero'),
(3, 'KIXabcHashVeterinario');

-- Usuarios.TokenRecuperacion
INSERT INTO "Usuarios"."TokenRecuperacion" (token, id_usuario, expiracion) VALUES 
('TK1A2B3C4D', 1, NOW() + INTERVAL '1 hour'),
('TK5E6F7G8H', 2, NOW() + INTERVAL '1 hour'),
('TK9I0J1K2L', 3, NOW() + INTERVAL '1 hour');

-- Usuarios.VerificacionDosPasos
INSERT INTO "Usuarios"."VerificacionDosPasos" (id_usuario, tipo_metodo, codigo_otp, expiracion_otp) VALUES 
(1, 'EMAIL', '482910', NOW() + INTERVAL '10 minutes'),
(2, 'SMS', '739204', NOW() + INTERVAL '10 minutes'),
(3, 'EMAIL', '615847', NOW() + INTERVAL '10 minutes');

-- Usuarios.AuditoriaUsuario
INSERT INTO "Usuarios"."AuditoriaUsuario" (id_usuario, tipo_evento, descripcion, ip_origen) VALUES
(1, 'LOGIN', 'Inicio de sesión exitoso', '192.168.1.10'),
(2, 'CAMBIO_CONTRASENA', 'El usuario actualizó su contraseña', '190.24.50.112'),
(3, 'LOGOUT', 'Cierre de sesión manual', '181.33.22.98');

-- Usuarios.PerfilExtendido
INSERT INTO "Usuarios"."PerfilExtendido" (id_usuario, direccion, ciudad, intereses, redes_sociales) VALUES 
(1, 'Calle 12 # 5-40', 'Villavicencio', 'Ganadería, Tecnología', 'https://linkedin.com/in/carlosmedina'),
(2, 'Vereda El Porvenir Km 3', 'Puerto López', 'Agricultura, Subastas', 'https://facebook.com/juanpablo.rios'),
(3, 'Carrera 8 # 20-15', 'Granada', 'Medicina veterinaria, Bovinos', 'https://instagram.com/mfernanda.vet');

-- Agropecuario.Subasta
INSERT INTO "Agropecuario"."Subasta" (fecha_subasta, ubicacion, precio, id_usuario_registra) VALUES 
('2025-03-10', 'Feria ganadera de Villavicencio, Meta', 2500000.00, 1),
('2025-03-17', 'Centro agropecuario Puerto López', 1800000.00, 2),
('2025-03-24', 'Subasta ganadera Granada, Meta', 3100000.00, 1);

-- Agropecuario.Tr_PrecioSubastaGanado
INSERT INTO "Agropecuario"."Tr_PrecioSubastaGanado" (id_subasta, id_categoria_ganado) VALUES 
(1, 1),
(2, 2),
(3, 3);

-- Veterinarias.Veterinaria
INSERT INTO "Veterinarias"."Veterinaria" (
    nombre_clinica, direccion, telefono, correo_publico, horario_atencion, id_usuario
) VALUES 
('Clínica VetCampo', 'Cra 14 # 8-22, Villavicencio', '3157891234', 'vetcampo@clinica.co', 'Lun-Sab 7am-6pm', 3),
('AgroVet Llanos', 'Calle 5 # 2-10, Granada', '3209874561', 'agrovet@llanos.co', 'Lun-Vie 8am-5pm', 3),
('Centro Veterinario del Meta', 'Av. Catama # 15-40, Villavicencio', '3104562378', 'cvmeta@vet.co', 'Lun-Dom 6am-8pm', 3);

-- Veterinarias.tr_VeterinariaEspecialidad
INSERT INTO "Veterinarias"."tr_VeterinariaEspecialidad" (id_veterinaria, id_especialidad) VALUES 
(1, 1),
(2, 2),
(3, 3);

-- Veterinarias.ServicioVeterinaria
INSERT INTO "Veterinarias"."ServicioVeterinaria" (id_veterinaria, nombre_servicio, descripcion, precio) VALUES 
(1, 'Vacunación bovina', 'Aplicación de vacunas contra aftosa y brucelosis', 45000.00),
(2, 'Cirugía de casco', 'Corrección de lesiones podales en bovinos', 320000.00),
(3, 'Plan nutricional', 'Diseño de dieta para engorde de ganado', 150000.00);

-- Veterinarias.tr_VeterinariaServicio
INSERT INTO "Veterinarias"."tr_VeterinariaServicio" (id_veterinaria, id_servicio_general) VALUES
(1, 1),
(2, 2),
(3, 3);

-- Veterinarias.ProductoVeterinaria
INSERT INTO "Veterinarias"."ProductoVeterinaria" (id_veterinaria, nombre_producto, descripcion, precio) VALUES
(1, 'Vacuna Aftosa Bivalente', 'Dosis de 2ml para bovinos mayores de 3 meses', 18000.00),
(2, 'Ivermectina 1%', 'Antiparasitario inyectable para bovinos y porcinos', 25000.00),
(3, 'Suplemento mineral bovino', 'Bolsa 25kg, mezcla de sales minerales para pastoreo', 87000.00);

-- Veterinarias.ResenaVeterinaria
INSERT INTO "Veterinarias"."ResenaVeterinaria" (id_veterinaria, id_usuario, calificacion, comentario) VALUES 
(1, 2, 5, 'Excelente atención, el veterinario conoce muy bien el manejo bovino.'),
(2, 1, 4, 'Buen servicio, aunque los tiempos de espera son un poco largos.'),
(3, 2, 5, 'El plan nutricional mejoró notablemente el peso de mi ganado.');

-- Contenido.VideoEducativo
INSERT INTO "Contenido"."VideoEducativo" (titulo, descripcion, url_video, id_usuario, estado) VALUES
('Cómo aplicar la vacuna contra la aftosa', 'Guía paso a paso para la correcta vacunación bovina en campo.', 'https://cdn.agrocampo.co/videos/vacuna-aftosa.mp4', 3, 'Publicado'),
('Manejo de praderas en época de verano', 'Técnicas para conservar la productividad de pasturas durante la sequía.', 'https://cdn.agrocampo.co/videos/praderas-verano.mp4', 1, 'Publicado'),
('Cómo leer los precios de la subasta ganadera', 'Explicación de indicadores y categorías en las ferias ganaderas del Meta.', 'https://cdn.agrocampo.co/videos/precios-subasta.mp4', 2, 'Borrador');

-- Contenido.TemaForo
INSERT INTO "Contenido"."TemaForo" (titulo, descripcion, id_autor, id_categoria, estado) VALUES
('¿Cuál es la mejor raza bovina para el Llano?', 'Quiero saber qué razas se adaptan mejor al clima del Meta para producción de carne.', 2, 1, 'Aprobado'),
('Precios de subasta marzo 2025 en Villavicencio', 'Comparto los precios que vi esta semana en la feria ganadera de Villavicencio.', 1, 3, 'Aprobado'),
('Problema con pasto estrella en invierno', 'El pasto se está encharcando, ¿qué me recomiendan para manejarlo?', 2, 2, 'Pendiente');

-- Contenido.RespuestaForo
INSERT INTO "Contenido"."RespuestaForo" (id_tema, id_autor, descripcion) VALUES 
(1, 3, 'Para el Meta recomiendo la raza Brahman o cruces con Senepol, son muy adaptables al calor y a los pastos nativos.'),
(2, 2, 'Gracias por la info, el novillo gordo estuvo entre $2.8M y $3.1M en mi finca.'),
(3, 3, 'Lo ideal es mejorar el drenaje de los potreros y rotar con pasto Brachiaria brizantha.');


INSERT INTO "Noticias"."Noticia" (
    titulo, tipo, cuerpo, url_video, imagen_destacada, fuente, fecha_noticia, id_usuario, estado
) VALUES 
(
    'Precio del kilo de carne bovina sube un 8% en el Meta',
    'Articulo',
    'El mercado ganadero del departamento del Meta registró un incremento del 8% en el precio del kilo en pie durante el primer trimestre del año, según datos de la feria ganadera de Villavicencio.',
    NULL,
    'https://cdn.agrocampo.co/img/noticia-precio-carne.jpg',
    'Fedegán',
    '2025-03-20',
    1,
    'Publicado'
),
(
    'ICA refuerza campaña de vacunación contra la fiebre aftosa',
    'Articulo',
    'El Instituto Colombiano Agropecuario inició el segundo ciclo de vacunación del año en los departamentos de los Llanos Orientales con meta de cubrir más de 2 millones de bovinos.',
    NULL,
    'https://cdn.agrocampo.co/img/noticia-ica-vacuna.jpg',
    'ICA Colombia',
    '2025-03-15',
    1,
    'Publicado'
),
(
    'Feria Ganadera Internacional de Villavicencio 2025',
    'Video',
    'Cobertura audiovisual de la feria ganadera internacional realizada en Villavicencio durante 2025.',
    'https://cdn.agrocampo.co/videos/feria-ganadera-2025.mp4',
    'https://cdn.agrocampo.co/img/noticia-feria-2025.jpg',
    'Cámara de Comercio Meta',
    '2025-03-01',
    1,
    'Publicado'
);

-- Reportes.ReporteActividad
INSERT INTO "Reportes"."ReporteActividad" (
    tipo_reporte, fecha_inicio, fecha_fin, id_usuario_solicitante, resultado
) VALUES 
('Actividad de usuarios', '2025-03-01', '2025-03-31', 1, 'Total inicios de sesión: 412. Usuarios activos: 38. Nuevos registros: 9.'),
('Precios de subasta', '2025-01-01', '2025-03-31', 1, 'Promedio novillo gordo: $2.950.000. Categoría con mayor volumen: BOV-CAR.'),
('Publicaciones de contenido', '2025-03-01', '2025-03-31', 1, 'Videos publicados: 5. Temas de foro: 12. Noticias: 8.');

-- Reportes.tr_ExportacionDatos
INSERT INTO "Reportes"."tr_ExportacionDatos" (
    id_usuario, id_reporte, formato, url_archivo, estado
) VALUES 
(1, 1, 'PDF', 'https://cdn.agrocampo.co/reportes/actividad-usuarios-mar25.pdf', 'Generado'),
(1, 2, 'XLSX', 'https://cdn.agrocampo.co/reportes/precios-subasta-q1-25.xlsx', 'Generado'),
(1, 3, 'PDF', NULL, 'Procesando');

-- Soporte.PreguntaFrecuente
INSERT INTO "Soporte"."PreguntaFrecuente" (pregunta, respuesta, id_categoria_faq) VALUES 
('¿Cómo recupero mi contraseña?', 'Ve a la pantalla de inicio de sesión, haz clic en "¿Olvidaste tu contraseña?" e ingresa tu correo registrado. Recibirás un código de recuperación válido por 1 hora.', 1),
('¿Cada cuánto se actualizan los precios de subasta?', 'Los precios se actualizan después de cada feria ganadera registrada en el sistema, generalmente de forma semanal según el calendario de subastas del Meta.', 2),
('¿Cómo puedo registrar mi clínica veterinaria?', 'Crea una cuenta con el rol Veterinario, dirígete a "Mi clínica" en el menú y completa el formulario de registro. Un administrador revisará y aprobará tu solicitud.', 3);
