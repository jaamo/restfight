const arenaSize = 10;

let arena = document.querySelector('.arena');

let eventHandlers = {
  'JOIN_GAME': (event) => {
    let robot = document.createElement('div');
    robot.classList.add('robot');
    robot.dataset.robotId = event.robot.robot_id;
    document.querySelector('body').appendChild(robot);
  }
}

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
    handleEvent(JSON.parse(event.data));
  };

}

function handleEvent(event) {
  console.log(event);
  log(event);
  if (typeof(eventHandlers[event.event_type]) != 'undefined') {
    eventHandlers[event.event_type](event);
  } else {
    log('Unknown event handler ' + event.event_type);
  }
}

function log(msg) {

  if (typeof(msg) == 'object') {
    msg = JSON.stringify(msg);
  }

  document.querySelector('.console').innerHTML = msg + "<br>" + document.querySelector('.console').innerHTML;

}

function init() {
  initArena();
  initWebSocket();
}

init();