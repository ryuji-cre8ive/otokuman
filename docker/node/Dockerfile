FROM node:14.19-alpine

WORKDIR /app

RUN apk update
RUN apk add yarn

ADD package.json ./

COPY . .

RUN yarn install --production=false
# RUN yarn global add @vue/cli@4.5.15