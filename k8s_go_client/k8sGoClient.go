package k8s_go_client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/informers"
	"path/filepath"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"fukiya/utilities"
)


// helper function to get kubeconfig out of kubeconfig path
func GetKubeConfig()(*kubernetes.Clientset, error){
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error loading kubeconfig: %v", err)
	}

	// creating the kube clientset
	kubeclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error loading kube client: %v", err)
	}

	return kubeclient, nil
} 

// helper function to get home dir
func homeDir() string {
	return "/root"
}

// helper function to watch pod activity
func GetPodsStatus(kubeclient *kubernetes.Clientset){
	// creating an informer
	factory := informers.NewSharedInformerFactory(kubeclient, 0)
	informer := factory.Core().V1().Pods().Informer()
	stopper := make(chan struct{})
	defer close(stopper)

	// creating an even handler for consistently look for pod status updates
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
				pod := newObj.(*v1.Pod)
				fmt.Printf("Pod Updated: %s/%s\n", pod.Namespace, pod.Name)
				checkPodStatus(pod)
		},
	})

	// starting the informer
	fmt.Println("Starting Get pod status informer")
	informer.Run(stopper)
}


// helper function used by GetPodsStatus
func checkPodStatus(pod *v1.Pod) {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == v1.PodReady && condition.Status != v1.ConditionTrue {
			fmt.Printf("Pod %s/%s is not ready: %s\n", pod.Namespace, pod.Name, condition.Reason)
			utilities.SendEmailAlert(pod)
		}
	}
}