FROM node:14-alpine as builder
WORKDIR src/
COPY package*.json ./
RUN npm_config_target_arch=x64 npm install
EXPOSE 8002
CMD ["npm","run","start"]


