package chrome

import (
  "encoding/binary"
  "io"
)

var byteOrder binary.ByteOrder = binary.LittleEndian

func Post(msg []byte, writer io.Writer) error {
  // Post message length in native byte order
  header := make([]byte, 4)
  byteOrder.PutUint32(header, (uint32)(len(msg)))
  if n, err := writer.Write(header); err != nil || n != len(header) {
    return err
  }

  // Post message body
  if n, err := writer.Write(msg); err != nil || n != len(msg) {
    return err
  }
  return nil
}

func Recieve(reader io.Reader) ([]byte, error) {
  // Read message length in native byte order
  var length uint32
  if err := binary.Read(reader, byteOrder, &length); err != nil {
    return nil, err
  }

  // Return if no message
  if length == 0 {
    return nil, nil
  }

  // Read message body
  recieved := make([]byte, length)
  if n, err := reader.Read(recieved); err != nil || n != len(recieved) {
    return nil, err
  }
  return recieved, nil
}
