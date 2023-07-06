package funcaptcha

import http "github.com/bogdanfinn/fhttp"

var headers = http.Header{
	"Accept":           []string{"*/*"},
	"Accept-Encoding":  []string{"gzip, deflate, br"},
	"Accept-Language":  []string{"en-US,en;q=0.5"},
	"Cache-Control":    []string{"no-cache"},
	"Connection":       []string{"keep-alive"},
	"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
	"Host":             []string{"client-api.arkoselabs.com"},
	"Origin":           []string{"https://client-api.arkoselabs.com"},
	"User-Agent":       []string{bv},
	"X-Requested-With": []string{"XMLHttpRequest"},
}

const bx_template string = `
[{
	"key": "api_type",
	"value": "js"
}, {
	"key": "p",
	"value": 1
}, {
	"key": "f",
	"value": "%s"
}, {
	"key": "n",
	"value": "%s"
}, {
	"key": "wh",
	"value": "%s"
}, {
	"key": "enhanced_fp",
	"value": [{
		"key": "webgl_extensions",
		"value": "%s"
	}, {
		"key": "webgl_extensions_hash",
		"value": "%s"
	}, {
		"key": "webgl_renderer",
		"value": "%s"
	}, {
		"key": "webgl_vendor",
		"value": "%s"
	}, {
		"key": "webgl_version",
		"value": "%s"
	}, {
		"key": "webgl_shading_language_version",
		"value": "%s"
	}, {
		"key": "webgl_aliased_line_width_range",
		"value": "%s"
	}, {
		"key": "webgl_aliased_point_size_range",
		"value": "%s"
	}, {
		"key": "webgl_antialiasing",
		"value": "%s"
	}, {
		"key": "webgl_bits",
		"value": "%s"
	}, {
		"key": "webgl_max_params",
		"value": "%s"
	}, {
		"key": "webgl_max_viewport_dims",
		"value": "%s"
	}, {
		"key": "webgl_unmasked_vendor",
		"value": "%s"
	}, {
		"key": "webgl_unmasked_renderer",
		"value": "%s"
	}, {
		"key": "webgl_vsf_params",
		"value": "%s"
	}, {
		"key": "webgl_vsi_params",
		"value": "%s"
	}, {
		"key": "webgl_fsf_params",
		"value": "%s"
	}, {
		"key": "webgl_fsi_params",
		"value": "%s"
	}, {
		"key": "webgl_hash_webgl",
		"value": "%s"
	}, {
		"key": "user_agent_data_brands",
		"value": "Not.A/Brand,Chromium,Google Chrome"
	}, {
		"key": "user_agent_data_mobile",
		"value": false
	}, {
		"key": "navigator_connection_downlink",
		"value": 10.0
	}, {
		"key": "navigator_connection_downlink_max",
		"value": null
	}, {
		"key": "network_info_rtt",
		"value": 150
	}, {
		"key": "network_info_save_data",
		"value": false
	}, {
		"key": "network_info_rtt_type",
		"value": null
	}, {
		"key": "screen_pixel_depth",
		"value": 24
	}, {
		"key": "navigator_device_memory",
		"value": 8
	}, {
		"key": "navigator_languages",
		"value": "en-US"
	}, {
		"key": "window_inner_width",
		"value": 0
	}, {
		"key": "window_inner_height",
		"value": 0
	}, {
		"key": "window_outer_width",
		"value": 1920
	}, {
		"key": "window_outer_height",
		"value": 1057
	}, {
		"key": "browser_detection_firefox",
		"value": false
	}, {
		"key": "browser_detection_brave",
		"value": false
	}, {
		"key": "audio_codecs",
		"value": "{\"ogg\":\"probably\",\"mp3\":\"probably\",\"wav\":\"probably\",\"m4a\":\"maybe\",\"aac\":\"probably\"}"
	}, {
		"key": "video_codecs",
		"value": "{\"ogg\":\"probably\",\"h264\":\"probably\",\"webm\":\"probably\",\"mpeg4v\":\"\",\"mpeg4a\":\"\",\"theora\":\"\"}"
	}, {
		"key": "media_query_dark_mode",
		"value": true
	}, {
		"key": "headless_browser_phantom",
		"value": false
	}, {
		"key": "headless_browser_selenium",
		"value": false
	}, {
		"key": "headless_browser_nightmare_js",
		"value": false
	}, {
		"key": "document__referrer",
		"value": ""
	}, {
		"key": "window__ancestor_origins",
		"value": ["https://chat.openai.com"]
	}, {
		"key": "window__tree_index",
		"value": [2]
	}, {
		"key": "window__tree_structure",
		"value": "[[],[],[]]"
	}, {
		"key": "window__location_href",
		"value": "https://tcr9i.chat.openai.com/v2/1.5.2/enforcement.%s.html#35536E1E-65B4-4D96-9D97-6ADB7EFF8147"
	}, {
		"key": "client_config__sitedata_location_href",
		"value": "https://chat.openai.com"
	}, {
		"key": "client_config__surl",
		"value": "https://tcr9i.chat.openai.com"
	}, {
		"key": "mobile_sdk__is_sdk"
	}, {
		"key": "client_config__language",
		"value": null
	}, {
		"key": "navigator_battery_charging",
		"value": true
	}, {
		"key": "audio_fingerprint",
		"value": "124.04347527516073"
	}]
}, {
	"key": "fe",
	"value": %s
}, {
	"key": "ife_hash",
	"value": "%s"
}, {
	"key": "cs",
	"value": 1
}, {
	"key": "jsbd",
	"value": "{\"HL\":5,\"NCE\":true,\"DT\":\"\",\"NWD\":\"false\",\"DOTO\":1,\"DMTO\":1}"
}]
`

// LinkedHashMap
var fe = []map[string]interface{}{
	{"DNT": "1"},
	{"L": "en-US"},
	{"D": 24},
	{"PR": 1}, // this.getPixelRatio();
	{"S": "0;0"},
	{"AS": false},
	{"TO": 0},
	{"SS": true},
	{"LS": true},
	{"IDB": true},
	{"B": false},
	{"ODB": true},
	{"CPUC": "unknown"},
	{"PK": "Linux x86_64"},
	{"CFP": cfp},
	{"FR": false},
	{"FOS": false},
	{"FB": false},
	{"JSF": "Arial,Arial Narrow,Bitstream Vera Sans Mono,Bookman Old Style,Century Schoolbook,Courier,Courier New,Helvetica,MS Gothic,MS PGothic,Palatino,Palatino Linotype,Times,Times New Roman"},
	{"P": p},
	{"T": "0;false;false"},
	{"H": 2},
	{"SWF": false},
}

const (
	p  = "Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF"
	bv = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
)

const (
	cfp = "-1471542875"
)
