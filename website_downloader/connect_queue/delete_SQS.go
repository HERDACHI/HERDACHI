package connect_queue

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// snippet-end:[sqs.go.delete_message.imports]

// GetQueueURL gets the URL of an Amazon SQS queue
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     queueName is the name of the queue
// Output:
//     If success, the URL of the queue and nil
//     Otherwise, an empty string and an error from the call to
func GetQueueURL2(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteMessage deletes a message from an Amazon SQS queue
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     queueURL is the URL of the queue
//     messageID is the ID of the message
// Output:
//     If success, nil
//     Otherwise, an error from the call to DeleteMessage
func DeleteMessage2(sess *session.Session, queueURL *string, messageHandle *string) error {

	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})

	if err != nil {
		return err
	}

	return nil
}

func main2() {

	// Example to execute (Rankleads in the SQS name)
	// go run delete_SQS.go -q Rankleads -m AQEBtSnrKnIpzET2KZKQDl9zV3mz4C9BZ7xFWFd95gOwjoTvXOiWTFktKfffhesztG50DMddwfZ9k135b0XfkXB5yUw7W9o1BUt2i9AWYqA0uYOE7mdJyt9+DM/h9UlF93HXQGPYheX/o66OINXijv8ZbbPq/5OMKhx9fW5s9x71k73wvaeOKVEz41oGN1F9rCUCQc0OFC6Ycb59A3JImeW4Luo6kzP66VJ53MBlqIridQj7etEUsImRRxXa4gIrib8/F3jO9XPxwgFjjpeE++RkijJdSIZGumuO6WEbSOlz0zKjf810d6vqDYJ19JpSC0WUA1M7fUmQhMsKku9pBFZ7F2DXjg/gjSPFHZ/4mQynvsrdSDsQF6+3Ua54sJ+8qFQtM/Ge1b2ZhEPSv5OR5Fmy5A==
	//
	//note:
	//-q is the queue name
	//-m is the queue messageHandle

	queue := flag.String("q", "", "The name of the queue")
	messageHandle := flag.String("m", "", "The receipt handle of the message")
	flag.Parse()

	//if queue name is empty or queue messageHandle is empty then return.
	//The messageHandle is automatically created when sqs message is created.
	//When we retrieve sqs message the ReceiptHandle is in ReceiptHandle property.
	if *queue == "" || *messageHandle == "" {
		fmt.Println("You must supply a queue name (-q QUEUE) and message receipt handle (-m MESSAGE-HANDLE)")
		return
	}

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	result, err := GetQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	// Delete sqs message
	err = DeleteMessage(sess, queueURL, messageHandle)
	if err != nil {
		fmt.Println("Got an error deleting the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message removed from queue based on URL " + *queueURL)
}
