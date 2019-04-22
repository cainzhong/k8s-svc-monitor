# k8s-svc-monitor
Monitor k8s service has a endpoint 

1. Use a thread to get all the k8s service every n seconds in all namespaces.  
   a. new added k8s svc?  
   b. deleted k8s svc?
2. Check every k8s svc's endpoint every n seconds.
   a. If the k8s svc has at least 1 endpoint, it means that the k8s svc can provide business service.
   b. If the k8s svc has no endpoint, it means that the k8s svc cannot provide business service. Mark the first detected time as the start time, we can calcuate the 'downtime duration' = 'current time' - 'start time'
3. Use UI to show the k8s svc downtime result.  
   **Eg.**  

   **Overview:**  
   ![avatar](https://github.com/cainzhong/k8s-svc-monitor/blob/master/assets/k8s_svc_downtime_overview.png)

   **Downtime detail for specified k8s svc:**  
   ![avatar](https://github.com/cainzhong/k8s-svc-monitor/blob/master/assets/k8s_svc_downtime_detail.png)
