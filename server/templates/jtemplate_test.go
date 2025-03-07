package templates

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jitsucom/jitsu/server/events"
	"io/ioutil"
	"strings"
	"testing"
)

type templateTestData = struct {
	name      string
	expFormat string
	template  string
	event     events.Event
	expected  interface{}
}

var templateTest = []templateTestData{
	{"go1", "go", "{{if .metric_type }}{{ .metric_type }}_test{{else}}{{ .app }}_web_test {{end}}", events.Event{"metric_type": "ogon", "app": "neagon"}, "ogon_test"},
	{"go2", "go", "Hello {{ .metric_type }}", events.Event{"metric_type": "ogon", "app": "neagon"}, "Hello ogon"},
	{"go3", "go", `{{if or (eq .event_type "user_identify") (eq .event_type "identify")}}
          {{"identifies"}}
        {{else}}
          {{if or (eq .event_type "page") (eq .event_type "pageview")}}
            {{"pages"}}
          {{else}}
            {{.event_type}}
          {{end}}
        {{end}}`, events.Event{"event_type": "page"}, "pages"},
	{"go_json", "go", "{\n\t\"blocks\": [\n\t\t{\n\t\t\t\"type\": \"header\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"plain_text\",\n\t\t\t\t\"text\": \"🤘 {{.event_type}} from {{.user.email}}\",\n\t\t\t\t\"emoji\": true\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"User {{.user.email}} has sent *{{.event_type}}* in <https://exmaple.com?email={{urlquery .user.email}}>\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Data*:```{{ json_indent_quote . }}```\"\n\t\t\t}\n\t\t}\n\t]\n}",
		events.Event{"event_type": "page", "user": object{"email": "test@test.com"}}, "{\n\t\"blocks\": [\n\t\t{\n\t\t\t\"type\": \"header\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"plain_text\",\n\t\t\t\t\"text\": \"🤘 page from test@test.com\",\n\t\t\t\t\"emoji\": true\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"User test@test.com has sent *page* in <https://exmaple.com?email=test%40test.com>\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Data*:```{\\n  \\\"event_type\\\": \\\"page\\\",\\n  \\\"user\\\": {\\n    \\\"email\\\": \\\"test@test.com\\\"\\n  }\\n}```\"\n\t\t\t}\n\t\t}\n\t]\n}"},
	{"js_json", "javascript", "var koza = \"🤘\"\nreturn {\n\t\"blocks\": [\n\t\t{\n\t\t\t\"type\": \"header\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"plain_text\",\n\t\t\t\t\"text\": `${koza} ${$.event_type} from ${$.user.email}`,\n\t\t\t\t\"emoji\": true\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": `User ${$.user.email} has sent *${$.event_type}* in <https://example.com?email=${encodeURIComponent($.user?.email)}>`\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"type\": \"section\",\n\t\t\t\"text\": {\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Data*:```\" + json_indent_quote($) + \"```\"\n\t\t\t}\n\t\t}\n\t]\n}",
		events.Event{"event_type": "page", "user": object{"email": "test@test.com"}},
		object{"blocks": []interface{}{object{"text": object{"emoji": true, "text": "🤘 page from test@test.com", "type": "plain_text"}, "type": "header"}, object{"text": object{"text": "User test@test.com has sent *page* in <https://example.com?email=test%40test.com>", "type": "mrkdwn"}, "type": "section"}, object{"text": object{"text": "*Data*:```{\\n  \\\"event_type\\\": \\\"page\\\",\\n  \\\"user\\\": {\\n    \\\"email\\\": \\\"test@test.com\\\"\\n  }\\n}```", "type": "mrkdwn"}, "type": "section"}}}},
	{"go_and_js_error", "error", "if (true) {{return func(){ .metric_type }}}", events.Event{"metric_type": "ogon", "app": "neagon"}, fmt.Errorf("2 errors occurred")},
	{"golike_js", "javascript", "if (true) {var a = function(){ metric_type }} return $.metric_type", events.Event{"metric_type": "ogon", "app": "neagon"}, "ogon"},
	{"js_load_error", "javascript", "class Rectangle {constructor() {}}", events.Event{"metric_type": "ogon", "app": "neagon"}, fmt.Errorf("JS LOADING ERROR")},
	{"js_throw", "javascript", "throw new Error(\"test_error_throw\");", events.Event{"metric_type": "ogon", "app": "neagon"}, fmt.Errorf("javascript error: Error: test_error_throw")},

	{"const: \"table_name\"", "constant", "\"table_name\"", events.Event{"metric_type": "ogon"}, "\"table_name\""},
	{"const: table_name", "constant", "table_name", events.Event{"metric_type": "ogon"}, "table_name"},
	{"const: table name (err)", "constant", "table name", events.Event{"metric_type": "ogon"}, fmt.Errorf("javascript error: ReferenceError")},
	{"const: \"table name\"", "constant", "\"table name\"", events.Event{"metric_type": "ogon"}, "\"table name\""},
	{"const: data_base.table1", "constant", "data_base.table1", events.Event{"metric_type": "ogon"}, "data_base.table1"},
	{"const: \"data base\".\"table1\"", "constant", "\"data base\".\"table1\"", events.Event{"metric_type": "ogon"}, "\"data base\".\"table1\""},
	{"const: url", "constant", "https://example.com/123/abc123", events.Event{"metric_type": "ogon"}, "https://example.com/123/abc123"},

	{"infinite loop", "javascript", "var i; while (1) {i++}", events.Event{"metric_type": "ogon"}, fmt.Errorf("javascript error: Error: Infinite(or too long) loop break on the line")},
	{"infinite recursion", "javascript", "var emailRegexp = /^$/\nfunction removeEmails(obj) {\n\tfor (const key in obj) {\n\t\tif (typeof obj[key] === \"object\") {\n\t\t\tremoveEmails($)\n\t\t} else if (typeof obj[key] === \"string\" && obj[key].match(emailRegexp)) { \n\t\t\tdelete obj[key]\n\t\t}\n\t}\n}\nremoveEmails($)\nreturn $", events.Event{"object": object{}}, fmt.Errorf("javascript error: <nil>")},
}

func TestJtemplate(t *testing.T) {
	for _, tt := range templateTest {
		t.Run(tt.name, func(t *testing.T) {
			test(t, tt)
		})
	}
}

func TestJavascriptTemplate(t *testing.T) {
	for _, jstt := range JSTemplateTest {
		file, err := ioutil.ReadFile(jstt.filename)
		if err != nil {
			t.Errorf("cannot open file with test data %s : %v", jstt.filename, err)
			continue
		}
		var tt = templateTestData{name: jstt.filename, expFormat: "javascript", template: string(file), event: jstt.input, expected: jstt.expected}
		t.Run(tt.name, func(t *testing.T) {
			test(t, tt)
		})
	}
}

func test(t *testing.T, data templateTestData) {
	t.Logf("Test %v. Template:\n%v\nInput: %v\nExpected: %v", data.name, data.template, data.event, data.expected)
	templateExecutor, err := SmartParse(data.name, data.template, JSONSerializeFuncs)
	if err != nil {
		testExpectedError(t, data, err)
		return
	}
	t.Logf("Format %s", templateExecutor.Format())
	value, err := templateExecutor.ProcessEvent(data.event)
	if err != nil {
		testExpectedError(t, data, err)
		return
	}
	if data.expFormat != templateExecutor.Format() {
		t.Errorf("Format %v != expected: %v", templateExecutor.Format(), data.expFormat)
	}
	if !cmp.Equal(value, data.expected) {
		t.Errorf("Not equals. %v != expected: %v\nDiff:%v", value, data.expected, cmp.Diff(value, data.expected))
	} else {
		t.Logf("%s", value)
	}
}

func testExpectedError(t *testing.T, data templateTestData, err error) {
	switch expErr := data.expected.(type) {
	case error:
		if strings.HasPrefix(err.Error(), expErr.Error()) {
			t.Logf("Expected error: %v", err)
		} else {
			t.Errorf("unexpected error: %v\n Expected: %v", err, data.expected)
		}
	default:
		t.Error(err)
	}
}
