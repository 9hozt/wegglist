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

void fileTest(){
    FILE * f;
    f = fopen("/tmp/test", "a")
    fprintf( f, "Appending data..." );
    fclose(f);
}

void unclosed()
{
    FILE * f;
    f = fopen("/tmp/test", "a")
    if (f == NULL){
        return 0;
    }
    fprintf( f, "Appending data..." );
}

void loopTest()
{
    char b[8];
    char t[90];
    read(0, &t, 90);
    for(int i=0;i<100;i++)
    {
        b[i] = t[i];
    }
}


int main(int argc, char *argv[])
{
    char *buffer;
    char stack[20];

    gets(stack);
    strcpy(stack,"a");
    printf("%s", buffer);
    
}
