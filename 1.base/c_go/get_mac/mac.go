package main

// #include <stdio.h>
// #include <stdlib.h>
//#include <string.h>
//#include <sys/ioctl.h>
//#include <net/if.h>
//#include <ctype.h>
//#include <time.h>
//#include <unistd.h>
//#include <sys/stat.h>
//#include <sys/types.h>
/*
void print(char *str) {
    printf("%s\n", str);
}
int  get_local_mac(char *mac, int len)
{
   int fd;
   int interfaceNum = 0;
   struct ifreq buf[16];
   struct ifconf ifc;
   struct ifreq ifrcopy;

   memset(mac, 0, len);
   if ((fd = socket(AF_INET, SOCK_DGRAM, 0)) < 0){
       perror("socket");
       return -1;
   }

   ifc.ifc_len = sizeof(buf);
   ifc.ifc_buf = (caddr_t)buf;
   if (!ioctl(fd, SIOCGIFCONF, (char *)&ifc)) {
       interfaceNum = ifc.ifc_len / sizeof(struct ifreq);
       printf("interface num = %d\n", interfaceNum);
       while (interfaceNum-- > 0) {
           printf("\ndevice name: %s\n", buf[interfaceNum].ifr_name);
           //ignore the interface that not up or not runing
           ifrcopy = buf[interfaceNum];
           if (ioctl(fd, SIOCGIFFLAGS, &ifrcopy) == 0) {
               if((ifrcopy.ifr_flags & IFF_UP) &&
                       !(ifrcopy.ifr_flags & IFF_LOOPBACK)){
                   if (!ioctl(fd, SIOCGIFHWADDR, (char *)(&buf[interfaceNum]))){
                       memset(mac, 0, len);
                       snprintf(mac, len, "%02X-%02X-%02X-%02X-%02X-%02X",
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[0],
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[1],
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[2],
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[3],
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[4],
                               (unsigned char)buf[interfaceNum].ifr_hwaddr.sa_data[5]);
                       printf("device mac: %s\n", mac);
                       break;
                   }
               }
           }
       }
   }
   close(fd);
   return 0;
}
*/
import "C"

//import "C" 和上面的注释不能有空格,否则会出现编译错误
//could not determine kind of name for C.test
import (
	"errors"
	"fmt"
	"unsafe"
)

func GetLocalMac() (mac string, err error) {
	data := make([]byte, 32, 32)
	cs := C.CString(string(data))
	sp := C.get_local_mac(cs, 32)
	if sp != 0 {
		err = errors.New("get mac error")
		return
	}
	mac = C.GoStringN(cs, 32)
	C.free(unsafe.Pointer(cs))
	return
}
func main() {
	mac, err := GetLocalMac()
	if err != nil {
		panic(err)
	}
	fmt.Println("mac:", mac)
}
