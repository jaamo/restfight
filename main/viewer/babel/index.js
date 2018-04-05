const arenaSize = 10;
const arenaCellWidth = 40;

// Arena DOM element.
let arena = document.querySelector('.arena');

// List of all robots.
let robots = {};

// Explosion.
let explosion = null;

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
    this.el.style.left = this.data.x * this.robotWidth + 'px';
    this.el.style.top = this.data.y * this.robotWidth + 'px';
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

  set x(x) {
    this.data.x = x;
    this.render();
  }

  set y(y) {
    this.data.y = y;
    this.render();
  }  
}

let eventHandlers = {

  'STATUS': (event) => {

    updateRobotsLegend(event.status.robots);

    updateArena(event.status.arena);

    if (event.status.status == 2) {
      document.querySelector('.gameover').classList.add('gameover--visible');
    } else {
      document.querySelector('.gameover').classList.remove('gameover--visible');      
    }

  },

  'SHOOT': (event) => {
    explosion.style.left = event.x * arenaCellWidth + 'px';
    explosion.style.top = event.y * arenaCellWidth + 'px';
    explosion.classList.add('blink');
    setTimeout(() => {
      explosion.classList.remove('blink');
    }, 5000);
  },

  'JOIN_GAME': (event) => {
    let robot = new Robot(event.robot, arenaCellWidth);
    robots[event.robot.robot_id] = robot;
  },

  'ROBOT_MOVED': (event) => {
    robots[event.robot.robot_id].x = event.robot.x;
    robots[event.robot.robot_id].y = event.robot.y;
  },

  'NEW_GAME': (event) => {
    initGame();
  }
  
}

/**
 * Render arena.
 */
function initGame() {
  arena.innerHTML = '';
  robots = {};
  for (let y = 0; y < arenaSize; y++) {
    for (let x = 0; x < arenaSize; x++) {
      let cell = document.createElement('div');
      cell.classList.add('cell');
      cell.dataset.x = x;
      cell.dataset.y = y;
      cell.style.left = x * arenaCellWidth + 'px';
      cell.style.top = y * arenaCellWidth + 'px';
      arena.appendChild(cell); 
    }
  }  

  explosion = document.createElement('div');
  explosion.classList.add('explosion');
  document.querySelector('.gameover').classList.remove('gameover--visible');
  // cell.classList.add('blink');
  arena.appendChild(explosion); 

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
  // console.log(event);
  log(event);
  if (typeof(eventHandlers[event.event_type]) != 'undefined') {
    eventHandlers[event.event_type](event);
  } else {
    log('Unknown event handler ' + event.event_type);
  }
}

function updateRobotsLegend(robots) {

  if (!robots) {
    return;
  }

  let legend = document.querySelector('.robots-legend');
  legend.innerHTML = '';
  robots.forEach((robot) => {
    legend.innerHTML += 'id: ' + robot.robot_id + '<br>';
    legend.innerHTML += 'health: ' + robot.health + '<br>';
    legend.innerHTML += 'x: ' + robot.x + '<br>';
    legend.innerHTML += 'y: ' + robot.y + '<br>';
    legend.innerHTML += '<br>';
  })

}

function updateArena(arena) {
  arena.forEach((row) => {
    row.forEach((cell) => {
      if (cell.type == 2) {
        document.querySelector('div[data-x="'+cell.x+'"][data-y="'+cell.y+'"]').classList.add('cell--obstacle')
      }
    })
  })
}

function log(msg) {

  console.log(msg);

  // if (typeof(msg) == 'object') {
  //   msg = JSON.stringify(msg);
  // }

  // document.querySelector('.console').innerHTML = msg + "<br>" + document.querySelector('.console').innerHTML;

}

function init() {
  initGame();
  initWebSocket();
}

init();