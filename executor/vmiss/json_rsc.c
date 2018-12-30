#include "json_rsc.h"

void instruction_json(char* allstr, char *opcode){
    char str[64] = "";
    sprintf(str,"\"instruction\": \"%s\"",opcode);
    strcat(allstr,str);
}

void reg_json(char* allstr, char* reg_name, unsigned short vmiss,unsigned short student){
    char str[64] = "";
    sprintf(str, "\"%s\":{ \"vmiss\": %d,\"student\": %d},",reg_name,vmiss,student);
    strcat(allstr,str);
}

void flg_json(char* allstr, char* flg_name, char vmiss, char student){
    char str[64] ="";
    sprintf(str, "\"%s\":{ \"vmiss\": %d,\"student\": %d},",flg_name,vmiss,student);
    strcat(allstr,str);
}

void dmem_json(char* allstr, unsigned short addr,unsigned char vmiss, unsigned char student){
    char str[64] ="";
    sprintf(str,"\"memory[%04x]\":{ \"vmiss\": %d,\"student\": %d},",addr,vmiss,student);
    strcat(allstr,str);
}
