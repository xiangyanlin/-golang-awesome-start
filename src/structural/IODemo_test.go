package structural

import "testing"

func TestOperateReaderRead(t *testing.T) {
	operateReaderRead()
}

func TestOperateReaderReadByte(t *testing.T) {
	operateReaderReadByte()
}

func TestOperateReaderReadBytes(t *testing.T) {
	operateReaderReadBytes()
}

func TestOperateReaderReadLine(t *testing.T) {
	operateReaderReadLine()
}

func TestOperateReaderReadRune(t *testing.T) {
	operateReaderReadRune()
}

func TestOperateReaderSlice(t *testing.T) {
	operateReaderReadSlice()
}

func TestOperateReaderString(t *testing.T) {
	operateReaderReadString()
}

func TestOperateReaderBuffered(t *testing.T) {
	operateReaderBuffered()
}

func TestOperateReaderPeek(t *testing.T) {
	operateReaderPeek()
}

func TestOperateWriterAvailable(t *testing.T) {
	operateWriterAvailable()
}

func TestOperateWriterBuffered(t *testing.T) {
	operateWriterBuffered()
}
