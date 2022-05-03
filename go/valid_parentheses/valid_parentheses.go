package valid_parentheses

// stack based LIFO approach; the last open char in the queue
// has to match the next closed counterpart
func isValid(s string) bool {

	queue := make([]rune, 0)
	for _, c := range s {

		// if open char
		if c == '{' || c == '(' || c == '[' {
			queue = append(queue, c)
		} else { // if closed char
			if len(queue) == 0 {
				return false
			}

			// set the expected open char of the current closing char
			var e rune
			switch c {
			case '}':
				e = '{'
			case ')':
				e = '('
			case ']':
				e = '['
			}

			// if the expected open char does not match the
			// actual last open char in queue; not valid
			if e != queue[len(queue)-1] {
				return false
			}

			// pop the last value off the queue
			queue = queue[:len(queue)-1]
		}
	}

	// if the queue is not empty after we've processed all the chars; not valid
	return len(queue) == 0
}
