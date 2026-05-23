FROM gradle:jdk21 AS build
WORKDIR /app

COPY build.gradle settings.gradle ./
RUN --mount=type=cache,target=/home/gradle/.gradle \
    gradle build -x test --no-daemon || true

COPY src ./src
RUN --mount=type=cache,target=/home/gradle/.gradle \
    gradle clean build -x test --no-daemon

FROM eclipse-temurin:21-jre-alpine
RUN apk --no-cache add curl
EXPOSE 9090
COPY --from=build /app/build/libs/*.jar app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]
