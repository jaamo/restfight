# Tutorial: Starting the server

In this tutorial we will launch the game server and run an example robot.

## Requirements

You will need to have go lang and node tools installed.

## 1. Clone the repo

Clone the repo to your computer:  
`git@github.com:jaamo/restfight.git`

## 2. Install go dependencies

`cd restfight`  
`RUN go get -d -v ./...`  
`RUN go install -v ./...`

## 3. Run the server

`./main`

## 4. Make sure the server is running

Load the url with your browser:  
`http://127.0.0.1:8000/viewer`

You should see the arena.

## 5. Run example robot

Open new terminal window, go to example-robot folder and run the first robot:
`node index.js`

If you now check your browser you should see the first robot on the arena.

## 6. Run another example robot

Open new terminal window and repeat step 5. You should now see two robots fighting!

## 7. Enjoy

You should now see two robots fighting :) If not... better to start debugging! Have fun.