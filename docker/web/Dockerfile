
FROM node:alpine AS build

WORKDIR /app
ARG ENV
ARG PACKTOKEN
ENV PATH /app/node_modules/.bin:$PATH
COPY ./web/package.json ./
COPY ./web/package-lock.json ./
COPY ./web/.npmrc ./
RUN npm install -g dotenv-cli
RUN npm install
COPY ./web ./
RUN npm run $ENV

FROM nginx:stable-alpine
COPY --from=build /app/build  /usr/share/nginx/html
COPY ./deps/nginx.web/nginx.conf  /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

