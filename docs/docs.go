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
        "/base/captchaPhone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "验证"
                ],
                "summary": "获取手机验证码",
                "parameters": [
                    {
                        "description": "传入参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ShotMessageCode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/response.DTO"
                        }
                    }
                }
            }
        },
        "/reservation/getRandomPhoneNumbers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "预约前台"
                ],
                "summary": "获取随机手机号",
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetsRandomPhoneNumbers"
                        }
                    }
                }
            }
        },
        "/reservation/getReservationRecord": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "预约前台"
                ],
                "summary": "获取预约记录",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetReservationRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetReservationRecord"
                        }
                    }
                }
            }
        },
        "/reservation/reservationPhone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "预约前台"
                ],
                "summary": "预约手机号",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ReservationPhone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/response.ReservationPhone"
                        }
                    }
                }
            }
        },
        "/reservation/searchPhone": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "预约前台"
                ],
                "summary": "模糊查找手机号",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SearchPhone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/response.SearchPhone"
                        }
                    }
                }
            }
        },
        "/wechat/getOpenID": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信"
                ],
                "summary": "获取微信openid",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetOpenID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询成功",
                        "schema": {
                            "$ref": "#/definitions/response.GetOpenID"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.PhoneNumber": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "status": {
                    "description": "手机号状态 0-未售 1-已售",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "request.GetOpenID": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "request.GetReservationRecord": {
            "type": "object",
            "properties": {
                "phone": {
                    "description": "手机号 加密",
                    "type": "string"
                },
                "verify_code": {
                    "description": "验证码",
                    "type": "string"
                }
            }
        },
        "request.ReservationPhone": {
            "type": "object",
            "properties": {
                "contact_phone": {
                    "description": "加密",
                    "type": "string"
                },
                "identify_card": {
                    "description": "加密",
                    "type": "string"
                },
                "open_id": {
                    "type": "string"
                },
                "real_name": {
                    "type": "string"
                },
                "reservation_phone_id": {
                    "type": "integer"
                },
                "salesman_id": {
                    "type": "integer"
                },
                "school": {
                    "type": "string"
                }
            }
        },
        "request.SearchPhone": {
            "type": "object",
            "properties": {
                "suffix": {
                    "type": "string"
                }
            }
        },
        "request.ShotMessageCode": {
            "type": "object",
            "required": [
                "number"
            ],
            "properties": {
                "number": {
                    "description": "手机号",
                    "type": "string"
                }
            }
        },
        "response.DTO": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应代码",
                    "type": "integer"
                },
                "data": {
                    "description": "相应数据"
                },
                "msg": {
                    "description": "响应消息",
                    "type": "string"
                }
            }
        },
        "response.GetOpenID": {
            "type": "object",
            "properties": {
                "open_id": {
                    "type": "string"
                }
            }
        },
        "response.GetReservationRecord": {
            "type": "object",
            "properties": {
                "contact_phone": {
                    "description": "联系电话 加密",
                    "type": "string"
                },
                "created_at": {
                    "description": "订单时间",
                    "type": "string"
                },
                "identify_card": {
                    "description": "身份证 加密",
                    "type": "string"
                },
                "real_name": {
                    "description": "真实名字",
                    "type": "string"
                },
                "reservation_phone": {
                    "description": "预定电话 加密",
                    "type": "string"
                },
                "salesman_id": {
                    "description": "推销员id",
                    "type": "integer"
                },
                "salesman_name": {
                    "description": "推销员名字",
                    "type": "string"
                },
                "salesman_phone": {
                    "description": "推销员电话 加密",
                    "type": "string"
                },
                "school": {
                    "description": "预约学校",
                    "type": "string"
                },
                "status": {
                    "description": "订单状态 0-未支付 1-支付成功 2-错误",
                    "type": "integer"
                }
            }
        },
        "response.GetsRandomPhoneNumbers": {
            "type": "object",
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PhoneNumber"
                    }
                }
            }
        },
        "response.ReservationPhone": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "nonceStr": {
                    "type": "string"
                },
                "package": {
                    "type": "string"
                },
                "paySign": {
                    "type": "string"
                },
                "prepay_id": {
                    "type": "string"
                },
                "signType": {
                    "type": "string"
                },
                "timeStamp": {
                    "type": "string"
                }
            }
        },
        "response.SearchPhone": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PhoneNumber"
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
	Version:          "1.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
