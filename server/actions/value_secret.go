package actions

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecretValueByARN(arn string) (secretsmanager.GetSecretValueOutput, error) {
	setting := GetAWSSetting()
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(setting.Region),
	}))

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(arn),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return secretsmanager.GetSecretValueOutput{}, err
	}
	if result == nil {
		return secretsmanager.GetSecretValueOutput{}, errors.New("Can't get secret")
	}
	return *result, nil
}