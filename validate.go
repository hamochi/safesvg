package safesvg

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

var svg_elements = map[string]struct{}{
	"svg":                 {},
	"altglyph":            {},
	"altglyphdef":         {},
	"altglyphitem":        {},
	"animatecolor":        {},
	"animatemotion":       {},
	"animatetransform":    {},
	"circle":              {},
	"clippath":            {},
	"defs":                {},
	"desc":                {},
	"ellipse":             {},
	"filter":              {},
	"font":                {},
	"g":                   {},
	"glyph":               {},
	"glyphref":            {},
	"hkern":               {},
	"image":               {},
	"line":                {},
	"lineargradient":      {},
	"marker":              {},
	"mask":                {},
	"metadata":            {},
	"mpath":               {},
	"path":                {},
	"pattern":             {},
	"polygon":             {},
	"polyline":            {},
	"radialgradient":      {},
	"rect":                {},
	"stop":                {},
	"switch":              {},
	"symbol":              {},
	"text":                {},
	"textpath":            {},
	"title":               {},
	"tref":                {},
	"tspan":               {},
	"use":                 {},
	"view":                {},
	"vkern":               {},
	"feBlend":             {},
	"feColorMatrix":       {},
	"feComponentTransfer": {},
	"feComposite":         {},
	"feConvolveMatrix":    {},
	"feDiffuseLighting":   {},
	"feDisplacementMap":   {},
	"feDistantLight":      {},
	"feFlood":             {},
	"feFuncA":             {},
	"feFuncB":             {},
	"feFuncG":             {},
	"feFuncR":             {},
	"feGaussianBlur":      {},
	"feMerge":             {},
	"feMergeNode":         {},
	"feMorphology":        {},
	"feOffset":            {},
	"fePointLight":        {},
	"feSpecularLighting":  {},
	"feSpotLight":         {},
	"feTile":              {},
	"feTurbulence":        {},
}

var svg_attributes = map[string]struct{}{
	"accent-height":               {},
	"accumulate":                  {},
	"additivive":                  {},
	"alignment-baseline":          {},
	"ascent":                      {},
	"attributename":               {},
	"attributetype":               {},
	"azimuth":                     {},
	"baseprofile":                 {},
	"basefrequency":               {},
	"baseline-shift":              {},
	"begin":                       {},
	"bias":                        {},
	"by":                          {},
	"class":                       {},
	"clip":                        {},
	"clip-path":                   {},
	"clip-rule":                   {},
	"color":                       {},
	"color-interpolation":         {},
	"color-interpolation-filters": {},
	"color-profile":               {},
	"color-rendering":             {},
	"cx":                          {},
	"cy":                          {},
	"d":                           {},
	"dx":                          {},
	"dy":                          {},
	"diffuseconstant":             {},
	"direction":                   {},
	"display":                     {},
	"divisor":                     {},
	"dur":                         {},
	"edgemode":                    {},
	"elevation":                   {},
	"end":                         {},
	"fill":                        {},
	"fill-opacity":                {},
	"fill-rule":                   {},
	"filter":                      {},
	"flood-color":                 {},
	"flood-opacity":               {},
	"font-family":                 {},
	"font-size":                   {},
	"font-size-adjust":            {},
	"font-stretch":                {},
	"font-style":                  {},
	"font-variant":                {},
	"font-weight":                 {},
	"fx":                          {},
	"fy":                          {},
	"g1":                          {},
	"g2":                          {},
	"glyph-name":                  {},
	"glyphref":                    {},
	"gradientunits":               {},
	"gradienttransform":           {},
	"height":                      {},
	"href":                        {},
	"id":                          {},
	"image-rendering":             {},
	"in":                          {},
	"in2":                         {},
	"k":                           {},
	"k1":                          {},
	"k2":                          {},
	"k3":                          {},
	"k4":                          {},
	"kerning":                     {},
	"keypoints":                   {},
	"keysplines":                  {},
	"keytimes":                    {},
	"lang":                        {},
	"lengthadjust":                {},
	"letter-spacing":              {},
	"kernelmatrix":                {},
	"kernelunitlength":            {},
	"lighting-color":              {},
	"local":                       {},
	"marker-end":                  {},
	"marker-mid":                  {},
	"marker-start":                {},
	"markerheight":                {},
	"markerunits":                 {},
	"markerwidth":                 {},
	"maskcontentunits":            {},
	"maskunits":                   {},
	"max":                         {},
	"mask":                        {},
	"media":                       {},
	"method":                      {},
	"mode":                        {},
	"min":                         {},
	"name":                        {},
	"numoctaves":                  {},
	"offset":                      {},
	"operator":                    {},
	"opacity":                     {},
	"order":                       {},
	"orient":                      {},
	"orientation":                 {},
	"origin":                      {},
	"overflow":                    {},
	"paint-order":                 {},
	"path":                        {},
	"pathlength":                  {},
	"patterncontentunits":         {},
	"patterntransform":            {},
	"patternunits":                {},
	"points":                      {},
	"preservealpha":               {},
	"preserveaspectratio":         {},
	"r":                           {},
	"rx":                          {},
	"ry":                          {},
	"radius":                      {},
	"refx":                        {},
	"refy":                        {},
	"repeatcount":                 {},
	"repeatdur":                   {},
	"restart":                     {},
	"result":                      {},
	"rotate":                      {},
	"scale":                       {},
	"seed":                        {},
	"shape-rendering":             {},
	"specularconstant":            {},
	"specularexponent":            {},
	"spreadmethod":                {},
	"stddeviation":                {},
	"stitchtiles":                 {},
	"stop-color":                  {},
	"stop-opacity":                {},
	"stroke-dasharray":            {},
	"stroke-dashoffset":           {},
	"stroke-linecap":              {},
	"stroke-linejoin":             {},
	"stroke-miterlimit":           {},
	"stroke-opacity":              {},
	"stroke":                      {},
	"stroke-width":                {},
	"style":                       {},
	"surfacescale":                {},
	"tabindex":                    {},
	"targetx":                     {},
	"targety":                     {},
	"transform":                   {},
	"text-anchor":                 {},
	"text-decoration":             {},
	"text-rendering":              {},
	"textlength":                  {},
	"type":                        {},
	"u1":                          {},
	"u2":                          {},
	"unicode":                     {},
	"version":                     {},
	"values":                      {},
	"viewbox":                     {},
	"visibility":                  {},
	"vert-adv-y":                  {},
	"vert-origin-x":               {},
	"vert-origin-y":               {},
	"width":                       {},
	"word-spacing":                {},
	"wrap":                        {},
	"writing-mode":                {},
	"xchannelselector":            {},
	"ychannelselector":            {},
	"x":                           {},
	"x1":                          {},
	"x2":                          {},
	"xmlns":                       {},
	"y":                           {},
	"y1":                          {},
	"y2":                          {},
	"z":                           {},
	"zoomandpan":                  {},

	"xlink:href":  {},
	"xml:id":      {},
	"xlink:title": {},
	"xml:space":   {},
	"xmlns:xlink": {},
}

// Validator is a struct with private variables for storing the whitelists
type Validator struct {
	whiteListElements   map[string]struct{}
	whiteListAttributes map[string]struct{}
}

// NewValidator creates a new validator with default whitelists
func NewValidator() Validator {
	vld := Validator{
		whiteListElements:   svg_elements,
		whiteListAttributes: svg_attributes,
	}
	return vld
}

// Validate validates a slice of bytes containing the svg data
func (vld Validator) Validate(b []byte) error {
	r := bytes.NewReader(b)
	return vld.ValidateReader(r)
}

// ValidateReader validates svg data from an io.Reader interface
func (vld Validator) ValidateReader(r io.Reader) error {
	t := xml.NewDecoder(r)
	var to xml.Token
	var err error

	for {
		to, err = t.Token()

		switch v := to.(type) {
		case xml.StartElement:
			if ok := validElements(v.Name.Local, vld.whiteListElements); !ok {
				return errors.New("Invalid element " + v.Name.Local)
			}

			if err := validAttributes(v.Attr, vld.whiteListAttributes); err != nil {
				return err
			}
		case xml.EndElement:
			if ok := validElements(v.Name.Local, vld.whiteListElements); !ok {
				return errors.New("Invalid element " + v.Name.Local)
			}
		case xml.CharData: //text

		case xml.Comment:

		case xml.ProcInst:

		case xml.Directive: //doctype etc

		}

		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return err
			}
		}

	}

	return nil
}

// WhitelistElements adds svg elements to the whitelist
func (vld *Validator) WhitelistElements(elements ...string) {
	for _, elemet := range elements {
		vld.whiteListElements[elemet] = struct{}{}
	}
}

// WhitelistAttributes adds svg attributes to the whitelist
func (vld *Validator) WhitelistAttributes(attributes ...string) {
	for _, attr := range attributes {
		vld.whiteListAttributes[attr] = struct{}{}
	}
}

// BlacklistElements removes svg elements from the whitelist
func (vld *Validator) BlacklistElements(elements ...string) {
	for _, elemet := range elements {
		delete(vld.whiteListElements, elemet)
	}
}

// BlacklistAttributes removes svg attributes from the whitelist
func (vld *Validator) BlacklistAttributes(attributes ...string) {
	for _, attr := range attributes {
		delete(vld.whiteListAttributes, attr)
	}
}

func validAttributes(attrs []xml.Attr, whiteListAttributes map[string]struct{}) error {
	var key string
	for _, attr := range attrs {
		if attr.Name.Space != "" {
			if attr.Name.Space == "http://www.w3.org/XML/1998/namespace" {
				attr.Name.Space = "xml"
			}
			key = attr.Name.Space + ":" + attr.Name.Local
		} else {
			key = attr.Name.Local
		}
		_, found := whiteListAttributes[strings.ToLower(key)]
		if !found {
			return errors.New("Invalid attribute " + attr.Name.Local)
		}
	}
	return nil
}

func validElements(elm string, whiteListElements map[string]struct{}) bool {
	_, found := whiteListElements[elm]
	return found
}
