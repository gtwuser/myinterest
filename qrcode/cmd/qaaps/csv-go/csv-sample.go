package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	js := `{
    "status": "success",
    "transactionId": "070c7c94-8e5c-4c92-9b00-f07ab23e74c9_1607016320226",
    "recordCount": 1,
 "report": [
       {
           "Id": 0,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-04-03T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 1,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-11T16:23:25.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 2,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-18T16:29:02.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 3,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-27T21:19:28.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 4,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-02T06:31:55.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 5,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-20T23:31:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 6,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-05T18:44:44.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 7,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-05-16T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 8,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-30T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 9,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-11T18:59:53.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 10,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-22T21:19:29.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 11,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-25T20:10:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 12,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-19T07:59:29.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 13,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-04T14:55:04.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 14,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-16T21:36:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 15,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-27T11:48:07.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 16,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-25T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 17,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-21T08:21:01.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 18,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-19T23:32:03.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 19,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-02T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 20,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-17T16:53:16.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 21,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-30T12:37:17.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 22,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-09T06:33:45.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 23,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2019-12-20T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 24,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-20T22:57:36.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 25,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-29T01:35:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 26,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-26T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 27,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-14T20:57:41.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 28,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-27T23:38:48.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 29,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2019-12-09T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 30,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2019-12-12T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 31,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-06T10:10:20.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 32,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-04-27T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 33,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-03-03T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 34,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-03-27T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 35,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-03T17:14:45.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 36,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-16T17:12:03.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 37,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-31T14:16:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 38,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-02-28T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 39,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-03-21T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 40,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-11T23:44:24.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 41,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-30T15:05:38.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 42,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-06T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 43,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-26T22:58:28.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 44,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-14T21:18:48.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 45,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-28T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 46,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-25T15:51:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 47,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-13T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 48,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-09T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 49,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-26T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 50,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-01-06T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 51,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-16T18:09:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 52,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-03T11:54:48.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 53,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-12T19:40:45.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 54,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-02T18:45:47.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 55,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-19T23:27:56.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 56,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-03-09T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 57,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-01T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 58,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-01-20T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 59,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-02-20T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 60,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-01T16:10:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 61,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-30T07:28:37.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 62,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-07T08:22:24.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 63,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-07T00:57:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 64,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-19T07:04:06.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 65,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-20T07:06:21.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 66,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-12T16:37:23.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 67,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-02T15:38:49.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 68,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-04T16:01:36.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 69,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-07T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 70,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-13T18:32:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 71,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-02-16T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 72,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-22T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 73,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-10T13:27:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 74,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-04T22:04:10.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 75,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-12T14:56:28.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 76,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-03-13T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 77,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-20T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 78,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-04T14:16:57.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 79,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-21T01:18:34.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 80,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-02-14T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 81,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-17T13:08:53.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 82,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-04T12:30:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 83,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-10T15:50:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 84,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-02T21:25:15.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 85,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-04T23:45:55.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 86,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-16T20:56:42.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 87,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-23T18:17:37.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 88,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-04-29T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 89,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-09T14:49:47.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 90,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-29T20:03:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 91,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-05T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 92,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-03T17:42:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 93,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-15T14:25:43.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 94,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-14T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 95,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-29T14:58:33.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 96,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-16T07:05:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 97,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-13T07:06:46.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 98,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-09T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 99,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-03-23T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 100,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-16T10:53:51.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 101,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-25T03:59:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 102,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T07:04:29.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 103,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-16T20:59:41.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 104,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-26T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 105,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-02T19:20:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 106,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-31T07:03:50.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 107,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-15T14:31:06.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 108,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-14T22:20:35.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 109,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-19T20:23:41.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 110,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-26T21:10:35.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 111,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-26T12:24:07.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 112,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-01-13T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 113,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-11T07:03:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 114,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-17T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 115,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-02-19T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 116,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-04T16:35:53.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 117,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-22T07:27:40.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 118,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-05-11T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 119,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-10T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 120,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-24T07:54:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 121,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-11T16:57:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 122,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-07T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 123,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-08T22:12:18.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 124,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-04T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 125,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-03T13:06:34.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 126,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-18T00:57:15.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 127,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-15T16:39:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 128,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-26T07:13:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 129,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-13T01:27:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 130,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T11:23:06.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 131,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-14T18:56:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 132,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-07T19:09:46.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 133,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-02T15:52:19.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 134,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-10T08:12:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 135,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2019-12-16T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 136,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-30T08:29:58.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 137,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-21T14:06:06.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 138,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-26T19:49:52.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 139,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-01-24T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 140,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-01-25T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 141,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-30T14:43:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 142,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-05T09:00:51.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 143,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-10T02:00:42.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 144,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-04T14:40:51.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 145,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-13T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 146,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2019-12-21T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 147,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-05T16:26:40.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 148,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-22T18:57:22.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 149,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-28T03:30:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 150,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-12T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 151,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-05T04:16:25.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 152,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-16T14:12:02.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 153,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-21T19:28:11.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 154,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-20T15:59:03.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 155,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-18T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 156,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-30T23:28:25.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 157,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-23T18:30:38.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 158,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-06T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 159,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-10T19:02:13.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 160,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-01T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 161,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 6,
           "dateOfLastLogin": "2020-03-16T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 162,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-01-16T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 163,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-18T20:04:55.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 164,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-04-13T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 165,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T21:06:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 166,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T12:09:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 167,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-24T13:51:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 168,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-17T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 169,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-03T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 170,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-25T21:19:30.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 171,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-01T17:36:40.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 172,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-01T16:24:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 173,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-20T16:15:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 174,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-12T15:23:22.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 175,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-04-12T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 176,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-18T17:32:45.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 177,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-25T16:02:17.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 178,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-18T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 179,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-30T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 180,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-18T11:10:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 181,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-03T21:28:21.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 182,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-09T17:49:08.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 183,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-19T20:27:48.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 184,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-19T07:22:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 185,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-24T09:28:42.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 186,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-18T17:42:19.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 187,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-02-24T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 188,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-03T19:02:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 189,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-03T16:10:51.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 190,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-11T17:52:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 191,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-28T13:34:22.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 192,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-26T09:52:34.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 193,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-01T20:43:37.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 194,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-13T18:20:29.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 195,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-02-17T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 196,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-02-10T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 197,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 4,
           "dateOfLastLogin": "2020-05-04T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 198,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-05T23:12:16.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 199,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-30T00:12:41.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 200,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-04T16:29:34.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 201,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-01-31T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 202,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-24T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 203,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-10T21:13:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 204,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-14T07:03:16.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 205,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-01T04:49:15.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 206,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-19T22:43:02.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 207,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-03-12T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 208,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-18T07:05:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 209,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-05T04:40:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 210,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-29T21:20:40.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 211,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-26T16:26:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 212,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T20:18:27.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 213,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-10T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 214,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-30T18:36:27.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 215,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-14T13:46:25.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 216,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-29T14:19:10.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 217,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-02T10:02:20.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 218,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-03T23:33:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 219,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-11T04:06:53.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 220,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-14T13:11:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 221,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-30T07:13:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 222,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-20T07:04:09.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 223,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-05T11:28:22.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 224,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-31T13:01:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 225,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-25T09:00:03.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 226,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-15T14:19:59.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 227,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-08T21:36:24.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 228,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-18T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 229,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-05-10T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 230,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-17T21:20:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 231,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-17T07:10:40.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 232,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-14T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 233,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-17T10:31:20.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 234,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-03T13:38:02.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 235,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-31T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 236,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-03-01T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 237,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-02-02T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 238,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-26T23:28:24.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 239,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-13T14:59:52.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 240,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-14T07:34:36.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 241,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-23T23:27:46.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 242,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-05T12:28:57.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 243,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-02T21:35:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 244,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-02-21T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 245,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-29T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 246,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-30T02:53:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 247,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-07T17:09:19.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 248,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-05T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 249,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-09T11:26:47.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 250,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-02-03T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 251,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-01-27T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 252,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 6,
           "dateOfLastLogin": "2020-04-20T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 253,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-02-22T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 254,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-21T16:21:30.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 255,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-26T14:19:28.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 256,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-04T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 257,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-28T16:48:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 258,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-22T19:11:35.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 259,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 5,
           "dateOfLastLogin": "2020-03-30T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 260,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-11T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 261,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 6,
           "dateOfLastLogin": "2020-03-02T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 262,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-08T18:42:14.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 263,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-12T07:04:41.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 264,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 3,
           "dateOfLastLogin": "2020-04-06T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 265,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-01-03T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 266,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-24T19:17:54.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 267,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-04T18:37:10.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 268,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-06-28T06:58:48.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 269,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-25T15:17:46.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 270,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-04T14:00:27.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 271,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2020-02-09T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 272,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-17T16:29:15.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 273,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-04-04T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 274,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-04T22:34:21.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 275,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-08-27T08:29:31.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 276,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-03-25T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 277,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-02-07T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 278,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-14T15:59:58.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 279,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-05T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 280,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 6,
           "dateOfLastLogin": "2020-05-18T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 281,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-05-29T19:01:26.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 282,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 2,
           "dateOfLastLogin": "2019-12-07T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 283,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2019-12-23T00:00:00.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 284,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-08T19:33:27.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 285,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-07-08T18:48:20.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 286,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-10-10T10:21:30.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       },
       {
           "Id": 287,
           "cxCustomerBuId": 838757,
           "cxCustomerBuName": "AUTOLIV INC SE",
           "cxCustomerId": 642241,
           "countOfPortalLoginsOnADate": 1,
           "dateOfLastLogin": "2020-09-20T10:06:12.000+0000",
           "ampguid": "eu-6061",
           "cxCustomerName": "AUTOLIV INC"
       }
   ]
}`
//
//		js := `{
//    "status": "success",
//    "transactionId": "514c3dea-1059-4122-82e3-be7ef853d61d_1607016343703",
//    "recordCount": 1,
//	   "report": [
//	       {
//	           "Id": 0,
//	           "ipBlocklist": 1,
//	           "telemetryReportedDate": "2020-10-19",
//	           "simpleCustomerDetectionLists": 3,
//	           "totalUser": 36,
//	           "tetraConnectorCount": 771,
//	           "fileFetchEnabled": "true",
//	           "usedApi": "TRUE",
//	           "cxCustomerBuId": 838757,
//	           "advancedCustomerDetectionLists": 2,
//	           "cxCustomerId": 642241,
//	           "autoAnalysisForLowPrevalence": "FALSE",
//	           "ipWhitelist": 1,
//	           "licenseEndDate": "2022-05-14",
//	           "exploitPreventionConnectorCount": 143,
//	           "nonDefaultGroups": 10,
//	           "companyName": "Autoliv",
//	           "ampguid": "eu-6061",
//	           "blockinglist": 2,
//	           "whitelist": 1,
//	           "cxCustomerName": "AUTOLIV INC",
//	           "dcConfigured": "FALSE",
//	           "exclusionsList": 6,
//	           "ctaEnabled": "TRUE",
//	           "nonDefaultPolicies": 8,
//	           "adminCount": 14,
//	           "mapConnectorCount": 21048,
//	           "userCount": 22,
//	           "cxCustomerBuName": "AUTOLIV INC SE",
//	           "licenseStartDate": "2019-05-15",
//	           "region": "eu",
//	           "ownTgApiKey": "FALSE"
//	       }
//	   ]
//	}`
	recordFile, err := os.Create("./superheroes.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	// 2. Initialize the writer
	writer := csv.NewWriter(recordFile)
	defer recordFile.Close()
	headers := make([]string, 0)
	anyJSON := make(map[string]interface{})
	_ = json.Unmarshal([]byte(js), &anyJSON)
	//fmt.Print(m)
	fmt.Println("First level ", reflect.TypeOf(anyJSON))
	for key, val := range anyJSON {
		fmt.Println("key:", key, "val", val)
		if key != "report"{
			continue
		}
		fmt.Println("Second level ", reflect.TypeOf(val))
		ma := val.([]interface{})
		for i, me := range ma {
			data := make([]string, 0)
			fmt.Println("index ", i, ":", me)
			fmt.Println("-----------------------------")
			fm := me.(map[string]interface{})
			if len(headers) == 0 {
				fmt.Println("headers are not set yet !!!")
				for k, v := range fm {
					fmt.Println(k, ":", v)
					//if strings.EqualFold("cxCustomerBuName", k){
					//	continue
					//}
					//if len(headers) <= len(fm) {
					fmt.Println("len of map ", len(fm), "len of headers ", len(headers))
					headers = append(headers, k)
					//}
				}
				writer.Write(headers)
				//writer.Write(data)
				writer.Flush()
			}
			//else {
			//for k, v := range fm {
			//	fmt.Println(k, ":", v)
				for _, header := range headers {
					fv := fm[header]
					switch t := fv.(type) {
					case float32, float64:
						fv = fmt.Sprintf("%f", t)
					}
					data = append(data, fv.(string))
				}
			//}
			writer.Write(data)
			writer.Flush()
			//}
		}
	}
}
