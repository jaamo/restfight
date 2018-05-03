# Tutorial: Starting docker server

You can run the server in a docker container. 

Build the image:  
`docker build . --tag restfight`

Run the image:  
`docker run -d -p 8000:8000 restfight`