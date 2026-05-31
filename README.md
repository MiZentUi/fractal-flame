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

![Result 1](images/result1.png)

![Result 2](images/result2.png)

![Result 3](images/result3.png)

![Result 4](images/result4.png)

![Result 5](images/result5.png)

![Result 6](images/result6.png)
