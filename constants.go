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
[
  { "key": "api_type", "value": "js" },
  { "key": "p", "value": 1 },
  { "key": "f", "value": "0f9f23887d9eccedec3ebc075286a587" },
  { "key": "n", "value": "%s" },
  {
    "key": "wh",
    "value": "f7a6d54dc1e4357cfa1951136b3cba7d|72627afbfd19a741c7da1732218301ac"
  },
  {
    "key": "enhanced_fp",
    "value": [
      {
        "key": "webgl_extensions",
        "value": "ANGLE_instanced_arrays;EXT_blend_minmax;EXT_color_buffer_half_float;EXT_disjoint_timer_query;EXT_float_blend;EXT_frag_depth;EXT_shader_texture_lod;EXT_texture_compression_bptc;EXT_texture_compression_rgtc;EXT_texture_filter_anisotropic;EXT_sRGB;KHR_parallel_shader_compile;OES_element_index_uint;OES_fbo_render_mipmap;OES_standard_derivatives;OES_texture_float;OES_texture_float_linear;OES_texture_half_float;OES_texture_half_float_linear;OES_vertex_array_object;WEBGL_color_buffer_float;WEBGL_compressed_texture_etc;WEBGL_compressed_texture_etc1;WEBGL_compressed_texture_s3tc;WEBGL_compressed_texture_s3tc_srgb;WEBGL_debug_renderer_info;WEBGL_debug_shaders;WEBGL_depth_texture;WEBGL_draw_buffers;WEBGL_lose_context;WEBGL_multi_draw"
      },
      {
        "key": "webgl_extensions_hash",
        "value": "409becf3d3b4cb7713d51f30f35dabd4"
      },
      { "key": "webgl_renderer", "value": "WebKit WebGL" },
      { "key": "webgl_vendor", "value": "WebKit" },
      { "key": "webgl_version", "value": "WebGL 1.0 (OpenGL ES 2.0 Chromium)" },
      {
        "key": "webgl_shading_language_version",
        "value": "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)"
      },
      { "key": "webgl_aliased_line_width_range", "value": "[1, 7.375]" },
      { "key": "webgl_aliased_point_size_range", "value": "[1, 255]" },
      { "key": "webgl_antialiasing", "value": "yes" },
      { "key": "webgl_bits", "value": "8,8,24,8,8,0" },
      {
        "key": "webgl_max_params",
        "value": "16,64,16384,1024,16384,32,16384,32,16,32,1024"
      },
      { "key": "webgl_max_viewport_dims", "value": "[16384, 16384]" },
      { "key": "webgl_unmasked_vendor", "value": "Google Inc. (Intel)" },
      {
        "key": "webgl_unmasked_renderer",
        "value": "ANGLE (Intel, Mesa Intel(R) Xe Graphics (TGL GT2), OpenGL 4.6)"
      },
      {
        "key": "webgl_vsf_params",
        "value": "23,127,127,23,127,127,23,127,127"
      },
      { "key": "webgl_vsi_params", "value": "0,31,30,0,31,30,0,31,30" },
      {
        "key": "webgl_fsf_params",
        "value": "23,127,127,23,127,127,23,127,127"
      },
      { "key": "webgl_fsi_params", "value": "0,31,30,0,31,30,0,31,30" },
      {
        "key": "webgl_hash_webgl",
        "value": "b8bd8e5311f1c56b26eafcd4efc2d6c4"
      },
      { "key": "user_agent_data_brands", "value": "Not.A/Brand,Chromium" },
      { "key": "user_agent_data_mobile", "value": false },
      { "key": "navigator_connection_downlink", "value": 5.2 },
      { "key": "navigator_connection_downlink_max", "value": null },
      { "key": "network_info_rtt", "value": 50 },
      { "key": "network_info_save_data", "value": false },
      { "key": "network_info_rtt_type", "value": null },
      { "key": "screen_pixel_depth", "value": 24 },
      { "key": "navigator_device_memory", "value": 8 },
      { "key": "navigator_languages", "value": "en-US,en" },
      { "key": "window_inner_width", "value": 0 },
      { "key": "window_inner_height", "value": 0 },
      { "key": "window_outer_width", "value": 1128 },
      { "key": "window_outer_height", "value": 691 },
      { "key": "browser_detection_firefox", "value": false },
      { "key": "browser_detection_brave", "value": false },
      {
        "key": "audio_codecs",
        "value": "{\"ogg\":\"probably\",\"mp3\":\"probably\",\"wav\":\"probably\",\"m4a\":\"maybe\",\"aac\":\"probably\"}"
      },
      {
        "key": "video_codecs",
        "value": "{\"ogg\":\"probably\",\"h264\":\"probably\",\"webm\":\"probably\",\"mpeg4v\":\"\",\"mpeg4a\":\"\",\"theora\":\"\"}"
      },
      { "key": "media_query_dark_mode", "value": true },
      { "key": "headless_browser_phantom", "value": false },
      { "key": "headless_browser_selenium", "value": false },
      { "key": "headless_browser_nightmare_js", "value": false },
      { "key": "document__referrer", "value": "" },
      {
        "key": "window__ancestor_origins",
        "value": ["https://chat.openai.com"]
      },
      { "key": "window__tree_index", "value": [0] },
      { "key": "window__tree_structure", "value": "[[]]" },
      {
        "key": "window__location_href",
        "value": "https://tcr9i.chat.openai.com/v2/1.5.2/enforcement.%s.html#35536E1E-65B4-4D96-9D97-6ADB7EFF8147"
      },
      {
        "key": "client_config__sitedata_location_href",
        "value": "https://chat.openai.com/"
      },
      {
        "key": "client_config__surl",
        "value": "https://tcr9i.chat.openai.com"
      },
      { "key": "mobile_sdk__is_sdk" },
      { "key": "client_config__language", "value": null },
      { "key": "navigator_battery_charging", "value": false },
      { "key": "audio_fingerprint", "value": "124.04347527516074" }
    ]
  },
  {
    "key": "fe",
    "value": [
      "DNT:unknown",
      "L:en-US",
      "D:24",
      "PR:2",
      "S:1128,752",
      "AS:1128,724",
      "TO:-480",
      "SS:true",
      "LS:true",
      "IDB:true",
      "B:false",
      "ODB:true",
      "CPUC:unknown",
      "PK:Linux x86_64",
      "CFP:-1849310688",
      "FR:false",
      "FOS:false",
      "FB:false",
      "JSF:Andale Mono,Arial,Arial Black,Calibri,Cambria,Comic Sans MS,Courier,Courier New,Georgia,Helvetica,Impact,MS Gothic,MS PGothic,Times,Times New Roman,Trebuchet MS,Verdana",
      "P:Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF",
      "T:0,false,false",
      "H:8",
      "SWF:false"
    ]
  },
  { "key": "ife_hash", "value": "f184377c7522d556751f69cfdd420f8b" },
  { "key": "cs", "value": 1 },
  {
    "key": "jsbd",
    "value": "{\"HL\":5,\"NCE\":true,\"DT\":\"\",\"NWD\":\"false\",\"DOTO\":1,\"DMTO\":1}"
  }
]

`

const (
	p  = "Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,PDF Viewer,WebKit built-in PDF"
	bv = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
)

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
	cfp = "-1471542875"
)
