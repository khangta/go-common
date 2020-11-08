package number


func IsPrime(num int) bool {
	//isPrime := true
	for i := 2; i < num - 1; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}
