package main

import (
	"bufio"
	"fmt"
	Heap "imooc-product/backend/排序/堆排序"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var filePath = "/Users/wangyingjie/Desktop/big_data.txt"

/**
文件全部读入内存
 */
func readAll(filePath string) {
	start1 := time.Now()
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//readAll spend :  45.840708ms
	fmt.Println("readAll spend : ", time.Now().Sub(start1))
}

/**
按行读取, 有带缓冲区的
 */
func readLine(filePath string) {
	start1 := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 200000)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

/**
大文件切分成多个小文件
 */
func sliceBigFile(filePath, writeDir string, count int)  {
	start1 := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 300000)
	for j := 0; ; j ++{
		outPutFile := fmt.Sprintf("%s/%d.txt", writeDir, j)

		writeFile, err := os.OpenFile(outPutFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 777)
		if err != nil{
			fmt.Println("Open file err =", err)
			writeFile.Close()
			return
		}
		writer := bufio.NewWriter(writeFile)
		heap := Heap.NewMaxHeap()

		//buf := make([]string, 0, count)
		//读取count个数据作为一个文件的行数
		for i := 0; i < count; i ++ {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				j = -1
				break
			}
			//buf = append(buf, string(line))
			number, _ := strconv.Atoi(string(line))
			//大顶堆排序
			heap.Insert(number)
		}

		for !heap.IsEmpty() {
			//按序输出到文件
			max := heap.ExtractMax()
			_, err = writer.WriteString(fmt.Sprintf("%d\r\n", max))
			if err != nil{
				fmt.Println("Write file err =", err)
				return
			}
		}

		if j == -1 {
			break
		}
		writer.Flush()
		writeFile.Close()
		fmt.Printf("导出文件: %s\n", outPutFile)
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

func topN(paths []string, num int, outPutFile string) error {
	var readers []*bufio.Reader
	var files []*os.File

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		reader := bufio.NewReader(file)
		files = append(files, file)
		readers = append(readers, reader)
	}

	writeFile, err := os.OpenFile(outPutFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 777)
	if err != nil {
		return err
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)

	//heap := Heap.NewMaxHeap()
	for k, reader := range readers{
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			files[k].Close()
			delete(files, k)
		}
		d, _ := strconv.Atoi(string(line))
		heap.Insert(d)
	}

	for num > 0 {

		writer.WriteString(item)
		num --
	}
}

func createBigData(outPutFile string) {
	writeFile, err := os.OpenFile(outPutFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 777)
	if err != nil{
		log.Fatal(err)
	}
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)

	for i := 0; i <= 100000000; i ++ {
		_, err = writer.WriteString(fmt.Sprintf("%d\r\n", rand.Int()))
		if err != nil{
			log.Fatal(err)
		}
	}
	writer.Flush()
}

func main() {
	//readAll(filePath)
	//readLine(filePath)
	sliceBigFile(filePath, "/Users/wangyingjie/Documents/code/go/src/imooc-product/storages", 10000000)
	//createBigData("/Users/wangyingjie/Desktop/big_data.txt")
}

//log.Fatal 与panic 的区别

//先判断err, 再使用defer

//删除元素

//https://zhuanlan.zhihu.com/p/44536667
/**
不要直接在循环中使用 defer
defer 是后定义的先执行，和栈类似。
如果在循环中调用 defer，可能会导致堆积了很多 defer，在循环结束后才会执行。
这中间如果有任何一个 defer 失败了怎么办？
多个 defer 执行的内容有没有依赖关系和冲突？
所以，除非万不得已，不要给自己增加复杂度。
不这么用就好了。
 */


//海量数据排序
//普通操作, 好像没任何方法, 只能全部加载到内存中进行排序
//分区操作,
//	切分: 将大文件切分成小文件, 并且每份小文件内部是排好序的,
//	归并: 每份文件获取一条数据存入内存的buffer, 输出最大值, 该值所属文件的下一个值加载到内存, 依次类推....

//海量数据topN, 求最大前N个
//堆操作: 建立一个N个节点的小顶堆, 遍历整个大文件, 将当前读到的数与堆顶元素进行比较, 比堆顶元素大的话放入堆顶, 执行shift-down操作
//分区操作:
//	切分: 将大文件切分成小文件, 并且每份小文件是排好序的并且取前N个
//	归并: 每份文件获取一条数据存入内存的buffer, 输出最大值, 该值所属文件的下一个值加载到内存, 依次类推....