/*
 * Copyright (c) 2025 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package huh

import (
	"encoding/xml"
	"errors"
	"reflect"
)

type nodeType int

const (
	nodeInvalid nodeType = iota
	nodeSelect
	nodeMultiSelect
	nodeNote
)

type node interface {
	Type() nodeType
}

// -----------------------------------------------------------------------------

type formDoc struct {
	Groups []*group `xml:"group"`
}

type noteNode struct {
	Title       string `xml:"title,attr"`
	Description string `xml:"description,attr"`
	Next        string `xml:"next,attr"`
}

func (n *noteNode) Type() nodeType {
	return nodeNote
}

type selectNode struct {
	ID      string        `xml:"id,attr"`
	Title   string        `xml:"title,attr"`
	Options []*optionNode `xml:"option"`
}

func (s *selectNode) Type() nodeType {
	return nodeSelect
}

type multiSelectNode struct {
	selectNode
	Limit int `xml:"limit,attr"`
}

func (ms *multiSelectNode) Type() nodeType {
	return nodeMultiSelect
}

type optionNode struct {
	Title    string `xml:"title,attr"`
	Value    string `xml:"value,attr"`
	Selected []byte `xml:"selected,attr"`
}

type group struct {
	Items []node
}

func (g *group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch se := token.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "select":
				var s selectNode
				if err := d.DecodeElement(&s, &se); err != nil {
					return err
				}
				g.Items = append(g.Items, &s)
			case "multiselect":
				var ms multiSelectNode
				if err := d.DecodeElement(&ms, &se); err != nil {
					return err
				}
				g.Items = append(g.Items, &ms)
			case "note":
				var n noteNode
				if err := d.DecodeElement(&n, &se); err != nil {
					return err
				}
				g.Items = append(g.Items, &n)
			default:
				return errors.New("unknown element " + se.Name.Local)
			}
		case xml.CharData: // ignore
		case xml.EndElement:
			if se.Name == start.Name {
				return nil
			}
			return errors.New("unexpected end element " + se.Name.Local)
		default:
			return errors.New("unexpected token " + reflect.TypeOf(token).String())
		}
	}
}

// -----------------------------------------------------------------------------
