# ips

[![Build Status](https://github.com/vicanso/ips/workflows/Test/badge.svg)](https://github.com/vicanso/ips/actions)

Check ip is exists in iplist.

## API

```go
ips := ips.NewWithoutMutex()
ips.Add("12.12.12.12", "12.12.12.13", "192.168.1.1/24")
fmt.Println(ips.Contains("192.168.1.1"))
ips.Replace("1.1.1.1")
ips.Reset()
```

If you need to change the ip anytime, you shuuld add `mutex`.

```go
ips := ips.New()
ips.Add("12.12.12.12")
ips.Add("12.12.12.13")
ips.Add("192.168.1.1/24")
fmt.Println(ips.Contains("192.168.1.1"))
ips.Replace("1.1.1.1")
ips.Reset()
```