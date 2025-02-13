// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

/*
 * Policy Manager Service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"log"
	"net/http"

	"emperror.dev/errors"
	opabl "fybrik.io/fybrik/connectors/open_policy_agent/lib"
	openapiclientmodels "fybrik.io/fybrik/pkg/taxonomy/model/base"
	"github.com/gdexlab/go-render/render"
	"github.com/gin-gonic/gin"
)

type ConnectorController struct {
	OpaServerURL  string
	OpaReader     *opabl.OpaReader
	CatalogReader *opabl.CatalogReader
}

func (e *ConnectorController) contactOPA(input openapiclientmodels.PolicyManagerRequest, creds string) (openapiclientmodels.PolicyManagerResponse, error) {
	policyToBeEvaluated := "dataapi/authz/verdict"
	eval, err := e.OpaReader.GetOPADecisions(&input, creds, e.CatalogReader, policyToBeEvaluated)
	if err != nil {
		return openapiclientmodels.PolicyManagerResponse{}, errors.Wrap(err, "GetOPADecisions error in contactOPA")
	}

	output := render.AsCode(eval)
	log.Println("received PolicyManagerResponse from GetOPADecisions:", output)

	return eval, nil
}

func (e *ConnectorController) GetPoliciesDecisions(c *gin.Context) {
	log.Println("in GetPoliciesDecisions of V2 OPA Connector!")
	input := new(openapiclientmodels.PolicyManagerRequest)
	if err := c.ShouldBindJSON(input); err != nil {
		log.Println("Error during ShouldBindJSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output := render.AsCode(input)
	log.Println("bound PolicyManagerRequest:", output)

	resp, err := e.contactOPA(*input, c.Request.Header["X-Request-Cred"][0])
	if err != nil {
		log.Println("Error during contactOPA: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, resp)
}
