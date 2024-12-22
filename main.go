package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema" // Correct import for GroupVersionResource
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func listPDFs(dynamicClient dynamic.Interface, ctx context.Context, namespace string) {
	// Correctly use the GroupVersionResource from the schema package
	resource := dynamicClient.Resource(
		schema.GroupVersionResource{
			Group:    "k8s.startkubernetes.com",
			Version:  "v1",
			Resource: "mypdfdocuments",
		},
	).Namespace(namespace)

	// List the Custom Resources
	pdfs, err := resource.List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pdfdocuments: %s", err.Error())
	}

	// Output the PDF documents
	fmt.Printf("\nThere are %d PDF documents in the '%s' namespace:\n", len(pdfs.Items), namespace)
	for i, pdf := range pdfs.Items {
		fmt.Printf("%d: %s\n", i+1, pdf.GetName())
	}
}

func main() {
	// Setup kubeconfig (or use in-cluster config if running inside the cluster)
	kubeconfig := flag.String("", "/path/to/your/kubeconfig", "") //
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// Handle error
		fmt.Printf("error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig() // inside..
		if err != nil {
			log.Fatalf("error %s getting in-cluster config", err.Error())
		}
	}

	// Create the dynamic client to interact with CRDs
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %s", err.Error())
	}

	// Set the namespace (can be customized as needed)
	namespace := "default"
	ctx := context.Background()

	// Optionally, list pods as well (standard Kubernetes resource)
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %s", err.Error())
	}
	listPods(*clientSet, ctx, namespace)
	listNS(*clientSet, ctx, namespace)

	// List the custom resources (PDFDocuments) in the given namespace
	listPDFs(dynamicClient, ctx, namespace)

}

// Function to list pods (unchanged)
func listPods(clientSet kubernetes.Clientset, ctx context.Context, namespace string) {
	pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		// Handle errors
		fmt.Printf("error %s, while listing all the pods from default namespace\n", err.Error())
	}
	fmt.Printf("\nThere are %d pods in the '%s' namespace:\n", len(pods.Items), namespace)
	for i, pod := range pods.Items {
		fmt.Println(i+1, pod.Name)
	}

}

func listNS(clientSet kubernetes.Clientset, ctx context.Context, namespace string) {
	namespaces, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		// Handle errors
		fmt.Printf("error %s, while listing all the namespaces \n", err.Error())
	}
	fmt.Printf("\nThere are %d pods in the '%s' namespace:\n", len(namespaces.Items), namespace)

	for i, ns := range namespaces.Items {
		fmt.Println(i+1, ns.Name)
	}

}
