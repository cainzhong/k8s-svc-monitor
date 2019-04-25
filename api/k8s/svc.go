package k8s

import (
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetAllSvc(namespace string) {
	clientset := GetClientset();
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d Service in the cluster\n", len(services.Items))

	if err != nil {
		panic(err.Error())
	}
	for i, length := 0, len(services.Items); i < length; i++ {
		fmt.Println(fmt.Sprintf("%s ", services.Items[i].Name))
	}
}

func GetAllServices(clientset *kubernetes.Clientset, namespace string) (*v1.ServiceList, error) {
	serviceList, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return serviceList, nil
}

var clientsetInit = false

func GetClientset() *kubernetes.Clientset {
	var clientset = &kubernetes.Clientset{}
	if clientsetInit == true {
		return clientset
	} else {
		var kubeconfig *string
		/*if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}*/
		kubeconfig = flag.String("kubeconfig", "C:/repositories/cainzhong/src/k8s-svc-monitor/assets/ssl/kubeconfig", "absolute path to the kubeconfig file")
		flag.Parse()

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		// create the clientset
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		clientsetInit = true
		return clientset
	}
}
