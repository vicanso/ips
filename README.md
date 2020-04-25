# ips

[![Build Status](https://img.shields.io/travis/vicanso/ips.svg?label=linux+build)](https://travis-ci.org/vicanso/ips)

Check ip is exists in iplist.

## API

```go
iPs := ips.New()
iPs.Add("12.12.12.12")
iPs.Add("12.12.12.13")
iPs.Add("192.168.1.1/24")
fmt.Println(iPs.Contains("192.168.1.1"))
```

If you need to change the ip anytime, you shuld add `mutex`.

```go
iPs := ips.New()
iPs.Mutex = new(sync.RWMutex)
iPs.Add("12.12.12.12")
iPs.Add("12.12.12.13")
iPs.Add("192.168.1.1/24")
fmt.Println(iPs.Contains("192.168.1.1"))
```