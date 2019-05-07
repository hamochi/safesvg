package safesvg

import (
	"bytes"
	"encoding/xml"
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
}

// ValidateBytes validates a slice of bytes containing the svg data
func ValidateBytes(b []byte) (bool, error) {
	r := bytes.NewReader(b)
	return ValidateReader(r)
}

// ValidateString validates a string containing the svg data
func ValidateString(s string) (bool, error) {
	r := strings.NewReader(s)
	return ValidateReader(r)
}

// ValidateReader validates svg data from a io.Reader interface
func ValidateReader(r io.Reader) (bool, error) {
	t := xml.NewDecoder(r)
	var to xml.Token
	var err error

	for {
		to, err = t.Token()

		switch v := to.(type) {
		case xml.StartElement:
			if ok := validElements(v.Name.Local); !ok {
				return false, nil
			}

			if ok := validAttributes(v.Attr); !ok {
				return false, nil
			}
		case xml.EndElement:
			if ok := validElements(v.Name.Local); !ok {
				return false, nil
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
				return false, err
			}
		}

	}

	return true, nil
}

func validAttributes(attrs []xml.Attr) bool {
	for _, attr := range attrs {
		_, found := svg_attributes[attr.Name.Local]
		return found
	}
	return false
}

func validElements(elm string) bool {
	_, found := svg_elements[elm]
	return found
}