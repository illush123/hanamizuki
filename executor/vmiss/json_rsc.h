#ifndef json_rsc_h
#define json_rsc_h

#include <stdio.h>
void instruction_json(char* allstr, char *opcode);
void reg_json(char* allstr, char* reg_name, unsigned short vmiss,unsigned short student);
void flg_json(char* allstr, char* flg_name, char vmiss, char student);
void dmem_json(char* allstr, unsigned short addr,unsigned char vmiss, unsigned char student);
#endif /* json_rsc_h */
