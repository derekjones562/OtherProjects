package main
import "fmt"
import "strings"
import "strconv"

/*
You are building an educational website and want to create a simple calculator for students to use. The calculator will only allow addition and subtraction of non-negative integers.

We also want to allow parentheses in our input. Given an expression string using the "+", "-", "(", and ")" operators like "5+(16-2)", write a function to parse the string and evaluate the result.

Sample output:
  calculate("5+16-((9-6)-(4-2))+1") => 21
  calculate("22+(2-4)") => 20
  calculate("6+9-12") => 3
  calculate("((1024))") => 1024
  calculate("1+(2+3)-(4-5)+6") => 13
  calculate("255") => 255

*/

func main() {

	const expression1 = "6+9-12"; // = 3
	const expression2 = "1+2-3+4-5+6-7"; // = -2
	const expression3 = "100+200+300"; // = 600
	const expression4 = "1-2-3-0"; // = -4
	const expression5 = "255"; // = 255

	const expression2_1 = "5+16-((9-6)-(4-2))+1";
	const expression2_2 = "22+(2-4)";
	const expression2_3 = "6+9-12";
	const expression2_4 = "((1024))";
	const expression2_5 = "1+(2+3)-(4-5)+6";
	const expression2_6 = "255";



	expressions := []string{expression1, expression2, expression3, expression4, expression5}
	for _, expression := range expressions {
		total, err := Calculate(expression)
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to Calculate expression: %s", err.Error()))
		} else {
			fmt.Println(total)
		}
	}

}

func Calculate(expression string) (int,error) {
	var total int
	//   for _, a := range strings.Split(expression, "+"){
	//     d, err := strconv.Atoi(a)
	//     if err != nil {
	//          for i, b := range strings.Split(a, "-"){
	//             c, err := strconv.Atoi(b)
	//             if err != nil {
	//               return 0, err
	//             }
	//             if i ==0 {
	//               c = c *-1
	//            }
	//           d = Subtraction(d, c)
	//           }
	//     }
	//     total += d
	//   }
	stack := []string{}
	for i, v := range expression {
		//     fmt.Println("v")
		//     fmt.Println(string(v))
		switch string(v) {
		case "+":
			left, err := strconv.Atoi(strings.Join(stack, ""))
			if err != nil {
				return 0, err
			}
			//       fmt.Println("left")
			//       fmt.Println(left)
			right, err := Calculate(expression[i+1:])
			//       fmt.Println("right")
			//       fmt.Println(right)
			if err != nil {
				return 0, err
			}
			return Add(left, right), nil
		case "-":
			left, err := strconv.Atoi(strings.Join(stack, ""))
			if err != nil {
				return 0, err
			}
			//       fmt.Println("left")
			//       fmt.Println(left)
			right, err := Calculate(expression[i+1:])
			if err != nil {
				return 0, err
			}
			return Subtraction(left, right), nil
		case "(":
		case ")":
		default:
			stack = append(stack, string(v))
		}
	}
	return total, nil
}

func Add(a, b int) int {
	return a + b
}

func Subtraction (a, b int) int{
	return a - b
}

