// Lists all deployments in a namespace
// Options:
//   --namespace NAMESPACE (optional)
//   --all-namespaces (optional)
//   --kubeconfig PATH_TO_CONFIG (optional)
//
// Resources:
// https://github.com/kubernetes/client-go/tree/master/examples

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	defaultKubeConfigFile := fmt.Sprintf("%v/.kube/config", os.Getenv("HOME"))

	kubeConfigFile := flag.String("kubeconfig", defaultKubeConfigFile, "Absolute path to the kubeconfig file")
	namespace := flag.String("namespace", "default", "Namespace to target")
	allNamespaces := flag.Bool("all-namespaces", false, "Target all namespaces")
	flag.Parse()

	if *allNamespaces == true {
		*namespace = apiv1.NamespaceAll
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfigFile)
	if err != nil {
		log.Fatalln("Error building Kubernetes configuration:", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error building Kubernetes client", err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(*namespace)

	getTargetNamespace := func() string {
		if *namespace == apiv1.NamespaceAll {
			return "All Namespaces"
		}
		return *namespace
	}

	fmt.Println("Listing deployments in namespace:", getTargetNamespace())

	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" - %s (%d out of %d Available Replicas)\n", d.Name, d.Status.AvailableReplicas, *d.Spec.Replicas)
	}
}
