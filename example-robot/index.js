const request = require('request');
 

function start() {

    request('http://127.0.0.1:8000/join', { json: true }, (err, res, body) => {
        if (err) { return console.log(err); }
        console.log(body.url);
        console.log(body.explanation);
      });

    // https.get(getUrl('/join', {weaponLevel: 1, engineLevel: 1, shieldLevel: 0})).then((response) => {
    //     response.json().then(function(data) {
    //         console.log(data);
    //     });
    // });

}



// function getUrl(url, params) {
//     let params = '?';
//     for (var i in json) {
//         params += i + '=' + json[i] + "&";
//     }
//     return url + params;
// }

start();