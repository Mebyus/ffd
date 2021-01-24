package royalroad

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const outdir = "out"

func gatherall(total, chapters int, inRes <-chan result, successRes chan<- error) {
	fmt.Println("Total, chapters: ", total, chapters)
	err := os.MkdirAll(outdir, 0766)
	if err != nil {
		successRes <- err
		return
	}
	inChapter := make([]chan result, chapters)
	chaptersSuccess := make(chan error, chapters)
	parts := chapterParts(total, chapters, 0)
	for i := 0; i < chapters; i++ {
		chaptername := fmt.Sprintf("the-daily-grind-%d.txt", i+1)
		chapterpath := filepath.Join(outdir, chaptername)
		outfile, err := os.Create(chapterpath)
		if err != nil {
			successRes <- err
			return
		}
		defer closeFile(outfile)
		currentParts := chapterParts(total, chapters, i+1)
		inChapter[i] = make(chan result, parts)
		go gather(outfile, currentParts, i*parts, inChapter[i], chaptersSuccess)
	}
	for i := 0; i < total; i++ {
		newResult := <-inRes
		chapter := newResult.index / parts
		inChapter[chapter] <- newResult
	}
	for i := 0; i < chapters; i++ {
		err = <-chaptersSuccess
		if err != nil {
			successRes <- err
			return
		}
	}
	successRes <- nil
	return
}

func chapterParts(total, chapters, i int) int {
	if chapters == 1 {
		return total
	}
	if total%chapters == 0 {
		return total / chapters
	}
	if i < chapters {
		return total / (chapters - 1)
	}
	return total % (chapters - 1)
}

func gather(dst io.Writer, total, start int, inRes <-chan result, successRes chan<- error) {
	current := 0
	results := make([]result, total)
	fmt.Println("Total, start: ", total, start)
	for current < total {
		newResult := <-inRes
		results[newResult.index-start] = newResult
		for current < total && results[current].url != "" {
			_, err := io.Copy(dst, results[current].content)
			if err != nil {
				successRes <- err
				return
			}
			current++
		}
	}
	successRes <- nil
	return
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
