// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "[TODO]",
        "contact": {
            "name": "API Support",
            "url": "[TODO]",
            "email": "[TODO]"
        },
        "license": {
            "name": "[TODO]",
            "url": "[TODO]"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/login": {
            "post": {
                "description": "requires email and password. Returns Authorization token in header as well",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "POST request for login",
                "parameters": [
                    {
                        "description": "request info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.SignUpINresponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "requires email and password for registration. Returns user info and Authorization token  in header",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "POST request for registration",
                "parameters": [
                    {
                        "description": "user info for sign up",
                        "name": "user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.RegistractionUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/respmodels.SignUpINresponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/reset-password": {
            "post": {
                "description": "requires registered email address. TODO! This endpoint may not work",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reset-password"
                ],
                "summary": "POST request to update password",
                "parameters": [
                    {
                        "description": "user email for update",
                        "name": "reset-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.PasswordResetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/{provider}": {
            "get": {
                "description": "requires param provider, for example google, facebook or apple  (at this moment apple not working) This request redirects to the provider's page for authorization, which in turn transmits a token in the parameters (token)",
                "consumes": [
                    "text/html"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "auth_with_provider get request for auth with provider"
                ],
                "summary": "GET request for auth with provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "provider for auth",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    }
                }
            }
        },
        "/open/advertisements/adv-filter": {
            "post": {
                "description": "endpoint for getting specific advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-filter"
                ],
                "summary": "POST request to get advertisement based on params in filter",
                "parameters": [
                    {
                        "description": "advertisement filter",
                        "name": "advertisement-filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.AdvertisementFilterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementPaginationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/open/advertisements/getall": {
            "get": {
                "description": "endpoint for getting all advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements-getall"
                ],
                "summary": "GET request to get 10 items sorted by creation date in desc order",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/open/advertisements/getbyid/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint to get advertisement based on it's id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "open/advertisements/getbyid/{id}"
                ],
                "summary": "GET request to get advertisement by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "advertisement ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/open/allcategories": {
            "get": {
                "description": "endpoint for getting all categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "open/allcategories"
                ],
                "summary": "GET all categories parents with children in array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/queries.GetCategoriesWithChildrenRow"
                            }
                        }
                    }
                }
            }
        },
        "/protected/advertisement-create": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement creation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-create"
                ],
                "summary": "POST request to create advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement information",
                        "name": "advertisement-create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.CreateAdvertisementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/advertisement-delete": {
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement deletion by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-delete"
                ],
                "summary": "DELETE request to delete advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement id",
                        "name": "advertisement-delete",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.DeleteAdvertisementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/advertisement-getmy": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for getting user advertisements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisements-getmy"
                ],
                "summary": "GET request to get user created advertisements",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/advertisement-patch": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "endpoint for advertisement update",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "advertisement-patch"
                ],
                "summary": "PATCH request to update advertisement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "advertisement information",
                        "name": "advertisement-patch",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.UpdateAdvertisementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.AdvertisementResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/change-email": {
            "post": {
                "description": "requires current password and new email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "change-email"
                ],
                "summary": "POST request to update email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user email for update",
                        "name": "change-email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.EmailChangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/change-password": {
            "post": {
                "description": "requires current password and new password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "change-password"
                ],
                "summary": "POST request to update password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user email for update",
                        "name": "change-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.PasswordChangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/create-password": {
            "patch": {
                "description": "requires token. TODO! This endpoint may not work",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "create-password"
                ],
                "summary": "PATCH request to create new password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "new user password",
                        "name": "create-password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.PasswordCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.StringResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/user-patch": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "requires valid token and user info for update. Returns user info and Authorization token in header",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user-patch"
                ],
                "summary": "PATCH request to update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "user info for update",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqmodels.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.SignUpINresponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        },
        "/protected/userinfo": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "requires valid token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userinfo"
                ],
                "summary": "Get request to see user info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/respmodels.UserInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/respmodels.FailedResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.PaginationInfo": {
            "type": "object",
            "properties": {
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "sort_by": {
                    "type": "string"
                },
                "sort_order": {
                    "type": "string"
                },
                "total_count": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "queries.GetCategoriesWithChildrenRow": {
            "type": "object",
            "properties": {
                "children": {},
                "parent_id": {
                    "type": "integer"
                },
                "parent_name": {
                    "type": "string"
                }
            }
        },
        "reqmodels.AdvertisementFilterRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "format": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "max_exp": {
                    "type": "integer"
                },
                "max_price": {
                    "type": "integer"
                },
                "min_exp": {
                    "type": "integer"
                },
                "min_price": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "sort_by": {
                    "type": "string"
                },
                "sort_order": {
                    "type": "string"
                },
                "time_length": {
                    "type": "integer"
                },
                "title_keyword": {
                    "type": "string"
                }
            }
        },
        "reqmodels.CreateAdvertisementRequest": {
            "type": "object",
            "properties": {
                "attachment": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "experience": {
                    "type": "integer"
                },
                "format": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "mobile_phone": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                }
            }
        },
        "reqmodels.DeleteAdvertisementRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "reqmodels.EmailChangeRequest": {
            "type": "object",
            "properties": {
                "currentPassword": {
                    "type": "string"
                },
                "newEmail": {
                    "type": "string"
                }
            }
        },
        "reqmodels.LoginUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "reqmodels.PasswordChangeRequest": {
            "type": "object",
            "properties": {
                "currentPassword": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string"
                }
            }
        },
        "reqmodels.PasswordCreateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "reqmodels.PasswordResetRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "reqmodels.RegistractionUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "reqmodels.UpdateAdvertisementRequest": {
            "type": "object",
            "properties": {
                "attachment": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "experience": {
                    "type": "integer"
                },
                "format": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "mobile_phone": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "telegram": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 50
                }
            }
        },
        "reqmodels.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                }
            }
        },
        "respmodels.AdvertisementPaginationResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/respmodels.ResponseAdvertismetPagin"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.AdvertisementResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/respmodels.ResponseAdvertismet"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.AdvertisementsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/respmodels.ResponseAdvertismet"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.FailedResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.ResponseAdvertismet": {
            "type": "object",
            "properties": {
                "attachment": {
                    "type": "string"
                },
                "category_name": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "experience": {
                    "type": "integer"
                },
                "format": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "mobile_phone": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "provider_id": {
                    "type": "integer"
                },
                "provider_name": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "respmodels.ResponseAdvertismetPagin": {
            "type": "object",
            "properties": {
                "advertisements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/respmodels.ResponseAdvertismet"
                    }
                },
                "pagination_info": {
                    "$ref": "#/definitions/entities.PaginationInfo"
                }
            }
        },
        "respmodels.ResponseUser": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "respmodels.SignUpINresponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "token": {
                            "type": "string"
                        }
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.StringResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "respmodels.UserInfoResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/respmodels.ResponseUser"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Study marketplace API",
	Description:      "Marketplace to connect students and teachers",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
