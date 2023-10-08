// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthcheck": {
            "get": {
                "tags": [
                    "Default API"
                ],
                "summary": "Healthcheck",
                "responses": {}
            }
        },
        "/user": {
            "post": {
                "tags": [
                    "Internal :: User"
                ],
                "summary": "Get single user by ID",
                "parameters": [
                    {
                        "description": "Create user input",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/user/": {
            "get": {
                "tags": [
                    "Internal :: User"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserPagination"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "Internal :: Auth"
                ],
                "summary": "Get single user by ID",
                "parameters": [
                    {
                        "description": "Login input",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "tags": [
                    "Internal :: User"
                ],
                "summary": "Get single user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.CreateUserDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginDTO": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "updatedBy": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UserPagination": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.User"
                    }
                },
                "page": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Auth Service",
	Description:      "Reservio 2.0 Auth Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}