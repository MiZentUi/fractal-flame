# Fractal Flame

## Usage

### dotenv file

```shell
POSTGRES_USER=postgres
POSTGRES_PASSWORD=admin

MINIO_ROOT_USER=minio
MINIO_ROOT_PASSWORD=minioadmin
JWT_SIGNING_KEY=BASE64
```

### Run

```shell
docker-compose up -d --build
```

### Port mapping

- 8080 - gateway
- 9090 - generator
- 5432 - generator postgres
- 9000 - generator minio
- 9001 - generator minio web

## Gallery

![Result 1](docs/images/result1.png)

![Result 2](docs/images/result2.png)

![Result 3](docs/images/result3.png)

![Result 4](docs/images/result4.png)

![Result 5](docs/images/result5.png)

![Result 6](docs/images/result6.png)

![Fractal Flame 1](docs/images/ff1.png)

![Fractal Flame 2](docs/images/ff2.png)

![Fractal Flame 3](docs/images/ff3.png)

![Fractal Flame 4](docs/images/ff4.png)

![Fractal Flame 5](docs/images/ff5.png)
