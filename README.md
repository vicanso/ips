# ips

Check ip is exists in iplist.

## API

```go
iPs := ips.New()
iPs.Add("12.12.12.12")
iPs.Add("12.12.12.13")
iPs.Add("192.168.1.1/24")
exists := iPs.Exists("192.168.1.1")
fmt.Println(exists)
```