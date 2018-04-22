/**
 * Example robot for RESTFight game. This robot is super dummy. 
 * The idea is to show how to interact with the API.
 */

const fetch = require('node-fetch');
const querystring = require('querystring');

// Base url. Point this to your server.
const baseUrl = 'http://127.0.0.1:8000/';

// Robot id.
let robotId = 0;

// Game status returned from the server.
let status = {};

// Start everything!
joinGame();

/** 
 * Joins a game.
 */
function joinGame() {

    console.log('Join game.');

    // Define our robot properties. Plese refer API docs & game rules for more information.
    const params = {
        engineLevel: 1, 
        shieldLevel: 1, 
        weaponLevel: 0
    };

    // Join the game. Robot is returned as a response. Save that for future use.
    fetch(baseUrl + 'join?' + querystring.stringify(params)).then(response => {
        response.json().then(json => { 
            robotId = json.robot_id;
            console.log('Joined to game. Robot id is ' + robotId);
            waitForTurn();
        });
    });    

}

/** 
 * Poll status until the game is started.
 */
function waitForTurn() {

    console.log('Wait for turn...');

    // Request status. Remember to pass robot ID to get more meaningful response.
    fetch(baseUrl + 'status?' + querystring.stringify({robot_id: robotId})).then(response => {
        response.json().then(json => {

            // Check if it's your turn. If not, wait for 5 seconds and retry.
            if (json.is_your_turn == 1) {
                console.log('It\'s our turn!')
                getStatusAndPlayTurn();
            } else {
                setTimeout(() => {
                    waitForTurn();
                }, 5000)
            }

        });
    });

}

/** 
 * Update game status before playing turn.
 */
function getStatusAndPlayTurn() {

    console.log('Get status.');

    fetch(baseUrl + 'status?' + querystring.stringify({robot_id: robotId})).then(response => {
        response.json().then(json => {

            // Save status.
            status = json;
            // console.log(status);

            // A little delay to make debugging easier.
            setTimeout(() => {
                playTurn();
            }, 500);

        });
    });    

}

/**
 * Play turn! This the place for magical AI code.
*/
function playTurn() {

    console.log('Play turn.');

    // No moves left or out of ammo => quit.
    if (status.robot.moves == 0 || status.robot.weapon_ammo == 0) {
        console.log('Out of moves: ' + status.robot.moves + '/' + status.robot.max_moves);
        fetch(baseUrl + 'endturn?' + querystring.stringify({robot_id: status.robot.robot_id})).then((response) => {
            response.text().then(text => {
                waitForTurn();
            });
        });
        return;
    }

    // Get enemy.
    let enemy = status.enemies[0];

    // If enemy is close enough, SHOOT!
    if (enemyInRange()) {

        console.log('Shoot to location ' + enemy.x + ' x ' + enemy.y + '.');

        let coords = {
            robot_id: status.robot.robot_id,
            x: enemy.x,
            y: enemy.y
        }        
        fetch(baseUrl + 'shoot?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                getStatusAndPlayTurn();
            });
        });        

    } 
    // Move horizontally.
    else if (enemy.x != status.robot.x) {

        let coords = {
            robot_id: status.robot.robot_id,
            x: status.robot.x,
            y: status.robot.y
        }
        
        if (enemy.x < status.robot.x) {
            coords.x--;
        } else {
            coords.x++;
        }

        console.log('Move horizontally to location ' + coords.x + ' x ' + coords.y + '.');

        fetch(baseUrl + 'move?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                getStatusAndPlayTurn();
            });
        });

    }     
    // Move vertically.
    else {

        console.log('Move vertically.');

        let coords = {
            robot_id: status.robot.robot_id,
            x: status.robot.x,
            y: status.robot.y
        }
        
        if (enemy.y < status.robot.y) {
            coords.y--;
        } else {
            coords.y++;
        }

        console.log('Move horizontally to location ' + coords.x + ' x ' + coords.y + '.');

        fetch(baseUrl + 'move?' + querystring.stringify(coords)).then((response) => {
            response.json().then(json => {
                getStatusAndPlayTurn();
            });
        });
        
    }
 
}

/**
 * Return true if enemy is in range.
 */
function enemyInRange() {

    let enemy = status.enemies[0];

    if (Math.abs(enemy.x - status.robot.x) <= status.robot.weapon_range 
        && Math.abs(enemy.y - status.robot.y) <= status.robot.weapon_range) {
        return true;
    } else {
        return false;
    }

}
