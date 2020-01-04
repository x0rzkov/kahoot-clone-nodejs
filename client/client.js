//Import dependencies
const path = require('path');
const http = require('http');
const express = require('express');
const publicPath = path.join(__dirname, '../public');

var app = express();
var server = http.createServer(app);
app.use(express.static(publicPath));

//Starting server on port 3000
server.listen(3000, () => {
    console.log("Server started on port 3000");
});