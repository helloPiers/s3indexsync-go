# s3indexsync

`go get hpgo.io/s3indexsync`

This is an s3 sync tool that uploads files called ".../index.html" additionally
to S3 objects with key ".../", where "..." is some prefix.

This is a slight hack, but it lets non-website S3 buckets work well behind
CloudFront.

Notes: 
 - the S3 console doesn't understand this (it breaks the pretence that keys are heirarchical)
 - index.html files are uploaded under the key ".../index.html" as well as ".../"
 - for the very top you'll need to set index.html as the default object in the CF distro
