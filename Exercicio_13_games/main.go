package main
import ("fmt"
 		"strings")
func main (){
	//games:= []string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"} //30
	games1 := []string{"0:1","0:2","0:3","0:4","1:2","1:3","1:4","2:3","2:4","3:4"} //0
	//games2 := []string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"} //15

	var points int
	for _, game:= range games1{
	  score := strings.Split(game,":")
	  if score[0] == score[1]{
		points += 1
	  } else if score[0] > score[1]{
		points += 3
	  }
	  
	}
	fmt.Println(points)
}