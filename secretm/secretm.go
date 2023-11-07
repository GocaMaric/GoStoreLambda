package secretm

import (
	"encoding/json"
	"fmt"

	"example.com/GoStoreLambda/awsgo"
	"example.com/GoStoreLambda/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(numberSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + numberSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(numberSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println(" > Lectura Secret OK " + numberSecret)
	return dataSecret, nil
}
