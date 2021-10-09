package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func TestString(t *testing.T) {
	t.Log(strings.Compare("abc", "dsc")) // -1
	t.Log(strings.Compare("abc", "abc")) // 0

	t.Log(strings.Contains("asdfg", "as")) // true
	t.Log(strings.Contains("asdfg", "ase")) // false

	t.Log(strings.ContainsAny("asdfg", "as")) // true
	t.Log(strings.ContainsAny("asdfg", "ase")) // true
	t.Log(strings.ContainsAny("asdfg", "rew")) // false

	t.Log(strings.Count("asdasdasd", "da")) // 2
	t.Log(strings.Count("ydededede", "ded")) // 2

	t.Log(strings.Fields("This is    my    str.field")) // [This is my str.field]
	
	t.Log(strings.FieldsFunc("This1is2a3first.", func(r rune) bool { // [This is a first.]
		return unicode.IsNumber(r)
	}))

	t.Log(strings.Split("This,is,string.", ",")) // [This is string.]
	t.Log(strings.SplitAfter("This,is,string.", ",")) // [This, is, string.]
	t.Log(strings.SplitN("This,is,string.", ",", 2)) // [This is,string.]
	t.Log(strings.SplitN("This,is,string.", ",", 0)) // []
	t.Log(strings.SplitN("This,is,string.", ",", -1)) // [This is string.]

	t.Log(strings.HasPrefix("Gopher", "Go")) // true
	t.Log(strings.HasPrefix("Gopher", "go")) // false
	t.Log(strings.HasPrefix("Gopher", "")) // true
	t.Log(strings.HasSuffix("Gopher", "er")) // true
	t.Log(strings.HasSuffix("Gopher", "ph")) // false
	t.Log(strings.HasSuffix("Gopher", "")) // true

	t.Log(strings.Index("Gopherhe", "he")) // 3
	t.Log(strings.Index("Gopher", "hh")) // -1
	t.Log(strings.IndexAny("Gopher", "hr")) // 3
	t.Log(strings.IndexFunc("Goph123er", func(r rune) bool { // 4
		return unicode.IsNumber(r)
	}))
	t.Log(strings.LastIndex("Gopherhe", "he")) // 6

	t.Log(strings.Join([]string{"This", "is", "join."}, " ")) // This is join.

	t.Log(strings.Repeat("Go", 3)) // GoGoGo

	t.Log(strings.Replace("veveveveve", "ve", "sa", 3)) // sasasaveve
	t.Log(strings.Replace("veveveveve", "ve", "sa", 0)) // veveveveve
	t.Log(strings.Replace("veveveveve", "ve", "sa", -2)) // sasasasasa
	t.Log(strings.ReplaceAll("veveveveve", "ve", "sa")) // sasasasasa

	t.Log(strings.ToUpper("name")) // NAME
	t.Log(strings.ToLower("VJUDGE")) // vjudge
	// 返回字符串s的副本，其中所有Unicode字母均使用c指定的大小写映射映射到其大写字母。
	t.Log(strings.ToLowerSpecial(unicode.TurkishCase, "China中国")) // china中国

	t.Log(strings.Title("thIs is my tiTle")) // ThIs Is My TiTle
	t.Log(strings.ToTitle("thIs is my tiTle")) // THIS IS MY TITLE

	t.Log(strings.Trim("   this  is    trim .    ", " ")) // this  is    trim .
	t.Log(strings.TrimLeft("   this  is    trim .    ", " ")) // this  is    trim .
	t.Log(strings.TrimRight("   this  is    trim .    ", " ")) //    this  is    trim .
	t.Log(strings.TrimSpace("   this  is    trim .    ")) // this  is    trim .
	t.Log(strings.TrimPrefix("   this  is    trim .    ", " ")) //
	t.Log(strings.TrimSuffix("   this  is    trim .    ", " ")) //


}


