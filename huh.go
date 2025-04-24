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
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

// Form is a collection of groups that are displayed one at a time on a "page".
//
// The form can navigate between groups and is complete once all the groups are
// complete.
type Form struct {
	*huh.Form
}

func New(text string, ret any) (_ Form, err error) {
	var doc formDoc
	d := xml.NewDecoder(strings.NewReader(text))
	d.Strict = false
	err = d.Decode(&doc)
	if err != nil {
		return
	}
	var groups []*huh.Group
	for _, g := range doc.Groups {
		var field huh.Field
		var fields = make([]huh.Field, 0, len(g.Items))
		for _, item := range g.Items {
			if field, err = newField(item, ret); err != nil {
				return
			}
			fields = append(fields, field)
		}
		groups = append(groups, huh.NewGroup(fields...))
	}
	return Form{huh.NewForm(groups...)}, nil
}

// -----------------------------------------------------------------------------

func newField(item node, ret any) (huh.Field, error) {
	switch item := item.(type) {
	case *selectNode:
		s, e := newSelect(item.ID, ret)
		if e != nil {
			return nil, e
		}
		return s, initSelect(s, item)
	case *multiSelectNode:
		ms, e := newMultiSelect(item.ID, ret)
		if e != nil {
			return nil, e
		}
		return ms, initMultiSelect(ms, item)
	default:
		return nil, errors.New("unknown node type")
	}
}

// -----------------------------------------------------------------------------

type huhOptions interface {
	AddOption(opt *optionNode) error
}

type optionsWrap[T comparable] []huh.Option[T]

func (p *optionsWrap[T]) AddOption(opt *optionNode) (err error) {
	var value T
	if err = initValue(&value, opt.Value); err != nil {
		return
	}
	title := opt.Title
	if title == "" {
		title = opt.Value
	}
	option := huh.NewOption(title, value)
	if opt.Selected != nil {
		option = option.Selected(toBool(opt.Selected, "selected"))
	}
	*p = append(*p, option)
	return
}

func toBool(val []byte, attr string) bool {
	switch string(val) {
	case attr, "true":
		return true
	case "false", "":
		return false
	}
	panic("invalid bool attribute " + attr)
}

func initValue(pv any, val string) (err error) {
	switch pv := pv.(type) {
	case *string:
		*pv = val
	case *int:
		*pv, err = strconv.Atoi(val)
	default:
		err = errors.New("initValue unsupported type: " + reflect.TypeOf(pv).String())
	}
	return
}

type huhSelect interface {
	huh.Field
	NewOptions() huhOptions
	Options(options huhOptions)
	Title(title string)
}

type huhMultiSelect interface {
	huhSelect
	Limit(n int)
}

func initSelect(s huhSelect, val *selectNode) (err error) {
	if val.Title != "" {
		s.Title(val.Title)
	}
	options := s.NewOptions()
	for _, opt := range val.Options {
		if err = options.AddOption(opt); err != nil {
			return
		}
	}
	s.Options(options)
	return
}

func initMultiSelect(ms huhMultiSelect, val *multiSelectNode) (err error) {
	if val.Limit > 0 {
		ms.Limit(val.Limit)
	}
	return initSelect(ms, &val.selectNode)
}

// -----------------------------------------------------------------------------

func newSelect(id string, ret any) (huhSelect, error) {
	v := reflect.ValueOf(ret)
	fld := v.Elem().FieldByName(id)
	if !fld.IsValid() {
		return nil, errors.New("invalid field " + id)
	}
	pfld := fld.Addr().Interface()
	switch kind := fld.Kind(); kind {
	case reflect.String:
		return newSelectWrap[string]().Value(pfld), nil
	case reflect.Int:
		return newSelectWrap[int]().Value(pfld), nil
	default:
		panic("Select: unsupported field type " + kind.String())
	}
}

type selectWrap[T comparable] struct {
	*huh.Select[T]
}

func newSelectWrap[T comparable]() selectWrap[T] {
	return selectWrap[T]{huh.NewSelect[T]()}
}

func (p selectWrap[T]) NewOptions() huhOptions {
	return new(optionsWrap[T])
}

func (p selectWrap[T]) Title(title string) {
	p.Select.Title(title)
}

func (p selectWrap[T]) Options(options huhOptions) {
	p.Select.Options(*options.(*optionsWrap[T])...)
}

func (p selectWrap[T]) Value(value any) huhSelect {
	p.Select.Value(value.(*T))
	return p
}

// -----------------------------------------------------------------------------

func newMultiSelect(id string, ret any) (huhMultiSelect, error) {
	v := reflect.ValueOf(ret)
	fld := v.Elem().FieldByName(id)
	if !fld.IsValid() {
		return nil, errors.New("invalid field " + id)
	}
	pfld := fld.Addr().Interface()
	switch kind := fld.Type().Elem().Kind(); kind {
	case reflect.String:
		return newMultiSelectWrap[string]().Value(pfld), nil
	case reflect.Int:
		return newMultiSelectWrap[int]().Value(pfld), nil
	default:
		panic("MultiSelect: unsupported field type " + kind.String())
	}
}

type multiSelectWrap[T comparable] struct {
	*huh.MultiSelect[T]
}

func newMultiSelectWrap[T comparable]() multiSelectWrap[T] {
	return multiSelectWrap[T]{huh.NewMultiSelect[T]()}
}

func (p multiSelectWrap[T]) NewOptions() huhOptions {
	return new(optionsWrap[T])
}

func (p multiSelectWrap[T]) Title(title string) {
	p.MultiSelect.Title(title)
}

func (p multiSelectWrap[T]) Options(options huhOptions) {
	p.MultiSelect.Options(*options.(*optionsWrap[T])...)
}

func (p multiSelectWrap[T]) Limit(n int) {
	p.MultiSelect.Limit(n)
}

func (p multiSelectWrap[T]) Value(value any) huhMultiSelect {
	p.MultiSelect.Value(value.(*[]T))
	return p
}

// -----------------------------------------------------------------------------
