#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "pahomqtt/src/MQTTClient.h"

/** Config MQTT Broker */
#define ADDRESS     "tcp://localhost:1883"
#define CLIENTID    "8d893304-5122-4000-9564-efdc08357a6f"
#define TOPIC       "topic/testc"
#define PAYLOAD     "Ahoy!"
#define QOS         2
#define TIMEOUT     10000L


volatile MQTTClient_deliveryToken deliveredtoken;

void delivered(void *context, MQTTClient_deliveryToken dt) {
    printf("Message with token value %d delivery confirmed\n", dt);
    deliveredtoken = dt;
}


int msgarrvd(void *context, char *topicName, int topicLen, MQTTClient_message *message) {
    printf("Message arrived\n");
    printf("    topic: %s\n", topicName);
    printf("    message: %.*s\n", message->payloadlen, (char*)message->payload);
    MQTTClient_freeMessage(&message);
    MQTTClient_free(topicName);

    return 1;
}


void connlost(void *context, char *cause) {
    printf("\nConnection lost\n");
    printf("     cause: %s\n", cause);
}


int main(int argc, char **argv) {
    MQTTClient client;
    MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;
    int rc;

    if ((rc = MQTTClient_create(&client, ADDRESS, CLIENTID,
        MQTTCLIENT_PERSISTENCE_NONE, NULL)) != MQTTCLIENT_SUCCESS) {
        printf("Failed to create client, return code %d\n", rc);
        rc = EXIT_FAILURE;
        goto exit;
    }

    /** Callback */
    if ((rc = MQTTClient_setCallbacks(client, NULL, connlost, msgarrvd, delivered)) != MQTTCLIENT_SUCCESS) {
        printf("Failed to set callbacks, return code %d\n", rc);
        rc = EXIT_FAILURE;
        goto destroy_exit;
    }

    conn_opts.keepAliveInterval         = 10;
    conn_opts.cleansession              = 1;
    conn_opts.username                  = "guest";
    conn_opts.password                  = "pass123";

    if ((rc = MQTTClient_connect(client, &conn_opts)) != MQTTCLIENT_SUCCESS) {
        printf("Failed to connect, return code %d\n", rc);
        rc = EXIT_FAILURE;
        goto destroy_exit;
    }

    printf("Subscribing to topic %s\nfor client %s using QOS%d\n\n"
           "Press Q <Enter> to quit\n\n", TOPIC, CLIENTID, QOS);

    if ((rc = MQTTClient_subscribe(client, TOPIC, QOS)) != MQTTCLIENT_SUCCESS) {
    	printf("Failed to subscribe, return code %d\n", rc);
    	rc = EXIT_FAILURE;

    } else {
        int ch;

        do {
        	ch = getchar();
    	} while (ch != 'Q' && ch != 'q');

        if ((rc = MQTTClient_unsubscribe(client, TOPIC)) != MQTTCLIENT_SUCCESS) {
        	printf("Failed to unsubscribe, return code %d\n", rc);
        	rc = EXIT_FAILURE;
        }
    }

     if ((rc = MQTTClient_disconnect(client, TIMEOUT)) != MQTTCLIENT_SUCCESS) {
    	printf("Failed to disconnect, return code %d\n", rc);
    	rc = EXIT_FAILURE;
    }

    destroy_exit:
        MQTTClient_destroy(&client);

    exit:
        return rc;
}
