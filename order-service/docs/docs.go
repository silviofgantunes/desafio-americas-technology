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
        "/orders": {
            "get": {
                "description": "List all orders",
                "produces": [
                    "application/json"
                ],
                "summary": "List all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Order details",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                }
            }
        },
        "/orders/limit/{id}": {
            "delete": {
                "responses": {}
            }
        },
        "/orders/user/{user_id}": {
            "get": {
                "description": "List orders by user",
                "produces": [
                    "application/json"
                ],
                "summary": "List orders by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Order"
                            }
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get an order by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get an order by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateOrderRequest": {
            "type": "object",
            "required": [
                "amount",
                "direction",
                "pair",
                "type",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "direction": {
                    "$ref": "#/definitions/models.Direction"
                },
                "pair": {
                    "$ref": "#/definitions/models.Pair"
                },
                "type": {
                    "$ref": "#/definitions/models.OrderType"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Direction": {
            "type": "string",
            "enum": [
                "buy",
                "sell"
            ],
            "x-enum-varnames": [
                "DirectionBuy",
                "DirectionSell"
            ]
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "direction": {
                    "$ref": "#/definitions/models.Direction"
                },
                "id": {
                    "type": "string"
                },
                "pair": {
                    "$ref": "#/definitions/models.Pair"
                },
                "type": {
                    "$ref": "#/definitions/models.OrderType"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.OrderType": {
            "type": "string",
            "enum": [
                "market",
                "limit"
            ],
            "x-enum-varnames": [
                "OrderTypeMarket",
                "OrderTypeLimit"
            ]
        },
        "models.Pair": {
            "type": "string",
            "enum": [
                "USD/BTC",
                "USD/ADA",
                "USD/ETH",
                "BTC/USD",
                "ETH/USD",
                "ADA/USD"
            ],
            "x-enum-varnames": [
                "PairUSDBTC",
                "PairUSDADA",
                "PairUSDETH",
                "PairBTCUSD",
                "PairETHUSD",
                "PairADAUSD"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}