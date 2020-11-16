package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TokenAccess struct {
	AccessToken string `json:"access_token"`
}

func main() {

	//	s := `{
	//  "portalAccounts": [
	//    {
	//      "customerId": "c1",
	//      "policies": [
	//        {
	//          "role": "admin",
	//          "assetGroups": [
	//            {
	//              "name": "agName1",
	//              "id": "ag1"
	//            },
	//            {
	//              "name": "agName3",
	//              "id": "ag3"
	//            }
	//          ]
	//        },
	//        {
	//          "role": "view",
	//          "assetGroups": [
	//            {
	//              "name": "agName2",
	//              "id": "ag2"
	//            }
	//          ]
	//        }
	//      ]
	//    }
	//  ]
	//}`
	s2 := `{
    "access_token": "OsGPz5F5JhRMXuR2qFmossfPrWZz",
    "token_type": "Bearer",
    "expires_in": 3599
}`

	t := &TokenAccess{}
	//cp := &denver.CPAccessPolicy{}
	//err := json.Unmarshal([]byte(s), cp)
	err := json.Unmarshal([]byte(s2), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	js, err := json.Marshal(t)
	fmt.Println(string(js))
}
