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

    updateRobotsLegend(event.status.robots, event.status.active_robot);

    updateArena(event.status.arena);

    updateRobots(event.status.robots);

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
    // let robot = new Robot(event.robot, arenaCellWidth);
    // robots[event.robot.robot_id] = robot;
    // console.log(robots);
  },

  'NEW_TURN': (event) => {

    updateRobotsLegend(event.status.robots, event.status.active_robot);
    
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

  let socket = new WebSocket("ws://" + window.location.host + "/socket");  
  
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

function updateRobotsLegend(robots, activeRobot) {

  if (!robots) {
    return;
  }
  robots.forEach((robot, i) => {

    let active = '';
    if (activeRobot == robot.robot_index) {
      active = '*';
    }

    document.querySelector('.js-robot'+i+'-title').innerHTML = active + 'robot' + i;
    document.querySelector('.js-robot'+i+'-x').innerHTML = robot.x;
    document.querySelector('.js-robot'+i+'-y').innerHTML = robot.y;
    document.querySelector('.js-robot'+i+'-health').innerHTML = robot.health;
    document.querySelector('.js-robot'+i+'-max-health').innerHTML = robot.max_health;
    document.querySelector('.js-robot'+i+'-max-moves').innerHTML = robot.max_moves;
    document.querySelector('.js-robot'+i+'-weapon-range').innerHTML = robot.weapon_range;
    document.querySelector('.js-robot'+i+'-weapon-power').innerHTML = robot.weapon_power;
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

function updateRobots(robotsData) {

  if (!robotsData) {
    return;
  }

  // Update robot locations.
  robotsData.forEach((robotData, i) => {

    // debugger;
    // console.log(robots);

    // Robot is missing. Add.
    if (typeof(robots[robotData.robot_id]) == "undefined") {
      // console.log('Robot ' + robotData.robot_id +' doesn\'t exists. Let\'s create one.');
      let robot = new Robot(robotData, arenaCellWidth);
      robots[robotData.robot_id] = robot; 
    } else {
      robots[robotData.robot_id].x = robotData.x
      robots[robotData.robot_id].y = robotData.y  
    }

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