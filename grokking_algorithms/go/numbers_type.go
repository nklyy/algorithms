package main

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type UnSigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Numbers interface {
	Signed | UnSigned | Float
}
