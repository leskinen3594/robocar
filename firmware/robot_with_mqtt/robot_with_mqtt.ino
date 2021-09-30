// Include WiFi library for esp32 or esp8266
#ifdef ESP32
  #include <WiFi.h>
#else
  #include <ESP8266WiFi.h>
#endif

// Include MQTT pub/sub library
#include <PubSubClient.h>
// JSON library
#include <ArduinoJson.h>
/* ------------------------------------------ End Header ------------------------------------------ */

/* ----------------------------------------- Configuration ---------------------------------------- */
// WiFi
const char* ssid    = "Anchelsea123_2.4GHz";
const char* passwd  = "0926624715";

// Get MAC address - 94:B9:7E:D5:AD:F4
const String mac_addr = WiFi.macAddress();

// Topics
// head + mac = subscriber
// head + usernamr = publisher
const String head_topic = "/mR_robot/";

// MQTT Broker
#define MQTT_SERVER   "192.168.1.22"
#define MQTT_PORT     1888
#define MQTT_USERNAME "bot"
#define MQTT_PASSWORD "P@ssw0rd"
#define MQTT_NAME     "Last_Order"

// define L298N or L293D motor control pins
int leftMotorForward    = 14;   // GPIO2(D4)  -> IN3
int rightMotorForward   = 12;   // GPIO15(D8) -> IN1
int leftMotorBackward   = 27;   // GPIO0(D3)  -> IN4
int rightMotorBackward  = 13;   // GPIO13(D7) -> IN2

// Ultrasonic HC-SR04
#define TRIG_PIN 26
#define ECHO_PIN 25
#define SOUND_SPEED 0.034
long duration;
float distanceCm;
unsigned long previousMillis = 0;
const long interval = 150;
bool dangerZone = false;
/* --------------------------------------- End Configuration -------------------------------------- */

WiFiClient client;
PubSubClient mqtt(client);

void setup() {
    Serial.begin(115200);

    /* initialize motor control pins as output */
    pinMode(leftMotorForward, OUTPUT);
    pinMode(rightMotorForward, OUTPUT);
    pinMode(leftMotorBackward, OUTPUT);
    pinMode(rightMotorBackward, OUTPUT);

    /* initialize ultrasonic pins */
    pinMode(TRIG_PIN, OUTPUT);
    pinMode(ECHO_PIN, INPUT);

    /* Connect to a WiFi network */
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, passwd);
    Serial.print("Connecting");
    while (WiFi.status() != WL_CONNECTED) {
      Serial.print(".");
      delay(3000);
    }
    // เชื่อมต่อสำเร็จ
    Serial.println("");
    Serial.println("WiFi connected");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
    Serial.print("ESP32 Board MAC Address: ");
    Serial.println(mac_addr);

    // connect to mqtt server
    connectMqtt();
}
/* --------------------------------------- End setup() -------------------------------------- */

void loop() {
    // ถ้าการเชื่อมต่อไวไฟหาด จะพยายามต่อใหม่จนกว่าจะสำเร็จ
    if (mqtt.connected() == false) {
        Serial.print( "WiFi Status : " );
        Serial.println( WiFi.status() );
        if (WiFi.status() != WL_CONNECTED) {
            reconnectWiFi();
        }

        Serial.println("MQTT reconnection...");
        connectMqtt();
    } // END reconnect

    /* Object detection */
    /*unsigned long currentMillis = millis();

    if (currentMillis - previousMillis >= interval) {
        previousMillis = currentMillis;

        clearTriggerPin();
        pingTriggerPin();

        distanceCm = getDistanceCm();

        if (distanceCm < 41.00) {
          Serial.print("Distance (cm): ");
          Serial.println(distanceCm);

          MotorStop();
          dangerZone = true;

        } else {
          dangerZone = false;
        }
    } // END detect*/

    // ถ้า MQTT ต่ออยู่แล้ว ก็ให้ทำงานใน callback
    mqtt.loop();
}
/* ---------------------------------------- End loop() --------------------------------------- */

// Ultrasonic
void clearTriggerPin() {
    digitalWrite(TRIG_PIN, LOW);
    delayMicroseconds(2);
}

void pingTriggerPin() {
    digitalWrite(TRIG_PIN, HIGH);
    delayMicroseconds(10);
    digitalWrite(TRIG_PIN, LOW);

    duration = pulseIn(ECHO_PIN, HIGH);
}

float getDistanceCm() {
    return duration * SOUND_SPEED / 2;
}
/* ------------------------------------ END Ultrasonic ------------------------------------ */

// Forward
void MotorForward(void){
  digitalWrite(leftMotorForward, HIGH);
  digitalWrite(rightMotorForward, HIGH);
  digitalWrite(leftMotorBackward, LOW);
  digitalWrite(rightMotorBackward, LOW);
}
/* --------------------------------- END MotorForward() --------------------------------- */

// Backward
void MotorBackward(void) {
  digitalWrite(leftMotorBackward, HIGH);
  digitalWrite(rightMotorBackward, HIGH);
  digitalWrite(leftMotorForward, LOW);
  digitalWrite(rightMotorForward, LOW);
}
/* --------------------------------- END MotorBackward() --------------------------------- */

// Turn Left
void TurnLeft(void) {
  digitalWrite(leftMotorForward, HIGH);
  digitalWrite(rightMotorForward, LOW);
  digitalWrite(rightMotorBackward, HIGH);
  digitalWrite(leftMotorBackward, LOW);
}
/* ------------------------------------ END TurnLeft() ------------------------------------ */

// Turn Right
void TurnRight(void){
  digitalWrite(leftMotorForward, LOW);
  digitalWrite(rightMotorForward, HIGH);
  digitalWrite(rightMotorBackward, LOW);
  digitalWrite(leftMotorBackward, HIGH);
}
/* ------------------------------------ END TurnRight() ------------------------------------ */

// Stop
void MotorStop(void) {
  digitalWrite(leftMotorForward, LOW);
  digitalWrite(leftMotorBackward, LOW);
  digitalWrite(rightMotorForward, LOW);
  digitalWrite(rightMotorBackward, LOW);
}
/* ------------------------------------ END MotorSTOP() ------------------------------------ */

// ถ้าไวไฟหลุด จะเรียกฟังก์ชันนี้ จนกว่าจะต่อใหม่ได้
void reconnectWiFi() {
    Serial.print("WiFi reconnecting... ");
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, passwd);
    while (WiFi.status() != WL_CONNECTED) {
      Serial.print(".");
      delay(3000);
    }
    Serial.println("Connected!");
}
/* ----------------------------------- End reconnectWiFi() ---------------------------------- */

// connect to MQTT broker
void connectMqtt() {
    // เซ็ตค่า mqtt server
    mqtt.setServer(MQTT_SERVER, MQTT_PORT);
    mqtt.setCallback(callback);

    Serial.print("MQTT connection... ");
    if (mqtt.connect(MQTT_NAME, MQTT_USERNAME, MQTT_PASSWORD)) {
      Serial.println("MQTT connected");

      /* ---- Subscriber ---- */
      // c_str() -> A pointer to the C-style version of the invoking String.
      // convert string to char array
      String sub_api = head_topic + mac_addr;
      const char* sub_topic = sub_api.c_str();
      mqtt.subscribe(sub_topic);  // Example topic: "/mR_robot/94:B9:7E:D5:AD:F4"

    } else {
      // ต่อ MQTT Server ล้มเหลว รอ 5000 mili-sec แล้วพยายามต่อใหม่
      Serial.println("failed");
      delay(5000);
    }
}
/* ------------------------------------ End connectMqtt() ----------------------------------- */

// ฟังก์ชันการทำงานของ MQTT
void callback(char* topic, byte* payload, unsigned int length) {
    // อ่านข้อความที่รับมา
    payload[length] = '\0';
    String topic_str = topic, payload_str = (char*)payload;
    Serial.println("sub - [" + topic_str + "]: " + payload_str);

    // username, message received from API
    char username[25] = "";
    char message[25]  = "";

    // Receiving JSON Data
    StaticJsonDocument<256> doc, reflect;
    deserializeJson(doc, payload, length);
    strlcpy(username, doc["username"], sizeof(username));
    strlcpy(message, doc["message"], sizeof(message));

    /*Serial.print("username = ");
    Serial.print(username);
    Serial.println("$");
    Serial.print("message = ");
    Serial.print(message);
    Serial.println("$");*/

    /* ---- Publisher ---- */
    String pub_api = head_topic + "ws/" + username;
    const char* pub_topic = pub_api.c_str();  // Example topic: "/mR_robot/ws/dolly"

    // Create publish data
    char out[256];
    reflect["mac_addr"] = mac_addr;
    reflect["username"] = username;
    
    // กำหนดเงื่อนไขที่แตกต่างกันตาม Topic
    /* ---- handshake ---- */
    if ( String(message) == "Ahoy!" ) {
      Serial.println( message );
      reflect["robot_msg"] = "Yo-ho!";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);
    }

    /* ---- Movement ---- */
    if ( String(message) == "forward" ) {
      Serial.println("Go ahead!!");
      MotorForward();
      reflect["robot_msg"] = "ok, forward";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);

    } else if ( String(message) == "backward" ) {
      Serial.println("Back back back!!");
      MotorBackward();
      reflect["robot_msg"] = "ok, backward";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);

    } else if ( String(message) == "left" ) {
      Serial.println("Turn Left");
      TurnLeft();
      reflect["robot_msg"] = "ok, left";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);

    } else if ( String(message) == "right" ) {
      Serial.println("Turn Right");
      TurnRight();
      reflect["robot_msg"] = "ok, right";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);

    } else if ( String(message) == "stop" ) {
      Serial.println("Stop!!");
      MotorStop();
      reflect["robot_msg"] = "ok, stop";
      size_t plength =  serializeJson(reflect, out);
      mqtt.publish(pub_topic, out, plength);
    }
}
/* ----------------------------------- End mqtt_callback() ----------------------------------- */
