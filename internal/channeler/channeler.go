package channeler

import (
	"bufio"
	"io"
)

func Channeler(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		buf := bufio.NewReader(r)

		for {
			l, err := buf.ReadString('\n')
			if l != "" {
				ch <- l
			}

			if err != nil {
				break
			}
		}

		close(ch)

	}()

	return ch
}
