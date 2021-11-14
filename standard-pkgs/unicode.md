# unicode
包含基本的字符判断函数
* unicode.IsControl(r rune) bool  // 是否控制字符
* unicode.IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 0-9
* unicode.IsGraphic(r rune) bool // 是否图形字符
* unicode.IsLetter(r rune) bool // 是否字母
* unicode.IsLower(r rune) bool // 是否小写字符
* unicode.IsMark(r rune) bool // 是否符号字符
* unicode.IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
* unicode.IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个
* unicode.IsPrint(r rune) bool // 是否可打印字符
* unicode.IsPunct(r rune) bool // 是否标点符号
* unicode.IsSpace(r rune) bool // 是否空格
* unicode.IsSymbol(r rune) bool // 是否符号字符
* unicode.IsTitle(r rune) bool // 是否 title case
* unicode.IsUpper(r rune) bool // 是否大写字符
* unicode.Is(rangeTab *RangeTable, r rune) bool // r 是否为 rangeTab 类型的字符
* unicode.In(r rune, ranges ...*RangeTable) bool  // r 是否为 ranges 中任意一个类型的字符












# unicode/utf8
主要负责 rune 和 byte 之间的转换


# unicode/utf16
负责 rune 和 uint16 数组之间的转换


