package euler

import (
	"fmt"
	"testing"
)

func TestP78(t *testing.T) {

	// Disabled for now
	t.SkipNow()

	/*
	   	memo := make(map[int64]*big.Int)
	   	pair := func(a, b int64) int64 {
	   		return a<<32 + b
	   	}
	   	zero := big.NewInt(0)
	   	one := big.NewInt(1)

	   	var countsums func(total int64, max int64, rec string) *big.Int
	   	countsums = func(total int64, max int64, rec string) *big.Int {
	   		if max > total {
	   			max = total
	   		}
	   		// fmt.Println(rec, total, max)
	   		if max <= 1 {
	   			return one
	   		}

	   		if sum, found := memo[pair(total, max)]; found {
	   			return sum
	   		}

	   		partial := max - 1
	   		sum := big.NewInt(1)
	   		for partial > 0 {
	   			sum.Add(sum, countsums(total-partial, partial, rec+"  "))
	   			partial--
	   		}

	   		if max == 2 {
	   			if sum.Int64() != total/2+1 {
	   				fmt.Println(total, max, "--->", sum, total/2+1)
	   			}
	   		}
	   		memo[pair(total, max)] = sum
	   		return sum
	   	}

	       onemillion := big.NewInt(1000000)
	       for i := int64(3); i < 1000; i++ {
	           cnt := countsums2(i, i, "")
	           fmt.Println("VAL:", i, cnt)
	           if cnt.Mod(cnt, onemillion).Cmp(zero) == 0 {
	               break
	           }
	       }
	*/

	memo := make(map[int64]int64, 256)
	memo[1] = 1
	memo[2] = 2

	var countsums3 func(total int64, max int64) int64
	countsums2 := func(total int64) int64 {
		if sum, found := memo[total]; found {
			return sum
		}

		sum := int64(1)
		// fmt.Println("BEGIN:", total)
		for partial := total - 1; partial > 0; partial-- {
			cur := countsums3(total-partial, partial)
			// fmt.Println("PARTIAL:", partial, total-partial, partial, "->", cur)
			sum += cur
		}
		// fmt.Println("END:", sum)
		memo[total] = sum % 1000000
		return sum
	}

	countsums3 = func(total int64, max int64) int64 {
		if total < max {
			max = total
		}
		if max == 1 {
			return 1
		}

		num_blocks := total / max
		return num_blocks*(memo[max]-1) + 1
	}

	for i := int64(3); i < 100; i++ {
		sum := countsums2(i)
		if i == int64(4) && sum != 5 {
			fmt.Println("ERROR:", i, sum, 5)
		}
		if i == int64(5) && sum != 7 {
			fmt.Println("ERROR:", i, sum, 7)
		}
		if i == int64(6) && sum != 11 {
			fmt.Println("ERROR:", i, sum, 11)
		}

		fmt.Println(i, sum)
		if sum%1000000 == 0 {
			break
		}
	}
}

/*

999 255519*835717

3-----------
1***
2** *
3* * *

4------------
1****
2*** *
4** **
4** * *
5* * * *

5------------
1*****
2**** *
3*** * *
4*** **
5** ** *
6** * * *
7* * * * *

6------------
1******
2***** *
3**** **
4**** * *
5*** ***
6*** ** *
7*** * * *
8** ** **
9** ** * *
0** * * * *
1* * * * * *

7------------
1*******
2****** *
3***** **
4***** * *
5**** ***
6**** ** *
7**** * * *
8*** *** *
9*** ** * *
0*** * * * *
1** ** ** *
2** ** * * *
3** * * * * *
4* * * * * * *














  ** ** **
  ** ** * *
  ** * * * *
  * * * * * *

  *** *** ***
  *** *** ** *
  *** *** * * *
  *** ** * * * *
  *** * * * * * *
  ** * * * * * * *
  * * * * * * * * *

*/
