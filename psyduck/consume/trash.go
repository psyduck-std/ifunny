package consume

import (
	"github.com/gastrodon/psyduck/sdk"
)

func Trash(parse func(interface{}) error) sdk.Consumer {
	return func(signal chan string) chan []byte {
		data := make(chan []byte, 32)

		go func() {
			for {
				select {
				case received := <-signal:
					panic(received)
				case <-data:
					continue
				}
			}
		}()

		return data
	}
}
