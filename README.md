# ipserv, An IP lookup API


## Build

```
$ go build

$ ./ipserv -h
Usage of ./ipserv:
  -f string
        ip data file (default "./ipv4_cn.ipdb")
  -ip string
        IP to use, default 0.0.0.0 (default "0.0.0.0")
  -port int
        Port to use, default 9527 (default 9527)

$ ./ipserv -f ./ipv4_cn.ipdb
```

## Usage

You can get ip database from [ipip](https://www.ipip.net/product/ip.html) , then use this tool to run some lookup api servce.

* /  txt response, only some IP attribute 
* /json json response, full IP attribute

Usercase:
Any text content that has IP in it, you can use HTTP POST method to send the content to this API, it will response you with the IP infomation.

# Bash alias, recommended

You can add some bash alias，in case you want to use it in shell
```
alias ipip-json="curl -s 127.0.0.1:9527/json -d@-"
alias ipip="curl -H 'Content-Type: text/plain' -s 127.0.0.1:9527/ -d@- "
```

Then run:
```
pc » echo 1.2.3.4 | ipip
#ip country_name region_name city_name owner_domain isp_domain china_admin_code country_code continent_code idc base_station country_code3 anycast
1.2.3.4 APNIC.NET APNIC.NET - apnic.net - - - - IDC - - ANYCAST


pc » echo 1.2.3.4 | ipip-json
{
    "1.2.3.4": {
        "anycast": "ANYCAST",
        "asn_info": [],
        "base_station": "",
        "china_admin_code": "",
        "city_name": "",
        "continent_code": "",
        "country_code": "",
        "country_code3": "",
        "country_name": "APNIC.NET",
        "currency_code": "",
        "currency_name": "",
        "european_union": "0",
        "idc": "IDC",
        "idd_code": "",
        "isp_domain": "",
        "latitude": "",
        "longitude": "",
        "owner_domain": "apnic.net",
        "region_name": "APNIC.NET",
        "timezone": "",
        "utc_offset": ""
    }
}
```

# Or lookup the IP in clickboard
```
pc » pbpaste | ipip
....
pc » pbpaste | ipip-json
....
```


# Or just use the raw API:
```
pc » echo "1.1.1.1 2.2.2.2 123.123.1.1 1.1.1.1 255.255.255.255 8.8.8.26 114.114.114.234----" | curl 127.1:9527/ -d@- -H 'Content-Type: text/plain'
#ip country_name region_name city_name owner_domain isp_domain china_admin_code country_code continent_code idc base_station country_code3 anycast
1.1.1.1 CLOUDFLARE.COM CLOUDFLARE.COM - apnic.net - - - - IDC - - ANYCAST
2.2.2.2 法国 法国 - - orange.com - FR EU - - FRA -
123.123.1.1 中国 北京 北京 - 联通 110000 CN AP - - CHN -
255.255.255.255 IPIP.NET 2020121605 - - - - - - - - - -
8.8.8.26 GOOGLE.COM GOOGLE.COM - google.com level3.com - - - IDC - - ANYCAST
114.114.114.234 114DNS.COM 114DNS.COM - greatbit.com - - - - IDC - - ANYCAST
```

json response：
```
pc » echo "1.1.1.1 test 114.114.114.114" | curl 127.1:9527/json -d@-
{
    "1.1.1.1": {
        "anycast": "ANYCAST",
        "asn_info": [
            {
                "asn": 13335,
                "cc": "US",
                "net": "CLOUDFLARENET",
                "org": "Cloudflare, Inc.",
                "reg": "arin"
            }
        ],
        "base_station": "",
        "china_admin_code": "",
        "city_name": "",
        "continent_code": "",
        "country_code": "",
        "country_code3": "",
        "country_name": "CLOUDFLARE.COM",
        "currency_code": "",
        "currency_name": "",
        "european_union": "0",
        "idc": "IDC",
        "idd_code": "",
        "isp_domain": "",
        "latitude": "",
        "longitude": "",
        "owner_domain": "apnic.net",
        "region_name": "CLOUDFLARE.COM",
        "timezone": "",
        "utc_offset": ""
    },
    "114.114.114.114": {
        "anycast": "ANYCAST",
        ...
    }
}%
```



