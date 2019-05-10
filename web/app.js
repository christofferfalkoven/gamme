const express = require('express');
var url = require('url');
var http = require('http');
var fs = require('fs');
const WebSocket = require('ws');

const app = express();
 
app.use(function (req, res) {
  var q = url.parse(req.url, true);
  var filename = "." + q.pathname;
  fs.readFile(filename, function(err, data) {
    if(err){
      res.writeHead(404, {'Content-Type': 'text/html'});
      return res.end("404 Not Found");
    }
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write(data);
    res.end();
  });
});

const server = http.createServer(app);
const wss = new WebSocket.Server({ server });
var clients = [];
var goserver = null;
var failedMess = {id:1, name:"fail", message:"failed to start simulation"};
var clientMess = {id:1, name:"start", message:"Add client to client list"};
var serverMess = {id:3, name:"start", message:"start the simulation"};

wss.on('connection', function connection(ws, req) {
  
  // You might use location.query.access_token to authenticate or share sessions
  // or req.headers.cookie (see http://stackoverflow.com/a/16395220/151312)
  ws.on('message', function incoming(message) {
    console.log("Passing message towards next server");
    var m = JSON.parse(message);
    if (m.name === "goserver" && m.id === 0){
      if (goserver === null){
        goserver = ws;
        console.log("goserver added!");
      }else{
        console.log("goserver already added!");
      }
    }
    else if (m.name === "client" && m.id === 0){
      if (clients.indexOf(ws) < 0){
        clients.push(ws);
        console.log("client added!");
      }else{
        console.log("client already added!");
      }
    }
    else if (m.name === "start" && m.id === 0){
      if(goserver === null){
        ws.send(JSON.stringify(failedMess));
        console.log("failed to start server");
      }else{
        goserver.send(JSON.stringify(serverMess));
        ws.send(JSON.stringify(clientMess));
      }
    }
    else if (m.id === 0){
      // received virus info.
      console.log("sending message to goserver");
      goserver.send(message);
    }
    else if (m.id === 1){
      console.log("Received a message, forwarding towards clients");
      clients.forEach(element => {
        try{
        element.send(message);
        }catch(e){
          console.log("Disconnected socket in client list, removing it");
          var index = clients.indexOf(element);
          clients.splice(index,1);
          console.log("Removed");
        }
      });
    }
    else if (JSON.parse(message).id == 3){
      goserver.send(message);
    }
  });
 
});

 
server.listen(8080, function listening() {
  console.log('Listening on %d', server.address().port);
});
