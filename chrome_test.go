package chrome

import (
  "testing"
  "bytes"
  "encoding/binary"
)

func TestReceive(t *testing.T) {
  message := bytes.NewBufferString("{\"text\": \"hello\"}")
  body := message.Bytes()

  header := make([]byte, 4)
  binary.LittleEndian.PutUint32(header, (uint32)(message.Len()))

  b := append(header, body...)

  reader := bytes.NewReader(b)

  received, err := Receive(reader)
  if err != nil {
    t.Error("Failed to receive message")
  }

  if ! bytes.Equal(received, body) {
    t.Error("Received message is different from original message")
  }
}

func TestPost(t *testing.T) {
  message := bytes.NewBufferString("{\"text\": \"hello\"}")
  body := message.Bytes()

  writer := &bytes.Buffer{}

  err := Post(body, writer)
  if err != nil {
    t.Error("Failed to post message")
  }

  header := make([]byte, 4)
  binary.LittleEndian.PutUint32(header, (uint32)(message.Len()))

  expected := append(header, body...)


  if ! bytes.Equal(expected, writer.Bytes()) {
    t.Error("Posted message is different from original message")
  }
}
