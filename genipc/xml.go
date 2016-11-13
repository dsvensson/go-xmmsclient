package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type XmlValueType []string

type XmlReturnValue struct {
	Doc  string       `xml:"documentation"`
	Type XmlValueType `xml:"type"`
}

type XmlArgumentType struct {
	Content string `xml:",innerxml"`
}

type XmlArgument struct {
	Name string       `xml:"name"`
	Doc  string       `xml:"documentation"`
	Type XmlValueType `xml:"type"`
}

type XmlMethod struct {
	Name        string         `xml:"name"`
	Doc         string         `xml:"documentation"`
	Arguments   []XmlArgument  `xml:"argument"`
	ReturnValue XmlReturnValue `xml:"return_value"`
}

type XmlObject struct {
	Name    string      `xml:"name"`
	Methods []XmlMethod `xml:"method"`
}

type XmlEnum struct {
	Name    string   `xml:"name"`
	Hint    string   `xml:"namespace-hint"`
	Members []string `xml:"member"`
}

type Query struct {
	IpcVersion int         `xml:"version,attr"`
	Offset     int         `xml:"constant>value"`
	Objects    []XmlObject `xml:"object"`
	Enums      []XmlEnum   `xml:"enum"`
}

func (c *XmlValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var signature []string

	done := false
	for !done {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			if elem.Name.Local != "type" {
				signature = append(signature, elem.Name.Local)
				if len(elem.Attr) > 0 && elem.Attr[0].Name.Local == "name" {
					signature = append(signature, elem.Attr[0].Value)
				}
			}
		case xml.EndElement:
			if elem.Name.Local == "type" {
				done = true
			}
		}
	}

	*c = signature

	return nil
}

func parseAPI(filename string) (*Query, error) {
	var query Query

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &query)
	if err != nil {
		return nil, err
	}

	return &query, nil
}
