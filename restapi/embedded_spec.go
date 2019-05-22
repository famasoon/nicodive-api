// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "nicodive-api",
    "version": "0.0.1"
  },
  "paths": {
    "/api/video/{id}": {
      "get": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object",
              "properties": {
                "commentCount": {
                  "type": "integer"
                },
                "description": {
                  "type": "string"
                },
                "embedurl": {
                  "type": "string"
                },
                "firstRetrive": {
                  "type": "string"
                },
                "length": {
                  "type": "string"
                },
                "mylistCount": {
                  "type": "integer"
                },
                "tags": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                },
                "title": {
                  "type": "string"
                },
                "viwCount": {
                  "type": "integer"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "nicodive-api",
    "version": "0.0.1"
  },
  "paths": {
    "/api/video/{id}": {
      "get": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object",
              "properties": {
                "commentCount": {
                  "type": "integer"
                },
                "description": {
                  "type": "string"
                },
                "embedurl": {
                  "type": "string"
                },
                "firstRetrive": {
                  "type": "string"
                },
                "length": {
                  "type": "string"
                },
                "mylistCount": {
                  "type": "integer"
                },
                "tags": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                },
                "title": {
                  "type": "string"
                },
                "viwCount": {
                  "type": "integer"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}