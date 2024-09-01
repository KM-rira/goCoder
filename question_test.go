package question_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"question"
	"testing"
)

func TestQuestion(t *testing.T) {
	inputDir := "./data/input"
	outputDir := "./data/output"
	files, err := os.ReadDir(inputDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for n, inputFile := range files {

		if inputFile.IsDir() {
			continue
		}

		inputFilePath := filepath.Join(inputDir, inputFile.Name())
		outputFilePath := filepath.Join(outputDir, fmt.Sprintf("output%d.txt", n+1))

		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			t.Fatalf("Failed to open input file: %v", err)
		}
		defer inputFile.Close()

		// ファイルが空かどうかを確認
		fileInfo, err := inputFile.Stat()
		if err != nil {
			t.Fatalf("Failed to get file stats: %v", err)
		}
		if fileInfo.Size() == 0 {
			break
		}

		inputFileHandle, err := os.Open(inputFilePath)
		if err != nil {
			t.Fatalf("Failed to open input file %s: %v", inputFilePath, err)
		}
		defer inputFileHandle.Close()

		// 標準出力をキャプチャするためのパイプを作成
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe: %v", err)
		}

		// 標準出力をパイプにリダイレクト
		stdout := os.Stdout
		os.Stdout = w
		// テストが終了したら元に戻すためのクリーンアップ処理
		defer func() {
			os.Stdout = stdout
			w.Close()
		}()

		// 標準入力を `inputFileHandle` に設定
		stdin := os.Stdin
		os.Stdin = inputFileHandle

		// test do
		question.Question()

		// 標準入力を元に戻す
		os.Stdin = stdin

		// パイプを閉じて出力をキャプチャ
		w.Close()
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			t.Fatalf("Failed to copy from pipe: %v", err)
		}
		r.Close()

		// 実際の出力を取得
		output := buf.String()

		// 期待される出力を読み込む
		expectedOutput, err := os.ReadFile(outputFilePath)
		if err != nil {
			t.Fatalf("Failed to read output file: %v", err)
		}

		// 実際の出力と期待される出力を比較
		if output != string(expectedOutput) {
			t.Errorf("Number: %d Expected output: %s, but got: %s", n+1, string(expectedOutput), output)
		} else {
			t.Logf("Number: %d Success", n+1)
		}
	}
}
