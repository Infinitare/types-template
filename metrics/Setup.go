package metrics

import (
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"os"
)

func Setup(service string) (func(), func()) {

	if os.Getenv("DISABLE_TRACES") == "true" {
		return func() {}, func() {}
	}

	tracer.Start(
		tracer.WithService(service),
		tracer.WithEnv(os.Getenv("ENV")),
		tracer.WithDebugStack(true),
		tracer.WithAnalytics(true),
		tracer.WithRuntimeMetrics(),
	)

	err := profiler.Start(
		profiler.WithService(service),
		profiler.WithEnv(os.Getenv("ENV")),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)

	if err != nil {
		panic(err)
	}

	return tracer.Stop, profiler.Stop

}
