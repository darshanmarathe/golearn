repackage main

func main() {
	marks := CheckIfCondition()

	Inline_If_exmple()

	Switch_case_example(marks)
}

func Switch_case_example(marks int) {
	switch marks {
	case 10:
		println("ten")
	case 30:
		println("thirty")

	}
}

func Inline_If_exmple() {
	if percent := 20; percent > 10 {
		println("good percent")
	}
}

func CheckIfCondition() int {
	marks := 30
	if marks >= 20 {
		println("passed")
	} else {
		println("failed")
	}
	return marks
}

func CheckMarks() int {
	marks := 30
	if marks >= 20 {
		println("this is more than cool code ")
	}
}