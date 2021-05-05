
FROM python:buster AS build

RUN \
  echo "deb https://deb.nodesource.com/node_14.x buster main" > /etc/apt/sources.list.d/nodesource.list && \
  wget -qO- https://deb.nodesource.com/gpgkey/nodesource.gpg.key | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" > /etc/apt/sources.list.d/yarn.list && \
  wget -qO- https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  apt-get update && \
  apt-get install -yqq nodejs yarn && \
  pip install -U pip && pip install pipenv && \
  npm i -g npm@^6

WORKDIR /app
ARG ENV
ARG PACKTOKEN
ENV PATH /app/node_modules/.bin:$PATH
COPY ./web/package.json ./
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

