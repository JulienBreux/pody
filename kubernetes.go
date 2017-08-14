package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getClientSet() *kubernetes.Clientset {
	// Use the current context in kubeconfig
	c, err := clientcmd.BuildConfigFromFlags("", *getKubeConfig())
	if err != nil {
		panic(err.Error())
	}

	// Create the client set
	cs, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}

	return cs
}

func getKubeConfig() *string {
	var kc *string
	if h := getHomeDir(); h != "" {
		kc = flag.String("kubeconfig", filepath.Join(h, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kc = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return kc
}

func getHomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getPods() (*v1.PodList, error) {
	cs := getClientSet()
	return cs.CoreV1().Pods(NAMESPACE).List(metav1.ListOptions{})
}

func columnHelperRestarts(cs []v1.ContainerStatus) string {
	r := 0
	for _, c := range cs {
		r = r + int(c.RestartCount)
	}
	return strconv.Itoa(r)
}

func columnHelperAge(t metav1.Time) string {
	delta := time.Now().Sub(t.Time)

	if delta.Hours() > 1 {
		if delta.Hours() > 24 {
			d := float64(delta.Hours() / 24)
			return fmt.Sprintf("%.0fd", d)
		} else {
			return fmt.Sprintf("%.0fh", delta.Hours())
		}
	} else if delta.Minutes() > 1 {
		return fmt.Sprintf("%.0fm", delta.Minutes())
	} else if delta.Seconds() > 1 {
		return fmt.Sprintf("%.0fs", delta.Seconds())
	}

	return "?"
}

func columnHelperStatus(status v1.PodStatus) string {
	return fmt.Sprintf("%s", status.Phase)
}

func columnHelperReady(cs []v1.ContainerStatus) string {
	cr := 0
	for _, c := range cs {
		if c.Ready {
			cr = cr + 1
		}
	}
	return fmt.Sprintf("%d/%d", cr, len(cs))
}
