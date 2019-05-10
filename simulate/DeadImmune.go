package simulate

/* import "time"

// UpdateHumanStateNew - renamed function so not to clash with original
// This version is trying to implement that a human can become immune or die before the virus completes its lifespan
func UpdateHumanStateNew(human *Human) {
	//find total number of health states = healthy + lifespan of virus + immune + dead
	var TotalNumberOfStates = 3 + human.Viruses[1].Lifespan

	//start by incrementing health state to infected
	human.State++
	//always sleep after updating state
	time.Sleep(time.Second * 3)

	//invalid operation for condition but will work if we change data type of state to int
	//we want this while-loop and condition because we want to keep incrementing the state until human is dead either by following the cycle or dying prematurely
	for human.State < TotalNumberOfStates-1 {
		//once in a infected state, try to decrement healthiness to see if human will die prematurely
		//start by calculating how much we decrement with based on initial healthiness - healthiness needs to be a float for these operations to work!
		//The following calculation will generate a high value if healthiness is low and vice versa
		var DecrementValue = 3/human.Healthiness * 100
		human.Healthiness = human.Healthiness - DecrementValue

		//check if we killed human prematurely - if so, set state to dead - if not, keep running the cycle = increment state as usual
		if human.Healthiness <= 0 {
			human.State = TotalNumberOfStates - 1
		} else {
			human.State++
			time.Sleep(time.Second * 3)
		}
	}
	//fmt.Println("human done updating!")
}
*/
