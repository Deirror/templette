// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package element

type Tag string

const (
	HTML       Tag = "html"
	Head       Tag = "head"
	Body       Tag = "body"
	Header     Tag = "header"
	Main       Tag = "main"
	Footer     Tag = "footer"
	Button     Tag = "button"
	A          Tag = "a"
	Div        Tag = "div"
	Option     Tag = "option"
	Select     Tag = "select"
	Fieldset   Tag = "fieldset"
	Form       Tag = "form"
	Input      Tag = "input"
	Label      Tag = "label"
	Audio      Tag = "audio"
	Canvas     Tag = "canvas"
	Figcaption Tag = "figcaption"
	Figure     Tag = "figure"
	Img        Tag = "img"
	Source     Tag = "source"
	Video      Tag = "video"
	Meter      Tag = "meter"
	Progress   Tag = "progress"
	Article    Tag = "article"
	Details    Tag = "details"
	Hr         Tag = "hr"
	Nav        Tag = "nav"
	Section    Tag = "section"
	Summary    Tag = "summary"
	Table      Tag = "table"
	TBody      Tag = "tbody"
	Th         Tag = "th"
	Td         Tag = "td"
	THead      Tag = "thead"
	Tr         Tag = "tr"
	P          Tag = "p"
	H1         Tag = "h1"
	H2         Tag = "h2"
	H3         Tag = "h3"
	H4         Tag = "h4"
	H5         Tag = "h5"
	H6         Tag = "h6"
	Span       Tag = "span"
	Pre        Tag = "pre"
	Code       Tag = "code"
	Blockquote Tag = "blockquote"
	Ul         Tag = "ul"
	Ol         Tag = "ol"
	Li         Tag = "li"
	Dl         Tag = "dl"
	Dt         Tag = "dt"
	Dd         Tag = "dd"
	Strong     Tag = "strong"
	B          Tag = "b"
	Em         Tag = "em"
	I          Tag = "i"
	Small      Tag = "small"
	Mark       Tag = "mark"
	Del        Tag = "del"
	Abbr       Tag = "abbr"
	Ins        Tag = "ins"
	Kbd        Tag = "kbd"
	S          Tag = "s"
	U          Tag = "u"
	Sub        Tag = "sub"
	Sup        Tag = "sup"
	Textarea   Tag = "textarea"
	Br         Tag = "br"
)

var validTags = map[Tag]struct{}{
	Button: {}, A: {}, Div: {}, Option: {}, Select: {}, Fieldset: {}, Form: {}, Input: {}, Label: {},
	Footer: {}, Header: {}, Main: {}, Audio: {}, Canvas: {}, Figcaption: {}, Figure: {}, Img: {}, Source: {}, Video: {},
	Meter: {}, Progress: {}, Article: {}, Details: {}, Hr: {}, Nav: {}, Section: {}, Summary: {}, Table: {}, TBody: {},
	Th: {}, Td: {}, THead: {}, Tr: {}, P: {}, H1: {}, H2: {}, H3: {}, H4: {}, H5: {}, H6: {}, Span: {}, Pre: {}, Code: {},
	Blockquote: {}, Ul: {}, Ol: {}, Li: {}, Dl: {}, Dt: {}, Dd: {}, Strong: {}, B: {}, Em: {}, I: {}, Small: {}, Mark: {},
	Del: {}, Abbr: {}, Ins: {}, Kbd: {}, S: {}, U: {}, Sub: {}, Sup: {}, Textarea: {}, Br: {}, HTML: {}, Head: {}, Body: {},
}

func (t Tag) IsValid() bool {
	_, ok := validTags[t]
	return ok
}
