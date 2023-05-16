#!/bin/bash

# build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o artifacts/simple-image/kube-sample-apiserver
echo "build kube-sample-apiserver end"

# build image
hubAddr="gymai/kube-sample-apiserver"

hookBranch=`git symbolic-ref --short -q HEAD | sed 's/\//-/g'`
hookRevision=`git rev-parse --short HEAD`

date=$(date +%Y%m%d%H%M)
branchTarget="${hubAddr}:${hookBranch}"
revisionTarget="${hubAddr}:${hookBranch}-${hookRevision}-${date}"

echo "branchTarget: ${branchTarget}, revisionTarget: ${revisionTarget}"

docker build -t ${branchTarget} -f ./artifacts/simple-image/Dockerfile .

#docker tag ${branchTarget} ${revisionTarget}

docker push ${branchTarget}
#docker push ${revisionTarget}

docker rmi ${branchTarget}
#docker rmi ${revisionTarget}

