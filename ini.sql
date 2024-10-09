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
        INSERT INTO users (name, email) VALUES ('Juan Perez', 'juan@example.com');
    END IF;
END $$;
