/* Serial number - 0      ; Codename - Dolly      */
/* Serial number - 20001  ; Codename - Last Order */

// Include WiFi library for esp32 or esp8266
#ifdef ESP32
  #include <WiFi.h>
#else
  #include <ESP8266WiFi.h>
#endif

// Include MQTT pub/sub library
#include <PubSubClient.h>
/* ---------------------------------------------- End Header ---------------------------------------------- */

/* --------------------------------------------- Configuration -------------------------------------------- */
// WiFi
const char* ssid    = "Anchelsea123_2.4GHz";
const char* passwd  = "0926624715";

// Get MAC address - 94:B9:7E:D5:AD:F4
const String mac_addr = WiFi.macAddress();

// Topics
String handshake  = "/caro_bot/handshake/" + mac_addr;
String movement   = "/caro_bot/movement/" + mac_addr;

// MQTT Broker
#define MQTT_SERVER   "192.168.1.22"
#define MQTT_PORT     1888
#define MQTT_USERNAME "bot"
#define MQTT_PASSWORD "P@ssw0rd"
#define MQTT_NAME     "Last_Order"
/* ------------------------------------------- End Configuration ------------------------------------------ */

WiFiClient client;
PubSubClient mqtt(client);

void setup() {
    Serial.begin(115200);

    /* Connect to a WiFi network */
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, passwd);
    Serial.print("Connecting");
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.print(".");
    }
    // เชื่อมต่อสำเร็จ
    Serial.println("");
    Serial.println("WiFi connected");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
    Serial.print("ESP32 Board MAC Address: ");
    Serial.println(mac_addr);

    // เซ็ตค่า mqtt server
    mqtt.setServer(MQTT_SERVER, MQTT_PORT);
    mqtt.setCallback(mqtt_callback);
}
/* ---------------------------------------------- End setup() --------------------------------------------- */

// ถ้าไวไฟหลุด จะเรียกฟังก์ชันนี้ จนกว่าจะต่อใหม่ได้
void reconnectWiFi() {
    Serial.print("Reconnecting");
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, passwd);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.print(".");
    }
    Serial.println("Connected!");
}
/* ------------------------------------------ End reconnectWiFi() ----------------------------------------- */

// ฟังก์ชันการทำงานของ MQTT
void mqtt_callback(char* topic, byte* payload, unsigned int length) {
    // อ่านข้อความที่รับมา
    payload[length] = '\0';
    String topic_str = topic, payload_str = (char*)payload;
    Serial.println("[" + topic_str + "]: " + payload_str);

    // กำหนดเงื่อนไขที่แตกต่างกันตาม Topic
    /* ---- handshake ---- */
    if ( topic_str == handshake ) {
      Serial.print( "sub - /caro_bot/handshake : " );
      Serial.println( payload_str );

      if ( payload_str == "Ahoy!" ) {
        mqtt.publish(topic_str.c_str(), "Yo-ho!");
      }

      /* ---- Movement ---- */
    } else if ( topic_str == movement ) {
      Serial.print( "sub - /caro_bot/movement : " );
      Serial.println( payload_str );

      if ( payload_str == "Forward" ) {
        Serial.println("Go ahead!!");

      } else if ( payload_str == "Backward" ) {
        Serial.println("Back back back!!");

      } else if ( payload_str == "Left" ) {
        Serial.println("Turn Left");

      } else if ( payload_str == "Right" ) {
        Serial.println("Turn Right");

      } else {
        Serial.println("Stop!!");
      }

    } // End topic_str == movement
}
/* ------------------------------------------- End mqtt_callback() ------------------------------------------ */

// เพิ่ม Mac address ลงฐานข้อมูลหากยังงไม่มี
void init() {
    /** Check;
     * SELECT `MAC_addr` FROM `Controller` WHERE `MAC_addr` LIKE "mac_addr";
     * IF NOT mac_addr IN database
     * THEN insert mac_addr TO database
     */

     //if (!getMAC) mqtt.publish("/caro_bot/addMAC", mac_addr.c_str());  // insert
}
/* ----------------------------------------------- End init() ----------------------------------------------- */

void loop() {
     // ถ้าการเชื่อมต่อไวไฟหาด จะพยายามต่อใหม่จนกว่าจะสำเร็จ
    if (mqtt.connected() == false) {
      Serial.print( "WiFi Status : " );
      Serial.println( WiFi.status() );
      if (WiFi.status() != WL_CONNECTED) {
        reconnectWiFi();
      }

      // พยายามเชื่อมต่อ MQTT server
      Serial.print("MQTT connection... ");
      if (mqtt.connect(MQTT_NAME, MQTT_USERNAME, MQTT_PASSWORD)) {
        // ต่อสำเร็จ แล้ว subscribe topics
        Serial.println("connected");

        /* ---- Subscriber ---- */
        const char* sub_handshake = handshake.c_str();  // c_str() -> A pointer to the C-style version of the invoking String.
        const char* sub_movement  = movement.c_str();
        mqtt.subscribe(sub_handshake);  // Topic: "/caro_bot/handshake/94:B9:7E:D5:AD:F4";
        mqtt.subscribe(sub_movement);

      } else {
        // ต่อ MQTT Server ล้มเหลว รอ 5000 mili-sec แล้วพยายามต่อใหม่
        Serial.println("failed");
        delay(5000);
      }

    } else {
      // ถ้า MQTT ต่ออยู่แล้ว ก็ให้ทำงานใน callback
      mqtt.loop();
    }
}
/* ------------------------------------------------ End loop() ----------------------------------------------- */
