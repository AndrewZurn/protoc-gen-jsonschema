package testdata

const Proto3Required = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "query",
        "page_number"
    ],
    "properties": {
        "query": {
            "type": "string"
        },
        "page_number": {
            "type": "integer"
        },
        "result_per_page": {
            "type": "integer"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
