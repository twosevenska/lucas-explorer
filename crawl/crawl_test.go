package crawl

import (
	"fmt"
	"testing"
)

type testEU struct {
	line    string
	address string
	ok      bool
}

func (expected testEU) isExpected(address string, ok bool) error {
	result := testEU{
		line:    expected.line,
		address: address,
		ok:      ok,
	}

	if expected != result {
		return fmt.Errorf("value mismatch. Expected VS Result\n    %+v\n    %+v\n", expected, result)
	}

	return nil
}

func TestExtractURL(t *testing.T) {
	cases := []testEU{
		testEU{
			line:    "<li><a href=\"http://info.cern.ch/hypertext/WWW/TheProject.html\">Browse the first website</a></li>",
			address: "http://info.cern.ch/hypertext/WWW/TheProject.html",
			ok:      true,
		},
		testEU{
			line:    "<p>From here you can:</p>",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "<li><a href=\"\">Browse the first website</a></li>",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "<meta charset=\"utf-8\">",
			address: "",
			ok:      false,
		},
		testEU{
			line:    "<a href=\"/issues\" class=\"issues\">Archives</a><span class=\"divider\">|</span>",
			address: "/issues",
			ok:      true,
		},
	}

	for _, c := range cases {
		a, ok := extractURL(c.line)
		err := c.isExpected(a, ok)
		if err != nil {
			t.Errorf("Failed assertion: \n %s", err.Error())
		}
	}
}

func TestExtractURLS(t *testing.T) {
	testCase := `
	<html><head></head><body><header>
	<title>http://info.cern.ch</title>
	</header>
	<h1>http://info.cern.ch - home of the first website</h1>
	<p>From here you can:</p>
	<ul>
	<li><a href="http://info.cern.ch/hypertext/WWW/TheProject.html">Browse the first website</a></li>
	<li><a href="http://line-mode.cern.ch/www/hypertext/WWW/TheProject.html">Browse the first website using the line-mode browser simulator</a></li>
	<li><a href="http://home.web.cern.ch/topics/birth-web">Learn about the birth of the web</a></li>
	<li><a href="http://home.web.cern.ch/about">Learn about CERN, the physics laboratory where the web was born</a></li>
	<li><a href="http://home.web.cern.ch/about">Learn about CERN, the physics laboratory where the web was born</a></li>
	<li><a HREF="http://home.web.cern.ch/about">Learn about CERN, the physics laboratory where the web was born</a></li>
	</ul>
	</body></html>
`

	expected := map[string]int{
		"http://info.cern.ch/hypertext/WWW/TheProject.html":          1,
		"http://line-mode.cern.ch/www/hypertext/WWW/TheProject.html": 1,
		"http://home.web.cern.ch/topics/birth-web":                   1,
		"http://home.web.cern.ch/about":                              3,
	}

	result := ExtractURLS(testCase, 2)

	for u, c := range expected {
		if len(result) != len(expected) {
			t.Errorf("Failed assertion: expected len different from result len: \n%+v\n%+v\n", expected, result)
		}

		if result[u] != c {
			t.Errorf("Failed assertion: expected len different from result len: \n%+v\n%+v\n", expected, result)
			t.Errorf("Failed assertion for %s: expected counter %d got %d", u, c, result[u])
		}
	}

}
