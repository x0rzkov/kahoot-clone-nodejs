FROM mhart/alpine-node:latest
WORKDIR /app
COPY . .
RUN npm install

EXPOSE 3000

CMD ["node", "server/server.js"]
