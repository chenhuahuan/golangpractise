package main



import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudSearch_BuildSuggesters() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.BuildSuggestersInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.BuildSuggesters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_CreateDomain() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.CreateDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.CreateDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DefineAnalysisScheme() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DefineAnalysisSchemeInput{
		AnalysisScheme: &cloudsearch.AnalysisScheme{ // Required
			AnalysisSchemeLanguage: aws.String("AnalysisSchemeLanguage"), // Required
			AnalysisSchemeName:     aws.String("StandardName"),           // Required
			AnalysisOptions: &cloudsearch.AnalysisOptions{
				AlgorithmicStemming:            aws.String("AlgorithmicStemming"),
				JapaneseTokenizationDictionary: aws.String("String"),
				StemmingDictionary:             aws.String("String"),
				Stopwords:                      aws.String("String"),
				Synonyms:                       aws.String("String"),
			},
		},
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DefineAnalysisScheme(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DefineExpression() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DefineExpressionInput{
		DomainName: aws.String("DomainName"), // Required
		Expression: &cloudsearch.Expression{ // Required
			ExpressionName:  aws.String("StandardName"),    // Required
			ExpressionValue: aws.String("ExpressionValue"), // Required
		},
	}
	resp, err := svc.DefineExpression(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DefineIndexField() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DefineIndexFieldInput{
		DomainName: aws.String("DomainName"), // Required
		IndexField: &cloudsearch.IndexField{ // Required
			IndexFieldName: aws.String("DynamicFieldName"), // Required
			IndexFieldType: aws.String("IndexFieldType"),   // Required
			DateArrayOptions: &cloudsearch.DateArrayOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			DateOptions: &cloudsearch.DateOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SortEnabled:   aws.Bool(true),
				SourceField:   aws.String("FieldName"),
			},
			DoubleArrayOptions: &cloudsearch.DoubleArrayOptions{
				DefaultValue:  aws.Float64(1.0),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			DoubleOptions: &cloudsearch.DoubleOptions{
				DefaultValue:  aws.Float64(1.0),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SortEnabled:   aws.Bool(true),
				SourceField:   aws.String("FieldName"),
			},
			IntArrayOptions: &cloudsearch.IntArrayOptions{
				DefaultValue:  aws.Int64(1),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			IntOptions: &cloudsearch.IntOptions{
				DefaultValue:  aws.Int64(1),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SortEnabled:   aws.Bool(true),
				SourceField:   aws.String("FieldName"),
			},
			LatLonOptions: &cloudsearch.LatLonOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SortEnabled:   aws.Bool(true),
				SourceField:   aws.String("FieldName"),
			},
			LiteralArrayOptions: &cloudsearch.LiteralArrayOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SourceFields:  aws.String("FieldNameCommaList"),
			},
			LiteralOptions: &cloudsearch.LiteralOptions{
				DefaultValue:  aws.String("FieldValue"),
				FacetEnabled:  aws.Bool(true),
				ReturnEnabled: aws.Bool(true),
				SearchEnabled: aws.Bool(true),
				SortEnabled:   aws.Bool(true),
				SourceField:   aws.String("FieldName"),
			},
			TextArrayOptions: &cloudsearch.TextArrayOptions{
				AnalysisScheme:   aws.String("Word"),
				DefaultValue:     aws.String("FieldValue"),
				HighlightEnabled: aws.Bool(true),
				ReturnEnabled:    aws.Bool(true),
				SourceFields:     aws.String("FieldNameCommaList"),
			},
			TextOptions: &cloudsearch.TextOptions{
				AnalysisScheme:   aws.String("Word"),
				DefaultValue:     aws.String("FieldValue"),
				HighlightEnabled: aws.Bool(true),
				ReturnEnabled:    aws.Bool(true),
				SortEnabled:      aws.Bool(true),
				SourceField:      aws.String("FieldName"),
			},
		},
	}
	resp, err := svc.DefineIndexField(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DefineSuggester() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DefineSuggesterInput{
		DomainName: aws.String("DomainName"), // Required
		Suggester: &cloudsearch.Suggester{ // Required
			DocumentSuggesterOptions: &cloudsearch.DocumentSuggesterOptions{ // Required
				SourceField:    aws.String("FieldName"), // Required
				FuzzyMatching:  aws.String("SuggesterFuzzyMatching"),
				SortExpression: aws.String("String"),
			},
			SuggesterName: aws.String("StandardName"), // Required
		},
	}
	resp, err := svc.DefineSuggester(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DeleteAnalysisScheme() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DeleteAnalysisSchemeInput{
		AnalysisSchemeName: aws.String("StandardName"), // Required
		DomainName:         aws.String("DomainName"),   // Required
	}
	resp, err := svc.DeleteAnalysisScheme(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DeleteDomain() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DeleteDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeleteDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DeleteExpression() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DeleteExpressionInput{
		DomainName:     aws.String("DomainName"),   // Required
		ExpressionName: aws.String("StandardName"), // Required
	}
	resp, err := svc.DeleteExpression(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DeleteIndexField() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DeleteIndexFieldInput{
		DomainName:     aws.String("DomainName"),       // Required
		IndexFieldName: aws.String("DynamicFieldName"), // Required
	}
	resp, err := svc.DeleteIndexField(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DeleteSuggester() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DeleteSuggesterInput{
		DomainName:    aws.String("DomainName"),   // Required
		SuggesterName: aws.String("StandardName"), // Required
	}
	resp, err := svc.DeleteSuggester(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeAnalysisSchemes() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeAnalysisSchemesInput{
		DomainName: aws.String("DomainName"), // Required
		AnalysisSchemeNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
		Deployed: aws.Bool(true),
	}
	resp, err := svc.DescribeAnalysisSchemes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeAvailabilityOptions() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeAvailabilityOptionsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Bool(true),
	}
	resp, err := svc.DescribeAvailabilityOptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeDomains() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeDomainsInput{
		DomainNames: []*string{
			aws.String("DomainName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeDomains(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeExpressions() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeExpressionsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Bool(true),
		ExpressionNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeExpressions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeIndexFields() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeIndexFieldsInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Bool(true),
		FieldNames: []*string{
			aws.String("DynamicFieldName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeIndexFields(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeScalingParameters() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeScalingParametersInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeScalingParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeServiceAccessPolicies() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeServiceAccessPoliciesInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Bool(true),
	}
	resp, err := svc.DescribeServiceAccessPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_DescribeSuggesters() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.DescribeSuggestersInput{
		DomainName: aws.String("DomainName"), // Required
		Deployed:   aws.Bool(true),
		SuggesterNames: []*string{
			aws.String("StandardName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeSuggesters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_IndexDocuments() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.IndexDocumentsInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.IndexDocuments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_ListDomainNames() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	var params *cloudsearch.ListDomainNamesInput
	resp, err := svc.ListDomainNames(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_UpdateAvailabilityOptions() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.UpdateAvailabilityOptionsInput{
		DomainName: aws.String("DomainName"), // Required
		MultiAZ:    aws.Bool(true),           // Required
	}
	resp, err := svc.UpdateAvailabilityOptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_UpdateScalingParameters() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.UpdateScalingParametersInput{
		DomainName: aws.String("DomainName"), // Required
		ScalingParameters: &cloudsearch.ScalingParameters{ // Required
			DesiredInstanceType:     aws.String("PartitionInstanceType"),
			DesiredPartitionCount:   aws.Int64(1),
			DesiredReplicationCount: aws.Int64(1),
		},
	}
	resp, err := svc.UpdateScalingParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearch_UpdateServiceAccessPolicies() {
	sess := session.Must(session.NewSession())

	svc := cloudsearch.New(sess)

	params := &cloudsearch.UpdateServiceAccessPoliciesInput{
		AccessPolicies: aws.String("PolicyDocument"), // Required
		DomainName:     aws.String("DomainName"),     // Required
	}
	resp, err := svc.UpdateServiceAccessPolicies(params)

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
}
