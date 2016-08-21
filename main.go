package main

import (
	"log"
	"os"

	"github.com/TechResearchID/go-in-action/array"
	"github.com/TechResearchID/go-in-action/math"
	"github.com/TechResearchID/go-in-action/slice"
)

func init() {
	log.Println("Init program")
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("Start program")
	log.Println(math.Max(1, 2))
	log.Println(math.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	array.DeclaringArray()
	array.DeclaringArrayUsingArrayLiteral()
	array.CalculatingSize()
	array.SpecificElement()
	array.AccessingArray()
	array.Pointer()
	array.AssigningArray()
	array.AssigningArrayPointer()
	array.TwoDimentionalArray()

	slice.MakeSlice()
	slice.MakeWithLiteral()
	slice.EmptySlice()
	slice.ChangeSlice()
	slice.TakeSliceAsSlice()
	slice.ChangeArraySlice()
	slice.AppendSlice()
	slice.AppendResizeSlice()
	slice.CapacitySet()
	slice.AppendTwoSlice()
	slice.IteratingSlice()
	slice.IteratingSliceIsCopy()
	slice.CreateMultidimentionSlice()
	slice.PassingSliceToFunction()
}
