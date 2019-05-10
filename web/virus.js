var socket = new WebSocket('ws://localhost:8080');
var startMessage = {id: 3,  name: "start", message:"start"};
var stopMessage = {id: 3,  name: "stop", message:"stop"};
var clientMessage = {id: 0, name: "client", message:"Adding client"};
var contagiousDays = 0;
var lifeSpan = 0;


socket.addEventListener('open', function (event) {
  //socket.send(JSON.stringify('Hello Server!'));
  console.log("connected");
  socket.send(JSON.stringify(clientMessage));
});

// Added event listener for receiving messages
socket.addEventListener('message', function (event) {
  message = JSON.parse(event.data);
  
  // Message sorting and handling
  //Handling startMessage
  if (message.id === 1 && message.name === "start"){
      console.log("Simulation have started");
  }
  //Handling failMessage
  else if (message.id === 1 && message.name === "fail"){
    //2sec timeout on retrying to start server.
    setTimeout(function(){
      socket.send(JSON.stringify(startMessage));
    }, 2000); 
  }
  //Handling boardMessage
  else if (message.id === 1 && message.name === "board"){
  fillBoard(context,message.data);
      if (booleanPos){
          boardArray = message.data;
      }
      console.log("Updated Board");
  }else{
    console.log("-----------------------------");
    console.log("The following gave an error:");
    console.log(message.name);
    console.log(message.id);
    console.log("-----------------------------");
  }
});

function createGrid(context){
  for(var x = 0.5; x<1001; x+=10){
    context.moveTo(x, 0);
    context.lineTo(x,1000);
  }
  for(var y = 0.5; y<1001; y+=10){
    context.moveTo(0, y);
    context.lineTo(1000,y);
  }
    context.strokeStyle = "#ddd";
    context.stroke();

}


function fillCoordinate(context,x,y, color) {
      x = (1+x)*10-9;
      y = (1+y)*10-9;
      fillSquare(context, x, y, color);
}

function fillSquare(context, x, y, color){
    context.fillStyle = color;
    context.fillRect(x,y,9,9);
}

function componentToHex(c) {
  var hex = c.toString(16);
  return hex.length === 1 ? "0" + hex : hex;
}

function rgbToHex(r, g, b) {
  return "#" + componentToHex(r) + componentToHex(g) + componentToHex(b);
}

function fillBoard(context, board){
    for (var row = 0; row<board.length; row++) {
      for(var col = 0; col<board[row].length; col++){
        switch (board[row][col].state) {
          case -2:
            fillCoordinate(context, col, row, "#FFA500");
            //console.log("filling context with black");
            break;
          case -1:
            fillCoordinate(context, col, row, "#000000");
            //console.log("filling context with black");
            break;
          case 0:
                fillCoordinate(context, col, row, "#FFFF00");
            break;
          case 1:
            fillCoordinate(context, col, row, "#ffffff");
            break;
          default:
            if (board[row][col].state > (lifeSpan+1 - contagiousDays)){
              var r = Math.floor(255 / contagiousDays) * (board[row][col].state -(lifeSpan+1-contagiousDays));
              fillCoordinate(context, col, row, rgbToHex(255 -r,0,0));
            }else {
              var gb = Math.floor(255 / (lifeSpan+1-contagiousDays))*board[row][col].state;
              fillCoordinate(context, col, row, rgbToHex(255,255-gb,255-gb));
            }
            break;
        }
      }
    }
}

function getProperties(){
  var formElements = document.getElementById("form").elements;
  var virusMessage = {};
  for(var i = 0; i<formElements.length; i++)
  {
    if(formElements[i].type==="number")
    {
      virusMessage[formElements[i].name] = Number(formElements[i].value);
    }else if (formElements[i].type==="text"){
      virusMessage[formElements[i].name] = formElements[i].value;
    }
  }
  virusMessage.id = 0;
  lifeSpan = virusMessage.lifespan;
  contagiousDays = virusMessage.contagiousDays;
  virusMessage.posX = posX;
  virusMessage.posY = posY;
  sendMessage(virusMessage);
  sendMessage(startMessage);

}

function stopSimulation(){
  sendMessage(stopMessage);
}

function sendMessage(message){
  socket.send(JSON.stringify(message));
}
var canvas = document.getElementById('canvas');
var context = canvas.getContext('2d');
canvas.height = 1000;
canvas.width = 1000;
context.scale(1,1);
createGrid(context);
var posX;
var posY;
var booleanPos = false;
var boardArray;

function getCursorPosition(canvas, event) {
    var rect = canvas.getBoundingClientRect();
    var x = (event.clientX - rect.left) / (rect.right - rect.left) * canvas.width;
    var y = (event.clientY - rect.top) / (rect.bottom - rect.top) * canvas.height;
    x = x/10;
    y = y/10;
    x = Math.floor(x);
    y = Math.floor(y);
    var clickArray = [];
    posY = y;
    posX = x;
    booleanPos = true;
    clickArray["X-value"] = x;
    clickArray["Y-value"] = y;
    clearDiv("coordDiv");
    showCoords(clickArray);
}


function fillSquare(context, x, y, color){
    context.fillStyle = color
    context.fillRect(x,y,9,9);
}

function showHuman(human){
    var arr = human[posY][posX];
    console.log("posX: "+posX+ " posY: "+posY);
    console.log("x-value-arr: "+arr.pos.x+" : y-value-arr: "+arr.pos.y);
    var div = document.getElementById("coordDiv");
        console.log("the state is: "+arr.state);
        if(arr.state === -2){
            var doc = document.createTextNode("This human is immune by birth" + "<br/>");
        }else if(arr.state === -1) {
            var doc = document.createTextNode("This human is dead: "+  "<br/>");
        }else if(arr.state === 0) {
            var doc = document.createTextNode("This human survived "+arr.virus[0].name+" and became immune: "+ "<br/>");
        }else if(arr.state === 1){
            var doc = document.createTextNode("This human is healthy : "+  "<br/>");
        }else if(arr.state >=2 && arr.state <= (arr.virus[0].lifespan-arr.virus[0].contagiousdays)){
            var doc = document.createTextNode("This human recently became infected by "+arr.virus[0].name+" and is not contagious: "+ "<br/>");
        }else{
            var doc = document.createTextNode("This human is infected by "+arr.virus[0].name+" and is contagious: "+ "<br/>");
        }
    div.innerHTML += doc.data;
}
canvas.addEventListener('click', function(event) {
    getCursorPosition(canvas, event);
    showHuman(boardArray);
});

function showCoords(array){
    var div = document.getElementById("coordDiv");
    for (var c in array){
        var x = array[c].toString();
        var y = c.toString();  //Y-Coord, X-coord,
        var doc = document.createTextNode(y+" : "+x+ '<br/>');
        div.innerHTML += doc.data;
    }
}
function clearDiv(elementID){
    document.getElementById(elementID).innerHTML = "";
}
