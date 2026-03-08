package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// runMain runs main() with the provided args, captures stdout and whether it panicked.
func runMain(args []string) (string, bool) {
    oldArgs := os.Args
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()

    os.Args = args
    os.Stdout = w

    outC := make(chan string)
    go func() {
        var buf bytes.Buffer
        _, _ = io.Copy(&buf, r)
        outC <- buf.String()
    }()

    panicked := false
    func() {
        defer func() {
            if rec := recover(); rec != nil {
                panicked = true
            }
        }()
        main()
    }()

    // close writer to allow reader goroutine to finish
    _ = w.Close()
    out := <-outC

    // restore globals
    os.Args = oldArgs
    os.Stdout = oldStdout

    return out, panicked
}

func TestMainIntegration(t *testing.T) {
    // helper to build paths relative to package directory
    td := func(parts ...string) string {
        p := append([]string{"testdata"}, parts...)
        return filepath.Join(p...)
    }

    cases := []struct {
        name         string
        args         []string
        wantContains string
        wantPanic    bool
    }{
        {name: "NilArgs", args: nil, wantPanic: true},
        {name: "LessThanTwoArgs", args: []string{"cmd"}, wantPanic: true},
        {name: "NonExistentFile", args: []string{"cmd", "does-not-exist.json"}, wantPanic: true},
        {name: "DirectoryPath", args: []string{"cmd", "testdata"}, wantPanic: true},

        {name: "Step1Valid", args: []string{"cmd", td("step1", "valid.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step1Invalid", args: []string{"cmd", td("step1", "invalid.json")}, wantContains: "❌ The provided JSON file is INVALID", wantPanic: false},

        {name: "Step2Valid", args: []string{"cmd", td("step2", "valid.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step2Valid2", args: []string{"cmd", td("step2", "valid2.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step2Invalid", args: []string{"cmd", td("step2", "invalid.json")}, wantContains: "❌ The provided JSON file is INVALID", wantPanic: false},
        {name: "Step2Invalid2", args: []string{"cmd", td("step2", "invalid2.json")}, wantContains: "❌ The provided JSON file is INVALID", wantPanic: false},

        {name: "Step3Valid", args: []string{"cmd", td("step3", "valid.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step3Invalid", args: []string{"cmd", td("step3", "invalid.json")}, wantContains: "❌ The provided JSON file is INVALID", wantPanic: false},

        {name: "Step4Valid", args: []string{"cmd", td("step4", "valid.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step4Valid2", args: []string{"cmd", td("step4", "valid2.json")}, wantContains: "✅ The provided JSON file is VALID", wantPanic: false},
        {name: "Step4Invalid", args: []string{"cmd", td("step4", "invalid.json")}, wantContains: "❌ The provided JSON file is INVALID", wantPanic: false},
    }

    for _, tc := range cases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            // run serially to avoid clobbering os.Args/os.Stdout
            out, pan := runMain(tc.args)

            assert.Equal(t, tc.wantPanic, pan, "case %s: panic mismatch; output: %q", tc.name, out)

            if tc.wantContains != "" {
                assert.Contains(t, out, tc.wantContains, "case %s: output mismatch", tc.name)
            }
        })
    }
}

func TestMainIntegration_ExtraCases(t *testing.T) {
    cases := []struct {
        name         string
        content      string
        wantContains string
        wantPanic    bool
    }{
        {name: "SingleIllegal", content: "'", wantPanic: false, wantContains: "INVALID"},
        {name: "TrailingComma", content: "[1,2,]", wantPanic: false, wantContains: "INVALID"},
        {name: "DeeplyNested", content: `{"a": [{"b": [1, {"c": [2, {"d": [3]}]}]}]}`, wantPanic: false, wantContains: "VALID"},
        {name: "MixedTypes", content: `{"str": "hello", "num": 123, "arr": [true, false, null, 1], "obj": {"k": "v"}, "bool": true, "nul": null}`, wantPanic: false, wantContains: "VALID"},
        {name: "IllegalInObject", content: `{"a": 'b'}`, wantPanic: false, wantContains: "INVALID"},
        {name: "IllegalAfterValid", content: `true False`, wantPanic: false, wantContains: "INVALID"},
    }

    tmp := "_tmp_test.json"
    defer os.Remove(tmp)

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            // Write content to temp file
            err := os.WriteFile(tmp, []byte(tc.content), 0644)
            assert.NoError(t, err)
            out, pan := runMain([]string{"cmd", tmp})
            assert.Equal(t, tc.wantPanic, pan, "case %s: panic mismatch; output: %q", tc.name, out)
            if tc.wantContains != "" {
                assert.Contains(t, out, tc.wantContains, "case %s: output mismatch", tc.name)
            }
        })
    }
}
