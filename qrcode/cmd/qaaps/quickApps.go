package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
	"os"
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

	// SQS Stubs
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	qL, err := svc.ListQueues(nil)

	if err != nil {
		logrus.Errorf("getting queue urls %v", err)
	}

	for _, url := range qL.QueueUrls {
		logrus.Info("Queue urls ", *url)
	}

	qc, err := svc.CreateQueue(&sqs.CreateQueueInput{
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
		QueueName: aws.String(`kapjoshi-qt`),
	})

	if err != nil {
		logrus.Errorf("creating queue %v", err)
	}

	logrus.Info("new queue created ", *qc.QueueUrl)

	// queue url with prefix
	qL, err = svc.ListQueues(&sqs.ListQueuesInput{
		QueueNamePrefix: aws.String(`kapjoshi`),
	})

	if err != nil {
		logrus.Errorf("listing created queue %v", err)
	}
	logrus.Info("Queue created with url ", *qL.QueueUrls[0])

	fQL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(`kapjoshi-qt`),
	})

	msg, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": {
				DataType:    aws.String(`String`),
				StringValue: aws.String(`yesterdays activities`),
			},
		},
		MessageBody: aws.String(`Been to old town yesterday`),
		QueueUrl:    fQL.QueueUrl,
	})

	if err != nil {
		logrus.Errorf("sending message %v", err)
	}

	logrus.Infof("Message sent output %v", msg)

	rm, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: fQL.QueueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
	})

	if err != nil {
		logrus.Errorf("recieving message %v", err)
	}

	logrus.Infof("recieve messages %v ", rm.Messages)
	logrus.Infof("recieve message details %v ", rm)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
