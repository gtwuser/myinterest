package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//url := flag.String(`url`, ``, `Url for data inspection`)
	//flag.Parse()
	//resp, err := http.Get(*url)
	//fmt.Println(`Url :`,*url)
	//if err != nil {
	//	log.Fatal("Url could not be reached due to err ", err)
	//}
	//defer resp.Body.Close()
	//
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal("Url ", *url, "output could not be read ", resp.StatusCode)
	//}

	//file, _ := ioutil.ReadFile(`./data.xml`)
	//content := string(file)
	//re := regexp.MustCompile(`<li.*?>.*?title="(.*?)"`)
	//allString := re.FindAllStringSubmatch(content, -1)
	//for i, r := range allString {
	//	fmt.Println(r[1], i)
	//}

	//fmt.Println("Done")
	//s := []string{"1808596"}
	//fmt.Println(strings.Join(s, `","`))

	//// TODO
	//file, _ := ioutil.ReadFile(`./data.xml`)
	//content := string(file)
	//re := regexp.MustCompile(`<li.*?>.*?title="(.*?)"`)
	//allString := re.FindAllStringSubmatch(content, -1)
	//for i, r := range allString {
	//	fmt.Println(r[1], i)
	//}

	//baseContentType := "application/grpc"
	//contentType := "application/grpc; application/json; charset=utf-8"
	//switch contentType[len(baseContentType)] {
	//case '+', ';':
	//	// this will return true for "application/grpc+" or "application/grpc;"
	//	// which the previous validContentType function tested to be valid, so we
	//	// just say that no content-subtype is specified in this case
	//	fmt.Println(contentType[len(baseContentType)+1:], true)
	//default:
	//	fmt.Println("", false)
	//}
	//const layout = "2006-01-02"
	//currDate := time.Now().Format(layout)
	//post := time.Now().AddDate(-1, 0, 0).Format(layout)
	//fmt.Println("Date", currDate, post)
	//str := []string{"Geeks", "For", "Geeks"}
	//
	//currentTime := time.Now()
	//
	//fmt.Println("Current Time in String: ", currentTime.String())
	//
	//fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	//fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	//// joining the string by separator
	//fmt.Println(strings.Join(str, `","`))
	//fmt.Println(contentType[len(baseContentType)+1:])
	//letters := [2][2]string{}
	//
	//// Assign all elements in 2 by 2 array.
	//letters[0][0] = "a"
	//letters[0][1] = "b"
	//letters[1][0] = "c"
	//letters[1][1] = "d"
	//
	//fmt.Println(letters[:])
	//var data = `{
	//"id": 12423434,
	//"Name": "Fernando"
	//}`
	//d := json.NewDecoder(strings.NewReader(data))
	//d.UseNumber()
	//var x interface{}
	//if err := d.Decode(&x); err != nil {
	//	logrus.Fatal(err)
	//}
	//fmt.Printf("decoded to %#v\n", x)
	//mm := x.(map[string]interface{})
	//fv := mm["id"]
	//switch t := fv.(type) {
	//case json.Number:
	//	fv = fmt.Sprintf("%v", t)
	//}
	//if fv == nil {
	//	fv = ""
	//}
	//var a float64 = 1234567689.9866
	//var fv interface{}
	//fv = a
	//switch t := fv.(type) {
	//case float32, float64:
	//	fv = fmt.Sprintf("%d", int(t.(float64)))
	//}
	//fmt.Println(fv)

	//bucket := os.Args[1]
	//filename := os.Args[2]
	//
	//file, err := os.Open(filename)
	//if err != nil {
	//	exitErrorf("Unable to open file %q, %v", err)
	//}
	//
	//defer file.Close()
	//sess, err := session.NewSession(&aws.Config{
	//	Region: aws.String("us-west-2")},
	//)
	//
	//uploader := s3manager.NewUploader(sess)
	//
	//_, err = uploader.Upload(&s3manager.UploadInput{
	//	Bucket: aws.String(bucket),
	//	Key:    aws.String(filename),
	//	Body:   file,
	//})
	//
	//logrus.Info("uploaded successfully ", err)

	//paths := []string{}
	//
	//filepath.Walk(filepath.Dir("."), func(path string, info os.FileInfo, err error) error {
	//	// We care only about files, not directories
	//	if !info.IsDir() {
	//		paths = append(paths, path)
	//	}
	//	return nil
	//})
	//
	//fmt.Println("Paths are ", paths)

	// SNS Stubs
	//// List topics
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))
	//
	//svc := sns.New(sess)
	//topics, err := svc.ListTopics(nil)
	//if err != nil {
	//	logrus.Fatal("Error fetching topics ", err)
	//}
	//
	//var kt string
	//for _, topic := range topics.Topics {
	//	arn := *topic.TopicArn
	//	if strings.Contains(arn, `kapjoshi`) {
	//		kt = arn
	//		logrus.Infof("Topic Arn %v", arn)
	//	}
	//}
	//
	//topic, err := svc.CreateTopic(&sns.CreateTopicInput{
	//	Name: aws.String(`kapjoshi-tt`),
	//})
	//if err != nil {
	//	logrus.Fatalf("Error while creating topic %v", err)
	//}
	//
	//logrus.Infof("Topic created %v", topic)
	//
	//resultSub, err := svc.Subscribe(&sns.SubscribeInput{
	//	Endpoint:              aws.String(`kapjoshi@cisco.com`),
	//	Protocol:              aws.String(`email`),
	//	ReturnSubscriptionArn: aws.Bool(true),
	//	TopicArn:              topic.TopicArn,
	//})
	//
	//if err != nil {
	//	logrus.Fatalf("Error while subscribing topic %v", err)
	//}
	//
	//logrus.Infof("Subscription details %v", *resultSub.SubscriptionArn)
	//
	//rs, err := svc.Subscribe(&sns.SubscribeInput{
	//	Endpoint:              aws.String(`kapil2kapil@gmail.com`),
	//	Protocol:              aws.String(`email`),
	//	ReturnSubscriptionArn: aws.Bool(true),
	//	TopicArn:              topic.TopicArn,
	//})
	//
	//if err != nil {
	//	logrus.Errorf("subscription issue %v", err)
	//}
	//
	//logrus.Infof("subscription details on the new gmail id %v", rs)
	//
	//topicList, err := svc.ListSubscriptionsByTopic(&sns.ListSubscriptionsByTopicInput{
	//	TopicArn: topic.TopicArn,
	//})
	//
	//if err != nil {
	//	logrus.Errorf("fetching topic arn's %v", err)
	//}
	//
	//for _, subscription := range topicList.Subscriptions {
	//	logrus.Infof("subscriptions found %v", *subscription.TopicArn)
	//}
	//
	//po, err := svc.Publish(&sns.PublishInput{
	//	Message:  aws.String(`hello new Lathikata!!`),
	//	TopicArn: aws.String(kt),
	//})
	//
	//if err != nil {
	//	logrus.Errorf("publishing message %v", err)
	//}
	//
	//logrus.Infof("published messages details %v", po.String())

	//}
	//	fmt.Println("Today is a weekday.")
	//default:
	//	fmt.Println("Today is Sunday.")
	//case time.Sunday:
	//	fmt.Println("Today is Saturday.")
	//case time.Saturday:
	//switch time.Now().Weekday() {
	//fmt.Print(result)
	//result := gjson.Get(j, "intersight_api_private_key")
	//}`
	//    "intersight_api_private_key": "oiasdoiasdoihasd"
	//	j := `{
	//s := "-----BEGIN RSA PRIVATE KEY----- oasdoiasoasd -----END RSA PRIVATE KEY-----"
	//s = base64.StdEncoding.EncodeToString([]byte(s))
	//validatePEM(s)
	//
	//const json = `{"name":{"first":"Tom", "last": "Anderson"},"age": 37,"children": ["Sara", "Alex", "Jack"],"fav.movie": "Dear Hunter","friends": [{"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}] }`
	//fmt.Print(gjson.Get(json, "@pretty"))
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "intersight")
	fmt.Println(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, _ := ioutil.ReadDir(dir)
	fmt.Println(len(files))
	defer os.RemoveAll(dir) // clean up
	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
	files, _ = ioutil.ReadDir(dir)
	fmt.Println(len(files))

}

//func validatePEM(s string) {
//	log.Info().Str("demo", "sample").Msg("Hi")
//	decStr, _ := base64.StdEncoding.DecodeString(s)
//	decode, _ := pem.Decode(decStr)
//	if decode != nil {
//		fmt.Printf("yes\n")
//	} else {
//		fmt.Printf("no\n")
//	}
//}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
