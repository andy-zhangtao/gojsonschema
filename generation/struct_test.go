/*
 *   Copyright (c) 2017. The gojsonschema Author ztao8607@gmail.com
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 *
 */

package generation

import (
	"testing"
	"github.com/andy-zhangtao/gojsonschema/parse"
)

func TestGenerate(t *testing.T) {
	jsons := `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "description": "",
    "type": "object",
    "properties": {
        "info": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "minLength": 1
                },
                "display_id": {
                    "type": "string",
                    "minLength": 1
                },
                "duration": {
                    "type": "number"
                },
                "ext": {
                    "type": "string",
                    "minLength": 1
                },
                "extractor": {
                    "type": "string",
                    "minLength": 1
                },
                "extractor_key": {
                    "type": "string",
                    "minLength": 1
                },
                "filesize": {
                    "type": "number"
                },
                "format": {
                    "type": "string",
                    "minLength": 1
                },
                "format_id": {
                    "type": "string",
                    "minLength": 1
                },
                "formats": {
                    "type": "array",
                    "uniqueItems": true,
                    "minItems": 1,
                    "items": {
                        "required": [
                            "ext",
                            "format",
                            "format_id",
                            "preference",
                            "protocol",
                            "url",
                            "http_headers"
                        ],
                        "properties": {
                            "ext": {
                                "type": "string",
                                "minLength": 1
                            },
                            "format": {
                                "type": "string",
                                "minLength": 1
                            },
                            "format_id": {
                                "type": "string",
                                "minLength": 1
                            },
                            "http_headers": {
                                "type": "object",
                                "properties": {
                                    "Accept": {
                                        "type": "string",
                                        "minLength": 1
                                    },
                                    "Accept-Charset": {
                                        "type": "string",
                                        "minLength": 1
                                    },
                                    "Accept-Encoding": {
                                        "type": "string",
                                        "minLength": 1
                                    },
                                    "Accept-Language": {
                                        "type": "string",
                                        "minLength": 1
                                    },
                                    "Referer": {
                                        "type": "string",
                                        "minLength": 1
                                    },
                                    "User-Agent": {
                                        "type": "string",
                                        "minLength": 1
                                    }
                                },
                                "required": [
                                    "Accept",
                                    "Accept-Charset",
                                    "Accept-Encoding",
                                    "Accept-Language",
                                    "Referer",
                                    "User-Agent"

                                ]
                            },
                            "preference": {
                                "type": "number"
                            },
                            "protocol": {
                                "type": "string",
                                "minLength": 1
                            },
                            "url": {
                                "type": "string",
                                "minLength": 1
                            }
                        }
                    }
                },
                "http_headers": {
                    "type": "object",
                    "properties": {
                        "Accept": {
                            "type": "string",
                            "minLength": 1
                        },
                        "Accept-Charset": {
                            "type": "string",
                            "minLength": 1
                        },
                        "Accept-Encoding": {
                            "type": "string",
                            "minLength": 1
                        },
                        "Accept-Language": {
                            "type": "string",
                            "minLength": 1
                        },
                        "Referer": {
                            "type": "string",
                            "minLength": 1
                        },
                        "User-Agent": {
                            "type": "string",
                            "minLength": 1
                        }
                    },
                    "required": [
                        "Accept",
                        "Accept-Charset",
                        "Accept-Encoding",
                        "Accept-Language",
                        "Referer",
                        "User-Agent"
                    ]
                },
                "id": {
                    "type": "string",
                    "minLength": 1
                },
                "playlist": {},
                "playlist_index": {},
                "protocol": {
                    "type": "string",
                    "minLength": 1
                },
                "requested_subtitles": {},
                "thumbnail": {
                    "type": "string",
                    "minLength": 1
                },
                "thumbnails": {
                    "type": "array",
                    "uniqueItems": true,
                    "minItems": 1,
                    "items": {
                        "required": [
                            "id",
                            "url"
                        ],
                        "properties": {
                            "id": {
                                "type": "string",
                                "minLength": 1
                            },
                            "url": {
                                "type": "string",
                                "minLength": 1
                            }
                        }
                    }
                },
                "timestamp": {
                    "type": "number"
                },
                "title": {
                    "type": "string",
                    "minLength": 1
                },
                "upload_date": {
                    "type": "string",
                    "minLength": 1
                },
                "uploader": {
                    "type": "string",
                    "minLength": 1
                },
                "uploader_id": {
                    "type": "string",
                    "minLength": 1
                },
                "url": {
                    "type": "string",
                    "minLength": 1
                },
                "webpage_url": {
                    "type": "string",
                    "minLength": 1
                },
                "webpage_url_basename": {
                    "type": "string",
                    "minLength": 1
                }
            },
            "required": [
                "description",
                "display_id",
                "duration",
                "ext",
                "extractor",
                "extractor_key",
                "filesize",
                "format",
                "format_id",
                "formats",
                "http_headers",
                "id",
                "protocol",
                "thumbnail",
                "thumbnails",
                "timestamp",
                "title",
                "upload_date",
                "uploader",
                "uploader_id",
                "url",
                "webpage_url",
                "webpage_url_basename"
            ]
        },
        "url": {
            "type": "string",
            "minLength": 1
        }
    },
    "required": [
        "info",
        "url"
    ]
}`


	mapStruct, mapTag, err := parse.ParseJson("GOGO",jsons)
	if err != nil{
		t.Error(err.Error())
	}

	data := Generate("GOGO", mapStruct, mapTag)

	if len(data) != 1784 {
		t.Error("Generate Error")
	}
}
