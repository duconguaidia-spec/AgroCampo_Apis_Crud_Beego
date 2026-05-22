CREATE SCHEMA IF NOT EXISTS Veterinarias;

CREATE TABLE IF NOT EXISTS Veterinarias.Especialidad (
    id_especialidad SERIAL PRIMARY KEY,
    nombre_especialidad VARCHAR(80) NOT NULL UNIQUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Veterinarias.ServicioGeneral (
    id_servicio_general SERIAL PRIMARY KEY,
    nombre_servicio VARCHAR(100) NOT NULL UNIQUE,
    descripcion VARCHAR(255),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS Veterinarias.Veterinaria (
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
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Veterinarias.tr_VeterinariaEspecialidad (
    id_vet_especialidad SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_especialidad INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_vet_especialidad_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES Veterinarias.Veterinaria(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_vet_especialidad_especialidad
        FOREIGN KEY (id_especialidad)
        REFERENCES Veterinarias.Especialidad(id_especialidad)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_veterinaria_especialidad UNIQUE (id_veterinaria, id_especialidad)
);

CREATE TABLE IF NOT EXISTS Veterinarias.tr_VeterinariaServicio (
    id_vet_servicio SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_servicio_general INT NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_vet_servicio_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES Veterinarias.Veterinaria(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_vet_servicio_general
        FOREIGN KEY (id_servicio_general)
        REFERENCES Veterinarias.ServicioGeneral(id_servicio_general)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_veterinaria_servicio_general UNIQUE (id_veterinaria, id_servicio_general)
);

CREATE TABLE IF NOT EXISTS Veterinarias.ServicioVeterinaria (
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
        REFERENCES Veterinarias.Veterinaria(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_servicio_precio
        CHECK (precio IS NULL OR precio >= 0)
);

CREATE TABLE IF NOT EXISTS Veterinarias.ProductoVeterinaria (
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
        REFERENCES Veterinarias.Veterinaria(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_producto_precio
        CHECK (precio IS NULL OR precio >= 0)
);

CREATE TABLE IF NOT EXISTS Veterinarias.ResenaVeterinaria (
    id_resena SERIAL PRIMARY KEY,
    id_veterinaria INT NOT NULL,
    id_usuario INT NOT NULL,
    calificacion INT NOT NULL,
    comentario VARCHAR(500),
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_resena_veterinaria
        FOREIGN KEY (id_veterinaria)
        REFERENCES Veterinarias.Veterinaria(id_veterinaria)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT ck_resena_calificacion
        CHECK (calificacion BETWEEN 1 AND 5)
);
