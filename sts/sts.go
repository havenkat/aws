package sts

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetIdentity() {
	sess, err := session.NewSession()
	if err != nil {
		log.Panic("session not created", err)
	}
	svc := sts.New(sess)
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
