// Copyright 2017, Dell EMC, Inc.

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "on-network api",
    "version": "0.0.1"
  },
  "paths": {
    "/about": {
      "get": {
        "description": "Information about on-network",
        "tags": [
          "/about"
        ],
        "summary": "Get about",
        "operationId": "aboutGet",
        "security": [
          {
            "Bearer": []
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully retrieved about",
            "schema": {
              "$ref": "#/definitions/About"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/checkVlan": {
      "post": {
        "description": "Check if vlan exists",
        "tags": [
          "/checkVlan"
        ],
        "summary": "Checks if vlan exists",
        "operationId": "checkVlan",
        "security": [
          {
            "Bearer": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CheckVlan"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully returned whether a vlan exists or not"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Allow users to log in, and to receive a Token\n",
        "tags": [
          "auth"
        ],
        "parameters": [
          {
            "description": "The username/password",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Login"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Login Success",
            "schema": {
              "$ref": "#/definitions/Token"
            }
          },
          "401": {
            "description": "Whether the user is not found or error while login",
            "schema": {
              "$ref": "#/definitions/LoginError"
            }
          }
        }
      }
    },
    "/switchConfig": {
      "post": {
        "description": "Get switch running config",
        "tags": [
          "/switchConfig"
        ],
        "summary": "Get switch running config",
        "operationId": "switchConfig",
        "security": [
          {
            "Bearer": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Switch"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully returned switch running config",
            "schema": {
              "$ref": "#/definitions/SwitchConfigResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/switchFirmware": {
      "post": {
        "description": "Get switch Firmware Version",
        "tags": [
          "/switchFirmware"
        ],
        "summary": "Get switch Firmware Version",
        "operationId": "switchFirmware",
        "security": [
          {
            "Bearer": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Switch"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully returned switch firmware version",
            "schema": {
              "$ref": "#/definitions/SwitchVersionResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/switchVersion": {
      "post": {
        "description": "Get switch Version",
        "tags": [
          "/switchVersion"
        ],
        "summary": "Get switch Firmware Version",
        "operationId": "switchVersion",
        "security": [
          {
            "Bearer": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Switch"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully returned switch firmware version"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/updateSwitch": {
      "post": {
        "description": "Update switch firmware based on specified switch type and firmware image",
        "tags": [
          "/updateSwitch"
        ],
        "summary": "Update switch firmware",
        "operationId": "updateSwitch",
        "security": [
          {
            "Bearer": []
          }
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UpdateSwitch"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Successfully issued update switch firmware",
            "schema": {
              "$ref": "#/definitions/Status"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "About": {
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "CheckVlan": {
      "type": "object",
      "required": [
        "endpoint",
        "vlanID"
      ],
      "properties": {
        "endpoint": {
          "$ref": "#/definitions/SwitchEndpoint"
        },
        "vlanID": {
          "type": "integer"
        }
      }
    },
    "ErrorResponse": {
      "required": [
        "message"
      ],
      "properties": {
        "errors": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "message": {
          "type": "string"
        }
      }
    },
    "FirmwareImage": {
      "type": "object",
      "required": [
        "imageType",
        "imageURL"
      ],
      "properties": {
        "imageType": {
          "type": "string",
          "enum": [
            "kickstart",
            "system"
          ]
        },
        "imageURL": {
          "type": "string"
        }
      }
    },
    "Login": {
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
    "LoginError": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Status": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Switch": {
      "type": "object",
      "required": [
        "endpoint"
      ],
      "properties": {
        "endpoint": {
          "$ref": "#/definitions/SwitchEndpoint"
        }
      }
    },
    "SwitchConfigResponse": {
      "type": "object",
      "properties": {
        "config": {
          "type": "string"
        }
      }
    },
    "SwitchEndpoint": {
      "type": "object",
      "required": [
        "ipaddress",
        "username",
        "password",
        "switchType"
      ],
      "properties": {
        "ipaddress": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "switchType": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "SwitchVersionResponse": {
      "type": "object",
      "properties": {
        "version": {
          "type": "string"
        }
      }
    },
    "Token": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "UpdateSwitch": {
      "type": "object",
      "required": [
        "endpoint",
        "firmwareImages",
        "switchModel"
      ],
      "properties": {
        "endpoint": {
          "$ref": "#/definitions/SwitchEndpoint"
        },
        "firmwareImages": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/FirmwareImage"
          }
        },
        "switchModel": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "description": "For accessing the API a valid JWT token must be passed in all the queries in\nthe 'Authorization' header.\n\n\nA valid JWT token is generated by the API and retourned as answer of a call\nto the route /login giving a valid user \u0026 password.\n\n\nThe following syntax must be used in the 'Authorization' header :\n\nBearer xxxxxx.yyyyyyy.zzzzzz\n",
      "type": "apiKey",
      "name": "authorization",
      "in": "header"
    }
  }
}`))
}
