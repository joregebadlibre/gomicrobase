# Usar la imagen oficial de PostgreSQL
FROM postgres:latest

# Establecer variables de entorno para PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Copiar archivos de inicialización (opcional)
COPY ./init.sql /docker-entrypoint-initdb.d/

# Exponer el puerto de PostgreSQL
EXPOSE 5432
