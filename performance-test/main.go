package main

func main() {
	iterations := 1000000

	TestStructKeyMap(iterations)
	TestStringKeyMap(iterations)

	TestReadOnly(iterations)
	TestReadWithLock(iterations)

	TestWriteOnly(iterations)
	TestWriteWithLock(iterations)

	TestNumbersToString_1(iterations)
	TestNumbersToString_2(iterations)

	TestKeyWithStringCombine(iterations)
	TestKeyWithFmtFormat(iterations)
}
