This project is to showcase how to deploy multiple service with kubernetes and rely on kube's internal dns to make them visible to each other

# To run

1. make sure you have a running kubernetes cluster (minikube will work, see [here](https://kubernetes.io/docs/tutorials/hello-minikube/) for details on how to run minikube).

2. run server as pod/service in kube cluster

   ```
   $ cd client && kubectl apply -f server.yaml
   ```

   You could verify server is running by curl ingress host/port (which is minikube's ip and port 80)

3. run client as pod/service in kube cluster

   ```
   $ cd server && kubectl apply -f client.yaml
   ```

4. verify the client & server are working by

   ```
   $ kubectl get pods
   $ kubectl logs -f <client-pod-name-from-output-above>
   ```

   you should see client prints out response from server every few seconds

5. Note that if you scale up server pods, `NodePort` will still route client request to the previous pod but now if you start a new client, NodePort should send it to new pod as kubernetes aims to distribute load evenly.
