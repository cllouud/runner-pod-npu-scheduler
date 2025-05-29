package pkg

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const PluginName = "runner-pod-npu-scheduler"

type RunnerPodNpuScheduler struct{}

var _ framework.FilterPlugin = &RunnerPodNpuScheduler{}

func (hs *RunnerPodNpuScheduler) Name() string {
	return PluginName
}

func (hs *RunnerPodNpuScheduler) Filter(ctx context.Context, cycleState *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	// 实现 Filter 逻辑
	fmt.Printf("%s", pod.Name)
	fmt.Printf("%s", pod.Namespace)
	fmt.Printf("%s", pod.APIVersion)
	fmt.Printf("%s", &pod.Spec.Containers[0].Resources)
	return framework.NewStatus(framework.Success, "")
}

func New(_ context.Context, _ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &RunnerPodNpuScheduler{}, nil
}
