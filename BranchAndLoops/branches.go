package main

func main() {
	marks := 30
	if marks >= 20 {
		println("passed")
	} else {
		println("failed")
	}

	if percent := 20; percent > 10 {
		println("good percent")
	}

	switch marks {
	case 10:
		println("ten")
	case 30:
		println("thirty")

	}

}
