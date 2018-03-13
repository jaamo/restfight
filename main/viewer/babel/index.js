const arenaSize = 10;
const arenaCellWidth = 40;

// Arena DOM element.
let arena = document.querySelector('.arena');

// List of all robots.
let robots = {};

/** 
 * Robot.
 */
class Robot {

  constructor(data, robotWidth) {
    this.robotWidth = robotWidth;
    this.setData(data);
    this.createElement();
    this.render();
  }

  createElement() {
    this.el = document.createElement('div');
    this.el.classList.add('robot');
    this.el.dataset.robotId = this.robot_id;
    document.querySelector('.arena').appendChild(this.el);
  }

  render() {
    this.el.style.left = this.x * this.robotWidth + 'px';
    this.el.style.top = this.x * this.robotWidth + 'px';
  }

  setData(data) {
    this.data = data;
  }

  get robot_id() {
    return this.data.robot_id;
  }

  get x() {
    return this.data.x;
  }

  get y() {
    return this.data.x;
  }

}

let eventHandlers = {
  'JOIN_GAME': (event) => {
    let robot = new Robot(event.robot, arenaCellWidth);
    robots[event.robot.robot_id] = robot;
  }
}

/**
 * Render arena.
 */
function initArena() {
  for (let y = 0; y < arenaSize; y++) {
    for (let x = 0; x < arenaSize; x++) {
      let cell = document.createElement('div');
      cell.classList.add('cell');
      cell.style.left = x * arenaCellWidth + 'px';
      cell.style.top = y * arenaCellWidth + 'px';
      arena.appendChild(cell); 
    }
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