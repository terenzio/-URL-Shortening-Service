// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Terence Liu",
            "url": "https://github.com/terenzio/URL-Shortening-Service",
            "email": "terenzio@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/redirect/{shortcode}": {
            "get": {
                "description": "NOTE: Copy the full url including the short code to the browser to be redirected. Do not use swagger as it does not support redirection.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "REDIRECT"
                ],
                "summary": "Redirects the user to the original URL based on the input short code.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short Code",
                        "name": "shortcode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "307": {
                        "description": "Redirected to original url - example: http://localhost:9000/api/v1/redirect/2v5ompxD",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Parameter missing - enter the short code in the URL path",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No original URL exists for the given short code",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/url/add": {
            "post": {
                "description": "Creates a shortened link for the given original URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Creates a shortened link for the given original URL.",
                "parameters": [
                    {
                        "description": "Original URL",
                        "name": "original_url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.AddURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shortened URL",
                        "schema": {
                            "$ref": "#/definitions/domain.AddSuccessResponse"
                        }
                    }
                }
            }
        },
        "/url/display": {
            "get": {
                "description": "Display the home page of the URL shortener",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Display the home page of the URL shortener",
                "responses": {
                    "200": {
                        "description": "URL Mappings",
                        "schema": {
                            "$ref": "#/definitions/domain.URLMapping"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AddSuccessResponse": {
            "type": "object",
            "properties": {
                "shortened_url": {
                    "type": "string"
                }
            }
        },
        "domain.AddURLRequest": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "domain.URLMapping": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                },
                "short_code": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "URL Shortening Service",
	Description:      "This is a URL shortening service that allows users to shorten long URLs especially built for TSMC.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
