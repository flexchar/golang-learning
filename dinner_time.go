// Implement the dining philosopher’s problem with the following constraints/modifications.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ChopStick is a type of chopstick
type ChopStick struct {
	sync.Mutex
}

// Philo is a type of philosopher
type Philo struct {
	number                int
	leftStick, rightStick *ChopStick
}

func main() {
	var waitGroup sync.WaitGroup

	// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	CSticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}

	// make philosphers
	philos := make([]*Philo, 5)

	placeholder := []int{0, 1, 2, 3, 4}
	rand.Seed(time.Now().UnixNano())
	// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	rand.Shuffle(len(placeholder), func(i, j int) { placeholder[i], placeholder[j] = placeholder[j], placeholder[i] })

	for num := 0; num < 5; num++ {
		n := placeholder[num]
		// Each philosopher is numbered, 1 through 5.
		philos[n] = &Philo{n, CSticks[n], CSticks[(n+1)%5]}
	}

	for index := 0; index < 5; index += 2 {

		channelOne := make(chan bool)
		channelTwo := make(chan bool)
		// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		go askHostToEat(channelOne)
		go askHostToEat(channelTwo)

		philo1 := <-channelOne
		philo2 := <-channelTwo

		// The host allows no more than 2 philosophers to eat concurrently.
		if philo1 && philo2 {
			waitGroup.Add(1)
			go philos[index].Eat(&waitGroup)

			if index+1 < 5 {
				waitGroup.Add(1)
				go philos[index+1].Eat(&waitGroup)
			}

		}

		waitGroup.Wait()
	}
}

func askHostToEat(p chan bool) {
	// simply approve
	p <- true
}

// Eat will make philosopher to eat food (lock chopsticks for 10ms)
func (p Philo) Eat(waitGroup *sync.WaitGroup) {
	// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for i := 0; i < 3; i++ {

		// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Printf("\nstarting to eat %d\n", p.number+1)

		p.leftStick.Lock()
		p.rightStick.Lock()

		time.Sleep(10 * time.Millisecond)

		// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
		fmt.Printf("\nfinishing eating %d\n", p.number+1)

		p.rightStick.Unlock()
		p.leftStick.Unlock()
	}

	waitGroup.Done()
}
