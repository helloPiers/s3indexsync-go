# s3indexsync

`go get hpgo.io/s3indexsync`

This is an S3 sync tool that uploads files called ".../index.html" additionally
to S3 objects with key ".../", where "..." is some prefix.

That is: it creates S3 objects whose key *ends with a slash*.

This is a hack which lets non-website S3 buckets with index.html pages work as
expected when behind CloudFront.

Notes: 
 - the S3 console doesn't understand this (it breaks the pretence that keys are hierarchical)
 - index.html files are uploaded under the key ".../index.html" as well as ".../"
 - for the very top you'll need to set index.html as the default object in the CF distro

Todo:
 - Any form of cleverness in the sync - currently every file is re-uploaded
