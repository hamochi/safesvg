package safesvg

import (
	"testing"
)

func Test_ValidSVGByte(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	err := v.Validate(validSvg)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

}

func Test_InvalidElements(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><script>window.alert('evil')</script><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	err := v.Validate(validSvg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}

}

func Test_WhiteListElements(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><newstuff></newstuff><script>window.alert('evil')</script><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	v.WhitelistElements("newstuff", "script")
	err := v.Validate(validSvg)
	if err != nil {
		t.Errorf("Unexptected error %v", err)
	}
}

func Test_WhiteListAttributes(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" stranger="things" foo="bar"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	v.WhitelistAttributes("stranger", "foo")
	err := v.Validate(validSvg)
	if err != nil {
		t.Errorf("Unexptected error %v", err)
	}
}

func Test_BlackListElements(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	v.BlacklistElements("path")
	err := v.Validate(validSvg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}
}

func Test_BlackListAttributes(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	v.BlacklistAttributes("xmlns", "width")
	err := v.Validate(validSvg)
	if err == nil {
		t.Errorf("Exptected validation error, got none")
	}
}

func Test_InvalidAttributes(t *testing.T) {
	validSvg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" random="notvalid" height="24" viewBox="0 0 24 24"><path fill="none" d="M0 0h24v24H0V0z"/><path d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/></svg>`)
	v := NewValidator()
	err := v.Validate(validSvg)
	if err == nil {
		t.Errorf("Expected validation error, got none")
	}

}
