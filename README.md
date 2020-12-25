# ipserv


## build

```
go build ipserv.go

./ipserv -f ./ipv4_cn.ipdb
```

## Usage

使用ipip的ipdb数据库，提供了几个http的查询接口

* /  返回text格式结果，仅有部分字段
* /json 返回json结果，全部字段

使用方式：
将任意含有IP的文本，用POST方式发送到上述接口，即可返回对应IP的查询结果。


```
wxh-rmbp » echo "1.1.1.1 2.2.2.2 123.123.1.1 1.1.1.1 255.255.255.255 8.8.8.26 114.114.114.234----" | curl 127.1:8080/ -d@- -H 'Content-Type: text/plain'
#ip country_name region_name city_name owner_domain isp_domain china_admin_code country_code continent_code idc base_station country_code3 anycast
1.1.1.1 CLOUDFLARE.COM CLOUDFLARE.COM - apnic.net - - - - IDC - - ANYCAST
2.2.2.2 法国 法国 - - orange.com - FR EU - - FRA -
123.123.1.1 中国 北京 北京 - 联通 110000 CN AP - - CHN -
255.255.255.255 IPIP.NET 2020121605 - - - - - - - - - -
8.8.8.26 GOOGLE.COM GOOGLE.COM - google.com level3.com - - - IDC - - ANYCAST
114.114.114.234 114DNS.COM 114DNS.COM - greatbit.com - - - - IDC - - ANYCAST
```

json格式：
```
wxh-rmbp » echo "1.1.1.1 test 114.114.114.114" | curl 127.1:8080/json -d@-
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


也可增加几个alias，方便shell中使用
```
alias ipip-json="curl -s 127.0.0.1:8080/json -d@-"
alias ipip="curl -H 'Content-Type: text/plain' -s 127.0.0.1:8080/ -d@- "
```

```
wxh-rmbp » echo 1.2.3.4 | ipip
#ip country_name region_name city_name owner_domain isp_domain china_admin_code country_code continent_code idc base_station country_code3 anycast
1.2.3.4 APNIC.NET APNIC.NET - apnic.net - - - - IDC - - ANYCAST


wxh-rmbp » echo 1.2.3.4 | ipip-json
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

# 查询剪切板中的IP
```
wxh-rmbp » pbpaste | ipip
....
wxh-rmbp » pbpaste | ipip-json
....
```