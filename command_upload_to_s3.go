package main

import (
	"fmt"
	"log"
	"time"
)

func uploadArtifactsToS3(params Params) {

	branchCheck := BranchCheck{Params: params}

	if branchCheck.shouldExecuteUploadToS3() {
		start := time.Now()
		log.Println("Upload to s3 start...")

		commit := params.Git.Commit
		aws := params.Config.AWS
		bucket := aws.ActifactBucket

		InitAWSSession(aws.Region)

		for _, component := range components {
			localFile := fmt.Sprintf("./shootr-%s/target/shootr-%s.jar", component, component)
			s3File := fmt.Sprintf("artifacts/shootr-%s.jar.%s", component, commit)
			uploadToS3(bucket, localFile, s3File)
		}

		log.Printf("Upload to s3 done in %s", time.Since(start))

	} else {
		log.Println("Skipping Upload to s3")
	}
}