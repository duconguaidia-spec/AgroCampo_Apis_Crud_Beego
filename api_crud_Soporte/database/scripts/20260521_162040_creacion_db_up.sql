
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

INSERT INTO CategoriaPreguntas (nombre_categoria, activo) VALUES
    ('General', true),
    ('Técnica', true),
    ('Administrativa', true);

INSERT INTO PreguntaFrecuente (pregunta, respuesta, id_categoria_faq) VALUES
    ('¿Cómo puedo restablecer mi contraseña?', 'Para restablecer tu contraseña, haz clic en "Olvidé mi contraseña" en la página de inicio de sesión y sigue las instrucciones.', 1),
    ('¿Cómo puedo contactar al soporte técnico?', 'Puedes contactar al soporte técnico enviando un correo a soporte@empresa.com', 2);
