/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ampolicy

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/pcf/logger"
	"free5gc/src/pcf/handler/message"
	"free5gc/src/pcf/util"

	"github.com/gin-gonic/gin"
)

func PoliciesPolAssoIdDelete(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["polAssoId"], _ = c.Params.Get("polAssoId")
	channelMsg := message.NewHttpChannelMessage(message.EventAMPolicyDelete, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// PoliciesPolAssoIdGet -
func PoliciesPolAssoIdGet(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["polAssoId"], _ = c.Params.Get("polAssoId")
	channelMsg := message.NewHttpChannelMessage(message.EventAMPolicyGet, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}

// PoliciesPolAssoIdUpdatePost -
func PoliciesPolAssoIdUpdatePost(c *gin.Context) {
	var policyAssociationUpdateRequest models.PolicyAssociationUpdateRequest
	err := c.ShouldBindJSON(&policyAssociationUpdateRequest)
	if err != nil {
		rsp := util.GetProblemDetail("Malformed request syntax", util.ERROR_REQUEST_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	req := http_wrapper.NewRequest(c.Request, policyAssociationUpdateRequest)
	req.Params["polAssoId"], _ = c.Params.Get("polAssoId")
	channelMsg := message.NewHttpChannelMessage(message.EventAMPolicyUpdate, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// PoliciesPost -
func PoliciesPost(c *gin.Context) {
	var policyAssociationRequest models.PolicyAssociationRequest
	err := c.ShouldBindJSON(&policyAssociationRequest)
	if err != nil {
		rsp := util.GetProblemDetail("Malformed request syntax", util.ERROR_REQUEST_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	if policyAssociationRequest.Supi == "" || policyAssociationRequest.NotificationUri == "" {
		rsp := util.GetProblemDetail("Miss Mandotory IE", util.ERROR_REQUEST_PARAMETERS)
		logger.HandlerLog.Errorln(rsp.Detail)
		c.JSON(int(rsp.Status), rsp)
		return
	}
	req := http_wrapper.NewRequest(c.Request, policyAssociationRequest)
	channelMsg := message.NewHttpChannelMessage(message.EventAMPolicyCreate, req)

	message.SendMessage(channelMsg)
	recvMsg := <-channelMsg.HttpChannel
	HTTPResponse := recvMsg.HTTPResponse
	for key, val := range HTTPResponse.Header {
		c.Header(key, val[0])
	}
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
