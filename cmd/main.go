/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"github.com/wonderivan/logger"
	"k8s-svc-monitor/api/k8s"
	"k8s-svc-monitor/api/result"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

)

func init() {
	// https://github.com/wonderivan/logger
	//DEBG, INFO
	logConfig := `{
  	"TimeFormat": "2019-04-15 20:25:05",
  	"Console": {
   	 "level": "INFO",
   	 "color": true
  	},
  	"File": {
    	"filename": "C:/repositories/cainzhong/src/lucky-draw/lucky-draw.log",
    	"level": "INFO",
    	"daily": true,
    	"maxlines": 1000000,
    	"maxsize": 3,
    	"maxdays": -1,
    	"append": true,
    	"permit": "0660"
 	 }
	}`
	logger.SetLogger(logConfig)
}

func main() {
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
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "itsma1"
	svcList, err := k8s.GetAllServices(clientset, namespace)
	if err != nil {
		logger.Error(err)
	}
	services := svcList.Items
	svcDowntime := &result.SvcDowntime{}
	svcDowntime.StartTimestamp = time.Now().Unix()
	svcDowntime.List = make([]result.Svc, len(services))
	for i:=0;i< len(services);i++ {
		svcDowntime.List[i].Name = services[i].Name
		// kubectl get endpoint -n itsma1
	}
	fmt.Println(fmt.Sprintf("%+v",svcDowntime))
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
