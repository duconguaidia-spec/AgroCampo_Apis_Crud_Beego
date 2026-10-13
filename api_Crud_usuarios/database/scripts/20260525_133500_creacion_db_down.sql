DROP TABLE IF EXISTS usuarios."AuditoriaUsuario";
DROP TABLE IF EXISTS usuarios."VerificacionDosPasos";
DROP TABLE IF EXISTS usuarios."TokenRecuperacion";
DROP TABLE IF EXISTS usuarios."PerfilExtendido";
DROP TABLE IF EXISTS usuarios."Contrasena";
DROP TABLE IF EXISTS usuarios."Usuario";
DROP TABLE IF EXISTS usuarios."Rol";

DROP SEQUENCE IF EXISTS usuarios."AuditoriaUsuario_id_auditoria_seq";
DROP SEQUENCE IF EXISTS usuarios."VerificacionDosPasos_id_verificacion_seq";
DROP SEQUENCE IF EXISTS usuarios."TokenRecuperacion_id_token_seq";
DROP SEQUENCE IF EXISTS usuarios."PerfilExtendido_id_perfil_seq";
DROP SEQUENCE IF EXISTS usuarios."Contrasena_id_contrasena_seq";
DROP SEQUENCE IF EXISTS usuarios."Usuario_id_usuario_seq";
DROP SEQUENCE IF EXISTS usuarios."Rol_id_rol_seq";

DROP SCHEMA IF EXISTS usuarios;
