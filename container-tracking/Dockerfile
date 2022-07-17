FROM node:16-alpine
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app/dist

# Install app dependencies from package.json. If modules are not included in the package.json file enter a RUN command. E.g. RUN npm install <module-name>
COPY *.json /usr/src/app/
RUN     npm install npm --location=global
RUN     npm install
RUN     npm install -g ts-node
# Bundle app source
COPY . /usr/src/app
COPY entrypoint.sh /usr/src/app

# Enter the command which should be used when the image starts up. E.g. CMD ["node", "app.js"]
EXPOSE ${GRPC_PORT}
ENTRYPOINT [ "npm","run","start"]
#CMD [ "npm","run","start"]


