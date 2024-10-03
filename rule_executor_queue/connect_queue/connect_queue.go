package connect_queue

import (
	"fmt"
	"log"

	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueURL gets the URL of an Amazon SQS queue
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     queueName is the name of the queue
// Output:
//     If success, the URL of the queue and nil
//     Otherwise, an empty string and an error from the call to
func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		fmt.Println("Error creating session to sqs:")
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func SessionQueue() (*session.Session, error) {

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	/*sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))*/

	regionAWS := os.Getenv("SQS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionAWS)},
	)

	if err != nil {
		log.Println("Error creating session to sqs:")
		log.Println(err)
		return nil, err
	}

	/*if err != nil {
		fmt.Println("Error creating session to sqs:")
		fmt.Println(err)
		return nil, err
	}*/

	return sess, nil
}

func QueueURL_SQS(nameQueue *string, session *session.Session) (queueURL *string, err error) {
	// Get URL of queue
	result, err := GetQueueURL(session, nameQueue)
	if err != nil {
		fmt.Println("Error creating session to sqs:")
		fmt.Println(err)
		return nil, err
	}

	queueURL = result.QueueUrl

	return queueURL, nil
}
