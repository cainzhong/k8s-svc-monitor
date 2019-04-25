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
	"encoding/json"
	"flag"
	"github.com/wonderivan/logger"
	"io/ioutil"
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
	for i := 0; i < len(services); i++ {
		svcDowntime.List[i].Name = services[i].Name
		detail := result.Detail{}

		endpoints, err := k8s.GetEndpoints(clientset, namespace, services[i].Name)
		if err != nil {
			logger.Error(err)
		}
		subsets := endpoints.Subsets
		detail.Timestamp = time.Now().Unix()
		detail.NumEndpoint = len(subsets)
		logger.Info("Endpoints %s has %d Subsets.", endpoints.Name, len(subsets))
		var details []result.Detail
		details = append(details, detail)
		svcDowntime.List[i].Details = details
	}

	MonitorEndpointsTimer(svcDowntime, clientset, namespace)
}

func MonitorEndpointsTimer(svcDowntime *result.SvcDowntime, clientset *kubernetes.Clientset, namespace string) {
	monitorEndpointsTimer := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-monitorEndpointsTimer.C:
			// 同步代码
			MonitorEndpoints(svcDowntime, clientset, namespace)

			filename := "./api/k8s_svc_downtime.json.json"
			logger.Info("Save it into file %s", filename)
			b, err := json.Marshal(svcDowntime)
			if err != nil {
				logger.Error(err)
			}
			ioutil.WriteFile(filename, b, 0644)
		}
	}
}

func MonitorEndpoints(svcDowntime *result.SvcDowntime, clientset *kubernetes.Clientset, namespace string) {
	svcList := svcDowntime.List
	for i := 0; i < len(svcList); i++ {
		detail := result.Detail{}

		endpoints, err := k8s.GetEndpoints(clientset, namespace, svcList[i].Name)
		if err != nil {
			logger.Error(err)
		}
		subsets := endpoints.Subsets
		detail.Timestamp = time.Now().Unix()
		detail.NumEndpoint = len(subsets)
		logger.Info("Endpoints %s has %d Subsets.", endpoints.Name, len(subsets))
		svcDowntime.List[i].Details = append(svcDowntime.List[i].Details, detail)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
