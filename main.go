package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func listPods(clientSet kubernetes.Clientset, ctx context.Context, namespace string) {
	pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		//handle errors
	}

	for i, pods := range pods.Items {
		fmt.Println(i+1, pods.Name)
	}
	fmt.Printf("\nThere are %d pods in the '%s' namespace\n", len(pods.Items), namespace)
}

func main() {

	kubeconfig := flag.String("", "Path/to/your/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handle error
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle errors
	}

	namespace := "kube-system"
	ctx := context.Background()

	listPods(*clientSet, ctx, namespace)
	//fmt.Printf("\nThere are %d pods in the '%s' namespace\n", len(pods.Items), namespace)

	deployments, err := clientSet.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})

	fmt.Printf("\nThere are %d deploys in the '%s' namespace\n", len(deployments.Items), namespace)
	for i, deplodeployments := range deployments.Items {
		fmt.Println(i+1, deplodeployments.Name)
	}

	ns, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	fmt.Printf("\nThere are %d ns in the cluster\n", len(ns.Items))
	for i, ns := range ns.Items {
		fmt.Println(i+1, ns.Name)
	}

}
