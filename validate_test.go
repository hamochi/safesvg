package safesvg_test

import (
	"github.com/hamochi/safesvg"
	"testing"
)

func Test_ValidSVGByte(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	err := v.Validate(svg)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

}

func Test_InvalidElements(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><script>window.alert('evil')</script><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}

}

func Test_WhiteListElements(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><newstuff></newstuff><script>window.alert('evil')</script><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	v.WhitelistElements("newstuff", "script")
	err := v.Validate(svg)
	if err != nil {
		t.Errorf("Unexptected error %v", err)
	}
}

func Test_WhiteListAttributes(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" stranger="things" foo="bar" bersion="2"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	v.WhitelistAttributes("stranger", "foo", "bersion")
	err := v.Validate(svg)
	if err != nil {
		t.Errorf("Unexptected error %v", err)
	}
}

func Test_BlackListElements(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	v.BlacklistElements("path")
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}
}

func Test_BlackListAttributes(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	v.BlacklistAttributes("xmlns", "width")
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}
}

func Test_InvalidAttributes(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" random="notvalid" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Expected validation error, got none")
	}

}

//https://en.wikipedia.org/wiki/Billion_laughs_attack
func Test_SVGBomb(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g id="a"><use/><use/><use/><use/><use/><use/><use/><use/><use/><use/></g><g id="b"><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/><use xlink:href="#a"/></g><g id="c"><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/><use xlink:href="#b"/></g><g id="d"><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/><use xlink:href="#c"/></g><g id="e"><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/><use xlink:href="#d"/></g><g id="f"><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/><use xlink:href="#e"/></g><g id="g"><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/><use xlink:href="#f"/></g><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/><use xlink:href="#g"/></svg>`)
	v := safesvg.NewValidator()
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Expected validation error, got none")
	}
}

func Test_Script(t *testing.T) {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><scRipt >alert("buu")</scRipt><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := safesvg.NewValidator()
	err := v.Validate(svg)
	if err == nil {
		t.Errorf("Expected validation error, got none")
	}
}
