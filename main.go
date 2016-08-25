package main

import (
	"log"
	"os"

	"github.com/TechResearchID/go-in-action/array"
	"github.com/TechResearchID/go-in-action/concurrency"
	"github.com/TechResearchID/go-in-action/concurrency/pool"
	"github.com/TechResearchID/go-in-action/concurrency/work"
	"github.com/TechResearchID/go-in-action/libs/execute"
	"github.com/TechResearchID/go-in-action/libs/ios"
	"github.com/TechResearchID/go-in-action/libs/jsons"
	"github.com/TechResearchID/go-in-action/libs/loglib"
	"github.com/TechResearchID/go-in-action/maps"
	"github.com/TechResearchID/go-in-action/math"
	"github.com/TechResearchID/go-in-action/slice"
	"github.com/TechResearchID/go-in-action/types"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
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

	maps.CreateEmptyMap()
	maps.InitializeMap()
	maps.MapOfSlice()
	maps.RetrieveValue()
	maps.IteratingMap()
	maps.RemoveItemInMap()
	maps.PassingMap()

	types.CreateUser()
	types.Method()
	types.SendNotification()
	types.ImplementByPointer()
	types.SampleEmbed()

	concurrency.GetCPUCoreNum()
	concurrency.OneLogical()
	concurrency.TimeSlice()
	concurrency.RaceCondition()
	concurrency.SafeRaceCondition()
	concurrency.AtomicLoadAndStore()
	concurrency.MutexRaceCondition()
	concurrency.CreateChannel()
	concurrency.TennisGame()
	concurrency.RunnerGame()
	concurrency.DistributeWorker()

	pool.Main()
	work.Main()
	// runner.Main()

	// loglib.Main()
	loglib.CustomeLogger()

	jsons.GoogleMapsPlaceSearch()
	jsons.ContactJSON()
	jsons.UnmarshalMap()
	jsons.PrettyEncoding()
	jsons.UglyEncoding()
	jsons.StructEncoding()

	ios.Simple()
	ios.Curl()

	execute.RunJavaVersion()
}
