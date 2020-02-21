#!/bin/bash
sleep 10
echo "I'm here" > /root/.background
while [ ! -f /root/.kube/config ]
do
  sleep 1
done
echo done > /root/.k8s-start
if [ -f /root/.kube/start ]; then
  /root/.kube/start
fi

kubectl wait --timeout=600s --for=condition=Ready nodes/master > /root/.k8s-ready

snap install --classic go