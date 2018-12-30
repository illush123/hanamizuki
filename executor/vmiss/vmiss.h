#ifndef vmiss_h
#define vmiss_h

#include <stdio.h>
#include "interpreter.h"

int wait_client(void);
int compare(int fd, VM* vm);

#endif /* vmiss_h */
