// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-17 13:32:09.761023 +0800 CST m=+0.026373289

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/menu/add": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增权限功能",
                "parameters": [
                    {
                        "description": "新增权限功能",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/check_form.AddMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/delete/:id": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除权限菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "权限功能ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取权限功能列表",
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/update/:id": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改权限菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "权限功能ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/menu/view/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取权限功能详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "权限功能ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录接口",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/check_form.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":true,\"returnCode\":200,\"msg\":\"OK\",data\":{}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "check_form.AddMenu": {
            "type": "object",
            "properties": {
                "component": {
                    "type": "string"
                },
                "hidden": {
                    "type": "integer"
                },
                "icon": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "check_form.LoginForm": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
