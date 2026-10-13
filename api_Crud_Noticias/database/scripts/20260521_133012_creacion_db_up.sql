CREATE SCHEMA IF NOT EXISTS Noticias;

CREATE TABLE Noticias.Noticias ( 
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
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);