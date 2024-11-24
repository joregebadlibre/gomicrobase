CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS personas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    email VARCHAR(100)
);


CREATE TABLE IF NOT EXISTS cuentas (
    id SERIAL PRIMARY KEY,
    persona_id INT,
    saldo NUMERIC
);

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'users') THEN
        INSERT INTO users (name, email) VALUES ('Juan ', 'juan@.com');
    END IF;
END $$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'personas') THEN
        INSERT INTO public.personas (id, nombre, email) VALUES(nextval('personas_id_seq'::regclass), 'JOrge', 'jorge@l.com');
    END IF;
END $$;
