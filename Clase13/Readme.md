# Step-by-Step Guide to Complete the Project: Sistema de Asignaciones de Cursos

## Objective
The goal of this project is to create a distributed system that simulates course assignments during peak hours, processes massive amounts of data, and visualizes real-time analytics using Kubernetes, NoSQL databases, and Grafana.

---

## Steps to Complete the Project

### 1. **Set Up the Development Environment**
   - Install Kubernetes and configure a cluster.
   - Install Docker for containerization.
   - Install Locust for traffic generation.
   - Install Kafka with Strimzi for message queue management.
   - Install MongoDB and Redis for NoSQL database management.
   - Install Grafana for real-time data visualization.
   - Install Go, Rust, and Python for development.

---

### 2. **Create the Kubernetes Namespace**
   - Create a namespace called `project` in Kubernetes:
     ```bash
     kubectl create namespace project
     ```

---

### 3. **Generate Traffic and Load Data**
   - Write a Locust script in Python to simulate user behavior and generate traffic.
   - Format the input data as JSON:
     ```json
     {
       "curso": "SO1",
       "facultad": "ingenieria",
       "carrera": "sistemas",
       "region": "NORTE"
     }
     ```
   - Store the JSON file as an OCI artifact using `oras`:
     ```bash
     oras push <registry_url>/project/test-data:latest ./test-data.json
     ```
   - Configure Locust to consume the JSON file and simulate 3,000–10,000 concurrent users.

---

### 4. **Set Up Harbor for Container Images**
   - Install and configure Harbor as a private container registry.
   - Create a project in Harbor named `project`.
   - Build Docker images for the following components:
     - API Rest + gRPC Client (Go)
     - gRPC Server (Go)
     - Data Processor (Rust)
   - Push the images to Harbor:
     ```bash
     docker push <harbor_url>/project/<image_name>:<tag>
     ```
   - Configure Kubernetes to access Harbor using a Secret:
     ```bash
     kubectl create secret docker-registry harbor-secret \
       --docker-server=<harbor_url> \
       --docker-username=<username> \
       --docker-password=<password> \
       --namespace=project
     ```

---

### 5. **Deploy the System Components**
   - **Ingress Configuration**:
     - Set up an Ingress to route traffic to the API Rest service.
     - Use NIP.io for a temporary DNS.
   - **Deployment 1: d-api-rest-grpc**:
     - Create a deployment with 3 containers:
       1. **API Rest + gRPC Client (Go)**:
          - Receives data via POST requests.
          - Sends data to the gRPC Server.
       2. **gRPC Server (Go)**:
          - Processes data and sends it to Kafka.
       3. **Data Processor (Rust)**:
          - Stores processed data in Redis for real-time visualization.
     - Apply the deployment in Kubernetes.
   - **Deployment 2: d-kafka-process**:
     - Create a deployment with a single container:
       - Reads messages from Kafka.
       - Stores logs in MongoDB.
     - Apply the deployment in Kubernetes.

---

### 6. **Set Up Grafana for Real-Time Visualization**
   - Deploy Grafana in Kubernetes using the provided YAML file:
     ```bash
     kubectl apply -f grafana-deployment.yaml
     ```
   - Create a dashboard with the following visualizations:
     - Total assignments in Guatemala.
     - Assignments by course, career, and region.
     - Pie charts for course assignments.
   - Add combo boxes to filter data by course, faculty, and region.

---

### 7. **Configure Autoscaling**
   - Enable Horizontal Pod Autoscaling (HPA) for both deployments:
     ```bash
     kubectl autoscale deployment d-api-rest-grpc --cpu-percent=50 --min=1 --max=3 --namespace=project
     kubectl autoscale deployment d-kafka-process --cpu-percent=50 --min=1 --max=3 --namespace=project
     ```

---

### 8. **Test the System**
   - Use Locust to simulate traffic and verify the system's performance.
   - Ensure data is processed correctly and visualized in Grafana.

---

### 9. **Document and Submit**
   - Push all source code to a private GitHub repository under the folder `proyecto2`.
   - Add the auxiliary user `Allenrovas` to the repository.
   - Submit the repository link via UEDI before **23:59 on January 2, 2024**.

---

## Notes
- Ensure all components are properly containerized and deployed in Kubernetes.
- Follow best practices for security and performance optimization.
- Any plagiarism will result in a score of 0 and will be reported.

--- 

## References
- [Strimzi Quickstart Guide](https://strimzi.io/quickstarts/)
- [Grafana Deployment YAML](https://github.com/PacktPublishing/Edge-Computing-Systems-with-Kubernetes/blob/main/ch11/yaml/grafana-deployment.yaml)
- [NoSQL Databases Setup](https://github.com/PacktPublishing/Edge-Computing-Systems-with-Kubernetes/tree/main/ch10/yaml)
- [k8slens](https://k8slens.dev/)
```