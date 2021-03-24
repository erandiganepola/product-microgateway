/*
 *  Copyright (c) 2020, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package model

import (
	"github.com/go-openapi/spec"
	"github.com/google/uuid"
	logger "github.com/wso2/micro-gw/loggers"
)

// SetInfoSwagger populates the MgwSwagger object with the properties within the openAPI v2
// (swagger) definition.
// The title, version, description, vendor extension map, endpoints based on host and schemes properties,
// and pathItem level information are populated here.
//
// for each pathItem; vendor extensions, available http Methods,
// are populated. Each resource corresponding to a pathItem, has the property called iD, which is a
// UUID.
//
// No operation specific information is extracted.
func (swagger *MgwSwagger) SetInfoSwagger(swagger2 spec.Swagger) {
	swagger.id = swagger2.ID
	if swagger2.Info != nil {
		swagger.description = swagger2.Info.Description
		swagger.title = swagger2.Info.Title
		swagger.version = swagger2.Info.Version
	}
	swagger.vendorExtensions = swagger2.VendorExtensible.Extensions
	swagger.resources = setResourcesSwagger(swagger2)
	swagger.apiType = HTTP
	swagger.xWso2Basepath = swagger2.BasePath

	// According to the definition, multiple schemes can be mentioned. Since the microgateway can assign only one scheme
	// https is prioritized over http. If it is ws or wss, the microgateway will print an error.
	// If the schemes property is not mentioned at all, http will be assigned. (Only swagger 2 version has this property)
	if swagger2.Host != "" {
		urlScheme := ""
		for _, scheme := range swagger2.Schemes {
			//TODO: (VirajSalaka) Introduce Constants
			if scheme == "https" {
				urlScheme = "https://"
				break
			} else if scheme == "http" {
				urlScheme = "http://"
			} else {
				//TODO: (VirajSalaka) Throw an error and stop processing
				logger.LoggerOasparser.Errorf("The scheme : %v for the swagger definition %v:%v is not supported", scheme,
					swagger2.Info.Title, swagger2.Info.Version)
			}
		}
		endpoint := getHostandBasepathandPort(urlScheme + swagger2.Host + swagger2.BasePath)
		swagger.productionUrls = append(swagger.productionUrls, endpoint)
	}
}

// setResourcesSwagger sets swagger (openapi v2) paths as mgwSwagger resources.
func setResourcesSwagger(swagger2 spec.Swagger) []Resource {
	var resources []Resource
	if swagger2.Paths != nil {
		for path, pathItem := range swagger2.Paths.Paths {
			var methodsArray []Operation
			methodFound := false
			if pathItem.Get != nil {
				methodsArray = append(methodsArray, NewOperation("GET", pathItem.Get.Security,
					pathItem.Get.Extensions))
				methodFound = true
			}
			if pathItem.Post != nil {
				methodsArray = append(methodsArray, NewOperation("POST", pathItem.Post.Security,
					pathItem.Post.Extensions))
				methodFound = true
			}
			if pathItem.Put != nil {
				methodsArray = append(methodsArray, NewOperation("PUT", pathItem.Put.Security,
					pathItem.Put.Extensions))
				methodFound = true
			}
			if pathItem.Delete != nil {
				methodsArray = append(methodsArray, NewOperation("DELETE", pathItem.Delete.Security,
					pathItem.Delete.Extensions))
				methodFound = true
			}
			if pathItem.Head != nil {
				methodsArray = append(methodsArray, NewOperation("HEAD", pathItem.Head.Security,
					pathItem.Head.Extensions))
				methodFound = true
			}
			if pathItem.Patch != nil {
				methodsArray = append(methodsArray, NewOperation("PATCH", pathItem.Patch.Security,
					pathItem.Patch.Extensions))
				methodFound = true
			}
			if pathItem.Options != nil {
				methodsArray = append(methodsArray, NewOperation("OPTION", pathItem.Options.Security,
					pathItem.Options.Extensions))
				methodFound = true
			}
			if methodFound {
				resource := setOperationSwagger(path, methodsArray, pathItem)
				resources = append(resources, resource)
			}
		}
	}
	return resources
}

func getSwaggerOperationLevelDetails(operation *spec.Operation, method string) Operation {
	var securityData []map[string][]string = operation.Security
	return NewOperation(method, securityData, operation.Extensions)
}

func setOperationSwagger(path string, methods []Operation, pathItem spec.PathItem) Resource {
	var resource Resource
	resource = Resource{
		path:    path,
		methods: methods,
		// TODO: (VirajSalaka) This will not solve the actual problem when incremental Xds is introduced (used for cluster names)
		iD: uuid.New().String(),
		// PathItem object in swagger 2 specification does not contain summary and description properties
		summary:     "",
		description: "",
		//schemes:          operation.Schemes,
		//tags:             operation.Tags,
		//security:         operation.Security,
		vendorExtensions: pathItem.VendorExtensible.Extensions,
	}
	return resource
}

//SetInfoSwaggerWebSocket populates the mgwSwagger object for web sockets
// TODO - (VirajSalaka) read cors config and populate mgwSwagger feild
func (swagger *MgwSwagger) SetInfoSwaggerWebSocket(apiData map[string]interface{}) {

	data := apiData["data"].(map[string]interface{})
	// UUID in the generated api.yaml file is considerd as swagger.id
	swagger.id = data["id"].(string)
	// Set apiType as WS for websockets
	swagger.apiType = "WS"
	// name and version in api.yaml corresponds to title and version respectively.
	swagger.title = data["name"].(string)
	swagger.version = data["version"].(string)
	// context value in api.yaml is assigned as xWso2Basepath
	swagger.xWso2Basepath = data["context"].(string) + "/" + swagger.version

	// productionURL & sandBoxURL values are extracted from endpointConfig in api.yaml
	endpointConfig := data["endpointConfig"].(map[string]interface{})
	if endpointConfig["sandbox_endpoints"] != nil {
		sandboxEndpoints := endpointConfig["sandbox_endpoints"].(map[string]interface{})
		sandBoxURL := sandboxEndpoints["url"].(string)
		sandBoxEndpoint := getHostandBasepathandPortWebSocket(sandBoxURL)
		swagger.sandboxUrls = append(swagger.sandboxUrls, sandBoxEndpoint)
	}
	if endpointConfig["production_endpoints"] != nil {
		productionEndpoints := endpointConfig["production_endpoints"].(map[string]interface{})
		productionURL := productionEndpoints["url"].(string)
		productionEndpoint := getHostandBasepathandPortWebSocket(productionURL)
		swagger.productionUrls = append(swagger.productionUrls, productionEndpoint)
	}

}
