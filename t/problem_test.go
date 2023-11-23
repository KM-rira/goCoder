package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestProblem(t *testing.T) {
	inputFile, err := os.Open("./data/input.txt")
	if err != nil {
		t.Fatalf("Failed to open input file: %v", err)
	}
	defer inputFile.Close()

	// 標準出力をキャプチャするためのパイプを作成
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// 標準出力をパイプにリダイレクト
	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
		w.Close()
	}()

	// 関数を実行
	os.Stdin = inputFile
	problem()

	// パイプを閉じて出力をキャプチャ
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()

	// 実際の出力を取得
	output := buf.String()

	// 期待される出力を読み込む
	expectedOutput, err := os.ReadFile("./data/output.txt")
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// 実際の出力と期待される出力を比較
	if output != string(expectedOutput) {
		t.Errorf("Expected output %s, but got %s", string(expectedOutput), output)
	}
}
