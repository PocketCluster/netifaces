# netifaces
Portable network interface information - Golang port of Python [netifaces 0.10.5](https://pypi.python.org/pypi/netifaces) 

> 1. What is this?
> 
> It’s been annoying me for some time that there’s no easy way to get the address(es) of the machine’s network interfaces from Python. There is a good reason for this difficulty, which is that it is virtually impossible to do so in a portable manner. However, it seems to me that there should be a package you can easy_install that will take care of working out the details of doing so on the machine you’re using, then you can get on with writing Python code without concerning yourself with the nitty gritty of system-dependent low-level networking APIs.

### Task

- [ ] inteface list
- [ ] addresses
- [ ] netmasks
- [x] gateways : IP6 portion `AF_INET6` is unconverted as it will be needed later. It's much better to handle `Py_*` function exception when needed.

### Native Tests

- OSX  
  run xcode
- Linux 

  ```sh
  cd xcode/netifaces
  gcc  ../../netifaces.c ./main.c
  ./a.out
  ```
  
### GO test

```sh
go test ./...
```
