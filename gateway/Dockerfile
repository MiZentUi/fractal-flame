FROM gradle:jdk21 AS build
WORKDIR /app

COPY build.gradle settings.gradle ./
COPY lombok.config ./
COPY openapi.yaml .openapi-generator-ignore ./
COPY src ./src
RUN gradle clean build -x test --no-daemon

FROM eclipse-temurin:21-jre-alpine
RUN apk --no-cache add curl
EXPOSE 8080
COPY --from=build /app/build/libs/*.jar app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]
