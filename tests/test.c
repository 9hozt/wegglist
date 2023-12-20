#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>


void doublefree(){
    char *buffer;
    buffer  = malloc(1337);
    if(buffer){
        free(buffer);
    }

    free(buffer);
}

void fakedoublefree(){
    char *buffer = malloc(1337);
    if(buffer){
        free(buffer);
        return;
    }

    free(buffer);
}

void formatString(){
    char *buffer;
    buffer = malloc(0xff);
    read(0,buffer,sizeof(buffer));
    printf(buffer);
}

void uaf()
{
    char *buffer;
    buffer  = malloc(1337);
    if(buffer){
        free(buffer);
    }

    printf("%s", buffer);
}

int main(int argc, char *argv[])
{
    char *buffer;
    char stack[20];

    gets(stack);
    strcpy(stack,"a");
    printf("%s", buffer);
    
}
