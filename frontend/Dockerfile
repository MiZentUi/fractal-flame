FROM openapitools/openapi-generator-cli:latest AS generator

COPY openapi.yaml .
COPY . .

RUN /usr/local/bin/docker-entrypoint.sh generate \
    -i openapi.yaml \
     -g typescript-axios -o src/api/generated

FROM node:24-alpine as build
WORKDIR /app
ARG API_URL
ENV VITE_API_URL=$API_URL

COPY --from=generator . .

RUN npm i
RUN npm run build

FROM nginx
ENV PORT=80


COPY --from=build /app/dist /usr/share/nginx/html
COPY ./nginx/templates/ /etc/nginx/templates/
EXPOSE $PORT
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:${PORT} || exit 1


