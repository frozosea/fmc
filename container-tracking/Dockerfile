FROM node:14-alpine as builder
MAINTAINER User="i20072004@gmail.com"
WORKDIR src/
COPY package*.json ./
RUN npm install
EXPOSE 8002
CMD ["npm","run","start"]


