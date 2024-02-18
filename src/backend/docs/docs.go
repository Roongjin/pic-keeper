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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/v1/login": {
            "post": {
                "description": "Admin login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Admin login via email and password",
                "parameters": [
                    {
                        "description": "email and password of the administrator",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The token will be returned inside the data field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Administrator does not exist",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Unhandled internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/v1/logout": {
            "put": {
                "description": "Logout endpoint for administrators",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Logout endpoint for administrators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session token is required",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Logged out successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Administrator is no longer existed",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Unhandled internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/v1/refresh": {
            "get": {
                "description": "Refresh session token for administrators",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Refresh session token for administrators",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Expired token is required",
                        "name": "ExpiredToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The refreshed session token will be returned inside the data field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "403": {
                        "description": "No Authorization header provided",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Re-login is needed or the administrator may no longer exist",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Unhandled internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/v1/verifications/unverified-photographers": {
            "get": {
                "description": "List all unverified photographers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "List all unverified photographers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session token is required",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of unverified photographers will be located inside the data field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Administrator is no longer existed",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Unhandled internal server error OR session token cannot be verified",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/v1/verifications/verify/:id": {
            "get": {
                "description": "Verify the specified photographer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Verify the specified photographer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Session token is required",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the to-be-verified photographer",
                        "name": "photographerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully verified the photographer",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Issues with finding the photographer in the database",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.JSONErrorResult": {
            "type": "object",
            "properties": {
                "error": {},
                "status": {
                    "type": "string",
                    "example": "failed"
                }
            }
        },
        "model.JSONSuccessResult": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "model.LoginCredentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "abc123"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pic-keeper APIs",
	Description:      "This is the back-end documentation of the pic-keeper project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
