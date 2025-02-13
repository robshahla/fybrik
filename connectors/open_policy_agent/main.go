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

package main

import (
	"log"
	"os"
	"strconv"
	"time"

	openapiserver "fybrik.io/fybrik/connectors/open_policy_agent/go"
	sw "fybrik.io/fybrik/connectors/open_policy_agent/go"
	opabl "fybrik.io/fybrik/connectors/open_policy_agent/lib"
	clients "fybrik.io/fybrik/pkg/connectors/clients"
	"github.com/hashicorp/go-retryablehttp"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	log.Printf("Env. variable extracted: %s - %s\n", key, value)
	return value
}

func main() {
	log.Println("in main of V2 OPA Connector: ")
	catalogConnectorAddress := getEnv("CATALOG_CONNECTOR_URL")
	catalogProviderName := getEnv("CATALOG_PROVIDER_NAME")

	timeOutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeOut, err := strconv.Atoi(timeOutInSecs)
	if err != nil {
		log.Println("error in strconv:", err.Error())
		return
	}
	opaServerURL := getEnv("OPA_SERVER_URL")

	connectionTimeout := time.Duration(timeOut) * time.Second
	dataCatalog, err := clients.NewGrpcDataCatalog(catalogProviderName, catalogConnectorAddress, connectionTimeout)
	if err != nil {
		log.Println("Error during creation of data catalog connection!")
		return
	}
	catalogReader := opabl.NewCatalogReader(&dataCatalog)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	standardClient := retryClient.HTTPClient // *http.Client

	opaReader := opabl.NewOpaReader(opaServerURL, standardClient)

	connController := &openapiserver.ConnectorController{
		OpaServerURL:  opaServerURL,
		OpaReader:     opaReader,
		CatalogReader: catalogReader,
	}

	router := sw.NewRouter(connController)

	log.Printf("Server started")
	log.Fatal(router.Run(":8080"))
}
