Description
-----------
`go-ok` takes an input file which has a list of IP addresses and port and checks if
HTTP GET is successful (200) for the host:port. If yes, it prints the IP address and port on the console


Usage
-----
```
$ go get github.com/jayapriya90/go-ok

$ go-ok -fpath=<path to input file>
```

Sample Input/Output
-------------------
$ go-ok -fpath=input.txt  
Input: [216.58.195.78:80 151.101.193.69:80 172.217.3.17:80]  
HTTP GET is successful for  
216.58.195.78:80  
172.217.3.17:80







