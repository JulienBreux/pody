package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get Kubernetes client set
func getClientSet() *kubernetes.Clientset {
	c := getConfig()

	// Use the current context in kubeconfig
	cc, err := clientcmd.BuildConfigFromFlags("", *c.kubeConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create the client set
	cs, err := kubernetes.NewForConfig(cc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cs
}

// Get pods (use namespace)
func getPods() (*v1.PodList, error) {
	cs := getClientSet()

	return cs.CoreV1().Pods(NAMESPACE).List(metav1.ListOptions{})
}

// Get the pod containers
func getPodContainers(p string) []string {
	var pc []string
	cs := getClientSet()

	pod, _ := cs.CoreV1().Pods(NAMESPACE).Get(p, metav1.GetOptions{})
	for _, c := range pod.Spec.Containers {
		pc = append(pc, c.Name)
	}

	return pc
}

// Delete pod
func deletePod(p string) error {
	cs := getClientSet()

	return cs.CoreV1().Pods(NAMESPACE).Delete(p, &metav1.DeleteOptions{})
}

// Get pod container logs
func getPodContainerLogs(p string, c string, o io.Writer) error {
	tl := int64(50)
	cs := getClientSet()

	opts := &v1.PodLogOptions{
		Container: c,
		TailLines: &tl,
	}

	req := cs.CoreV1().Pods(NAMESPACE).GetLogs(p, opts)

	readCloser, err := req.Stream()
	if err != nil {
		return err
	}

	_, err = io.Copy(o, readCloser)

	readCloser.Close()

	return err
}

// Column helper: Restarts
func columnHelperRestarts(cs []v1.ContainerStatus) string {
	r := 0
	for _, c := range cs {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

// Column helper: Age
func columnHelperAge(t metav1.Time) string {
	d := time.Now().Sub(t.Time)

	if d.Hours() > 1 {
		if d.Hours() > 24 {
			ds := float64(d.Hours() / 24)
			return fmt.Sprintf("%.0fd", ds)
		} else {
			return fmt.Sprintf("%.0fh", d.Hours())
		}
	} else if d.Minutes() > 1 {
		return fmt.Sprintf("%.0fm", d.Minutes())
	} else if d.Seconds() > 1 {
		return fmt.Sprintf("%.0fs", d.Seconds())
	}

	return "?"
}

// Column helper: Status
func columnHelperStatus(s v1.PodStatus) string {
	return fmt.Sprintf("%s", s.Phase)
}

// Column helper: Ready
func columnHelperReady(cs []v1.ContainerStatus) string {
	cr := 0
	for _, c := range cs {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(cs))
}
