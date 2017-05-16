package main



import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudSearchDomain_Search() {
	sess := session.Must(session.NewSession(&aws.Config{
  Credentials: credentials.AnonymousCredentials, }))

	svc := cloudsearchdomain.New(sess,aws.NewConfig().WithRegion("us-west-2"),aws.NewConfig().WithEndpoint("search-test2-yoahfnmxlf3lk65n5gucgw6nuq.us-west-2.cloudsearch.amazonaws.com"))

	params := &cloudsearchdomain.SearchInput{
		Query:        aws.String("french"), // Required
	}
	resp, err := svc.Search(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}


func ExampleCloudSearchDomain_Search2() {
	sess := session.Must(session.NewSession(&aws.Config{
  Credentials: credentials.AnonymousCredentials, }))

	svc := cloudsearchdomain.New(sess,aws.NewConfig().WithRegion("us-west-2"),aws.NewConfig().WithEndpoint("search-test2-yoahfnmxlf3lk65n5gucgw6nuq.us-west-2.cloudsearch.amazonaws.com"))


	params := &cloudsearchdomain.SearchInput{
		Query:        aws.String("french"), // Required
		Cursor:       aws.String("initial"),
		//Expr:         aws.String("{\"expression1\":\"_score*rating\", \"expression2\":\"(1/rank)*year\"}"),
		Facet:        aws.String("{'m49':{'sort':'count','size':6}}"),
		//Highlight:    aws.String("'title': {'format': 'text','max_phrases': 2,'pre_tag': '','post_tag':''}"),
		//Partial:      aws.Bool(true),
		//QueryOptions: aws.String("QueryOptions"),
		//QueryParser:  aws.String("QueryParser"),
		Return:       aws.String("m49"),
		Size:         aws.Int64(15),
		Sort:         aws.String("m49 desc"),
		//Stats:        aws.String("Stat"),
	}
	resp, err := svc.Search(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearchDomain_Suggest() {
	sess := session.Must(session.NewSession(&aws.Config{
  Credentials: credentials.AnonymousCredentials, }))

	svc := cloudsearchdomain.New(sess,aws.NewConfig().WithRegion("us-west-2"),aws.NewConfig().WithEndpoint("search-test2-yoahfnmxlf3lk65n5gucgw6nuq.us-west-2.cloudsearch.amazonaws.com"))

	params := &cloudsearchdomain.SuggestInput{
		Query:     aws.String("fren"),     // Required
		Suggester: aws.String("test2"), // Required
		Size:      aws.Int64(10),
	}
	resp, err := svc.Suggest(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearchDomain_UploadDocuments() {
	sess := session.Must(session.NewSession())

	svc := cloudsearchdomain.New(sess)

	params := &cloudsearchdomain.UploadDocumentsInput{
		ContentType: aws.String("ContentType"),          // Required
		Documents:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.UploadDocuments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}




func main() {

	//aws.DefaultConfig.Region = "us-west-2"
	//ExampleCloudSearchDomain_Suggest();
	ExampleCloudSearchDomain_Search();
	fmt.Println("-------------------------Delemiter-------------------------------------")
	ExampleCloudSearchDomain_Search2()

}
