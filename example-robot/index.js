// const request = require('request');
const fetch = require('node-fetch');
const querystring = require('querystring');

const baseUrl = 'http://127.0.0.1:8000/';

let robot = {};
let status = {};

function newGame() {

    console.log('newGame');

    // New game.
    fetch(baseUrl + 'new').then(response => {
        joinGame();
    });

}

function joinGame() {

    console.log('joinGame');

    const params = {
        engineLevel: 1, 
        shieldLevel: 1, 
        weaponLevel: 0
    };

    fetch(baseUrl + 'join?' + querystring.stringify(params)).then(response => {
        response.json().then(json => { 
            robot = json;
            console.log(robot);
            waitForTurn();
        });
    });    

}

function waitForTurn() {

    console.log('waitForTurn');

    fetch(baseUrl + 'status').then(response => {
        response.json().then(json => {

            json.arena = [];

            if (json.status == 1 && json.active_robot == robot.robot_index) {

                getStatusAndPlayTurn();

            } else {

                setTimeout(() => {
                    waitForTurn();
                }, 5000)

            }

        });
    });

}

function getStatusAndPlayTurn() {

    console.log('getStatusAndPlayTurn');
    fetch(baseUrl + 'status').then(response => {
        response.json().then(json => {
            status = json;
            robot = getRobot();
            setTimeout(() => {
                playTurn();
            }, 500);
        });
    });    

}

function playTurn() {

    console.log('playTurn');

    // No moves left. Quit!
    if (robot.moves == robot.max_moves || robot.weapon_ammo == 0) {
        console.log('Out of moves: ' + robot.moves + '/' + robot.max_moves);
        fetch(baseUrl + 'endturn?' + querystring.stringify({robot_id: robot.robot_id})).then((response) => {
            response.text().then(text => {
                console.log(text);
                waitForTurn();
            });
        });
        return;
    }

    // Get enemy.
    let enemy = getEnemy();

    // If enemy is close enough, SHOOT!
    if (enemyInRange()) {

        console.log('SHOOT!');

        let coords = {
            robot_id: robot.robot_id,
            x: enemy.x,
            y: enemy.y
        }
        
        fetch(baseUrl + 'shoot?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                console.log(json);
                getStatusAndPlayTurn();
            });
        });        

    } 
    // Move horizontally.
    else if (enemy.x != robot.x) {

        console.log('Move horizontally.');

        let coords = {
            robot_id: robot.robot_id,
            x: robot.x,
            y: robot.y
        }
        
        if (enemy.x < robot.x) {
            coords.x--;
        } else {
            coords.x++;
        }

        fetch(baseUrl + 'move?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                console.log(json);
                getStatusAndPlayTurn();
            });
        });

    }     
    // Move vertically.
    else {

        console.log('Move vertically.');

        let coords = {
            robot_id: robot.robot_id,
            x: robot.x,
            y: robot.y
        }
        
        if (enemy.y < robot.y) {
            coords.y--;
        } else {
            coords.y++;
        }

        fetch(baseUrl + 'move?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                console.log(json);
                getStatusAndPlayTurn();
            });
        });
        
    }

    // for (let i = 0; i < robot.max_moves; i++) {

    //     let coords = {
    //         robot_id: robot.robot_id,
    //         x: robot.x + i,
    //         y: robot.y
    //     }
    //     // let response = await fetch(baseUrl + 'move?' + querystring.stringify(coords));
    //     // let json = await response.json();

    //     console.log(json);
        
    // }
 
}

function enemyInRange() {

    let enemy = getEnemy();

    if (Math.abs(enemy.x - robot.x) <= robot.weapon_range 
        && Math.abs(enemy.y - robot.y) <= robot.weapon_range) {
        return true;
    } else {
        return false;
    }

}

/** 
 * Return enemy robot.
 */
function getEnemy() {

    for (r of status.robots) {
        if (r.robot_index != robot.robot_index) {
            return r;
        }
    }

}

function getRobot() {

    for (r of status.robots) {
        if (r.robot_index == robot.robot_index) {
            return r;
        }
    }

}

// newGame();
joinGame();