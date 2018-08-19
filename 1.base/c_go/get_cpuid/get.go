package main

// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <sys/ioctl.h>
// #include <net/if.h>
// #include <ctype.h>
// #include <time.h>
// #include <unistd.h>
// #include <sys/stat.h>
// #include <sys/types.h>
/*
int get_cpu_id_by_asm(char *cpu_id, int len){
    int ret;
    unsigned int s1 = 0;
    unsigned int s2 = 0;
    memset(cpu_id, 0, sizeof(len));
    asm volatile (
            "movl $0x01, %%eax; \n\t"
            "xorl %%edx, %%edx; \n\t"
            "cpuid; \n\t"
            "movl %%edx, %0; \n\t"
            "movl %%eax, %1; \n\t"
            : "=m"(s1), "=m"(s2)
            );

    if (0 == s1 && 0 == s2){
        return -1;
    }

    char cpu[32] = { 0 };
    snprintf(cpu, sizeof(cpu), "%08X%08X", htonl(s2), htonl(s1));
    memcpy(cpu_id, cpu, 32);
    return 0;
}

static int parse_cpu_id(const char * file_name, const char * match_words, char * cpu_id, int len){
    int ret = -1;
    FILE * pf;
    char line[4096] = { 0 };
    memset(cpu_id, 0, len);
    pf = fopen(file_name, "r");
    if (!pf)
        return -1;
    while (fgets(line, sizeof(line), pf))
    {

        const char * cpu = strstr(line, match_words);
        if (NULL == cpu){
            continue;
        }
        ret = 0;
        cpu += strlen(match_words);
        while ('\0' != cpu[0]){
            if (' ' != cpu[0]) {
                *cpu_id++ = cpu[0];
            }
            ++cpu;
        }
        if(ret == 0)
            break;
    }
    fclose(pf);
    return ret;
}

static int  get_cpu_id_by_system(char *cpu_id, int len){
    int ret;
    memset(cpu_id, 0, len);
    const char * dmidecode_result = ".dmidecode_result.txt";
    char command[512] = { 0 };
    snprintf(command, sizeof(command), "dmidecode -t 4 | grep ID > %s", dmidecode_result);

    if (0 == system(command)){
        ret = parse_cpu_id(dmidecode_result, "ID:", cpu_id, len);
    }

    unlink(dmidecode_result);

    return ret;
}

int get_cpu_id() {
    char cpuid[32] = {0};
    #if 0
    if (get_cpu_id_by_asm(cpuid, sizeof(cpuid)) == 0){
        //memcpy(cpuid, id, sizeof(cpuid));
        printf("cpuid %s\r\n", cpuid);
    }
    #endif
    if (0 == getuid()){
        if (get_cpu_id_by_system(cpuid, sizeof(cpuid)) == 0){
           // memcpy(cpuid, id, sizeof(cpuid));
            printf("cpuid %s\r\n", cpuid);
        }
    }
    return 0;
}
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	//var id [32]char
	id := "hewewe"
	cs := C.CString(id)
	C.get_cpu_id()
	//C.print(cs)
	C.free(unsafe.Pointer(cs))
}
