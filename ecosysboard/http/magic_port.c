#include <sys/types.h>
#include <sys/socket.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <errno.h>
#include <netinet/in.h>
#include "magic_port.h"

int get_first_open_port()
{
    int sock = socket(AF_INET, SOCK_STREAM, 0);
    if(sock < 0) {
      printf("socket error\n");
      return -1;
     }

    struct sockaddr_in serv_addr;
    bzero((char *) &serv_addr, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_addr.s_addr = INADDR_ANY;
    serv_addr.sin_port = 0;
    if (bind(sock, (struct sockaddr *) &serv_addr, sizeof(serv_addr)) < 0) {
      if(errno == EADDRINUSE) {
        printf("the port is not available. already to other process\n");
        return -1;
      } else {
        printf("could not bind to process (%d) %s\n", errno, strerror(errno));
        return -1;
      }
     }

    socklen_t len = sizeof(serv_addr);
    if (getsockname(sock, (struct sockaddr *)&serv_addr, &len) == -1) {
      perror("getsockname");
      return -1;
     }

    if (close (sock) < 0 ) {
      printf("did not close: %s\n", strerror(errno));
      return -1;
     }
     return ntohs(serv_addr.sin_port);
}