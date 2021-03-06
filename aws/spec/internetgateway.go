/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CreateInternetgateway struct {
	_      string `action:"create" entity:"internetgateway" awsAPI:"ec2" awsCall:"CreateInternetGateway" awsInput:"ec2.CreateInternetGatewayInput" awsOutput:"ec2.CreateInternetGatewayOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
}

func (cmd *CreateInternetgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateInternetgateway) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateInternetGatewayOutput).InternetGateway.InternetGatewayId)
}

type DeleteInternetgateway struct {
	_      string `action:"delete" entity:"internetgateway" awsAPI:"ec2" awsCall:"DeleteInternetGateway" awsInput:"ec2.DeleteInternetGatewayInput" awsOutput:"ec2.DeleteInternetGatewayOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"InternetGatewayId" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteInternetgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type AttachInternetgateway struct {
	_      string `action:"attach" entity:"internetgateway" awsAPI:"ec2" awsCall:"AttachInternetGateway" awsInput:"ec2.AttachInternetGatewayInput" awsOutput:"ec2.AttachInternetGatewayOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"InternetGatewayId" awsType:"awsstr" templateName:"id" required:""`
	Vpc    *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc" required:""`
}

func (cmd *AttachInternetgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type DetachInternetgateway struct {
	_      string `action:"detach" entity:"internetgateway" awsAPI:"ec2" awsCall:"DetachInternetGateway" awsInput:"ec2.DetachInternetGatewayInput" awsOutput:"ec2.DetachInternetGatewayOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"InternetGatewayId" awsType:"awsstr" templateName:"id" required:""`
	Vpc    *string `awsName:"VpcId" awsType:"awsstr" templateName:"vpc" required:""`
}

func (cmd *DetachInternetgateway) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
