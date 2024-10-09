CREATE TABLE personas (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    email VARCHAR(100)
);

CREATE TABLE cuentas (
    id SERIAL PRIMARY KEY,
    persona_id INT REFERENCES personas(id),
    saldo NUMERIC
);
