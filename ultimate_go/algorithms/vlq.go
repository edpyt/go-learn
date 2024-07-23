package main

import (
  "fmt"
  "math"
  "math/bits"
)

func main() {
	inputs := [][]byte{
		[]byte{0x7F},
		[]byte{0x81, 0x00},
		[]byte{0xC0, 0x00},
		[]byte{0xFF, 0x7F},
		[]byte{0x81, 0x80, 0x00},
		[]byte{0xFF, 0xFF, 0x7F},
		[]byte{0x81, 0x80, 0x80, 0x00},
		[]byte{0xC0, 0x80, 0x80, 0x00},
		[]byte{0xFF, 0xFF, 0xFF, 0x7F},
		[]byte{0x82, 0x00},
		[]byte{0x81, 0x10},
	}

	for _, input := range inputs {
	  decoded, err := DecodeVariant(input)
	  if err != nil {
	    fmt.Println(err)
	    return
	  }

	  encoded := EncodeVariant(decoded)
	  fmt.Printf("input 0x%x, decoded: %d, encoded: 0x%x\n", input, decoded, encoded)
	}
}

func DecodeVariant(input []byte) (uint32, error) {
  const lastBitSet = 0x80

  var d uint32
  var bitPos int

  for i := len(input) - 1; i >= 0; i-- {
    n := uint8(input[i])

    for checkBit := 0; checkBit < 7; checkBit++ {
      n = bits.RotateLeft8(n, -1)

      if n >= lastBitSet {
        switch {
        case bitPos == 0:
          d++
        default:
          base10 := math.Pow(2, float64(bitPos))
          d += uint32(base10)
        }
      }

      bitPos++
    }
  }
  
  return d, nil
}

func EncodeVariant(n uint32) []byte {
  const maxBytes = 4
  const eightBitSet = 0x80
  const lastBitSet = 0x80000000

  encoded := make([]byte, maxBytes)

  for bytePos := maxBytes - 1; bytePos >= 0; bytePos-- {
    var d uint8

    for checkBit := 0; checkBit < 7; checkBit++ {
      n = bits.RotateLeft32(n, -1)

      if n >= lastBitSet {
        switch {
        case checkBit == 0:
          d++
        default:
          base10 := math.Pow(2, float64(checkBit))
          d += uint8(base10)
        }
      }
    }

    if bytePos < 3 {
      d += eightBitSet
    }

    encoded[bytePos] =  d
  }

  for bytePos, b := range encoded {
    if b == eightBitSet {
      continue
    }
    encoded = encoded[bytePos:]
    break
  }
  
  return encoded
}
