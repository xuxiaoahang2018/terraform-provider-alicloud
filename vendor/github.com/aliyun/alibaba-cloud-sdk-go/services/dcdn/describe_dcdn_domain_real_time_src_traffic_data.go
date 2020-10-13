package dcdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDcdnDomainRealTimeSrcTrafficData invokes the dcdn.DescribeDcdnDomainRealTimeSrcTrafficData API synchronously
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficData(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) (response *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, err error) {
	response = CreateDescribeDcdnDomainRealTimeSrcTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRealTimeSrcTrafficDataWithChan invokes the dcdn.DescribeDcdnDomainRealTimeSrcTrafficData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficDataWithChan(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) (<-chan *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRealTimeSrcTrafficData(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDcdnDomainRealTimeSrcTrafficDataWithCallback invokes the dcdn.DescribeDcdnDomainRealTimeSrcTrafficData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficDataWithCallback(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest, callback func(response *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRealTimeSrcTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRealTimeSrcTrafficData(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDcdnDomainRealTimeSrcTrafficDataRequest is the request struct for api DescribeDcdnDomainRealTimeSrcTrafficData
type DescribeDcdnDomainRealTimeSrcTrafficDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnDomainRealTimeSrcTrafficDataResponse is the response struct for api DescribeDcdnDomainRealTimeSrcTrafficData
type DescribeDcdnDomainRealTimeSrcTrafficDataResponse struct {
	*responses.BaseResponse
	RequestId                         string                            `json:"RequestId" xml:"RequestId"`
	DomainName                        string                            `json:"DomainName" xml:"DomainName"`
	StartTime                         string                            `json:"StartTime" xml:"StartTime"`
	EndTime                           string                            `json:"EndTime" xml:"EndTime"`
	DataInterval                      string                            `json:"DataInterval" xml:"DataInterval"`
	RealTimeSrcTrafficDataPerInterval RealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval" xml:"RealTimeSrcTrafficDataPerInterval"`
}

// CreateDescribeDcdnDomainRealTimeSrcTrafficDataRequest creates a request to invoke DescribeDcdnDomainRealTimeSrcTrafficData API
func CreateDescribeDcdnDomainRealTimeSrcTrafficDataRequest() (request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) {
	request = &DescribeDcdnDomainRealTimeSrcTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRealTimeSrcTrafficData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainRealTimeSrcTrafficDataResponse creates a response to parse from DescribeDcdnDomainRealTimeSrcTrafficData response
func CreateDescribeDcdnDomainRealTimeSrcTrafficDataResponse() (response *DescribeDcdnDomainRealTimeSrcTrafficDataResponse) {
	response = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
