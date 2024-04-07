FROM node:20.10-alpine as build

WORKDIR /build

COPY . /build

RUN npm install

RUN npm run build

FROM node:20.10-alpine

WORKDIR /app

COPY package.json /app
COPY package-lock.json /app

COPY --from=build /build/build /app

RUN npm ci --omit dev

CMD [ "node", "index.js" ]