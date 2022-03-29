FROM            scratch

ADD             "work/dracut/rootfs.tar" "/"

CMD             ["/usr/local/bin/girie"]
