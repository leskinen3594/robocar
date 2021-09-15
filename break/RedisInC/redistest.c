#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <hiredis/hiredis.h>

/** Used for sleep() */
/** Cross platform */
#ifdef _WIN32
    #include <Windows.h>
#else
    #include <unistd.h>
#endif


int main(int argc, char **argv) {
    redisContext *c;
    redisReply *reply;
	const char *host = (argc > 1) ? argv[1] : "127.0.0.1";

    int port = (argc > 2) ? atoi(argv[2]) : 6379;

	/** printf("host = %s\n", host); */
	/** printf("port = %d\n", port); */

	c = redisConnect(host, port);

    if (c == NULL || c->err) {
        if (c) {
            printf("Connection error: %s\n", c->errstr);
            redisFree(c);

        } else {
            printf("Connection error: can't allocate redis context\n");
        }
        exit(1);
    }


    /* AUTH */
    reply = redisCommand(c, "AUTH %s", "pass123");
    if (reply->type == REDIS_REPLY_ERROR) {
        printf("Access denied.\n");
    }
    printf("Auth: %s\n", reply->str);
    freeReplyObject(reply);

    /* PING */
    reply = redisCommand(c, "PING");
    printf("PING: %s\n", reply->str);
    freeReplyObject(reply);

    /* SET a key */
    /** reply = redisCommand(c,"SET %s %s", "test", "hello world"); */
    /** printf("SET: %s\n", reply->str); */
    /** freeReplyObject(reply); */

    /* GET a key */
    /** reply = redisCommand(c, "GET test"); */
    /** printf("GET test: %s\n", reply->str); */
    /** freeReplyObject(reply); */

    /* Publisher */
    for (int i = 0; i < 10; i++) {
        reply = redisCommand(c, "PUBLISH topic1 %s%d", "test", i+1);
        printf("Publish: %s\n", reply->str);
        freeReplyObject(reply);
        sleep(1);   // delay 1 second
    }


    /* Disconnects and frees the context */
    redisFree(c);

    return 0;
}
