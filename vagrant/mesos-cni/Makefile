.PHONY : mesos-cni-demo mesos-cni-destroy

mesos-cni-demo: mesos-cni-destroy
	vagrant up && vagrant ssh mesos-node1 -c 'cd /opt/gopath/src/github.com/haoyixin/netplugin && make host-restart'

mesos-cni-destroy:
	vagrant destroy -f
