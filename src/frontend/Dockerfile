FROM node:16

RUN mkdir -p /app
COPY . /app

WORKDIR /app
RUN yarn install

EXPOSE 3000

CMD [ "yarn", "dev" ]