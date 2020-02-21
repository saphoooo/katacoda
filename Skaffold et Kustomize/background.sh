while [ ! -f /root/.kube/config ]
do
  sleep 1
done
echo done > /root/k8s-start
if [ -f /root/.kube/start ]; then
  /root/.kube/start
fi

snap install --classic go