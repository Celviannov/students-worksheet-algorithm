# EXPLANATION.md - Go, Sequencing & Introduction to Genetic Algorithms


## Lab 01 - Campus Festival Registration: What is your data?

I declared four variables, each with the type that matches what the value really is: age as an int because it is a whole-number count, mass as a float64 because it is a fractional measurement, checked-in status as a bool because it has only two states, and student ID as a string because its leading zeros are data, not value. The program prints them as 19 | 62.5 | true | 00123.

**Explain:** A student ID is an identifier, not a quantity - you never add, subtract, or average IDs. If treated as an int, Go silently drops the leading zeros (00123 becomes 123), and worse, the conversion succeeds with no error, so the corruption is invisible until the wrong record is fetched.

**Reflection:** Choosing a type is choosing which operations make sense.


## Lab 02 - Campus Café: Does order matter?

The program computes the subtotal first (unitPrice x quantity = 36000) and only then subtracts the discount to get the total (31000); changing the quantity to three gives a subtotal of 54000 and a total of 49000 by editing one line.

**Explain:** The subtotal must be computed first because the discount is defined as a reduction of the subtotal, not of the unit price - in Go, reversing the lines is even a compile error, because subtotal would not exist yet. If the discount exceeded the subtotal, the raw subtraction would produce a negative total, which a real checkout system should clamp to zero or reject as invalid data; this lab assumes the discount never exceeds the subtotal.

**Reflection:** The right calculation in the wrong order may still lead to the wrong decision.


## Lab 03 - Equal Volunteer Teams: Can everyone fit?

After reading two integers, the program validates that both are positive and prints invalid input otherwise; if the input is valid, it prints whether y % x == 0, so 5 20 gives true, 20 5 gives false, 123 123 gives true, and 0 20 gives invalid input.

**Explain:** Checking x > 0 before the modulo is essential because Go panics at runtime on integer division or remainder by zero, and since x is only known at runtime, a guard must run before the operation - the pattern is guard first, compute second. (Go technically allows negative operands in %, but negative group sizes are meaningless here, so validation rejects them anyway.)


## Lab 04 - Population Dashboard: A model is a simplification

The program applies the flow rule final = initial + births + arrivals - deaths - departures to the data 1000, 20, 5, 10, 4 and prints Population: 1011.

**Explain:** A negative result would mean more people left than were present and added - a signal of bad input data or of the model being pushed beyond its assumptions, since a negative population is physically meaningless. None of the five inputs can reasonably be negative, because they are all counts; only the result can go negative, and that is precisely the alarm bell.

**Reflection:** A useful mathematical model is not the whole of reality; its assumptions matter - real demography would also involve age structure, migration timing, and measurement error.


## Lab 05 - A Composer's Pattern: Number sequences

Without a loop, the program computes each new term as the sum of the previous two: for starting values 5 and 3 it produces 8, 11, and 19; the same rule verified for 1 1 (giving 2 3 5) and 10 20 (giving 30 50 80).

**Explain:** Generating a number sequence is different from merely executing instructions sequentially because each term has a data dependency on the previous ones - fourth literally consumes third, so the order is forced by the values, not just by the clock, and reordering the lines would not even compile. An optional for loop with a sliding window (a, b = b, next) expresses the same rule once instead of three times, which is the bridge to the genetic-algorithm labs where operations repeat across generations.


## Lab 06 - Gravity on Mars: Why units matter

The program computes Earth weight as mass x 9.8 = 686.0 N and Mars weight as mass x 9.8 x 0.38 = 260.68 N, printing each to one decimal place with fmt.Printf, so the output reads Earth: 686.0 N and Mars: 260.7 N.

**Explain:** The input is in kilograms while the output is in newtons because they are different physical quantities: kilograms measure mass (the amount of matter, unchanged between planets), while newtons measure weight (the gravitational force on that mass, which changes with gravity). float64 is used because masses, constants, and forces are continuous measurements that can be fractional, and an int would truncate them. One subtlety: because 9.8 and 0.38 have no exact binary representation, the computed Mars value is really about 260.68000000000006 - %.1f rounds it for display, which is a reminder never to compare floats with ==.


## Lab 07 - Tiny genetic algorithm operations

The fitness function counts the 1 bits in a candidate; crossing parents 1001 and 0110 after the second bit produces the child 1010 (fitness 2), and mutating its second bit produces 1110 (fitness 3), with both states printed before and after mutation - matching the worksheet's manual walkthrough exactly.

**Explain:** This demonstrates GA operations, not a full genetic algorithm, because a real GA runs selection, crossover, and mutation inside a loop over many generations with fitness as feedback, while here the parents are hardcoded, the cut point fixed, and there is no selection step or repetition - it is a microscope slide of the three tools before Lab 09 assembles them into a machine. Mutation does **not** always improve fitness: it helped here (2 to 3), but flipping the same bit in 1110 would drop fitness back to 2, which is exactly why a GA relies on populations and selection rather than any single mutation.

**Reflection:** Evolution changes candidates; the objective function defines what counts as better.


## Lab 08 - Cube-root Check: A precise rule

The program reads integers x and y and verifies xxx == y using exact integer multiplication, so 5 125 prints true, -3 -27 prints true, and 12 12 prints false.

**Explain:** A negative number can have a negative cube because the sign multiplies three times: minus times minus is plus, and plus times minus is minus - odd powers preserve the base's sign (a square, by contrast, can never be negative, which is why the problem specifies a cube). For very large integers, though, x-cubed overflows silently once x exceeds about 1.29 million in Go's 64-bit int, wrapping around to wrong values so a large true cube could print false; arbitrary-precision arithmetic (math/big) or division-based verification would fix that, but for this lab's small cases plain multiplication is exact. This lab and Lab 06 are two sides of one lesson: floats give fast approximate answers, integers give exact answers within a limited range.


## Lab 09 - Advanced Bonus: Evolve a population

Instead of editing one bit string once, the program maintains a population of six random four-bit candidates and repeats a full cycle per generation - compute fitness, select parents by tournament (three random picks, fittest wins), cross them at a random cut point, mutate each bit with probability 0.1, and carry the best candidate forward unchanged (elitism) - printing the best fitness every generation, stopping early if fitness reaches 4 or after a hard limit of 20 generations.

**Explain:** The algorithm does not guarantee finding 1111 within the budget because the search is stochastic: mutation may never flip the needed bit, crossover can only mix genes that parents already carry, and with a small population tournament selection can lose a good gene by chance - the 20-generation limit simply stops the loop, reached goal or not, and a guarantee would require exhaustively checking all 16 strings, which defeats the purpose of a search. Random exploration helps because selection alone can stall: if no candidate has a 1 in some position, crossover can never create one, so blind mutation is the only source of new genetic material. But the same blindness can make candidates worse - flipping a 1 to 0 in 1110 drops fitness from 3 to 2 - which is why elitism matters: the champion survives unchanged, so the population can never regress below its current best, and worse individuals exist only as exploration noise to be filtered by selection.

**Boundary-case testing:** running with several seeds (42, 1, 7, 99) shows runs that differ in how many generations they need - and some that fail to reach fitness 4 within 20 generations - directly demonstrating the non-guarantee.
