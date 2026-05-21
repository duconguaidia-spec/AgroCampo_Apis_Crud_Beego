-- Down migration: eliminar tablas creadas en la migración de creación
BEGIN;

-- Eliminar respuestas de foro primero (depende de temas)
DROP TABLE IF EXISTS public.respuestaforo CASCADE;

-- Eliminar temas de foro
DROP TABLE IF EXISTS public.temaforo CASCADE;

-- Eliminar videos educativos
DROP TABLE IF EXISTS public.videoeducativo CASCADE;

-- Finalmente eliminar categorías
DROP TABLE IF EXISTS public.categoria CASCADE;

COMMIT;

-- Fin de down migration
