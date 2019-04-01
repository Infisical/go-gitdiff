	"io"
	t.Run("read", func(t *testing.T) {
		if err := p.Next(); err != nil {
			t.Fatalf("error advancing parser: %v", err)
		if p.lineno != 1 {
			t.Fatalf("incorrect line number: expected %d, actual: %d", 1, p.lineno)
		}

		line := p.Line(0)
		if err := p.Next(); err != nil {
			t.Fatalf("error advancing parser: %v", err)
		}
		if p.lineno != 2 {
			t.Fatalf("incorrect line number: expected %d, actual: %d", 2, p.lineno)

		line = p.Line(0)
		if err := p.Next(); err != nil {
			t.Fatalf("error advancing parser: %v", err)
		}
		if p.lineno != 3 {
			t.Fatalf("incorrect line number: expected %d, actual: %d", 3, p.lineno)
		}
		line = p.Line(0)
		if line != "the third line\n" {
			t.Fatalf("incorrect third line: %s", line)

		// reading after the last line should return EOF
		if err := p.Next(); err != io.EOF {
			t.Fatalf("expected EOF, but got: %v", err)
		}
		if p.lineno != 4 {
			t.Fatalf("incorrect line number: expected %d, actual: %d", 4, p.lineno)
		// reading again returns EOF again and does not advance the line
		if err := p.Next(); err != io.EOF {
			t.Fatalf("expected EOF, but got: %v", err)
		if p.lineno != 4 {
			t.Fatalf("incorrect line number: expected %d, actual: %d", 4, p.lineno)
		}
	})

	t.Run("peek", func(t *testing.T) {
		p := newParser()

		if err := p.Next(); err != nil {
			t.Fatalf("error advancing parser: %v", err)
		}

		line := p.Line(1)
		if line != "the second line\n" {
		if err := p.Next(); err != nil {
			t.Fatalf("error advancing parser: %v", err)

		line = p.Line(0)
		if line != "the second line\n" {

	t.Run("emptyInput", func(t *testing.T) {
		p := &parser{r: bufio.NewReader(strings.NewReader(""))}
		if err := p.Next(); err != io.EOF {
			t.Fatalf("expected EOF, but got: %v", err)
		}
	})
}

func TestParserAdvancment(t *testing.T) {
	tests := map[string]struct {
		Input   string
		Parse   func(p *parser) error
		EndLine string
	}{
		"ParseGitFileHeader": {
			Input: `diff --git a/dir/file.txt b/dir/file.txt
index 9540595..30e6333 100644
--- a/dir/file.txt
+++ b/dir/file.txt
@@ -1,2 +1,3 @@
context line
`,
			Parse: func(p *parser) error {
				_, err := p.ParseGitFileHeader()
				return err
			},
			EndLine: "@@ -1,2 +1,3 @@\n",
		},
		"ParseTraditionalFileHeader": {
			Input: `--- dir/file.txt
+++ dir/file.txt
@@ -1,2 +1,3 @@
context line
`,
			Parse: func(p *parser) error {
				_, err := p.ParseTraditionalFileHeader()
				return err
			},
			EndLine: "@@ -1,2 +1,3 @@\n",
		},
		"ParseTextFragmentHeader": {
			Input: `@@ -1,2 +1,3 @@
context line
`,
			Parse: func(p *parser) error {
				_, err := p.ParseTextFragmentHeader()
				return err
			},
			EndLine: "context line\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := &parser{r: bufio.NewReader(strings.NewReader(test.Input))}
			p.Next()

			if err := test.Parse(p); err != nil {
				t.Fatalf("unexpected error while parsing: %v", err)
			}

			if test.EndLine != p.Line(0) {
				t.Errorf("incorrect position after parsing\nexpected: %q\nactual: %q", test.EndLine, p.Line(0))
			}
		})
	}
func TestParseTextFragmentHeader(t *testing.T) {
		"trailingComment": {
			Input: "@@ -21,5 +28,9 @@ func test(n int) {\n",
				Comment:     "func test(n int) {",
			p := &parser{r: bufio.NewReader(strings.NewReader(test.Input))}
			p.Next()

			frag, err := p.ParseTextFragmentHeader()
			if !reflect.DeepEqual(test.Output, frag) {
				t.Fatalf("incorrect fragment\nexpected: %+v\nactual: %+v", test.Output, frag)