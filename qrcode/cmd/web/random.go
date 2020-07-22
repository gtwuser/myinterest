package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)



type CPToken struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}



func main() {
	clientId := `cxcp-nonprod-api`
	clientSec := `cxcp-nonprod-api$$`
	tokenUrl := `https://cloudsso-test.cisco.com/as/token.oauth2`
	//cpAccessPolicyUrl := `https://api-dev1.us.csco.cloud/controlpoint/ctrlpt/v1/access-policy-discovery/pmagi`

	tokReq, _ := http.NewRequest(`POST`, tokenUrl, nil)
	q := tokReq.URL.Query()
	q.Add("grant_type", "client_credentials")
	q.Add("client_id", clientId)
	q.Add("client_secret", clientSec)
	tokReq.URL.RawQuery = q.Encode()
	fmt.Println(`URL 1 `, tokReq.URL.RawQuery)
	fmt.Println(`URL 2 `, tokReq.URL.String())
	// Send req using http Client
	client := &http.Client{}

	start := time.Now()
	fmt.Println("Token URL ", tokenUrl)
	fmt.Println("Token URL ", tokReq.URL.String())
	log.Printf("CP token request to %v\n", tokReq.Host)
	resp, err := client.Do(tokReq)
	if err != nil {
		log.Printf("Error while fetching token from %v, %v", tokReq.Host, err)
		//return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Failed to close response body %v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body %v\n", err)
		//return nil, err
	}

	elapsed := time.Now().Sub(start)

	log.Printf("Time took for getting token from %v, %v\n", tokReq.Host, elapsed)

	if resp.StatusCode != http.StatusOK {
		stringBody := string(body)
		log.Printf("Error getting token from CP code=%d  response=%v\n", resp.StatusCode, stringBody)
		//return nil, fmt.Errorf("Error getting token from CP code=%d  response=%v", resp.StatusCode, stringBody)
	}

	var cpToken CPToken
	err = json.Unmarshal(body, &cpToken)
	if err != nil {
		log.Printf("Error marshalling token %v\n", err)
		//return nil, err
	}

	log.Printf(`Token recieved is :%v`, cpToken.AccessToken)
}
