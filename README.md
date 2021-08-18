# Paho MQTT in C

<br />

### Download and Building from source
``` can rename directory.
git clone https://github.com/eclipse/paho.mqtt.c.git
cd paho.mqtt.c
make
sudo make install
```

<br />

### Compile the code
`gcc <filename.c> -o <output_file> -lpaho-mqtt3c`


<br /><br />

**NOTE**
- Publisher and Subscriber do not same CLIENTID!
- [reference](https://www.eclipse.org/paho/index.php?page=clients/c/index.php)