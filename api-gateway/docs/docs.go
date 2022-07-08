// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/login": {
            "post": {
                "description": "login user by username and password, tokens expires is unix timestamp",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login user by username and password",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginUserResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/auth.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "refresh token by refresh token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "register user by username and password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register user by username and password",
                "parameters": [
                    {
                        "description": "body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/auth.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/addBillNo": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add bill numbers on track. Every day in your selected time track bill numbers and send email with result about it.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "add bill numbers on track",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.AddOnTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.AddOnTrackResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/addContainer": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add containers on track. Every day in your selected time track container and send email with result about it.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "add containers on track",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.AddOnTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.AddOnTrackResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/addEmail": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add new email to tracking, bill or container doesn't matter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "add new email to tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.AddEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/deleteBillNumbers": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete numbers from tracking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "delete bill numbers from tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.DeleteFromTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/deleteContainers": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete containers from tracking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "delete containers from tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.DeleteFromTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/deleteEmail": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete email from tracking, bill or container doesn't matter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "delete email from tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.DeleteEmailFromTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/getInfo": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get info about number on tracking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "get info about number on tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.GetInfoAboutTrackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.GetInfoAboutTrackResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/schedule/updateTime": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update time of tracking, bill or container doesn't matter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedule Tracking"
                ],
                "summary": "update time of tracking",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.UpdateTrackingTimeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schedule_tracking.BaseAddOnTrackResponse"
                            }
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schedule_tracking.BaseResponse"
                        }
                    }
                }
            }
        },
        "/tracking/trackByBillNumber": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "tracking by bill number, if eta not found will be 0",
                "tags": [
                    "Tracking"
                ],
                "summary": "Track by bill number",
                "parameters": [
                    {
                        "enum": [
                            "AUTO",
                            "FESO",
                            "SKLU",
                            "SITC",
                            "HALU"
                        ],
                        "type": "string",
                        "default": "FESO",
                        "description": "scac code",
                        "name": "scac",
                        "in": "query"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 9,
                        "type": "string",
                        "default": "FLCE405711",
                        "description": "bill number",
                        "name": "number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tracking.BillNumberResponse"
                        }
                    },
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/tracking/trackByContainerNumber": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "tracking by container number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tracking"
                ],
                "summary": "Track by container number",
                "parameters": [
                    {
                        "enum": [
                            "AUTO",
                            "FESO",
                            "SKLU",
                            "SITC",
                            "HALU",
                            "MAEU",
                            "MSCU",
                            "COSU",
                            "ONEY",
                            "KMTU"
                        ],
                        "type": "string",
                        "default": "SKLU",
                        "description": "scac code",
                        "name": "scac",
                        "in": "query"
                    },
                    {
                        "maxLength": 11,
                        "minLength": 10,
                        "type": "string",
                        "default": "TEMU2094051",
                        "description": "container number",
                        "name": "number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tracking.ContainerNumberResponse"
                        }
                    },
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/user/addBillNumbers": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add bill numbers to account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add bill numbers to account",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AddContainers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/addContainers": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add containers to account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add containers to account",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AddContainers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/deleteBillNumbers": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete bill numbers from account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete bill numbers from account",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.DeleteNumbers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/deleteContainers": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete containers from account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete containers from account",
                "parameters": [
                    {
                        "description": "info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.DeleteNumbers"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/getAllBillsContainers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all bill numbers and containers from account",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all bill numbers and containers from account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.AllContainersAndBillNumbers"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.BaseResponse": {
            "type": "object",
            "required": [
                "error",
                "success"
            ],
            "properties": {
                "error": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "auth.LoginUserResponse": {
            "type": "object",
            "required": [
                "refreshToken",
                "refreshTokenExpires",
                "token",
                "tokenExpires",
                "token_type"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "refreshTokenExpires": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "tokenExpires": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "auth.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "auth.User": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schedule_tracking.AddEmailRequest": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schedule_tracking.AddOnTrackRequest": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "schedule_tracking.AddOnTrackResponse": {
            "type": "object",
            "properties": {
                "alreadyOnTrack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schedule_tracking.BaseAddOnTrackResponse"
                    }
                }
            }
        },
        "schedule_tracking.BaseAddOnTrackResponse": {
            "type": "object",
            "properties": {
                "nextRunTime": {
                    "type": "integer"
                },
                "number": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schedule_tracking.BaseResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schedule_tracking.DeleteEmailFromTrackRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "schedule_tracking.DeleteFromTrackRequest": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "schedule_tracking.GetInfoAboutTrackRequest": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                }
            }
        },
        "schedule_tracking.GetInfoAboutTrackResponse": {
            "type": "object",
            "properties": {
                "emails": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "nextRunTime": {
                    "type": "integer"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "schedule_tracking.UpdateTrackingTimeRequest": {
            "type": "object",
            "properties": {
                "newTime": {
                    "type": "string"
                },
                "numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "tracking.BaseInfoAboutMoving": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string"
                },
                "operation_name": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "vessel": {
                    "type": "string"
                }
            }
        },
        "tracking.BillNumberResponse": {
            "type": "object",
            "properties": {
                "Scac": {
                    "type": "string"
                },
                "billNo": {
                    "type": "string"
                },
                "eta_final_delivery": {
                    "type": "integer"
                },
                "info_about_moving": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tracking.BaseInfoAboutMoving"
                    }
                }
            }
        },
        "tracking.ContainerNumberResponse": {
            "type": "object",
            "properties": {
                "Scac": {
                    "type": "string"
                },
                "container": {
                    "type": "string"
                },
                "container_size": {
                    "type": "string"
                },
                "info_about_moving": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tracking.BaseInfoAboutMoving"
                    }
                }
            }
        },
        "user.AddContainers": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "maxItems": 28,
                    "minItems": 10,
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "user.AllContainersAndBillNumbers": {
            "type": "object",
            "properties": {
                "bill_numbers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.Container"
                    }
                },
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.Container"
                    }
                }
            }
        },
        "user.BaseResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "user.Container": {
            "type": "object",
            "required": [
                "id",
                "is_on_track",
                "number"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_on_track": {
                    "type": "boolean"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "user.DeleteNumbers": {
            "type": "object",
            "properties": {
                "numberIds": {
                    "type": "array",
                    "maxItems": 28,
                    "minItems": 10,
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FindMyCargo API",
	Description:      "API server for application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
