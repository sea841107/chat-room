FROM node:16
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm run install
COPY . .
RUN npm run build
EXPOSE 8080
CMD [ "npm", "run", "start" ]