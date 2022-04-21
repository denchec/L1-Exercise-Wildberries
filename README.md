11. Что выведет данная программа и почему? 
 1. - Надо передать указатель на wg, потому что мы хотим сказать, что мы обращаемся к прошлой wg и не определяем новую
 2. - wg можно удалить, горутина и так имеет доступ к wg

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(wg sync.WaitGroup, i int) { // *sync.WaitGroup
      fmt.Println(i)
      wg.Done()
    }(wg, i)                            // (&wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
