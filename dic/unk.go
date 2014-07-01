package dic

func ReferUnkContent(a_id int) (content Content, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	return UnkContents[a_id], nil
}

func GetUnkCost(a_id int) (cost Cost, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	cost = UnkCosts[a_id]
	return
}

var UnkContents []Content = []Content{
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"感動詞", "*", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"記号", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "サ変接続", "*", "*", "*", "*", "*", "*", "*"},
	Content{"感動詞", "*", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "サ変接続", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "数", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "人名", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "一般", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "地域", "一般", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "組織", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "固有名詞", "一般", "*", "*", "*", "*", "*", "*"},
	Content{"感動詞", "*", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "数", "*", "*", "*", "*", "*", "*", "*"},
	Content{"記号", "空白", "*", "*", "*", "*", "*", "*", "*"},
	Content{"名詞", "サ変接続", "*", "*", "*", "*", "*", "*", "*"},
}

var UnkCosts []Cost = []Cost{
	Cost{1285, 1285, 13398}, Cost{1293, 1293, 18706}, Cost{3, 3, 15235}, Cost{1288, 1288, 15673}, Cost{1289, 1289, 18188},
	Cost{1292, 1292, 13835}, Cost{1289, 1289, 12615}, Cost{1292, 1292, 8492}, Cost{1293, 1293, 12600}, Cost{1285, 1285, 7966},
	Cost{1288, 1288, 9866}, Cost{5, 5, 4769}, Cost{1285, 1285, 7884}, Cost{1292, 1292, 8573}, Cost{1289, 1289, 12697},
	Cost{1288, 1288, 10029}, Cost{1293, 1293, 12681}, Cost{1293, 1293, 17882}, Cost{1288, 1288, 14787}, Cost{1289, 1289, 18060},
	Cost{1285, 1285, 13069}, Cost{1283, 1283, 20223}, Cost{3, 3, 16989}, Cost{1292, 1292, 14761}, Cost{1283, 1283, 17290},
	Cost{1292, 1292, 12649}, Cost{1285, 1285, 11426}, Cost{1288, 1288, 15295}, Cost{1289, 1289, 17340}, Cost{1293, 1293, 17611},
	Cost{1295, 1295, 27473}, Cost{1289, 1289, 13581}, Cost{1285, 1285, 9461}, Cost{1293, 1293, 13661}, Cost{1292, 1292, 10922},
	Cost{1288, 1288, 10521}, Cost{3, 3, 14138}, Cost{1295, 1295, 27386}, Cost{9, 9, 8903}, Cost{1283, 1283, 17585},
}

var UnkIndex map[CharacterClass][2]int = map[CharacterClass][2]int{
	GREEK:        [2]int{12, 5},
	HIRAGANA:     [2]int{17, 7},
	KANJINUMERIC: [2]int{30, 1},
	KATAKANA:     [2]int{31, 6},
	ALPHA:        [2]int{0, 6},
	CYRILLIC:     [2]int{6, 5},
	DEFAULT:      [2]int{11, 1},
	KANJI:        [2]int{24, 6},
	NUMERIC:      [2]int{37, 1},
	SPACE:        [2]int{38, 1},
	SYMBOL:       [2]int{39, 1},
}
