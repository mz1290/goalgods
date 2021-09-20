package main

import "github.com/mz1290/goalgods/ds/queue"

func hotpotato(names []string, num int) string {
	circle := new(queue.Queue)

	for _, name := range names {
		circle.Enqueue(name)
	}

	for {
		for i := 0; i < num; i++ {
			// Save name
			name := circle.Peek()

			// Pass potato (remove first)
			circle.Dequeue()

			// Get at the end of line
			circle.Enqueue(name)
		}

		// Check if we have a winner
		if circle.Size() == 1 {
			break
		}

		// Permanently remove person from game
		circle.Dequeue()
	}

	return circle.Dequeue().(string)
}
