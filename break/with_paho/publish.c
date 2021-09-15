#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "pahomqtt/src/MQTTClient.h"

/** MQTT broker config */
#define ADDRESS     "tcp://localhost:1883"
#define CLIENTID    "df125c9d-f966-45fc-a257-67522f3d2913"
#define TOPIC       "topic/testc"
#define PAYLOAD     "Ahoy!"
#define QOS         1
#define TIMEOUT     10000L


int main(int argc, char **argv) {
    /**
     * Create client object and token
     * Initialize connect options and publish message
     */
    MQTTClient client;
    MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;
    MQTTClient_message pubmsg           = MQTTClient_message_initializer;
    MQTTClient_deliveryToken token;
    int rc;

    /** Create Client */
    if ((rc = MQTTClient_create(&client, ADDRESS, CLIENTID,
        MQTTCLIENT_PERSISTENCE_NONE, NULL)) != MQTTCLIENT_SUCCESS) {
         printf("Failed to create client, return code %d\n", rc);
         exit(EXIT_FAILURE);
    }

    /** Setting Options and Authen */
    conn_opts.keepAliveInterval = 10;
    conn_opts.cleansession      = 1;
    conn_opts.username          = "guest";
    conn_opts.password          = "pass123";

    /** Connect to broker */
    if ((rc = MQTTClient_connect(client, &conn_opts)) != MQTTCLIENT_SUCCESS) {
        printf("Failed to connect, return code %d\n", rc);
        exit(EXIT_FAILURE);
    }

    /** Setup playload */
    pubmsg.payload    = PAYLOAD;
    pubmsg.payloadlen = (int)strlen(PAYLOAD);
    pubmsg.qos        = QOS;
    pubmsg.retained   = 0;

    /** Publish message */
    if ((rc = MQTTClient_publishMessage(client, TOPIC, &pubmsg, &token)) != MQTTCLIENT_SUCCESS) {
         printf("Failed to publish message, return code %d\n", rc);
         exit(EXIT_FAILURE);
    }

    /** Waiting for publish message */
    printf("Waiting for up to %d seconds for publication of %s\n"
            "on topic %s for client with ClientID: %s\n",
            (int)(TIMEOUT/1000), PAYLOAD, TOPIC, CLIENTID);

    /**
     * The delivery token can be used in a waitForCompletion call
     * to synchronize on the completion of the MQTT packet exchange
     */
    rc = MQTTClient_waitForCompletion(client, token, TIMEOUT);
    printf("Message with delivery token %d delivered\n", token);

    /** Disconnect */
    if ((rc = MQTTClient_disconnect(client, TIMEOUT)) != MQTTCLIENT_SUCCESS)
        printf("Failed to disconnect, return code %d\n", rc);

    /** free memory */
    MQTTClient_destroy(&client);

    return rc;
}
