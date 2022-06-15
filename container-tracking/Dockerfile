FROM node:14-alpine as builder
MAINTAINER User="i20072004@gmail.com"
WORKDIR src/
COPY package*.json ./
RUN npm install
EXPOSE 8080
CMD ["npm","run","start"]


