package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// testing for Get method
func TestRequestGet(t *testing.T) {
	caseEmpty := "/"
	caseSetHeader := "/set_header"

	// run fake http sever to accept requests
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != http.MethodGet {
			t.Errorf("Expected method %q; got %q", http.MethodGet, r.Method)
		}
		if r.Header == nil {
			t.Error("Expected non-nil SuperAgent Header")
		}
		switch r.URL.Path {
		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case caseEmpty:
			t.Logf("case %v ", caseEmpty)
		case caseSetHeader:
			t.Logf("case %v ", caseSetHeader)
			if r.Header.Get("Header1") != "foo-key" {
				t.Errorf("Expected 'Header1' == %q; got %q", "foo-key", r.Header.Get("Header1"))
			}
		}
	}))

	defer ts.Close()

	if err := New().Get(ts.URL+caseEmpty, nil); err != nil {
		t.Errorf("Expected NoError %v", err)
	}

	hdr := make(Header)
	hdr["Header1"] = "foo-key"

	if err := New().Get(ts.URL+caseSetHeader, nil, hdr); err != nil {
		t.Errorf("Expected NoError %v", err)
	}
}
