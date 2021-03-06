package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ipipdotnet/ipdb-go"
)

//KEYS : result keys.
var KEYS = []string{
	"#ip",          // raw ip
	"country_name", //国家名字
	"region_name",  //省名字
	"city_name",    //城市名字
	"owner_domain", //所有者
	"isp_domain",   //运营商
	//"latitude",         //纬度
	//"longitude",        //经度
	//"timezone",         //时区
	//"utc_offset",       //UTC时区
	"china_admin_code", //中国行政区划代码
	//"idd_code",         //国家电话号码前缀
	"country_code",   //国家2位代码
	"continent_code", //大洲代码
	"idc",            //IDC
	"base_station",   //基站
	"country_code3",  //国家3位代码
	//"european_union",   //是否为欧盟成员国
	//"currency_code",    //当前国家货币代码
	//"currency_name",    //当前国家货币名称
	"anycast", //ANYCAST
}
var db *ipdb.City

func findIPs(text string) map[string]bool {
	re := regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`)
	re.Longest()
	ips := re.FindAllString(text, -1)
	//fmt.Printf("%q\n", ips)
	ret := make(map[string]bool)
	for _, ip := range ips {
		ret[ip] = true
	}
	return ret
}

func ipRet(ip string, loc map[string]string) []string {
	vals := make([]string, 0, len(KEYS))
	vals = append(vals, ip)
	for i, k := range KEYS {
		if i == 0 {
			continue
		}
		val, ok := loc[k]
		if !ok || val == "" {
			val = "-"
		}

		vals = append(vals, val)
	}
	return vals
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	ips := findIPs(string(body))
	fmt.Fprintf(w, strings.Join(KEYS, " ")+"\n")
	for ip := range ips {
		if loc, err := db.FindMap(ip, "CN"); err != nil {
			continue
		} else {
			vals := ipRet(ip, loc)
			fmt.Fprintf(w, strings.Join(vals, " ")+"\n")
		}
	}

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ips := findIPs(string(body))
	resp := make(map[string]interface{})

	for ip := range ips {
		if loc, err := db.FindMap(ip, "CN"); err != nil {
			continue
		} else {
			ipInfo := make(map[string]interface{})

			for key, val := range loc {
				if key == "asn_info" {
					var asnInfo interface{}
					json.Unmarshal([]byte(val), &asnInfo)
					ipInfo[key] = asnInfo
					continue
				}
				ipInfo[key] = val
			}
			resp[ip] = ipInfo
		}
	}
	respJSON, _ := json.MarshalIndent(resp, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func main() {
	f := flag.String("f", "./ipv4_cn.ipdb", "ip data file")
	flag.Parse()
	var err error
	db, err = ipdb.NewCity(*f)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/json", jsonHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
