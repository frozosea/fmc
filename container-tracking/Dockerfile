FROM node:14-alpine as builder
WORKDIR src/
COPY package*.json ./
RUN npm install
EXPOSE 8002
CMD ["npm","run","start"]


