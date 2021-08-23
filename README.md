&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;
![](./image/kontext.gif?raw=true "")

<h1>Kontext:</h1>

Every "kubectl" command needs a kubeconfig file to know which cluster to execute to command on.  The default kubeconfig is: "$HOME/.kube/config".  Alternatively some folks uses multiple kubeconfig files if the cluster is very small and/or if they have very good memory!

I don't have good memory.  So, **Kontext** helps me manage the default kubeconfig file better.  The use-cases are listed below...

BTW, to find out more about kubeconfig files, please read [this documentation](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/).


_**Use case 1:**_ &nbsp; You have a new cluster that you want to explore.  And you just copied a "kubeconfig" yaml for that new cluster from rancher or somewhere to your clipboard.  Now you want to merge that context to your "~/.kube/config" file.

    kontext -p 

    -p is for paste. Kontext will backup existing config file and will append the new-config.yml from clip-board to ~/.kube/config

    This essentially replaces these commands:
    - pbpaste > ~/.kube/new-file.yml
    - cp ~/.kube/config ~/.kube/config.YYYY-MM-DDTHH:MI:SS
    - KUBECONFIG=~/.kube/config:~/.kube/new-file.yml kubectl config view --flatten > ~/.kube/config


_**Use case 2:**_ &nbsp; So, your organization added a new cluster, and you were given the "kubeconfig" yaml for that cluster.  Now you want to add the context to your "~/.kube/config" file.

    kontext -f new-config.yml

    -f is for file. Kontext will backup existing config file and will append new-config.yml into ~/.kube/config

    This essentially replaces these commands:
    - cp ~/.kube/config ~/.kube/config.YYYY-MM-DDTHH:MI:SS
    - KUBECONFIG=~/.kube/config:~/.kube/new-config.yml kubectl config view --flatten > ~/.kube/config


_**Use case 3:**_ &nbsp; So, you joined a new company, and they have several clusters managed by Rancher.  You collected the "kubeconfig" files from each cluster and saved them in a directory called "all-configs".  Now you want to merge them into your ~/.kube/config file.

    kontext -d ./all-configs/

    -d is for directory. Kontext will backup the existing config file and add all the ".yaml" or ".yml" files found in 
    "all-configs" folder and merges them into ~/.kube/config


_**Use case 4:**_ &nbsp; So, you were playing in AWS EKS and no longer need your "my-test-eks" cluster.  And would like to remove that context from "~/.kube/config".

    kontext -r my-test-eks
    
    -r is for remove. Kontext will backup the existing config file and will remove the "cluster" references.  
    Not sure if there is a comparable kubectl command for this.


_**Use case 5:**_ &nbsp; So, you are working on multiple clusters and want to switch to "staging-b".

    kontext 

    View the contexts. The green one is the current context.  Use arrow keys to switch. Or you can do a text search as well.
    This essentially replaces two kubectl commands:

    - kubectl config current-context 
    - kubectl config use-context my-cluster-name


_**Use case 6:**_ &nbsp; So, you want o work on "kong" name-space for next 1 hour.

    kontext -n kong

    -n kong: sets the default name-space to kong.

    kontext -n -

    "-n -", a 'dash' sets the name-space back to default.  These commands replaces this kubectl command:

    kubectl config set-context --current --namespace=kong


<h3>Missing config file error:</h3>

Let us say your kubeconfig file is: "/etc/config" instead of "$USER_HOME/.kube/config" then set the "KONTEXT_FILE" env var like this:

    export KONTEXT_FILE=/etc/config

And then run "kontext" so kontext knows what file to work on.

<h3> Installation: </h3>

You have the go-lang binary:

    go install github.com/ihawke/kontext

You don't have go-lang binary and want the executable:

    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-OS-ARCH

    Where OS-ARCH in:  darwin-amd64, darwin-arm64, windows-386.exe, windows-amd64.exe, linux-386, linux-amd64, linux-arm64, openbsd-386, openbsd-amd64

    So, for most Mac amd the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-darwin-amd64

    So, for Mac Applie Silicon the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-darwin-arm64

    For most linux amd the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-linux-amd64

    For most windows amd the command will be (please add .exe after downloading):
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-windows-amd64

<h3> Issues/Features: </h3>

Please open an issue or a feature request here:  https://github.com/iHawke/kontext/issues

I'll be happy to look into them.
