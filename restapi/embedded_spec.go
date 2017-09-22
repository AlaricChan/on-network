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
    "/updateSwitch": {
      "post": {
        "description": "Update switch firmware based on specified switch type and firmware image",
        "tags": [
          "/updateSwitch"
        ],
        "summary": "Update switch firmware",
        "operationId": "updateSwitch",
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
            "description": "Successfully issued update switch firmware"
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
    "Error": {
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
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
    "SwitchEndpoint": {
      "type": "object",
      "required": [
        "ip",
        "username",
        "password",
        "switchType"
      ],
      "properties": {
        "ip": {
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
    "UpdateSwitch": {
      "type": "object",
      "required": [
        "endpoint",
        "imageURL",
        "switchModel"
      ],
      "properties": {
        "endpoint": {
          "$ref": "#/definitions/SwitchEndpoint"
        },
        "imageURL": {
          "type": "string"
        },
        "switchModel": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "APIKeyHeader": []
    }
  ]
}`))
}
