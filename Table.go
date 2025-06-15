package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Table struct {
	Head   []string
	Aligns []byte    // 0左对齐 1居中 2右对齐
	Body   [][]*Cell // body[y][x]
}

func NewTable(lines []string) (*Table, bool) {
	var table Table

	// 预先处理第一行的title
	table.Head = splitTableLine(lines[0])
	colNum := table.colNum()

	// 处理第二行，识别列对齐方式
	secondColumns := splitTableLine(lines[1])
	if len(secondColumns) != colNum {
		return nil, false
	}
	for _, e := range secondColumns {
		if len(e) < 2 {
			return nil, false
		}
		table.Aligns = append(table.Aligns, getColumnTextAlign(e))
	}

	// 处理表格正文
	for _, line := range lines[2:] {
		cellLine := splitTableLine(line)
		if len(cellLine) != colNum {
			return nil, false
		}
		bodyLine := []*Cell{}
		for _, e := range cellLine {
			bodyLine = append(bodyLine, NewCell(e))
		}
		table.Body = append(table.Body, bodyLine)
	}

	return table.calcMarkCells(), true
}

func (t *Table) Render() string {
	var builder strings.Builder
	builder.WriteString(`<table>`)

	// 预先处理第一行的title
	builder.WriteString(`<tr>`)
	for _, e := range t.Head {
		builder.WriteString(fmt.Sprintf(`<th>%s</th>`, e))
	}
	builder.WriteString(`</tr>`)

	// 渲染表格正文
	for _, row := range t.Body {
		builder.WriteString(`<tr>`)
		for i, cell := range row {
			if cell.Mark {
				continue
			}
			builder.WriteString(cell.Render(t.Aligns[i]))
		}
		builder.WriteString(`</tr>`)
	}

	builder.WriteString(`</table>`)
	return builder.String()
}

func (t *Table) calcMarkCells() *Table {
	var endx, endy int

	for starty, row := range t.Body {
		for startx, cell := range row {
			if cell.Mark || (cell.RowSpan < 2 && cell.ColSpan < 2) {
				continue
			}
			endx = startx + cell.ColSpan - 1
			if endx > t.colNum() {
				endx = t.colNum() + 1
			}
			endy = starty + cell.RowSpan - 1
			if endy > len(t.Body) {
				endy = len(t.Body) + 1
			}
			t.markCell(startx, starty, endx, endy)
		}
	}

	return t
}

func (t *Table) markCell(startx, starty, endx, endy int) {
	mainCell := true
	for x := startx; x <= endx; x++ {
		for y := starty; y <= endy; y++ {
			if mainCell {
				mainCell = false
				continue
			}
			t.Body[y][x].Mark = true
		}
	}
}

func (t *Table) colNum() int {
	return len(t.Head)
}

type Cell struct {
	Text    string
	RowSpan int  // 如果没有合并单元格则值为1
	ColSpan int  // 如果没有合并单元格则值为1
	Mark    bool // 若mark则不进行计算，因为实际上是被合并掉的单元格
}

// 格式：content@span:2,1
//
// 如果没有合并单元格span为1
func NewCell(textWithSpan string) *Cell {
	var cell Cell
	if textWithSpan == "" {
		return &cell
	}

	text, spanText, _ := strings.Cut(textWithSpan, "@span:")
	cell.Text = text

	if spanText == "" {
		return &cell
	}

	colSpanText, rowSpanText, _ := strings.Cut(spanText, ",")
	rowSpan, _ := strconv.Atoi(rowSpanText)
	colSpan, _ := strconv.Atoi(colSpanText)
	if rowSpan == 0 {
		rowSpan = 1
	}
	if colSpan == 0 {
		colSpan = 1
	}
	cell.RowSpan = rowSpan
	cell.ColSpan = colSpan
	return &cell
}

func (c *Cell) Render(align byte) string {
	var builder strings.Builder
	builder.WriteString(`<td`)

	if align > 0 {
		alignStr := fmt.Sprintf(` style="text-align: %s"`, []string{"auto", "left", "center", "right"}[align])
		builder.WriteString(alignStr)
	}

	if c.RowSpan > 1 {
		builder.WriteString(fmt.Sprintf(` rowspan="%d"`, c.RowSpan))
	}
	if c.ColSpan > 1 {
		builder.WriteString(fmt.Sprintf(` colspan="%d"`, c.ColSpan))
	}
	builder.WriteString(">")
	builder.WriteString(c.Text)
	builder.WriteString("</td>")
	return builder.String()
}

// 查找并替换章节文本中的表格
func ProcChapterContent(text string) string {
	lines := strings.Split(text, "\n")
	beginIndex := -1
	lastEndIndex := 0
	result := []string{}

	for i := range lines {
		// 该行包含管道符，可能是表格的一部分
		if strings.ContainsRune(lines[i], '|') {
			lines[i] = strings.TrimSpace(lines[i])
			lines[i] = trimTunnelSpace(lines[i])

			if beginIndex == -1 {
				beginIndex = i
				continue
			}

			// 在第二行时判断是否符合表格格式
			if i == beginIndex+1 && !containsTunnelHyphen(lines[i]) {
				beginIndex = -1
			}
			continue
		}

		if beginIndex == -1 {
			continue
		}
		// 表格后无空行（格式错误）或表格没有正文（全表小于三行）
		if lines[i] != "" || i == beginIndex+2 {
			beginIndex = -1
			continue
		}

		// 将上一个表格后至该表格前的内容写入result
		result = append(result, lines[lastEndIndex:beginIndex]...)
		lastEndIndex = i

		// 正式处理表格随后写入result
		table, ok := NewTable(lines[beginIndex:i])
		// 编码成功写入编码后的结果，否则写入raw data
		if ok {
			result = append(result, table.Render())
		} else {
			result = append(result, lines[beginIndex:i]...)
		}
		beginIndex = -1
	}

	if beginIndex == -1 {
		result = append(result, lines[lastEndIndex:]...)
	} else {
		// 处理以表格结尾的情况
		table, ok := NewTable(lines[beginIndex:])
		// 编码成功写入编码后的结果，否则写入raw data
		if ok {
			result = append(result, table.Render())
		} else {
			result = append(result, lines[beginIndex:]...)
		}
	}

	return strings.Join(result, "\n")
}

// 识别列对齐方式
//
// ---- 自动0
//
// :--- 左对齐1
//
// :--: 居中对齐2
//
// ---: 右对齐3
func getColumnTextAlign(e string) byte {
	alignLeft := e[0] == ':'
	alignRight := e[len(e)-1] == ':'

	if alignRight {
		if alignLeft {
			return 2
		}
		return 3
	}
	if alignLeft {
		return 1
	}
	return 0
}

// 判断一行文本是否同时包含'|'和'-'
func containsTunnelHyphen(line string) bool {
	var tunnel, hyphen bool
	for _, e := range line {
		switch e {
		case '|':
			tunnel = true
		case '-':
			hyphen = true
		}

		if tunnel && hyphen {
			return true
		}
	}
	return false
}

// 将表格行分割为单元格，去除多余空格
func splitTableLine(line string) []string {
	// line = trimTunnelSpace(line)
	return regexp.MustCompile(` *\| *`).Split(line, -1)
}

// 去掉字符串两端的空格和管道符和\r
func trimTunnelSpace(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == '|' || r == ' '
	})
}
