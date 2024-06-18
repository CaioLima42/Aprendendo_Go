package main

	import (
	"fmt"
	 //"example.com/myapp/sample"
	 //temp "gopkg.in/yaml.v2"
	 //"temp"
	 //"sync"
	)
	type TestStruct struct {
		Sample string
		SampleInt int
		}

		func (t *TestStruct) ChangeSample(s string) {
		t.Sample = s
		fmt.Println(t)
	}
	
	/*var messages = make(chan string)
	func createPing() {
		messages <- "tread1"
		fmt.Println("eu, funcao 1, enviei uma mensagem para a main")
		wg.Done()
	}
	func runcode(){
		messages <- "tread2"
		fmt.Println("eu funcao 1, enviei uma mensagem para a main")
		wg.Done()
	}
	var wg sync.WaitGroup
*/
	func exampleFunc() (int, error) {
		erro := "vai toma no cu"
		return 1, fmt.Errorf("this is an example error %v", erro)
	}
	func main() {
		/*
		palavra := "flamengo"
		var palavra2 string = "cruzeiro"
		fmt.Println(palavra2)
		fmt.Println(palavra)
		fmt.Println("Hello, 世界")
		output, _ := temp.Marshal(map[string]string{"a": "b", "c": "d"})
		fmt.Println(string(output))
		*/
		/*
		primes := []int{2, 3, 5, 7, 11}
		primes = append(primes, 12)
		fmt.Println(primes)
		*/
		/*
		primes := make([]int,6)
		primes = append(primes, 5)
		fmt.Println(primes)
		primes[0] = 2
		primes[0] = 3
		fmt.Println(primes)
		*/
		/*
		mappedItems := make(map[string]string)
		mappedItems["test"] = "test"
		fmt.Println(mappedItems["test"])
		*/
		/*
		mappedItems := map[string]string{"valor":"valor", "tipo":"tipo"}
		mappedItems["test"] = "teste"
		fmt.Println(mappedItems)
		*/
		/*
		minhaMae := func (a int) int {
			a++
			return a
		}
		mappedItems := make(map[string]func(a int)int)
		mappedItems["flamengo"] = minhaMae
		num := 5
		fmt.Printf("Double of %d: %d\n", num, mappedItems["flamengo"](num))
		*/
		/*
		testVar := TestStruct{"aa", 1}
		fmt.Println(testVar)
		testVar.ChangeSample("bb")
		fmt.Println(testVar)
		*/
		/*
		type Subtask struct {
			Param1 string
			Param2 string
			Status string
			}
		allTasks := []Subtask{{Status: "incomplete"},{Status:"completed"}}
		for _, x := range allTasks {
			if x.Status != "completed" {
			 fmt.Println("Main task is still incomplete")
			 }
		}
		*/
		/*
		items := map[string]string{
			"key1": "value1",
			"key2": "value2",
		   }
		   for idx, x := range items {
			fmt.Println(idx)
			fmt.Println(x)
		   }
		   */
		/*
		fmt.Println("Program Begin")
		e := sample.ExposedStruct{
			ExposedSample: "sample", //parametro publico(letra maiuscula)
			unexposedSample: "sample", // parametro privado(letra minuscula)
		}
		fmt.Println(e)
		*/
		/*
		wg.Add(2)
		go createPing()
		go runcode()
		msg := <-messages
		wg.Wait()
		fmt.Println(msg)
		*/
		saida, entrada := exampleFunc()
		fmt.Println(saida)
		fmt.Println(entrada)
		}
