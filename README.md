<h1>Kontext:</h1>

Every "kubectl" command needs a kubeconfig file to know which cluster to go to.  The default kubeconfig is: "$HOME/.kube/config".  Alternatively some folks uses multiple kubeconfig files if the cluster is very small and/or if they have very good memory!

I don't have good memory.  So, **Kontext** helps me manage the default kubeconfig file better.  The use-cases are listed below...

BTW, to find out more about kubeconfig files, please read [this documentation](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/).


_**Use case 1:**_ &nbsp; &nbsp; So, you just copied a "kubeconfig" yaml for a cluster to your clipboard.  Now you want to paste that context to your "~/.kube/config" file.

    kontext -p 

    -p is for paste.  Kontext will backup existing config file and will merge the new-config.yml from clip-board to ~/.kube/config


_**Use case 2:**_ &nbsp; &nbsp; So, your organization added a new cluster, and you were given the "kubeconfig" yaml for that cluster.  Now you want to add the context to your "~/.kube/config" file.

    kontext -a new-config.yml

    -a is for add.  Kontext will backup existing config file and will merge new-config.yml into ~/.kube/config


_**Use case 1:**_ &nbsp; &nbsp; So, you joined a new company, and they have several clusters managed by Rancher.  You collected the "kubeconfig" files from each cluster and saved them in a directory called "all-configs".  Now you want to merge them into your ~/.kube/config file.

    kontext -d ./all-configs/

    -d is for directory.  Kontext will backup the existing config file and add all the ".yaml" or ".yml" files found in "all-configs" folder and merges them into ~/.kube/config


_**Use case 4:**_ &nbsp; &nbsp; So, you were playing in AWS EKS and no longer need your "my-test-eks" cluster.  And would like to remove that context from "~/.kube/config".

    kontext -f my-test-eks
    
    -r is for remove.  Kontext will backup the existing config file and will remove the "cluster" references.  Not sure if there is a kubectl command for this.


_**Use case 5:**_ &nbsp; &nbsp; So, you are working on multiple clusters and want to switch to "staging-b".

    kontext 

    View the contexts.  The green one is the current one.  Use arrow keys to switch.  This essentially replaces to commands:
    kubectl config current-context 
    kubectl config use-context my-cluster-name


_**Use case 6:**_ &nbsp; &nbsp; So, you want o work on "kong" name-space for next 1 hour.

    kontext -n kong

    -n kong: sets the default name-space to kong.

    kontext -n -

    "-n -", a 'dash' sets the name-space back to default.  These commands replaces this kubectl command:
    kubectl config set-context --current --namespace=kong


<h3> Installation </h3>

You have the go-lang binary:

    go install github.com/ihawke/kontext

You don't have go binary and want the executable (coming soon):

    curl -O github.com/ihawke/kontext/binary/kontext_os_arch

