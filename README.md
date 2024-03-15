## Ecommerce

# Create database using Docker

    docker run --name ecommerce -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123secret -d postgres:16-alpine