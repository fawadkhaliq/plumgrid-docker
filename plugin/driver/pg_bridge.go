package driver

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	Log "github.com/Sirupsen/logrus"
)

func (d *driver) pgBridgeCreate(ID string) {
	cookieJar, _ := cookiejar.New(nil)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	url := "https://" + vif + "/0/login"
	Log.Infof("URL:> %s", url)

	var jsonStr = []byte(`{"userName":"` + username + `", "password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Cookies:", resp.Cookies())
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//== PUT call

	url2 := "https://" + vif + "/0/connectivity/domain/" + tenant + "/ne/bri" + ID
	fmt.Println("URL:>", url2)

	var jsonStr1 = []byte(`{
				"action": {
					"action1": {
						"action_text": "create_and_link_ifc(DYN_1)"
					}
				},
				"ifc": {},
				"mobility": "true",
				"ne_description": "PLUMgrid Bridge",
				"ne_dname": "bridge-` + ID[:10] + `",
				"ne_group": "Bridge",
				"ne_type": "bridge"
	}`)

	req2, err := http.NewRequest("PUT", url2, bytes.NewBuffer(jsonStr1))
	req2.Header.Set("Accept", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	resp2, err2 := client.Do(req2)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer resp2.Body.Close()

	fmt.Println("response Status:", resp2.Status)
	//fmt.Println("response Headers:", resp2.Header)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("response Body:", string(body2))

	//== PUT call

	url3 := "https://" + vif + "/0/connectivity/domain/" + tenant + "/rule_group/cnf" + ID
	fmt.Println("URL:>", url3)

	var jsonStr3 = []byte(`{
					"mark_disabled": false,
					"ne_dest": "/ne/bri` + ID + `/action/action1",
					"ne_dname": "cnf-vmgroup-` + ID[10:] + `",
					"ne_type": "cnf-vmgroup",
					"rule": {
						"rules` + ID + `": {
							"add_context": "",
							"criteria": "pgtag2",
							"match": "bridge-` + ID[:10] + `"
						}
					}
	}`)

	req3, err := http.NewRequest("PUT", url3, bytes.NewBuffer(jsonStr3))
	req3.Header.Set("Accept", "application/json")
	req3.Header.Set("Content-Type", "application/json")

	resp3, err3 := client.Do(req3)
	if err3 != nil {
		fmt.Println(err3)
	}
	defer resp3.Body.Close()

	fmt.Println("response Status:", resp3.Status)
	//fmt.Println("response Headers:", resp3.Header)
	body3, _ := ioutil.ReadAll(resp3.Body)
	fmt.Println("response Body:", string(body3))

}

func (d *driver) pgBridgeDestroy(ID string) {
	cookieJar, _ := cookiejar.New(nil)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	url := "https://" + vif + "/0/login"
	Log.Infof("URL:> %s", url)

	var jsonStr = []byte(`{"userName":"` + username + `", "password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Cookies:", resp.Cookies())
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//== Delete call

	url2 := "https://" + vif + "/0/connectivity/domain/" + tenant + "/ne/bri" + ID
	fmt.Println("URL:>", url2)

	req2, err := http.NewRequest("DELETE", url2, nil)
	req2.Header.Set("Accept", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	resp2, err2 := client.Do(req2)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer resp2.Body.Close()

	fmt.Println("response Status:", resp2.Status)
	//fmt.Println("response Headers:", resp2.Header)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println("response Body:", string(body2))

	//== DELETE call

	url3 := "https://" + vif + "/0/connectivity/domain/" + tenant + "/rule_group/cnf" + ID
	fmt.Println("URL:>", url3)

	req3, err := http.NewRequest("DELETE", url3, nil)
	req3.Header.Set("Accept", "application/json")
	req3.Header.Set("Content-Type", "application/json")

	resp3, err3 := client.Do(req3)
	if err3 != nil {
		fmt.Println(err3)
	}
	defer resp3.Body.Close()

	fmt.Println("response Status:", resp3.Status)
	//fmt.Println("response Headers:", resp3.Header)
	body3, _ := ioutil.ReadAll(resp3.Body)
	fmt.Println("response Body:", string(body3))

}
