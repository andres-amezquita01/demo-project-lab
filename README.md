# Final demo DevOps project lab
Final demo project lab, this is a CI/CD process for a Golang application.
Run terraform apply in staging and production and you will see an infrastructure like this:
![image](https://github.com/andres-amezquita01/demo-project-lab/assets/56136585/d84aba03-d546-4dfd-8110-d1e1ec19d202)
![image](https://github.com/andres-amezquita01/demo-project-lab/assets/56136585/2f96c222-c86f-433e-b509-8514d7c67b38)


About monitoring, in this project was implemented prometheus and grafana, for scraping metrics of the infrastructure of the app (containers), jenkins and golang app, you will find the configuration of each dashboard in ./terraform/monitoring:
* Monitoring of Go-app
![image](https://github.com/andres-amezquita01/demo-project-lab/assets/56136585/67d27b92-d994-49b8-adfd-2ab2ad1738a7)
* Monitoring of Jenkins
![image](https://github.com/andres-amezquita01/demo-project-lab/assets/56136585/918dd320-7905-4204-9302-c80c3233c3cb)
* Monitoring of infrastructure:
![image](https://github.com/andres-amezquita01/demo-project-lab/assets/56136585/3dc975b3-4c12-4d40-88f4-b4b13e1935a9)

and finally, I you want to get more useful information about DevOps:
* https://dev-ops-project-lab.vercel.app/
