package currency

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return fmt.Errorf("currency value is empty")
	}

	normalized := strings.ReplaceAll(trimmed, ",", ".")
	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", raw, err)
	}
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return fmt.Errorf("currency value must be finite")
	}

	*cf = CommaFloat(val)
	return nil
}

type OptionalInt int

func (oi *OptionalInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		*oi = 0
		return nil
	}

	val, err := strconv.Atoi(trimmed)
	if err != nil {
		return fmt.Errorf("parse int %q: %w", raw, err)
	}

	*oi = OptionalInt(val)
	return nil
}

type Valute struct {
	NumCode  OptionalInt `xml:"NumCode"  json:"num_code"`
	CharCode string      `xml:"CharCode" json:"char_code"`
	Value    CommaFloat  `xml:"Value"    json:"value"`
}
