package main

import (
	"flag"
	"os"

	"k8s.io/klog/v2"

	plugin "runner-pod-npu-scheduler/pkg"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	defer klog.Flush()

	command := app.NewSchedulerCommand(
		app.WithPlugin(plugin.PluginName, plugin.New),
	)

	if err := command.Execute(); err != nil {
		klog.Errorf("Error executing scheduler command: %v", err)
		os.Exit(1)
		return
	}
}
