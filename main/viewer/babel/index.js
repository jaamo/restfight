const arenaSize = 10;
let arena = document.querySelector('.arena');

/**
 * Render arena.
 */
function initArena() {
  for (let y = 0; y < arenaSize; y++) {
    let row = document.createElement('div');
    row.classList.add('row');
    for (let x = 0; x < arenaSize; x++) {
      let cell = document.createElement('div');
      cell.classList.add('cell');
      cell.innerHTML = '';
      row.appendChild(cell); 
    }
    arena.appendChild(row);
  }  
}

function initWebSocket() {

  let socket = new WebSocket("ws://127.0.0.1:8000/socket");  
  
  socket.onopen = function (event) {
    log('Connection to server opened.');
    socket.send("Here's some text that the server is urgently awaiting!"); 
  };  

  socket.onmessage = function(event) {
    log(event.data);    
  };

}

function log(msg) {
  document.querySelector('.console').innerHTML = msg + "<br>" + document.querySelector('.console').innerHTML;
}

function init() {
  initArena();
  initWebSocket();
}

init();