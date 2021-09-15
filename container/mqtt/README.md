# MQTT on docker

<br />

- **First step**
  edit your volume path in ***docker-compose.yml***

<br />

- **run docker-compose file**
  `docker-compose up -d`

<br />

- **execute docker container**
  `docker exec -it <container_name> sh`

<br />

- **Create user and password**
  `mosquitto_passwd -c /mosquitto/config/mosquitto.passwd <username>`

<br />

- In ***config/mosquitto.conf***
  ``` 
    persistence true
    persistence_location /mosquitto/data/
    log_dest file /mosquitto/log/mosquitto.log

    password_file /mosquitto/config/mosquitto.passwd
    allow_anonymous false

    listener 1883
    listener 9001
    protocol websockets
  ```