// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article/getArticleDetails": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取文章详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article文章"
                ],
                "summary": "获取文章详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/article/getArticleList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取文章列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article文章"
                ],
                "summary": "获取文章列表",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.AllListForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/article/getUserArticleList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取用户文章列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article文章"
                ],
                "summary": "获取用户文章列表",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.UserArticleListForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/article/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "保存文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article文章"
                ],
                "summary": "保存文章",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.SaveForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/article/saveDraft": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "保存草稿",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article文章"
                ],
                "summary": "保存草稿",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.SaveDraftForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/articleColumn/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除专栏",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articleColumn专栏"
                ],
                "summary": "根据id删除专栏",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "专栏id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/articleColumn/details": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取专栏详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articleColumn专栏"
                ],
                "summary": "获取专栏详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "专栏id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/articleColumn/getAllArticleColumnList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部专栏",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articleColumn专栏"
                ],
                "summary": "获取全部专栏",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/articleColumn/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "保存专栏",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articleColumn专栏"
                ],
                "summary": "保存专栏",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/articleColumn.SaveForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/enum/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除指定枚举",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enum枚举"
                ],
                "summary": "根据id删除指定枚举",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "枚举id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/enum/details": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取枚举详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enum枚举"
                ],
                "summary": "获取枚举详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "枚举id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/enum/getAllEnums": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取全部数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enum枚举"
                ],
                "summary": "获取全部数据",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/enum/getEnumsByType": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据分类查询枚举",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enum枚举"
                ],
                "summary": "根据分类查询枚举",
                "parameters": [
                    {
                        "type": "string",
                        "description": "枚举类型code",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/enum/save": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "增加、编辑",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enum枚举"
                ],
                "summary": "增加、编辑",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/enum.EnumsForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/file/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "文件上传",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file文件"
                ],
                "summary": "文件上传",
                "parameters": [
                    {
                        "type": "file",
                        "description": "请求对象",
                        "name": "param",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/user/getUserDetails": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取用户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user用户"
                ],
                "summary": "获取用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取user用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user用户"
                ],
                "summary": "获取user用户列表",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ListForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user用户"
                ],
                "summary": "用户登陆",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        },
        "/user/sendVerificationCode": {
            "post": {
                "description": "发送验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user用户"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "description": "请求对象",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.VerificationCodeForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseResultInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.EmptyInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "article.AllListForm": {
            "type": "object",
            "required": [
                "pageNo",
                "pageSize"
            ],
            "properties": {
                "classify": {
                    "type": "string"
                },
                "pageNo": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 1
                },
                "pageSize": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 10
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "article.SaveDraftForm": {
            "type": "object",
            "required": [
                "title",
                "userId"
            ],
            "properties": {
                "classify": {
                    "type": "string"
                },
                "collectionColumn": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "coverImg": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "article.SaveForm": {
            "type": "object",
            "required": [
                "classify",
                "collectionColumn",
                "content",
                "coverImg",
                "summary",
                "tags",
                "title",
                "userId"
            ],
            "properties": {
                "classify": {
                    "type": "string"
                },
                "collectionColumn": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "coverImg": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "article.UserArticleListForm": {
            "type": "object",
            "required": [
                "pageNo",
                "pageSize"
            ],
            "properties": {
                "pageNo": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 1
                },
                "pageSize": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 10
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "articleColumn.SaveForm": {
            "type": "object",
            "required": [
                "coverImg",
                "introduction",
                "name"
            ],
            "properties": {
                "coverImg": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "enum.EnumsForm": {
            "type": "object",
            "required": [
                "name",
                "typeCode",
                "typeName",
                "value"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "typeCode": {
                    "type": "string"
                },
                "typeName": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "user.ListForm": {
            "type": "object",
            "required": [
                "pageNo",
                "pageSize"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "pageNo": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 1
                },
                "pageSize": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 10
                }
            }
        },
        "user.LoginForm": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "mk@163.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3,
                    "example": "123456"
                }
            }
        },
        "user.RegisterForm": {
            "type": "object",
            "required": [
                "code",
                "email",
                "password"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "user.VerificationCodeForm": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "utils.EmptyInfo": {
            "type": "object"
        },
        "utils.ResponseResultInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
