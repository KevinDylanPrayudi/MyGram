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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/comment": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show List of Comment by json Comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Show List of Comment",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.GetCommentResult"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create by json comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Create comment",
                "parameters": [
                    {
                        "description": "Show List of Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.AddComment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.AddCommentResult"
                        }
                    }
                }
            }
        },
        "/comment/{commentId}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update by json comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Update an comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateComment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateCommentResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete by comment ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Delete an comment",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.DeleteCommentResult"
                        }
                    }
                }
            }
        },
        "/photo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show List of Photo by json photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Show List of Photo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.GetPhotoResult"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create by json photo",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Create Photo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title of photo",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "caption of photo",
                        "name": "caption",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.AddPhotoResult"
                        }
                    }
                }
            }
        },
        "/photo/{photoId}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update by json photo",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Update an photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title of photo",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "caption of photo",
                        "name": "caption",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "account image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.UpdatePhotoResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete by photo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Delete an photo",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "photo ID",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.DeletePhotoResult"
                        }
                    }
                }
            }
        },
        "/socialmedias": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show List of Social Media by json Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Show List of Social Media",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.GetSocialMediaResult"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Social Mediaby json Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Create Social Media",
                "parameters": [
                    {
                        "description": "Show List of SocialMedia",
                        "name": "socialmedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.AddSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.AddSocialMediaResult"
                        }
                    }
                }
            }
        },
        "/socialmedias/{socialmediaId}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update by json Social Media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Update an Social Media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "socialmediaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Social Media",
                        "name": "photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateSocialMediaResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete by socialmedia ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedia"
                ],
                "summary": "Delete an socialmedia",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "socialmedia ID",
                        "name": "socialmediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.DeleteSocialMediaResult"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login by json users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login an users",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.LoginUserResult"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "add by json users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add an users",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.AddUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.AddUserResult"
                        }
                    }
                }
            }
        },
        "/users/{userId}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update by json users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update an users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.UpdateUserResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete by User ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete an User",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.DeleteUserResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.AddComment": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                },
                "photo_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.AddCommentResult": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "message"
                },
                "photo_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.AddPhotoResult": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "caption"
                },
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "photo_url": {
                    "type": "string",
                    "example": "photo url"
                },
                "title": {
                    "type": "string",
                    "example": "title"
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.AddSocialMedia": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "testing"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "testing"
                }
            }
        },
        "structs.AddSocialMediaResult": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "testing"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "testing"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "structs.AddUser": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "format": "int64",
                    "example": 9
                },
                "email": {
                    "type": "string",
                    "example": "testing@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "testing"
                }
            }
        },
        "structs.AddUserResult": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "format": "int64",
                    "example": 9
                },
                "email": {
                    "type": "string",
                    "example": "testing@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "testing"
                }
            }
        },
        "structs.DeleteCommentResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Your comment has been successfully deleted"
                }
            }
        },
        "structs.DeletePhotoResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Your photo has been successfully deleted"
                }
            }
        },
        "structs.DeleteSocialMediaResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Your social media has been successfully deleted"
                }
            }
        },
        "structs.DeleteUserResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Your account has been successfully deleted"
                }
            }
        },
        "structs.GetCommentResult": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "message"
                },
                "photo_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user": {
                    "type": "object",
                    "properties": {
                        "email": {
                            "type": "string",
                            "example": "testing@gmail.com"
                        },
                        "id": {
                            "type": "integer",
                            "format": "int64",
                            "example": 1
                        },
                        "username": {
                            "type": "string",
                            "example": "testing"
                        }
                    }
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.GetPhotoResult": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "caption"
                },
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "photo_url": {
                    "type": "string",
                    "example": "photo_url"
                },
                "title": {
                    "type": "string",
                    "example": "title"
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user": {
                    "type": "object",
                    "properties": {
                        "email": {
                            "type": "string",
                            "example": "testing@gmail.com"
                        },
                        "username": {
                            "type": "string",
                            "example": "testing"
                        }
                    }
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.GetSocialMediaResult": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "date"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "testing"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "testing"
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user": {
                    "type": "object",
                    "properties": {
                        "email": {
                            "type": "string",
                            "example": "testing@gmail.com"
                        },
                        "id": {
                            "type": "integer",
                            "format": "int64",
                            "example": 1
                        },
                        "username": {
                            "type": "string",
                            "example": "testing"
                        }
                    }
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "structs.LoginUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "testing@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "structs.LoginUserResult": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "result of generated token"
                }
            }
        },
        "structs.UpdateComment": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "update message"
                }
            }
        },
        "structs.UpdateCommentResult": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "update message"
                },
                "photo_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.UpdatePhotoResult": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "update caption"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "photo_url": {
                    "type": "string",
                    "example": "update photo url"
                },
                "title": {
                    "type": "string",
                    "example": "update title"
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user_id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                }
            }
        },
        "structs.UpdateSocialMedia": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "update testing"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "update testing"
                }
            }
        },
        "structs.UpdateSocialMediaResult": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "update testing"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "update testing"
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "structs.UpdateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "testing1@gmail.com"
                },
                "username": {
                    "type": "string",
                    "example": "testing1"
                }
            }
        },
        "structs.UpdateUserResult": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "format": "int64",
                    "example": 9
                },
                "email": {
                    "type": "string",
                    "example": "testing@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "date"
                },
                "username": {
                    "type": "string",
                    "example": "testing"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "mygram-production-7b69.up.railway.app",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyGram api",
	Description:      "This is a API to reach out my final assignment",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
