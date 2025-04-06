package app

import (
	"testing"
	"log/slog"
	"os"
	"bytes"
	"io"

	"github.com/stretchr/testify/require"
)

func mustCreateApp(t *testing.T) *App {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)

	counter, err := New(logger)
	require.NoError(t, err)

	return counter
}

func TestRun(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		cli := mustCreateApp(t)

		origStdout := os.Stdout
		pipeReader, pipeWriter, err := os.Pipe()
		require.NoError(t, err)

		os.Stdout = pipeWriter

		file := createTempFileWithContent(t, "a b b a, a c d e f g")
		defer os.Remove(file.Name())

		cli.Run([]string{"binary_name", "2", file.Name()})

		pipeWriter.Close()
		os.Stdout = origStdout

		var buf bytes.Buffer
		io.Copy(&buf, pipeReader)

		output := buf.String()

		require.Equal(t, "a: 3\nb: 2\n", output)
	})
}

func createTempFileWithContent(t *testing.T, content string) *os.File {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "testoutput-*.txt")
	require.NoError(t, err)

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)

	tmpFile.Sync()
	tmpFile.Seek(0, io.SeekStart)

	return tmpFile
}
