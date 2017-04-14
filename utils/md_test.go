package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestLineIndex(t *testing.T) {
	content := `xx
yyy
search
zzz
`
	want := 2
	got := LineIndex(content, "search")
	if want != got {
		t.Errorf("LineIndex fail, want=%v, got=%v", want, got)
	}
}

func TestGetAllMdTitles(t *testing.T) {
	content := `# one
yyy
## two
# three
zzz
`

	got := GetAllMdTitles(content)
	iwant := 3
	igot := len(got)
	if iwant != igot {
		t.Errorf("GetAllMdTitles fail, len() want=%v, got=%v", iwant, igot)
	}

	sgot := got[0].Title
	swant := "one"
	if swant != sgot {
		t.Errorf("GetAllMdTitles fail, [0].Title want=%v, got=%v", swant, sgot)
	}
	iwant = 1
	igot = got[0].Power
	if iwant != igot {
		t.Errorf("GetAllMdTitles fail, [0].Power want=%v, got=%v", iwant, igot)
	}

	sgot = got[1].Title
	swant = "two"
	if swant != sgot {
		t.Errorf("GetAllMdTitles fail, [1].Title want=%v, got=%v", swant, sgot)
	}
	iwant = 2
	igot = got[1].Power
	if iwant != igot {
		t.Errorf("GetAllMdTitles fail, [0].Power want=%v, got=%v", iwant, igot)
	}

	sgot = got[2].Title
	swant = "three"
	if swant != sgot {
		t.Errorf("GetAllMdTitles fail, [1].Title want=%v, got=%v", swant, sgot)
	}
	iwant = 1
	igot = got[2].Power
	if iwant != igot {
		t.Errorf("GetAllMdTitles fail, [0].Power want=%v, got=%v", iwant, igot)
	}
}

func TestMakeTitleTree(t *testing.T) {
	content := `# one
## two
# three
# four
## four 1
### four 1-1
### four 1-2
## four 2
### four 2-1
### four 2-2
# five
###### five 1-1
`

	titles := GetAllMdTitles(content)
	root := MakeTitleTree(titles)
	got := root.Items

	// fmt.Println(got[0])
	// fmt.Println(got[1])
	// fmt.Println(got[2])
	// fmt.Println(got[2].Items[0])
	// fmt.Println(got[2].Items[0].Items[0])
	// fmt.Println(got[2].Items[0].Items[1])
	// fmt.Println(got[2].Items[1])
	// fmt.Println(got[3])

	iwant := 4
	igot := len(got)
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, len() want=%v, got=%v", iwant, igot)
	}

	swant := "one"
	sgot := got[0].Title
	if swant != sgot {
		t.Errorf("MakeTitleTree fail, [0].Title want=%v, got=%v", swant, sgot)
	}

	iwant = 1
	igot = got[0].Power
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, [0].Power want=%v, got=%v", iwant, igot)
	}

	iwant = 1
	igot = len(got[0].Items)
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, len([0].Items) want=%v, got=%v", iwant, igot)
	}

	swant = "two"
	sgot = got[0].Items[0].Title
	if swant != sgot {
		t.Errorf("MakeTitleTree fail, [0]Items[0].Title want=%v, got=%v", swant, sgot)
	}

	iwant = 2
	igot = got[0].Items[0].Power
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, [0]Items[0].Power want=%v, got=%v", iwant, igot)
	}

	iwant = 1
	igot = got[1].Power
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, [1].Power want=%v, got=%v", iwant, igot)
	}

	swant = "three"
	sgot = got[1].Title
	if swant != sgot {
		t.Errorf("MakeTitleTree fail, [0].Title want=%v, got=%v", swant, sgot)
	}

	iwant = 0
	igot = len(got[1].Items)
	if iwant != igot {
		t.Errorf("MakeTitleTree fail, len([1].Items) want=%v, got=%v", iwant, igot)
	}
}

func TestTraverse(t *testing.T) {
	content := `# one
## two
# three
# four
## four 1
### four 1-1
### four 1-2
## four 2
### four 2-1
#### four 2-2
# five
###### five 1-1
`

	titles := GetAllMdTitles(content)
	root := MakeTitleTree(titles)
	got := ""
	root.Traverse(PowerLess(5, func(n *MdTitleTree) {
		// got += MakeTOCItem("  ", n) + "\n"
		link := GetMdLinkHash(n.Title)
		x := strings.Repeat("  ", n.Power)
		got += fmt.Sprintf("%v- [%v](#%v)\n", x, n.Title, link)
	}))
	want := `  - [one](#one)
    - [two](#two)
  - [three](#three)
  - [four](#four)
    - [four 1](#four-1)
      - [four 1-1](#four-1-1)
      - [four 1-2](#four-1-2)
    - [four 2](#four-2)
      - [four 2-1](#four-2-1)
        - [four 2-2](#four-2-2)
  - [five](#five)
`
	if want != got {
		t.Errorf("TestTraverse failed, want=\n%v\ngot\n%v", want, got)
	}
}

func TestGetMdLinkHash(t *testing.T) {
	sgot := GetMdLinkHash("/$ .>")
	swant := "-"
	if swant != sgot {
		t.Errorf("GetMdLinkHash fail, want=%v, got=%v", swant, sgot)
	}
}
